package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	dataURL = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/04-12-2020.csv"

	locationFlag       string
	secondLocationFlag string
)

func getData(url, field, place string) int64 {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	lines := strings.Split(string(body), "\n")
	fields := strings.Split(lines[0], ",")
	field_number := 0

	for n, f := range fields {
		if f == field {
			field_number = n
		}
	}

	var result int64

	for _, v := range lines {

		if place == "global" {
			d := strings.Split(v, ",")
			if len(v) >= field_number {
				r, e := strconv.ParseInt(d[field_number], 10, 10)
				if e == nil {
					result = result + r
				}
			}
		} else {
			fields := strings.Split(v, ",")
			if len(fields) > 3 {
				province := fields[2]
				country := fields[3]

				if strings.Contains(province, place) || strings.Contains(country, place) {

					d := strings.Split(v, ",")
					r, e := strconv.ParseInt(d[field_number], 10, 10)
					if e == nil {
						result = result + r
					} else {
						return int64(0)
					}
				}
			}
		}
	}

	return result
}

func printData(url, field, locationOne, locationTwo, locationThree string) {

	if locationOne != "" {
		resultOne := getData(url, field, locationOne)
		fmt.Print(resultOne)
	}

	if locationTwo != "" {
		resultTwo := getData(url, field, locationTwo)
		fmt.Print("/", resultTwo)
	}

	if locationThree != "" {
		resultThree := getData(url, field, locationThree)
		fmt.Print(":", resultThree, " ")
	}

}

func main() {

	var activeFlag bool
	var confirmedFlag bool
	var deadFlag bool
	var globalFlag bool
	var recoveredFlag bool
	var globalString = ""

	flag.BoolVar(&activeFlag, "a", false, "Specify username. Default is root")
	flag.BoolVar(&confirmedFlag, "c", false, "Specify username. Default is root")
	flag.BoolVar(&deadFlag, "d", false, "Specify pass. Default is password")
	flag.BoolVar(&recoveredFlag, "r", false, "Specify pass. Default is password")
	flag.BoolVar(&globalFlag, "g", false, "Specify pass. Default is password")
	flag.StringVar(&locationFlag, "l", "", "Specify pass. Default is password")
	flag.StringVar(&secondLocationFlag, "o", "", "Specify pass. Default is password")

	flag.Parse()

	if locationFlag == "" {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	if globalFlag {
		globalString = "global"
	}

	if activeFlag {
		printData(dataURL, "Active", locationFlag, secondLocationFlag, globalString)
	}

	if confirmedFlag {
		printData(dataURL, "Confirmed", locationFlag, secondLocationFlag, globalString)
	}

	if deadFlag {
		printData(dataURL, "Deaths", locationFlag, secondLocationFlag, globalString)
	}

	if recoveredFlag {
		printData(dataURL, "Recovered", locationFlag, secondLocationFlag, globalString)
	}

	os.Exit(0)
}
