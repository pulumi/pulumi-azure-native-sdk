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

// ApiRelease details.
//
// Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-09-01-preview.
//
// Other available API versions: 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WorkspaceApiRelease struct {
	pulumi.CustomResourceState

	// Identifier of the API the release belongs to.
	ApiId pulumi.StringPtrOutput `pulumi:"apiId"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The time the API was released. The date conforms to the following format: yyyy-MM-ddTHH:mm:ssZ as specified by the ISO 8601 standard.
	CreatedDateTime pulumi.StringOutput `pulumi:"createdDateTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Release Notes
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the API release was updated.
	UpdatedDateTime pulumi.StringOutput `pulumi:"updatedDateTime"`
}

// NewWorkspaceApiRelease registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceApiRelease(ctx *pulumi.Context,
	name string, args *WorkspaceApiReleaseArgs, opts ...pulumi.ResourceOption) (*WorkspaceApiRelease, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:WorkspaceApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:WorkspaceApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230501preview:WorkspaceApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230901preview:WorkspaceApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240501:WorkspaceApiRelease"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240601preview:WorkspaceApiRelease"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkspaceApiRelease
	err := ctx.RegisterResource("azure-native:apimanagement:WorkspaceApiRelease", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceApiRelease gets an existing WorkspaceApiRelease resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceApiRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceApiReleaseState, opts ...pulumi.ResourceOption) (*WorkspaceApiRelease, error) {
	var resource WorkspaceApiRelease
	err := ctx.ReadResource("azure-native:apimanagement:WorkspaceApiRelease", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceApiRelease resources.
type workspaceApiReleaseState struct {
}

type WorkspaceApiReleaseState struct {
}

func (WorkspaceApiReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceApiReleaseState)(nil)).Elem()
}

type workspaceApiReleaseArgs struct {
	// Identifier of the API the release belongs to.
	ApiId string `pulumi:"apiId"`
	// Release Notes
	Notes *string `pulumi:"notes"`
	// Release identifier within an API. Must be unique in the current API Management service instance.
	ReleaseId *string `pulumi:"releaseId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a WorkspaceApiRelease resource.
type WorkspaceApiReleaseArgs struct {
	// Identifier of the API the release belongs to.
	ApiId pulumi.StringInput
	// Release Notes
	Notes pulumi.StringPtrInput
	// Release identifier within an API. Must be unique in the current API Management service instance.
	ReleaseId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Workspace identifier. Must be unique in the current API Management service instance.
	WorkspaceId pulumi.StringInput
}

func (WorkspaceApiReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceApiReleaseArgs)(nil)).Elem()
}

type WorkspaceApiReleaseInput interface {
	pulumi.Input

	ToWorkspaceApiReleaseOutput() WorkspaceApiReleaseOutput
	ToWorkspaceApiReleaseOutputWithContext(ctx context.Context) WorkspaceApiReleaseOutput
}

func (*WorkspaceApiRelease) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceApiRelease)(nil)).Elem()
}

func (i *WorkspaceApiRelease) ToWorkspaceApiReleaseOutput() WorkspaceApiReleaseOutput {
	return i.ToWorkspaceApiReleaseOutputWithContext(context.Background())
}

func (i *WorkspaceApiRelease) ToWorkspaceApiReleaseOutputWithContext(ctx context.Context) WorkspaceApiReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceApiReleaseOutput)
}

type WorkspaceApiReleaseOutput struct{ *pulumi.OutputState }

func (WorkspaceApiReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkspaceApiRelease)(nil)).Elem()
}

func (o WorkspaceApiReleaseOutput) ToWorkspaceApiReleaseOutput() WorkspaceApiReleaseOutput {
	return o
}

func (o WorkspaceApiReleaseOutput) ToWorkspaceApiReleaseOutputWithContext(ctx context.Context) WorkspaceApiReleaseOutput {
	return o
}

// Identifier of the API the release belongs to.
func (o WorkspaceApiReleaseOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApiRelease) pulumi.StringPtrOutput { return v.ApiId }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o WorkspaceApiReleaseOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApiRelease) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The time the API was released. The date conforms to the following format: yyyy-MM-ddTHH:mm:ssZ as specified by the ISO 8601 standard.
func (o WorkspaceApiReleaseOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApiRelease) pulumi.StringOutput { return v.CreatedDateTime }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceApiReleaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApiRelease) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Release Notes
func (o WorkspaceApiReleaseOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkspaceApiRelease) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceApiReleaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApiRelease) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The time the API release was updated.
func (o WorkspaceApiReleaseOutput) UpdatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkspaceApiRelease) pulumi.StringOutput { return v.UpdatedDateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceApiReleaseOutput{})
}
