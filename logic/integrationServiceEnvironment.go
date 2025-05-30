// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logic

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The integration service environment.
//
// Uses Azure REST API version 2019-05-01. In version 2.x of the Azure Native provider, it used API version 2019-05-01.
type IntegrationServiceEnvironment struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Managed service identity properties.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Gets the resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The integration service environment properties.
	Properties IntegrationServiceEnvironmentPropertiesResponseOutput `pulumi:"properties"`
	// The sku.
	Sku IntegrationServiceEnvironmentSkuResponsePtrOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Gets the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationServiceEnvironment registers a new resource with the given unique name, arguments, and options.
func NewIntegrationServiceEnvironment(ctx *pulumi.Context,
	name string, args *IntegrationServiceEnvironmentArgs, opts ...pulumi.ResourceOption) (*IntegrationServiceEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroup == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20190501:IntegrationServiceEnvironment"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource IntegrationServiceEnvironment
	err := ctx.RegisterResource("azure-native:logic:IntegrationServiceEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationServiceEnvironment gets an existing IntegrationServiceEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationServiceEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationServiceEnvironmentState, opts ...pulumi.ResourceOption) (*IntegrationServiceEnvironment, error) {
	var resource IntegrationServiceEnvironment
	err := ctx.ReadResource("azure-native:logic:IntegrationServiceEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationServiceEnvironment resources.
type integrationServiceEnvironmentState struct {
}

type IntegrationServiceEnvironmentState struct {
}

func (IntegrationServiceEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationServiceEnvironmentState)(nil)).Elem()
}

type integrationServiceEnvironmentArgs struct {
	// Managed service identity properties.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The integration service environment name.
	IntegrationServiceEnvironmentName *string `pulumi:"integrationServiceEnvironmentName"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The integration service environment properties.
	Properties *IntegrationServiceEnvironmentProperties `pulumi:"properties"`
	// The resource group.
	ResourceGroup string `pulumi:"resourceGroup"`
	// The sku.
	Sku *IntegrationServiceEnvironmentSku `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IntegrationServiceEnvironment resource.
type IntegrationServiceEnvironmentArgs struct {
	// Managed service identity properties.
	Identity ManagedServiceIdentityPtrInput
	// The integration service environment name.
	IntegrationServiceEnvironmentName pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The integration service environment properties.
	Properties IntegrationServiceEnvironmentPropertiesPtrInput
	// The resource group.
	ResourceGroup pulumi.StringInput
	// The sku.
	Sku IntegrationServiceEnvironmentSkuPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (IntegrationServiceEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationServiceEnvironmentArgs)(nil)).Elem()
}

type IntegrationServiceEnvironmentInput interface {
	pulumi.Input

	ToIntegrationServiceEnvironmentOutput() IntegrationServiceEnvironmentOutput
	ToIntegrationServiceEnvironmentOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentOutput
}

func (*IntegrationServiceEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironment)(nil)).Elem()
}

func (i *IntegrationServiceEnvironment) ToIntegrationServiceEnvironmentOutput() IntegrationServiceEnvironmentOutput {
	return i.ToIntegrationServiceEnvironmentOutputWithContext(context.Background())
}

func (i *IntegrationServiceEnvironment) ToIntegrationServiceEnvironmentOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationServiceEnvironmentOutput)
}

type IntegrationServiceEnvironmentOutput struct{ *pulumi.OutputState }

func (IntegrationServiceEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationServiceEnvironment)(nil)).Elem()
}

func (o IntegrationServiceEnvironmentOutput) ToIntegrationServiceEnvironmentOutput() IntegrationServiceEnvironmentOutput {
	return o
}

func (o IntegrationServiceEnvironmentOutput) ToIntegrationServiceEnvironmentOutputWithContext(ctx context.Context) IntegrationServiceEnvironmentOutput {
	return o
}

// The Azure API version of the resource.
func (o IntegrationServiceEnvironmentOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Managed service identity properties.
func (o IntegrationServiceEnvironmentOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The resource location.
func (o IntegrationServiceEnvironmentOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Gets the resource name.
func (o IntegrationServiceEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The integration service environment properties.
func (o IntegrationServiceEnvironmentOutput) Properties() IntegrationServiceEnvironmentPropertiesResponseOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) IntegrationServiceEnvironmentPropertiesResponseOutput {
		return v.Properties
	}).(IntegrationServiceEnvironmentPropertiesResponseOutput)
}

// The sku.
func (o IntegrationServiceEnvironmentOutput) Sku() IntegrationServiceEnvironmentSkuResponsePtrOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) IntegrationServiceEnvironmentSkuResponsePtrOutput { return v.Sku }).(IntegrationServiceEnvironmentSkuResponsePtrOutput)
}

// The resource tags.
func (o IntegrationServiceEnvironmentOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Gets the resource type.
func (o IntegrationServiceEnvironmentOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationServiceEnvironment) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(IntegrationServiceEnvironmentOutput{})
}
