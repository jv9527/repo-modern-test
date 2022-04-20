package interfaces

import "github.com/tokopedia/go-test/domain/segmentation/model"

type SegmentationDBRepository interface {
	GetSegmentData(param model.SegmentDBRepoGetSegmentDataParam) (model.SegmentDBRepoGetSegmentDataResult, error)
}
