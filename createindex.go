package shogun

import "fmt"

type CreateIndexQuery interface {
	Index(indexName string) *CreateIndexBuilder
	On(tableName, field string) *CreateIndexBuilder
	String() string
	Build() string
	SetDriver(sqlDriver Driver) *CreateIndexBuilder
	GetDriver() string
}

type CreateIndexBuilder struct {
	driver    Driver
	action    string
	name      string
	tableName string
	field     string
}

func NewIndexBuilder() *CreateIndexBuilder {
	return DefaultDriver.NewIndexBuilder()
}

func newIndexBuilder() *CreateIndexBuilder {
	return &CreateIndexBuilder{
		driver: DefaultDriver,
		action: "CREATE INDEX",
	}
}

func Index(indexName string) *CreateIndexBuilder {
	return newIndexBuilder().Index(indexName)
}

func (i *CreateIndexBuilder) Index(indexName string) *CreateIndexBuilder {
	i.name = indexName
	return i
}

func (i *CreateIndexBuilder) On(tableName, field string) *CreateIndexBuilder {
	i.tableName = tableName
	i.field = field
	return i
}

func (i CreateIndexBuilder) String() string {
	return i.Build()
}

func (i *CreateIndexBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString(i.action + " ")
	buf.WriteLeadingString(i.name + " ")
	buf.WriteLeadingString("ON ")
	buf.WriteString(fmt.Sprintf("%s(%s)", i.tableName, i.field))
	buf.WriteString(";")

	return buf.String()
}

// Sets a new driver
func (i *CreateIndexBuilder) SetDriver(sqlDriver Driver) *CreateIndexBuilder {
	i.driver = sqlDriver
	return i
}

// Returns current driver being used
func (i CreateIndexBuilder) GetDriver() Driver {
	return i.driver
}
