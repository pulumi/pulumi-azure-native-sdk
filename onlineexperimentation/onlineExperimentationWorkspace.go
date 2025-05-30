// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onlineexperimentation

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An online experimentation workspace resource.
//
// Uses Azure REST API version 2025-05-31-preview.
type OnlineExperimentationWorkspace struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The managed service identities assigned to this resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties OnlineExperimentationWorkspacePropertiesResponseOutput `pulumi:"properties"`
	// The SKU (Stock Keeping Unit) assigned to this resource.
	Sku OnlineExperimentationWorkspaceSkuResponsePtrOutput `pulumi:"sku"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOnlineExperimentationWorkspace registers a new resource with the given unique name, arguments, and options.
func NewOnlineExperimentationWorkspace(ctx *pulumi.Context,
	name string, args *OnlineExperimentationWorkspaceArgs, opts ...pulumi.ResourceOption) (*OnlineExperimentationWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:onlineexperimentation/v20250531preview:OnlineExperimentationWorkspace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource OnlineExperimentationWorkspace
	err := ctx.RegisterResource("azure-native:onlineexperimentation:OnlineExperimentationWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOnlineExperimentationWorkspace gets an existing OnlineExperimentationWorkspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOnlineExperimentationWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OnlineExperimentationWorkspaceState, opts ...pulumi.ResourceOption) (*OnlineExperimentationWorkspace, error) {
	var resource OnlineExperimentationWorkspace
	err := ctx.ReadResource("azure-native:onlineexperimentation:OnlineExperimentationWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OnlineExperimentationWorkspace resources.
type onlineExperimentationWorkspaceState struct {
}

type OnlineExperimentationWorkspaceState struct {
}

func (OnlineExperimentationWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*onlineExperimentationWorkspaceState)(nil)).Elem()
}

type onlineExperimentationWorkspaceArgs struct {
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The resource-specific properties for this resource.
	Properties *OnlineExperimentationWorkspaceProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU (Stock Keeping Unit) assigned to this resource.
	Sku *OnlineExperimentationWorkspaceSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the OnlineExperimentationWorkspace
	WorkspaceName *string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a OnlineExperimentationWorkspace resource.
type OnlineExperimentationWorkspaceArgs struct {
	// The managed service identities assigned to this resource.
	Identity ManagedServiceIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties OnlineExperimentationWorkspacePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The SKU (Stock Keeping Unit) assigned to this resource.
	Sku OnlineExperimentationWorkspaceSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the OnlineExperimentationWorkspace
	WorkspaceName pulumi.StringPtrInput
}

func (OnlineExperimentationWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*onlineExperimentationWorkspaceArgs)(nil)).Elem()
}

type OnlineExperimentationWorkspaceInput interface {
	pulumi.Input

	ToOnlineExperimentationWorkspaceOutput() OnlineExperimentationWorkspaceOutput
	ToOnlineExperimentationWorkspaceOutputWithContext(ctx context.Context) OnlineExperimentationWorkspaceOutput
}

func (*OnlineExperimentationWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((**OnlineExperimentationWorkspace)(nil)).Elem()
}

func (i *OnlineExperimentationWorkspace) ToOnlineExperimentationWorkspaceOutput() OnlineExperimentationWorkspaceOutput {
	return i.ToOnlineExperimentationWorkspaceOutputWithContext(context.Background())
}

func (i *OnlineExperimentationWorkspace) ToOnlineExperimentationWorkspaceOutputWithContext(ctx context.Context) OnlineExperimentationWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OnlineExperimentationWorkspaceOutput)
}

type OnlineExperimentationWorkspaceOutput struct{ *pulumi.OutputState }

func (OnlineExperimentationWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OnlineExperimentationWorkspace)(nil)).Elem()
}

func (o OnlineExperimentationWorkspaceOutput) ToOnlineExperimentationWorkspaceOutput() OnlineExperimentationWorkspaceOutput {
	return o
}

func (o OnlineExperimentationWorkspaceOutput) ToOnlineExperimentationWorkspaceOutputWithContext(ctx context.Context) OnlineExperimentationWorkspaceOutput {
	return o
}

// The Azure API version of the resource.
func (o OnlineExperimentationWorkspaceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineExperimentationWorkspace) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The managed service identities assigned to this resource.
func (o OnlineExperimentationWorkspaceOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *OnlineExperimentationWorkspace) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o OnlineExperimentationWorkspaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineExperimentationWorkspace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o OnlineExperimentationWorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineExperimentationWorkspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o OnlineExperimentationWorkspaceOutput) Properties() OnlineExperimentationWorkspacePropertiesResponseOutput {
	return o.ApplyT(func(v *OnlineExperimentationWorkspace) OnlineExperimentationWorkspacePropertiesResponseOutput {
		return v.Properties
	}).(OnlineExperimentationWorkspacePropertiesResponseOutput)
}

// The SKU (Stock Keeping Unit) assigned to this resource.
func (o OnlineExperimentationWorkspaceOutput) Sku() OnlineExperimentationWorkspaceSkuResponsePtrOutput {
	return o.ApplyT(func(v *OnlineExperimentationWorkspace) OnlineExperimentationWorkspaceSkuResponsePtrOutput {
		return v.Sku
	}).(OnlineExperimentationWorkspaceSkuResponsePtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o OnlineExperimentationWorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *OnlineExperimentationWorkspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o OnlineExperimentationWorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OnlineExperimentationWorkspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o OnlineExperimentationWorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OnlineExperimentationWorkspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(OnlineExperimentationWorkspaceOutput{})
}
