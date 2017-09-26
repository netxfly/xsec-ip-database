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
	"xsec-evil-ips/settings"
	"xsec-evil-ips/logger"

	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"github.com/patrickmn/go-cache"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"path/filepath"
	"fmt"
)

var (
	DATA_TYPE string
	DATA_NAME string
	DATA_HOST string
	DATA_PORT int
	USERNAME  string
	PASSWORD  string
	SSL_MODE  string

	Engine               *xorm.Engine
	CACHE_IPS, CACHE_DNS *cache.Cache
)

func init() {
	cfg := settings.Cfg
	sec := cfg.Section("DATABASE")

	DATA_TYPE = sec.Key("DATA_TYPE").MustString("sqlite")
	DATA_NAME = sec.Key("DATA_NAME").MustString("data")
	DATA_HOST = sec.Key("DATA_HOST").MustString("DATA_HOST")
	DATA_PORT = sec.Key("DATA_PORT").MustInt(3306)
	USERNAME = sec.Key("USERNAME").MustString("USERNAME")
	PASSWORD = sec.Key("PASSWORD").MustString("PASSWORD")
	SSL_MODE = sec.Key("SSL_MODE").MustString("disable")

	err := NewDbEngine()
	if err == nil {
		Engine.Sync2(new(IpList))
		Engine.Sync2(new(DomainList))
	}

	CACHE_IPS = cache.New(cache.NoExpiration, cache.DefaultExpiration)
	CACHE_DNS = cache.New(cache.NoExpiration, cache.DefaultExpiration)

}

// init a database instance
func NewDbEngine() (err error) {
	switch DATA_TYPE {
	case "sqlite":
		cur, _ := filepath.Abs(".")
		dataSourceName := fmt.Sprintf("%v/%v/%v.db", cur, "data", DATA_NAME)
		logger.Logger.Infof("sqlite db: %v", dataSourceName)
		Engine, err = xorm.NewEngine("sqlite3", dataSourceName)
		Engine.Logger().SetLevel(core.LOG_OFF)
		err = Engine.Ping()

	case "mysql":
		dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
			USERNAME, PASSWORD, DATA_HOST, DATA_PORT, DATA_NAME)

		Engine, err = xorm.NewEngine("mysql", dataSourceName)
		err = Engine.Ping()
	case "postgres":
		dbSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", USERNAME, PASSWORD, DATA_HOST,
			DATA_PORT, DATA_NAME, SSL_MODE)
		Engine, err = xorm.NewEngine("postgres", dbSourceName)
		err = Engine.Ping()

	default:
		cur, _ := filepath.Abs(".")
		dataSourceName := fmt.Sprintf("%v/%v/%v.db", cur, "data", DATA_NAME)
		logger.Logger.Infof("sqlite db: %v", dataSourceName)
		Engine, err = xorm.NewEngine("sqlite3", dataSourceName)
		err = Engine.Ping()
	}

	return err
}
