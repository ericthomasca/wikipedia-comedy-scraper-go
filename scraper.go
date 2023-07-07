package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type ComedySpecial struct {
	Date    string
	Details string
}

var months = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func main() {
	c := colly.NewCollector()

	var comedySpecials []ComedySpecial

	c.OnHTML("div.mw-parser-output", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, li *colly.HTMLElement) {
			for _, month := range months {
				if strings.HasPrefix(li.Text, month) {
					parts := strings.Split(li.Text, ":")

					date := strings.TrimSpace(parts[0])
					details := strings.TrimSpace(parts[1])

					// Trim quotations
					details = strings.ReplaceAll(details, "\"", "")

					// Remove everthing after the last period
					lastIndex := strings.LastIndex(details, ".")
					if lastIndex != -1 {
						details = details[:lastIndex]
					}

					comedySpecial := ComedySpecial{
						Date:    date,
						Details: details,
					}

					comedySpecials = append(comedySpecials, comedySpecial)
					break
				}
			}
		})
	})

	url := "https://en.wikipedia.org/wiki/2023_in_stand-up_comedy"

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	err = saveToCSV(comedySpecials)
	if err != nil {
		log.Fatal(err)
	}
}

func saveToCSV(comedySpecials []ComedySpecial) error {
	file, err := os.Create("specials.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Date", "Details"}
	err = writer.Write(header)
	if err != nil {
		return err
	}

	for _, comedySpecial := range comedySpecials {
		row := []string{comedySpecial.Date, comedySpecial.Details}
		err := writer.Write(row)
		if err != nil {
			return err
		}
	}

	fmt.Println("Data saved to specials.csv.")
	return nil
}
