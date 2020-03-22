package coronagoaway

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// SimpleDate is easiest way to structure for figuring out which CSV to query.
type SimpleDate struct {
	Month string
	Day   string
	Year  string
}

// UnstructuredCoronaData is pre-processed from the query, pure CSV.
type UnstructuredCoronaData struct {
	Date string
	Data string
}

// CreateSimpleDate generates a simple date, and appends 0 to the start to follow their naming conventions.
func CreateSimpleDate(intMonth int, intDay int, intYear int) SimpleDate {
	var month, day, year string
	if intMonth < 10 {
		month = "0" + strconv.Itoa(intMonth)
	} else {
		month = strconv.Itoa(intMonth)
	}

	if intDay < 10 {
		day = "0" + strconv.Itoa(intDay)
	} else {
		day = strconv.Itoa(intDay)
	}

	year = strconv.Itoa(intYear)

	return SimpleDate{
		Month: month,
		Day:   day,
		Year:  year,
	}
}

// GetDataForDate gets the data from the response.
// TODO: Throw some error if not on a valid date
func GetDataForDate(date SimpleDate) UnstructuredCoronaData {
	url := BuildURLFromDate(date)
	resp, err := SendRequest(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return UnstructuredCoronaData{
		Data: string(body),
		Date: buildDateString(date),
	}
}

func buildDateString(date SimpleDate) string {
	const dateTemplate = `{{.Month}}-{{.Day}}-{{.Year}}.csv`
	tmpl, err := template.New("dateFormat").Parse(dateTemplate)
	if err != nil {
		log.Println("Template could not be parsed")
	}
	var buffer bytes.Buffer
	err = tmpl.ExecuteTemplate(&buffer, "dateFormat", date)
	if err != nil {
		log.Println("Template could not be executed")
	}
	return buffer.String()
}

// BuildURLFromDate basically just lets us scrape that specific URL for our data.
// Example:
// https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/01-22-2020.csv
func BuildURLFromDate(date SimpleDate) string {
	str := buildDateString(date)
	return "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/" + str
}

// SendRequest actually sends a simple http GET request to the URL with small error handle.
func SendRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("URL was " + url)
		log.Println(err)
	}
	return resp, err
}
