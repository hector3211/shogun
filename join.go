package shogun

import (
	"fmt"
	"strings"
)

type Join int

const (
	NOJOIN Join = iota
	INNER
	FULL
	LEFT
	RIGHT
	CROSS
)

func (j Join) String() string {
	switch j {
	case INNER:
		return "INNER"
	case FULL:
		return "FULL"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	case CROSS:
		return "CROSS"
	default:
		return ""
	}
}

type Table struct {
	name   string
	fields []string
}

type JoinQuery interface {
	JSelect(tableName, targetField string) *JoinBuilder
	JFrom(tableName string) *JoinBuilder
	Join(typeOfJoin Join, tableName string) *JoinBuilder
	OnCondition(tableNameA, tableFieldA string, condition ConditionToken, tableNameB, tableFieldB string, arg interface{}) *JoinBuilder
	Equal() *JoinBuilder
	NotEqual() *JoinBuilder
	And() *JoinBuilder
	String() string
	Build() string
	SetDriver(sqlDriver Driver) *JoinBuilder
	GetDriver() string
}

type JoinBuilder struct {
	driver         Driver
	typeOfJoin     Join
	selectedTables []Table
	fromTable      string
	joinTable      string
	and            bool
	condition      ConditionToken
	conditionStmts []string
}

// Creates a new instance of the JoinBuilder struct
func NewJoinBuilder() *JoinBuilder {
	return &JoinBuilder{
		driver:         DefaultDriver,
		typeOfJoin:     NOJOIN,
		selectedTables: make([]Table, 0),
		conditionStmts: make([]string, 0),
	}
}

// Creates a new instance of the JoinBuilder struct
func newJoinBuilder() *JoinBuilder {
	return &JoinBuilder{
		driver:         DefaultDriver,
		typeOfJoin:     NOJOIN,
		selectedTables: make([]Table, 0),
		conditionStmts: make([]string, 0),
	}
}

// Sets the tables and fields that query will select returning new instance of JoinBuilder
func JSelect(tableName, targetField string) *JoinBuilder {
	return newJoinBuilder().JSelect(tableName, targetField)
}

// Sets the tables and fields that query will select
func (j *JoinBuilder) JSelect(tableName, targetField string) *JoinBuilder {
	if tableExists(j.selectedTables, tableName) {
		j.selectedTables = addTableField(j.selectedTables, tableName, targetField)
	} else {
		j.selectedTables = append(j.selectedTables, Table{name: tableName, fields: []string{targetField}})
	}
	return j
}

// Sets the table that query will target
func (j *JoinBuilder) JFrom(tableName string) *JoinBuilder {
	j.fromTable = tableName
	return j
}

// Sets the join table
func (j *JoinBuilder) Join(typeOfJoin Join, tableName string) *JoinBuilder {
	j.typeOfJoin = typeOfJoin
	j.joinTable = tableName
	return j
}

// Loads a query condition
func (j *JoinBuilder) OnCondition(tableNameA, tableFieldA string, condition ConditionToken, tableNameB, tableFieldB string, arg interface{}) *JoinBuilder {
	// If no table B provided then we compare with arg
	if tableNameB == "" && tableFieldB == "" {
		if arg != "" {
			var argFormat string
			switch v := arg.(type) {
			case string:
				argFormat = fmt.Sprintf("'%s'", v)
			case int:
				argFormat = fmt.Sprintf("%d", v)
			case float32:
				argFormat = fmt.Sprintf("%f", v)
			case bool:
				argFormat = strings.ToUpper(fmt.Sprintf("%v", v))
			}
			j.conditionStmts = append(j.conditionStmts, fmt.Sprintf("%s.%s %s %s", tableNameA, tableFieldA, condition.String(), argFormat))
			return j
		}
	}

	j.conditionStmts = append(j.conditionStmts, fmt.Sprintf("%s.%s %s %s.%s", tableNameA, tableFieldA, condition.String(), tableNameB, tableFieldB))
	return j
}

// Sets join condition to "="
func (j *JoinBuilder) Equal() *JoinBuilder {
	j.condition = EQUAL
	return j
}

// Sets join condition to "!="
func (j *JoinBuilder) NotEqual() *JoinBuilder {
	j.condition = NOTEQUAL
	return j
}

// Toggles the ability to have multiple conditions
func (j *JoinBuilder) And() *JoinBuilder {
	j.and = true
	return j
}

// Sets a new driver
func (j *JoinBuilder) SetDriver(sqlDriver Driver) *JoinBuilder {
	j.driver = sqlDriver
	return j
}

// Returns current driver being used
func (j JoinBuilder) GetDriver() Driver {
	return j.driver
}

// Returns the query in a string format
func (j JoinBuilder) String() string {
	return j.Build()
}

// Builds out the final query
func (j *JoinBuilder) Build() string {
	buf := newStringBuilder()

	// Select
	if len(j.selectedTables) > 0 {
		buf.WriteLeadingString("SELECT ")

		for i := 0; i < len(j.selectedTables); i++ {
			currentTable := j.selectedTables[i]

			for j := 0; j < len(currentTable.fields); j++ {
				currentValue := currentTable.fields[j]
				buf.WriteString(fmt.Sprintf("%s.%s", currentTable.name, currentValue))
				if j < len(currentTable.fields)-1 {
					buf.WriteString(",")
				}
			}
			if i < len(j.selectedTables)-1 {
				buf.WriteString(",")
			}

		}
	}

	// FROM
	if len(j.fromTable) != 0 {
		buf.WriteLeadingString("FROM ")
		buf.WriteString(string(j.fromTable))
	}

	// Join
	if j.typeOfJoin != NOJOIN {
		buf.WriteLeadingString(fmt.Sprintf("%s JOIN %s ", j.typeOfJoin.String(), j.joinTable))
	}

	// On
	if len(j.conditionStmts) > 0 {
		buf.WriteLeadingString("ON ")
		for idx, con := range j.conditionStmts {
			buf.WriteString(con)
			if idx < len(j.conditionStmts)-1 && j.and {
				buf.WriteLeadingString("AND ")
			}
		}
	}
	buf.WriteString(";")
	return buf.String()
}

func tableExists(tables []Table, tableName string) bool {
	for _, t := range tables {
		if t.name == tableName {
			return true
		}
	}

	return false
}

func addTableField(table []Table, tableName, newField string) []Table {
	for i := 0; i < len(table); i++ {
		if table[i].name == tableName {
			table[i].fields = append(table[i].fields, newField)
		}
	}

	return table
}
