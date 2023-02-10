package routers

import (
	v1 "fiber-nuzn-api/controllers/v1"

	"github.com/gofiber/fiber/v2"
)

func SetRoute(app *fiber.App) {

	main := v1.NewDefaultController()
	group := app.Group("/v1")
	group.Get("/getList", main.GetList)    // 列表
	group.Post("/category", main.Category) // 详情

	Appv1 := app.Group("/v1/basic")
	// 登录相关的
	loginRouter := Appv1.Group("/login")
	{
		login := v1.NewLoginController()
		loginRouter.Post("/", login.Login)               // 登录
		loginRouter.Post("/userinfo", login.GetUserInfo) //获取用户信息
	}
	//v1.Use(middleware.Disable)
	// rbac 权限相关
	{
		// 用户
		userRouter := Appv1.Group("/user")
		{
			user := v1.NewUserController()
			userRouter.Post("/create", user.Create) // 创建账号
			userRouter.Post("/del", user.Del)       // 删除账号
			userRouter.Post("/update", user.Update) // 更新账号
			userRouter.Post("/list", user.GetList)  // 获取账号列

		}
		// 菜单（菜单列表）
		menuRouter := Appv1.Group("/menu")
		{
			menu := v1.NewMenuController()
			menuRouter.Post("/create", menu.Create) // 创建菜单
			menuRouter.Post("/del", menu.Del)       // 删除菜单
			menuRouter.Post("/update", menu.Update) // 编辑菜单
			menuRouter.Post("/list", menu.GetList)  // 获取菜单列表

		}
		// 权限（路径信息）
		roleRouter := Appv1.Group("/role")
		{
			role := v1.NewRoleController()
			roleRouter.Post("/create", role.Create)                                                   // 创建角色
			roleRouter.Post("/del", role.Del)                                                         // 删除角色
			roleRouter.Post("/update", role.Update)                                                   // 更新角色
			roleRouter.Post("/list", role.GetList)                                                    // 获取角色列表
			roleRouter.Post("/setCurrentRoleAuthorization", role.SetCurrentRoleAuthorization)         // 授权当前角色
			roleRouter.Post("/getCurrentRoleAuthorizationMenu", role.GetCurrentRoleAuthorizationMenu) // 获取当前角色授权菜单
		}
	}
	// 文件相关
	Files := Appv1.Group("/files")
	{
		uploadRouter := Files.Group("/upload")
		{
			upload := v1.NewUploadController()
			uploadRouter.Post("/formFile", upload.SaveUpload) // 上传任意文件
			uploadRouter.Post("/chunkFile", upload.ChunkFile) // 上传切片
			uploadRouter.Post("/mergeFile", upload.MergeFile) // 合并切片
		}
	}
	// tree 树相关的
	treeRouter := Appv1.Group("/tree")
	{
		tree := v1.NewTreeController()
		treeRouter.Post("/", tree.GetList) // tree
	}

}
