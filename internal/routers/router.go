package routers

import (
    "github.com/gorilla/mux"
    "go-auth-app/internal/middleware"
     authHandlers "go-auth-app/internal/handlers/auth"
     adminHandlers "go-auth-app/internal/handlers/admin"
     userHandlers "go-auth-app/internal/handlers/users"
    "net/http"
     _ "go-auth-app/cmd/docs"

    httpSwagger "github.com/swaggo/http-swagger"

)

func SetupRoutes() *mux.Router {
    r := mux.NewRouter()

    r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)


    r.HandleFunc("/register", authHandlers.RegisterHandler).Methods("POST")
    r.HandleFunc("/login", authHandlers.LoginHandler).Methods("POST")
    r.HandleFunc("/auth/logout", authHandlers.LogoutHandler).Methods("POST")
    r.HandleFunc("/auth/refresh", authHandlers.RefreshHandler).Methods("POST")

    r.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(userHandlers.ProfileHandler)))

    adminRouter := r.PathPrefix("/admin").Subrouter()
    adminRouter.Use(middleware.RoleMiddleware("admin"))
    adminRouter.HandleFunc("", adminHandlers.AdminHandler).Methods("GET")
    adminRouter.HandleFunc("/users", adminHandlers.GetAllUsers).Methods("GET")
    adminRouter.HandleFunc("/users", adminHandlers.CreateUser).Methods("POST")
    adminRouter.HandleFunc("/users/{id}", adminHandlers.UpdateUser).Methods("PUT")
    adminRouter.HandleFunc("/users/{id}", adminHandlers.DeleteUser).Methods("DELETE")
    adminRouter.HandleFunc("/dashboard", adminHandlers.AdminDashboard).Methods("GET")

    return r
}