package mariner

import "github.com/khulnasoft-lab/vul-db/pkg/vulnsrc/mariner/oval"

type operator string

type Entry struct {
	PkgName  string
	Version  string
	Operator operator
	Metadata oval.Metadata
}
