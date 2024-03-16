package admin

import (
	"net/http"

	"github.com/jcwiamcool/my-go-scaffold/configs"
	"github.com/jcwiamcool/my-go-scaffold/internal/code"
	"github.com/jcwiamcool/my-go-scaffold/internal/pkg/core"
	"github.com/jcwiamcool/my-go-scaffold/internal/repository/redis"
	"github.com/jcwiamcool/my-go-scaffold/pkg/errors"
)

type logoutResponse struct {
	Username string `json:"username"` // 用户账号
}

// Logout 管理员登出
// @Summary 管理员登出
// @Description 管理员登出
// @Tags API.admin
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Success 200 {object} logoutResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/logout [post]
// @Security LoginToken
func (h *handler) Logout() core.HandlerFunc {
	return func(c core.Context) {
		res := new(logoutResponse)
		res.Username = c.SessionUserInfo().UserName

		if !h.cache.Del(configs.RedisKeyPrefixLoginUser+c.GetHeader(configs.HeaderLoginToken), redis.WithTrace(c.Trace())) {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AdminLogOutError,
				code.Text(code.AdminLogOutError)).WithError(errors.New("cache del err")),
			)
			return
		}

		c.Payload(res)
	}
}
