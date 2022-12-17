package base

/**
 * @Author: Kim
 * @Description:
 * @File:  resposeBase
 * @Date: 12/17/2022 9:28 AM
 */

type ResponseBase struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
