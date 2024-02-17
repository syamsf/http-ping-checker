package internal

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"log"
	"os"
	"strings"
)

func WriteToCsv(result []HitResult) {
	file, err := os.Create("checked.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := gocsv.MarshalFile(result, file); err != nil {
		panic(err)
	}
}

func Parser(fileName string) []DnsRecord {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error while reading the file. Message: ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading record. Message: ", err)
	}

	var dnsList []DnsRecord
	for i, record := range records {
		if i == 0 {
			continue
		}

		dnsList = append(dnsList, DnsRecord{
			hostname:   strings.TrimSuffix(record[0], "."),
			recordType: record[1],
		})
	}

	return dnsList
}
