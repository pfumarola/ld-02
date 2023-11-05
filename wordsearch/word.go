package wordsearch

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"sync"
)

const PATH = "./collection/"

type Word struct {
	Word              string      `json:"word"`
	TextFrequency     int         `json:"TF"`
	DocumentFrequency int         `json:"DF"`
	Mutex             *sync.Mutex `json:"-"`
}

func (w *Word) countOccurrence(content []byte) int {
	re := regexp.MustCompile(`\b` + w.Word + `\b`)
	return len(re.FindAll(content, -1))
}

func (w *Word) searchInDocument(file fs.DirEntry, fileWg *sync.WaitGroup) {
	defer fileWg.Done()

	content, err := os.ReadFile(PATH + file.Name())
	if err != nil {
		log.Println(err)
	}
	occurency := w.countOccurrence(content)

	if occurency > 0 {
		w.Mutex.Lock()
		w.DocumentFrequency++
		w.TextFrequency += occurency
		w.Mutex.Unlock()
	}
}

func (w *Word) searchInDocuments(wg *sync.WaitGroup) {
	defer wg.Done()

	dir, err := os.Open(PATH)

	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := dir.ReadDir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileWg := sync.WaitGroup{}

	for _, file := range files {
		if !file.IsDir() {
			fileWg.Add(1)
			go w.searchInDocument(file, &fileWg)
		}
	}
	fileWg.Wait()
	db.Channel <- *w
}
