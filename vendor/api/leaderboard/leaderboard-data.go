/*
Package leaderbaord - Leaderboard data file where database related stuff perform
*/

package leaderboard

import "common"

// GetCompUsers  :
func GetCompUsers() ([]CompUser, error) {
	var compInfo []CompUser
	sqlStr := "SELECT user_id, user_name, user_score FROM tbl_comp_users"
	rows, err := common.DBCon.Query(sqlStr)
	if err != nil {
		return compInfo, err
	}
	defer rows.Close()
	for rows.Next() {
		var objCompInfo CompUser
		err := rows.Scan(
			&objCompInfo.UserID,
			&objCompInfo.UserName,
			&objCompInfo.UserScore,
		)
		if err != nil {
			return compInfo, err
		}
		compInfo = append(compInfo, objCompInfo)
	}

	return compInfo, nil
}
