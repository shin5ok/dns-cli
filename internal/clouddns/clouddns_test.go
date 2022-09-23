package clouddns

import (
	"os"
	"reflect"
	"testing"
)

type mockZoneInfo struct {
	ZoneInfo
}

var existedRecord = &Record{
	RType: "A",
	RData: []string{"192.168.0.1"},
	RKey:  "foo.example.com.",
}

func (r *mockZoneInfo) Get(key string) (*Record, error) {
	existedKey := existedRecord.RKey
	if key != existedKey {
		return &Record{}, ErrNotFound
	}
	return &Record{
		RType: "A",
		RData: []string{"192.168.0.1"},
		RKey:  existedKey,
		// TTL:    60,
		// Status: "OK",
	}, nil
}

func (r *mockZoneInfo) Set(*Record) error {
	return nil
}

func (r *mockZoneInfo) Create(*Record) error {
	return nil
}

func TestZoneInfo_Get(t *testing.T) {
	type fields struct {
		Domain      string
		ProjectId   string
		ManagedZone string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Record
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				Domain:      os.Getenv("DNS_DOMAIN"),
				ProjectId:   os.Getenv("GOOGLE_CLOUD_PROJECT"),
				ManagedZone: os.Getenv("DNS_ZONE"),
			},
			args: args{
				key: "foo.example.com.",
			},
			want: &Record{
				RData: []string{"192.168.0.1"},
				RKey:  "foo.example.com.",
				RType: "A",
			},
		},
		{
			name: "test2",
			fields: fields{
				Domain:      os.Getenv("DNS_DOMAIN"),
				ProjectId:   os.Getenv("GOOGLE_CLOUD_PROJECT"),
				ManagedZone: os.Getenv("DNS_ZONE"),
			},
			args: args{
				key: "notfound.example.com.",
			},
			want:    &Record{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &mockZoneInfo{
				ZoneInfo: ZoneInfo{
					Domain:      tt.fields.Domain,
					ProjectId:   tt.fields.ProjectId,
					ManagedZone: tt.fields.ManagedZone,
				},
			}
			got, err := i.Get(tt.args.key)

			if (err != nil) != tt.wantErr {
				t.Errorf("ZoneInfo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZoneInfo.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
