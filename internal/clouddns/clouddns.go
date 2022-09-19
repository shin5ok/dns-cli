package clouddns

import (
	"context"
	"fmt"
	"log"

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

func (i *ZoneInfo) makeClient(ctx context.Context) *dns.Service {
	dnsService, err := dns.NewService(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return dnsService
}

func (i *ZoneInfo) Get(key string) (*Record, error) {

	ctx := context.Background()

	dnsService := i.makeClient(ctx)

	responseRecordSet, err := dnsService.ResourceRecordSets.Get(i.ProjectId, i.ManagedZone, key, "A").Context(ctx).Do()
	if err != nil {
		fmt.Printf("%v", err)
		return &Record{}, nil
	}
	// log.Printf("%#v\n", responseRecordSet)

	return &Record{
		RData: responseRecordSet.Rrdatas,
		RType: responseRecordSet.Type,
		RKey:  responseRecordSet.Name,
	}, nil
}

func (i *ZoneInfo) Set(r *Record) error {

	ctx := context.Background()

	dnsService := i.makeClient(ctx)
	recordSet := dns.ResourceRecordSet{
		Name:    r.RKey,
		Rrdatas: r.RData,
		Ttl:     int64(r.TTL),
		Type:    r.RType,
	}

	_, err := dnsService.ResourceRecordSets.Patch(i.ProjectId, i.ManagedZone, r.RKey, r.RType, &recordSet).Context(ctx).Do()
	if err != nil {
		return err
	}
	return nil
}
