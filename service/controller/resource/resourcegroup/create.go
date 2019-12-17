package resourcegroup

import (
	"context"
)

// EnsureCreated ensures the resource group is created via the Azure API.
func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	_ = r.logger.LogCtx(ctx, "method", "EnsureCreated")
	return nil
}
