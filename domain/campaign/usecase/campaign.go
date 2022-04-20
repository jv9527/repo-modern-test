package usecase

import (
	iCampaign "github.com/tokopedia/go-test/domain/campaign/interfaces"
	"github.com/tokopedia/go-test/domain/campaign/model"
	iSegment "github.com/tokopedia/go-test/domain/segmentation/interfaces"
)

type Usecase struct {
	campaignSQL   iCampaign.CampaignSQLRepository
	campaignCache iCampaign.CampaignCacheRepository
	segmentDB     iSegment.SegmentationDBRepository
}

func New(campaignSQL iCampaign.CampaignSQLRepository, campaignCache iCampaign.CampaignCacheRepository, segmentDB iSegment.SegmentationDBRepository) *Usecase {
	return &Usecase{
		campaignSQL:   campaignSQL,
		campaignCache: campaignCache,
		segmentDB:     segmentDB,
	}
}

func (u *Usecase) CreateCampaign(param model.CampaignUcCreateCampaignParam) (model.CampaignUcCreateCampaignResult, error) {
	createCampaignResult, err := u.campaignSQL.CreateCampaignData(model.CampaignSQLRepoCreateCampaignDataParam{})
	if err != nil {
		return model.CampaignUcCreateCampaignResult{}, nil
	}

	if createCampaignResult.Campaign.NeedSegment {
		//segment, err := u.segmentDB.GetSegmentData(mSegment.SegmentDBRepoGetSegmentDataParam{})
		if err != nil {
			return model.CampaignUcCreateCampaignResult{}, nil
		}
	}

	return model.CampaignUcCreateCampaignResult{}, nil
}
func (u *Usecase) SendCampaign(param model.CampaignUcSendCampaignParam) (model.CampaignUcSendCampaignResult, error) {
	return model.CampaignUcSendCampaignResult{}, nil
}
