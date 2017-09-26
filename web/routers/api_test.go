/*

Copyright (c) 2017 xsec.io

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THEq
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

*/

package routers_test

import (
	"testing"
	"time"
	"net/http"
	"net/url"

	"xsec-evil-ips/util"
)

func TestUpdateIp(t *testing.T) {
	u := "http://127.0.0.1:8000/api/ip/"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	k := "aadcbfbc837757a9a24ac96cf9171c8b"
	ip := "212.129.58.111"
	pro := "xsec test pro"

	t.Log(http.PostForm(u, url.Values{"timestamp": {timestamp}, "secureKey": {util.MakeSign(timestamp, k)}, "ip": {ip}, "pro": {pro}}))
}

func TestUpdateDomain(t *testing.T) {
	u := "http://127.0.0.1:8000/api/domain/"
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	k := "aadcbfbc837757a9a24ac96cf9171c8b"
	domain := "www.hosting2balooonba.com"
	pro := "xsec test pro"

	t.Log(http.PostForm(u, url.Values{"timestamp": {timestamp}, "secureKey": {util.MakeSign(timestamp, k)}, "domain": {domain}, "pro": {pro}}))
}