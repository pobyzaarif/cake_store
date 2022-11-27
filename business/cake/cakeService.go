package cake

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/pobyzaarif/cake_store/business"
)

type (
	service struct {
		repository Repository
		validate   *validator.Validate
	}

	Service interface {
		Create(ic business.InternalContext, createCake CakeCreateSpec) (cake Cake, err error)

		FindAll(ic business.InternalContext) (cakes []Cake, err error)

		FindByID(ic business.InternalContext, id int) (cake Cake, err error)

		Update(ic business.InternalContext, updateCake CakeUpdateSpec) (err error)

		Delete(ic business.InternalContext, id int) (err error)
	}
)

func NewService(repository Repository) Service {
	return &service{
		repository,
		validator.New(),
	}
}

func (s *service) Create(ic business.InternalContext, createCake CakeCreateSpec) (cake Cake, err error) {
	err = s.validate.Struct(createCake)
	if err != nil {
		err = business.ErrInvalidSpec
		return
	}

	newCake := Cake{}
	newCake.Title = createCake.Title
	newCake.Description = createCake.Description
	newCake.Rating = createCake.Rating
	newCake.Image = createCake.Image

	return s.repository.Create(ic, newCake)
}

func (s *service) FindAll(ic business.InternalContext) (cakes []Cake, err error) {
	return s.repository.FindAll(ic)
}

func (s *service) FindByID(ic business.InternalContext, id int) (cake Cake, err error) {
	if id == 0 {
		err = business.ErrInvalidSpec
		return
	}

	return s.repository.FindByID(ic, id)
}

func (s *service) Update(ic business.InternalContext, updateCake CakeUpdateSpec) (err error) {
	err = s.validate.Struct(updateCake)
	if err != nil {
		err = business.ErrInvalidSpec
		return
	}

	_, err = s.repository.FindByID(ic, updateCake.ID)
	if err != nil {
		return
	}

	newCake := Cake{}
	newCake.ID = updateCake.ID
	newCake.Title = updateCake.Title
	newCake.Description = updateCake.Description
	newCake.Rating = updateCake.Rating
	newCake.Image = updateCake.Image

	return s.repository.Update(ic, newCake)
}

func (s *service) Delete(ic business.InternalContext, id int) (err error) {
	if id == 0 {
		err = business.ErrInvalidSpec
		return
	}

	_, err = s.repository.FindByID(ic, id)
	if err != nil {
		return
	}

	return s.repository.Delete(ic, id)
}
