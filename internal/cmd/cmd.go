package cmd

import (
	"context"

	"chatBox/internal/controller/chat"
	"chatBox/internal/controller/diary"
	"chatBox/internal/controller/promise"
	"chatBox/internal/controller/user"
	"chatBox/internal/middleware"
	"chatBox/internal/logic/chat/ws"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// WebSocket handler
			s.BindHandler("/ws/chat", ws.WSHandler)

			// Static files for uploaded resources
			s.AddStaticPath("/resource", "resource")

			// Create controllers
			userCtrl := user.NewV1()
			chatCtrl := chat.NewV1()
			diaryCtrl := diary.NewV1()
			promiseCtrl := promise.NewV1()

			// Public routes (no authentication)
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)

				// Bind public routes
				group.Bind(
					userCtrl.Register,
					userCtrl.Login,
				)
			})

			// Protected routes (require authentication)
			s.Group("/api/v1", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(middleware.Middleware.Auth)

				// Bind all protected routes
				group.Bind(
					userCtrl.GenerateInvite,
					userCtrl.BindCouple,
					userCtrl.GetProfile,
					userCtrl.UpdateProfile,
					userCtrl.GetCoupleInfo,
					chatCtrl.GetHistory,
					chatCtrl.MarkRead,
					chatCtrl.Upload,
					chatCtrl.Clear,
					diaryCtrl.Create,
					diaryCtrl.List,
					diaryCtrl.Get,
					diaryCtrl.Update,
					diaryCtrl.Delete,
					diaryCtrl.Upload,
					promiseCtrl.Create,
					promiseCtrl.List,
					promiseCtrl.Get,
					promiseCtrl.Update,
					promiseCtrl.Delete,
					promiseCtrl.ToggleComplete,
				)
			})

			s.Run()
			return nil
		},
	}
)
