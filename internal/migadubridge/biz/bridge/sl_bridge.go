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
	"migadu-bridge/internal/pkg/utils"
	"migadu-bridge/pkg/api/enum"
	"migadu-bridge/pkg/api/manage/bridge/sl"
)

func (b *bridgeBiz) SLAliasRandomNew(c *gin.Context, req *sl.AliasRandomNewReq) (*sl.Alias, int, error) {
	if req.Authentication == "" {
		log.C(c).Error("Biz SLAliasRandomNew token is nil")
		return nil, http.StatusUnauthorized, errors.New("token is nil")
	}

	token, err := b.checkToken(c, enum.ProviderEnumSimpleLogin, req.Authentication)
	if err != nil {
		log.C(c).WithError(err).Error("SLAliasRandomNew checkToken")
		return nil, http.StatusUnauthorized, errors.New("check token error")
	}
	logId, err := b.log(c, token, req.Note)
	if err != nil {
		log.C(c).WithError(err).Error("SLAliasRandomNew log")
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
			log.C(c).WithError(err).Error("Biz GetGetRWordsDefault")
			return nil, http.StatusBadRequest, err
		}
	}
	client, err := migadu.MigaduClient()
	if err != nil {
		log.C(c).WithError(err).Error("Biz NewMigaduClient")
		return nil, http.StatusBadRequest, err
	}

	alias, err := client.GetAlias(c, localPart)
	if err != nil && !utils.IsMigaduHttpErr(err, http.StatusNotFound, "no such alias") {
		log.C(c).WithError(err).Error("Biz GetAlias")
		return nil, http.StatusBadRequest, err
	} else if err == nil && alias != nil {
		log.C(c).WithField("localPart", localPart).Error("Biz HasAlias")
		return nil, http.StatusBadRequest, errors.New("alias already exists")
	}

	alias, err = client.NewAlias(c, localPart, []string{token.TargetEmail})
	if err != nil {
		log.C(c).WithError(err).Error("Biz NewAlias")
		return nil, http.StatusBadRequest, err
	}

	if err = b.logAlias(c, logId, alias.Address); err != nil {
		log.C(c).WithError(err).Error("Biz logAlias")
		return nil, http.StatusBadRequest, err
	}

	mailbox := sl.MailBox{
		Email: token.TargetEmail,
		Id:    1,
	}

	now := time.Now()
	return &sl.Alias{
		Alias:             alias.Address,
		CreationDate:      now.Format("2006-01-02 15:04:05-07:00"),
		CreationTimestamp: now.Unix(),
		Email:             alias.Address,
		Name:              localPart,
		Enabled:           true,
		Id:                1,
		Mailbox:           mailbox,
		Mailboxes:         []sl.MailBox{mailbox},
		Note:              req.Note,
	}, http.StatusCreated, nil
}
