package interfaces

import "github.com/tokopedia/go-test/domain/campaign/model"

type CampaignUsecase interface {
	CreateCampaign(param model.CampaignUcCreateCampaignParam) (model.CampaignUcCreateCampaignResult, error)
	SendCampaign(param model.CampaignUcSendCampaignParam) (model.CampaignUcSendCampaignResult, error)
}
