package bridge

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/internal/pkg/migadu"
	"migadu-bridge/internal/pkg/model"
	"migadu-bridge/internal/pkg/rwords"
	"migadu-bridge/pkg/api/enum"
	"migadu-bridge/pkg/api/manage/bridge/sl"
)

type BridgeBiz interface {
	SLAliasRandomNew(c *gin.Context, req *sl.AliasRandomNewReq) (*sl.Alias, int, error)
}

type bridgeBiz struct {
	ds store.IStore
}

func New(ds store.IStore) BridgeBiz {
	return &bridgeBiz{ds: ds}
}

func (b bridgeBiz) checkToken(c *gin.Context, mockProvider enum.ProviderEnum, tokenString string) (*model.Token, error) {
	token, err := b.ds.Token().GetByToken(c, mockProvider, tokenString)
	if err != nil {
		log.C(c).Errorf("BridgeBiz CheckToken error: %s", err.Error())
		return nil, err
	}
	return token, nil
}

func (b bridgeBiz) SLAliasRandomNew(c *gin.Context, req *sl.AliasRandomNewReq) (*sl.Alias, int, error) {
	if req.Authentication == "" {
		log.C(c).Errorf("BridgeBiz SLAliasRandomNew token is nil")
		return nil, http.StatusUnauthorized, errors.New("token is nil")
	}

	token, err := b.checkToken(c, enum.ProviderEnumSimpleLogin, req.Authentication)
	if err != nil {
		log.C(c).Errorf("SLAliasRandomNew checkToken error: %s", err.Error())
		return nil, http.StatusUnauthorized, errors.New("check token error")
	}

	parts := strings.Split(token.TargetEmail, "@")
	if len(parts) != 2 {
		log.C(c).Errorf("Invalid email format: %s", token.TargetEmail)
		return nil, http.StatusBadRequest, errors.New("invalid email format")
	}

	if req.Hostname != "" && req.Hostname != parts[1] {
		log.C(c).Errorf("Invalid email format: %s", token.TargetEmail)
		return nil, http.StatusUnauthorized, errors.New("hostname not eq targetEmail")
	}

	var (
		//random    bool
		localPart string
	)
	if req.Word != "" {
		localPart = req.Word
		//random = false
	} else if req.UUID != "" {
		localPart = req.UUID
		//random = false
	} else {
		localPart, err = rwords.GetGetRWordsDefault()
		if err != nil {
			log.C(c).Errorf("BridgeBiz GetGetRWordsDefault error: %s", err.Error())
			return nil, http.StatusBadRequest, err
		}
		//random = true
	}
	client, err := migadu.MigaduClient()
	if err != nil {
		log.C(c).Errorf("BridgeBiz NewMigaduClient error: %s", err.Error())
		return nil, http.StatusBadRequest, err
	}
	// TODO 1. 验证真实报错 2. add retry
	alias, err := client.GetAlias(c, localPart)
	if err != nil {
		log.C(c).Errorf("BridgeBiz GetAlias error: %s", err.Error())
		return nil, http.StatusBadRequest, err
	}

	alias, err = client.NewAlias(c, localPart, []string{req.Note})
	if err != nil {
		log.C(c).Errorf("BridgeBiz NewAlias error: %s", err.Error())
		return nil, http.StatusBadRequest, err
	}
	now := time.Now()
	return &sl.Alias{
		CreationDate:      now.Format("2006-01-02 15:04:05-07:00"),
		CreationTimestamp: now.Unix(),
		Email:             alias.Address,
		Name:              localPart,
		Enabled:           true,
		Id:                1,
		Mailbox: struct {
			Email string `json:"email"`
			Id    int    `json:"id"`
		}{
			Email: token.TargetEmail,
		},
		Note: req.Note,
	}, http.StatusCreated, nil
}
