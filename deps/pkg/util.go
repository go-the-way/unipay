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
	"fmt"
	"strconv"
	"time"
)

type number interface{ uint | byte }

func IfFunc(ok bool, fn func()) {
	if ok {
		fn()
	}
}

func IfGt0Func[T number](n T, fn func())   { IfFunc(n > 0, fn) }
func IfNotEmptyFunc(str string, fn func()) { IfFunc(str != "", fn) }
func TimeNow() time.Time                   { return time.Now() }
func TimeNowStr() string                   { return TimeNow().Format("2006-01-02 15:04:05") }
func TimeNowNumStr() string                { return TimeNow().Format("20060102150405") }
func TimeNowStamp() string                 { return fmt.Sprintf("%d", TimeNow().Unix()) }
func TimeNowStampLong() string             { return fmt.Sprintf("%d", TimeNow().UnixMilli()) }
func ParseTime(str string) (t time.Time) {
	t, _ = time.Parse("2006-01-02 15:04:05", str)
	return
}
func FormatTime(t time.Time) (str string) {
	return t.Format("2006-01-02 15:04:05")
}
func FromUnix(unixStr string) time.Time {
	unix, _ := strconv.ParseInt(unixStr, 10, 64)
	if len(unixStr) == 10 {
		return time.UnixMilli(unix)
	}
	return time.UnixMicro(unix)
}
func GetTimeMap() map[string]any {
	// 时间变量 => Time.
	// 当前时间`2006-01-02 15:04:05` => NowTime
	// 当前时间`20060102150405` => NowTimeNum
	// 当前时间戳`1705976043` => NowTimestamp
	// 当前时间戳`1705976043000` => NowTimestampLong
	return map[string]any{
		"NowTime":          TimeNowStr(),
		"NowTimeNum":       TimeNowNumStr(),
		"NowTimestamp":     TimeNowStamp(),
		"NowTimestampLong": TimeNowStampLong(),
	}
}
