package api

import (
    "github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
    userRouter := router.PathPrefix("/user").Subrouter()
    RegisterUserRoutes(userRouter)

    adminRouter := router.PathPrefix("/admin").Subrouter()
    RegisterAdminRoutes(adminRouter)
}
