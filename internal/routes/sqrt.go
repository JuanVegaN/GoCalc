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
*   This file contains the route localhost:3000/sqrt
*
 */

package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/JuanVegaN/GoCalc.git/pkg/mathops"
)

func Sqrt(w http.ResponseWriter, r *http.Request) {
	var req JsonReq
	var resultAdd float64
	var ops string = "Sqrt Entero/Flotante"
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error en  el JSON Request")
	}
	json.Unmarshal(reqBody, &req)
	resultAdd = mathops.Sqrt(req.FirstOperator)
	resp := JsonResp{
		Operation:      ops,
		FirstOperator:  req.FirstOperator,
		SecondOperator: req.SecondOperator,
		Result:         resultAdd,
		Message:        "La raiz cuadrada se ha realizado con exito",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
