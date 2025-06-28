package message

type Req struct {
	ServiceName string
	MethodName  string
	Arg         []byte
}

type Resp struct {
	Data     []byte
	Err      string
	Metadata map[string]string
}
