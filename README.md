
# xsec ip database

xsec ip database为一个恶意IP和域名库(Malicious ip database)，它获取恶意IP和域名的方式有以下几种：

1. 通过爬虫定期拉取网络中公开的恶意ip库（可能过增加新爬虫的方式订阅新的IP库）
1. 支持与自有的其他安全产品联动（HIDS、WAF、蜜罐、防火墙等产品），实时更新IP库

## 功能说明

1. 启动后会定期更新ip库，默认为1小时更新一次
1. 支持将恶意ip信息写入postgres, sqlite, mysql, mongodb数据库
1. 支持恶意ip信息导出、导入
1. 提供了恶意ip和dns检测以及与其他安全产品联动的接口

## 使用方法

```bash
$ ./main 
[xorm] [info]  2017/09/26 13:22:58.220496 PING DATABASE mysql
NAME:
   xsec Malicious ip database - A Malicious ip database

USAGE:
   main [global options] command [command options] [arguments...]
   
VERSION:
   20170925
   
AUTHOR(S):
   netxfly <x@xsec.io> 
   
COMMANDS:
     serve    startup evil ips
     dump     Fetch all evil ips info and save to file
     load     load ips from file
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version


```
- serve会启动程序，抓取完恶意ip和域名信息后会启动web接口
- dump，将恶意ip和域名导出到当前目录，文件名分别为ips和dns
- load，将ips和dns中的信息导入内存并启动WEB接口

### 运行截图

- 直接启动程序

![](https://docs.xsec.io/images/evil_ips/serve.png)

- 导出恶意ip信息到当前目录中，使用场景为部分URL是被墙了的，需要先在国外的VPS中导出文件拖回国内使用

![](https://docs.xsec.io/images/evil_ips/dump.png)

- 导入恶意ip信息并启动WEB接口
![](https://docs.xsec.io/images/evil_ips/load.png)

- 恶意IP检测及实时提交测试
![](https://docs.xsec.io/images/evil_ips/api_ip.png)

- 恶意域名检测及提交测试
![](https://docs.xsec.io/images/evil_ips/api_dns.png)

其中测试与其他安全产品联动的测试代码的内容如下：

```go

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
```

提交的参数需要有以下几个参数，而且安全产品的key必须与恶意IP库的相同，否则不会接受提交的恶意ip或域名信息。

- timestamp
- secureKey
- ip/domain，表示恶意ip或域名
- pro，表示需要联动的安全产品名称

### Demo

1. 恶意IP检测，[http://xsec.io:8000/api/ip/212.129.58.111](http://xsec.io:8000/api/ip/212.129.58.111)
1. 恶意域名检测，[http://xsec.io:8000/api/domain/www.hosting2balooonba.com](http://xsec.io:8000/api/domain/www.hosting2balooonba.com)

项目地址：https://github.com/netxfly/xsec-ip-database

## 更新记录

### 2017/9/28

- 恶意域名的种子中新增了360 netlab提供的DGA，使得域名记录直接上到了百万级。

![](https://docs.xsec.io/images/evil_ips/netlab_360.png)
![](https://docs.xsec.io/images/evil_ips/netlab_360_check.png)

- 因为data.netlab.360.com在国内，而且体积在70M以上，所以从vps中的拉取速度很慢，建议下载到本地，将`feeds/netlab360.go`中的URL改为本地地址。

```go
url := "http://data.netlab.360.com/feeds/dga/dga.txt"
// url := "http://127.0.0.1:8000/dga.txt"
```

- 如果vps内存不足，会在将恶意IP和域名导出到文件中时报错，解决方案为增加swap分区。
![](https://docs.xsec.io/images/evil_ips/swap.png)