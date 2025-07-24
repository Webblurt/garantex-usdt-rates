package grpc

import (
	"context"
	"log"
	"time"

	"github.com/Webblurt/garantex-usdt-rates/internal/rates"
	"github.com/Webblurt/garantex-usdt-rates/internal/storage"
	ratespb "github.com/Webblurt/garantex-usdt-rates/proto/ratespb"
)

type RateServer struct {
	ratespb.UnimplementedRateServiceServer
	storage *storage.Storage
	apiURL  string
}

func NewRateServer(storage *storage.Storage, apiURL string) *RateServer {
	return &RateServer{
		storage: storage,
		apiURL:  apiURL,
	}
}

func (s *RateServer) GetRates(ctx context.Context, req *ratespb.GetRatesRequest) (*ratespb.GetRatesResponse, error) {
	rate, err := rates.FetchRates(s.apiURL)
	if err != nil {
		log.Printf("error fetching rates: %v", err)
		return nil, err
	}

	err = s.storage.SaveRate(rate)
	if err != nil {
		log.Printf("error saving rate: %v", err)
		return nil, err
	}

	return &ratespb.GetRatesResponse{
		Ask:       rate.Ask,
		Bid:       rate.Bid,
		Timestamp: rate.Timestamp.Format(time.RFC3339),
	}, nil
}

func (s *RateServer) HealthCheck(ctx context.Context, req *ratespb.HealthCheckRequest) (*ratespb.HealthCheckResponse, error) {
	return &ratespb.HealthCheckResponse{
		Status:    "OK",
		Timestamp: time.Now().Format(time.RFC3339),
	}, nil
}
