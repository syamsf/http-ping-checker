package internal

type DnsRecord struct {
	hostname   string
	recordType string
}

type HitResult struct {
	Url         string              `csv:"url"`
	RecordType  string              `csv:"record_type"`
	Ping        PingResult          `csv:"ping"`
	HttpResult  TempHttpCheckResult `csv:"http"`
	HttpsResult TempHttpCheckResult `csv:"https"`
}

type TempHttpCheckResult struct {
	StatusCode   int    `csv:"status_code"`
	Status       bool   `csv:"status"`
	ErrorMessage string `csv:"error_message"`
}

type PingResult struct {
	IsSuccess    bool   `csv:"is_success"`
	ErrorMessage string `csv:"error_message"`
}
