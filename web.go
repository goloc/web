package main

import (
	"encoding/json"
	"flag"
	"fmt"
	martini "github.com/go-martini/martini"
	core "github.com/goloc/core"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(16)

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
	m.Get("/search/:search", func(params martini.Params) []byte {
		list := mi.Search(params["search"], 5)
		json, _ := json.Marshal(list.ToArray())
		return json
	})
	m.Run()
}
