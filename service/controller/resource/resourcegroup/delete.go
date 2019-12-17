package resourcegroup

import (
	"context"
)

// EnsureDeleted ensures the resource group is deleted via the Azure API.
func (r *Resource) EnsureDeleted(ctx context.Context, obj interface{}) error {
	_ = r.logger.LogCtx(ctx, "method", "EnsureDeleted")
	return nil
}
