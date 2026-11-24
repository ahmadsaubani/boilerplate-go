package Example

import (
	"database/sql"
)

func (u example) DBStatus() sql.DBStats {
	return u.repo.Example.DBStatus()
}

func (u example) IsExists(yourName string) (res bool) {
	res = yourName != ""
	return
}
