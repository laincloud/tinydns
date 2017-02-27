package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/laincloud/tinydns/dns"
	"github.com/mijia/sweb/log"
)

const (
	version = "2.0.1"
)

func main() {
	var debug, printVersion bool
	flag.BoolVar(&debug, "debug", false, "Set true to open debug logging")
	flag.BoolVar(&printVersion, "v", false, "Print version of tinydnsd")
	flag.Parse()
	if printVersion == true {
		fmt.Println("networkd version: " + version)
		fmt.Println("Go version: " + runtime.Version())
		fmt.Println("Go OS/Arch: " + runtime.GOOS + "/" + runtime.GOARCH)
		os.Exit(0)
	}
	if debug {
		log.EnableDebug()
	}
	log.Infof("DNS-Creator starts. version: %s", version)
	notify := make(chan interface{})
	coreInfoWatcher := dns.NewCoreInfoWatcher(notify)
	dependsWatcher := dns.NewDependsWatcher(notify)
	fqdnsWatcher := dns.NewFqdnsWatcher(notify)
	dnsCreator := dns.NewCreator(notify)
	go coreInfoWatcher.Watch()
	go dependsWatcher.Watch()
	go fqdnsWatcher.Watch()
	dnsCreator.CreateDNS()
}
