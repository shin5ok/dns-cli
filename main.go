package main

import (
	"flag"
	"fmt"

	dns "github.com/shin5ok/dnscli/internal/dns"
)

func main() {
	data := flag.String("data", "", "")
	key := flag.String("key", "", "")
	flag.Parse()
	rr := dns.Record{
		RType: "A",
		RData: *data,
		RKey:  *key,
		TTL:   60,
	}

	dnsRr := dns.Rr{}
	err := dnsRr.Set(&rr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Registered: %+v\n", rr)
}
