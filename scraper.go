package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Datascraped struct {
	CompanyTitle          string `json:"company_title"`
	PreviousClose         string `json:"previous_close"`
	Open                  string `json:"open"`
	Bid                   string `json:"bid"`
	Ask                   string `json:"ask"`
	DaysRange             string `json:"days_range"`
	Week52Range           string `json:"week_52_range"`
	Volume                string `json:"volume"`
	AvgVolume             string `json:"avg_volume"`
	MarketCap             string `json:"market_cap"`
	Beta                  string `json:"beta"`
	PERatio               string `json:"pe_ratio"`
	EPS                   string `json:"eps"`
	EarningsDate          string `json:"earnings_date"`
	ForwardDividendYield  string `json:"forward_dividend_yield"`
	ExDividendDate        string `json:"ex_dividend_date"`
	OneYearTargetEstimate string `json:"one_year_target_estimate"`
}

func main() {

	companies := []string{"MSFT", "AAPL", "NVDA", "META"}
	datascraped := Datascraped{}

	datascrapes := make([]Datascraped, 0, 1)

	c := colly.NewCollector(colly.AllowedDomains("www.finance.yahoo.com", "finance.yahoo.com"))

	c.SetRequestTimeout(10 * time.Second)

	c.OnHTML("h1.yf-xxbei9", func(h *colly.HTMLElement) {
		datascraped.CompanyTitle = h.Text

	})

	c.OnHTML("div.container ul li", func(h *colly.HTMLElement) {

		title := h.ChildText("span.label")
		value := h.ChildText("span.value")

		if title != "" {

			if value != "" {
				switch title {

				case "Previous Close":
					datascraped.PreviousClose = value

				case "Open":
					datascraped.Open = value

				case "Bid":
					datascraped.Bid = value

				case "Ask":
					datascraped.Ask = value

				case "Day's Range":
					datascraped.DaysRange = value

				case "52 Week Range":
					datascraped.Week52Range = value

				case "Volume":
					datascraped.Volume = value

				case "Avg. Volume":
					datascraped.AvgVolume = value

				case "Market Cap (intraday)":
					datascraped.MarketCap = value

				case "Beta (5Y Monthly)":
					datascraped.Beta = value

				case "PE Ratio (TTM)":
					datascraped.PERatio = value

				case "EPS (TTM)":
					datascraped.EPS = value

				case "Earnings Date":
					datascraped.EarningsDate = value

				case "Forward Dividend & Yield":
					datascraped.ForwardDividendYield = value

				case "Ex-Dividend Date":
					datascraped.ExDividendDate = value

				case "1y Target Est":
					datascraped.OneYearTargetEstimate = value

				}
			} else {

				switch title {
				case "Previous Close":
					datascraped.PreviousClose = "N/A"

				case "Market Cap (intraday)":
					datascraped.MarketCap = "N/A"

				}
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visitin %s", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while scraping: %s\n", err.Error())
	})

	c.OnScraped(func(r *colly.Response) {

		datascrapes = append(datascrapes, datascraped)
		datascraped = Datascraped{}
	})

	for _, company := range companies {
		c.Visit(URLscraping(company))
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(datascrapes)

	content, err := json.Marshal(datascrapes)
	if err != nil {
		fmt.Println(err.Error())
	}
	os.WriteFile("scraped.json", content, 0644)

}

func URLscraping(company string) string {
	return "https://finance.yahoo.com/quote/" + company
}
