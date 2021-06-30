// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	"github.com/aws/smithy-go/middleware"
)

// Pause an alert.
func (c *Client) PauseAlert(ctx context.Context, params *PauseAlertInput, optFns ...func(*Options)) (*PauseAlertOutput, error) {
	if params == nil {
		params = &PauseAlertInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PauseAlert", params, optFns, addOperationPauseAlertMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PauseAlertOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PauseAlertInput struct {

	// The alert ID.
	//
	// This member is required.
	Id *string

	// Whether to pause.
	//
	// This member is required.
	Paused bool

	OrgId int64
}

type PauseAlertOutput struct {

	// The alert ID.
	//
	// This member is required.
	Id *string

	// Message.
	//
	// This member is required.
	Message *string

	// Alert state.
	//
	// This member is required.
	State *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPauseAlertMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPauseAlertValidationMiddleware(stack); err != nil {
		return err
	}
	return nil
}
