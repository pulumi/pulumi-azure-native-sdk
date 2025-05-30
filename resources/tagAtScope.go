// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resources

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Wrapper resource for tags API requests and responses.
//
// Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
//
// Other available API versions: 2020-10-01, 2021-01-01, 2021-04-01, 2022-09-01, 2023-07-01, 2024-07-01, 2024-11-01, 2025-03-01, 2025-04-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native resources [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type TagAtScope struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the tags wrapper resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of tags.
	Properties TagsResponseOutput `pulumi:"properties"`
	// The type of the tags wrapper resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagAtScope registers a new resource with the given unique name, arguments, and options.
func NewTagAtScope(ctx *pulumi.Context,
	name string, args *TagAtScopeArgs, opts ...pulumi.ResourceOption) (*TagAtScope, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resources/v20191001:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200601:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20200801:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20201001:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210101:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20210401:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20220901:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20230701:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20240301:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20240701:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20241101:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20250301:TagAtScope"),
		},
		{
			Type: pulumi.String("azure-native:resources/v20250401:TagAtScope"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource TagAtScope
	err := ctx.RegisterResource("azure-native:resources:TagAtScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagAtScope gets an existing TagAtScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagAtScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagAtScopeState, opts ...pulumi.ResourceOption) (*TagAtScope, error) {
	var resource TagAtScope
	err := ctx.ReadResource("azure-native:resources:TagAtScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagAtScope resources.
type tagAtScopeState struct {
}

type TagAtScopeState struct {
}

func (TagAtScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagAtScopeState)(nil)).Elem()
}

type tagAtScopeArgs struct {
	// The set of tags.
	Properties Tags `pulumi:"properties"`
	// The resource scope.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a TagAtScope resource.
type TagAtScopeArgs struct {
	// The set of tags.
	Properties TagsInput
	// The resource scope.
	Scope pulumi.StringInput
}

func (TagAtScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagAtScopeArgs)(nil)).Elem()
}

type TagAtScopeInput interface {
	pulumi.Input

	ToTagAtScopeOutput() TagAtScopeOutput
	ToTagAtScopeOutputWithContext(ctx context.Context) TagAtScopeOutput
}

func (*TagAtScope) ElementType() reflect.Type {
	return reflect.TypeOf((**TagAtScope)(nil)).Elem()
}

func (i *TagAtScope) ToTagAtScopeOutput() TagAtScopeOutput {
	return i.ToTagAtScopeOutputWithContext(context.Background())
}

func (i *TagAtScope) ToTagAtScopeOutputWithContext(ctx context.Context) TagAtScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagAtScopeOutput)
}

type TagAtScopeOutput struct{ *pulumi.OutputState }

func (TagAtScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagAtScope)(nil)).Elem()
}

func (o TagAtScopeOutput) ToTagAtScopeOutput() TagAtScopeOutput {
	return o
}

func (o TagAtScopeOutput) ToTagAtScopeOutputWithContext(ctx context.Context) TagAtScopeOutput {
	return o
}

// The Azure API version of the resource.
func (o TagAtScopeOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *TagAtScope) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the tags wrapper resource.
func (o TagAtScopeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TagAtScope) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The set of tags.
func (o TagAtScopeOutput) Properties() TagsResponseOutput {
	return o.ApplyT(func(v *TagAtScope) TagsResponseOutput { return v.Properties }).(TagsResponseOutput)
}

// The type of the tags wrapper resource.
func (o TagAtScopeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TagAtScope) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TagAtScopeOutput{})
}
