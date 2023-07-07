# Wikipedia Comedy Scraper Go

![build](https://img.shields.io/github/actions/workflow/status/ericthomasca/wikipedia-comedy-scraper-go/go.yml)  ![golangci-lint](https://img.shields.io/github/actions/workflow/status/ericthomasca/wikipedia-comedy-scraper-go/golangci-lint.yml?label=golangci-lint)

A simple web scraper built in Go. It scrapes the Wikipedia page of stand-up comedy specials in 2023 and converts it to a csv... for now.

## TODO (in no particular order)

- Move the data to a database.
- Serve up an api with the info.
- Break down details to get comedian name, special name, network.
- Break down date string into a datetime.
