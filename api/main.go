package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

type Product struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type Address struct {
	AddressData string `json:"address"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var addr Address
	_ = json.NewDecoder(r.Body).Decode(&addr)
	address := addr.AddressData
	log.Println("address is:" + address)

	allProducts := make([]Product, 0)

	collector := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	collector.OnHTML(".octopus-dlp-asin-section", func(element *colly.HTMLElement) {
		title := element.ChildAttr(".a-size-base", "title")
		priceStr := "" + element.ChildText(".a-price-whole") + element.ChildText(".a-price-fraction")
		priceStr = strings.ReplaceAll(priceStr, ",", "")
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			price = 0
		}
		product := Product{
			Title: title,
			Price: price,
		}
		allProducts = append(allProducts, product)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	collector.Visit(address)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allProducts)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/address", homePage).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
