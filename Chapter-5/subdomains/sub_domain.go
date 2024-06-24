package main

import (
	"flag"
	"fmt"
	"os"
)

type Result struct {
	HostName  string `json:"hostname"`
	IPaddress string `json:"ipadderess"`
}

func main() {
	var (
		flDomain      = flag.String("domain", "", "The domain to perform guessing against.")
		flWordlist    = flag.String("wordlist", "", "The wordlist to use for guessing.")
		flWorkerCount = flag.Int("c", 100, "The amount of workers to use.")
		flServerAddr  = flag.String("server", "8.8.8.8:53", "The DNS server to use.")
	)
	flag.Parse()

	if *flDomain == "" || *flWordlist == "" {
		fmt.Print("Domain and wordlist are required.Provied them.")
		os.Exit(1)
	}
}
