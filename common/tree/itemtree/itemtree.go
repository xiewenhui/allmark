// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package itemtree

import (
	"fmt"
	"github.com/andreaskoch/allmark2/common/route"
	"github.com/andreaskoch/allmark2/common/tree"
	"github.com/andreaskoch/allmark2/common/tree/treeutil"
	"github.com/andreaskoch/allmark2/model"
)

func New() *ItemTree {
	return &ItemTree{
		*tree.New("", nil),
	}
}

type ItemTree struct {
	tree.Tree
}

func (nodeTree *ItemTree) Root() *model.Item {
	rootNode := nodeTree.Tree.Root()
	if rootNode == nil {
		return nil
	}

	return nodeToItem(rootNode)
}

func (nodeTree *ItemTree) Insert(item *model.Item) {

	if item == nil {
		return
	}

	// convert the route to a path
	path := treeutil.RouteToPath(item.Route())

	nodeTree.Tree.Insert(path, item)
}

func (itemTree *ItemTree) Delete(item *model.Item) (bool, error) {

	if item == nil {
		return false, fmt.Errorf("The supplied item is empty.")
	}

	itemRoute := item.Route()
	return itemTree.Tree.Delete(itemRoute.Components())
}

func (nodeTree *ItemTree) GetItem(route *route.Route) *model.Item {

	// locate the node
	node := nodeTree.getNode(route)
	if node == nil {
		return nil
	}

	// cast the value
	return nodeToItem(node)
}

func (nodeTree *ItemTree) GetChildItems(route *route.Route) []*model.Item {

	childItems := make([]*model.Item, 0)

	node := nodeTree.getNode(route)
	if node == nil {
		return childItems
	}

	for _, childNode := range node.Childs() {
		item := nodeToItem(childNode)
		if item == nil {
			// todo: log error
			continue
		}

		childItems = append(childItems, item)
	}

	return childItems
}

func (nodeTree *ItemTree) getNode(route *route.Route) *tree.Node {

	if route == nil {
		return nil
	}

	// convert the route to a path
	path := treeutil.RouteToPath(route)

	// locate the node
	node := nodeTree.Tree.GetNode(path)
	if node == nil {
		return nil
	}

	return node
}

func nodeToItem(node *tree.Node) *model.Item {

	if node.Value() == nil {
		return nil
	}

	return node.Value().(*model.Item)
}
