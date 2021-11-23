package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Product struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func main() {
	allProducts := make([]Product, 0)

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

	collector.Visit("https://www.amazon.com.tr/deal/7b0d0301?showVariations=true&smid=A1UNQM1SR2CHM&pf_rd_r=X5EDD9E248NAS06S3GVC&pf_rd_p=6c9a22ab-31e0-4aec-ad02-11e5945c72a4&pd_rd_r=33e42bb7-1f31-4108-a90e-22b8a0699969&pd_rd_w=jOmCO&pd_rd_wg=cytiS&ref_=pd_gw_unk")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allProducts)

	writeJSON(allProducts)

}

func writeJSON(data []Product) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create JSON file")
		return
	}

	_ = ioutil.WriteFile("data.json", file, 0644)
}
