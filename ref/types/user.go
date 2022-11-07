package types

type User struct {
	Name string
	Age  int
	Comp Company
}

type Address struct {
	Addr string
}

type Company struct {
	Boss string
	Addr Address
}
