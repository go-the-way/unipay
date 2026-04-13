// Copyright 2026 unipay Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package debug

import (
	"log"
	"os"
	"path/filepath"
)

var (
	debugLogger    = log.New(os.Stdout, "[unipay-debug-log] ", log.LstdFlags)
	debugLogEnable = os.Getenv("UNIPAY_DEBUG_LOG_ENABLE") == "T"
	debugLogDir    = os.Getenv("UNIPAY_DEBUG_LOG_DIR")
)

func Enabled() bool { return debugLogEnable }

func Store(filename, content string) {
	filepath := filepath.Join(debugLogDir, filename)
	if err := os.WriteFile(filepath, []byte(content), 0700); err != nil {
		debugLogger.Println("=======================>saved err:", filepath, err)
	} else {
		debugLogger.Println("=======================>saved ok:", filepath)
	}
}

func init() {
	if !debugLogEnable {
		return
	}
	if debugLogDir == "" {
		debugLogDir = "/tmp/unipay_debug_logs"
	}
	if !dirExists(debugLogDir) {
		if err := os.MkdirAll(debugLogDir, os.ModePerm); err != nil {
			debugLogger.Println("=======================>create save dir err:", err)
		}
	}
	debugLogger.Println("=======================>current save dir:", debugLogDir)
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return info.IsDir()
}
