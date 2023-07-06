package main

import (
	"github.com/balazsgrill/extractld"
	"github.com/balazsgrill/extractld/app"
	baseapp "github.com/balazsgrill/oauthenticator/app"
)

func main() {
	list := &extractld.UrlProcessorList{}
	main := &app.ExtractorApp{
		List:      list,
		Processor: list,
	}
	baseapp.Main(main)
}
