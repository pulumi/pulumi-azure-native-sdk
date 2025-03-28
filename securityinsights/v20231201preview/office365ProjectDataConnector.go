// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231201preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Office Microsoft Project data connector.
type Office365ProjectDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes Office365ProjectConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'Office365Project'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOffice365ProjectDataConnector registers a new resource with the given unique name, arguments, and options.
func NewOffice365ProjectDataConnector(ctx *pulumi.Context,
	name string, args *Office365ProjectDataConnectorArgs, opts ...pulumi.ResourceOption) (*Office365ProjectDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataTypes == nil {
		return nil, errors.New("invalid value for required argument 'DataTypes'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Office365Project")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240301:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240401preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240901:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20241001preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250101preview:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250301:Office365ProjectDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:Office365ProjectDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Office365ProjectDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20231201preview:Office365ProjectDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOffice365ProjectDataConnector gets an existing Office365ProjectDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOffice365ProjectDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Office365ProjectDataConnectorState, opts ...pulumi.ResourceOption) (*Office365ProjectDataConnector, error) {
	var resource Office365ProjectDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20231201preview:Office365ProjectDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Office365ProjectDataConnector resources.
type office365ProjectDataConnectorState struct {
}

type Office365ProjectDataConnectorState struct {
}

func (Office365ProjectDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*office365ProjectDataConnectorState)(nil)).Elem()
}

type office365ProjectDataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes Office365ProjectConnectorDataTypes `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'Office365Project'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Office365ProjectDataConnector resource.
type Office365ProjectDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes Office365ProjectConnectorDataTypesInput
	// The kind of the data connector
	// Expected value is 'Office365Project'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (Office365ProjectDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*office365ProjectDataConnectorArgs)(nil)).Elem()
}

type Office365ProjectDataConnectorInput interface {
	pulumi.Input

	ToOffice365ProjectDataConnectorOutput() Office365ProjectDataConnectorOutput
	ToOffice365ProjectDataConnectorOutputWithContext(ctx context.Context) Office365ProjectDataConnectorOutput
}

func (*Office365ProjectDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**Office365ProjectDataConnector)(nil)).Elem()
}

func (i *Office365ProjectDataConnector) ToOffice365ProjectDataConnectorOutput() Office365ProjectDataConnectorOutput {
	return i.ToOffice365ProjectDataConnectorOutputWithContext(context.Background())
}

func (i *Office365ProjectDataConnector) ToOffice365ProjectDataConnectorOutputWithContext(ctx context.Context) Office365ProjectDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Office365ProjectDataConnectorOutput)
}

type Office365ProjectDataConnectorOutput struct{ *pulumi.OutputState }

func (Office365ProjectDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Office365ProjectDataConnector)(nil)).Elem()
}

func (o Office365ProjectDataConnectorOutput) ToOffice365ProjectDataConnectorOutput() Office365ProjectDataConnectorOutput {
	return o
}

func (o Office365ProjectDataConnectorOutput) ToOffice365ProjectDataConnectorOutputWithContext(ctx context.Context) Office365ProjectDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o Office365ProjectDataConnectorOutput) DataTypes() Office365ProjectConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) Office365ProjectConnectorDataTypesResponseOutput {
		return v.DataTypes
	}).(Office365ProjectConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o Office365ProjectDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'Office365Project'.
func (o Office365ProjectDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o Office365ProjectDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o Office365ProjectDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o Office365ProjectDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o Office365ProjectDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Office365ProjectDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(Office365ProjectDataConnectorOutput{})
}
