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

package models

import (
	"github.com/patrickmn/go-cache"
	"github.com/urfave/cli"

	"xsec-evil-ips/logger"
	//"xsec-evil-ips/web"

	"encoding/gob"
	"github.com/toolkits/slice"
)

func init() {
	gob.Register(DomainList{})
	gob.Register(IpList{})
}

func CacheStatus(cache *cache.Cache) (count int, items map[string]cache.Item) {
	count = cache.ItemCount()
	items = cache.Items()
	return count, items
}

func Status() {
	{
		count, _ := CacheStatus(CACHE_IPS)
		logger.Logger.Infof("Evil Ips count:%v", count)
	}

	{
		count, _ := CacheStatus(CACHE_DNS)
		logger.Logger.Infof("Evil Dns count:%v", count)
	}
}

func SaveToFile(ctx *cli.Context) (err error) {
	CACHE_IPS.SaveFile("ips")
	CACHE_DNS.SaveFile("dns")
	return err
}

func SaveToDB() {
	ClearDB()

	Num := 5000
	domainList := make([]DomainList, 0)
	{
		n, items := CacheStatus(CACHE_DNS)
		for _, v := range items {
			d := v.Object.(DomainList)
			domainList = append(domainList, d)
		}

		if n%Num == 0 {
			batch := n / Num
			for i := 0; i < batch; i++ {
				domains := domainList[i*Num:(i+1)*Num]
				InsertDomains2Db(domains)
				//log.Println(ret, err, i*Num, (i+1)*Num)
			}
		} else {
			batch := n / Num
			for i := 0; i < batch; i++ {
				domains := domainList[i*Num:(i+1)*Num]
				InsertDomains2Db(domains)
			}
			InsertDomains2Db(domainList[batch*Num:n])
		}
	}

	ipList := make([]IpList, 0)
	{
		n, items := CacheStatus(CACHE_IPS)
		for _, v := range items {
			i := v.Object.(IpList)
			ipList = append(ipList, i)
		}

		if n%Num == 0 {
			batch := n / Num
			for i := 0; i < batch; i++ {
				ips := ipList[i*Num:(i+1)*Num]
				InsertIps2Db(ips)

			}
		} else {
			batch := n / Num
			for i := 0; i < batch; i++ {
				ips := ipList[i*Num:(i+1)*Num]
				InsertIps2Db(ips)

			}
			InsertIps2Db(ipList[batch*Num:n])
		}
	}
}

func SaveEvilDns(evilDns EvilDns, err error) {
	if err == nil {
		domains := evilDns.Domains
		src := evilDns.Src
		for _, d := range domains {
			infos := make([]Source, 0)
			infos = append(infos, src)
			domain := NewDomainList(d, infos)
			item, found := CACHE_DNS.Get(d)
			if found {
				v := item.(DomainList)
				infos := v.Info

				sliceSource := make([]string, 0)
				for _, s := range infos {
					sliceSource = append(sliceSource, s.Source)
				}
				if !slice.ContainsString(sliceSource, src.Source) {
					infos = append(infos, src)
				}
				domain = NewDomainList(d, infos)

			}

			CACHE_DNS.Set(d, domain, cache.NoExpiration)
		}
	}
}

func SaveEvilIps(evilIps EvilIps, err error) {
	if err == nil {
		ips := evilIps.Ips
		src := evilIps.Src
		for _, ip := range ips {
			infos := make([]Source, 0)
			infos = append(infos, src)
			ipList := NewIpList(ip, infos)
			item, found := CACHE_IPS.Get(ip)
			if found {
				v := item.(IpList)
				infos := v.Info

				sliceSource := make([]string, 0)
				for _, s := range infos {
					sliceSource = append(sliceSource, s.Source)
				}
				if !slice.ContainsString(sliceSource, src.Source) {
					infos = append(infos, src)
				}
				ipList = NewIpList(ip, infos)
			}
			CACHE_IPS.Set(ip, ipList, cache.NoExpiration)
		}
	}
}
