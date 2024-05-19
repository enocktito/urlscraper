# URL Scraper

## Overview
This Go program connects to one or more HTTP URLs provided as command line parameters, extracts all the links from each URL, and outputs them either as one absolute URL per line or as a JSON hash. The JSON hash format uses the base domain as the key and an array of relative paths as the value.
Based on https://github.com/gocolly/colly

## Features
- Connects to specified URLs.
- Extracts and lists all links.
- Outputs results either as plain text (stdout) or in JSON format.

## Installation
1. Ensure you have [Go installed](https://golang.org/doc/install) and/or [Docker installed](https://docs.docker.com/engine/install/)
2. Clone this repository or copy the provided code into your Go workspace.

## Build with go
```sh
$ go mod tidy
$ go build
```

## Build with docker
```sh
$ docker buildx build -t urlscraper:0.1 --no-cache -f Dockerfile . 
```

## Usage
### Command Line Options
- `-u`: Specify the URL to scrape. Can be used multiple times for multiple URLs.
- `-o`: Specify the output format. Options are `stdout` or `json`. Default is `stdout`.

### Examples
#### Example 1: Output one absolute URL per line
```sh
$ ./urlScraper -u "https://news.ycombinator.com/" -o "stdout"
https://news.ycombinator.com/newest
https://news.ycombinator.com/newcomments
https://news.ycombinator.com/ask
...
```

#### Example 2: Output as a JSON hash
```sh
$ ./urlScraper -u "https://news.ycombinator.com/" -u "https://arstechnica.com/" -o "json"
{
  "https://news.ycombinator.com": [
    "/newest",
    "/newcomments",
    "/ask",
    ...
  ],
  "https://arstechnica.com": [
    "/",
    "/civis/",
    "/store/product/subscriptions/",
    ...
  ]
}
```
#### RUN With docker
```sh
$ docker run docker.io/sedjro/urlscraper:0.1 -u "https://linkchecker.github.io/" -u "https://redis-py.readthedocs.io/" -o "json" > output/result.json

$ docker run docker.io/sedjro/urlscraper:0.1 -u "https://linkchecker.github.io/" -u "https://redis-py.readthedocs.io/" -o "stdout" > output/result.txt
```


Credits: https://github.com/gocolly/colly, https://www.scrapingbee.com/