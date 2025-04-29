package bridge

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/internal/pkg/migadu"
	"migadu-bridge/internal/pkg/rwords"
	"migadu-bridge/pkg/api/enum"
	"migadu-bridge/pkg/api/manage/bridge/sl"
)

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
		Mailbox: sl.MailBox{
			Email: token.TargetEmail,
		},
		Note: req.Note,
	}, http.StatusCreated, nil
}
