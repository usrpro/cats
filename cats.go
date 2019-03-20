package cats

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"sort"
)

// Category model.
type Category struct {
	ID       int
	Name     string
	Parent   int
	Path     string
	Children []*Category
}

func (c *Category) generatePath(cm *CategoryMap) {
	path := new(bytes.Buffer)

	if c.Parent != 0 {
		path.WriteString(cm.Get(c.Parent).Path)
		path.WriteString("/")
	} else {
		path.WriteString("/")
	}

	path.WriteString(c.Name)
	c.Path = path.String()

	for _, ch := range c.Children {
		ch.generatePath(cm)
	}
}

// CategoryMap a map of Category pointers and associates an index with it, for ordered output.
// The order of the index is based on the order items where added.
type CategoryMap struct {
	cats  map[int]*Category
	index []int
}

// NewCm initializes and returs a pointer to a new CategoryMap.
func NewCm() *CategoryMap {
	cats := make(map[int]*Category)
	return &CategoryMap{
		cats: cats,
	}
}

// Set a Category point to the map.
// It will also be appended to the index, keeping the order of calling this function.
func (cm *CategoryMap) Set(c *Category) {
	cm.cats[c.ID] = c
	cm.index = append(cm.index, c.ID)
}

// Get a Category pointer by its id.
func (cm *CategoryMap) Get(id int) (c *Category) {
	return cm.cats[id]
}

// Sort re-indexes the CategoryMap, increasing order on category id.
// It is advised to add the categories in a sorted way instead of using this method.
func (cm *CategoryMap) Sort() {
	sort.Ints(cm.index)
}

// Index returns a struct of category id. It allows for a sorted range loop.
func (cm *CategoryMap) Index() []int {
	return cm.index
}

// Tree creates decendant tree by population the Category's children.
// It returns a slice of Category pointers, representing the root of the tree.
func (cm *CategoryMap) Tree(offset int) (root []*Category) {
	for _, i := range cm.Index() {
		c := cm.Get(i)
		// Are we at the root of the tree?
		if c.Parent == offset {
			root = append(root, c)
			continue
		}
		p := cm.Get(c.Parent)
		// Append Category to its parent's children
		p.Children = append(p.Children, c)
	}
	return
}

// JSONTree returns a JSON document, representing the parent /
// child relationship of the categories in a tree.
// It returns an error if json.Marshall does.
func (cm *CategoryMap) JSONTree(offset int) ([]byte, error) {
	//return json.MarshalIndent(cm.Tree(offset), "", "  ")
	return json.Marshal(cm.Tree(offset))
}

// XMLTree returns a XML document, representing the parent /
// child relationship of the categories in a tree.
// It returns an error if XML.Marshall does.
func (cm *CategoryMap) XMLTree(offset int) ([]byte, error) {
	return xml.Marshal(cm.Tree(offset))
}

// GeneratePaths generate paths for each category
func (cm *CategoryMap) GeneratePaths() {
	root := cm.Tree(0)
	for _, c := range root {
		c.generatePath(cm)
	}
}
