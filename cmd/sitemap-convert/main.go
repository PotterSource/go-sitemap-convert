package main

import (
	"flag"
	"fmt"
	"github.com/Pottersource/go-sitemap-convert/sitemap"
	"os"
)

func main() {
	// Define subcommands
	csvCmd := flag.NewFlagSet("csv", flag.ExitOnError)
	jsonCmd := flag.NewFlagSet("json", flag.ExitOnError)

	// Define flags for the csv subcommand
	csvOutput := csvCmd.String("o", "sitemap.csv", "Output CSV file path")
	csvURL := csvCmd.String("url", "", "Sitemap URL to fetch")

	// Define flags for the json subcommand
	jsonOutput := jsonCmd.String("o", "sitemap.json", "Output JSON file path")
	jsonURL := jsonCmd.String("url", "", "Sitemap URL to fetch")

	// Verify that a subcommand has been provided
	if len(os.Args) < 2 {
		fmt.Println("csv or json subcommand is required")
		fmt.Println("Usage:")
		fmt.Println("  {cmd} csv -url {sitemap URL} -o {output file path}")
		fmt.Println("  {cmd} json -url {sitemap URL} -o {output file path}")
		os.Exit(1)
	}

	// Switch based on the subcommand
	switch os.Args[1] {
	case "csv":
		_ = csvCmd.Parse(os.Args[2:])
		handleCSVCommand(csvCmd, *csvURL, *csvOutput)
	case "json":
		_ = jsonCmd.Parse(os.Args[2:])
		handleJSONCommand(jsonCmd, *jsonURL, *jsonOutput)
	default:
		fmt.Println("Unknown subcommand:", os.Args[1])
		fmt.Println("Available subcommands: csv, json")
		os.Exit(1)
	}
}

func handleCSVCommand(cmd *flag.FlagSet, sitemapURL, outputPath string) {
	if sitemapURL == "" {
		fmt.Println("Please provide a sitemap URL using the -url flag")
		cmd.PrintDefaults()
		os.Exit(1)
	}

	// Fetch and parse the sitemap
	siteMap, err := sitemap.GetSitemapFromUrl(sitemapURL)
	if err != nil {
		fmt.Printf("Error fetching sitemap: %v\n", err)
		os.Exit(1)
	}

	// Output to CSV file
	err = siteMap.OutputCSV(outputPath)
	if err != nil {
		fmt.Printf("Error writing CSV file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Sitemap URLs successfully exported to %s\n", outputPath)
}

func handleJSONCommand(cmd *flag.FlagSet, sitemapURL, outputPath string) {
	if sitemapURL == "" {
		fmt.Println("Please provide a sitemap URL using the -url flag")
		cmd.PrintDefaults()
		os.Exit(1)
	}

	// Fetch and parse the sitemap
	siteMap, err := sitemap.GetSitemapFromUrl(sitemapURL)
	if err != nil {
		fmt.Printf("Error fetching sitemap: %v\n", err)
		os.Exit(1)
	}

	// Output to JSON file
	err = siteMap.OutputJSON(outputPath)
	if err != nil {
		fmt.Printf("Error writing JSON file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Sitemap URLs successfully exported to %s\n", outputPath)
}
