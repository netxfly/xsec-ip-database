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

func FetchDnsFromBambenekconsulting() (evilDns models.EvilDns, err error) {
	url := "http://osint.bambenekconsulting.com/feeds/c2-dommasterlist-high.txt"
	src := "bambenekconsulting.com"
	desc := "C&Cs domain"
	check := ",Domain used by"

	evilDns.Src.Source = src
	evilDns.Src.Desc = desc

	resp, err := util.GetPage(url)
	if err == nil {
		ret, err := ioutil.ReadAll(resp)
		if err == nil {
			lines := strings.Split(string(ret), "\n")
			for _, line := range lines {
				if strings.Contains(line, "#") {
					continue
				}
				tmp := strings.Split(line, check)
				if len(tmp) > 1 {
					evilDns.Domains = append(evilDns.Domains, tmp[0])
				}
			}
		}
	}
	return evilDns, err
}

func FetchDGADnsFromBambenekconsulting() (evilDns models.EvilDns, err error) {
	url := "http://osint.bambenekconsulting.com/feeds/dga-feed.txt"
	src := "bambenekconsulting.com"
	desc := "Domain feed of known DGA domains from -2 to +3 days"
	check := ",Domain used by"

	evilDns.Src.Source = src
	evilDns.Src.Desc = desc

	resp, err := util.GetPage(url)
	if err == nil {
		ret, err := ioutil.ReadAll(resp)
		if err == nil {
			lines := strings.Split(string(ret), "\n")
			for _, line := range lines {
				if strings.Contains(line, "#") {
					continue
				}
				tmp := strings.Split(line, check)
				if len(tmp) > 1 {
					evilDns.Domains = append(evilDns.Domains, tmp[0])
				}
			}
		}
	}
	return evilDns, err
}

func FetchIpFromBambenekconsulting() (evilIps models.EvilIps, err error) {
	url := "http://osint.bambenekconsulting.com/feeds/c2-ipmasterlist-high.txt"
	src := "bambenekconsulting.com"
	desc := "C&Cs IP"
	check := ",IP used by"

	evilIps.Src.Source = src
	evilIps.Src.Desc = desc

	resp, err := util.GetPage(url)
	if err == nil {
		ret, err := ioutil.ReadAll(resp)
		if err == nil {
			lines := strings.Split(string(ret), "\n")
			for _, line := range lines {
				if strings.Contains(line, "#") {
					continue
				}
				tmp := strings.Split(line, check)
				if len(tmp) > 1 {
					evilIps.Ips = append(evilIps.Ips, tmp[0])
				}
			}
		}
	}
	return evilIps, err
}
