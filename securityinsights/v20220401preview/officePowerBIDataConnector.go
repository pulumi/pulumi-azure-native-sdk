// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents Office Microsoft PowerBI data connector.
type OfficePowerBIDataConnector struct {
	pulumi.CustomResourceState

	// The available data types for the connector.
	DataTypes OfficePowerBIConnectorDataTypesResponseOutput `pulumi:"dataTypes"`
	// Etag of the azure resource
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The kind of the data connector
	// Expected value is 'OfficePowerBI'.
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

// NewOfficePowerBIDataConnector registers a new resource with the given unique name, arguments, and options.
func NewOfficePowerBIDataConnector(ctx *pulumi.Context,
	name string, args *OfficePowerBIDataConnectorArgs, opts ...pulumi.ResourceOption) (*OfficePowerBIDataConnector, error) {
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
	args.Kind = pulumi.String("OfficePowerBI")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:OfficePowerBIDataConnector"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:OfficePowerBIDataConnector"),
		},
	})
	opts = append(opts, aliases)
	var resource OfficePowerBIDataConnector
	err := ctx.RegisterResource("azure-native:securityinsights/v20220401preview:OfficePowerBIDataConnector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOfficePowerBIDataConnector gets an existing OfficePowerBIDataConnector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOfficePowerBIDataConnector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OfficePowerBIDataConnectorState, opts ...pulumi.ResourceOption) (*OfficePowerBIDataConnector, error) {
	var resource OfficePowerBIDataConnector
	err := ctx.ReadResource("azure-native:securityinsights/v20220401preview:OfficePowerBIDataConnector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OfficePowerBIDataConnector resources.
type officePowerBIDataConnectorState struct {
}

type OfficePowerBIDataConnectorState struct {
}

func (OfficePowerBIDataConnectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*officePowerBIDataConnectorState)(nil)).Elem()
}

type officePowerBIDataConnectorArgs struct {
	// Connector ID
	DataConnectorId *string `pulumi:"dataConnectorId"`
	// The available data types for the connector.
	DataTypes OfficePowerBIConnectorDataTypes `pulumi:"dataTypes"`
	// The kind of the data connector
	// Expected value is 'OfficePowerBI'.
	Kind string `pulumi:"kind"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tenant id to connect to, and get the data from.
	TenantId string `pulumi:"tenantId"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a OfficePowerBIDataConnector resource.
type OfficePowerBIDataConnectorArgs struct {
	// Connector ID
	DataConnectorId pulumi.StringPtrInput
	// The available data types for the connector.
	DataTypes OfficePowerBIConnectorDataTypesInput
	// The kind of the data connector
	// Expected value is 'OfficePowerBI'.
	Kind pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The tenant id to connect to, and get the data from.
	TenantId pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (OfficePowerBIDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*officePowerBIDataConnectorArgs)(nil)).Elem()
}

type OfficePowerBIDataConnectorInput interface {
	pulumi.Input

	ToOfficePowerBIDataConnectorOutput() OfficePowerBIDataConnectorOutput
	ToOfficePowerBIDataConnectorOutputWithContext(ctx context.Context) OfficePowerBIDataConnectorOutput
}

func (*OfficePowerBIDataConnector) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficePowerBIDataConnector)(nil)).Elem()
}

func (i *OfficePowerBIDataConnector) ToOfficePowerBIDataConnectorOutput() OfficePowerBIDataConnectorOutput {
	return i.ToOfficePowerBIDataConnectorOutputWithContext(context.Background())
}

func (i *OfficePowerBIDataConnector) ToOfficePowerBIDataConnectorOutputWithContext(ctx context.Context) OfficePowerBIDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficePowerBIDataConnectorOutput)
}

type OfficePowerBIDataConnectorOutput struct{ *pulumi.OutputState }

func (OfficePowerBIDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficePowerBIDataConnector)(nil)).Elem()
}

func (o OfficePowerBIDataConnectorOutput) ToOfficePowerBIDataConnectorOutput() OfficePowerBIDataConnectorOutput {
	return o
}

func (o OfficePowerBIDataConnectorOutput) ToOfficePowerBIDataConnectorOutputWithContext(ctx context.Context) OfficePowerBIDataConnectorOutput {
	return o
}

// The available data types for the connector.
func (o OfficePowerBIDataConnectorOutput) DataTypes() OfficePowerBIConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) OfficePowerBIConnectorDataTypesResponseOutput { return v.DataTypes }).(OfficePowerBIConnectorDataTypesResponseOutput)
}

// Etag of the azure resource
func (o OfficePowerBIDataConnectorOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// The kind of the data connector
// Expected value is 'OfficePowerBI'.
func (o OfficePowerBIDataConnectorOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource
func (o OfficePowerBIDataConnectorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o OfficePowerBIDataConnectorOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tenant id to connect to, and get the data from.
func (o OfficePowerBIDataConnectorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o OfficePowerBIDataConnectorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OfficePowerBIDataConnector) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OfficePowerBIDataConnectorOutput{})
}