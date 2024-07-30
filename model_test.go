package shogun

import "testing"

type MyModel struct {
	ID   int    `orm:"id"`
	Name string `orm:"string"`
}

func TestModelOne(t *testing.T) {}
