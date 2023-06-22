package parser

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"github.com/go-gota/gota/dataframe"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var DefaultColumns = "accession,id,gene_names,protein_name,organism_name,organism_id,length,xref_refseq,xref_geneid,xref_ensembl,go_id,go_p,go_c,go_f,cc_subcellular_location,ft_topo_dom,ft_carbohyd,mass,cc_mass_spectrometry,sequence,ft_var_seq,cc_alternative_products"
var baseUrl = "https://rest.uniprot.org/idmapping/run"
var checkStatusUrl = "https://rest.uniprot.org/idmapping/status/"
var nextLinkPattern, _ = regexp.Compile("<(.*)>;")

type Parser struct {
	pollInterval   int
	session        *http.Client
	format         string
	columns        string
	includeIsoform bool
}

func NewParser(pollInterval int) *Parser {
	session := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return &Parser{pollInterval, session, "tsv", DefaultColumns, false}
}

func (p *Parser) SetFormat(format string) error {
	if format != "tsv" && format != "fasta" {
		return errors.New("Invalid format")
	}
	p.format = format
	return nil
}

func (p *Parser) SetColumns(columns string) {
	p.columns = columns
}

func (p *Parser) SetIncludeIsoform(includeIsoform bool) {
	p.includeIsoform = includeIsoform
}

func (p *Parser) GetColumns() string {
	return p.columns
}

func (p *Parser) GetFormat() string {
	return p.format
}

func (p *Parser) GetIncludeIsoform() bool {
	return p.includeIsoform
}

func (p *Parser) GetPollInterval() int {
	return p.pollInterval
}

func (p *Parser) GetSession() *http.Client {
	return p.session
}

func (p *Parser) Parse(ids []string, segment int) {
	portions := make([][]string, 0)
	for i := 0; i < len(ids); i += segment {
		end := i + segment
		if end > len(ids) {
			end = len(ids)
		}
		portions = append(portions, ids[i:end])
	}

	resultLinkChan := make(chan string)
	defer close(resultLinkChan)
	resultDFChan := make(chan dataframe.DataFrame)
	defer close(resultDFChan)

	for _, portion := range portions {
		params := url.Values{
			"from": {"UniProtKB_AC-ID"},
			"to":   {"UniProtKB"},
			"ids":  {strings.Join(portion, ",")},
		}
		paramsString := strings.NewReader(params.Encode())

		req, err := http.NewRequest("POST", baseUrl, paramsString)
		if err != nil {
			panic(err)
		}

		req.Header.Add("User-Agent", "uniprotparser-go")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := p.session.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			result := &JobIDResponse{}
			err = json.NewDecoder(resp.Body).Decode(result)
			if err != nil {
				log.Fatal(err)
			}
			job := Job{session: p.session, JobURL: checkStatusUrl + result.JobID, pollInterval: p.pollInterval}
			go func() {
				resultLink := ""
				for {
					resultLink = job.CheckJob()
					if resultLink != "" {
						resultLinkChan <- resultLink
						break
					}
					time.Sleep(time.Duration(p.pollInterval) * time.Second)
				}
			}()
		}
	}

	resultArray := make([]dataframe.DataFrame, 0)
	for i := 0; i < len(portions); i++ {
		result := <-resultLinkChan
		p.processLinkRecursive(result, &resultArray)
	}
	log.Println(resultArray)
	finDF := dataframe.New()
	for _, df := range resultArray {
		if finDF.Nrow() == 0 {
			finDF = df
			continue
		}
		finDF = finDF.RBind(df)
	}
	log.Println(finDF)
	log.Println("Done")
}

func (p *Parser) processLinkRecursive(resultLink string, resultArray *[]dataframe.DataFrame) {
	table, reader := p.getUniProtResult(resultLink)
	d := p.readResultIntoDataframe(reader)
	nextLink := p.getNextLink(table)
	if nextLink != "" {
		log.Println("Next link: " + nextLink)
		p.processLinkRecursive(nextLink, resultArray)
	}
	*resultArray = append(
		*resultArray,
		d,
	)
}

func (p *Parser) getNextLink(table *http.Response) string {
	nextLink := table.Header.Get("link")
	if nextLink != "" {
		nextLink = nextLinkPattern.FindStringSubmatch(nextLink)[1]
	}
	return nextLink
}

func (p *Parser) readResultIntoDataframe(reader io.ReadCloser) dataframe.DataFrame {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanRunes)
	var buf bytes.Buffer
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
	}

	d := dataframe.ReadCSV(bufio.NewReader(&buf), dataframe.WithDelimiter('\t'), dataframe.HasHeader(true))
	return d
}

func (p *Parser) getUniProtResult(result string) (*http.Response, io.ReadCloser) {
	req, err := http.NewRequest("GET", result, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "uniprotparser-go")
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Accept-Encoding", "gzip,deflate")
	req.Header.Add("Accept", "text/plain")
	q := req.URL.Query()
	q.Add("format", p.format)
	q.Add("fields", p.columns)
	q.Add("size", strconv.Itoa(500))
	if p.includeIsoform {
		q.Add("includeIsoform", "true")
	} else {
		q.Add("includeIsoform", "false")
	}
	req.URL.RawQuery = q.Encode()
	log.Println("Processing " + req.URL.String())
	table, err := p.session.Do(req)
	if err != nil {
		panic(err)
	}
	defer table.Body.Close()
	var reader io.ReadCloser
	switch table.Header.Get("Content-Encoding") {
	case "gzip":
		gz, err := gzip.NewReader(table.Body)
		if err != nil {
			panic(err)
		}
		defer gz.Close()
		reader = gz
	default:
		reader = table.Body
	}
	return table, reader
}
