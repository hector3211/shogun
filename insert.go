package shogun

type InsertBuilder struct {
	Action    string
	TableName string
	Columns   []string
	Values    [][]string
}

// Creates a new instance of the InsertBuilder struct
func NewInsertBuilder() *InsertBuilder {
	return &InsertBuilder{
		Action: "INSERT",
	}
}

// Sets table name
func (i *InsertBuilder) Table(tableName string) *InsertBuilder {
	i.TableName = tableName
	return i
}

func (i *InsertBuilder) Cols(Columns ...string) *InsertBuilder {
	i.Columns = append(i.Columns, Columns...)
	return i
}

func (i *InsertBuilder) Vals(values ...string) *InsertBuilder {
	i.Values = append(i.Values, values)
	return i
}
