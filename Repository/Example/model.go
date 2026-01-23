package Example

import (
	"boilerplate-go/Config"
	"database/sql"
)

type Repository interface {
	repositoryRedis
	DBStatus() sql.DBStats
}

// CONSTRUCTOR STRUCT
type example struct {
	dbCon Config.DbConInterface
}

// CONSTRUCTOR FUNCTION FOR USER REPOSITORY
func RepositoryNew(dbCon Config.DbConInterface) Repository {
	return &example{
		dbCon: dbCon,
	}
}
