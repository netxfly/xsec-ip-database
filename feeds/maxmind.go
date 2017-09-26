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

	"github.com/PuerkitoBio/goquery"
)

func FetchFrommaxmind() (evilIps models.EvilIps, err error) {
	url := "https://www.maxmind.com/en/high-risk-ip-sample-list"
	src := "www.maxmind.com"
	desc := "bad reputation (suspicious)"

	evilIps.Src.Source = src
	evilIps.Src.Desc = desc

	resp, err := util.GetPage(url)
	if err == nil {
		doc, err := goquery.NewDocumentFromReader(resp)
		if err == nil {
			TRs := doc.Find("a.span3")
			TRs.Each(func(_ int, sec *goquery.Selection) {
				if len(sec.Nodes) > 0 {
					ip := sec.Nodes[0].FirstChild.Data
					evilIps.Ips = append(evilIps.Ips, ip)
				}
			})
		}
	}
	return evilIps, err
}
