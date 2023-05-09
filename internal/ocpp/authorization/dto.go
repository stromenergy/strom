package authorization

import (
	"encoding/json"
	"errors"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
)

type IDTagInfo struct {
	ExpiryDate  *types.OcppTime        `json:"expiryDate,omitempty"`
	ParentIDTag *string                `json:"parentIdTag,omitempty"`
	Status      db.AuthorizationStatus `json:"status"`
}

type AuthorizeReq struct {
	IDTag string `json:"idTag"`
}

type AuthorizeConf struct {
	IDTagInfo IDTagInfo `json:"idTagInfo"`
}

func UnmarshalAuthorizeReq(payload interface{}) (*AuthorizeReq, error) {
	authorizeReq := &AuthorizeReq{}

	switch typedPayload := payload.(type) {
	case []byte:
		if err := json.Unmarshal(typedPayload, authorizeReq); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("Invalid type")
	}

	return authorizeReq, nil
}
