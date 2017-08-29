package contract

type HelloWorldRequest struct {
	Name string
}

type HelloWorldResponse struct {
	Message string
}

type UserRequest struct {
	Id int32
}

type UserResponse struct {
	Name string
}
