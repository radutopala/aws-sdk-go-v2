// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new ReplicationConfigurationTemplate.
func (c *Client) CreateReplicationConfigurationTemplate(ctx context.Context, params *CreateReplicationConfigurationTemplateInput, optFns ...func(*Options)) (*CreateReplicationConfigurationTemplateOutput, error) {
	if params == nil {
		params = &CreateReplicationConfigurationTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReplicationConfigurationTemplate", params, optFns, c.addOperationCreateReplicationConfigurationTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReplicationConfigurationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReplicationConfigurationTemplateInput struct {

	// Request to associate the default Application Migration Service Security group
	// with the Replication Settings template.
	//
	// This member is required.
	AssociateDefaultSecurityGroup *bool

	// Request to configure bandwidth throttling during Replication Settings template
	// creation.
	//
	// This member is required.
	BandwidthThrottling int64

	// Request to create Public IP during Replication Settings template creation.
	//
	// This member is required.
	CreatePublicIP *bool

	// Request to configure data plane routing during Replication Settings template
	// creation.
	//
	// This member is required.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// Request to configure the Staging Disk EBS volume type to "gp2" during
	// Replication Settings template creation.
	//
	// This member is required.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// Request to configure EBS enryption during Replication Settings template
	// creation.
	//
	// This member is required.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// Request to configure the Replication Server instance type during Replication
	// Settings template creation.
	//
	// This member is required.
	ReplicationServerInstanceType *string

	// Request to configure the Replication Server Secuirity group ID during
	// Replication Settings template creation.
	//
	// This member is required.
	ReplicationServersSecurityGroupsIDs []string

	// Request to configure the Staging Area subnet ID during Replication Settings
	// template creation.
	//
	// This member is required.
	StagingAreaSubnetId *string

	// Request to configure Staiging Area tags during Replication Settings template
	// creation.
	//
	// This member is required.
	StagingAreaTags map[string]string

	// Request to use Dedicated Replication Servers during Replication Settings
	// template creation.
	//
	// This member is required.
	UseDedicatedReplicationServer *bool

	// Request to configure an EBS enryption key during Replication Settings template
	// creation.
	EbsEncryptionKeyArn *string

	// Request to configure tags during Replication Settings template creation.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateReplicationConfigurationTemplateOutput struct {

	// Replication Configuration template template ID.
	//
	// This member is required.
	ReplicationConfigurationTemplateID *string

	// Replication Configuration template ARN.
	Arn *string

	// Replication Configuration template associate default Application Migration
	// Service Security group.
	AssociateDefaultSecurityGroup *bool

	// Replication Configuration template bandwidth throtting.
	BandwidthThrottling int64

	// Replication Configuration template create Public IP.
	CreatePublicIP *bool

	// Replication Configuration template data plane routing.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// Replication Configuration template use dedault large Staging Disk type.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// Replication Configuration template EBS encryption.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// Replication Configuration template EBS encryption key ARN.
	EbsEncryptionKeyArn *string

	// Replication Configuration template server instance type.
	ReplicationServerInstanceType *string

	// Replication Configuration template server Security Groups IDs.
	ReplicationServersSecurityGroupsIDs []string

	// Replication Configuration template Staging Area subnet ID.
	StagingAreaSubnetId *string

	// Replication Configuration template Staging Area Tags.
	StagingAreaTags map[string]string

	// Replication Configuration template Tags.
	Tags map[string]string

	// Replication Configuration template use Dedicated Replication Server.
	UseDedicatedReplicationServer *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReplicationConfigurationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateReplicationConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateReplicationConfigurationTemplate{}, middleware.After)
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
	if err = addOpCreateReplicationConfigurationTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReplicationConfigurationTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReplicationConfigurationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "CreateReplicationConfigurationTemplate",
	}
}
