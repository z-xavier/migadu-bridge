package bridge

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/xid"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/common"
	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/internal/pkg/migadu"
	"migadu-bridge/internal/pkg/model"
	"migadu-bridge/internal/pkg/rwords"
	"migadu-bridge/pkg/api/enum"
	"migadu-bridge/pkg/api/manage/bridge/addy"
	"migadu-bridge/pkg/api/manage/bridge/sl"
)

type Biz interface {
	SLAliasRandomNew(c *gin.Context, req *sl.AliasRandomNewReq) (*sl.Alias, int, error)
}

type bridgeBiz struct {
	ds store.IStore
}

func New(ds store.IStore) Biz {
	return &bridgeBiz{ds: ds}
}

func (b *bridgeBiz) checkToken(c *gin.Context, mockProvider enum.ProviderEnum, tokenString string) (*model.Token, error) {
	token, err := b.ds.Token().GetActiveToken(c, mockProvider, tokenString)
	if err != nil {
		log.C(c).Errorf("Biz CheckToken error: %s", err.Error())
		return nil, err
	}

	if err = b.ds.Token().UpdateById(c, token.Id, map[string]any{
		"last_called_at": c.GetTime(common.XRequestTime),
		"status":         enum.TokenStatusActive,
	}); err != nil {
		log.C(c).Errorf("Biz UpdateToken error: %s", err.Error())
		return nil, err
	}

	return token, nil
}

func (b *bridgeBiz) log(c *gin.Context, token *model.Token) (string, error) {
	logId, err := b.ds.CallLog().Create(c, &model.CallLog{
		TokenId:     token.Id,
		RequestPath: c.Request.URL.Path,
		RequestAt:   c.GetTime(common.XRequestTime),
	})
	if err != nil {
		log.C(c).Errorf("Biz CreateCallLog error: %s", err.Error())
		return "", err
	}
	return logId, nil
}

func (b *bridgeBiz) logAlias(c *gin.Context, logId, genAlias string) error {
	if err := b.ds.CallLog().Update(c, logId, &model.CallLog{
		GenAlias: genAlias,
	}); err != nil {
		log.C(c).Errorf("Biz UpdateCallLog error: %s", err.Error())
		return err
	}
	return nil
}

func (b *bridgeBiz) AddyAliases(c *gin.Context, req *addy.CreateAliasReq) (*addy.CreateAliasResp, int, error) {
	if req.Authorization == "" {
		log.C(c).Errorf("Biz AddyAliases token is nil")
		return nil, http.StatusUnauthorized, errors.New("token is nil")
	}
	if req.XRequestedWith != "XMLHttpRequest" {
		log.C(c).Errorf("Biz AddyAliases XRequestedWith is nil")
		return nil, http.StatusBadRequest, errors.New("XRequestedWith not eq XMLHttpRequest")
	}

	token, err := b.checkToken(c, enum.ProviderEnumAddy, req.Authorization)
	if err != nil {
		log.C(c).Errorf("AddyAliases checkToken error: %s", err.Error())
		return nil, http.StatusUnauthorized, errors.New("check token error")
	}
	logId, err := b.log(c, token)
	if err != nil {
		log.C(c).Errorf("AddyAliases log error: %s", err.Error())
		return nil, http.StatusBadRequest, err
	}

	parts := strings.Split(token.TargetEmail, "@")
	if len(parts) != 2 {
		log.C(c).Errorf("Invalid email format: %s", token.TargetEmail)
		return nil, http.StatusBadRequest, errors.New("invalid email format")
	}

	if req.Domain != "" && req.Domain != parts[1] {
		log.C(c).Errorf("Invalid email format: %s", token.TargetEmail)
		return nil, http.StatusUnauthorized, errors.New("domain not eq targetEmail")
	}

	var localPart string
	switch req.Format {
	case addy.AliasFormatRandomCharacters:
		{
			localPart = xid.New().String()
		}
	case addy.AliasFormatRandomWords:
		{
			localPart, err = rwords.GetGetRWordsDefault()
			if err != nil {
				log.C(c).Errorf("Biz GetRWords error: %s", err.Error())
				return nil, http.StatusBadRequest, err
			}
		}
	case addy.AliasFormatUUID:
		{
			localPart = uuid.NewString()
		}
	case addy.AliasFormatCustom:
		{
			localPart = req.LocalPart
		}
	}
	if localPart == "" {
		log.C(c).Errorf("Biz AddyAliases localPart is nil")
		return nil, http.StatusBadRequest, errors.New("localPart is nil")
	}

	client, err := migadu.MigaduClient()
	if err != nil {
		log.C(c).Errorf("Biz NewMigaduClient error: %s", err.Error())
		return nil, http.StatusBadRequest, err
	}

	alias, err := client.GetAlias(c, localPart)
	if err != nil && !strings.HasPrefix(err.Error(), "no such alias") {
		log.C(c).Errorf("Biz GetAlias error: %s", err.Error())
		return nil, http.StatusBadRequest, err
	} else if err == nil && alias != nil {
		log.C(c).Errorf("Biz HasAlias localPart: %s", localPart)
		return nil, http.StatusBadRequest, errors.New("alias already exists")
	}

	alias, err = client.NewAlias(c, localPart, []string{req.Description})
	if err != nil {
		log.C(c).Errorf("Biz NewAlias error: %s", err.Error())
		return nil, http.StatusBadRequest, err
	}

	if err = b.logAlias(c, logId, alias.Address); err != nil {
		log.C(c).Errorf("Biz logAlias error: %s", err.Error())
		return nil, http.StatusBadRequest, err
	}

	now := time.Now()
	return &addy.CreateAliasResp{
		Data: &addy.Alias{
			Id:          c.GetString(common.XRequestIDKey),
			UserId:      token.Id,
			LocalPart:   localPart,
			Domain:      req.Domain,
			Email:       alias.Address,
			Active:      true,
			Description: req.Description,
			CreatedAt:   now.Format(time.DateTime),
			UpdatedAt:   now.Format(time.DateTime),
		},
	}, http.StatusCreated, nil
}

func (b *bridgeBiz) SLAliasRandomNew(c *gin.Context, req *sl.AliasRandomNewReq) (*sl.Alias, int, error) {
	if req.Authentication == "" {
		log.C(c).Errorf("Biz SLAliasRandomNew token is nil")
		return nil, http.StatusUnauthorized, errors.New("token is nil")
	}

	token, err := b.checkToken(c, enum.ProviderEnumSimpleLogin, req.Authentication)
	if err != nil {
		log.C(c).Errorf("SLAliasRandomNew checkToken error: %s", err.Error())
		return nil, http.StatusUnauthorized, errors.New("check token error")
	}
	logId, err := b.log(c, token)
	if err != nil {
		log.C(c).Errorf("SLAliasRandomNew log error: %s", err.Error())
		return nil, http.StatusBadRequest, err
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

	var localPart string
	if req.Word != "" {
		localPart = req.Word
	} else if req.UUID != "" {
		localPart = req.UUID
	} else {
		localPart, err = rwords.GetGetRWordsDefault()
		if err != nil {
			log.C(c).Errorf("Biz GetGetRWordsDefault error: %s", err.Error())
			return nil, http.StatusBadRequest, err
		}
	}
	client, err := migadu.MigaduClient()
	if err != nil {
		log.C(c).Errorf("Biz NewMigaduClient error: %s", err.Error())
		return nil, http.StatusBadRequest, err
	}

	alias, err := client.GetAlias(c, localPart)
	if err != nil && !strings.HasPrefix(err.Error(), "no such alias") {
		log.C(c).Errorf("Biz GetAlias error: %s", err.Error())
		return nil, http.StatusBadRequest, err
	} else if err == nil && alias != nil {
		log.C(c).Errorf("Biz HasAlias localPart: %s", localPart)
		return nil, http.StatusBadRequest, errors.New("alias already exists")
	}

	alias, err = client.NewAlias(c, localPart, []string{req.Note})
	if err != nil {
		log.C(c).Errorf("Biz NewAlias error: %s", err.Error())
		return nil, http.StatusBadRequest, err
	}

	if err = b.logAlias(c, logId, alias.Address); err != nil {
		log.C(c).Errorf("Biz logAlias error: %s", err.Error())
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
