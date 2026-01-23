package Example

import (
	"database/sql"
)

func (u example) DBStatus() sql.DBStats {
	return u.dbCon.PostgreMainCon().Stats()
}
