package quote_client

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/bigkevmcd/go-configparser"
	"github.com/specialagentsoftware/random_quote_service_go/quote_model"
)

func QcInit() map[int]quote_model.Quote {
	/*
		QcInit is a poor man's constructor for the quote_model object. This function retrives the path to the
		csv file from the config file, opens that file, reads it in and calls a function to create the map of quote objects
		and returns that map of quotes
	*/
	p, err := configparser.NewConfigParserFromFile("config/config.cfg")
	check(err)
	v, err := p.Get("DEFAULT", "CsvFilePath")
	check(err)
	data, err := os.ReadFile(v)
	check(err)
	ds := string(data)
	quote_map := parsecsv(ds)
	return quote_map
}

func Output(quote_map map[int]quote_model.Quote) string {
	/*
		output takes a quote map and returns a formatted string representing a random
		quote from the passed in map.
	*/
	message := ""
	randomInt := getnewrandom(len(quote_map))
	author, category, quote := quote_model.GetQuoteInfo(quote_map[randomInt])
	message = fmt.Sprintf("<div style='margin: 25px 50px 75px 100px'><ul style='list-style-type: none'><li>Author: %s</li><li>Category: %s</li> <li>Quote: %s</li><ul></div>", author, category, quote)
	return message
}

func parsecsv(data string) map[int]quote_model.Quote {
	/*
		parsecsv is a helper function that takes in a data string and generates a map of quote objects.
	*/
	read := csv.NewReader(strings.NewReader(data))
	quote_map := make(map[int]quote_model.Quote)
	record_count := 1

	for {
		record, err := read.Read()
		if err == io.EOF {
			break
		}
		check(err)
		response := quote_model.NewQuote(record[0], record[1], record[2])
		quote_map[record_count] = response
		record_count++
		continue
	}
	return quote_map
}

func getnewrandom(max int) int {
	/*
		getnewrandom is a helper function that takes a max int and generates / returns a random integer from a range
	*/
	randSpeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randSpeed.Intn(max-0) + 0
}

func check(e error) {
	/*
		check is a utility function that takes an error as a parameter and logs a fatal error if one has been provided
	*/
	if e != nil {
		log.Fatal(e)
	}
}
