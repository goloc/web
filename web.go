package main

import (
	//"encoding/json"
	"flag"
	"fmt"
	martini "github.com/go-martini/martini"
	core "github.com/goloc/core"
)

func main() {
	inputFile := flag.String("in", "", "input file")
	flag.Parse()
	if *inputFile == "" {
		fmt.Printf("Input file is mandatory\n")
	}
	if *inputFile == "" {
		fmt.Printf("\nExecute help: web -help\n")
		return
	}
	mi := core.NewMemindexFromFile(*inputFile)

	m := martini.Classic()
	m.Get("/search/:search", func(params martini.Params) string {

		var str string
		results := mi.Search(params["search"], 10, 600, 300)
		for _, res := range results {
			str += res.Localisation.GetName()
		}
		return str
	})
	m.Run()
}
