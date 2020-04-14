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
	dataConfirmedURLs = []string{
		"https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_US.csv",
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

func getData(field, place string) int64 {

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
			r, e := strconv.ParseInt(d[len(d)-1], 0, 32)
			if e == nil {
				result = result + r
			}

		} else {

			d := strings.Split(v, ",")
			if len(d) > 1 {
				if strings.Contains(d[0], place) || strings.Contains(d[1], place) {
					r, e := strconv.ParseInt(d[len(d)-1], 0, 32)
					if e == nil {
						result = result + r
					}
				}
			}
		}
	}

	return result
}

func printActiveData(field, locationOne, locationTwo, locationThree string) {

	if locationOne != "" {
		confirmed := getData("Confirmed", locationOne)
		recovered := getData("Recovered", locationOne)
		fmt.Print(confirmed - recovered)
	}

	if locationTwo != "" {
		confirmed := getData("Confirmed", locationTwo)
		recovered := getData("Recovered", locationTwo)
		if locationOne != "" {
			fmt.Print("/")
		}
		fmt.Print(confirmed - recovered)
	}

	if locationThree != "" {
		confirmed := getData("Confirmed", locationThree)
		recovered := getData("Recovered", locationThree)
		if locationOne != "" || locationTwo != "" {
			fmt.Print(":")
		}
		fmt.Print(confirmed - recovered)
	}
}

func printData(field, locationOne, locationTwo, locationThree string) {

	if locationOne != "" {
		resultOne := getData(field, locationOne)
		fmt.Print(resultOne)
	}

	if locationTwo != "" {
		if locationOne != "" {
			fmt.Print("/")
		}
		resultTwo := getData(field, locationTwo)
		fmt.Print(resultTwo)
	}

	if locationThree != "" {
		if locationOne != "" || locationTwo != "" {
			fmt.Print(":")
		}
		resultThree := getData(field, locationThree)
		fmt.Print(resultThree, " ")
	}

}

func main() {

	var activeFlag bool
	var confirmedFlag bool
	var deadFlag bool
	var globalFlag bool
	var recoveredFlag bool
	var globalString = ""

	flag.BoolVar(&activeFlag, "a", false, "Show 'active' data.")
	flag.BoolVar(&confirmedFlag, "c", false, "Show 'confirmed' data.")
	flag.BoolVar(&deadFlag, "d", false, "Show 'dead' data.")
	flag.BoolVar(&recoveredFlag, "r", false, "Show 'recovered' data.")
	flag.BoolVar(&globalFlag, "g", false, "Show global location data")
	flag.StringVar(&locationFlag, "l", "", "Specify the primary location")
	flag.StringVar(&secondLocationFlag, "o", "", "Specify the secondary/alternative location")

	flag.Parse()

	if globalFlag {
		globalString = "global"
	}

	if activeFlag {
		printActiveData("Active", locationFlag, secondLocationFlag, globalString)
		fmt.Println()
	}

	if confirmedFlag {
		printData("Confirmed", locationFlag, secondLocationFlag, globalString)
		fmt.Println()
	}

	if deadFlag {
		printData("Deaths", locationFlag, secondLocationFlag, globalString)
		fmt.Println()
	}

	if recoveredFlag {
		printData("Recovered", locationFlag, secondLocationFlag, globalString)
		fmt.Println()
	}

	os.Exit(0)
}
