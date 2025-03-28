// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Google Cloud Platform data connector.
type GCPDataConnector struct {
	pulumi.CustomResourceState

	// The auth section of the connector.
	Auth GCPAuthPropertiesResponseOutput `pulumi:"auth"`
	// The name of the connector definition that represents the UI config.
	ConnectorDefinitionName pulumi.StringOutput `pulumi:"connectorDefinitionName"`
	// The configuration of the destination of the data.
	DcrConfig DCRConfigurationResponsePtrOutput `pulumi:"dcrConfig"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'GCP'.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The request section of the connector.
	Request GCPRequestPropertiesResponseOutput `pulumi:"request"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGCPDataConnector registers a new resource with the given unique name, arguments, and options.
func NewGCPDataConnector(ctx *pulumi.Context,
	name string, args *GCPDataConnectorArgs, opts ...pulumi.ResourceOption) (*GCPDataConnector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Auth == nil {
		return nil, errors.New("invalid value for required argument 'Auth'")
	}
	if args.ConnectorDefinitionName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorDefinitionName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.Request == nil {
		return nil, errors.New("invalid value for required argument 'Request'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("GCP")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221201preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230201preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230301preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230401preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230501preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230701preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230801preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20230901preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231001preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231101:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20231201preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240101preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240301:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240401preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20240901:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20241001preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250101preview:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20250301:GCPDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:GCPDataConnector"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource GCPDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20230601preview:GCPDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGCPDataConnector gets an existing GCPDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGCPDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GCPDataConnectorState, opts ...pulumi.ResourceOption) (*GCPDataConnector, error) {
	var resource GCPDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20230601preview:GCPDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GCPDataConnector resources.
type gcpdataConnectorState struct {
}

type GCPDataConnectorState struct {
}

func (GCPDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpdataConnectorState)(nil)).Elem()
}

type gcpdataConnectorArgs struct {
	// The auth section of the connector.
	Auth GCPAuthProperties `pulumi:"auth"`
	// The name of the connector definition that represents the UI config.
	ConnectorDefinitionName string `pulumi:"connectorDefinitionName"`
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The configuration of the destination of the data.
	DcrConfig *DCRConfiguration `pulumi:"dcrConfig"`
	// The kind of the data connector
	// Expected value is 'GCP'.
	Kind string `pulumi:"kind"`
	// The request section of the connector.
	Request GCPRequestProperties `pulumi:"request"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a GCPDataConnector resource.
type GCPDataConnectorArgs struct {
	// The auth section of the connector.
	Auth GCPAuthPropertiesInput
	// The name of the connector definition that represents the UI config.
	ConnectorDefinitionName pulumi.StringInput
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The configuration of the destination of the data.
	DcrConfig DCRConfigurationPtrInput
	// The kind of the data connector
	// Expected value is 'GCP'.
	Kind pulumi.StringInput
	// The request section of the connector.
	Request GCPRequestPropertiesInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (GCPDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpdataConnectorArgs)(nil)).Elem()
}

type GCPDataConnectorInput interface {
	pulumi.Input

	ToGCPDataConnectorOutput() GCPDataConnectorOutput
	ToGCPDataConnectorOutputWithContext(ctx context.Context) GCPDataConnectorOutput
}

func (*GCPDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**GCPDataConnector)(nil)).Elem()
}

func (i *GCPDataConnector) ToGCPDataConnectorOutput() GCPDataConnectorOutput {
	return i.ToGCPDataConnectorOutputWithContext(context.Background())
}

func (i *GCPDataConnector) ToGCPDataConnectorOutputWithContext(ctx context.Context) GCPDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GCPDataConnectorOutput)
}

type GCPDataConnectorOutput struct{ *pulumi.OutputState }

func (GCPDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GCPDataConnector)(nil)).Elem()
}

func (o GCPDataConnectorOutput) ToGCPDataConnectorOutput() GCPDataConnectorOutput {
	return o
}

func (o GCPDataConnectorOutput) ToGCPDataConnectorOutputWithContext(ctx context.Context) GCPDataConnectorOutput {
	return o
}

// The auth section of the connector.
func (o GCPDataConnectorOutput) Auth() GCPAuthPropertiesResponseOutput {
	return o.ApplyT(func(v *GCPDataConnector) GCPAuthPropertiesResponseOutput { return v.Auth }).(GCPAuthPropertiesResponseOutput)
}

// The name of the connector definition that represents the UI config.
func (o GCPDataConnectorOutput) ConnectorDefinitionName() pulumi.StringOutput {
	return o.ApplyT(func(v *GCPDataConnector) pulumi.StringOutput { return v.ConnectorDefinitionName }).(pulumi.StringOutput)
}

// The configuration of the destination of the data.
func (o GCPDataConnectorOutput) DcrConfig() DCRConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GCPDataConnector) DCRConfigurationResponsePtrOutput { return v.DcrConfig }).(DCRConfigurationResponsePtrOutput)
}

// Etag of the azure resource
func (o GCPDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GCPDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'GCP'.
func (o GCPDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *GCPDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o GCPDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GCPDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The request section of the connector.
func (o GCPDataConnectorOutput) Request() GCPRequestPropertiesResponseOutput {
	return o.ApplyT(func(v *GCPDataConnector) GCPRequestPropertiesResponseOutput { return v.Request }).(GCPRequestPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o GCPDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GCPDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o GCPDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GCPDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GCPDataConnectorOutput{})
}
