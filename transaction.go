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
	return &TransactionBuilder{
		driver: DefaultDriver,
		action: "BEGIN TRANSACTION",
	}
}

func InsertTransaction(insert *InsertBuilder) *TransactionBuilder {
	return NewTransactionBuilder().InsertTransaction(insert)
}

func UpdateTransaction(update *UpdateBuilder) *TransactionBuilder {
	return NewTransactionBuilder().UpdateTransaction(update)
}

func (t *TransactionBuilder) InsertTransaction(insert *InsertBuilder) *TransactionBuilder {
	t.insert = append(t.insert, insert)
	return t
}

func (t *TransactionBuilder) UpdateTransaction(update *UpdateBuilder) *TransactionBuilder {
	t.update = append(t.update, update)
	return t
}

func (t *TransactionBuilder) Commit() *TransactionBuilder {
	t.commit = true
	return t
}

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

func (t TransactionBuilder) String() string {
	return t.Build()
}

func (t TransactionBuilder) Build() string {
	buf := newStringBuilder()
	buf.WriteLeadingString(fmt.Sprintf("%s; ", t.action))

	if len(t.update) > 0 {
		for _, query := range t.update {
			buf.WriteString(query.Build())
		}
	}

	buf.WriteString(" ")
	if len(t.insert) > 0 {
		for _, query := range t.insert {
			buf.WriteString(query.Build())
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
