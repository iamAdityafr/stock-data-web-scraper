# Stock Data Web Scraper in Golang

This project is a web scraper built in Golang using the Colly library to gather stock data from Yahoo Finance. The slice of strings of company ticker symbols are added from which it automatically iterates over each one, and collects relevant stock information. The scraped data is then saved in a structured .json file within the present working directory.

## Key Features:

✨ **Dynamic Stock Selection** – Ticker symbols of companies are added in the slice of strings which can easily be added or removed as per the requirement and the scraper collects the latest data for each from Yahoo Finance.

📁 **JSON Storage** – Collected data is saved as .json for easy access and further analysis.

**🔄 Customizable Output Format** – The scraped data is currently stored as `.json`, could be configured to save in other formats like CSV.


## Installation:
Install Golang -

• Ensure you have Golang installed on your system. 

Set Up Colly -

• Install the Colly library by running:

```
go get -u github.com/gocolly/colly
```
Clone the Repository -

• Clone this repository to your local machine:

```
git clone https://github.com/iamAdityafr/stock-data-web-scraper.git

cd stock-data-web-scraper
```
### Usage:
Add Ticker Symbols: In the code, add the ticker symbols of the companies whose stock data you want to scrape. These should be added to the slice of strings within the code by the name ```companies```.

• Run the Scraper:

```
go run main.go
```
This will start the scraping process, fetching stock data for each specified company from Yahoo Finance.

Access the Data: Once scraping is complete, the data will be saved as data.json in the present working directory. You can open this file to view the JSON-formatted stock data.


