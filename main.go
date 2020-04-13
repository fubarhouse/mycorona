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
	dataURL = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/04-14-2020.csv"

	dataConfirmedURLs = []string{
		"https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/04-12-2020.csv",
		"https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv",
	}
	dataDeathsURLs = []string{
		"https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_US.csv",
		"https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_global.csv",
	}
	dataRecoveredURLs = []string{
		"https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_recovered_global.csv",
	}

	locationFlag       string
	secondLocationFlag string
)

var dataConfirmed []string
var dataDeaths []string
var dataRecovered []string

func init() {

	for _, url := range dataConfirmedURLs {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		for _, line := range strings.Split(string(body), "\n") {
			dataConfirmed = append(dataConfirmed, line)
		}
		resp.Body.Close()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}

	for _, url := range dataDeathsURLs {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		for _, line := range strings.Split(string(body), "\n") {
			dataDeaths = append(dataDeaths, line)
		}
		resp.Body.Close()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}

	for _, url := range dataRecoveredURLs {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		for _, line := range strings.Split(string(body), "\n") {
			dataRecovered = append(dataRecovered, line)
		}
		resp.Body.Close()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}

}

func getData(url, field, place string) int64 {

	var data []string
	switch field {
	case "Recovered":
		data = dataRecovered
	case "Deaths":
		data = dataDeaths
	case "Active":
		break
	case "Confirmed":
		data = dataConfirmed
	}

	var result int64

	for _, v := range data {

		if place == "global" {
			d := strings.Split(v, ",")
			r, e := strconv.ParseInt(d[len(d)-1], 10, 10)
			if e == nil {
				result = result + r
			}

		} else {
			fields := strings.Split(v, ",")
			if len(fields) > 3 {
				province := fields[0]
				country := fields[1]

				if strings.Contains(province, place) || country == place {

					d := strings.Split(v, ",")
					r, e := strconv.ParseInt(d[len(d)-1], 10, 10)
					if e == nil {
						result = result + r
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
		//locationFlag = "Australian Capital Territory"
	}

	if globalFlag {
		globalString = "global"
	}

	if activeFlag {
		printData(dataURL, "Active", locationFlag, secondLocationFlag, globalString)
		fmt.Println()
	}

	if confirmedFlag {
		printData(dataURL, "Confirmed", locationFlag, secondLocationFlag, globalString)
		fmt.Println()
	}

	if deadFlag {
		printData(dataURL, "Deaths", locationFlag, secondLocationFlag, globalString)
		fmt.Println()
	}

	if recoveredFlag {
		printData(dataURL, "Recovered", locationFlag, secondLocationFlag, globalString)
	}

	os.Exit(0)
}
