package router

import "net/http"

type Node struct {
	isRoot    bool
	character byte
	children  []*Node
	handlers  map[string]http.Handler
}

func newNode(character byte) *Node {
	return &Node{
		character: character,
		children:  []*Node{},
		handlers:  make(map[string]http.Handler),
	}
}

type Router struct {
	tree *Node
}

func NewRouter() *Router {
	return &Router{
		tree: &Node{
			isRoot:   true,
			handlers: nil,
		},
	}
}