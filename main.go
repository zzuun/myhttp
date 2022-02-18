package main

import (
	"crypto/md5"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Response struct {
	Add  string
	Hash string
	err  error
}

func main() {
	worker := *flag.Int("parallel", 1, "parallel requests to be made")
	flag.Parse()
	addrList := flag.Args()

	reqChan := make(chan string)
	resChan := make(chan Response)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := range reqChan {
				resp, err := MakeHTTPRequest(r)
				resChan <- Response{Add: r, Hash: ConvertBytesToMD5(resp), err: err}
			}
		}()
	}

	go func() {
		defer close(reqChan)
		for i, _ := range addrList {
			reqChan <- addrList[i]
		}
	}()

	go func() {
		for r := range resChan {
			if r.err == nil {
				fmt.Println(fmt.Sprintf("%s  %s", r.Add, r.Hash))
			}
		}
	}()

	wg.Wait()
	close(resChan)
}

func MakeHTTPRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not get %s: %v", url, err)
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Println("failed to close res body ", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("error calling endpoint")
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("error in reading the response")
	}

	return b, nil
}

func ConvertBytesToMD5(bytes []byte) string {
	return fmt.Sprintf("%x", md5.Sum(bytes))
}
