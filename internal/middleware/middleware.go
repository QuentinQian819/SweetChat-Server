package middleware

import (
	"chatBox/internal/logic/jwt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMiddleware struct{}

var Middleware = &sMiddleware{}

// Auth JWT认证中间件
func (s *sMiddleware) Auth(r *ghttp.Request) {
	// 获取token
	token := r.Header.Get("Authorization")
	if token == "" {
		token = gconv.String(r.Get("token"))
	}

	// 去掉 Bearer 前缀
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	if token == "" {
		r.Response.WriteJsonExit(g.Map{
			"code":    401,
			"message": "请先登录",
			"data":    nil,
		})
		return
	}

	// 解析token
	claims, err := jwt.ParseToken(token)
	if err != nil {
		r.Response.WriteJsonExit(g.Map{
			"code":    401,
			"message": "登录已过期，请重新登录",
			"data":    nil,
		})
		return
	}

	// 将用户信息存入上下文
	r.SetCtxVar("userId", claims.UserId)
	r.SetCtxVar("phone", claims.Phone)

	r.Middleware.Next()
}
