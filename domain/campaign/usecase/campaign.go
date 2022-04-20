package usecase

import (
	iCampaign "github.com/tokopedia/go-test/domain/campaign/interfaces"
	"github.com/tokopedia/go-test/domain/campaign/model"
	iSegment "github.com/tokopedia/go-test/domain/segmentation/interfaces"
	mSegment "github.com/tokopedia/go-test/domain/segmentation/model"
)

type Usecase struct {
	campaignPostgres iCampaign.CampaignSQLRepository
	campaignRedis    iCampaign.CampaignCacheRepository
	segmentCassandra iSegment.SegmentationDBRepository
}

func New(campaignPostgres iCampaign.CampaignSQLRepository, campaignRedis iCampaign.CampaignCacheRepository, segmentCassanra iSegment.SegmentationDBRepository) *Usecase {
	return &Usecase{
		campaignPostgres: campaignPostgres,
		campaignRedis:    campaignRedis,
		segmentCassandra: segmentCassanra,
	}
}

func (u *Usecase) CreateCampaign(param model.CampaignUcCreateCampaignParam) (model.CampaignUcCreateCampaignResult, error) {
	createCampaignResult, err := u.campaignPostgres.CreateCampaignData(model.CampaignSQLRepoCreateCampaignDataParam{})
	if err != nil {
		return model.CampaignUcCreateCampaignResult{}, nil
	}

	if createCampaignResult.Campaign.NeedSegment {
		segment, err := u.segmentCassandra.GetSegmentData(mSegment.SegmentDBRepoGetSegmentDataParam{})
		if err != nil {
			return model.CampaignUcCreateCampaignResult{}, nil
		}
	}

	return model.CampaignUcCreateCampaignResult{}, nil
}
func (u *Usecase) SendCampaign(param model.CampaignUcSendCampaignParam) (model.CampaignUcSendCampaignResult, error) {
	return model.CampaignUcSendCampaignResult{}, nil
}
