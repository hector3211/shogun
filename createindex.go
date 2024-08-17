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

// creates new instance of CreateIndexBuilder
func NewIndexBuilder() *CreateIndexBuilder {
	return DefaultDriver.NewIndexBuilder()
}

// creates new instance of CreateIndexBuilder
func newIndexBuilder() *CreateIndexBuilder {
	return &CreateIndexBuilder{
		driver: DefaultDriver,
		action: "CREATE INDEX",
	}
}

// Sets index name returning a new instance of CreateIndexBuilder
func Index(indexName string) *CreateIndexBuilder {
	return newIndexBuilder().Index(indexName)
}

// Sets index name
func (i *CreateIndexBuilder) Index(indexName string) *CreateIndexBuilder {
	i.name = indexName
	return i
}

// Sets the table and field associated with the pending index
func (i *CreateIndexBuilder) On(tableName, field string) *CreateIndexBuilder {
	i.tableName = tableName
	i.field = field
	return i
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

// Returns query in a string format
func (i CreateIndexBuilder) String() string {
	return i.Build()
}

// Builds out the final query
func (i *CreateIndexBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteString(fmt.Sprintf("%s %s ON %s(%s);", i.action, i.name, i.tableName, i.field))

	return buf.String()
}
