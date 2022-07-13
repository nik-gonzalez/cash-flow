package models

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestJsonDateTime_MarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	validTime, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	if err != nil {
		t.Fatalf("error parsing time: %v", err)
	}
	marshalledValidTime, err := json.Marshal(validTime)
	if err != nil {
		t.Fatalf("error marshalling valid time: %v", err)
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "valid time in bytes marshalled to correct JsonDateTime",
			fields:  fields{Time: validTime},
			want:    marshalledValidTime,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := JsonDateTime{
				Time: tt.fields.Time,
			}
			got, err := mt.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonDateTime_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Time time.Time
	}
	type args struct {
		bs []byte
	}
	validTimeTxt := "2006-01-02T15:04:05+07:00"
	validTime, err := time.Parse(time.RFC3339, validTimeTxt)
	if err != nil {
		t.Errorf("error parsing time: %s; error: %v", validTimeTxt, err)
	}
	marshalledValidTime, err := json.Marshal(validTime)
	if err != nil {
		t.Errorf("error marshalling time: %s; error: %v", validTime, err)
	}
	invalidMarshalledDate, err := json.Marshal("")
	if err != nil {
		t.Errorf("error marshalling empty string; error: %v", err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "valid date time should unmarshall to valid JSONDate",
			fields:  fields{Time: validTime},
			args:    args{bs: marshalledValidTime},
			wantErr: false,
		},
		{
			name:    "invalid date time should result in error",
			args:    args{bs: []byte("")},
			wantErr: true,
		},
		{
			name:    "invalid marshalled date time should result in error",
			args:    args{bs: invalidMarshalledDate},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mt := &JsonDateTime{
				Time: tt.fields.Time,
			}
			if err := mt.UnmarshalJSON(tt.args.bs); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
