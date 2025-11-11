package coinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jsonflowdev/cryptoberus/internal/models"
)

const (
	baseURL            = "https://api.coingecko.com/api/v3/coins/markets"
	vsCurrency         = "usd"
	orderBy            = "market_cap_desc"
	apiPerPage         = 100
	defaultHTTPTimeout = 10 * time.Second
)

func GetTopCoinsByMarketCap(n int) ([]models.Coin, error) {
	if n <= 0 {
		return nil, errors.New("n must be > 0")
	}

	client := &http.Client{Timeout: defaultHTTPTimeout}

	var all []models.Coin
	for page, want := 1, n; want > 0; page, want = page+1, n-len(all) {
		perPage := want
		if perPage > apiPerPage {
			perPage = apiPerPage
		}

		url := fmt.Sprintf(
			"%s?vs_currency=%s&order=%s&per_page=%d&page=%d&sparkline=false",
			baseURL, vsCurrency, orderBy, perPage, page,
		)

		resp, err := client.Get(url)
		if err != nil {
			return nil, fmt.Errorf("request page %d: %w", page, err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("page %d: unexpected status %d", page, resp.StatusCode)
		}

		var pageCoins []models.Coin
		if err := json.NewDecoder(resp.Body).Decode(&pageCoins); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("decode page %d: %w", page, err)
		}
		resp.Body.Close()

		if len(pageCoins) == 0 {
			break
		}
		all = append(all, pageCoins...)
	}

	if len(all) > n { // safety slice
		all = all[:n]
	}
	return all, nil
}

func GetStableCoins() (map[string]bool, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://api.coingecko.com/api/v3/coins/markets" +
		"?vs_currency=usd&category=stablecoins&order=market_cap_desc&per_page=250&page=1&sparkline=false"

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var stableCoinsList []models.Coin
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&stableCoinsList)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	stableCoins := make(map[string]bool)
	for _, coin := range stableCoinsList {
		stableCoins[strings.ToLower(coin.Symbol)] = true
	}

	return stableCoins, nil
}
