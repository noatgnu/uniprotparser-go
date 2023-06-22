package parser

import (
	"reflect"
	"testing"
)

func TestNewAccession(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    *Accession
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Test accession id with isoform", args{"O60271-4"}, &Accession{"O60271-4", "O60271", "-4"}, false},
		{"Test accession id without isoform", args{"O60271"}, &Accession{"O60271", "O60271", ""}, false},
		{"Test accession id with isoform", args{"1111111"}, &Accession{"1111111", "", ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAccession(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAccession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccession() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		a *Accession
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Test accession id with isoform", args{&Accession{"O60271-4", "O60271", "-4"}}, "O60271-4"},
		{"Test accession id without isoform", args{&Accession{"O60271", "O60271", ""}}, "O60271"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.a); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
