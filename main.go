// package main is the binary for the mycorona application and provides
// a simple cli to grab granular covid19 information for configured
// data inputs such as locations, and types of data. It will return results
// only, so basically the last entry in each dataset for the provided input
// or regex match. This package is designed to be consumed by a polybar
// script configuration item.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	// dataConfirmedURLs are time-series data exports updated every 24 hours (UTC 00:00).
	// These one contain data for confirmed cases of covid19.
	dataConfirmedURLs = []string{
		"https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_US.csv",
		"https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv",
	}
	// dataDeathsURLs are time-series data exports updated every 24 hours (UTC 00:00).
	// These one contain data for deaths from covid19.
	dataDeathsURLs = []string{
		"https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_US.csv",
		"https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_global.csv",
	}
	// dataRecoveredURLs are time-series data exports updated every 24 hours (UTC 00:00).
	// These one contain data for recovered cases of covid19.
	dataRecoveredURLs = []string{
		"https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_recovered_global.csv",
	}

	// locationFlag is a string to select a location from the datasets
	// The mechanism that selects the location also supports regex.
	locationFlag string

	// secondLocationFlag is a string to select a location from the datasets
	// The mechanism that selects the location also supports regex.
	secondLocationFlag string

	// dataConfirmed contains the data for confirmed cases of covid19.
	dataConfirmed []string

	// dataDeaths contains the data for covid deaths.
	dataDeaths []string

	// dataRecovered contains data for the recovered covid19 cases
	dataRecovered []string

	// activeFlag will indicate the user wants to see data for property 'active'
	activeFlag bool

	// confirmedFlag will indicate the user wants to see data for property 'confirmed'
	confirmedFlag bool

	// deadFlag will indicate the user wants to see data for property 'dead'
	deadFlag bool

	// globalFlag will indicate the user wants to see global data for all selected properties
	globalFlag bool

	// recoveredFlag will indicate the user wants to see data for property 'recovered'
	recoveredFlag bool
)

// init will on initialisation, populate the data.
// a caching layer is yet to be explored.
func init() {

	// Populate the data for confirmed cases of covid19
	// from the urls in dataConfirmedURLs.
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

	// Populate the data for cases of covid19 which
	// have resulted in death. Data will be from the
	// imported from the urls in dataDeathsURLs.
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

	// Populate the data for cases recovered from covid19
	// from the urls in dataRecoveredURLs.
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

// intFormat will take an int64 and return a formatted integer
// (string) which is simply the comma-separated equivalent.
func intFormat(num int64) string {
	str := fmt.Sprintf("%d", num)
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1,$2")
	}
	return str
}

// getData will process the data variable and find the entries which can be
// linked to up to two given location matching a regexp (providence/country)
// It will grab the last column from each of these rows and add it to the
// result. if the global flag is true, it will count all results in the
// last column for the given row of the dataset.
func getData(data []string, place string) int64 {

	var result int64

	for _, v := range data {

		d := strings.Split(v, ",")
		if len(d) > 1 {
			checkOne, _ := regexp.MatchString(place, d[0])
			checkTwo, _ := regexp.MatchString(place, d[1])
			if checkOne || checkTwo {
				r, e := strconv.ParseInt(d[len(d)-1], 0, 32)
				if e == nil {
					result = result + r
				}
			}
		}
	}

	return result
}

// printActiveData will print the qualified result from active
// cases for the dataset and location. Active cases are a simple
// operation of the confirmed cases with the recovered cases
// subtracted from it. It will do this for the three potential
// locations. Two optional and one results with the global data.
func printActiveData() {

	if locationFlag != "" {
		confirmed := getData(dataConfirmed, locationFlag)
		recovered := getData(dataRecovered, locationFlag)
		fmt.Print(intFormat(confirmed - recovered))
	}

	if secondLocationFlag != "" {
		confirmed := getData(dataConfirmed, secondLocationFlag)
		recovered := getData(dataRecovered, secondLocationFlag)
		if locationFlag != "" {
			fmt.Print("/")
		}
		fmt.Print(intFormat(confirmed - recovered))
	}

	if globalFlag {
		confirmed := getData(dataConfirmed, ".*")
		recovered := getData(dataRecovered, ".*")
		if locationFlag != "" || secondLocationFlag != "" {
			fmt.Print(":")
		}
		fmt.Print(intFormat(confirmed - recovered))
	}
}

// printData will query the data for the result and print it in the
// standard format which is LOCATION1/LOCATION2:GLOBAL where any of
// those fields can be absent.
func printData(data []string) {

	if locationFlag != "" {
		resultOne := getData(data, locationFlag)
		fmt.Print(intFormat(resultOne))
	}

	if secondLocationFlag != "" {
		if locationFlag != "" {
			fmt.Print("/")
		}
		resultTwo := getData(data, secondLocationFlag)
		fmt.Print(intFormat(resultTwo))
	}

	if globalFlag {
		if locationFlag != "" || secondLocationFlag != "" {
			fmt.Print(":")
		}
		resultThree := getData(data, ".*")
		fmt.Print(intFormat(resultThree), " ")
	}

}

func main() {

	flag.BoolVar(&activeFlag, "a", false, "Show 'active' data.")
	flag.BoolVar(&confirmedFlag, "c", false, "Show 'confirmed' data.")
	flag.BoolVar(&deadFlag, "d", false, "Show 'dead' data.")
	flag.BoolVar(&recoveredFlag, "r", false, "Show 'recovered' data.")
	flag.BoolVar(&globalFlag, "g", false, "Show global location data")
	flag.StringVar(&locationFlag, "l", "", "Specify the primary location")
	flag.StringVar(&secondLocationFlag, "o", "", "Specify the secondary/alternative location")

	flag.Parse()

	if activeFlag {
		printActiveData()
		fmt.Println()
	}

	if confirmedFlag {
		printData(dataConfirmed)
		fmt.Println()
	}

	if deadFlag {
		printData(dataDeaths)
		fmt.Println()
	}

	if recoveredFlag {
		printData(dataRecovered)
		fmt.Println()
	}

	os.Exit(0)
}
