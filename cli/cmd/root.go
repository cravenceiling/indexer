package cmd

import (
	"flag"

	"github.com/cravenceiling/indexer/cli/cmd/program"
	"github.com/cravenceiling/indexer/cli/internal/parser"
)

var (
	flags   = Flags{}
	indexer = program.Indexer{}
)

type Flags struct {
	directory *string
	zincURL   *string
	port      *string
	user      *string
	password  *string
	_index    *string
	_type     *string
}

func Execute() {
	req := program.HttpRequest{
		Creds: program.Credentials{
			User:     *flags.user,
			Password: *flags.password,
		},
		BaseURL: *flags.zincURL,
		Index:   *flags._index,
		Type:    *flags._type,
		Port:    *flags.port,
	}

	indexer.Parser = parser.Parser{}
	indexer.Index(*flags.directory, req)
}

func init() {
	flags.directory = flag.String("dir", "enron_mail_20110402", "path to email directory")
	flags.zincURL = flag.String("zincurl", "localhost", "zincsearch host url")
	flags.port = flag.String("port", "4080", "zincsearch host port")
	flags.user = flag.String("user", "admin", "zincsearch username")
	flags.password = flag.String("password", "Complexpass#123", "zincsearch password")
	flags._index = flag.String("index", "enron", "index name")
	flags._type = flag.String("type", "_doc", "request payload type")

	flag.Parse()
}
