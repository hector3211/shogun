package shogun

type Driver string

const (
	POSTGRES Driver = "postgres"
	SQLITE   Driver = "sqlite3"
)

var DefaultDriver = SQLITE

func (d Driver) NewCreateBuilder() *CreateTableBuilder {
	t := newCreateTableBuilder()
	t.SetDriver(d)
	return t
}

func (d Driver) NewIndexBuilder() *CreateIndexBuilder {
	t := newIndexBuilder()
	t.SetDriver(d)
	return t
}

func (d Driver) NewInsertBuilder() *InsertBuilder {
	t := newInsertBuilder()
	t.SetDriver(d)
	return t
}

func (d Driver) NewSelectBuilder() *SelectBuilder {
	t := newSelectbuilder()
	t.SetDriver(d)
	return t
}

func (d Driver) NewUpdateBuilder() *UpdateBuilder {
	t := newUpdateBuilder()
	t.SetDriver(d)
	return t
}

func (d Driver) NewDeleteBuilder() *DeleteBuilder {
	t := newDeleteBuilder()
	t.SetDriver(d)
	return t
}

func (d Driver) NewJoinBuilder() *JoinBuilder {
	t := newJoinBuilder()
	t.SetDriver(d)
	return t
}

func (d Driver) NewTransactionBuilder() *TransactionBuilder {
	t := newTransactionBuilder()
	t.SetDriver(d)
	return t
}
