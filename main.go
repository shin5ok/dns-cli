package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shin5ok/dnscli/internal/clouddns"
)

var defaultConfig = map[string]string{
	"domain": os.Getenv("DNS_DOMAIN"),
	"ttl":    os.Getenv("DNS_DEFAULT_TTL"),
}

func main() {
	data := flag.String("data", "", "")
	key := flag.String("key", "", "")
	domain := flag.String("domain", defaultConfig["domain"], `Do "export DNS_DOMAIN=<your domain>" to set default managed domain`)
	ttl := flag.Int64("ttl", 60, "")
	flag.Parse()
	rr := clouddns.Record{
		RType: "A",
		RData: *data,
		RKey:  *key,
		TTL:   int(*ttl),
	}

	dnsRr := clouddns.Rr{
		Domain: *domain,
	}
	err := dnsRr.Set(&rr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Registered: %+v\n", rr)
}
