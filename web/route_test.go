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
			path:   "/user",
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
						path:       "user",
						handleFunc: mockHandleFunc,
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

func TestRouter_AddRoutePanic(t *testing.T) {
	r := newRouter()
	mockHandleFunc := func(ctx *Context) {}

	assert.PanicsWithValue(t, "[route] empty path", func() {
		r.addRoute(http.MethodGet, "", mockHandleFunc)
	})

	assert.PanicsWithValue(t, "[route] path not start with '/'", func() {
		r.addRoute(http.MethodGet, "home", mockHandleFunc)
	})

	registered := "/"
	r.addRoute(http.MethodGet, registered, mockHandleFunc)
	assert.PanicsWithValue(
		t,
		fmt.Sprintf(fmt.Sprintf("[route] path '%s' has already registered", registered)),
		func() {
			r.addRoute(http.MethodGet, registered, mockHandleFunc)
		},
	)

	registered = "/user/edit"
	r.addRoute(http.MethodPost, registered, mockHandleFunc)
	assert.PanicsWithValue(
		t,
		fmt.Sprintf(fmt.Sprintf("[route] path '%s' has already registered", registered)),
		func() {
			r.addRoute(http.MethodPost, registered, mockHandleFunc)
		},
	)
}

func TestRouter_findRoute(t *testing.T) {

	r := newRouter()
	mockHandleFunc := func(ctx *Context) {}

	r.addRoute(http.MethodGet, "/", mockHandleFunc)
	r.addRoute(http.MethodGet, "/user/get", mockHandleFunc)
	r.addRoute(http.MethodGet, "/user/list", mockHandleFunc)
	r.addRoute(http.MethodPost, "/user/edit", mockHandleFunc)

	tcs := []struct {
		name           string
		method         string
		path           string
		wantRes        bool
		wantHandleFunc HandleFunc
	}{
		{
			name:           "index",
			method:         http.MethodGet,
			path:           "/",
			wantRes:        true,
			wantHandleFunc: mockHandleFunc,
		}, {
			name:           "userGet",
			method:         http.MethodGet,
			path:           "/user/get",
			wantRes:        true,
			wantHandleFunc: mockHandleFunc,
		}, {
			name:           "userList",
			method:         http.MethodGet,
			path:           "/user/get",
			wantRes:        true,
			wantHandleFunc: mockHandleFunc,
		}, {
			name:           "userEdit",
			method:         http.MethodPost,
			path:           "/user/edit",
			wantRes:        true,
			wantHandleFunc: mockHandleFunc,
		}, {
			name:           "user",
			method:         http.MethodPost,
			path:           "/user",
			wantRes:        false,
			wantHandleFunc: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			hf, res := r.findRoute(tc.method, tc.path)

			assert.Equal(t, tc.wantRes, res)

			if !res {
				assert.Nil(t, hf)
				return
			}
			assert.True(t, reflect.ValueOf(hf) == reflect.ValueOf(tc.wantHandleFunc))
		})
	}
}
