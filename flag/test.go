package main

import (
	"flag"
	"fmt"
	"os"
)

//
var (
	h bool

	v, V bool
	t, T bool
	q    *bool

	s string
	p string
	c string
	g string

	n int64
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")
	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration,dump it and exit")

	//另一中绑定方式
	q = flag.Bool("q", false, "suppress non-error messages during configuration testing")
	//注意’signal‘,默认是 -S string 有了`signal`之后，变为-s signal
	flag.StringVar(&s, "s", "", "send `signal` to a master")
	flag.StringVar(&p, "p", "/usr/local/nginx", "set configuration `file`")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set confiruration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuraton file")

	flag.Int64Var(&n, "n", 0, "set `number`")
	//改变默认的usage
	flag.Usage = Usage
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
	}
	if len(s) != 0 {
		fmt.Println(s)
	}
	if n != 0 {
		fmt.Println(n)
	}
}

func Usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0 Usage:nginx [-hvVtTq] [-n number] [-s signal] [-c filename] [-p prefix] [-g directives]
Options:
`)
	flag.PrintDefaults()
}
