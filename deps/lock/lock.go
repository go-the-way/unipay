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

package lock

import "sync"

var (
	m  = map[string]struct{}{}
	mu = &sync.RWMutex{}
)

func mInit()                              { m = map[string]struct{}{} }
func RLock()                              { mu.RLock() }
func RUnlock()                            { mu.RUnlock() }
func Lock()                               { mu.Lock() }
func Unlock()                             { mu.Unlock() }
func Set(key string)                      { m[key] = struct{}{} }
func SetWithLock(key string)              { Lock(); Set(key); Unlock() }
func Have(key string) (have bool)         { _, have = m[key]; return }
func HaveWithLock(key string) (have bool) { RLock(); have = Have(key); RUnlock(); return }
func Delete(key string)                   { delete(m, key) }
func DeleteWithLock(key string)           { Lock(); Delete(key); Unlock() }
func Clear()                              { mInit() }
func ClearWithLock()                      { Lock(); Clear(); Unlock() }
