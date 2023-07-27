package boot

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/kysion/oss-library/example/router"
	"github.com/kysion/oss-library/oss_global"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 导入权限树
			//sys_service.SysPermission().ImportPermissionTree(ctx, oss_consts.PermissionTree, nil)

			// 初始化路由
			apiPrefix := g.Cfg().MustGet(ctx, "service.apiPrefix").String()
			s.Group(apiPrefix, func(group *ghttp.RouterGroup) {
				// 中间件注册
				group.Middleware(ghttp.MiddlewareHandlerResponse)

				// 匿名路由绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					// group.Group("oss", func(group *ghttp.RouterGroup) { group.Bind(oss_controller.Oss) })

					// 路由绑定
					router.ModulesGroup(oss_global.Global.Modules, group)
				})

			})

			s.Run()
			return nil
		},
	}
)
