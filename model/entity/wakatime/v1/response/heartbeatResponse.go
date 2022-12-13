package response

/**
 * @Author: Kim
 * @Description:
 * @File:  heartbeat
 * @Date: 12/12/2022 9:27 PM
 */

type HeartbeatBulkResponse struct {
	Responses [][]interface{} `json:"responses"`
}

//such as
//{
//"responses": [
//		[
//			null,
//			201
//		],
//		[
//			null,
//			201
//		],
//	]
//}
