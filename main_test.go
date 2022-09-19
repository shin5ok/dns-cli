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
	mockRr := &MockRr{}
	err := mockRr.Set(&setRecord)
	if err != nil {
		t.Error(err)
	}
	getRecord, err := mockRr.Get("foo")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(getRecord.RData, []string{"192.168.0.1"}) {
		t.Error("RData:", getRecord.RData)
		t.Error("getRecord.RData is not equal expected")
	}
	t.Log(getRecord)
}

type MockRr struct {
}

func (r *MockRr) Get(key string) (*clouddns.Record, error) {
	return &clouddns.Record{
		RType:  "A",
		RData:  []string{"192.168.0.1"},
		RKey:   "foo",
		TTL:    60,
		Status: "OK",
	}, nil
}

func (r *MockRr) Set(*clouddns.Record) error {
	return nil
}
