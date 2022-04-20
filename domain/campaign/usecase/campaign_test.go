package usecase

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/tokopedia/go-test/domain/campaign/interfaces"
	mockCampaign "github.com/tokopedia/go-test/domain/campaign/interfaces/mock"
	"github.com/tokopedia/go-test/domain/campaign/model"
	iSegment "github.com/tokopedia/go-test/domain/segmentation/interfaces"
	mockSegmentation "github.com/tokopedia/go-test/domain/segmentation/interfaces/mock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSegmentDBRepo := mockSegmentation.NewMockSegmentationDBRepository(ctrl)
	mockCampaignCacheRepo := mockCampaign.NewMockCampaignCacheRepository(ctrl)
	mockCampaignSQLRepo := mockCampaign.NewMockCampaignSQLRepository(ctrl)

	type args struct {
		campaignPostgres interfaces.CampaignSQLRepository
		campaignRedis    interfaces.CampaignCacheRepository
		segmentCassanra  iSegment.SegmentationDBRepository
	}
	tests := []struct {
		name  string
		args  args
		isNil bool
	}{
		{
			name: "test case 1",
			args: args{
				campaignPostgres: mockCampaignSQLRepo,
				campaignRedis:    mockCampaignCacheRepo,
				segmentCassanra:  mockSegmentDBRepo,
			},
			isNil: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.campaignPostgres, tt.args.campaignRedis, tt.args.segmentCassanra) == nil; !reflect.DeepEqual(got, tt.isNil) {
				t.Errorf("New() = %v, isNil %v", got, tt.isNil)
			}
		})
	}
}

func TestUsecase_CreateCampaign(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSegmentDBRepo := mockSegmentation.NewMockSegmentationDBRepository(ctrl)
	mockCampaignCacheRepo := mockCampaign.NewMockCampaignCacheRepository(ctrl)
	mockCampaignSQLRepo := mockCampaign.NewMockCampaignSQLRepository(ctrl)

	u := &Usecase{
		campaignPostgres: mockCampaignSQLRepo,
		campaignRedis:    mockCampaignCacheRepo,
		segmentCassandra: mockSegmentDBRepo,
	}

	type args struct {
		param model.CampaignUcCreateCampaignParam
	}
	tests := []struct {
		name    string
		args    args
		want    model.CampaignUcCreateCampaignResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := u.CreateCampaign(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.CreateCampaign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.CreateCampaign() = %v, want %v", got, tt.want)
			}
		})
	}
}
