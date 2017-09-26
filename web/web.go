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

package web

import (
	"fmt"
	"net/http"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"github.com/urfave/cli"
	"gopkg.in/macaron.v1"

	"xsec-evil-ips/web/routers"
	"xsec-evil-ips/models"
	"xsec-evil-ips/logger"
)

func RunWeb(ctx *cli.Context) (err error) {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner())
	m.Use(csrf.Csrfer())
	m.Use(cache.Cacher())

	m.Get("/", routers.Index)
	m.Get("/api/ip/:ip", routers.CheckIp)
	m.Post("/api/ip/", routers.UpdateIp)

	m.Get("/api/domain/:domain", routers.CheckDomain)
	m.Post("/api/domain/", routers.UpdateDomain)

	logger.Logger.Infof("run server on %v", fmt.Sprintf("%v:%v", HTTP_HOST, HTTP_PORT))
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", HTTP_HOST, HTTP_PORT), m)

	return err
}

func LoadFromFile(ctx *cli.Context) (err error) {
	models.Status()
	models.CACHE_IPS.LoadFile("ips")
	models.CACHE_DNS.LoadFile("dns")
	models.Status()
	models.SaveToDB()
	RunWeb(ctx)
	return err
}
