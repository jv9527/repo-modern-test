package scylla

import "github.com/tokopedia/go-test/domain/segmentation/model"

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) GetSegmentData(param model.SegmentDBRepoGetSegmentDataParam) (model.SegmentDBRepoGetSegmentDataResult, error) {
	return model.SegmentDBRepoGetSegmentDataResult{}, nil
}
