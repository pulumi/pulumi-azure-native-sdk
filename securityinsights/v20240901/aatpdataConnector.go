// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240901

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents AATP (Azure Advanced Threat Protection) data connector.
type AATPDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorResponsePtrOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'AzureAdvancedThreatProtection'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAATPDataConnector registers a new resource with the given unique name, arguments, and options.
func NewAATPDataConnector(ctx *pulumi.Context,
	name string, args *AATPDataConnectorArgs, opts ...pulumi.ResourceOption) (*AATPDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("AzureAdvancedThreatProtection")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230601preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240301:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240401preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20241001preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250101preview:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250301:AATPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:AATPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource AATPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20240901:AATPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAATPDataConnector gets an existing AATPDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAATPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AATPDataConnectorState, opts ...pulumi.ResourceOption) (*AATPDataConnector, error) {
	var resource AATPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20240901:AATPDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AATPDataConnector resources.
type aatpdataConnectorState struct {
}

type AATPDataConnectorState struct {
}

func (AATPDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*aatpdataConnectorState)(nil)).Elem()
}

type aatpdataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes *AlertsDataTypeOfDataConnector `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'AzureAdvancedThreatProtection'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId *string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a AATPDataConnector resource.
type AATPDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes AlertsDataTypeOfDataConnectorPtrInput
	// The kind of the data connector
	// Expected value is 'AzureAdvancedThreatProtection'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (AATPDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aatpdataConnectorArgs)(nil)).Elem()
}

type AATPDataConnectorInput interface {
	pulumi.Input

	ToAATPDataConnectorOutput() AATPDataConnectorOutput
	ToAATPDataConnectorOutputWithContext(ctx context.Context) AATPDataConnectorOutput
}

func (*AATPDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**AATPDataConnector)(nil)).Elem()
}

func (i *AATPDataConnector) ToAATPDataConnectorOutput() AATPDataConnectorOutput {
	return i.ToAATPDataConnectorOutputWithContext(context.Background())
}

func (i *AATPDataConnector) ToAATPDataConnectorOutputWithContext(ctx context.Context) AATPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AATPDataConnectorOutput)
}

type AATPDataConnectorOutput struct{ *pulumi.OutputState }

func (AATPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AATPDataConnector)(nil)).Elem()
}

func (o AATPDataConnectorOutput) ToAATPDataConnectorOutput() AATPDataConnectorOutput {
	return o
}

func (o AATPDataConnectorOutput) ToAATPDataConnectorOutputWithContext(ctx context.Context) AATPDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o AATPDataConnectorOutput) DataTypes() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyT(func(v *AATPDataConnector) AlertsDataTypeOfDataConnectorResponsePtrOutput { return v.DataTypes }).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
}

// Etag of the azure resource
func (o AATPDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AATPDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'AzureAdvancedThreatProtection'.
func (o AATPDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *AATPDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o AATPDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AATPDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o AATPDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AATPDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o AATPDataConnectorOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AATPDataConnector) pulumi.StringPtrOutput { return v.TenantId }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o AATPDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AATPDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AATPDataConnectorOutput{})
}
