package interfaces

import "github.com/tokopedia/go-test/domain/campaign/model"

type CampaignCacheRepository interface {
	CreateCampaignData(param model.CampaignCacheRepoCreateCampaignDataParam) (model.CampaignCacheRepoCreateCampaignDataResult, error)
	GetCampaignData(param model.CampaignCacheRepoGetCampaignDataParam) (model.CampaignCacheRepoGetCampaignDataResult, error)
}

type CampaignSQLRepository interface {
	CreateCampaignData(param model.CampaignSQLRepoCreateCampaignDataParam) (model.CampaignSQLRepoCreateCampaignDataResult, error)
	GetCampaignData(param model.CampaignSQLRepoGetCampaignDataParam) (model.CampaignSQLRepoGetCampaignDataResult, error)
}
