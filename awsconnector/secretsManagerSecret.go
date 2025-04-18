// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package awsconnector

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Microsoft.AwsConnector resource
//
// Uses Azure REST API version 2024-12-01. In version 2.x of the Azure Native provider, it used API version 2024-12-01.
type SecretsManagerSecret struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties SecretsManagerSecretPropertiesResponseOutput `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSecretsManagerSecret registers a new resource with the given unique name, arguments, and options.
func NewSecretsManagerSecret(ctx *pulumi.Context,
	name string, args *SecretsManagerSecretArgs, opts ...pulumi.ResourceOption) (*SecretsManagerSecret, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:awsconnector/v20241201:SecretsManagerSecret"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SecretsManagerSecret
	err := ctx.RegisterResource("azure-native:awsconnector:SecretsManagerSecret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretsManagerSecret gets an existing SecretsManagerSecret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretsManagerSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretsManagerSecretState, opts ...pulumi.ResourceOption) (*SecretsManagerSecret, error) {
	var resource SecretsManagerSecret
	err := ctx.ReadResource("azure-native:awsconnector:SecretsManagerSecret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretsManagerSecret resources.
type secretsManagerSecretState struct {
}

type SecretsManagerSecretState struct {
}

func (SecretsManagerSecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretsManagerSecretState)(nil)).Elem()
}

type secretsManagerSecretArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Name of SecretsManagerSecret
	Name *string `pulumi:"name"`
	// The resource-specific properties for this resource.
	Properties *SecretsManagerSecretProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SecretsManagerSecret resource.
type SecretsManagerSecretArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Name of SecretsManagerSecret
	Name pulumi.StringPtrInput
	// The resource-specific properties for this resource.
	Properties SecretsManagerSecretPropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SecretsManagerSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretsManagerSecretArgs)(nil)).Elem()
}

type SecretsManagerSecretInput interface {
	pulumi.Input

	ToSecretsManagerSecretOutput() SecretsManagerSecretOutput
	ToSecretsManagerSecretOutputWithContext(ctx context.Context) SecretsManagerSecretOutput
}

func (*SecretsManagerSecret) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretsManagerSecret)(nil)).Elem()
}

func (i *SecretsManagerSecret) ToSecretsManagerSecretOutput() SecretsManagerSecretOutput {
	return i.ToSecretsManagerSecretOutputWithContext(context.Background())
}

func (i *SecretsManagerSecret) ToSecretsManagerSecretOutputWithContext(ctx context.Context) SecretsManagerSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretsManagerSecretOutput)
}

type SecretsManagerSecretOutput struct{ *pulumi.OutputState }

func (SecretsManagerSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretsManagerSecret)(nil)).Elem()
}

func (o SecretsManagerSecretOutput) ToSecretsManagerSecretOutput() SecretsManagerSecretOutput {
	return o
}

func (o SecretsManagerSecretOutput) ToSecretsManagerSecretOutputWithContext(ctx context.Context) SecretsManagerSecretOutput {
	return o
}

// The Azure API version of the resource.
func (o SecretsManagerSecretOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretsManagerSecret) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o SecretsManagerSecretOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretsManagerSecret) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SecretsManagerSecretOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretsManagerSecret) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource-specific properties for this resource.
func (o SecretsManagerSecretOutput) Properties() SecretsManagerSecretPropertiesResponseOutput {
	return o.ApplyT(func(v *SecretsManagerSecret) SecretsManagerSecretPropertiesResponseOutput { return v.Properties }).(SecretsManagerSecretPropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o SecretsManagerSecretOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SecretsManagerSecret) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o SecretsManagerSecretOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SecretsManagerSecret) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SecretsManagerSecretOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretsManagerSecret) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretsManagerSecretOutput{})
}
