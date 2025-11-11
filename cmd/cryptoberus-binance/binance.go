package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/jsonflowdev/cryptoberus/internal/configs"
	"github.com/jsonflowdev/cryptoberus/internal/platform/coinmarketcap"
	"github.com/olekukonko/tablewriter"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application failed: %v", err)

	}

}

func setupLogger() error {
	logFile, err := os.OpenFile("development.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(logFile)
	defer logFile.Close()

	return nil
}

func run() error {
	viper, err := configs.Load(".")
	if err != nil {
		return fmt.Errorf(" can not load config: %w", err)
	}

	if viper.App.Environment == "development" {
		err = setupLogger()
		if err != nil {
			return fmt.Errorf("failed to setup logger: %w", err)
		}
	}

	coins, err := coinmarketcap.GetTopCoinsByMarketCap(500)
	if err != nil {
		return fmt.Errorf("failed to get top coins by market cap: %w", err)
	}
	color.New(color.FgCyan, color.Bold).Printf("Top Coins by Market Cap:\n")
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Rank", "Name", "Symbol", "Market Cap"})
	for _, coin := range coins {
		table.Append([]string{
			fmt.Sprintf("%d", coin.MarketCapRank),
			color.New(color.FgGreen).Sprintf("%s", coin.Name),
			coin.Symbol,
			fmt.Sprintf("$%s", humanize.Commaf(coin.MarketCap)),
		})
	}
	table.Render()

	return nil

}
