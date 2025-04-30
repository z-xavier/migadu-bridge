// Package tokens implements the tokens API endpoints
package tokens

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/biz"
	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/common"
	"migadu-bridge/internal/pkg/errmsg"
	"migadu-bridge/internal/pkg/log"
	v1 "migadu-bridge/pkg/api/manage/v1"
)

// Controller 定义了 controller 层需要实现的方法.
type Controller struct {
	b biz.IBiz
}

// New 创建一个 token controller.
func New(ds store.IStore) *Controller {
	return &Controller{b: biz.NewBiz(ds)}
}

// Create godoc
// @Summary Create a new token
// @Description Create a new token for API access
// @Tags tokens
// @Accept json
// @Produce json
// @Param request body v1.CreateTokenReq true "Token creation request"
// @Success 200 {object} v1.Token "Success"
// @Failure 400 {object} v1.Response "Bad request"
// @Failure 500 {object} v1.Response "Internal server error"
// @Router /api/v1/tokens [post]
func (tc *Controller) Create(c *gin.Context) (any, error) {
	log.C(c).Info("create token begin")

	var r v1.CreateTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).WithError(err).Error("create token request parse")
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().Create(c, &r)
}

// Delete godoc
// @Summary Delete a token
// @Description Delete a token by ID
// @Tags tokens
// @Accept json
// @Produce json
// @Param tokenId path string true "Token ID"
// @Success 200 {object} nil "Success"
// @Failure 400 {object} v1.Response "Bad request"
// @Failure 404 {object} v1.Response "Token not found"
// @Failure 500 {object} v1.Response "Internal server error"
// @Router /api/v1/tokens/{tokenId} [delete]
func (tc *Controller) Delete(c *gin.Context) (any, error) {
	log.C(c).Info("delete token begin")

	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Error("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	return nil, tc.b.Token().Delete(c, tokenId)
}

// Put godoc
// @Summary Update a token
// @Description Update a token by ID (full update)
// @Tags tokens
// @Accept json
// @Produce json
// @Param tokenId path string true "Token ID"
// @Param request body v1.PutTokenReq true "Token update request"
// @Success 200 {object} v1.Token "Success"
// @Failure 400 {object} v1.Response "Bad request"
// @Failure 404 {object} v1.Response "Token not found"
// @Failure 500 {object} v1.Response "Internal server error"
// @Router /api/v1/tokens/{tokenId} [put]
func (tc *Controller) Put(c *gin.Context) (any, error) {
	log.C(c).Info("put token begin")
	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Error("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	var r v1.PutTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).WithError(err).Error("put token request parse")
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().Put(c, tokenId, &r)
}

// Patch godoc
// @Summary Partially update a token
// @Description Partially update a token by ID
// @Tags tokens
// @Accept json
// @Produce json
// @Param tokenId path string true "Token ID"
// @Param request body v1.PatchTokenReq true "Token patch request"
// @Success 200 {object} v1.Token "Success"
// @Failure 400 {object} v1.Response "Bad request"
// @Failure 404 {object} v1.Response "Token not found"
// @Failure 500 {object} v1.Response "Internal server error"
// @Router /api/v1/tokens/{tokenId} [patch]
func (tc *Controller) Patch(c *gin.Context) (any, error) {
	log.C(c).Info("patch token begin")
	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Error("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	var r v1.PatchTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).WithError(err).Error("patch token request parse")
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().Patch(c, tokenId, &r)
}

// List godoc
// @Summary List tokens
// @Description Get a list of tokens with pagination and filtering
// @Tags tokens
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1) minimum(1)
// @Param pageSize query int false "Page size" default(10) minimum(1) maximum(200)
// @Param orderBy query []string false "Order by fields"
// @Param targetEmail query string false "Filter by target email"
// @Param mockProvider query string false "Filter by mock provider"
// @Param description query string false "Filter by description"
// @Param expiryAtBegin query int64 false "Filter by expiry time (begin)"
// @Param expiryAtEnd query int64 false "Filter by expiry time (end)"
// @Param lastCalledAtBegin query int64 false "Filter by last called time (begin)"
// @Param lastCalledAtEnd query int64 false "Filter by last called time (end)"
// @Param updatedAtBegin query int64 false "Filter by updated time (begin)"
// @Param updatedAtEnd query int64 false "Filter by updated time (end)"
// @Param status query string false "Filter by status"
// @Success 200 {object} v1.ListTokenResp "Success"
// @Failure 400 {object} v1.Response "Bad request"
// @Failure 500 {object} v1.Response "Internal server error"
// @Router /api/v1/tokens [get]
func (tc *Controller) List(c *gin.Context) (any, error) {
	log.C(c).Info("list token begin")

	var r v1.ListTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).WithError(err).Error("list token request parse")
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().List(c, &r)
}

// Get godoc
// @Summary Get a token
// @Description Get a token by ID
// @Tags tokens
// @Accept json
// @Produce json
// @Param tokenId path string true "Token ID"
// @Success 200 {object} v1.Token "Success"
// @Failure 400 {object} v1.Response "Bad request"
// @Failure 404 {object} v1.Response "Token not found"
// @Failure 500 {object} v1.Response "Internal server error"
// @Router /api/v1/tokens/{tokenId} [get]
func (tc *Controller) Get(c *gin.Context) (any, error) {
	log.C(c).Info("get token begin")
	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Error("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	return tc.b.Token().Get(c, tokenId)
}
