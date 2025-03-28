// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents OfficeIRM (Microsoft Insider Risk Management) data connector.
type OfficeIRMDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'OfficeIRM'.
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

// NewOfficeIRMDataConnector registers a new resource with the given unique name, arguments, and options.
func NewOfficeIRMDataConnector(ctx *pulumi.Context,
	name string, args *OfficeIRMDataConnectorArgs, opts ...pulumi.ResourceOption) (*OfficeIRMDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	args.Kind = pulumi.String("OfficeIRM")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240301:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240401preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240901:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20241001preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250101preview:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250301:OfficeIRMDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:OfficeIRMDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OfficeIRMDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20230701preview:OfficeIRMDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOfficeIRMDataConnector gets an existing OfficeIRMDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOfficeIRMDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OfficeIRMDataConnectorState, opts ...pulumi.ResourceOption) (*OfficeIRMDataConnector, error) {
	var resource OfficeIRMDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20230701preview:OfficeIRMDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OfficeIRMDataConnector resources.
type officeIRMDataConnectorState struct {
}

type OfficeIRMDataConnectorState struct {
}

func (OfficeIRMDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*officeIRMDataConnectorState)(nil)).Elem()
}

type officeIRMDataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'OfficeIRM'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a OfficeIRMDataConnector resource.
type OfficeIRMDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorPtrInput
	// The kind of the data connector
	// Expected value is 'OfficeIRM'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (OfficeIRMDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*officeIRMDataConnectorArgs)(nil)).Elem()
}

type OfficeIRMDataConnectorInput interface {
	pulumi.Input

	ToOfficeIRMDataConnectorOutput() OfficeIRMDataConnectorOutput
	ToOfficeIRMDataConnectorOutputWithContext(ctx context.Context) OfficeIRMDataConnectorOutput
}

func (*OfficeIRMDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeIRMDataConnector)(nil)).Elem()
}

func (i *OfficeIRMDataConnector) ToOfficeIRMDataConnectorOutput() OfficeIRMDataConnectorOutput {
	return i.ToOfficeIRMDataConnectorOutputWithContext(context.Background())
}

func (i *OfficeIRMDataConnector) ToOfficeIRMDataConnectorOutputWithContext(ctx context.Context) OfficeIRMDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeIRMDataConnectorOutput)
}

type OfficeIRMDataConnectorOutput struct{ *pulumi.OutputState }

func (OfficeIRMDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeIRMDataConnector)(nil)).Elem()
}

func (o OfficeIRMDataConnectorOutput) ToOfficeIRMDataConnectorOutput() OfficeIRMDataConnectorOutput {
	return o
}

func (o OfficeIRMDataConnectorOutput) ToOfficeIRMDataConnectorOutputWithContext(ctx context.Context) OfficeIRMDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o OfficeIRMDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o OfficeIRMDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'OfficeIRM'.
func (o OfficeIRMDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o OfficeIRMDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o OfficeIRMDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o OfficeIRMDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o OfficeIRMDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficeIRMDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OfficeIRMDataConnectorOutput{})
}
