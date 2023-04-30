package authorization

import (
	"context"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Authorization) AuthorizeReq(client *ws.Client, message types.Message) {
	authorizeReq, err := unmarshalAuthorizeReq(message.Payload)

	if err != nil {
		util.LogError("STR055: Error unmarshaling AuthorizeReq", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeFormationViolation, "", types.NoError{})
		callError.Send(client)
		return
	}

	ctx := context.Background()
	idTagInfo := s.GetIDTagInfo(ctx, authorizeReq.IDTag, db.AuthorizationStatusInvalid)

	authorizationConf := AuthorizationConf{
		IDTagInfo: idTagInfo,
	}

	callResult := types.NewMessageCallResult(message.UniqueID, authorizationConf)
	callResult.Send(client)
}
