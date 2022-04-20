package postgre

import "github.com/tokopedia/go-test/domain/campaign/model"

type Repository struct {
}

func New(connString string) *Repository {
	return &Repository{}
}

func (r *Repository) CreateCampaignData(param model.CampaignSQLRepoCreateCampaignDataParam) (model.CampaignSQLRepoCreateCampaignDataResult, error) {
	return model.CampaignSQLRepoCreateCampaignDataResult{}, nil
}

func (r *Repository) GetCampaignData(param model.CampaignSQLRepoGetCampaignDataParam) (model.CampaignSQLRepoGetCampaignDataResult, error) {
	return model.CampaignSQLRepoGetCampaignDataResult{}, nil
}
