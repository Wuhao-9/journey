package dirsearch

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"race-proj/code/dsflag"
	"race-proj/setting"
	"runtime"
	"sync"
)

type SearchInfo struct {
	Url  string `json:"url"`
	Code int    `json:"code"`
}

var postfix []string = make([]string, 0, 10000)
var mutex *sync.Mutex = &sync.Mutex{}

func init() {
	dictFile, err := os.OpenFile(setting.DictPath, os.O_RDONLY, 0000)
	if err != nil {
		log.Fatalf("fail to open dict file: %s", err)
	}
	defer dictFile.Close()

	scanner := bufio.NewScanner(dictFile)
	for scanner.Scan() {
		postfix = append(postfix, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("fail to scan dict: %s", err)
	}
}

func Search(url string, how int) []*SearchInfo {
	if url[len(url)-1] != '/' {
		url += "/"
	}

	var result []*SearchInfo = make([]*SearchInfo, 0)
	ave := len(postfix) / int(setting.DefGRNum)
	wg := &sync.WaitGroup{}
	for i := 0; i < int(setting.DefGRNum); i++ {
		wg.Add(1)
		if uint(i) == setting.DefGRNum-1 {
			go issue_request(i*ave, len(postfix), url, how, &result, wg)
		} else {
			go issue_request(i*ave, (i+1)*ave, url, how, &result, wg)
		}
	}
	wg.Wait()
	return result
}

func issue_request(begin, end int, url string, how int, cter *[]*SearchInfo, grp *sync.WaitGroup) {
	defer grp.Done()
	switch how {
	case dsflag.Search_Postfix:
		handle_postfix(begin, end, url, cter)
	case dsflag.Search_SubDomain:

	}
	runtime.Goexit()
}

func handle_postfix(begin, end int, url string, cter *[]*SearchInfo) {
	for _, str := range postfix[begin:end] {
		resp, err := http.Get(url + str)
		if err != nil {
			continue
		} else if resp.StatusCode == http.StatusNotFound {
			continue
		}
		mutex.Lock()
		*cter = append(*cter, &SearchInfo{
			Url:  url + str,
			Code: resp.StatusCode,
		})
		mutex.Unlock()
	}
}
