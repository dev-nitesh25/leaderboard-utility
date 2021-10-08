/*
Package leaderbaord - Leaderboard data type file where leaderbaord related datatypes applies
*/

package leaderboard

import "common"

// CompUser :
type CompUser struct {
	UserID    int     `json:"user_id"`
	UserName  string  `json:"user_name"`
	UserScore float64 `json:"user_score"`
}

// CompLeaderboard :
type CompLeaderboard struct {
	Leaderboard []common.UserRank `json:"leaderboard"`
}
