package coronagoaway

import (
	"testing"
)

func TestSimpleDateQuery(t *testing.T) {
	data := GetDataForDate(CreateSimpleDate(3, 23, 2020))
	coronaObject, _ := GetCoronaDataFromCsv(data.Data, data.Date)
	testObj := CoronaData{
		Date:        "",
		FIPS:        "45001",
		Admin2:      "Abbeville",
		Province:    "South Carolina",
		Country:     "US",
		LastUpdate:  "2020-03-23 23:19:34",
		Latitude:    "34.22333378",
		Longitude:   "-82.46170658",
		Confirmed:   "1",
		Deaths:      "0",
		Recovered:   "0",
		Active:      "0",
		CombinedKey: "Abbeville, South Carolina, US",
	}

	if testObj != coronaObject[0] {
		t.Fatal("IT BROKE")
	}
}
