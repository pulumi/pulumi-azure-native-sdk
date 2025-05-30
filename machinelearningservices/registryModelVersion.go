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

// Azure Resource Manager resource envelope.
//
// Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
//
// Other available API versions: 2022-10-01-preview, 2022-12-01-preview, 2023-02-01-preview, 2023-04-01, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type RegistryModelVersion struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// [Required] Additional attributes of the entity.
	ModelVersionProperties ModelVersionResponseOutput `pulumi:"modelVersionProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistryModelVersion registers a new resource with the given unique name, arguments, and options.
func NewRegistryModelVersion(ctx *pulumi.Context,
	name string, args *RegistryModelVersionArgs, opts ...pulumi.ResourceOption) (*RegistryModelVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ModelName == nil {
		return nil, errors.New("invalid value for required argument 'ModelName'")
	}
	if args.ModelVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'ModelVersionProperties'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.ModelVersionProperties = args.ModelVersionProperties.ToModelVersionTypeOutput().ApplyT(func(v ModelVersionType) ModelVersionType { return *v.Defaults() }).(ModelVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20240101preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20240401:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20240401preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20240701preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20241001:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20241001preview:RegistryModelVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20250101preview:RegistryModelVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource RegistryModelVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices:RegistryModelVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistryModelVersion gets an existing RegistryModelVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistryModelVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryModelVersionState, opts ...pulumi.ResourceOption) (*RegistryModelVersion, error) {
	var resource RegistryModelVersion
	err := ctx.ReadResource("azure-native:machinelearningservices:RegistryModelVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegistryModelVersion resources.
type registryModelVersionState struct {
}

type RegistryModelVersionState struct {
}

func (RegistryModelVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryModelVersionState)(nil)).Elem()
}

type registryModelVersionArgs struct {
	// Container name.
	ModelName string `pulumi:"modelName"`
	// [Required] Additional attributes of the entity.
	ModelVersionProperties ModelVersionType `pulumi:"modelVersionProperties"`
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName string `pulumi:"registryName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a RegistryModelVersion resource.
type RegistryModelVersionArgs struct {
	// Container name.
	ModelName pulumi.StringInput
	// [Required] Additional attributes of the entity.
	ModelVersionProperties ModelVersionTypeInput
	// Name of Azure Machine Learning registry. This is case-insensitive
	RegistryName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier.
	Version pulumi.StringPtrInput
}

func (RegistryModelVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryModelVersionArgs)(nil)).Elem()
}

type RegistryModelVersionInput interface {
	pulumi.Input

	ToRegistryModelVersionOutput() RegistryModelVersionOutput
	ToRegistryModelVersionOutputWithContext(ctx context.Context) RegistryModelVersionOutput
}

func (*RegistryModelVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryModelVersion)(nil)).Elem()
}

func (i *RegistryModelVersion) ToRegistryModelVersionOutput() RegistryModelVersionOutput {
	return i.ToRegistryModelVersionOutputWithContext(context.Background())
}

func (i *RegistryModelVersion) ToRegistryModelVersionOutputWithContext(ctx context.Context) RegistryModelVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryModelVersionOutput)
}

type RegistryModelVersionOutput struct{ *pulumi.OutputState }

func (RegistryModelVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistryModelVersion)(nil)).Elem()
}

func (o RegistryModelVersionOutput) ToRegistryModelVersionOutput() RegistryModelVersionOutput {
	return o
}

func (o RegistryModelVersionOutput) ToRegistryModelVersionOutputWithContext(ctx context.Context) RegistryModelVersionOutput {
	return o
}

// The Azure API version of the resource.
func (o RegistryModelVersionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryModelVersion) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o RegistryModelVersionOutput) ModelVersionProperties() ModelVersionResponseOutput {
	return o.ApplyT(func(v *RegistryModelVersion) ModelVersionResponseOutput { return v.ModelVersionProperties }).(ModelVersionResponseOutput)
}

// The name of the resource
func (o RegistryModelVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryModelVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o RegistryModelVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *RegistryModelVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o RegistryModelVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegistryModelVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegistryModelVersionOutput{})
}
