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
		Name:   "V1",
	},
	{
		ID:     2,
		Parent: 0,
		Name:   "V2",
	},
	{
		ID:     3,
		Parent: 0,
		Name:   "V3",
	},
	{
		ID:     4,
		Parent: 1,
		Name:   "V4",
	},
	{
		ID:     5,
		Parent: 1,
		Name:   "V5",
	},
	{
		ID:     6,
		Parent: 1,
		Name:   "V6",
	},
	{
		ID:     7,
		Parent: 2,
		Name:   "V7",
	},
	{
		ID:     8,
		Parent: 2,
		Name:   "V8",
	},
	{
		ID:     9,
		Parent: 2,
		Name:   "V9",
	},
	{
		ID:     10,
		Parent: 3,
		Name:   "V10",
	},
	{
		ID:     11,
		Parent: 3,
		Name:   "V11",
	},
	{
		ID:     13,
		Parent: 3,
		Name:   "V13",
	},
}

var jsonTreeExp = []byte(`[{"ID":1,"Name":"V1","Parent":0,"Path":"","Children":[{"ID":4,"Name":"V4","Parent":1,"Path":"","Children":null},{"ID":5,"Name":"V5","Parent":1,"Path":"","Children":null},{"ID":6,"Name":"V6","Parent":1,"Path":"","Children":null}]},{"ID":2,"Name":"V2","Parent":0,"Path":"","Children":[{"ID":7,"Name":"V7","Parent":2,"Path":"","Children":null},{"ID":8,"Name":"V8","Parent":2,"Path":"","Children":null},{"ID":9,"Name":"V9","Parent":2,"Path":"","Children":null}]},{"ID":3,"Name":"V3","Parent":0,"Path":"","Children":[{"ID":10,"Name":"V10","Parent":3,"Path":"","Children":null},{"ID":11,"Name":"V11","Parent":3,"Path":"","Children":null},{"ID":13,"Name":"V13","Parent":3,"Path":"","Children":null}]}]`)

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

var xmlTreeExp = []byte("<Category><ID>1</ID><Name>V1</Name><Parent>0</Parent><Path></Path><Children><ID>4</ID><Name>V4</Name><Parent>1</Parent><Path></Path></Children><Children><ID>5</ID><Name>V5</Name><Parent>1</Parent><Path></Path></Children><Children><ID>6</ID><Name>V6</Name><Parent>1</Parent><Path></Path></Children></Category><Category><ID>2</ID><Name>V2</Name><Parent>0</Parent><Path></Path><Children><ID>7</ID><Name>V7</Name><Parent>2</Parent><Path></Path></Children><Children><ID>8</ID><Name>V8</Name><Parent>2</Parent><Path></Path></Children><Children><ID>9</ID><Name>V9</Name><Parent>2</Parent><Path></Path></Children></Category><Category><ID>3</ID><Name>V3</Name><Parent>0</Parent><Path></Path><Children><ID>10</ID><Name>V10</Name><Parent>3</Parent><Path></Path></Children><Children><ID>11</ID><Name>V11</Name><Parent>3</Parent><Path></Path></Children><Children><ID>13</ID><Name>V13</Name><Parent>3</Parent><Path></Path></Children></Category>")

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

func TestGeneratePaths(t *testing.T) {
	cm := NewCm()
	for _, c := range treeInput {
		p := c
		cm.Set(&p)
	}

	cm.GeneratePaths()

	exp := []string{"/V1", "/V2", "/V3", "/V1/V4", "/V1/V5", "/V1/V6", "/V2/V7", "/V2/V8",
		"/V2/V9", "/V3/V10", "/V3/V11", "/V3/V13"}

	for i, c := range cm.Index() {
		if cm.cats[c].Path != exp[i] {
			t.Fatalf("Failed. Expected %v got %v\n", exp[i], cm.cats[c].Path)
		}
	}
}
