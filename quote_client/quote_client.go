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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Qc() string {
	p, _ := configparser.NewConfigParserFromFile("config/config.cfg")
	v, _ := p.Get("DEFAULT", "CsvFilePath")
	data, err := os.ReadFile(v)
	check(err)
	ds := string(data)
	read := csv.NewReader(strings.NewReader(ds))
	message := ""
	record_map := make(map[int]quote_model.Quote)
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
		record_map[record_count] = response
		record_count++
		continue
	}
	randomInt := getnewrandom()
	author, category, quote := quote_model.GetQuoteInfo(record_map[randomInt])
	message = fmt.Sprintf("<div style='margin: 25px 50px 75px 100px'><ul style='list-style-type: none'><li>Author: %s</li><li>Category: %s</li> <li>Quote: %s</li><ul></div>", author, category, quote)
	return message
}

func getnewrandom() int {
	randSpeed := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randSpeed.Intn(48390-0) + 0
}
