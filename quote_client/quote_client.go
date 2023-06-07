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
	message := ""
	randomInt := getnewrandom(len(quote_map))
	author, category, quote := quote_model.GetQuoteInfo(quote_map[randomInt])
	message = fmt.Sprintf("<div style='margin: 25px 50px 75px 100px'><ul style='list-style-type: none'><li>Author: %s</li><li>Category: %s</li> <li>Quote: %s</li><ul></div>", author, category, quote)
	return message
}

func parsecsv(data string) map[int]quote_model.Quote {
	read := csv.NewReader(strings.NewReader(data))
	quote_map := make(map[int]quote_model.Quote)
	record_count := 1

	for {
		record, err := read.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		response := quote_model.NewQuote(record[0], record[1], record[2])
		quote_map[record_count] = response
		record_count++
		continue
	}
	return quote_map
}

func getnewrandom(max int) int {
	randSpeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randSpeed.Intn(max-0) + 0
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
