package coronagoaway

import (
	"encoding/csv"
	"log"
	"strings"
)

// CoronaData is the structure of the data in the table + a Date.
type CoronaData struct {
	Date           string
	Province       string
	Country        string
	LastUpdate     string
	ConfirmedCases string
	Deaths         string
	Recovered      string
}

// ReadCsvData returns the lines from the file to read
func ReadCsvData(data string) ([][]string, error) {
	lines, err := csv.NewReader(strings.NewReader(data)).ReadAll()
	if err != nil {
		log.Println("Error occurred while parsing CSV")
		return [][]string{}, err
	}
	return lines, nil
}

// GetCoronaDataFromCsv gets the actual data as a slice of CoronaData objects from the CSV
func GetCoronaDataFromCsv(data string, date string) ([]CoronaData, error) {
	lines, err := ReadCsvData(data)
	if err != nil {
		return []CoronaData{}, err
	}
	ret := make([]CoronaData, 0)

	for i, line := range lines {
		// Skip first line (headers)
		if i > 0 {
			curr := CoronaData{
				Date:           date,
				Province:       line[0],
				Country:        line[1],
				LastUpdate:     line[2],
				ConfirmedCases: line[3],
				Deaths:         line[4],
				Recovered:      line[5],
			}
			ret = append(ret, curr)
		}
	}

	return ret, nil
}
