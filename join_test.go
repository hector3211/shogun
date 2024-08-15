package shogun

import (
	"testing"
)

func TestJoinRight(t *testing.T) {
	query := "SELECT orders.orders_id,orders.orders_date,customers.customers_name FROM orders RIGHT JOIN customers ON orders.customer_id = customers.customers_id AND orders.customer_id != customers.customers_id;"

	stmt := NewJoinBuilder().
		JSelect("orders", "orders_id").
		JSelect("customers", "customers_name").
		JSelect("orders", "orders_date").
		JFrom("orders").
		Join(RIGHT, "customers").
		OnCondition("orders", "customer_id", EQUAL, "customers", "customers_id", "").
		And().
		OnCondition("orders", "customer_id", NOTEQUAL, "customers", "customers_id", "").
		Build()

	if stmt != query {
		t.Fatalf("TestJoinRight failed! wanted %s got %s", query, stmt)
	}
}

func TestJoinLeft(t *testing.T) {
	query := "SELECT employees.id,employees.name,departments.name FROM employees LEFT JOIN departments ON employees.department_id = departments.id;"

	stmt := NewJoinBuilder().
		JSelect("employees", "id").
		JSelect("employees", "name").
		JSelect("departments", "name").
		JFrom("employees").
		Join(LEFT, "departments").
		OnCondition("employees", "department_id", EQUAL, "departments", "id", "").
		Build()

	if stmt != query {
		t.Fatalf("TestJoinLeft failed! wanted %s got %s", query, stmt)
	}
}

func TestJoinInner(t *testing.T) {
	query := "SELECT employees.id,employees.name,departments.name FROM employees INNER JOIN departments ON employees.department_id = departments.id;"

	stmt := JSelect("employees", "id").
		JSelect("employees", "name").
		JSelect("departments", "name").
		JFrom("employees").
		Join(INNER, "departments").
		OnCondition("employees", "department_id", EQUAL, "departments", "id", "").
		Build()

	if stmt != query {
		t.Fatalf("TestJoinInner failed! wanted %s got %s", query, stmt)
	}
}

func TestJoinDriver(t *testing.T) {
	query := "SELECT employees.id,employees.name,departments.name FROM employees INNER JOIN departments ON employees.department_id = departments.id AND employees.name = 'hector';"

	stmt := SQLITE.NewJoinBuilder().
		JSelect("employees", "id").
		JSelect("employees", "name").
		JSelect("departments", "name").
		JFrom("employees").
		Join(INNER, "departments").
		OnCondition("employees", "department_id", EQUAL, "departments", "id", "").
		And().
		OnCondition("employees", "name", EQUAL, "", "", "hector").
		Build()

	if stmt != query {
		t.Fatalf("TestJoinDriver failed! wanted %s got %s", query, stmt)
	}
}
