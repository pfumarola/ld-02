package wordsearch

import (
	"encoding/json"
	"log"
	"os"
)

type row struct {
	Word          string
	TimesSearched int
	LastDF        int
	LastTF        int
	TFList        []int
	DFList        []int
}

var db DB

type DB struct {
	Channel chan Word
}

func (_db *DB) InitDB() {
	_db.Channel = make(chan Word)
	go _db.storeInDBRoutine()
	db = *_db
}

func (db *DB) storeInDBRoutine() {
	defer close(db.Channel)
	for {
		word := <-db.Channel
		db.store(&word)
	}
}

func (db *DB) store(word *Word) {
	// read the file
	pwd, _ := os.Getwd()
	dbContent, err := os.ReadFile("./db.json")
	if err != nil {
		log.Println(err)
	}
	// unmarshal the file with the file struct
	var f []row
	err = json.Unmarshal(dbContent, &f)
	if err != nil {
		log.Println(err)
	}
	// search the word in the file
	found := false
	for i, row := range f {
		if row.Word == word.Word {
			found = true
			f[i].TimesSearched++
			f[i].LastDF = word.DocumentFrequency
			f[i].LastTF = word.TextFrequency
			f[i].TFList = append(f[i].TFList, word.TextFrequency)
			f[i].DFList = append(f[i].DFList, word.DocumentFrequency)
		}
	}
	// if not found, add it to the file
	if !found {
		f = append(f, row{
			Word:          word.Word,
			TimesSearched: 1,
			LastDF:        word.DocumentFrequency,
			LastTF:        word.TextFrequency,
			TFList:        []int{word.TextFrequency},
			DFList:        []int{word.DocumentFrequency},
		})
	}
	// marshal the file
	content, err := json.Marshal(f)
	if err != nil {
		log.Println(err)
	}
	// write the file
	err = os.WriteFile(pwd+"/db.json", content, 0644)
	if err != nil {
		log.Println(err)
	}
}
