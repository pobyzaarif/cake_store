package cake_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/pobyzaarif/cake_store/business"
	"github.com/pobyzaarif/cake_store/business/cake"
	cakeRepoMock "github.com/pobyzaarif/cake_store/business/cake/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	cakeRepo       cakeRepoMock.Repository
	cakeService    cake.Service
	ic             business.InternalContext
	cakeResult     cake.Cake
	cakelistResult []cake.Cake
)

func TestMain(m *testing.M) {
	cakeResult = cake.Cake{
		ID:          1,
		Title:       "abc",
		Description: "abc",
		Rating:      7,
		Image:       "http://abc.com",
		ObjectMetadata: business.ObjectMetadata{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	cakelistResult = append(cakelistResult, cakeResult)

	cakeService = cake.NewService(&cakeRepo)

	os.Exit(m.Run())
}

func TestCakeServiceCreate(t *testing.T) {
	createCakeSpec := cake.CakeCreateSpec{
		Title:       "abc",
		Description: "abc",
		Rating:      7,
		Image:       "abc", // this is not valid url,
	}

	t.Run("create failed invalid spec", func(t *testing.T) {
		_, err := cakeService.Create(ic, createCakeSpec)
		assert.NotNil(t, err)
		assert.Equal(t, business.ErrInvalidSpec, err)
	})

	createCakeSpec.Image = "http://abc.com"
	t.Run("create cake success", func(t *testing.T) {
		cakeRepo.On("Create", mock.Anything, mock.Anything).Return(cakeResult, nil).Once()

		_, err := cakeService.Create(ic, createCakeSpec)
		assert.Nil(t, err)
	})
}

func TestCakeServiceFindAll(t *testing.T) {
	t.Run("find all cakes", func(t *testing.T) {
		cakeRepo.On("FindAll", mock.Anything).Return(cakelistResult, nil).Once()

		_, err := cakeService.FindAll(ic)
		assert.Nil(t, err)
	})
}

func TestCakeServiceFindByID(t *testing.T) {
	t.Run("find cake by id failed invalid spec", func(t *testing.T) {
		_, err := cakeService.FindByID(ic, 0)
		assert.NotNil(t, err)
		assert.Equal(t, business.ErrInvalidSpec, err)
	})

	t.Run("find cake by id success", func(t *testing.T) {
		cakeRepo.On("FindByID", mock.Anything, mock.Anything).Return(cakeResult, nil).Once()

		_, err := cakeService.FindByID(ic, 1)
		assert.Nil(t, err)
	})
}

func TestCakeServiceUpdate(t *testing.T) {
	updateCakeSpec := cake.CakeUpdateSpec{
		ID:          1,
		Title:       "abc",
		Description: "abc",
		Rating:      7,
		Image:       "abc", // this is not valid url
	}
	t.Run("update failed invalid spec", func(t *testing.T) {
		err := cakeService.Update(ic, updateCakeSpec)
		assert.NotNil(t, err)
		assert.Equal(t, business.ErrInvalidSpec, err)
	})

	updateCakeSpec.Image = "http://abc.com"
	t.Run("update failed id is not found", func(t *testing.T) {
		cakeRepo.On("FindByID", mock.Anything, mock.Anything).Return(cakeResult, errors.New("data was not found")).Once()

		err := cakeService.Update(ic, updateCakeSpec)
		assert.NotNil(t, err)
	})

	t.Run("update cake success", func(t *testing.T) {
		cakeRepo.On("FindByID", mock.Anything, mock.Anything).Return(cakeResult, nil).Once()
		cakeRepo.On("Update", mock.Anything, mock.Anything).Return(nil).Once()

		err := cakeService.Update(ic, updateCakeSpec)
		assert.Nil(t, err)
	})
}

func TestCakeServiceDelete(t *testing.T) {
	t.Run("delete failed invalid spec", func(t *testing.T) {
		err := cakeService.Delete(ic, 0) // 0 is not valid id
		assert.NotNil(t, err)
		assert.Equal(t, business.ErrInvalidSpec, err)
	})

	t.Run("delete failed id is not found", func(t *testing.T) {
		cakeRepo.On("FindByID", mock.Anything, mock.Anything).Return(cakeResult, errors.New("data was not found")).Once()

		err := cakeService.Delete(ic, 1)
		assert.NotNil(t, err)
	})

	t.Run("delete cake success", func(t *testing.T) {
		cakeRepo.On("FindByID", mock.Anything, mock.Anything).Return(cakeResult, nil).Once()
		cakeRepo.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()

		err := cakeService.Delete(ic, 1)
		assert.Nil(t, err)
	})
}
