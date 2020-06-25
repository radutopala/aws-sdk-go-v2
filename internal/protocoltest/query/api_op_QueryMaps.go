// Code generated by smithy-go-codegen DO NOT EDIT.
package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This test serializes simple and complex maps.
func (c *Client) QueryMaps(ctx context.Context, params *QueryMapsInput, optFns ...func(*Options)) (*QueryMapsOutput, error) {
	stack := middleware.NewStack("QueryMaps", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	stack.Initialize.Add(newServiceMetadataMiddleware_opQueryMaps(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "QueryMaps",
			Err:           err,
		}
	}
	out := result.(*QueryMapsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type QueryMapsInput struct {
	MapArg                  map[string]*string
	RenamedMapArg           map[string]*string
	ComplexMapArg           map[string]*types.GreetingStruct
	MapWithXmlMemberName    map[string]*string
	FlattenedMap            map[string]*string
	FlattenedMapWithXmlName map[string]*string
	MapOfLists              map[string][]*string
}

type QueryMapsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func newServiceMetadataMiddleware_opQueryMaps(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Query Protocol",
		ServiceID:      "queryprotocol",
		EndpointPrefix: "queryprotocol",
		OperationName:  "QueryMaps",
	}
}
