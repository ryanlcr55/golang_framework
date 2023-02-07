package respositories

type ITransaction interface {
	Begin() any
	Rollback()
	Commit() error
}
