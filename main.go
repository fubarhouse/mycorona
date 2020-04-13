package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var (
	confirmURL = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv"
	deathURL   = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_deaths_global.csv"
	recoverURL = "https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_recovered_global.csv"

	locationFlag string
)

func getData(url string) string {
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

	for _, v := range strings.Split(string(body), "\n") {
		if strings.Contains(v, locationFlag) {
			d := strings.Split(v, ",")
			return d[len(d)-1]
		}
	}

	return ""
}

func main() {

	var confirmedFlag bool
	var deadFlag bool
	var recoveredFlag bool

	flag.BoolVar(&confirmedFlag, "c", true, "Specify username. Default is root")
	flag.BoolVar(&deadFlag, "d", true, "Specify pass. Default is password")
	flag.BoolVar(&recoveredFlag, "r", true, "Specify pass. Default is password")
	flag.StringVar(&locationFlag, "l", "", "Specify pass. Default is password")

	flag.Parse()

	if locationFlag == "" {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	if confirmedFlag {
		confirmed := getData(confirmURL)
		fmt.Print(confirmed, " ")
	}
	if deadFlag {
		dead := getData(deathURL)
		fmt.Print(dead, " ")
	}
	if recoveredFlag {
		recovered := getData(recoverURL)
		fmt.Print(recovered, " ")
	}
	return
}
