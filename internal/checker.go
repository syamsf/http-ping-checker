package internal

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Check(item *DnsRecord) *HitResult {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	tempResult := &HitResult{
		Url:        item.hostname,
		RecordType: item.recordType,
	}

	protocol := map[string]string{"https": "https://", "http": "http://"}

	wg.Add(3)

	for key, value := range protocol {
		go func(currKey string, currProtocol string) {
			defer wg.Done()

			res := goLookupHttp(currProtocol, item)

			mutex.Lock()
			if currKey == "https" {
				tempResult.HttpsResult = res
			} else {
				tempResult.HttpResult = res
			}
			mutex.Unlock()
		}(key, value)
	}

	go func() {
		defer wg.Done()
		tempResult.Ping = goPingCmd(item.hostname)
	}()

	wg.Wait()

	aggregatedResult := &HitResult{
		Url:         item.hostname,
		RecordType:  item.recordType,
		HttpResult:  tempResult.HttpResult,
		HttpsResult: tempResult.HttpsResult,
		Ping:        tempResult.Ping,
	}

	return aggregatedResult
}

func goLookupHttp(currProtocol string, item *DnsRecord) TempHttpCheckResult {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	var client = &http.Client{Transport: tr, Timeout: 30 * time.Second}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", currProtocol, item.hostname), nil)
	if err != nil {
		return TempHttpCheckResult{ErrorMessage: err.Error()}
	}

	res, err := client.Do(req)
	if err != nil {
		return TempHttpCheckResult{ErrorMessage: err.Error()}
	}

	defer res.Body.Close()

	return TempHttpCheckResult{
		StatusCode: res.StatusCode,
		Status:     true,
	}
}

func goPingCmd(hostname string) PingResult {
	pingResult := PingResult{}

	err := PingCmd(hostname)
	if err != nil {
		pingResult.ErrorMessage = err.Error()
		return pingResult
	}

	pingResult.IsSuccess = true
	return pingResult
}
