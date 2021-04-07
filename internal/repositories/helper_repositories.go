package repositories

// RepoReader is an interface to override sql.Rows.Scan and sql.Row.Scan
type RepoReader interface {
	Scan(dest ...interface{}) error
}
