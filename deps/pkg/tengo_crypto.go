// Copyright 2024 unipay Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pkg

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"strings"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"

	_ "crypto/md5"
)

var cryptoModule = make(map[string]tengo.Object,
	len(hashNames)*4+ // # of hashes
		1*5*2+ // of block ciphers
		3, // of utilities
)

func init() {
	ReloadCryptoAlgorithms()
	LoadCustomCryptoAlgorithms()
	registerBlockCipher(&blockCiph{
		MName:      "aes",
		MBlockSize: aes.BlockSize,
		MKeySizes:  []int{128 / 8, 192 / 8, 256 / 8},
		NewBlock: func(key []byte) cipher.Block {
			c, err := aes.NewCipher(key)
			if err != nil {
				// key sizes are checked beforehand
				panic(err)
			}
			return c
		},
	})
	cryptoModule["pad_pkcs7"] = &tengo.UserFunction{Name: "pad_pkcs7", Value: padPKCS7}
	cryptoModule["unpad_pkcs7"] = &tengo.UserFunction{Name: "unpad_pkcs7", Value: unpadPKCS7}
	cryptoModule["rand_bytes"] = &tengo.UserFunction{Name: "rand_bytes", Value: randBytes}
	stdlib.BuiltinModules["crypto"] = cryptoModule
}

// TODO keyderiv
// TODO asymmetric ciphers

var ErrMalformedPadding = &tengo.Error{Value: &tengo.String{Value: fmt.Sprintf("malformed padding")}}
var ErrKeySize = errors.New("invalid key size")
var ErrIVSize = errors.New("invalid iv size")
var ErrDataMultipleBlockSize = errors.New("data must be multiple of block size")

var hashNames = map[crypto.Hash]string{
	crypto.MD4:         "MD4",
	crypto.MD5:         "MD5",
	crypto.SHA1:        "SHA1",
	crypto.SHA224:      "SHA224",
	crypto.SHA256:      "SHA256",
	crypto.SHA384:      "SHA384",
	crypto.SHA512:      "SHA512",
	crypto.MD5SHA1:     "MD5SHA1",
	crypto.RIPEMD160:   "RIPEMD160",
	crypto.SHA3_224:    "SHA3_224",
	crypto.SHA3_256:    "SHA3_256",
	crypto.SHA3_384:    "SHA3_384",
	crypto.SHA3_512:    "SHA3_512",
	crypto.SHA512_224:  "SHA512_224",
	crypto.SHA512_256:  "SHA512_256",
	crypto.BLAKE2s_256: "BLAKE2s_256",
	crypto.BLAKE2b_256: "BLAKE2b_256",
	crypto.BLAKE2b_384: "BLAKE2b_384",
	crypto.BLAKE2b_512: "BLAKE2b_512",
}

func ReloadCryptoAlgorithms() {
	for h, n := range hashNames {
		if !h.Available() {
			continue
		}

		n = strings.ToLower(n)

		if _, ok := cryptoModule[n]; ok {
			continue
		}

		registerHash(n, h.New)
	}
}

func LoadCustomCryptoAlgorithms() {
	cryptoModule["sha256WithRSA"] = &tengo.UserFunction{Name: "sha256WithRSA", Value: newSHA256WithRSA}
}

func padPKCS7(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 2 {
		return nil, tengo.ErrWrongNumArguments
	}

	data, ok := tengo.ToByteSlice(args[0])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "data",
			Expected: "bytes",
			Found:    args[0].TypeName(),
		}
	}

	l, ok := tengo.ToInt(args[1])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "length",
			Expected: "int",
			Found:    args[1].TypeName(),
		}
	}

	if l <= 0 || l > 255 {
		return nil, tengo.ErrIndexOutOfBounds
	}

	padLen := l - (len(data) % l)
	return &tengo.Bytes{Value: append(data, bytes.Repeat([]byte{byte(padLen)}, padLen)...)}, nil
}

func unpadPKCS7(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 2 {
		return nil, tengo.ErrWrongNumArguments
	}

	data, ok := tengo.ToByteSlice(args[0])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "data",
			Expected: "bytes",
			Found:    args[0].TypeName(),
		}
	}

	l, ok := tengo.ToInt(args[1])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "length",
			Expected: "int",
			Found:    args[1].TypeName(),
		}
	}

	if l <= 0 || l > 255 {
		return nil, tengo.ErrIndexOutOfBounds
	}

	if len(data)%l != 0 {
		return ErrMalformedPadding, nil
	}

	padLen := int(data[len(data)-1])

	if padLen >= len(data) {
		return ErrMalformedPadding, nil
	}

	for _, el := range data[len(data)-padLen:] {
		if el != byte(padLen) {
			// recoverable error
			return ErrMalformedPadding, nil
		}
	}

	return &tengo.Bytes{Value: data[:len(data)-padLen]}, nil
}

func randBytes(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 1 {
		return nil, tengo.ErrWrongNumArguments
	}

	l, ok := tengo.ToInt(args[0])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{
			Name:     "length",
			Expected: "int",
			Found:    args[0].TypeName(),
		}
	}

	if l < 0 {
		return nil, tengo.ErrIndexOutOfBounds
	}

	bs := make([]byte, l)
	_, err := rand.Read(bs)
	if err != nil {
		return nil, err
	}

	return &tengo.Bytes{Value: bs}, nil
}

func registerBlockCipher(ciph *blockCiph) {
	registerCipherFuncs(ciph.Name()+"_cbc", &blockModeCiph{
		Cipher:       ciph,
		NewEncrypter: cipher.NewCBCEncrypter,
		NewDecrypter: cipher.NewCBCDecrypter,
	})
	registerCipherFuncs(ciph.Name()+"_ctr", &streamCiph{
		MName: ciph.Name(),
		NewEncStream: func(key, iv []byte) cipher.Stream {
			return cipher.NewCTR(ciph.NewBlock(key), iv)
		},
		MKeySizes: ciph.KeySizes(),
		MIVSize:   ciph.MBlockSize,
	})
	registerCipherFuncs(ciph.Name()+"_ofb", &streamCiph{
		MName: ciph.Name(),
		NewEncStream: func(key, iv []byte) cipher.Stream {
			return cipher.NewOFB(ciph.NewBlock(key), iv)
		},
		MKeySizes: ciph.KeySizes(),
		MIVSize:   ciph.MBlockSize,
	})
	if ciph.BlockSize() == 16 {
		registerAEAD(ciph.Name()+"_gcm", func(key []byte) cipher.AEAD {
			gcm, err := cipher.NewGCM(ciph.NewBlock(key))
			if err != nil {
				// Will never occur
				panic(err)
			}
			return gcm
		}, ciph.KeySizes())
	}
}

func registerCipherFuncs(n string, ciph cipherI) {
	cryptoModule["encrypt_"+n] = &tengo.UserFunction{Name: "encrypt_" + n, Value: newCrypterFunc(ciph, true)}
	cryptoModule["decrypt_"+n] = &tengo.UserFunction{Name: "decrypt_" + n, Value: newCrypterFunc(ciph, false)}
}

func registerAEAD(n string, newCipher func(key []byte) cipher.AEAD, keySizes []int) {
	cryptoModule["seal_"+n] = &tengo.UserFunction{Name: "seal_" + n, Value: newAEADCrypterFunc(newCipher, keySizes, true)}
	cryptoModule["open_"+n] = &tengo.UserFunction{Name: "open_" + n, Value: newAEADCrypterFunc(newCipher, keySizes, false)}
}

func newCrypterFunc(ciph cipherI, encrypter bool) tengo.CallableFunc {
	return func(args ...tengo.Object) (tengo.Object, error) {
		if ciph.IVSize() > 0 {
			if len(args) != 3 {
				return nil, tengo.ErrWrongNumArguments
			}
		} else {
			if len(args) != 2 {
				return nil, tengo.ErrWrongNumArguments
			}
		}

		data, ok := tengo.ToByteSlice(args[0])
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     "data",
				Expected: "bytes",
				Found:    args[0].TypeName(),
			}
		}

		if ciph.BlockSize() > 0 && len(data)%ciph.BlockSize() != 0 {
			return nil, ErrDataMultipleBlockSize
		}

		key, ok := tengo.ToByteSlice(args[1])
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     "key",
				Expected: "bytes",
				Found:    args[1].TypeName(),
			}
		}

		var iv []byte

		if ciph.IVSize() > 0 {
			iv, ok = tengo.ToByteSlice(args[2])
			if !ok {
				return nil, tengo.ErrInvalidArgumentType{
					Name:     "iv",
					Expected: "bytes",
					Found:    args[2].TypeName(),
				}
			}
			if len(iv) != ciph.IVSize() {
				return nil, ErrIVSize
			}
		}

		for _, l := range ciph.KeySizes() {
			if l == len(key) {
				if encrypter {
					ciph.Encrypt(data, key, iv)
				} else {
					ciph.Decrypt(data, key, iv)
				}

				return &tengo.Bytes{
					Value: data,
				}, nil
			}
		}

		// probably unrecoverable
		return nil, ErrKeySize
	}
}

func newAEADCrypterFunc(newCipher func(key []byte) cipher.AEAD, keySizes []int, encrypter bool) tengo.CallableFunc {
	return func(args ...tengo.Object) (ret tengo.Object, err error) {
		if len(args) < 3 || len(args) > 4 {
			return nil, tengo.ErrWrongNumArguments
		}

		data, ok := tengo.ToByteSlice(args[0])
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     "data",
				Expected: "bytes",
				Found:    args[0].TypeName(),
			}
		}

		key, ok := tengo.ToByteSlice(args[1])
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     "key",
				Expected: "bytes",
				Found:    args[1].TypeName(),
			}
		}

		var ciph cipher.AEAD

		for _, l := range keySizes {
			if l == len(key) {
				ciph = newCipher(key)
			}
		}

		if ciph == nil {
			return nil, ErrKeySize
		}

		iv, ok := tengo.ToByteSlice(args[2])
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     "iv",
				Expected: "bytes",
				Found:    args[2].TypeName(),
			}
		}
		if len(iv) != ciph.NonceSize() {
			return nil, ErrIVSize
		}

		var addData []byte
		if len(args) == 4 {
			addData, ok = tengo.ToByteSlice(args[3])
			if !ok {
				return nil, tengo.ErrInvalidArgumentType{
					Name:     "additional_data",
					Expected: "bytes",
					Found:    args[3].TypeName(),
				}
			}
		}

		if encrypter {
			data = ciph.Seal(data[:0], iv, data, addData)
			return &tengo.Bytes{Value: data}, nil
		} else {
			data, err = ciph.Open(data[:0], iv, data, addData)
			return &tengo.Bytes{Value: data}, err
		}

	}
}

func registerHash(n string, newHash func() hash.Hash) {
	cryptoModule[n] = &tengo.UserFunction{Name: n, Value: newHashFunc(newHash, false)}
	cryptoModule[n+"_hex"] = &tengo.UserFunction{Name: n + "_hex", Value: newHashFunc(newHash, true)} // See #216
	cryptoModule["hmac_"+n] = &tengo.UserFunction{Name: "hmac_" + n, Value: newHMACFunc(hmacByHash(newHash), false)}
	cryptoModule["hmac_"+n+"_hex"] = &tengo.UserFunction{Name: "hmac_" + n + "_hex", Value: newHMACFunc(hmacByHash(newHash), true)}
}

func hmacByHash(newHash func() hash.Hash) func(key []byte) hash.Hash {
	return func(key []byte) hash.Hash {
		return hmac.New(newHash, key)
	}
}

func newHashFunc(newHash func() hash.Hash, returnHex bool) tengo.CallableFunc {
	return func(args ...tengo.Object) (tengo.Object, error) {
		if len(args) != 1 {
			return nil, tengo.ErrWrongNumArguments
		}

		inp, ok := tengo.ToByteSlice(args[0])
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     "data",
				Expected: "bytes",
				Found:    args[0].TypeName(),
			}
		}

		h := newHash()
		h.Write(inp)

		out := make([]byte, 0, h.Size())

		out = h.Sum(out)

		if returnHex {
			return &tengo.String{
				Value: hex.EncodeToString(out),
			}, nil
		} else {
			return &tengo.Bytes{
				Value: out,
			}, nil
		}
	}
}

func newSHA256WithRSA(args ...tengo.Object) (tengo.Object, error) {
	if len(args) != 2 {
		return nil, tengo.ErrWrongNumArguments
	}
	// 待加密的字符串
	text, ok := tengo.ToString(args[0])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{Name: "text", Expected: "string", Found: args[0].TypeName()}
	}
	//私钥 private_key
	privateKey, ok := tengo.ToString(args[1])
	if !ok {
		return nil, tengo.ErrInvalidArgumentType{Name: "private_key", Expected: "string", Found: args[1].TypeName()}
	}
	priKeyPEM := `
-----BEGIN PRIVATE KEY-----
` + privateKey + `
-----END PRIVATE KEY-----
`
	block, _ := pem.Decode([]byte(priKeyPEM))
	if block == nil {
		return nil, errors.New("encrypt error: block is nil")
	}

	priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to parse RSA private key: %v", err))
	}

	h := sha256.New()
	h.Write([]byte(text))
	d := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), crypto.SHA256, d)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to RSA sign: %v", err))
	}

	signatureStr := base64.StdEncoding.EncodeToString(signature)

	return &tengo.String{Value: signatureStr}, nil
}

func formatPrivateKeyToPEM(privateKey string) string {
	// 移除空格或换行符
	privateKey = strings.ReplaceAll(privateKey, "\n", "")
	privateKey = strings.ReplaceAll(privateKey, "\r", "")

	// 每 64 个字符换行
	var formattedKey string
	for i := 0; i < len(privateKey); i += 64 {
		end := i + 64
		if end > len(privateKey) {
			end = len(privateKey)
		}
		formattedKey += privateKey[i:end] + "\n"
	}
	// 拼接 PEM 格式头尾
	return "-----BEGIN PRIVATE KEY-----\n" + formattedKey + "-----END PRIVATE KEY-----"
}

func newHMACFunc(newMac func(key []byte) hash.Hash, returnHex bool) tengo.CallableFunc {
	return func(args ...tengo.Object) (tengo.Object, error) {
		if len(args) != 2 {
			return nil, tengo.ErrWrongNumArguments
		}

		inp, ok := tengo.ToByteSlice(args[0])
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     "data",
				Expected: "bytes",
				Found:    args[0].TypeName(),
			}
		}

		key, ok := tengo.ToByteSlice(args[1])
		if !ok {
			return nil, tengo.ErrInvalidArgumentType{
				Name:     "key",
				Expected: "bytes",
				Found:    args[1].TypeName(),
			}
		}

		h := newMac(key)
		h.Write(inp)

		out := make([]byte, 0, h.Size())

		out = h.Sum(out)

		if returnHex {
			return &tengo.String{
				Value: hex.EncodeToString(out),
			}, nil
		} else {
			return &tengo.Bytes{
				Value: out,
			}, nil
		}
	}
}

type cipherI interface {
	Name() string

	IVSize() int
	BlockSize() int
	KeySizes() []int

	Encrypt(data, key, iv []byte)
	Decrypt(data, key, iv []byte)
}

type blockCiph struct {
	MName      string
	MBlockSize int
	MKeySizes  []int
	NewBlock   func(key []byte) cipher.Block
}

func (c *blockCiph) Name() string {
	return c.MName
}

func (c *blockCiph) IVSize() int {
	return -1
}

func (c *blockCiph) BlockSize() int {
	return c.MBlockSize
}

func (c *blockCiph) KeySizes() []int {
	return c.MKeySizes
}

func (c *blockCiph) Encrypt(data, key, iv []byte) {
	c.NewBlock(key).Encrypt(data, data)
}

func (c *blockCiph) Decrypt(data, key, iv []byte) {
	c.NewBlock(key).Decrypt(data, data)
}

type blockModeCiph struct {
	Cipher       cipherI
	NewEncrypter func(b cipher.Block, iv []byte) cipher.BlockMode
	NewDecrypter func(b cipher.Block, iv []byte) cipher.BlockMode
}

func (c *blockModeCiph) Name() string {
	return c.Cipher.Name()
}

func (c *blockModeCiph) IVSize() int {
	return c.BlockSize()
}

func (c *blockModeCiph) BlockSize() int {
	return c.Cipher.BlockSize()
}

func (c *blockModeCiph) KeySizes() []int {
	return c.Cipher.KeySizes()
}

func (c *blockModeCiph) Encrypt(data, key, iv []byte) {
	ciph := c.Cipher
	cbc := c.NewEncrypter(ciph.(*blockCiph).NewBlock(key), iv)
	cbc.CryptBlocks(data, data)
}

func (c *blockModeCiph) Decrypt(data, key, iv []byte) {
	ciph := c.Cipher
	cbc := c.NewDecrypter(ciph.(*blockCiph).NewBlock(key), iv)
	cbc.CryptBlocks(data, data)
}

type streamCiph struct {
	MName        string
	MIVSize      int
	MKeySizes    []int
	NewEncStream func(key, iv []byte) cipher.Stream
	NewDecStream func(key, iv []byte) cipher.Stream
}

func (c *streamCiph) Name() string {
	return c.MName
}

func (c *streamCiph) IVSize() int {
	return c.MIVSize
}

func (c *streamCiph) BlockSize() int {
	return -1
}

func (c *streamCiph) KeySizes() []int {
	return c.MKeySizes
}

func (c *streamCiph) Encrypt(data, key, iv []byte) {
	c.NewEncStream(key, iv).XORKeyStream(data, data)
}

func (c *streamCiph) Decrypt(data, key, iv []byte) {
	if c.NewDecStream == nil {
		c.NewEncStream(key, iv).XORKeyStream(data, data)
		return
	}
	c.NewDecStream(key, iv).XORKeyStream(data, data)
}
