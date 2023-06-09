package gateway

import (
	"context"

	"go.mod/domain/common"
)

type IMQPublisher interface {
	ServiceUpdated(ctx context.Context, event *common.Event) error
}
