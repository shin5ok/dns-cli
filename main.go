package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/shin5ok/dnscli/internal/clouddns"
)

var defaultConfig = map[string]string{
	"domain":  os.Getenv("DNS_DOMAIN"),
	"zone":    os.Getenv("DNS_ZONE"),
	"project": os.Getenv("GOOGLE_CLOUD_PROJECT"),
}

type config struct {
	Project string `json:"project"`
	Domain  string `json:"domain"`
	Zone    string `json:"zone"`
}

func envinfo(c config) {
	data, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

type DNSMain struct {
	Client clouddns.Recorder
}

func main() {
	data := flag.String("data", "", "")
	key := flag.String("key", "", "")
	domain := flag.String("domain", defaultConfig["domain"], `ig: example.com`)
	zone := flag.String("zone", defaultConfig["zone"], `ig: exapmple-com`)
	projectId := flag.String("project", defaultConfig["project"], "")
	ttl := flag.Int64("ttl", 60, "")
	env := flag.Bool("env", false, "")
	flag.Parse()

	if *env {
		c := config{
			Domain:  *domain,
			Project: *projectId,
			Zone:    *zone,
		}
		envinfo(c)
		os.Exit(0)
	}

	rr := clouddns.Record{
		RType: "A",
		RData: []string{*data},
		RKey:  *key,
		TTL:   int(*ttl),
	}

	zoneInfo := clouddns.ZoneInfo{
		Domain:      *domain,
		ProjectId:   *projectId,
		ManagedZone: *zone,
	}

	newRr, err := run(&zoneInfo, rr, *key)
	if err != nil {
		fmt.Println(err)
		return
	}

	jsonResult, _ := json.Marshal(&newRr)
	fmt.Printf("Updated: %v\n", string(jsonResult))

}

func run(zoneInfo clouddns.Recorder, rr clouddns.Record, key string) (clouddns.Record, error) {

	v := DNSMain{
		Client: zoneInfo,
	}

	_, err := v.Client.Get(key)
	if errors.Is(err, clouddns.ErrNotFound) {
		err := v.Client.Create(&rr)
		if err != nil {
			return clouddns.Record{}, err
		}
	} else {
		err := v.Client.Set(&rr)
		if err != nil {
			return clouddns.Record{}, err
		}
	}

	return rr, nil
}
