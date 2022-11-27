package cake

import (
	"database/sql"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pobyzaarif/cake_store/business"
	"github.com/pobyzaarif/cake_store/business/cake"
)

type (
	MySQLRepository struct {
		*sql.DB
	}
)

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db,
	}
}

func (repo *MySQLRepository) Create(ic business.InternalContext, createCake cake.Cake) (cake cake.Cake, err error) {
	cakeQuery := `INSERT INTO cakes (
		id,
		title,
		description,
		rating,
		image,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?)`

	timeNow := time.Now()
	res, err := repo.DB.Exec(cakeQuery,
		createCake.ID,
		createCake.Title,
		createCake.Description,
		createCake.Rating,
		createCake.Image,
		timeNow,
		timeNow,
	)

	if err != nil {
		return
	}

	id, _ := res.LastInsertId()
	createCake.ID = int(id)
	createCake.ObjectMetadata.CreatedAt = timeNow.UTC().Truncate(time.Second)
	createCake.ObjectMetadata.UpdatedAt = timeNow.UTC().Truncate(time.Second)

	cake = createCake

	return
}

func (repo *MySQLRepository) FindAll(ic business.InternalContext) (cakes []cake.Cake, err error) {
	row, err := repo.DB.Query("SELECT * FROM cakes ORDER BY rating DESC, title")
	if err != nil {
		return
	}

	defer row.Close()

	for row.Next() {
		var scanCake cake.Cake

		err := row.Scan(
			&scanCake.ID,
			&scanCake.Title,
			&scanCake.Description,
			&scanCake.Rating,
			&scanCake.Image,
			&scanCake.ObjectMetadata.CreatedAt,
			&scanCake.ObjectMetadata.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		cakes = append(cakes, scanCake)
	}

	spew.Dump(cakes)

	return
}

func (repo *MySQLRepository) FindByID(ic business.InternalContext, id int) (cake cake.Cake, err error) {
	selectQuery := "SELECT * FROM cakes WHERE id = ?"

	err = repo.DB.
		QueryRow(selectQuery, id).
		Scan(
			&cake.ID,
			&cake.Title,
			&cake.Description,
			&cake.Rating,
			&cake.Image,
			&cake.ObjectMetadata.CreatedAt,
			&cake.ObjectMetadata.UpdatedAt,
		)

	if err == sql.ErrNoRows {
		err = business.ErrNotFound
		return
	}

	return
}

func (repo *MySQLRepository) Update(ic business.InternalContext, updateCake cake.Cake) (err error) {
	cakeQuery := `UPDATE cakes 
		SET 
			title = ?,
			description = ?,
			rating = ?,
			image = ?,
			updated_at = ?
	WHERE id = ?`

	timeNow := time.Now()
	_, err = repo.DB.Exec(cakeQuery,
		updateCake.Title,
		updateCake.Description,
		updateCake.Rating,
		updateCake.Image,
		timeNow,
		updateCake.ID,
	)

	return
}

func (repo *MySQLRepository) Delete(ic business.InternalContext, id int) (err error) {
	cakeQuery := "DELETE FROM cakes WHERE id = ?"
	_, err = repo.DB.Exec(cakeQuery, id)

	return
}
