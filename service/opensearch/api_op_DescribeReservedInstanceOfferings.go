// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists available reserved OpenSearch instance offerings.
func (c *Client) DescribeReservedInstanceOfferings(ctx context.Context, params *DescribeReservedInstanceOfferingsInput, optFns ...func(*Options)) (*DescribeReservedInstanceOfferingsOutput, error) {
	if params == nil {
		params = &DescribeReservedInstanceOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedInstanceOfferings", params, optFns, c.addOperationDescribeReservedInstanceOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedInstanceOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for parameters to DescribeReservedInstanceOfferings
type DescribeReservedInstanceOfferingsInput struct {

	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults int32

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string

	// The offering identifier filter value. Use this parameter to show only the
	// available offering that matches the specified reservation identifier.
	ReservedInstanceOfferingId *string

	noSmithyDocumentSerde
}

// Container for results from DescribeReservedInstanceOfferings
type DescribeReservedInstanceOfferingsOutput struct {

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string

	// List of reserved OpenSearch instance offerings
	ReservedInstanceOfferings []types.ReservedInstanceOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedInstanceOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeReservedInstanceOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeReservedInstanceOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedInstanceOfferings(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeReservedInstanceOfferingsAPIClient is a client that implements the
// DescribeReservedInstanceOfferings operation.
type DescribeReservedInstanceOfferingsAPIClient interface {
	DescribeReservedInstanceOfferings(context.Context, *DescribeReservedInstanceOfferingsInput, ...func(*Options)) (*DescribeReservedInstanceOfferingsOutput, error)
}

var _ DescribeReservedInstanceOfferingsAPIClient = (*Client)(nil)

// DescribeReservedInstanceOfferingsPaginatorOptions is the paginator options for
// DescribeReservedInstanceOfferings
type DescribeReservedInstanceOfferingsPaginatorOptions struct {
	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedInstanceOfferingsPaginator is a paginator for
// DescribeReservedInstanceOfferings
type DescribeReservedInstanceOfferingsPaginator struct {
	options   DescribeReservedInstanceOfferingsPaginatorOptions
	client    DescribeReservedInstanceOfferingsAPIClient
	params    *DescribeReservedInstanceOfferingsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedInstanceOfferingsPaginator returns a new
// DescribeReservedInstanceOfferingsPaginator
func NewDescribeReservedInstanceOfferingsPaginator(client DescribeReservedInstanceOfferingsAPIClient, params *DescribeReservedInstanceOfferingsInput, optFns ...func(*DescribeReservedInstanceOfferingsPaginatorOptions)) *DescribeReservedInstanceOfferingsPaginator {
	if params == nil {
		params = &DescribeReservedInstanceOfferingsInput{}
	}

	options := DescribeReservedInstanceOfferingsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedInstanceOfferingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedInstanceOfferingsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeReservedInstanceOfferings page.
func (p *DescribeReservedInstanceOfferingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedInstanceOfferingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeReservedInstanceOfferings(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeReservedInstanceOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeReservedInstanceOfferings",
	}
}
