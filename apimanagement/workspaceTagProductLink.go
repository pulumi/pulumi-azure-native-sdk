// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Tag-product link details.
//
// Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WorkspaceTagProductLink struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Full resource Id of a product.
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceTagProductLink registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceTagProductLink(ctx *pulumi.Context,
	name string, args *WorkspaceTagProductLinkArgs, opts ...pulumi.ResourceOption) (*WorkspaceTagProductLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.TagId == nil {
		return nil, errors.New("invalid value for required argument 'TagId'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:WorkspaceTagProductLink"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceTagProductLink"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230501preview:WorkspaceTagProductLink"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230901preview:WorkspaceTagProductLink"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240501:WorkspaceTagProductLink"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240601preview:WorkspaceTagProductLink"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceTagProductLink
	err := ctx.RegisterResource("azure-native:apimanagement:WorkspaceTagProductLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceTagProductLink gets an existing WorkspaceTagProductLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceTagProductLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceTagProductLinkState, opts ...pulumi.ResourceOption) (*WorkspaceTagProductLink, error) {
	var resource WorkspaceTagProductLink
	err := ctx.ReadResource("azure-native:apimanagement:WorkspaceTagProductLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceTagProductLink resources.
type workspaceTagProductLinkState struct {
}

type WorkspaceTagProductLinkState struct {
}

func (WorkspaceTagProductLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceTagProductLinkState)(nil)).Elem()
}

type workspaceTagProductLinkArgs struct {
	// Full resource Id of a product.
	ProductId string `pulumi:"productId"`
	// Tag-product link identifier. Must be unique in the current API Management service instance.
	ProductLinkId *string `pulumi:"productLinkId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceTagProductLink resource.
type WorkspaceTagProductLinkArgs struct {
	// Full resource Id of a product.
	ProductId pulumi.StringInput
	// Tag-product link identifier. Must be unique in the current API Management service instance.
	ProductLinkId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId pulumi.StringInput
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceTagProductLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceTagProductLinkArgs)(nil)).Elem()
}

type WorkspaceTagProductLinkInput interface {
	pulumi.Input

	ToWorkspaceTagProductLinkOutput() WorkspaceTagProductLinkOutput
	ToWorkspaceTagProductLinkOutputWithContext(ctx context.Context) WorkspaceTagProductLinkOutput
}

func (*WorkspaceTagProductLink) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceTagProductLink)(nil)).Elem()
}

func (i *WorkspaceTagProductLink) ToWorkspaceTagProductLinkOutput() WorkspaceTagProductLinkOutput {
	return i.ToWorkspaceTagProductLinkOutputWithContext(context.Background())
}

func (i *WorkspaceTagProductLink) ToWorkspaceTagProductLinkOutputWithContext(ctx context.Context) WorkspaceTagProductLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceTagProductLinkOutput)
}

type WorkspaceTagProductLinkOutput struct{ *pulumi.OutputState }

func (WorkspaceTagProductLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceTagProductLink)(nil)).Elem()
}

func (o WorkspaceTagProductLinkOutput) ToWorkspaceTagProductLinkOutput() WorkspaceTagProductLinkOutput {
	return o
}

func (o WorkspaceTagProductLinkOutput) ToWorkspaceTagProductLinkOutputWithContext(ctx context.Context) WorkspaceTagProductLinkOutput {
	return o
}

// The Azure API version of the resource.
func (o WorkspaceTagProductLinkOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceTagProductLink) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceTagProductLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceTagProductLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Full resource Id of a product.
func (o WorkspaceTagProductLinkOutput) ProductId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceTagProductLink) pulumi.StringOutput { return v.ProductId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceTagProductLinkOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceTagProductLink) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceTagProductLinkOutput{})
}
