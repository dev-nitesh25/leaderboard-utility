/*
Package main - Routes File from where end point call to their respective functions
*/

package main

import (
	"net/http"

	"api/leaderboard"

	"github.com/julienschmidt/httprouter"
)

func addRouteHandlers(router *httprouter.Router) {

	//Index Routes to check API
	router.GET("/", index)

	//User Leaderboard related API
	router.GET("/import-leaderboard", leaderboard.AddCompLeaderboard)
	router.GET("/leaderboard/:startpos/:limit", leaderboard.GetCompLeaderboard)

}

// Idex URL to check API Running or Not
func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<html><body style='font-family:Arial'><h1>Leaderboard Check API</h1></body></html>"))
}
