package response

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Body struct {
	Status int32       `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type DodyOption interface {
	apply(*Body)
}

type funcOption struct {
	f func(*Body)
}

func (fdo *funcOption) apply(do *Body) {
	fdo.f(do)
}

func newFuncOption(f func(*Body)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithStatus(s int32) DodyOption {
	return newFuncOption(func(o *Body) {
		o.Status = int32(s)
	})
}

func WithMsg(s string) DodyOption {
	return newFuncOption(func(o *Body) {
		o.Msg = s
	})
}

func WithResp(r *Response) {
	WithStatus(int32(r.Code))
	WithMsg(r.Msg)
}

func WithData(s interface{}) DodyOption {
	return newFuncOption(func(o *Body) {
		o.Data = s
	})
}

type repsData struct {
	opts Body
}

//默认参数
func defaultSucOptions() Body {
	return Body{
		0,
		GetErrMsg(0),
		//map[string]interface{}{},
		[]string{},
	}
}

func defaultErrOptions() Body {
	return Body{
		1,
		GetErrMsg(1),
		//map[string]interface{}{},
		[]string{},
	}
}

func Suc(ctx *gin.Context, opts ...DodyOption) {
	cc := &repsData{
		opts: defaultSucOptions(),
	}
	//循环调用opts
	for _, opt := range opts {
		opt.apply(&cc.opts)
	}
	responseBody := NewResponseBody()
	responseBody.SetStatus(cc.opts.Status)
	defer RecoverResponse(ctx, responseBody)
	if cc.opts.Msg != "" {
		responseBody.SetErrMsg(cc.opts.Msg)
	}
	responseBody.SetData(cc.opts.Data)
	return
}

func Err(ctx *gin.Context, opts ...DodyOption) {
	cc := &repsData{
		opts: defaultErrOptions(),
	}
	//循环调用opts
	for _, opt := range opts {
		opt.apply(&cc.opts)
	}
	responseBody := NewResponseBody()
	responseBody.SetStatus(cc.opts.Status)
	defer RecoverResponse(ctx, responseBody)
	if cc.opts.Status != 0 {
		responseBody.SetErrMsg(GetErrMsg(cc.opts.Status))
	}
	if cc.opts.Msg != "" {
		responseBody.SetErrMsg(cc.opts.Msg)
	}
	responseBody.SetData(cc.opts.Data)
	return
}

func NewResponseBody() *Body {
	return &Body{
		Status: ErrnoSuccess,
		Msg:    GetErrMsg(ErrnoSuccess),
		Data:   map[string]interface{}{},
	}

}

func (res *Body) SetData(data interface{}) {
	res.Data = data
}

func (res *Body) SetStatus(statusCode int32) {
	res.Status = statusCode
}

func (res *Body) SetErrMsg(errMsg string) {
	res.Msg = errMsg
}

func Success(ctx *gin.Context, data interface{}, msg string) {
	responseBody := NewResponseBody()
	defer RecoverResponse(ctx, responseBody)
	if msg != "" {
		responseBody.SetErrMsg(msg)
		fmt.Println(responseBody)
	}
	responseBody.SetData(data)
}

func Error(ctx *gin.Context, statusCode int32, msg string, data interface{}) {

	responseBody := NewResponseBody()
	defer RecoverResponse(ctx, responseBody)
	responseBody.Status = statusCode
	responseBody.Msg = msg
	fmt.Println(responseBody)
	if responseBody.Status != 0 && msg == "" {
		responseBody.SetErrMsg(GetErrMsg(statusCode))
		fmt.Println(msg)
	}
	responseBody.SetData(data)
}

func RecoverResponse(ctx *gin.Context, responseBody *Body) {
	if err := recover(); err != nil {
		responseBody.SetStatus(ErrnoUnKnown)
	}
	resp, err := json.Marshal(responseBody)
	if err != nil {
		ctx.Data(http.StatusOK, "application/json;charset=utf-8", []byte(`{"error":1,"msg":"unknown"}`))
		return
	}
	ctx.Data(http.StatusOK, "application/json;charset=utf-8", resp)
	return
}
