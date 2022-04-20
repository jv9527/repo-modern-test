package model

type CampaignUcCreateCampaignParam struct {
}

type CampaignUcCreateCampaignResult struct {
}

type CampaignUcSendCampaignParam struct {
}

type CampaignUcSendCampaignResult struct {
}

type CampaignCacheRepoCreateCampaignDataParam struct {
}

type CampaignCacheRepoCreateCampaignDataResult struct {
}

type CampaignCacheRepoGetCampaignDataParam struct {
}

type CampaignCacheRepoGetCampaignDataResult struct {
}

type CampaignSQLRepoCreateCampaignDataParam struct {
}

type CampaignSQLRepoCreateCampaignDataResult struct {
	Campaign Campaign
}

type CampaignSQLRepoGetCampaignDataParam struct {
}

type CampaignSQLRepoGetCampaignDataResult struct {
}

type Campaign struct {
	NeedSegment bool
}
