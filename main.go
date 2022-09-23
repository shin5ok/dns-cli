package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/shin5ok/dnscli/internal/clouddns"
)

var defaultConfig = map[string]string{
	"domain":  os.Getenv("DNS_DOMAIN"),
	"zone":    os.Getenv("DNS_ZONE"),
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

func main() {
	data := flag.String("data", "", "")
	key := flag.String("key", "", "")
	domain := flag.String("domain", defaultConfig["domain"], `ig: example.com`)
	zone := flag.String("zone", defaultConfig["zone"], `ig: exapmple-com`)
	projectId := flag.String("project", defaultConfig["project"], "")
	ttl := flag.Int64("ttl", 60, "")
	flag.Parse()

	rr := clouddns.Record{
		RType: "A",
		RData: []string{*data},
		RKey:  *key,
		TTL:   int(*ttl),
	}

	dnsRr := clouddns.ZoneInfo{
		Domain:      *domain,
		ProjectId:   *projectId,
		ManagedZone: *zone,
	}

	_, err := dnsRr.Get(*key)
	if errors.Is(err, clouddns.ErrNotFound) {
		err = dnsRr.Create(&rr)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		err = dnsRr.Set(&rr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("Registered: %#v\n", rr)
}
