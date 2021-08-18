package diaries

import "gin-api/domain/diaries"

type Interactor interface {
	Add(in AddInput) (err error)
	Get(in GetInput) (out GetInput, err error)
	GetList(in GetListInput) (out GetListInput, err error)
	Update(in UpdateInput) (err error)
}

type interactor struct {
	repository diaries.Repository
}

func NewInteractor(repository diaries.Repository) Interactor {
	return &interactor{
		repository: repository,
	}
}

func (i *interactor) Get(in GetInput) (out GetInput, err error) {
	panic("implement me")
}

func (i *interactor) GetList(in GetListInput) (out GetListInput, err error) {
	panic("implement me")
}

func (i *interactor) Update(in UpdateInput) (err error) {
	panic("implement me")
}

func (i *interactor) Add(in AddInput) (err error) {
	err = add(i.repository, in)
	return
}
