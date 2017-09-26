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

package cmd

import (
	"github.com/urfave/cli"

	"xsec-evil-ips/feeds"
	"xsec-evil-ips/web"
)

var Serve = cli.Command{
	Name:        "serve",
	Usage:       "startup evil ips",
	Description: "startup evil ips",
	Action:      feeds.Startup,
}

var RunWeb = cli.Command{
	Name:        "web",
	Usage:       "startup web interface",
	Description: "startup web interface",
	Action:      web.RunWeb,
}

var SaveFile = cli.Command{
	Name:        "dump",
	Usage:       "Fetch all evil ips info and save to file",
	Description: "Fetch all evil ips info and save to file",
	Action:      feeds.Dump,
}

var LoadFile = cli.Command{
	Name:        "load",
	Usage:       "load ips from file",
	Description: "load ips from file",
	Action:      web.LoadFromFile,
}
