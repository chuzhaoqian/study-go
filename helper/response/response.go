package response

import (
	"encoding/json"
	"fmt"
	"study-go/helper/status"
)

type Data map[string]interface{}

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

func NewResp() *Resp {
	resp := new(Resp)
	m := make(Data)
	resp.Data = m
	return resp
}

func NewRespMess(code int, message string) *Resp {
	resp := NewResp()
	resp.Code = code
	resp.Message = message
	return resp
}

func NewError(err error) *Resp {
	code := status.Error()
	return NewStatus(code, fmt.Sprintf("%s", err))
}

func NewStatus(code int, message string) *Resp {
	resp := NewResp()
	resp.Code = code
	resp.Message = message
	return resp
}

func NewData(code int, message string, data Data) *Resp {
	resp := NewResp()
	resp.Code = code
	resp.Message = message
	resp.Data = data
	return resp
}

func NewOKData(data Data) *Resp {
	code, message, err := status.Ok()
	if nil != err {
		return NewError(err)
	}
	return NewData(code, message, data)
}

func NewErrorResp(code int) *Resp {
	newCode, message, err := status.Status(code)
	if nil != err {
		return NewError(err)
	}
	return NewRespMess(newCode, message)
}

func NewValidError(message string) *Resp {
	var code int = status.InValid()
	return NewStatus(code, message)
}

func Json(resp *Resp) ([]byte, error) {
	j, err := json.Marshal(resp)
	if nil != err {
		return NewErrorJson(err)
	}
	return j, err
}

func String(resp *Resp) (string, error) {
	str, err := Json(resp)
	return string(str), err
}

func NewErrorJson(err error) ([]byte, error) {
	resp := NewError(err)
	return json.Marshal(resp)
}

func NewErrorRespJson(code int) ([]byte, error) {
	resp := NewErrorResp(code)
	return Json(resp)
}

func NewErrorRespJsonStr(code int) (string, error) {
	j, err := NewErrorRespJson(code)
	return string(j), err
}

func NewOKDataJson(data Data) ([]byte, error) {
	resp := NewOKData(data)
	return Json(resp)
}

func NewOKDataJsonStr(data Data) (string, error) {
	j, err := NewOKDataJson(data)
	return string(j), err
}

func NewValidErrorJson(message string) ([]byte, error) {
	resp := NewValidError(message)
	return Json(resp)
}

func NewValidErrorJsonStr(message string) (string, error) {
	j, err := NewValidErrorJson(message)
	return string(j), err
}
