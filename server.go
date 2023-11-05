package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pfumarola/ld-02/wordsearch"
)

type Server struct {
}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/search", searchHandler)
	r.Run()
}

func searchHandler(c *gin.Context) {
	var list []string

	err := json.Unmarshal([]byte(c.Query("list")), &list)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bad request"})
		return
	}

	wordList := wordsearch.WordList{}
	wordList.InitWordList(list)

	wordList.Search()

	c.JSON(http.StatusOK, gin.H{
		"list": wordList.List,
	})
}
