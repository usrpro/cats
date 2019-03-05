package cats

import (
	"reflect"
	"testing"
)

var setInput = []*Category{
	{
		ID:   2,
		Name: "Foo",
	},
	{
		ID:   3,
		Name: "Spanac",
	},
	{
		ID:   2,
		Name: "Bar",
	},
}

var getExp = &Category{
	ID:   2,
	Name: "Bar",
}

func TestSetGet(t *testing.T) {
	cm := NewCm()
	for _, c := range setInput {
		cm.Set(c)
	}
	got := cm.Get(2)
	if got.Name != getExp.Name {
		t.Fatal("Get result.\nExpected:\n", getExp, "\nGot:\n", got)
	}
}

var sortInput = []*Category{
	{ID: 5},
	{ID: 3},
	{ID: 2},
	{ID: 4},
	{ID: 0},
	{ID: 1},
}

func TestSortIndex(t *testing.T) {
	cm := NewCm()
	for _, c := range sortInput {
		cm.Set(c)
	}
	cm.Sort()
	if len(cm.Index()) != len(sortInput) {
		t.Fatal("Length of index expected:", len(sortInput), "Got:", len(cm.Index()))
	}
	for k, v := range cm.Index() {
		if k != v {
			t.Fatal("Incorrect sort order:", cm.Index())
		}
	}
}

var treeInput = []Category{
	{
		ID:     1,
		Parent: 0,
	},
	{
		ID:     2,
		Parent: 0,
	},
	{
		ID:     3,
		Parent: 0,
	},
	{
		ID:     4,
		Parent: 1,
	},
	{
		ID:     5,
		Parent: 1,
	},
	{
		ID:     6,
		Parent: 1,
	},
	{
		ID:     7,
		Parent: 2,
	},
	{
		ID:     8,
		Parent: 2,
	},
	{
		ID:     9,
		Parent: 2,
	},
	{
		ID:     10,
		Parent: 3,
	},
	{
		ID:     11,
		Parent: 3,
	},
	{
		ID:     13,
		Parent: 3,
	},
}

var jsonTreeExp = []byte(`[{"ID":1,"Name":"","Parent":0,"Children":[{"ID":4,"Name":"","Parent":1,"Children":null},{"ID":5,"Name":"","Parent":1,"Children":null},{"ID":6,"Name":"","Parent":1,"Children":null}]},{"ID":2,"Name":"","Parent":0,"Children":[{"ID":7,"Name":"","Parent":2,"Children":null},{"ID":8,"Name":"","Parent":2,"Children":null},{"ID":9,"Name":"","Parent":2,"Children":null}]},{"ID":3,"Name":"","Parent":0,"Children":[{"ID":10,"Name":"","Parent":3,"Children":null},{"ID":11,"Name":"","Parent":3,"Children":null},{"ID":13,"Name":"","Parent":3,"Children":null}]}]`)

func TestJSONTree(t *testing.T) {
	cm := NewCm()
	for _, c := range treeInput {
		p := c
		cm.Set(&p)
	}
	got, err := cm.JSONTree(0)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, jsonTreeExp) {
		t.Fatal(
			"JSONTree\nExpected:\n",
			string(jsonTreeExp),
			"\nGot:\n",
			string(got),
		)
	}
}

var xmlTreeExp = []byte("<Category><ID>1</ID><Name></Name><Parent>0</Parent><Children><ID>4</ID><Name></Name><Parent>1</Parent></Children><Children><ID>5</ID><Name></Name><Parent>1</Parent></Children><Children><ID>6</ID><Name></Name><Parent>1</Parent></Children></Category><Category><ID>2</ID><Name></Name><Parent>0</Parent><Children><ID>7</ID><Name></Name><Parent>2</Parent></Children><Children><ID>8</ID><Name></Name><Parent>2</Parent></Children><Children><ID>9</ID><Name></Name><Parent>2</Parent></Children></Category><Category><ID>3</ID><Name></Name><Parent>0</Parent><Children><ID>10</ID><Name></Name><Parent>3</Parent></Children><Children><ID>11</ID><Name></Name><Parent>3</Parent></Children><Children><ID>13</ID><Name></Name><Parent>3</Parent></Children></Category>")

func TestXMLTree(t *testing.T) {
	cm := NewCm()
	for _, c := range treeInput {
		p := c
		cm.Set(&p)
	}
	got, err := cm.XMLTree(0)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, xmlTreeExp) {
		t.Fatal(
			"XMLTree\nExpected:\n",
			string(xmlTreeExp),
			"\nGot:\n",
			string(got),
		)
	}
}
