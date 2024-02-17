package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
	"time"
	"web-crawl/internal"
)

func main() {
	start := time.Now()
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fileName := os.Getenv("FILENAME")
	dnsList := internal.Parser(fileName)

	checkStatus(dnsList)

	elapsedTime := time.Since(start)
	log.Printf("Elapsed Time: %s", elapsedTime.String())
}

func checkStatus(dnsList []internal.DnsRecord) {
	var result []internal.HitResult
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for _, item := range dnsList {
		log.Printf("%s", item)
		wg.Add(1)

		go func(dnsRecordItem internal.DnsRecord) {
			defer wg.Done()

			res := internal.Check(&dnsRecordItem)

			mutex.Lock()
			result = append(result, *res)
			mutex.Unlock()
		}(item)
	}

	wg.Wait()

	internal.WriteToCsv(result)
}
