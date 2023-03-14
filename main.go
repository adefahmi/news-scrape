package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(scrape()))
	})

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)

}

func scrape() string {
	doc, err := goquery.NewDocument("https://www.bmkg.go.id/")
	if err != nil {
		log.Fatal(err)
	}

	// Find the table by its ID or class name
	div := doc.Find(".press-release-home-bg")

	// create array
	var data []string

	// Loop through each row in the table
	div.Find(".blog-thumb").Each(func(i int, row *goquery.Selection) {
		// get image src
		imgSrc := getImgSrc(row)
		// get title and link
		titleAndLink := getTitleAndLink(row)

		// add to data array with format json comma separated
		data = append(data, fmt.Sprintf(`{"imgSrc":"%s","title":"%s","link":"%s","date":"%s"}`, imgSrc, titleAndLink[0], titleAndLink[1], titleAndLink[2]))

	})

	var result = "[" + strings.Join(data, ",") + "]"

	return result
}

// function get image src from row
func getImgSrc(row *goquery.Selection) string {
	col := row.Find(".blog-thumb-mkg")
	img := col.Find("img")
	src, _ := img.Attr("src")

	return src
}

// function get title and link from row return array
func getTitleAndLink(row *goquery.Selection) []string {
	col := row.Find(".blog-thumb-desc")
	title := col.Find("h3").Text()
	link, _ := col.Find("a").Attr("href")

	// add domain to link
	link = "https://www.bmkg.go.id" + link

	// get date from ul li
	date := col.Find("ul").Find("li").Text()

	return []string{title, link, date}
}
