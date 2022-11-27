package cake

import (
	"github.com/pobyzaarif/cake_store/business"
)

type Repository interface {
	Create(ic business.InternalContext, createCake Cake) (cake Cake, err error)

	FindAll(ic business.InternalContext) (cakes []Cake, err error)

	FindByID(ic business.InternalContext, id int) (cake Cake, err error)

	Update(ic business.InternalContext, updateCake Cake) (err error)

	Delete(ic business.InternalContext, id int) (err error)
}
