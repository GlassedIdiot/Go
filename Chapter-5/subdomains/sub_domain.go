package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/miekg/dns"
)

type Result struct {
	HostName  string `json:"hostname"`
	IPaddress string `json:"ipadderess"`
}

func A_Name(fqdn, dns_server string) ([]string, error) {
	var dns_msg dns.Msg
	var ip_addr []string

	dns_msg.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
	response, err := dns.Exchange(&dns_msg, "1.1.1.1:53")
	if err != nil {
		return ip_addr, err
	}

	for _, answer := range response.Answer {
		if a, ok := answer.(*dns.A); ok {
			ip_addr = append(ip_addr, a.A.String())
		}
	}
	return ip_addr, nil
}

func CName(fqdn, dns_server string) ([]string, error) {
	var dns_msg dns.Msg
	var ip_addr []string

	dns_msg.SetQuestion(dns.Fqdn(fqdn), dns.TypeCNAME)
	response, err := dns.Exchange(&dns_msg, "1.1.1.1:53")
	if err != nil {
		return ip_addr, err
	}

	for _, answer := range response.Answer {
		if a, ok := answer.(*dns.CNAME); ok {
			ip_addr = append(ip_addr, a.Target)
		}
	}
	return ip_addr, nil
}

func main() {
	var (
		flDomain   = flag.String("domain", "", "The domain to perform guessing against.")
		flWordlist = flag.String("wordlist", "", "The wordlist to use for guessing.")
		// flWorkerCount = flag.Int("c", 100, "The amount of workers to use.")
		// flServerAddr  = flag.String("server", "8.8.8.8:53", "The DNS server to use.")
	)
	flag.Parse()

	if *flDomain == "" || *flWordlist == "" {
		fmt.Print("Domain and wordlist are required.Provied them.")
		os.Exit(1)
	}
}
