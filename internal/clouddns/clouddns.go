package clouddns

import (
	"context"
	"errors"
	"fmt"
	"log"

	dns "google.golang.org/api/dns/v1"
	"google.golang.org/api/googleapi"
)

var (
	AlreadyExisted = errors.New("dnscli: already existed")
	NotFound       = errors.New("dnscli: not found")
	FatalError     = errors.New("dnscli: fatal error")
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
	Create(*Record) error
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
		var gError *googleapi.Error
		if match := errors.As(err, &gError); match {
			eR := &Record{}
			switch gError.Code {
			case 409:
				return eR, fmt.Errorf("%w", AlreadyExisted)
			case 404:
				return eR, fmt.Errorf("%w", NotFound)
			default:
				return eR, fmt.Errorf("%w", FatalError)
			}
		}
	}

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
		var gError *googleapi.Error
		if match := errors.As(err, &gError); match {
			switch gError.Code {
			case 409:
				return fmt.Errorf("%w", AlreadyExisted)
			case 404:
				return fmt.Errorf("%w", NotFound)
			default:
				return fmt.Errorf("%w", FatalError)
			}
		}
	}
	return nil
}

func (i *ZoneInfo) Create(r *Record) error {

	ctx := context.Background()

	dnsService := i.makeClient(ctx)
	recordSet := dns.ResourceRecordSet{
		Name:    r.RKey,
		Rrdatas: r.RData,
		Ttl:     int64(r.TTL),
		Type:    r.RType,
	}

	_, err := dnsService.ResourceRecordSets.Create(i.ProjectId, i.ManagedZone, &recordSet).Context(ctx).Do()
	if err != nil {
		var gError *googleapi.Error
		if match := errors.As(err, &gError); match {
			switch gError.Code {
			case 409:
				return fmt.Errorf("%w", AlreadyExisted)
			case 404:
				return fmt.Errorf("%w", NotFound)
			default:
				return fmt.Errorf("%w", FatalError)
			}
		}
	}
	return nil
}
