package main

import (
	"flag"
	"os"
	"fmt"
)

type Flags struct {
	QueryStringQuery string
	Server string
	Index string
	Template string
	ConfigFile string
	FilterName string
	From string
	To string
}

func parseFlags() (*Flags) {
	flags := &Flags{}

	flag.StringVar(&flags.From, "from", "now-15m", "Elasticsearch date for gte")
	flag.StringVar(&flags.To, "to", "now", "Elasticsearch date for lte")
	flag.StringVar(&flags.QueryStringQuery, "query", "*", "Elasticsearch query string query")
	flag.StringVar(&flags.FilterName, "filter-name", "", "If specified use the esfilter's filter as the query")
	flag.StringVar(&flags.ConfigFile, "config", "", "Use configuration file created by esfilters")
	flag.StringVar(&flags.Server, "server", "http://localhost:9200", "Specify elasticsearch server to query")
	flag.StringVar(&flags.Index, "index", "", "Specify the elasticsearch index to query")
	flag.StringVar(&flags.Template, "template", "{{ . | json }}", "Specify Go text/template. You can use the function 'json' or 'json_indent'.")

	flag.Parse()

	if flags.Index == "" {
		fmt.Fprintln(os.Stderr, "Flag \"-index\" is required")
		flag.Usage()
		os.Exit(2)
	}

	if flags.Template == "" {
		fmt.Fprintln(os.Stderr, "Flag \"-template\" cannot be empty")
		flag.Usage()
		os.Exit(2)
	}

	if flags.To == "" {
		fmt.Fprintln(os.Stderr, "Flag \"-to\" cannot be empty")
		flag.Usage()
		os.Exit(2)
	}

	if flags.From == "" {
		fmt.Fprintln(os.Stderr, "Flag \"-from\" cannot be empty")
		flag.Usage()
		os.Exit(2)
	}

	if flags.FilterName != "" && flags.ConfigFile == "" {
		fmt.Fprintln(os.Stderr, "When \"-filter-name\" flag is used, flag \"-config\" has to be specified")
		flag.Usage()
		os.Exit(2)
	}

	if flags.FilterName != "" && (flags.QueryStringQuery != "*" && flags.QueryStringQuery != "") {
		fmt.Fprintln(os.Stderr, "Flags \"-filter-name\" and \"-query\" are mutually exclusive")
		flag.Usage()
		os.Exit(2)
	}

	flags.Template = fmt.Sprintf("%s\n", flags.Template)

	return flags
}

func init() {
	flag.Usage = func () {
		fmt.Fprintf(os.Stderr, "Usage of %s: [-config=file] [-query=Query | <-config=file> <-filter-name=FilterName>] <-server=Url> <-index=Index> [-to=date] [-from=date] [-template=Template]\n", os.Args[0])
		flag.PrintDefaults()
	}
}