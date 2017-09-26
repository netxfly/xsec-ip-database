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

package feeds

import (
	"xsec-evil-ips/models"
	"xsec-evil-ips/util"

	"io/ioutil"
	"strings"
)

func FetchFromDataplane() (evilIps models.EvilIps, err error) {
	urls := []string{
		"https://dataplane.org/dnsrd.txt",
		"https://dataplane.org/dnsrdany.txt",
		"https://dataplane.org/dnsversion.txt",
		"https://dataplane.org/sipinvitation.txt",
		"https://dataplane.org/sipquery.txt",
		"https://dataplane.org/sipregistration.txt",
		"https://dataplane.org/sshclient.txt",
		"https://dataplane.org/sshpwauth.txt",
		"https://dataplane.org/vncrfb.txt",
	}

	src := "dataplane.org"
	desc := "known attacker"
	check := "|"

	evilIps.Src.Source = src
	evilIps.Src.Desc = desc

	for _, url := range urls {
		resp, err := util.GetPage(url)
		if err == nil {
			ret, err := ioutil.ReadAll(resp)
			if err == nil {
				lines := strings.Split(string(ret), "\n")
				for _, line := range lines {
					if strings.HasPrefix(line, "#") || !strings.Contains(line, ".") {
						continue
					}
					tmp := strings.Split(line, check)
					if len(tmp) == 5 {
						ip := strings.TrimSpace(tmp[2])
						evilIps.Ips = append(evilIps.Ips, ip)
					}
				}
			}
		}
	}
	return evilIps, err
}
