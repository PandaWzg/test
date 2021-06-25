package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RespError interface {
	GetCode() int
}

type Response struct {
	Result Result
	ctx    *gin.Context
}

const (
	SuccessCOde = 1000
	ErrorCode   = -1
)

func NewCode(code int) RespError {
	r := &Response{}
	r.Result.Code = code
	return r
}

func New(ctx *gin.Context) *Response {
	r := &Response{
		ctx: ctx,
	}
	r.Result.Code = SuccessCOde
	return r
}

func (r *Response) GetCode() int {
	return r.Result.Code
}

func (r *Response) Error(code int) *Response {
	r.Result.Code = code
	r.Result.Data = nil
	r.Result.Message = GetError(code)
	return r
}

func (r *Response) Data(data interface{}) *Response {
	if data != nil {
		r.Result.Data = data
	} else {
		r.Result.Data = nil
	}
	return r
}

func (r *Response) SetMessage(msg string) *Response {
	r.Result.Message = msg
	return r
}

func (r *Response) Format() {
	if r.Result.Message == "" {
		r.SetMessage("success")
	}
	r.ctx.JSON(http.StatusOK, &r.Result)
}
