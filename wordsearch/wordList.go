package wordsearch

import (
	"sync"
)

type WordList struct {
	List []Word `json:"list"`
}

func (wl *WordList) InitWordList(stringList []string) {
	wl.List = make([]Word, len(stringList))

	for i, stringWord := range stringList {
		wl.List[i] = Word{stringWord, 0, 0, &sync.Mutex{}}
	}
}

func (wl *WordList) Search() {
	wg := sync.WaitGroup{}

	for i := 0; i < len(wl.List); i++ {
		wg.Add(1)
		go wl.List[i].searchInDocuments(&wg)
	}

	wg.Wait()
}
