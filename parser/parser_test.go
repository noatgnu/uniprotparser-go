package parser

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewParser(t *testing.T) {
	type args struct {
		pollInterval int
	}
	tests := []struct {
		name string
		args args
		want *Parser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParser(tt.args.pollInterval); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetColumns(t *testing.T) {
	type fields struct {
		pollInterval   int
		session        *http.Client
		format         string
		columns        string
		includeIsoform bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				pollInterval:   tt.fields.pollInterval,
				session:        tt.fields.session,
				format:         tt.fields.format,
				columns:        tt.fields.columns,
				includeIsoform: tt.fields.includeIsoform,
			}
			if got := p.GetColumns(); got != tt.want {
				t.Errorf("GetColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetFormat(t *testing.T) {
	type fields struct {
		pollInterval   int
		session        *http.Client
		format         string
		columns        string
		includeIsoform bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				pollInterval:   tt.fields.pollInterval,
				session:        tt.fields.session,
				format:         tt.fields.format,
				columns:        tt.fields.columns,
				includeIsoform: tt.fields.includeIsoform,
			}
			if got := p.GetFormat(); got != tt.want {
				t.Errorf("GetFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetIncludeIsoform(t *testing.T) {
	type fields struct {
		pollInterval   int
		session        *http.Client
		format         string
		columns        string
		includeIsoform bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				pollInterval:   tt.fields.pollInterval,
				session:        tt.fields.session,
				format:         tt.fields.format,
				columns:        tt.fields.columns,
				includeIsoform: tt.fields.includeIsoform,
			}
			if got := p.GetIncludeIsoform(); got != tt.want {
				t.Errorf("GetIncludeIsoform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetPollInterval(t *testing.T) {
	type fields struct {
		pollInterval   int
		session        *http.Client
		format         string
		columns        string
		includeIsoform bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				pollInterval:   tt.fields.pollInterval,
				session:        tt.fields.session,
				format:         tt.fields.format,
				columns:        tt.fields.columns,
				includeIsoform: tt.fields.includeIsoform,
			}
			if got := p.GetPollInterval(); got != tt.want {
				t.Errorf("GetPollInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetSession(t *testing.T) {
	type fields struct {
		pollInterval   int
		session        *http.Client
		format         string
		columns        string
		includeIsoform bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				pollInterval:   tt.fields.pollInterval,
				session:        tt.fields.session,
				format:         tt.fields.format,
				columns:        tt.fields.columns,
				includeIsoform: tt.fields.includeIsoform,
			}
			if got := p.GetSession(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_Parse(t *testing.T) {
	type fields struct {
		pollInterval   int
		session        *http.Client
		format         string
		columns        string
		includeIsoform bool
	}
	type args struct {
		ids     []string
		segment int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			"Test accession id",
			fields{
				5,
				&http.Client{
					CheckRedirect: func(req *http.Request, via []*http.Request) error {
						return http.ErrUseLastResponse
					},
				},
				"tsv",
				"accession,id,gene_names,protein_name,organism_name,organism_id,length,xref_refseq",
				true},
			args{[]string{"O60271", "O60271-4"},
				1000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				pollInterval:   tt.fields.pollInterval,
				session:        tt.fields.session,
				format:         tt.fields.format,
				columns:        tt.fields.columns,
				includeIsoform: tt.fields.includeIsoform,
			}
			p.Parse(tt.args.ids, tt.args.segment)
		})
	}
}

func TestParser_SetColumns(t *testing.T) {
	type fields struct {
		pollInterval   int
		session        *http.Client
		format         string
		columns        string
		includeIsoform bool
	}
	type args struct {
		columns string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				pollInterval:   tt.fields.pollInterval,
				session:        tt.fields.session,
				format:         tt.fields.format,
				columns:        tt.fields.columns,
				includeIsoform: tt.fields.includeIsoform,
			}
			p.SetColumns(tt.args.columns)
		})
	}
}

func TestParser_SetFormat(t *testing.T) {
	type fields struct {
		pollInterval   int
		session        *http.Client
		format         string
		columns        string
		includeIsoform bool
	}
	type args struct {
		format string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				pollInterval:   tt.fields.pollInterval,
				session:        tt.fields.session,
				format:         tt.fields.format,
				columns:        tt.fields.columns,
				includeIsoform: tt.fields.includeIsoform,
			}
			if err := p.SetFormat(tt.args.format); (err != nil) != tt.wantErr {
				t.Errorf("SetFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_SetIncludeIsoform(t *testing.T) {
	type fields struct {
		pollInterval   int
		session        *http.Client
		format         string
		columns        string
		includeIsoform bool
	}
	type args struct {
		includeIsoform bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				pollInterval:   tt.fields.pollInterval,
				session:        tt.fields.session,
				format:         tt.fields.format,
				columns:        tt.fields.columns,
				includeIsoform: tt.fields.includeIsoform,
			}
			p.SetIncludeIsoform(tt.args.includeIsoform)
		})
	}
}
