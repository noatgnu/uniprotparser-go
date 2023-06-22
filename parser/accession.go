package parser

import (
	"fmt"
	"regexp"
)

type Accession struct {
	Raw     string
	Acc     string
	Isoform string
}

var regexPattern, _ = regexp.Compile("([OPQ][0-9][A-Z0-9]{3}[0-9]|[A-NR-Z][0-9]([A-Z][A-Z0-9]{2}[0-9]){1,2})(-\\d+)?")

func NewAccession(raw string) (*Accession, error) {
	match := regexPattern.FindStringSubmatch(raw)
	if len(match) == 0 {
		return &Accession{
			Raw:     raw,
			Acc:     "",
			Isoform: "",
		}, fmt.Errorf("invalid accession id")
	} else if match[3] == "" {
		return &Accession{
			Raw:     raw,
			Acc:     match[1],
			Isoform: "",
		}, nil
	}
	return &Accession{
		Raw:     raw,
		Acc:     match[1],
		Isoform: match[3],
	}, nil
}

func ToString(a *Accession) string {
	if a.Acc == "" {
		return a.Raw
	}

	if a.Isoform == "" {
		return a.Acc
	}
	return fmt.Sprintf("%s%s", a.Acc, a.Isoform)
}
