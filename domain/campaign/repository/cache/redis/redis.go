package redis

import "github.com/tokopedia/go-test/domain/campaign/model"

type Repository struct {
}

func New(addr string) *Repository {
	return &Repository{}
}

func (r *Repository) CreateCampaignData(param model.CampaignCacheRepoCreateCampaignDataParam) (model.CampaignCacheRepoCreateCampaignDataResult, error) {
	return model.CampaignCacheRepoCreateCampaignDataResult{}, nil
}

func (r *Repository) GetCampaignData(param model.CampaignCacheRepoGetCampaignDataParam) (model.CampaignCacheRepoGetCampaignDataResult, error) {
	return model.CampaignCacheRepoGetCampaignDataResult{}, nil
}
