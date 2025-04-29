package call_log

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/pkg/api/enum"
	v1 "migadu-bridge/pkg/api/manage/v1"
)

type Biz interface {
	List(*gin.Context, *v1.ListCallLogReq) (*v1.ListCallLogResp, error)
}

type callLogBiz struct {
	ds store.IStore
}

func New(ds store.IStore) Biz {
	return &callLogBiz{ds: ds}
}

func (c *callLogBiz) List(context *gin.Context, req *v1.ListCallLogReq) (*v1.ListCallLogResp, error) {
	if len(req.OrderBy) == 0 {
		req.OrderBy = []string{"updated_at:desc"}
	}

	cond := map[string][]any{}
	if req.TargetEmail != "" {
		cond["tokens.target_email like ?"] = []any{"%" + req.TargetEmail + "%"}
	}

	if req.MockProvider != "" {
		cond["tokens.mock_provider = ?"] = []any{req.MockProvider}
	}

	if req.RequestPath != "" {
		cond["call_logs.request_path like ?"] = []any{"%" + req.RequestPath + "%"}
	}

	if req.RequestIp != "" {
		cond["call_logs.request_ip = ?"] = []any{req.RequestIp}
	}

	if req.RequestAtBegin != 0 {
		cond["call_logs.updated_at >= ?"] = []any{time.Unix(req.RequestAtBegin, 0)}
	}

	if req.RequestAtEnd != 0 {
		cond["call_logs.updated_at <= ?"] = []any{time.Unix(req.RequestAtEnd, 0)}
	}

	var orderBy []any
	for _, o := range req.OrderBy {
		parts := strings.Split(o, ":")
		if len(parts) == 2 {
			tableName := ""
			if parts[0] == "target_email" || parts[0] == "mock_provider" {
				tableName = "tokens"
			} else {
				tableName = "call_logs"
			}

			orderBy = append(orderBy, fmt.Sprintf("%s.%s %s", tableName, parts[0], parts[1]))
		}
	}

	count, dbModels, err := c.ds.CallLog().ListAndTokenWithPage(context, req.Page, req.PageSize, cond, orderBy)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	callLogs := make([]*v1.CallLog, 0, len(dbModels))
	for _, tmp := range dbModels {
		callLogs = append(callLogs, &v1.CallLog{
			Id:           cast.ToString(tmp["call_logs.id"]),
			TokenId:      cast.ToString(tmp["tokens.id"]),
			TargetEmail:  cast.ToString(tmp["tokens.target_email"]),
			MockProvider: enum.ProviderEnum(cast.ToString(tmp["tokens.mock_provider"])),
			GenAlias:     cast.ToString(tmp["call_logs.gen_alias"]),
			RequestPath:  cast.ToString(tmp["call_logs.request_path"]),
			RequestIp:    cast.ToString(tmp["call_logs.request_ip"]),
			RequestAt:    cast.ToTime(tmp["call_logs.request_at"]).Unix(),
		})
	}

	return &v1.ListCallLogResp{
		Page: v1.Page{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    count,
		},
		List: callLogs,
	}, nil
}
