package response

var (
	OK       = response(200, "ok") // 通用成功
	ERROR    = response(500, "")   // 通用错误
	ErrParam = response(10001, "参数有误")
	//ErrSignParam = response(10002, "签名参数有误")
	//
	//// 模块级错误码 - 用户模块
	//ErrUserService = response(20100, "用户服务异常")
	//ErrUserPhone   = response(20101, "用户手机号不合法")
	//ErrUserCaptcha = response(20102, "用户验证码有误")
	//
	//// 库存模块
	//ErrOrderService = response(20200, "订单服务异常")
	//ErrOrderOutTime = response(20201, "订单超时")
)
