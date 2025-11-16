package main

import (
	"testing"

	"github.com/cravenceiling/indexer/cli/cmd/program"
	"github.com/cravenceiling/indexer/cli/internal/parser"
)

var (
	re = program.HttpRequest{
		Creds: program.Credentials{
			User:     "admin",
			Password: "Complexpass#123",
		},
		BaseURL: "localhost",
		Index:   "profiling",
		Type:    "_doc",
		Port:    "4080",
	}

	directory = "../enron_mail_20110402"
	indexer   = program.Indexer{
		Parser: parser.Parser{},
	}
)

func BenchmarkXxx(b *testing.B) {
	for b.Loop() {
		indexer.Index(directory, re)
	}
}
