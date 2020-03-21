package coronagoaway

import (
	"encoding/csv"
	"strings"
	"log"
)

// TODO: convert types?
type CoronaData struct {
	Province string
	Country string
	LastUpdate string
	ConfirmedCases string
	Deaths string
	Recovered string
}

func ReadCsvData(data string) ([][]string, error) {
	lines, err := csv.NewReader(strings.NewReader(data)).ReadAll()
	if err != nil {
		log.Println("Error occurred while parsing CSV")
		return [][]string{}, err
	}
	return lines, nil
}

func GetCoronaDataFromCsv(data string) ([]CoronaData, error) {
	lines, err := ReadCsvData(data)
	if err != nil {
		return []CoronaData{}, err
	}
	ret := make([]CoronaData, 0)

	for i, line := range lines {
		// Skip first line (headers)
		if i > 0 {
			curr := CoronaData{
				Province: line[0],
				Country: line[1],
				LastUpdate: line[2],
				ConfirmedCases: line[3],
				Deaths: line[4],
				Recovered: line[5],
			}
			ret = append(ret, curr)
		}
	}

	return ret, nil
}
