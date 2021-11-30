package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

type Product struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	vars := mux.Vars(r)
	address := vars["address"]
	log.Println("address is:" + address)

	allProducts := make([]Product, 0)

	cronJob := cron.New()
	cronJob.AddFunc("@every 8s", func() {
		collector := colly.NewCollector(
			colly.AllowURLRevisit(),
		)

		collector.Limit(&colly.LimitRule{
			DomainGlob:  "https://www.amazon.com.tr/*",
			Delay:       3 * time.Second,
			RandomDelay: 3 * time.Second,
		})

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

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", " ")
		enc.Encode(allProducts)

		writeJSON(allProducts)
		collector.Visit("https://www.amazon.com.tr/gp/bestsellers/?ref_=nav_em_cs_bestsellers_0_1_1_2" + address)
	})
	cronJob.Start()
	time.Sleep(time.Second * 100)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/address/{address}", homePage)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}

func writeJSON(data []Product) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create JSON file")
		return
	}

	_ = ioutil.WriteFile("data.json", file, 0644)
}
