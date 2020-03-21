package coronagoaway

import (
	"text/template"
	"bytes"
	"log"
	"net/http"
	"io/ioutil"
	"strconv"
)

type SimpleDate struct {
	Month string
	Day string
	Year string
}

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
		Day: day,
		Year: year,
	}
}

// TODO: Throw some error if not on a valid date
func GetDataForDate(date SimpleDate) string {
	url := BuildUrlFromDate(date)
	resp, err := SendRequest(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

// Example:
// https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/01-22-2020.csv
func BuildUrlFromDate(date SimpleDate) string {
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
	return "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/" + buffer.String()
}

func SendRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("URL was " + url)
		log.Println(err)
	}
	return resp, err
}
