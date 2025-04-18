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

// Issue Contract details.
//
// Uses Azure REST API version 2022-09-01-preview. In version 2.x of the Azure Native provider, it used API version 2022-08-01.
//
// Other available API versions: 2021-04-01-preview, 2021-08-01, 2021-12-01-preview, 2022-04-01-preview, 2022-08-01, 2023-03-01-preview, 2023-05-01-preview, 2023-09-01-preview, 2024-05-01, 2024-06-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native apimanagement [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ApiIssue struct {
	pulumi.CustomResourceState

	// A resource identifier for the API the issue was created for.
	ApiId pulumi.StringPtrOutput `pulumi:"apiId"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Date and time when the issue was created.
	CreatedDate pulumi.StringPtrOutput `pulumi:"createdDate"`
	// Text describing the issue.
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Status of the issue.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// The issue title.
	Title pulumi.StringOutput `pulumi:"title"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// A resource identifier for the user created the issue.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewApiIssue registers a new resource with the given unique name, arguments, and options.
func NewApiIssue(ctx *pulumi.Context,
	name string, args *ApiIssueArgs, opts ...pulumi.ResourceOption) (*ApiIssue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Title == nil {
		return nil, errors.New("invalid value for required argument 'Title'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:apimanagement/v20170301:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180101:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20180601preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20190101:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210101preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220401preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220801:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20220901preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230301preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230501preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20230901preview:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240501:ApiIssue"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20240601preview:ApiIssue"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ApiIssue
	err := ctx.RegisterResource("azure-native:apimanagement:ApiIssue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiIssue gets an existing ApiIssue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiIssue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiIssueState, opts ...pulumi.ResourceOption) (*ApiIssue, error) {
	var resource ApiIssue
	err := ctx.ReadResource("azure-native:apimanagement:ApiIssue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiIssue resources.
type apiIssueState struct {
}

type ApiIssueState struct {
}

func (ApiIssueState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueState)(nil)).Elem()
}

type apiIssueArgs struct {
	// A resource identifier for the API the issue was created for.
	ApiId string `pulumi:"apiId"`
	// Date and time when the issue was created.
	CreatedDate *string `pulumi:"createdDate"`
	// Text describing the issue.
	Description string `pulumi:"description"`
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId *string `pulumi:"issueId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Status of the issue.
	State *string `pulumi:"state"`
	// The issue title.
	Title string `pulumi:"title"`
	// A resource identifier for the user created the issue.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a ApiIssue resource.
type ApiIssueArgs struct {
	// A resource identifier for the API the issue was created for.
	ApiId pulumi.StringInput
	// Date and time when the issue was created.
	CreatedDate pulumi.StringPtrInput
	// Text describing the issue.
	Description pulumi.StringInput
	// Issue identifier. Must be unique in the current API Management service instance.
	IssueId pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Status of the issue.
	State pulumi.StringPtrInput
	// The issue title.
	Title pulumi.StringInput
	// A resource identifier for the user created the issue.
	UserId pulumi.StringInput
}

func (ApiIssueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiIssueArgs)(nil)).Elem()
}

type ApiIssueInput interface {
	pulumi.Input

	ToApiIssueOutput() ApiIssueOutput
	ToApiIssueOutputWithContext(ctx context.Context) ApiIssueOutput
}

func (*ApiIssue) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIssue)(nil)).Elem()
}

func (i *ApiIssue) ToApiIssueOutput() ApiIssueOutput {
	return i.ToApiIssueOutputWithContext(context.Background())
}

func (i *ApiIssue) ToApiIssueOutputWithContext(ctx context.Context) ApiIssueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiIssueOutput)
}

type ApiIssueOutput struct{ *pulumi.OutputState }

func (ApiIssueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiIssue)(nil)).Elem()
}

func (o ApiIssueOutput) ToApiIssueOutput() ApiIssueOutput {
	return o
}

func (o ApiIssueOutput) ToApiIssueOutputWithContext(ctx context.Context) ApiIssueOutput {
	return o
}

// A resource identifier for the API the issue was created for.
func (o ApiIssueOutput) ApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringPtrOutput { return v.ApiId }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o ApiIssueOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Date and time when the issue was created.
func (o ApiIssueOutput) CreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringPtrOutput { return v.CreatedDate }).(pulumi.StringPtrOutput)
}

// Text describing the issue.
func (o ApiIssueOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The name of the resource
func (o ApiIssueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Status of the issue.
func (o ApiIssueOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

// The issue title.
func (o ApiIssueOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringOutput { return v.Title }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ApiIssueOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A resource identifier for the user created the issue.
func (o ApiIssueOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiIssue) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiIssueOutput{})
}
