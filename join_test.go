package shogun

import "testing"

// func TestJoinRight(t *testing.T) {
// 	query := "SELECT orders.order_id,customers.customers_name,orders.order_date FROM orders RIGHT JOIN customers ON orderes.customer_id = customers.id;"
//
// 	stmt := NewJoinBuilder().JSelect("orders", "order_id").JSelect("customers", "customer_name").JSelect("orders", "order_date").JFrom("orderes").Join(RIGHT, "customers").OnTable("orders", "customer_id").OnTable("customers", "customer_id").Build()
//
// 	if stmt != query {
// 		t.Fatalf("TestJoinRight failed! wanted %s got %s", query, stmt)
// 	}
// }
//
// func TestJoinLeft(t *testing.T) {
// 	query := "SELECT employees.id,employees.name,departments.name FROM employees LEFT JOIN departments ON employees.department_id = departments.id;"
//
// 	stmt := NewJoinBuilder().JSelect("employees", "id").JSelect("employees", "name").JSelect("departments", "name").JFrom("employees").Join(LEFT, "departments").OnTable("employees", "department_id").OnTable("departments", "id").Build()
//
// 	if stmt != query {
// 		t.Fatalf("TestJoinLeft failed! wanted %s got %s", query, stmt)
// 	}
// }
//
// func TestJoinInner(t *testing.T) {
// 	query := "SELECT employees.id,employees.name,departments.name FROM employees INNER JOIN departments ON employees.department_id = departments.id;"
//
// 	stmt := JSelect("employees", "id").JSelect("employees", "name").JSelect("departments", "name").JFrom("employees").Join(INNER, "departments").OnTable("employees", "department_id").OnTable("departments", "id").Build()
//
// 	if stmt != query {
// 		t.Fatalf("TestJoinInner failed! wanted %s got %s", query, stmt)
// 	}
// }

func TestJoinDriver(t *testing.T) {
	query := "SELECT employees.id,employees.name,departments.name FROM employees INNER JOIN departments ON employees.department_id = departments.id;"

	stmt := SQLITE.NewJoinBuilder().JSelect("employees", "id").JSelect("employees", "name").JSelect("departments", "name").JFrom("employees").Join(INNER, "departments").OnTable("employees", "department_id").OnTable("departments", "id").Build()

	if stmt != query {
		t.Fatalf("TestJoinDriver failed! wanted %s got %s", query, stmt)
	}
}
