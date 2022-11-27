package cake

import (
	"github.com/pobyzaarif/cake_store/business/cake"
	"github.com/pobyzaarif/cake_store/config"
)

//RepositoryFactory Will return business.item.Repository based on active database connection
func RepositoryFactory(dbCon *config.DatabaseConnection) (cakeRepo cake.Repository) {

	if config.GetAPPConfig().DBDriver == "mysql" {
		cakeRepo = NewMySQLRepository(dbCon.MySQLDB)
	}

	return
}
