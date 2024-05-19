package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Scan struct {
	Collector *colly.Collector
	Links     map[string]int
}

func prefixTimmer(url string, domain string) string {
	url = strings.TrimPrefix(url, "http://")
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, domain)

	return url
}

func (s *Scan) generateCollector(url string, allowedDomains ...string) {

	// Define default allowed domains
	domain := prefixTimmer(url, "")
	domain = strings.TrimSuffix(domain, "/")
	if len(allowedDomains) == 0 {
		allowedDomains = append(allowedDomains, domain)
	}

	// Instantiate default collector
	s.Collector = colly.NewCollector(
		// Visit only domains allowed
		colly.AllowedDomains(allowedDomains...),
	)

	// On every a element which has href attribute call callback
	s.Collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		s.Collector.Visit(e.Request.AbsoluteURL(link))
	})

	// Create a callback on the XPath query searching for the URLs
	s.Collector.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
		// knownUrls = append(knownUrls, e.Text)
		url := prefixTimmer(e.Text, domain)
		if _, found := s.Links[url]; !found {
			s.Collector.Visit(e.Text)
		}
	})

	// Before making a request register the path"
	s.Collector.OnRequest(func(r *colly.Request) {
		link := r.URL.Path
		if link != "" {
			s.Links[link]++
		}
	})
}

func main() {
	// Create variables
	var urls []string
	scan := Scan{colly.NewCollector(), make(map[string]int)}
	scanResults := make(map[string][]string)

	// Define flags
	flag.Func("u", "URL to scrape. Can be specified multiple times.", func(s string) error {
		urls = append(urls, s)
		return nil
	})
	output := flag.String("o", "stdout", "Output format: stdout or json.")
	flag.Parse()

	// Handle the case where no url were provided
	if len(urls) == 0 {
		panic("Please provide at least one URL using the -u flag.")
	}

	// Generate scan result
	for _, url := range urls {
		scan.generateCollector(url)
		scan.Collector.Visit(url)
		scan.Collector.Visit(strings.TrimSuffix(url, "/") + "/sitemap.xml")
		// paths := getUrlPaths(url)
		for path := range scan.Links {
			scanResults[url] = append(scanResults[url], path)
		}
	}

	// Print the scan result
	switch *output {
	case "json":
		jsonOutput, err := json.MarshalIndent(scanResults, "", "  ")
		if err != nil {
			fmt.Println("Error generating JSON output:", err)
			return
		}
		fmt.Println(string(jsonOutput))
	case "stdout":
		for domain, paths := range scanResults {
			for _, path := range paths {
				fmt.Printf("%s%s\n", strings.TrimSuffix(domain, "/"), path)
			}
		}
	}
}
