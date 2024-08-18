package shogun

import "fmt"

type TransactionQuery interface {
	String() string
	Build() string
	SetDriver(sqlDriver Driver) *TransactionBuilder
	GetDriver() string
}

type TransactionBuilder struct {
	driver   Driver
	action   string
	update   []*UpdateBuilder
	insert   []*InsertBuilder
	commit   bool
	rollback bool
}

func NewTransactionBuilder() *TransactionBuilder {
	return DefaultDriver.NewTransactionBuilder()
}

func newTransactionBuilder() *TransactionBuilder {
	return &TransactionBuilder{
		driver: DefaultDriver,
		action: "BEGIN TRANSACTION",
		update: make([]*UpdateBuilder, 0),
		insert: make([]*InsertBuilder, 0),
	}
}

// Sets the update statement returning new instance of TransactionBuilder
func UpdateTransaction(update *UpdateBuilder) *TransactionBuilder {
	return NewTransactionBuilder().UpdateTransaction(update)
}

// Sets the Insert statement returning new instance of TransactionBuilder
func InsertTransaction(insert *InsertBuilder) *TransactionBuilder {
	return NewTransactionBuilder().InsertTransaction(insert)
}

// Loads up an update statement
func (t *TransactionBuilder) UpdateTransaction(update *UpdateBuilder) *TransactionBuilder {
	t.update = append(t.update, update)
	return t
}

// Loads up an Insert statement
func (t *TransactionBuilder) InsertTransaction(insert *InsertBuilder) *TransactionBuilder {
	t.insert = append(t.insert, insert)
	return t
}

// Sets COMMIT
func (t *TransactionBuilder) Commit() *TransactionBuilder {
	t.commit = true
	return t
}

// Sets ROLLBACK
func (t *TransactionBuilder) RollBack() *TransactionBuilder {
	t.rollback = true
	return t
}

// Sets a new driver
func (t *TransactionBuilder) SetDriver(sqlDriver Driver) *TransactionBuilder {
	t.driver = sqlDriver
	return t
}

// Returns current driver being used
func (t TransactionBuilder) GetDriver() Driver {
	return t.driver
}

// Returns the query in a string format
func (t TransactionBuilder) String() string {
	return t.Build()
}

// Builds out the final query
func (t TransactionBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString(fmt.Sprintf("%s; ", t.action))

	if len(t.update) > 0 {
		for i := 0; i < len(t.update); i++ {
			query := t.update[i]
			buf.WriteLeadingString(query.Build() + " ")
		}
	}

	if len(t.insert) > 0 {
		for i := 0; i < len(t.insert); i++ {
			query := t.insert[i]
			if i >= 1 {
				buf.WriteLeadingString(query.Build() + " ")
			} else {
				buf.WriteString(query.Build())

			}
		}
	}

	if t.commit == true {
		buf.WriteString(" COMMIT;")
	}

	if t.rollback == true {
		buf.WriteString(" ROLLBACK;")
	}

	return buf.String()
}
