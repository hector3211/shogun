# Shogun

---

A simple and light SQL query builder

### Examples

---

##### INSERT

```go
stmt := NewInsertBuilder().Insert("users").Columns("name").Values("maddog")

fmt.Println(stmt.Build())
// OUTPUT
// "INSERT INTO users (name) VALUES ('maddog');"


// InsertBuilder can also be constructed like this also
stmtTwo := Insert("users").Columns("name", "age").Values("maddog", 20)

fmt.Println(stmtTwo.Build())
// OUTPUT
// "INSERT INTO users (name,age) VALUES ('maddog',20);"


// Create an UPSERT
stmtThree := Insert("users").
    Columns("id", "name").
    Values(1, "Alice").
    OnConflict("id").
    DoUpdate("name", "NewAlice")

fmt.Println(stmtThree.Build())
// OUTPUT
// "INSERT INTO users (id,name) VALUES (1,'Alice') ON CONFLICT(id) DO UPDATE SET name = 'NewAlice';"
```

##### SELECT

```go
stmt := NewSelectBuilder().Select("id,name").From("users")

fmt.Println(stmt.Build())
// OUTPUT
// "SELECT id,name FROM users;"

stmtTwo := NewSelectBuilder().
    Select("id", "name").
    From("users").
    Where(
        Equal("name", "john"),
        And(),
        GreaterThan("id", 10),
    )

fmt.Println(stmtTwo.Build())
// OUTPUT
// "SELECT id,name FROM users WHERE name = 'john' AND id > 10;"

// SelectBuilders can also be constructed like this also
stmtThree := Select("*").From("users").OrderBy("name").Asc().Limit(15)

fmt.Println(stmtThree.String())
// OUTPUT
// "SELECT * FROM users ORDER BY name ASC LIMIT 15;"
```

##### UPDATE

```go
stmt := NewUpdateBuilder().
    Update("users").
    Set(Equal("name", "maddog")).
    Where(Equal("name", "john"))

fmt.Println(stmt.Build())
// OUTPUT
// "UPDATE users SET name = 'maddog' WHERE name = 'john';"


stmtTwo := Update("users").
    Set(Equal("name", "maddog")).
    Where(Equal("name", "john"), Or(), Equal("id", 1))

fmt.Println(stmtTwo.String())
// OUTPUT
// "UPDATE users SET name = 'maddog' WHERE name = 'john' OR id = 1;"
```

##### DELETE

```go
stmt := Delete("users").Where(Equal("id", 1), And(), Equal("name", "john")).Build()

fmt.Println(stmt.Build())
// OUTPUT
// "DELETE users WHERE id = 1 AND name = 'john';"


stmtTwo := NewDeleteBuilder().Delete("users").Where(
    Equal("id", 1),
    Or(),
    Equal("name", "john"),
    And(),
    Equal("verifyEmail", false)).Build()

fmt.Println(stmtTwo.String())
// OUTPUT
// "DELETE users WHERE id = 1 OR name = 'john' AND verifyEmail = FALSE;"
```

##### JOIN

```go
stmt := NewJoinBuilder().
    JSelect("orders", "orders_id").
    JSelect("customers", "customers_name").
    JSelect("orders", "orders_date").
    JFrom("orders").
    Join(RIGHT, "customers").
    OnCondition("orders", "customer_id", EQUAL, "customers", "customers_id", nil).
    And().
    OnCondition("orders", "customer_id", NOTEQUAL, "customers", "customers_id", nil)

fmt.Println(stmt.Build())
// OUTPUT
// "SELECT orders.orders_id,orders.orders_date,customers.customers_name FROM orders RIGHT JOIN customers ON orders.customer_id = customers.customers_id AND orders.customer_id != customers.customers_id;"
```
