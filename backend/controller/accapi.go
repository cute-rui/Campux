package controller

import (
	"github.com/RockChinQ/Campux/backend/service"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct {
	APIRouter
	AccountService service.AccountService
}

func NewAccountRouter(rg *gin.RouterGroup, as service.AccountService) *AccountRouter {
	ar := &AccountRouter{
		AccountService: as,
	}

	group := rg.Group("/account")

	// bind routes
	group.POST("/create", ar.CreateAccount)
	group.POST("/login", ar.LoginAccount)
	group.PUT("/reset", ar.ResetPassword)

	return ar
}

// 创建账户
func (ar *AccountRouter) CreateAccount(c *gin.Context) {
	// 取body的json里的uin
	var body AccountCreateBody

	if err := c.ShouldBindJSON(&body); err != nil {
		ar.Fail(c, 1, err.Error())
		return
	}

	// 创建账户
	pwd, err := ar.AccountService.CreateAccount(body.Uin)

	if err != nil {
		ar.Fail(c, 1, err.Error())
		return
	}

	ar.Success(c, gin.H{
		"passwd": pwd,
	})
}

// 登录
func (ar *AccountRouter) LoginAccount(c *gin.Context) {
	// 取body的json里的uin和pwd
	var body AccountLoginBody

	if err := c.ShouldBindJSON(&body); err != nil {
		ar.Fail(c, 1, err.Error())
		return
	}

	// 检查账户
	token, err := ar.AccountService.CheckAccount(body.Uin, body.Passwd)

	if err != nil {
		ar.Fail(c, 1, err.Error())
		return
	}

	ar.Success(c, gin.H{
		"token": token,
	})
}

// 重置密码
func (ar *AccountRouter) ResetPassword(c *gin.Context) {
	// 取body的json里的uin
	var body AccountCreateBody

	if err := c.ShouldBindJSON(&body); err != nil {
		ar.Fail(c, 1, err.Error())
		return
	}

	// 重置密码
	pwd, err := ar.AccountService.ResetPassword(body.Uin)

	if err != nil {
		ar.Fail(c, 1, err.Error())
		return
	}

	ar.Success(c, gin.H{
		"passwd": pwd,
	})
}
