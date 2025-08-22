package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Stock struct {
	Company string
	Price   string
	Change  string
}

func main() {

	ticker := []string{
		"MSFT", "IBM", "GE", "UNP", "COST", "MCD", "V", "DIS", "INTC", "AXP", "AAPL", "BA", "CSCO", "GS", "JPM", "CRM", "VZ",
	}

	var stocks []Stock

	// c := colly.NewCollector()'

	c := colly.NewCollector(
		colly.Async(true), // default is sync=false
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
		// fmt.Println("Visiting:", r.URL)
	})

	c.SetRequestTimeout(30 * time.Second)

	c.Limit(&colly.LimitRule{
    DomainGlob:  "*finance.yahoo.*",
    Parallelism: 2,
    Delay:       2 * time.Second, // adds delay between requests
})


c.OnResponse(func(r *colly.Response) {
    fmt.Println("Response status:", r.StatusCode)
    fmt.Println(string(r.Body)[:500]) // print first 500 chars
})


	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		// var stock Stock
		stock := Stock{}

		// Extract company name
		stock.Company = e.ChildText("h1")
		fmt.Println("Company:", stock.Company)

		// Extract current price
		stock.Price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("Price:", stock.Price)

		// Extract percentage change
		stock.Change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println("Change:", stock.Change)

		stocks = append(stocks, stock)
	})

	// Error handling
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	for _, t := range ticker {
		c.Visit("https://finance.yahoo.com/quote/" + t + "/")
		// fmt.Println("https://finance.yahoo.com/quote/"+t+"/")
	}
	c.Wait()
	fmt.Println("\nCollected Stocks:", stocks)

	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatal("Failed to create ouput CSV file", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	headers := []string{
		"company", "price", "change",
	}

	writer.Write(headers)

	for _, stock := range stocks {
		record := []string{
			stock.Company,
			stock.Price,
			stock.Change,
		}
		writer.Write(record)
	}

	defer writer.Flush()
}
