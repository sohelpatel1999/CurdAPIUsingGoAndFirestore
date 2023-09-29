package main

import (
	"fmt"
	"gowithcurd/firebase"
	"gowithcurd/routes"
)

func main() {
	fmt.Print("sohel first curd")
	firebase.FirestoreClient = firebase.InitializeFirestore()
	router := routes.SetupRoutes()
	fmt.Println("server Started!")
	router.Run(":8081")
}
