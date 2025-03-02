package store

import (
	"context"

	"github.com/rs/xid"
	"gorm.io/gorm"

	"migadu-bridge/internal/pkg/model"
)

type CallLogStore interface {
	Create(ctx context.Context, callLog *model.CallLog) (string, error)
	ListAndTokenWithPage(ctx context.Context, page, pageSize int64, cond map[string][]any, orderBy []any) (int64, []map[string]any, error)
	ListByTokenId(ctx context.Context, tokenIdList []string) ([]*model.CallLog, error)
}

type callLogStore struct {
	db *gorm.DB
}

func NewCallLogStore(db *gorm.DB) CallLogStore {
	return &callLogStore{db}
}

func (t *callLogStore) Create(ctx context.Context, callLog *model.CallLog) (string, error) {
	callLog.Id = xid.New().String()
	if err := t.db.WithContext(ctx).Create(callLog).Error; err != nil {
		return "", err
	}
	return callLog.Id, nil
}

func (t *callLogStore) ListAndTokenWithPage(ctx context.Context, page, pageSize int64, cond map[string][]any, orderBy []any) (int64, []map[string]any, error) {
	db := t.db.WithContext(ctx).Model(&model.CallLog{}).
		Joins("LEFT JOIN `tokens` ON `call_logs`.`token_id` = `tokens`.`id` AND `tokens`.`deleted_at` IS NULL").
		Select("call_logs.id, call_logs.request_path, call_logs.gen_alias, call_logs.request_ip, call_logs.request_at, tokens.id, tokens.target_email, tokens.mock_provider")

	if cond != nil {
		for k, v := range cond {
			db = db.Where(k, v...)
		}
	}

	session := db.Session(&gorm.Session{})

	var count int64
	if err := session.Count(&count).Error; err != nil {
		return 0, nil, err
	}
	if count == 0 {
		return 0, nil, gorm.ErrRecordNotFound
	}

	var result []map[string]any

	for _, o := range orderBy {
		session = session.Order(o)
	}
	if err := session.
		Offset(int((page - 1) * pageSize)).
		Limit(int(pageSize)).
		Find(&result).Error; err != nil {
		return 0, nil, err
	}
	return count, result, nil
}

func (t *callLogStore) ListByTokenId(ctx context.Context, tokenIdList []string) ([]*model.CallLog, error) {
	var callLogList []*model.CallLog
	if err := t.db.WithContext(ctx).
		Where("token_id IN ?", tokenIdList).
		Find(&callLogList).Error; err != nil {
		return nil, err
	}
	return callLogList, nil
}
