package cake

import "github.com/pobyzaarif/cake_store/business"

type (
	Cake struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Rating      int    `json:"rating"`
		Image       string `json:"image"`

		business.ObjectMetadata
	}

	CakeCreateSpec struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
		Rating      int    `json:"rating" validate:"number"`
		Image       string `json:"image" validate:"url"`
	}

	CakeUpdateSpec struct {
		ID          int    `json:"id" validate:"required"`
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
		Rating      int    `json:"rating" validate:"number"`
		Image       string `json:"image" validate:"url"`
	}
)
