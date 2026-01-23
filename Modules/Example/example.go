package Example

import "database/sql"

func (u example) DBStatus() sql.DBStats {
	return u.repo.Example.DBStatus()
}
