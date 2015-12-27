package main

import (
    // Standard library packages
    "net/http"

    "gopkg.in/mgo.v2"
      

    // Third party packages
    "github.com/julienschmidt/httprouter"
    "Test/controllers"
)

func getSession() *mgo.Session {
    // Connect to our local mongo
    s, err := mgo.Dial("mongodb://localhost")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    return s
}


func main() {
    // Instantiate a new router
    r := httprouter.New()

    // Get a UserController instance
    uc := controllers.NewUserController(getSession())

    // Get a user resource
    r.GET("/user/:id", uc.GetUser)

    r.POST("/user", uc.CreateUser)

    r.DELETE("/user/:id", uc.RemoveUser)

    // Fire up the server
    http.ListenAndServe("localhost:3000", r)
}
