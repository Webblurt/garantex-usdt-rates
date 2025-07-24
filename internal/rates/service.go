package rates

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type Rate struct {
	Ask       float64
	Bid       float64
	Timestamp time.Time
}

func FetchRates(apiURL string) (*Rate, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("non-200 response from Grinex API")
	}

	var data struct {
		Asks [][]string `json:"asks"`
		Bids [][]string `json:"bids"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data.Asks) == 0 || len(data.Bids) == 0 {
		return nil, errors.New("no ask or bid data")
	}

	ask, err := strconv.ParseFloat(data.Asks[0][0], 64)
	if err != nil {
		return nil, err
	}
	bid, err := strconv.ParseFloat(data.Bids[0][0], 64)
	if err != nil {
		return nil, err
	}

	return &Rate{
		Ask:       ask,
		Bid:       bid,
		Timestamp: time.Now().UTC(),
	}, nil
}
