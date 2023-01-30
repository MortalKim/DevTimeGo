package duration

/**
 * @Author: kim
 * @Description:
 * @File:  durationSearchParams
 * @Date: 1/30/23 9:01 PM
 */

type SearchParams struct {
	//User's id
	UserID string `form:"user_id"`
	Date   int64  `form:"date"`
}
