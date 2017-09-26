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
	"xsec-evil-ips/logger"

	"sync"
	"time"
	"xsec-evil-ips/web"
	"github.com/urfave/cli"
)

type EvilIpFunc func() (evilIps models.EvilIps, err error)

type EvilDnsFunc func() (evilDns models.EvilDns, err error)

var (
	EvilIpFuncs  []EvilIpFunc
	EvilDnsFuncs []EvilDnsFunc
)

func Init() {
	// Evil ips func
	EvilIpFuncs = append(EvilIpFuncs, FetchFromAlienvault)
	EvilIpFuncs = append(EvilIpFuncs, FetchBadips)
	EvilIpFuncs = append(EvilIpFuncs, FetchBlocklist)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromBotscout)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromBruteforceblocker)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromCinsscore)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromCruzitWebAttacks)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromcyBersweat)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromDataplane)
	EvilIpFuncs = append(EvilIpFuncs, FetchIpsFromdShield)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromEmergingthreats)
	EvilIpFuncs = append(EvilIpFuncs, FetchIpsFromEmergingthreats)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromFeodotracker)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromGreensnow)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromMalwaredomainlist)
	EvilIpFuncs = append(EvilIpFuncs, FetchFrommaxmind)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromRutgers)
	EvilIpFuncs = append(EvilIpFuncs, FetchFromZeustracker)

	// Evil dns func
	EvilDnsFuncs = append(EvilDnsFuncs, FetchDnsFromBambenekconsulting)
	EvilDnsFuncs = append(EvilDnsFuncs, FetchFromCybercrime)
	EvilDnsFuncs = append(EvilDnsFuncs, FetchDomainsFromdShield)

}

func FetchEvilIps() {
	var wg sync.WaitGroup
	startTime := time.Now()
	wg.Add(len(EvilIpFuncs))
	for _, fn := range EvilIpFuncs {
		go func(fn EvilIpFunc) {
			models.SaveEvilIps(fn())
			wg.Done()
		}(fn)
	}
	wg.Wait()
	logger.Logger.Infof("Fetch Evil ips Done, used time: %v", time.Since(startTime))
}

func FetchEvilDns() {
	var wg sync.WaitGroup
	startTime := time.Now()
	wg.Add(len(EvilDnsFuncs))
	for _, fn := range EvilDnsFuncs {
		go func(fn EvilDnsFunc) {
			models.SaveEvilDns(fn())
			wg.Done()
		}(fn)
	}
	wg.Wait()
	logger.Logger.Infof("Fetch Evil Dns Done, used time: %v", time.Since(startTime))
}

func FetchAll(ctx *cli.Context) {

	for {
		go func(ctx *cli.Context) {
			FetchEvilDns()
			FetchEvilIps()
			models.Status()
			models.SaveToDB()
			models.SaveToFile(ctx)
		}(ctx)

		// update ip database interval, default 1 hour
		time.Sleep(60 * 60 * time.Second)
	}
}

func Startup(ctx *cli.Context) (err error) {
	Init()
	go FetchAll(ctx)
	web.RunWeb(ctx)
	return err
}

func Dump(ctx *cli.Context) (err error) {
	Init()
	FetchEvilDns()
	FetchEvilIps()
	models.Status()
	models.SaveToFile(ctx)
	return err
}
