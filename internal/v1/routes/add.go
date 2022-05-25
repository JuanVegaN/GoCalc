/*
* File: add.go
*
* Author1:  Juan Vega, (jvega@vozy.co)
* Date:     May 2022
* Partner:  I worked alone
* Course:   Vozy assesment
*
* Summary of File:
*
*   This file contains the route localhost:3000/add
*
 */

package routes

import (
	hdl "github.com/JuanVegaN/GoCalc.git/internal/v1/handlers"
	"github.com/gorilla/mux"
)

func addRoutes(router *mux.Router) {

	// router.HandleFunc("/", routes.Index) => helat check
	router.HandleFunc("/add", hdl.AddHandler).Methods("POST")

}
