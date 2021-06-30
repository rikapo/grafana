// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	"github.com/aws/smithy-go/middleware"
	"github.com/grafana/grafana/smithy/build/go/grafana/types"
)

// Get alerts.
func (c *Client) ListAlerts(ctx context.Context, params *ListAlertsInput, optFns ...func(*Options)) (*ListAlertsOutput, error) {
	if params == nil {
		params = &ListAlertsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAlerts", params, optFns, addOperationListAlertsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAlertsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAlertsInput struct {
	DashboardIds []int64

	DashboardQuery *string

	DashboardTags []string

	FolderIds []int64
}

type ListAlertsOutput struct {

	// This member is required.
	Items []types.AlertSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAlertsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}
