package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goloc/goloc"
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
	mi := goloc.NewMemindexFromFile(*inputFile)

	router := gin.Default()
	router.GET("/locations/:id", func(c *gin.Context) {
		loc := mi.Get(c.Params.ByName("id"))
		c.JSON(200, loc)
	})
	router.GET("/places/:search", func(c *gin.Context) {
		list := mi.Search(c.Params.ByName("search"), 5, nil)
		c.JSON(200, list.ToArray())
	})
	router.Run(":3000")
}
