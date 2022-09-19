package clouddns

import (
	"context"
	"log"

	"golang.org/x/oauth2/google"
	dns "google.golang.org/api/dns/v1"
)

type Record struct {
	RType  string
	RData  []string
	RKey   string
	TTL    int
	Status string
}

type Recorder interface {
	Get(string) (*Record, error)
	Set(*Record) error
}

type ZoneInfo struct {
	Domain      string
	ProjectId   string
	ManagedZone string
}

func (i *ZoneInfo) Get(key string) (*Record, error) {
	return &Record{}, nil
}

func (i *ZoneInfo) Set(r *Record) error {

	ctx := context.Background()
	c, err := google.DefaultClient(ctx, dns.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	dnsService, err := dns.New(c)
	if err != nil {
		log.Fatal(err)
	}
	recordSet := dns.ResourceRecordSet{
		Name:    r.RKey,
		Rrdatas: r.RData,
	}
	// log.Println(i.ProjectId, i.ManagedZone, r.RKey, r.RType, &recordSet)
	_, err = dnsService.ResourceRecordSets.Patch(i.ProjectId, i.ManagedZone, r.RKey, r.RType, &recordSet).Context(ctx).Do()
	if err != nil {
		return err
	}
	return nil
}
