// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Uses Azure REST API version 2025-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2024-04-01-preview.
//
// Other available API versions: 2024-07-01-preview, 2024-10-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ConnectionRaiBlocklistItem struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// RAI Custom Blocklist Item properties.
	Properties RaiBlocklistItemPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnectionRaiBlocklistItem registers a new resource with the given unique name, arguments, and options.
func NewConnectionRaiBlocklistItem(ctx *pulumi.Context,
	name string, args *ConnectionRaiBlocklistItemArgs, opts ...pulumi.ResourceOption) (*ConnectionRaiBlocklistItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.RaiBlocklistName == nil {
		return nil, errors.New("invalid value for required argument 'RaiBlocklistName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20240401preview:ConnectionRaiBlocklist"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20240401preview:ConnectionRaiBlocklistItem"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20240701preview:ConnectionRaiBlocklistItem"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20241001preview:ConnectionRaiBlocklistItem"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20250101preview:ConnectionRaiBlocklistItem"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices:ConnectionRaiBlocklist"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ConnectionRaiBlocklistItem
	err := ctx.RegisterResource("azure-native:machinelearningservices:ConnectionRaiBlocklistItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionRaiBlocklistItem gets an existing ConnectionRaiBlocklistItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionRaiBlocklistItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionRaiBlocklistItemState, opts ...pulumi.ResourceOption) (*ConnectionRaiBlocklistItem, error) {
	var resource ConnectionRaiBlocklistItem
	err := ctx.ReadResource("azure-native:machinelearningservices:ConnectionRaiBlocklistItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionRaiBlocklistItem resources.
type connectionRaiBlocklistItemState struct {
}

type ConnectionRaiBlocklistItemState struct {
}

func (ConnectionRaiBlocklistItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionRaiBlocklistItemState)(nil)).Elem()
}

type connectionRaiBlocklistItemArgs struct {
	// Friendly name of the workspace connection
	ConnectionName string `pulumi:"connectionName"`
	// RAI Custom Blocklist Item properties.
	Properties RaiBlocklistItemProperties `pulumi:"properties"`
	// Api version used by proxy call
	ProxyApiVersion *string `pulumi:"proxyApiVersion"`
	// Name of the RaiBlocklist Item
	RaiBlocklistItemName *string `pulumi:"raiBlocklistItemName"`
	// The name of the RaiBlocklist.
	RaiBlocklistName string `pulumi:"raiBlocklistName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure Machine Learning Workspace Name
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a ConnectionRaiBlocklistItem resource.
type ConnectionRaiBlocklistItemArgs struct {
	// Friendly name of the workspace connection
	ConnectionName pulumi.StringInput
	// RAI Custom Blocklist Item properties.
	Properties RaiBlocklistItemPropertiesInput
	// Api version used by proxy call
	ProxyApiVersion pulumi.StringPtrInput
	// Name of the RaiBlocklist Item
	RaiBlocklistItemName pulumi.StringPtrInput
	// The name of the RaiBlocklist.
	RaiBlocklistName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Azure Machine Learning Workspace Name
	WorkspaceName pulumi.StringInput
}

func (ConnectionRaiBlocklistItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionRaiBlocklistItemArgs)(nil)).Elem()
}

type ConnectionRaiBlocklistItemInput interface {
	pulumi.Input

	ToConnectionRaiBlocklistItemOutput() ConnectionRaiBlocklistItemOutput
	ToConnectionRaiBlocklistItemOutputWithContext(ctx context.Context) ConnectionRaiBlocklistItemOutput
}

func (*ConnectionRaiBlocklistItem) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionRaiBlocklistItem)(nil)).Elem()
}

func (i *ConnectionRaiBlocklistItem) ToConnectionRaiBlocklistItemOutput() ConnectionRaiBlocklistItemOutput {
	return i.ToConnectionRaiBlocklistItemOutputWithContext(context.Background())
}

func (i *ConnectionRaiBlocklistItem) ToConnectionRaiBlocklistItemOutputWithContext(ctx context.Context) ConnectionRaiBlocklistItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionRaiBlocklistItemOutput)
}

type ConnectionRaiBlocklistItemOutput struct{ *pulumi.OutputState }

func (ConnectionRaiBlocklistItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionRaiBlocklistItem)(nil)).Elem()
}

func (o ConnectionRaiBlocklistItemOutput) ToConnectionRaiBlocklistItemOutput() ConnectionRaiBlocklistItemOutput {
	return o
}

func (o ConnectionRaiBlocklistItemOutput) ToConnectionRaiBlocklistItemOutputWithContext(ctx context.Context) ConnectionRaiBlocklistItemOutput {
	return o
}

// The Azure API version of the resource.
func (o ConnectionRaiBlocklistItemOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionRaiBlocklistItem) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o ConnectionRaiBlocklistItemOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionRaiBlocklistItem) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// RAI Custom Blocklist Item properties.
func (o ConnectionRaiBlocklistItemOutput) Properties() RaiBlocklistItemPropertiesResponseOutput {
	return o.ApplyT(func(v *ConnectionRaiBlocklistItem) RaiBlocklistItemPropertiesResponseOutput { return v.Properties }).(RaiBlocklistItemPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ConnectionRaiBlocklistItemOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ConnectionRaiBlocklistItem) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ConnectionRaiBlocklistItemOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ConnectionRaiBlocklistItem) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionRaiBlocklistItemOutput{})
}
