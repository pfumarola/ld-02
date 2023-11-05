package main

import (
	"github.com/pfumarola/ld-02/wordsearch"
)

func main() {

	wsdb := wordsearch.DB{}
	wsdb.InitDB()
	server := Server{}
	defer server.Run()
}
