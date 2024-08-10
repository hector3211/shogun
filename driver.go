package shogun

type Driver string

const (
	POSTGRES Driver = "postgres"
	SQLITE   Driver = "sqlite3"
)

var DefaultDriver = SQLITE

func (d Driver) NewCreateBuilder() *CreateTableBuilder {
	c := newCreateTableBuilder()
	c.SetDriver(d)
	return c
}

func (d Driver) NewInsertBuilder() *InsertBuilder {
	i := newInsertBuilder()
	i.SetDriver(d)
	return i
}

func (d Driver) NewSelectBuilder() *SelectBuilder {
	s := newSelectbuilder()
	s.SetDriver(d)
	return s
}

func (d Driver) NewUpdateBuilder() *UpdateBuilder {
	u := newUpdateBuilder()
	u.SetDriver(d)
	return u
}
