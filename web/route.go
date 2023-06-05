package web

import (
	"fmt"
	"strings"
)

type router struct {
	methodTrees map[string]*node
}

func newRouter() *router {
	return &router{
		methodTrees: map[string]*node{},
	}
}

func (r *router) addRoute(method string, path string, handleFunc HandleFunc) {

	if path == "" {
		panic("[route] empty path")
	}

	if path[0] != '/' {
		panic("[route] path not start with '/'")
	}

	root, ok := r.methodTrees[method]
	if !ok {
		root = &node{path: "/"}
		r.methodTrees[method] = root
	}

	if path == "/" {
		if root.handleFunc != nil {
			panic(fmt.Sprintf("[route] path '%s' has already registered", path))
		}
		root.handleFunc = handleFunc
		return
	}

	segs := strings.Split(strings.Trim(path, "/"), "/")
	for _, seg := range segs {
		root = root.createChild(seg)
	}

	if root.handleFunc != nil {
		panic(fmt.Sprintf("[route] path '%s' has already registered", path))
	}
	root.handleFunc = handleFunc

}

func (r *router) findRoute(method string, path string) (HandleFunc, bool) {

	root, ok := r.methodTrees[method]
	if !ok {
		return nil, false
	}

	if path == "/" {
		if root.handleFunc == nil {
			return nil, false
		}
		return root.handleFunc, true
	}

	segs := strings.Split(strings.Trim(path, "/"), "/")
	for _, seg := range segs {
		root, ok = root.findChild(seg)
		if !ok || root.handleFunc == nil {
			return nil, false
		}
	}

	return root.handleFunc, true
}

type node struct {
	path       string
	children   map[string]*node
	handleFunc HandleFunc
}

func (n *node) createChild(path string) *node {

	if n.children == nil {
		n.children = map[string]*node{}
	}

	if child, ok := n.children[path]; ok {
		return child
	}

	child := &node{
		path: path,
	}
	n.children[path] = child
	return child
}

func (n *node) findChild(path string) (*node, bool) {

	if n.children == nil {
		return nil, false
	}

	if child, ok := n.children[path]; ok {
		return child, true
	}

	return nil, false
}
