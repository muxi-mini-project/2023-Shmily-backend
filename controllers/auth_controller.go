package controllers

import (
	"2023-Shmily-backend/models/user"
	"2023-Shmily-backend/pkg/auth"
	"2023-Shmily-backend/pkg/flash"
	"2023-Shmily-backend/pkg/logger"
	"2023-Shmily-backend/pkg/route"
	"2023-Shmily-backend/pkg/view"
	"2023-Shmily-backend/requests"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// AuthController 处理用户认证
type AuthController struct {
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {

	// 1. 初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	// 2. 表单规则
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		// 3. 表单不通过 —— 重新显示表单
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		// 4. 验证成功，创建数据
		_user.Create()

		if _user.ID > 0 {
			// 登录用户并跳转到首页
			flash.Success("恭喜您注册成功！")
			auth.Login(_user)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "注册失败，请联系管理员")
		}
	}
}

// Login 显示登录表单
func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}

// DoLogin 处理登录表单提交
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {

	// 1. 初始化表单数据
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// 2. 尝试登录
	if err := auth.Attempt(email, password); err == nil {
		// 登录成功
		flash.Success("欢迎回来！")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// 3. 失败，显示错误提示
		view.RenderSimple(w, view.D{
			"Error":    err.Error(),
			"Email":    email,
			"Password": password,
		}, "auth.login")
	}
}

// Logout 退出登录
func (*AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	auth.Logout()
	flash.Success("您已退出登录")
	http.Redirect(w, r, "/", http.StatusFound)
}

// SendEmail 发送邮件表单
func (*AuthController) SendEmail(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.sendemail")
}

// DoSendEmail 发送邮件
func (*AuthController) DoSendEmail(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	// 1. 根据 Email 获取用户
	_user, err := user.GetByEmail(email)

	// 2. 如果出现错误
	var e error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e = errors.New("账号不存在")
		} else {
			e = errors.New("内部错误，请稍后尝试")
		}
		view.RenderSimple(w, view.D{
			"Error": e.Error(),
			"Email": email,
		}, "auth.sendemail")
	} else {
		id := _user.ID
		// TODO: 发送带id链接邮件到邮箱
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "邮件发送成功，ID: "+strconv.FormatUint(id, 10))
	}
}

// ResetPassword 重置密码表单
func (*AuthController) ResetPassword(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	view.RenderSimple(w, view.D{
		"ID": id,
	}, "auth.reset")
}

// DoResetPassword 重置密码
func (*AuthController) DoResetPassword(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	_user, err := user.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 用户未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		password := r.PostFormValue("password")
		password_confirm := r.PostFormValue("password_confirm")
		errs := validateResetPasswordFormData(password, password_confirm)
		if len(errs) > 0 {
			view.RenderSimple(w, view.D{
				"Errors":          errs,
				"Password":        password,
				"PasswordConfirm": password_confirm,
			}, "auth.reset")
		} else {
			_user.Password = password
			rowsAffected, err := _user.Update()
			if err != nil {
				// 数据库错误
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
				return
			}
			if rowsAffected > 0 {
				showURL := route.Name2URL("auth.login")
				http.Redirect(w, r, showURL, http.StatusFound)
			} else {
				fmt.Fprint(w, "您没有做任何更改！")
			}
		}
	}
}

func validateResetPasswordFormData(password string, password_confirm string) map[string][]string {
	errs := make(map[string][]string)
	if password == "" {
		errs["password"] = append(errs["password"], "密码为必填项")
	}
	if len(password) < 6 {
		errs["password"] = append(errs["password"], "密码长度需大于 6")
	}
	if password != password_confirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配！")
	}
	return errs
}
