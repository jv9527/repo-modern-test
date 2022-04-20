package http

import (
	"log"
	"net/http"

	"github.com/tokopedia/go-test/domain/campaign/model"
	uCampaign "github.com/tokopedia/go-test/domain/campaign/usecase"
)

type Handler struct {
	usecaseCampaign uCampaign.Usecase
}

func New(usecaseCampaign uCampaign.Usecase) *Handler {
	return &Handler{
		usecaseCampaign: usecaseCampaign,
	}
}

func (h *Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) {
	_, err := h.usecaseCampaign.CreateCampaign(model.CampaignUcCreateCampaignParam{})
	if err != nil {
		log.Println("[CreateCampaign] error occured! message:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) SendCampaign(w http.ResponseWriter, r *http.Request) {
	_, err := h.usecaseCampaign.SendCampaign(model.CampaignUcSendCampaignParam{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
