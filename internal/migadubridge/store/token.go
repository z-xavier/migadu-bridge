package store

import (
	"context"
	"errors"

	"github.com/rs/xid"
	"gorm.io/gorm"

	"migadu-bridge/internal/pkg/model"
)

type TokenStore interface {
	Create(ctx context.Context, token *model.Token) (string, error)
	DeleteById(ctx context.Context, id string) error
	UpdateById(ctx context.Context, id string, updates map[string]any) error
	GetById(ctx context.Context, id string) (*model.Token, error)
	GetByToken(ctx context.Context, token string) (*model.Token, error) // TODO scanner token
	ListId(ctx context.Context, cond map[string][]any) ([]int64, error)
	ListWithPage(ctx context.Context, page, pageSize int64, cond map[string][]any, orderBy []any) (int64, []*model.Token, error)
}

type tokenStore struct {
	db *gorm.DB
}

func NewTokenStore(db *gorm.DB) TokenStore {
	return &tokenStore{db}
}

func (t *tokenStore) Create(ctx context.Context, token *model.Token) (string, error) {
	token.Id = xid.New().String()
	if err := t.db.WithContext(ctx).Create(token).Error; err != nil {
		return "", err
	}
	return token.Id, nil
}

func (t *tokenStore) DeleteById(ctx context.Context, id string) error {
	if err := t.db.WithContext(ctx).Model(&model.Token{}).Delete("id", id).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

func (t *tokenStore) UpdateById(ctx context.Context, id string, updates map[string]any) error {
	return t.db.WithContext(ctx).Model(&model.Token{}).Where("id", id).Updates(updates).Error
}

func (t *tokenStore) GetById(ctx context.Context, id string) (*model.Token, error) {
	var token model.Token
	if err := t.db.WithContext(ctx).Model(&model.Token{}).Where("id", id).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (t *tokenStore) GetByToken(ctx context.Context, token string) (*model.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (t *tokenStore) ListId(ctx context.Context, cond map[string][]any) ([]int64, error) {
	db := t.db.WithContext(ctx).Model(&model.Token{})

	if cond != nil {
		for k, v := range cond {
			db = db.Where(k, v...)
		}
	}

	var ids []int64
	if err := db.Find(&ids).Error; err != nil {
		return nil, err
	}
	return ids, nil
}

func (t *tokenStore) ListWithPage(ctx context.Context, page, pageSize int64, cond map[string][]any, orderBy []any) (int64, []*model.Token, error) {
	db := t.db.WithContext(ctx).Model(&model.Token{})

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

	var tokens []*model.Token

	for _, o := range orderBy {
		session = session.Order(o)
	}
	if err := session.
		Offset(int((page - 1) * pageSize)).
		Limit(int(pageSize)).
		Find(&tokens).Error; err != nil {
		return 0, nil, err
	}
	return count, tokens, nil
}
