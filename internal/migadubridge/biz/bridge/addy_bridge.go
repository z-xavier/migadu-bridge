package bridge

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/xid"

	"migadu-bridge/internal/pkg/common"
	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/internal/pkg/migadu"
	"migadu-bridge/internal/pkg/rwords"
	"migadu-bridge/internal/pkg/utils"
	"migadu-bridge/pkg/api/enum"
	"migadu-bridge/pkg/api/manage/bridge/addy"
)

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
		log.C(c).WithError(err).Error("AddyAliases checkToken")
		return nil, http.StatusUnauthorized, errors.New("check token error")
	}
	logId, err := b.log(c, token)
	if err != nil {
		log.C(c).WithError(err).Error("AddyAliases log")
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
				log.C(c).WithError(err).Error("Biz GetRWords")
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
		log.C(c).WithError(err).Error("Biz NewMigaduClient")
		return nil, http.StatusBadRequest, err
	}

	alias, err := client.GetAlias(c, localPart)
	if err != nil && !utils.IsMigaduHttpErr(err, http.StatusNotFound, "no such alias") {
		log.C(c).WithError(err).Error("Biz GetAlias")
		return nil, http.StatusBadRequest, err
	} else if err == nil && alias != nil {
		log.C(c).Errorf("Biz HasAlias localPart: %s", localPart)
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
