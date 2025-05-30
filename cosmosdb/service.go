// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Properties for the database account.
//
// Uses Azure REST API version 2024-11-15.
//
// Other available API versions: 2021-04-01-preview, 2021-07-01-preview, 2021-10-15-preview, 2021-11-15-preview, 2022-02-15-preview, 2022-05-15, 2022-05-15-preview, 2022-08-15, 2022-08-15-preview, 2022-11-15, 2022-11-15-preview, 2023-03-01-preview, 2023-03-15, 2023-03-15-preview, 2023-04-15, 2023-09-15, 2023-09-15-preview, 2023-11-15, 2023-11-15-preview, 2024-02-15-preview, 2024-05-15, 2024-05-15-preview, 2024-08-15, 2024-09-01-preview, 2024-12-01-preview, 2025-04-15, 2025-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cosmosdb [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type Service struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the database account.
	Name pulumi.StringOutput `pulumi:"name"`
	// Services response resource.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210401preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20210701preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20211015preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20211115preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220215preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220515:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220515preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220815:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20220815preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20221115:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20221115preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230301preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230315:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230315preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230415:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230915:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20230915preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20231115:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20231115preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240215preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240515:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240515preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240815:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20240901preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20241115:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20241201preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20250415:Service"),
		},
		{
			Type: pulumi.String("azure-native:cosmosdb/v20250501preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230415:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20230915preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20231115:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20231115preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240215preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240515:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240515preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240815:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20240901preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20241115:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20241201preview:Service"),
		},
		{
			Type: pulumi.String("azure-native:documentdb:Service"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("azure-native:cosmosdb:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("azure-native:cosmosdb:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Properties in ServiceResourceCreateUpdateParameters.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Cosmos DB service name.
	ServiceName *string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Properties in ServiceResourceCreateUpdateParameters.
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Cosmos DB service name.
	ServiceName pulumi.StringPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// The Azure API version of the resource.
func (o ServiceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the database account.
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Services response resource.
func (o ServiceOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Service) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

// The type of Azure resource.
func (o ServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceOutput{})
}
