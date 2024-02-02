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

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	serverAddr = bindEnv("SERVER_ADDR", ":9988")
	appKey     = bindEnv("APP_KEY", "BmnXsm843uA9WjWh22CWIXbrASo")
	appSecret  = bindEnv("APP_SECRET", "Ne4WZgphE1GicyYgQAYn0ZqhwvA")
	domainUrl  = bindEnv("DOMAIN_URL", "")
)

func bindEnv(name, defVal string) (bindVal string) {
	bindVal = defVal
	if val := os.Getenv(name); val != "" {
		bindVal = val
	}
	return
}

func getPublicIp() string {
	resp, err := http.Get("http://ipinfo.io/ip")
	if err != nil {
		panic("获取公网ip错误：" + err.Error())
	}
	buf, _ := io.ReadAll(resp.Body)
	return strings.TrimSpace(string(buf))
}

func init() {
	if domainUrl == "" {
		domainUrl = fmt.Sprintf("http://%s%s", getPublicIp(), serverAddr)
	}
}
