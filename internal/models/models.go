package models

import (
	"sync"
	"time"
)

type Position struct {
	Symbol           string
	Size             float64
	EntryPrice       float64
	MarkPrice        float64
	LiquidationPrice float64
	MarginRatio      float64
	Margin           float64
	PNL              float64
	PNLPercent       float64
	Isolated         bool
}

type TypePosition struct {
	Flag     bool
	SydeType string
	Coin     string
}

type Config struct {
	TrustCoins   []string
	Interval     string
	DataPoints   int
	ReconnectMin time.Duration
	ReconnectMax time.Duration
}

type SymbolData struct {
	Mu                 sync.Mutex
	ClosingPrices      []float64
	HighPrices         []float64
	LowPrices          []float64
	RealtimePrices     []float64
	HistogramHistory   []float64
	MacdDiffs          []float64
	ActiveSymbolsPrint map[string]SymbolStatus
	Candles            []Candle
}

type SymbolStatus struct {
	LastActive time.Time
	Symbol     string
}

type Candle struct {
	Time   int64
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

type Coin struct {
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	MarketCapRank int     `json:"market_cap_rank"`
	MarketCap     float64 `json:"market_cap"`
}

func (c *Coin) changeSymbol() string {
	return c.Symbol + "USDT"
}
