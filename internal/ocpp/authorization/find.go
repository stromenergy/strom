package authorization

import (
	"context"

	"github.com/stromenergy/strom/internal/db"
	"github.com/stromenergy/strom/internal/util"
)

func (s *Authorization) GetIDTagInfo(ctx context.Context, token string, defaultStatus db.AuthorizationStatus) IDTagInfo {
	idTagInfo := IDTagInfo{
		Status: defaultStatus,
	}

	if idTag, err := s.repository.GetIDTagByToken(ctx, token); err == nil {
		idTagInfo.Status = idTag.Status

		if idTag.ParentIDTagID.Valid {
			if parentIDTag, err := s.repository.GetIDTag(ctx, idTag.ParentIDTagID.Int64); err == nil {
				idTagInfo.ParentIDTag = util.NilString(parentIDTag.Token)
			}
		}
	}

	return idTagInfo
}