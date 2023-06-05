package web

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestRouter_AddRoute(t *testing.T) {

	tcs := []struct {
		method string
		path   string
	}{
		{
			method: http.MethodGet,
			path:   "/",
		}, {
			method: http.MethodGet,
			path:   "/home",
		}, {
			method: http.MethodGet,
			path:   "/user/get",
		}, {
			method: http.MethodGet,
			path:   "/user/list",
		}, {
			method: http.MethodPost,
			path:   "/user/edit",
		},
	}

	r := newRouter()
	mockHandleFunc := func(ctx *Context) {}

	for _, tc := range tcs {
		r.addRoute(tc.method, tc.path, mockHandleFunc)
	}

	wantRouter := &router{
		methodTrees: map[string]*node{
			http.MethodGet: {
				path:       "/",
				handleFunc: mockHandleFunc,
				children: map[string]*node{
					"home": {
						path:       "home",
						handleFunc: mockHandleFunc,
					},
					"user": {
						path: "user",
						children: map[string]*node{
							"get": {
								path:       "get",
								handleFunc: mockHandleFunc,
							},
							"list": {
								path:       "list",
								handleFunc: mockHandleFunc,
							},
						},
					},
				},
			},
			http.MethodPost: {
				path: "/",
				children: map[string]*node{
					"user": {
						path: "user",
						children: map[string]*node{
							"edit": {
								path:       "edit",
								handleFunc: mockHandleFunc,
							},
						},
					},
				},
			},
		},
	}

	res, msg := wantRouter.equal(r)
	assert.True(t, res, msg)

}

func (r *router) equal(target *router) (bool, string) {
	for method, tree := range r.methodTrees {
		dst, ok := target.methodTrees[method]
		if !ok {
			return false, fmt.Sprintf("method tree '%s' unmatched", method)
		}

		res, msg := tree.equal(dst)
		if !res {
			return res, msg
		}
	}

	return true, ""
}

func (n *node) equal(target *node) (bool, string) {
	if n.path != target.path {
		return false, fmt.Sprintf("child node path '%s' unmatched", n.path)
	}

	if len(n.children) != len(target.children) {
		return false, "children length unmatched"
	}

	if reflect.ValueOf(n.handleFunc) != reflect.ValueOf(target.handleFunc) {
		return false, "node handle func unmatched"
	}

	for path, nd := range n.children {
		dst, ok := target.children[path]
		if !ok {
			return false, fmt.Sprintf("")
		}

		res, msg := nd.equal(dst)
		if !res {
			return res, msg
		}
	}

	return true, ""
}
