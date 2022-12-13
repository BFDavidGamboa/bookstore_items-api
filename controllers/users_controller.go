package controllers

type usersControllerInterface interface {
	Create()
}

type usersController struct{}

var (
	UsersController usersControllerInterface = &usersController{}
)

func (*usersController) Create() {
}
