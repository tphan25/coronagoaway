package coronagoaway

import (
	"encoding/csv"
	"log"
	"reflect"
	"strings"
)

// CoronaData is the structure of the data in the table + a Date.
// Data schema keeps changing (unannounced?)
// 1: Province/State,Country/Region,Last Update,Confirmed,Deaths,Recovered
// 2: Province/State,Country/Region,Last Update,Confirmed,Deaths,Recovered,Latitude,Longitude
// 3: FIPS,Admin2,Province_State,Country_Region,Last_Update,Lat,Long_,Confirmed,Deaths,Recovered,Active,Combined_Key
type CoronaData struct {
	Date        string
	FIPS        string `titles:"FIPS"`
	Admin2      string `titles:"Admin2"`
	Province    string `titles:"Province/State&Province_State"`
	Country     string `titles:"Country/Region&Country_Region"`
	LastUpdate  string `titles:"Last Update&Last_Update"`
	Latitude    string `titles:"Lat&Latitude"`
	Longitude   string `titles:"Long_&Longitude"`
	Confirmed   string `titles:"Confirmed"`
	Deaths      string `titles:"Deaths"`
	Recovered   string `titles:"Recovered"`
	Active      string `titles:"Active"`
	CombinedKey string `titles:"Combined_Key"`
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
			// https://stackoverflow.com/questions/6395076/using-reflect-how-do-you-set-the-value-of-a-struct-field
			curr := CoronaData{
				Date:       date,
				Province:   line[0],
				Country:    line[1],
				LastUpdate: line[2],
				Confirmed:  line[3],
				Deaths:     line[4],
				Recovered:  line[5],
			}
			ret = append(ret, curr)
		}
	}

	return ret, nil
}

// mapNamesToIndices will find which index at which some data resides, based on header name
func mapNamesToIndices(headerLine []string) map[string]int {
	mapHeadersToFieldNames := make(map[string]string)

	c := CoronaData{}
	v := reflect.ValueOf(c)
	typeOfC := v.Type()
	for i := 0; i < v.NumField(); i++ {
		currField := typeOfC.Field(i)
		currFieldName := typeOfC.Field(i).Name
		currTag := currField.Tag.Get("titles")
		// Fields not reflected in CSV headers at all
		if currTag == "" {
			continue
		}
		arr := strings.Split(currTag, "&")
		for _, title := range arr {
			mapHeadersToFieldNames[title] = currFieldName
		}
	}

	ret := make(map[string]int)
	for i, column := range headerLine {
		ret[column] = i
	}
	return ret
}
