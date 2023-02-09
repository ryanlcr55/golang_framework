package respositories

type ITrxHandler interface {
	Begin() any
	RollBack()
	Commit() error
}
