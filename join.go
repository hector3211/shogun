package shogun

import (
	"fmt"
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

type JoinBuilder struct {
	driver         Driver
	typeOfJoin     Join
	selectedTables []Table
	fromTable      string
	joinTable      string
	onSelected     []Table
	condition      ConditionToken
}

func NewJoinBuilder() *JoinBuilder {
	return &JoinBuilder{
		driver:         DefaultDriver,
		typeOfJoin:     NOJOIN,
		selectedTables: make([]Table, 0),
		onSelected:     make([]Table, 0),
	}
}

func newJoinBuilder() *JoinBuilder {
	return &JoinBuilder{
		driver:         DefaultDriver,
		typeOfJoin:     NOJOIN,
		selectedTables: make([]Table, 0),
		onSelected:     make([]Table, 0),
	}
}

func JSelect(tableName, targetField string) *JoinBuilder {
	return newJoinBuilder().JSelect(tableName, targetField)
}

// TODO: doesnt read all fieds from the same table
func (j *JoinBuilder) JSelect(tableName, targetField string) *JoinBuilder {
	if tableExists(j.selectedTables, tableName) {
		j.selectedTables = addTableField(j.selectedTables, tableName, targetField)
	} else {
		j.selectedTables = append(j.selectedTables, Table{name: tableName, fields: []string{targetField}})
	}
	return j
}

func (j *JoinBuilder) JFrom(tableName string) *JoinBuilder {
	j.fromTable = tableName
	return j
}

func (j *JoinBuilder) Join(typeOfJoin Join, tableName string) *JoinBuilder {
	j.typeOfJoin = typeOfJoin
	j.joinTable = tableName
	return j
}

func (j *JoinBuilder) OnTable(tableName, targetField string) *JoinBuilder {
	if tableExists(j.onSelected, tableName) == true {
		j.onSelected = addTableField(j.onSelected, tableName, targetField)
	} else {
		j.onSelected = append(j.onSelected, Table{name: tableName, fields: []string{targetField}})
	}
	return j
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
	for _, t := range table {
		if t.name == tableName {
			t.fields = append(t.fields, newField)
		}
	}

	return table
}

func (j JoinBuilder) String() string {
	return j.Build()
}

func (j *JoinBuilder) Build() string {
	buf := newStringBuilder()
	if len(j.selectedTables) > 0 {
		buf.WriteLeadingString("SELECT ")
		for i := 0; i < len(j.selectedTables); i++ {
			currentTable := j.selectedTables[i]
			for j := 0; j < len(currentTable.fields); j++ {
				currentValue := currentTable.fields[j]
				buf.WriteString(fmt.Sprintf("%s.%s", currentTable.name, currentValue))
			}

			if i < len(j.selectedTables)-1 {
				buf.WriteString(",")
			}
		}
	}

	if j.fromTable != "" {
		buf.WriteLeadingString("FROM ")
		buf.WriteString(j.fromTable)
	}

	if j.typeOfJoin != NOJOIN {
		buf.WriteLeadingString(fmt.Sprintf("%s JOIN %s ", j.typeOfJoin.String(), j.joinTable))
	}

	if len(j.onSelected) > 0 {
		buf.WriteLeadingString("ON ")
		for idx, t := range j.onSelected {
			currentTable := t.name
			for i := 0; i < len(t.fields); i++ {
				currentValue := t.fields[i]
				buf.WriteString(fmt.Sprintf("%s.%s", currentTable, currentValue))
				// if i < len(v)-1 {
				// 	buf.WriteString(",")
				// }
				if idx < len(j.onSelected)-1 {
					buf.WriteLeadingString("= ")
				}
			}
		}
	}
	buf.WriteString(";")

	return buf.String()
}

func (j *JoinBuilder) SetDriver(sqlDriver Driver) *JoinBuilder {
	j.driver = sqlDriver
	return j
}

func (j JoinBuilder) GetDriver() Driver {
	return j.driver
}
