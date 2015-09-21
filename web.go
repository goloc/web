package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goloc/goloc"
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
	mi := goloc.NewMemindexFromFile(*inputFile)

	router := gin.Default()

	router.GET("/locations/:id", func(c *gin.Context) {
		loc := mi.Get(c.Params.ByName("id"))
		c.JSON(200, loc)
	})

	router.GET("/places/:search", func(c *gin.Context) {
		if list, err := mi.Search(goloc.Parameters{"search": c.Params.ByName("search"), "limit": 5}); err == nil {
			c.JSON(200, list.ToArray())
		} else {
			c.JSON(500, err)
		}
	})

	router.Run(":3000")
}
