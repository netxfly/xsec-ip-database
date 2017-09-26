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

// bad ip or dns source info
type Source struct {
	Desc   string
	Source string
}

// evil ips
type EvilIps struct {
	Ips []string
	Src Source
}

// evil dns
type EvilDns struct {
	Domains []string
	Src     Source
}

type IpList struct {
	Id   int64
	Ip   string `xorm:"ip"`
	Info []Source `xorm:"info"`
}

type DomainList struct {
	Id     int64
	Domain string `xorm:"domain"`
	Info   []Source `xorm:"info"`
}

func NewIpList(ip string, info []Source) (IpList) {
	infos := make([]Source, 0)
	infos = append(infos, info...)
	return IpList{Ip: ip, Info: infos}
}

func (i *IpList) IsExist() (has bool, err error) {
	var iplist IpList
	has, err = Engine.Table("ip_list").Where("ip=?", i.Ip).Get(&iplist)
	return has, err
}

func (i *IpList) Update() (err error) {
	var iplist IpList
	has, err := Engine.Table("ip_list").Where("ip=?", i.Ip).Get(&iplist)
	if err == nil && has {
		Id := iplist.Id
		iplist.Info = append(iplist.Info, i.Info[0])
		_, err = Engine.Table("ip_list").ID(Id).Update(&iplist)
	}
	return err
}

func (i *IpList) Insert() (err error) {
	_, err = Engine.Table("ip_list").Insert(i)
	return err
}

func NewDomainList(domain string, info []Source) (DomainList) {
	infos := make([]Source, 0)
	infos = append(infos, info...)
	return DomainList{Domain: domain, Info: infos}
}

func (d *DomainList) IsExist() (has bool, err error) {
	var domainList DomainList
	has, err = Engine.Table("domain_list").Where("domain=?", d.Domain).Get(&domainList)
	return has, err
}

func (d *DomainList) Update() (err error) {
	var domainList DomainList
	has, err := Engine.Table("domain_list").Where("domain=?", d.Domain).Get(&domainList)
	if err == nil && has {
		Id := domainList.Id
		domainList.Info = append(domainList.Info, d.Info[0])
		_, err = Engine.Table("domain_list").ID(Id).Update(&domainList)
	}
	return err
}

func (d *DomainList) Insert() (err error) {
	_, err = Engine.Table("domain_list").Insert(d)
	return err
}

func InsertIps2Db(ips []IpList) (int64, error) {
	return Engine.Table("ip_list").Insert(ips)
}

func InsertDomains2Db(domains []DomainList) (int64, error) {
	return Engine.Table("domain_list").Insert(domains)
}

func ClearDB() (err error) {
	_, err = Engine.Exec("delete from ip_list")
	_, err = Engine.Exec("delte from domain_list")
	return err
}
