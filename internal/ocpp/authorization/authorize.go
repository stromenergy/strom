package authorization

import (
	"context"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/ocpp/types"
	"github.com/stromenergy/strom/internal/util"
	"github.com/stromenergy/strom/internal/ws"
)

func (s *Authorization) AuthorizeReq(client *ws.Client, message types.Message) {
	authorizeReq, err := UnmarshalAuthorizeReq(message.Payload)

	if err != nil {
		util.LogError("STR055: Error unmarshaling AuthorizeReq", err)
		callError := types.NewMessageCallError(message.UniqueID, types.ErrorCodeFormationViolation, "", types.NoError{})
		callError.Send(client)
		return
	}

	ctx := context.Background()
	idTagInfo := s.GetIDTagInfo(ctx, authorizeReq.IDTag, db.AuthorizationStatusInvalid)

	authorizeConf := AuthorizeConf{
		IDTagInfo: idTagInfo,
	}

	callResult := types.NewMessageCallResult(message.UniqueID, authorizeConf)
	callResult.Send(client)
}
