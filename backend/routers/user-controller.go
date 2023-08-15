package routers

import (
	"moflo-be/services"
)

func init() {
	UserRouter.HandleFunc("/register", services.UserRegister)
	UserRouter.HandleFunc("/login", services.UserLogin)
}
