package shogun

type UpdateBuilder struct {
	Action    string
	Table     string
	SetCond   [][]interface{}
	WhereCond [][]interface{}
}

func NewUpdateBuilder() *UpdateBuilder {
	return &UpdateBuilder{
		Action: "UPDATE",
	}
}

func (u *UpdateBuilder) Update(tableName string) *UpdateBuilder {
	u.Table = tableName
	return u
}

func (u *UpdateBuilder) Set(values ...interface{}) *UpdateBuilder {
	u.SetCond = append(u.SetCond, values)
	return u
}

func (u *UpdateBuilder) Where(conditions ...interface{}) *UpdateBuilder {
	u.WhereCond = append(u.WhereCond, conditions)
	return u
}
