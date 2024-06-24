package main

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg

	// fqdn = FUlly Qualified domain name.
	fqdn := dns.Fqdn("github.com")

	// Here we set the question to send to the DNS server
	msg.SetQuestion(fqdn, dns.TypeA)

	// Sending the message. In a recursive mode.
	response, err := dns.Exchange(&msg, "1.1.1.1:53")
	if err != nil {
		log.Fatalf("%s", err)
	}
  
  
	fmt.Print(response.Answer)
}
