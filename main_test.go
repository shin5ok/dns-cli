package main

import (
	"reflect"
	"testing"

	"github.com/shin5ok/dnscli/internal/clouddns"
)

func Test_main(t *testing.T) {
	setRecord := clouddns.Record{
		RType:  "A",
		RData:  []string{"192.168.0.1"},
		RKey:   "foo",
		TTL:    60,
		Status: "OK",
	}

	mockRr := &MockZoneInfo{}
	v := DNSMain{
		Client: mockRr,
	}

	err := v.Client.Set(&setRecord)
	if err != nil {
		t.Error(err)
	}
	getRecord, err := v.Client.Get("foo")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(getRecord.RData, []string{"192.168.0.1"}) {
		t.Error("RData:", getRecord.RData)
		t.Error("getRecord.RData is not equal expected")
	}
	t.Log(getRecord)
}

type MockZoneInfo struct {
	clouddns.Recorder
}

func (r *MockZoneInfo) Get(key string) (*clouddns.Record, error) {
	return &clouddns.Record{
		RType:  "A",
		RData:  []string{"192.168.0.1"},
		RKey:   "foo",
		TTL:    60,
		Status: "OK",
	}, nil
}

func (r *MockZoneInfo) Set(*clouddns.Record) error {
	return nil
}
