package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	formatter "github.com/pedro-leitao/jwksformatter/formatter"
)

func main() {
	uri := flag.String("uri", "http://prd-keystore-obd.s3-website-eu-west-1.amazonaws.com/00158000016i44jAAA/00158000016i44jAAA.jwks", "the URI where the JWKS lives")
	templ := flag.String("templ", "templates/csv.tmpl", "the template file we will use for formatting")
	flag.Parse()

	// Load the template
	templFormat, err := ioutil.ReadFile(*templ)
	if err != nil {
		log.Fatalf("failed to load template: %v", err)
	}

	// Load the JWKS
	var keyset formatter.JWKS

	if err := keyset.Load(*uri); err != nil {
		log.Fatalf("failed to load JWKS: %v", err)
	}

	// Apply the template and output the formatted result
	result, err := keyset.Format(string(templFormat))
	if err != nil {
		log.Fatalf("failed to format JWKS: %v", err)
	}

	fmt.Println(result)
}
