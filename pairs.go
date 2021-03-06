package finance

import (
	"fmt"
	"strings"
)

const (
	// USDGBP pair.
	USDGBP = "USDGBP=X"
	// USDEUR pair.
	USDEUR = "USDEUR=X"
	// USDAUD pair.
	USDAUD = "USDAUD=X"
	// USDCHF pair.
	USDCHF = "USDCHF=X"
	// USDJPY pair.
	USDJPY = "USDJPY=X"
	// USDCAD pair.
	USDCAD = "USDCAD=X"
	// USDSGD pair.
	USDSGD = "USDSGD=X"
	// USDNZD pair.
	USDNZD = "USDNZD=X"
	// USDHKD pair.
	USDHKD = "USDHKD=X"
	// GBPUSD pair.
	GBPUSD = "GBPUSD=X"
	// GBPEUR pair.
	GBPEUR = "GBPEUR=X"
	// GBPAUD pair.
	GBPAUD = "GBPAUD=X"
	// GBPCHF pair.
	GBPCHF = "GBPCHF=X"
	// GBPJPY pair.
	GBPJPY = "GBPJPY=X"
	// GBPCAD pair.
	GBPCAD = "GBPCAD=X"
	// GBPSGD pair.
	GBPSGD = "GBPSGD=X"
	// GBPNZD pair.
	GBPNZD = "GBPNZD=X"
	// GBPHKD pair.
	GBPHKD = "GBPHKD=X"
	// EURUSD pair.
	EURUSD = "EURUSD=X"
	// EURGBP pair.
	EURGBP = "EURGBP=X"
	// EURAUD pair.
	EURAUD = "EURAUD=X"
	// EURCHF pair.
	EURCHF = "EURCHF=X"
	// EURJPY pair.
	EURJPY = "EURJPY=X"
	// EURCAD pair.
	EURCAD = "EURCAD=X"
	// EURSGD pair.
	EURSGD = "EURSGD=X"
	// EURNZD pair.
	EURNZD = "EURNZD=X"
	// EURHKD pair.
	EURHKD = "EURHKD=X"
	// AUDUSD pair.
	AUDUSD = "AUDUSD=X"
	// AUDGBP pair.
	AUDGBP = "AUDGBP=X"
	// AUDEUR pair.
	AUDEUR = "AUDEUR=X"
	// AUDCHF pair.
	AUDCHF = "AUDCHF=X"
	// AUDJPY pair.
	AUDJPY = "AUDJPY=X"
	// AUDCAD pair.
	AUDCAD = "AUDCAD=X"
	// AUDSGD pair.
	AUDSGD = "AUDSGD=X"
	// AUDNZD pair.
	AUDNZD = "AUDNZD=X"
	// AUDHKD pair.
	AUDHKD = "AUDHKD=X"
	// CHFGBP pair.
	CHFGBP = "CHFGBP=X"
	// CHFEUR pair.
	CHFEUR = "CHFEUR=X"
	// CHFAUD pair.
	CHFAUD = "CHFAUD=X"
	// CHFJPY pair.
	CHFJPY = "CHFJPY=X"
	// CHFCAD pair.
	CHFCAD = "CHFCAD=X"
	// CHFSGD pair.
	CHFSGD = "CHFSGD=X"
	// CHFNZD pair.
	CHFNZD = "CHFNZD=X"
	// CHFHKD pair.
	CHFHKD = "CHFHKD=X"
	// JPYUSD pair.
	JPYUSD = "JPYUSD=X"
	// JPYGBP pair.
	JPYGBP = "JPYGBP=X"
	// JPYEUR pair.
	JPYEUR = "JPYEUR=X"
	// JPYAUD pair.
	JPYAUD = "JPYAUD=X"
	// JPYCHF pair.
	JPYCHF = "JPYCHF=X"
	// JPYCAD pair.
	JPYCAD = "JPYCAD=X"
	// JPYSGD pair.
	JPYSGD = "JPYSGD=X"
	// JPYNZD pair.
	JPYNZD = "JPYNZD=X"
	// JPYHKD pair.
	JPYHKD = "JPYHKD=X"
	// CADUSD pair.
	CADUSD = "CADUSD=X"
	// CADGBP pair.
	CADGBP = "CADGBP=X"
	// CADEUR pair.
	CADEUR = "CADEUR=X"
	// CADAUD pair.
	CADAUD = "CADAUD=X"
	// CADCHF pair.
	CADCHF = "CADCHF=X"
	// CADJPY pair.
	CADJPY = "CADJPY=X"
	// CADSGD pair.
	CADSGD = "CADSGD=X"
	// CADNZD pair.
	CADNZD = "CADNZD=X"
	// CADHKD pair.
	CADHKD = "CADHKD=X"
	// SGDUSD pair.
	SGDUSD = "SGDUSD=X"
	// SGDGBP pair.
	SGDGBP = "SGDGBP=X"
	// SGDEUR pair.
	SGDEUR = "SGDEUR=X"
	// SGDAUD pair.
	SGDAUD = "SGDAUD=X"
	// SGDCHF pair.
	SGDCHF = "SGDCHF=X"
	// SGDJPY pair.
	SGDJPY = "SGDJPY=X"
	// SGDCAD pair.
	SGDCAD = "SGDCAD=X"
	// SGDNZD pair.
	SGDNZD = "SGDNZD=X"
	// SGDHKD pair.
	SGDHKD = "SGDHKD=X"
	// NZDUSD pair.
	NZDUSD = "NZDUSD=X"
	// NZDGBP pair.
	NZDGBP = "NZDGBP=X"
	// NZDEUR pair.
	NZDEUR = "NZDEUR=X"
	// NZDAUD pair.
	NZDAUD = "NZDAUD=X"
	// NZDCHF pair.
	NZDCHF = "NZDCHF=X"
	// NZDJPY pair.
	NZDJPY = "NZDJPY=X"
	// NZDCAD pair.
	NZDCAD = "NZDCAD=X"
	// NZDSGD pair.
	NZDSGD = "NZDSGD=X"
	// NZDHKD pair.
	NZDHKD = "NZDHKD=X"
	// HKDUSD pair.
	HKDUSD = "HKDUSD=X"
	// HKDGBP pair.
	HKDGBP = "HKDGBP=X"
	// HKDEUR pair.
	HKDEUR = "HKDEUR=X"
	// HKDAUD pair.
	HKDAUD = "HKDAUD=X"
	// HKDCHF pair.
	HKDCHF = "HKDCHF=X"
	// HKDJPY pair.
	HKDJPY = "HKDJPY=X"
	// HKDCAD pair.
	HKDCAD = "HKDCAD=X"
	// HKDSGD pair.
	HKDSGD = "HKDSGD=X"
	// HKDNZD pair.
	HKDNZD = "HKDNZD=X"
)

// GetCurrencyPairQuote fetches a single currency pair's quote from Yahoo Finance.
func GetCurrencyPairQuote(symbol string) (*FXPairQuote, error) {

	params := map[string]string{
		"s": symbol,
		"f": strings.Join(fXFields[:], ""),
		"e": ".csv",
	}

	table, err := getPairsQuotesTable(buildURL(quoteURL, params))
	if err != nil {
		return nil, err
	}
	return generatePairQuotes(table)[0], nil
}

// getPairsQuotesTable fetches the pairs data table from the endpoint.
func getPairsQuotesTable(url string) ([][]string, error) {

	table, err := requestCSV(url)
	if err != nil {
		return nil, fmt.Errorf("request pairs table error:  (error was: %s)\n", err.Error())
	}
	return table, nil
}

// generatePairQuotes turns the raw table data of a pair quote into proper pair quote structs.
func generatePairQuotes(table [][]string) (pairs []*FXPairQuote) {

	for _, row := range table {
		pairs = append(pairs, newFXPairQuote(row))
	}
	return pairs
}
