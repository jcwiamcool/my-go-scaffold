package authorized

import (
	"net/http"

	"github.com/jcwiamcool/my-go-scaffold/internal/code"
	"github.com/jcwiamcool/my-go-scaffold/internal/pkg/core"
)

type updateUsedRequest struct {
	Id   string `form:"id"`   // 主键ID
	Used int32  `form:"used"` // 是否启用 1:是 -1:否
}

type updateUsedResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// UpdateUsed 更新调用方为启用/禁用
// @Summary 更新调用方为启用/禁用
// @Description 更新调用方为启用/禁用
// @Tags API.authorized
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id formData string true "hashID"
// @Param used formData int true "是否启用 1:是 -1:否"
// @Success 200 {object} updateUsedResponse
// @Failure 400 {object} code.Failure
// @Router /api/authorized/used [patch]
// @Security LoginToken
func (h *handler) UpdateUsed() core.HandlerFunc {
	return func(c core.Context) {
		req := new(updateUsedRequest)
		res := new(updateUsedResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		id := int32(ids[0])

		err = h.authorizedService.UpdateUsed(c, id, req.Used)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedUpdateError,
				code.Text(code.AuthorizedUpdateError)).WithError(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
