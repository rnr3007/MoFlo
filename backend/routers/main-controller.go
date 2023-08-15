package routers

import "github.com/gorilla/mux"

var Router *mux.Router = mux.NewRouter()

var UserRouter *mux.Router = Router.PathPrefix("/api/v1/user").Subrouter()
