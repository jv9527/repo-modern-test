package main

import (
	"github.com/gorilla/mux"
	dCampaign "github.com/tokopedia/go-test/domain/campaign/delivery/http"
)

func initHttpRouter(campaignHandler dCampaign.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/create/campaign", campaignHandler.CreateCampaign)
	r.HandleFunc("/api/v1/send/campaign", campaignHandler.SendCampaign)

	return r
}
