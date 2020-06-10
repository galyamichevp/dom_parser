package core

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//PTable - ...
func parseTable(data string) [][]string {
	var headings, row []string
	var rows [][]string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}

	// Find each table
	doc.Find("div").Each(func(index int, item *goquery.Selection) {
		item.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
			tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
				rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
					headings = append(headings, tableheading.Text())
				})
				rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
					row = append(row, tablecell.Text())
				})
				rows = append(rows, row)
				row = nil
			})
		})
	})

	return rows
}

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

// Create a goquery document from the HTTP response
// document, err := goquery.NewDocumentFromReader(rr)
// if err != nil {
// 	log.Fatal("Error loading HTTP response body. ", err)
// }

// // Find all links and process them with the function
// // defined earlier
// document.Find("tr").Each(processElement)
