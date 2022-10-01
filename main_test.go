package main

import (
	"reflect"
	"testing"

	"github.com/shin5ok/dnscli/internal/clouddns"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	setRecord := clouddns.Record{
		RType:  "A",
		RData:  []string{"192.168.0.1"},
		RKey:   "foo",
		TTL:    60,
		Status: "OK",
	}

	mockZoneInfo := &MockZoneInfo{}
	v := DNSMain{
		Client: mockZoneInfo,
	}

	err := v.Client.Set(&setRecord)
	assert.NoError(t, err)

	getRecord, err := v.Client.Get("foo")
	assert.NoError(t, err)

	assert.True(t, reflect.DeepEqual(getRecord.RData, []string{"192.168.0.1"}))
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
