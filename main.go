package main

import (
	"flag"
	"fmt"

	"github.com/shin5ok/dns-cli/internal/dns"
)

func main() {
	data := flag.String("data", "", "")
	key := flag.String("key", "", "")
	flag.Parse()
	rr := dns.Record{
		RType: "A",
		RData: data,
		RKey:  key,
		TTL:   60,
	}
	err := dns.Set(rr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Registered: %+v\n", rr)
}
