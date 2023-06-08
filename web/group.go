package web

type Group struct {
	server *HttpServer
	prefix string
	parent *Group
}

func newGroup(server *HttpServer, prefix string) *Group {

	if prefix[0] != '/' {
		panic("[route] group prefix must start with '/'")
	}

	return &Group{
		server: server,
		prefix: prefix,
	}
}

func (g *Group) Group(prefix string) *Group {
	if prefix[0] != '/' {
		panic("[route] group prefix must start with '/'")
	}

	child := newGroup(g.server, prefix)
	child.parent = g

	return child
}

func (g *Group) Get(path string, handleFunc HandleFunc) {
	g.server.Get(g.GetAbsolutePrefix()+path, handleFunc)
}

func (g *Group) Post(path string, handleFunc HandleFunc) {
	g.server.Post(g.GetAbsolutePrefix()+path, handleFunc)
}

func (g *Group) Put(path string, handleFunc HandleFunc) {
	g.server.Put(g.GetAbsolutePrefix()+path, handleFunc)
}

func (g *Group) Delete(path string, handleFunc HandleFunc) {
	g.server.Delete(g.GetAbsolutePrefix()+path, handleFunc)
}

func (g *Group) GetAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}

	return g.parent.GetAbsolutePrefix() + g.prefix
}
