package routes

import (
	controllers "2023-Shmily-backend/controllers"
	middlewares "2023-Shmily-backend/midwares"
	"github.com/gorilla/mux"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {

	// 静态页面
	pc := new(controllers.PagesController)
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// 文章相关页面
	ac := new(controllers.ArticlesController)
	r.HandleFunc("/", ac.Index).Methods("GET").Name("home")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	r.HandleFunc("/articles/create", middlewares.Auth(ac.Create)).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles", middlewares.Auth(ac.Store)).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", middlewares.Auth(ac.Edit)).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(ac.Update)).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", middlewares.Auth(ac.Delete)).Methods("POST").Name("articles.delete")

	// 文章分类
	cc := new(controllers.CategoriesController)
	r.HandleFunc("/categories/create", middlewares.Auth(cc.Create)).Methods("GET").Name("categories.create")
	r.HandleFunc("/categories", middlewares.Auth(cc.Store)).Methods("POST").Name("categories.store")
	r.HandleFunc("/categories", middlewares.Auth(cc.Store)).Methods("POST").Name("categories.store")
	r.HandleFunc("/categories/{id:[0-9]+}", cc.Show).Methods("GET").Name("categories.show")

	// 用户认证
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", middlewares.Guest(auc.Register)).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", middlewares.Guest(auc.DoRegister)).Methods("POST").Name("auth.doregister")
	r.HandleFunc("/auth/login", middlewares.Guest(auc.Login)).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", middlewares.Guest(auc.DoLogin)).Methods("POST").Name("auth.dologin")
	r.HandleFunc("/auth/logout", middlewares.Auth(auc.Logout)).Methods("POST").Name("auth.logout")

	r.HandleFunc("/auth/sendemail", auc.SendEmail).Methods("GET").Name("auth.sendemail")
	r.HandleFunc("/auth/dosendemail", auc.DoSendEmail).Methods("POST").Name("auth.dosendemail")
	r.HandleFunc("/auth/reset/{id:[0-9]+}", auc.ResetPassword).Methods("GET").Name("auth.reset")
	r.HandleFunc("/auth/doreset", auc.DoResetPassword).Methods("POST").Name("auth.doreset")

	// 中间件：强制内容类型为 HTML
	r.Use(middlewares.ForceHTML)

	// 开始会话
	r.Use(middlewares.StartSession)
}
