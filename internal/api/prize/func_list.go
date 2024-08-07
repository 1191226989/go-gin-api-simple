package prize

import (
	"net/http"

	"go-gin-api-simple/internal/code"
	"go-gin-api-simple/internal/pkg/core"
	"go-gin-api-simple/internal/services/prize"
	"go-gin-api-simple/pkg/timeutil"

	"github.com/spf13/cast"
)

type listRequest struct {
	Page     int     `form:"page"`      // 第几页
	PageSize int     `form:"page_size"` // 每页显示条数
	Name     string  `form:"name"`      // 奖品名称
	Worth    float64 `form:"worth"`     // 奖品价值
	Content  string  `form:"content"`   // 奖品描述
}

type listData struct {
	Id        int     `json:"id"`         // ID
	HashID    string  `json:"hashid"`     // hashid
	Name      string  `json:"name"`       // 奖品名称
	Image     string  `json:"image"`      // 奖品图片
	Worth     float64 `json:"worth"`      // 奖品价值
	Content   string  `json:"content"`    // 奖品描述
	IsUsed    int     `json:"is_used"`    // 是否启用 1:是 0:否
	CreatedAt string  `json:"created_at"` // 创建时间
	UpdatedAt string  `json:"updated_at"` // 更新时间
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PerPageCount int `json:"per_page_count"`
	} `json:"pagination"`
}

// List 奖品列表
// @Summary 奖品列表
// @Description 奖品列表
// @Tags API.prize
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param page query int true "第几页" default(1)
// @Param page_size query int true "每页显示条数" default(10)
// @Param name query string false "奖品名称"
// @Param worth query number false "奖品价值"
// @Param content query string false "奖品描述"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/prize [get]
// @Security LoginToken
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listRequest)
		res := new(listResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		page := req.Page
		if page == 0 {
			page = 1
		}

		pageSize := req.PageSize
		if pageSize == 0 {
			pageSize = 10
		}

		searchData := new(prize.SearchData)
		searchData.Page = page
		searchData.PageSize = pageSize
		searchData.Name = req.Name
		searchData.Worth = req.Worth
		searchData.Content = req.Content

		resListData, err := h.prizeService.PageList(c, searchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.PrizeListError,
				code.Text(code.PrizeListError)).WithError(err),
			)
			return
		}

		resCountData, err := h.prizeService.PageListCount(c, searchData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.PrizeListError,
				code.Text(code.PrizeListError)).WithError(err),
			)
			return
		}
		res.Pagination.Total = cast.ToInt(resCountData)
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page
		res.List = make([]listData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			data := listData{
				Id:        cast.ToInt(v.Id),
				HashID:    hashId,
				Name:      v.Name,
				Image:     v.Image,
				Worth:     v.Worth,
				Content:   v.Content,
				IsUsed:    cast.ToInt(v.IsUsed),
				CreatedAt: v.CreatedAt.Format(timeutil.CSTLayout),
				UpdatedAt: v.UpdatedAt.Format(timeutil.CSTLayout),
			}

			res.List[k] = data
		}

		c.Payload(res)
	}
}
