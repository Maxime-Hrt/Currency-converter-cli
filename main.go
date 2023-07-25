package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var fromCurrency string
var destCurrencies string
var amount float64

func main() {
	setDestination()

	API_KEY := getApiKey()

	body := getResult(API_KEY)

	fmt.Println(body)

}

func setDestination() {
	// Variables from cmd
	var toCurrencies string

	// Flags for arguments
	flag.StringVar(&fromCurrency, "from", "", "Currency to convert from")
	flag.StringVar(&toCurrencies, "to", "", "Currencies to convert to (comma-seperated)")
	flag.Float64Var(&amount, "amount", 0.0, "Amount to convert")
	flag.Parse()

	// Verifications
	if fromCurrency == "" {
		log.Fatal("Currency to convert from (-from) must be specified")
	}

	if toCurrencies == "" {
		log.Fatal("Currencies to convert to (-to) must be specified")
	}

	// Spliting currencies with comma and convert to link
	destCurrencies = strings.Join(strings.Split(toCurrencies, ","), "%2C")
}

func getApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("API_KEY")
}

func getResult(API_KEY string) string {
	link := fmt.Sprintf("https://api.freecurrencyapi.com/v1/latest?apikey=%s&currencies=%s&base_currency=%s",
		API_KEY,
		destCurrencies,
		fromCurrency,
	)
	res, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Currency API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	return string(body)
}
