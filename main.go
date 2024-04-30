package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func MakeRequest(url string, numRequests int, concurrency int) {
	var statusOk, statusErr int
	statusOk = 0
	statusErr = 0
	now := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numRequests/concurrency; j++ {
				_, err := http.Get(url)
				if err != nil {
					statusErr++
				}
				statusOk++
			}
		}()
	}

	wg.Wait()
	fmt.Println("Relatório Final do Teste de Carga")
	fmt.Println("Tempo total gasto na execução:", time.Since(now).Seconds(), "segundos")
	fmt.Println("Quantidade Total de requests realizados:", numRequests)
	fmt.Println("Quantidade de requests com status HTTP 200:", statusOk)
	fmt.Println("Distribuição de outros códigos de status HTTP (4xx, 5xx):", statusErr)
}

func main() {
	url := flag.String("url", "http://google.com", "URL to make requests")
	numRequests := flag.Int("requests", 100, "Number of requests to make")
	concurrency := flag.Int("concurrency", 10, "Number of requests to make concurrently")
	flag.Parse()

	MakeRequest(*url, *numRequests, *concurrency)
}
