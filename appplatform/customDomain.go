// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Custom domain resource payload.
//
// Uses Azure REST API version 2024-01-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-05-01-preview.
//
// Other available API versions: 2023-05-01-preview, 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview, 2023-12-01, 2024-05-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native appplatform [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type CustomDomain struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the custom domain resource.
	Properties CustomDomainPropertiesResponseOutput `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCustomDomain registers a new resource with the given unique name, arguments, and options.
func NewCustomDomain(ctx *pulumi.Context,
	name string, args *CustomDomainArgs, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppName == nil {
		return nil, errors.New("invalid value for required argument 'AppName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform/v20200701:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20201101preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210601preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20210901preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220401:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221201:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230101preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230301preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230501preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230701preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20230901preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20231101preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20231201:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20240101preview:CustomDomain"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20240501preview:CustomDomain"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CustomDomain
	err := ctx.RegisterResource("azure-native:appplatform:CustomDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomDomain gets an existing CustomDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomDomainState, opts ...pulumi.ResourceOption) (*CustomDomain, error) {
	var resource CustomDomain
	err := ctx.ReadResource("azure-native:appplatform:CustomDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomDomain resources.
type customDomainState struct {
}

type CustomDomainState struct {
}

func (CustomDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainState)(nil)).Elem()
}

type customDomainArgs struct {
	// The name of the App resource.
	AppName string `pulumi:"appName"`
	// The name of the custom domain resource.
	DomainName *string `pulumi:"domainName"`
	// Properties of the custom domain resource.
	Properties *CustomDomainProperties `pulumi:"properties"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CustomDomain resource.
type CustomDomainArgs struct {
	// The name of the App resource.
	AppName pulumi.StringInput
	// The name of the custom domain resource.
	DomainName pulumi.StringPtrInput
	// Properties of the custom domain resource.
	Properties CustomDomainPropertiesPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the Service resource.
	ServiceName pulumi.StringInput
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customDomainArgs)(nil)).Elem()
}

type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput
}

func (*CustomDomain) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (i *CustomDomain) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i *CustomDomain) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

type CustomDomainOutput struct{ *pulumi.OutputState }

func (CustomDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (o CustomDomainOutput) ToCustomDomainOutput() CustomDomainOutput {
	return o
}

func (o CustomDomainOutput) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return o
}

// The Azure API version of the resource.
func (o CustomDomainOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The name of the resource.
func (o CustomDomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Properties of the custom domain resource.
func (o CustomDomainOutput) Properties() CustomDomainPropertiesResponseOutput {
	return o.ApplyT(func(v *CustomDomain) CustomDomainPropertiesResponseOutput { return v.Properties }).(CustomDomainPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o CustomDomainOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CustomDomain) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o CustomDomainOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomDomain) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomDomainOutput{})
}
