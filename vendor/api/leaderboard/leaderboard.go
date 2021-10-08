/*
Package leaderbaord - Leaderboard go file where business logic writter
*/
package leaderboard

import (
	"common"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//AddCompLeaderboard
func AddCompLeaderboard(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//------Get All Comp Users
	objUsers, err := GetCompUsers()
	if err != nil {
		common.OutputResponse(w, r, http.StatusInternalServerError, "Error in get UserComp"+err.Error())
		return
	}
	//------Get All Comp Users

	//------Create Map and insert in Redis
	usersMap := make(map[string]float64)
	for _, userInfo := range objUsers {
		userIDStr := strconv.Itoa(userInfo.UserID)
		userComboStr := userInfo.UserName + "|" + userIDStr
		usersMap[userComboStr] = userInfo.UserScore
	}

	compRedisKey := "comp:cricketcomp:leaderboard"
	err = common.CacheZADDWithFloat64(compRedisKey, usersMap)
	if err != nil {
		common.OutputResponse(w, r, http.StatusInternalServerError, "Error in set Leaderboard"+err.Error())
		return
	}
	//------Create Map and insert in Redis

	common.OutputResponse(w, r, http.StatusOK, "Leaderboard set successfully")

}

//GetCompLeaderboard
func GetCompLeaderboard(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	startPos, _ := strconv.Atoi(p.ByName("startpos"))
	limit, _ := strconv.Atoi(p.ByName("limit"))

	//------Get comp users from Redis
	compRedisKey := "comp:cricketcomp:leaderboard"
	usersLB, err := common.CacheZRANGEBYSCORE(compRedisKey, -99999, 99999, startPos, limit, true)
	if err != nil {
		common.OutputResponse(w, r, http.StatusInternalServerError, "Error in get Leaderboard"+err.Error())
		return
	}
	//------Get comp users from Redis

	finalJson := common.JSONMSGWrappedObj(http.StatusOK, usersLB)
	common.OutputResponseJSONObject(w, http.StatusOK, finalJson)
}
