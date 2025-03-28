// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20241001preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Microsoft Threat Intelligence data connector.
type MSTIDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes MSTIDataConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'MicrosoftThreatIntelligence'.
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

// NewMSTIDataConnector registers a new resource with the given unique name, arguments, and options.
func NewMSTIDataConnector(ctx *pulumi.Context,
	name string, args *MSTIDataConnectorArgs, opts ...pulumi.ResourceOption) (*MSTIDataConnector, error) {
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
	args.Kind = pulumi.String("MicrosoftThreatIntelligence")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240301:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240401preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240901:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250101preview:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250301:MSTIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:MSTIDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource MSTIDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20241001preview:MSTIDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMSTIDataConnector gets an existing MSTIDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMSTIDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MSTIDataConnectorState, opts ...pulumi.ResourceOption) (*MSTIDataConnector, error) {
	var resource MSTIDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20241001preview:MSTIDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MSTIDataConnector resources.
type mstidataConnectorState struct {
}

type MSTIDataConnectorState struct {
}

func (MSTIDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*mstidataConnectorState)(nil)).Elem()
}

type mstidataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes MSTIDataConnectorDataTypes `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'MicrosoftThreatIntelligence'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MSTIDataConnector resource.
type MSTIDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes MSTIDataConnectorDataTypesInput
	// The kind of the data connector
	// Expected value is 'MicrosoftThreatIntelligence'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (MSTIDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mstidataConnectorArgs)(nil)).Elem()
}

type MSTIDataConnectorInput interface {
	pulumi.Input

	ToMSTIDataConnectorOutput() MSTIDataConnectorOutput
	ToMSTIDataConnectorOutputWithContext(ctx context.Context) MSTIDataConnectorOutput
}

func (*MSTIDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnector)(nil)).Elem()
}

func (i *MSTIDataConnector) ToMSTIDataConnectorOutput() MSTIDataConnectorOutput {
	return i.ToMSTIDataConnectorOutputWithContext(context.Background())
}

func (i *MSTIDataConnector) ToMSTIDataConnectorOutputWithContext(ctx context.Context) MSTIDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorOutput)
}

type MSTIDataConnectorOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnector)(nil)).Elem()
}

func (o MSTIDataConnectorOutput) ToMSTIDataConnectorOutput() MSTIDataConnectorOutput {
	return o
}

func (o MSTIDataConnectorOutput) ToMSTIDataConnectorOutputWithContext(ctx context.Context) MSTIDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o MSTIDataConnectorOutput) DataTypes() MSTIDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *MSTIDataConnector) MSTIDataConnectorDataTypesResponseOutput { return v.DataTypes }).(MSTIDataConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o MSTIDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'MicrosoftThreatIntelligence'.
func (o MSTIDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MSTIDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o MSTIDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MSTIDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o MSTIDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MSTIDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o MSTIDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *MSTIDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o MSTIDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MSTIDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MSTIDataConnectorOutput{})
}
