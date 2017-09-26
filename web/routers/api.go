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

package routers

import (
	"gopkg.in/macaron.v1"
	"github.com/patrickmn/go-cache"

	"strings"

	"xsec-evil-ips/models"
	"xsec-evil-ips/logger"
	"xsec-evil-ips/settings"
	"xsec-evil-ips/util"
	"github.com/toolkits/slice"
)

type IplistApi struct {
	Evil bool          `json:"evil"`
	Data models.IpList `json:"data"`
}

type DnsApi struct {
	Evil bool              `json:"evil"`
	Data models.DomainList `json:"data"`
}

func CheckIp(ctx *macaron.Context) {
	var ipApi IplistApi
	ip := strings.TrimSpace(ctx.Params("ip"))
	v, has := models.CACHE_IPS.Get(ip)
	data, ok := v.(models.IpList)

	if settings.DEBUG {
		logger.Logger.Infof("ip: %v, ret: %v", ip, data)
	}

	if ok {
		ipApi.Evil = has
		ipApi.Data = data
	}
	ctx.JSON(200, ipApi)
}

func CheckDomain(ctx *macaron.Context) {
	var domainApi DnsApi
	domain := strings.TrimSpace(ctx.Params("domain"))
	v, has := models.CACHE_DNS.Get(domain)
	data, ok := v.(models.DomainList)

	if settings.DEBUG {
		logger.Logger.Infof("ip: %v, ret: %v", domain, data)
	}

	if ok {
		domainApi.Evil = has
		domainApi.Data = data
	}
	ctx.JSON(200, domainApi)
}

func UpdateIp(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	timestamp := ctx.Req.Form.Get("timestamp")
	secureKey := ctx.Req.Form.Get("secureKey")
	ip := ctx.Req.Form.Get("ip")
	pro := ctx.Req.Form.Get("pro")
	if secureKey == util.MakeSign(timestamp, settings.SECRET) {
		var info []models.Source
		srcPro := models.Source{Desc: "real time attacker", Source: pro}
		info = append(info, srcPro)
		ipList := models.NewIpList(ip, info)
		item, found := models.CACHE_IPS.Get(ip)
		if found {
			v := item.(models.IpList)
			infos := v.Info

			sliceSource := make([]string, 0)
			for _, s := range infos {
				sliceSource = append(sliceSource, s.Source)
			}
			if !slice.ContainsString(sliceSource, srcPro.Source) {
				infos = append(infos, srcPro)
			}
			ipList = models.NewIpList(ip, infos)
		}
		models.CACHE_IPS.Set(ip, ipList, cache.NoExpiration)
		ips := make([]models.IpList, 0)
		ips = append(ips, ipList)
		models.InsertIps2Db(ips)
		ctx.JSON(200, ipList)
	} else {
		ctx.JSON(200, "error")
	}
}

func UpdateDomain(ctx *macaron.Context) {
	ctx.Req.ParseForm()
	timestamp := ctx.Req.Form.Get("timestamp")
	secureKey := ctx.Req.Form.Get("secureKey")
	domain := ctx.Req.Form.Get("domain")
	pro := ctx.Req.Form.Get("pro")
	if secureKey == util.MakeSign(timestamp, settings.SECRET) {
		var info []models.Source
		srcPro := models.Source{Desc: "real time attacker", Source: pro}
		info = append(info, srcPro)
		d := models.NewDomainList(domain, info)

		item, found := models.CACHE_DNS.Get(domain)
		if found {
			v := item.(models.DomainList)
			infos := v.Info

			sliceSource := make([]string, 0)
			for _, s := range infos {
				sliceSource = append(sliceSource, s.Source)
			}
			if !slice.ContainsString(sliceSource, srcPro.Source) {
				infos = append(infos, srcPro)
			}
			d = models.NewDomainList(domain, infos)
		}

		models.CACHE_DNS.Set(domain, d, cache.NoExpiration)
		domains := make([]models.DomainList, 0)
		domains = append(domains, d)
		models.InsertDomains2Db(domains)
		ctx.JSON(200, domains)
	} else {
		ctx.JSON(200, "error")
	}
}
