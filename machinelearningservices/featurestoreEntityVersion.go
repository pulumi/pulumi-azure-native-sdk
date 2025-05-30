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
// Uses Azure REST API version 2024-10-01. In version 2.x of the Azure Native provider, it used API version 2023-04-01-preview.
//
// Other available API versions: 2023-02-01-preview, 2023-04-01-preview, 2023-06-01-preview, 2023-08-01-preview, 2023-10-01, 2024-01-01-preview, 2024-04-01, 2024-07-01-preview, 2024-10-01-preview, 2025-01-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native machinelearningservices [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type FeaturestoreEntityVersion struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// [Required] Additional attributes of the entity.
	FeaturestoreEntityVersionProperties FeaturestoreEntityVersionResponseOutput `pulumi:"featurestoreEntityVersionProperties"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFeaturestoreEntityVersion registers a new resource with the given unique name, arguments, and options.
func NewFeaturestoreEntityVersion(ctx *pulumi.Context,
	name string, args *FeaturestoreEntityVersionArgs, opts ...pulumi.ResourceOption) (*FeaturestoreEntityVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FeaturestoreEntityVersionProperties == nil {
		return nil, errors.New("invalid value for required argument 'FeaturestoreEntityVersionProperties'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.FeaturestoreEntityVersionProperties = args.FeaturestoreEntityVersionProperties.ToFeaturestoreEntityVersionTypeOutput().ApplyT(func(v FeaturestoreEntityVersionType) FeaturestoreEntityVersionType { return *v.Defaults() }).(FeaturestoreEntityVersionTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:FeaturestoreEntityVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:FeaturestoreEntityVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:FeaturestoreEntityVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:FeaturestoreEntityVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20231001:FeaturestoreEntityVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20240101preview:FeaturestoreEntityVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20240401:FeaturestoreEntityVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20240401preview:FeaturestoreEntityVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20240701preview:FeaturestoreEntityVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20241001:FeaturestoreEntityVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20241001preview:FeaturestoreEntityVersion"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20250101preview:FeaturestoreEntityVersion"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FeaturestoreEntityVersion
	err := ctx.RegisterResource("azure-native:machinelearningservices:FeaturestoreEntityVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeaturestoreEntityVersion gets an existing FeaturestoreEntityVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeaturestoreEntityVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeaturestoreEntityVersionState, opts ...pulumi.ResourceOption) (*FeaturestoreEntityVersion, error) {
	var resource FeaturestoreEntityVersion
	err := ctx.ReadResource("azure-native:machinelearningservices:FeaturestoreEntityVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeaturestoreEntityVersion resources.
type featurestoreEntityVersionState struct {
}

type FeaturestoreEntityVersionState struct {
}

func (FeaturestoreEntityVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*featurestoreEntityVersionState)(nil)).Elem()
}

type featurestoreEntityVersionArgs struct {
	// [Required] Additional attributes of the entity.
	FeaturestoreEntityVersionProperties FeaturestoreEntityVersionType `pulumi:"featurestoreEntityVersionProperties"`
	// Container name. This is case-sensitive.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Version identifier. This is case-sensitive.
	Version *string `pulumi:"version"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a FeaturestoreEntityVersion resource.
type FeaturestoreEntityVersionArgs struct {
	// [Required] Additional attributes of the entity.
	FeaturestoreEntityVersionProperties FeaturestoreEntityVersionTypeInput
	// Container name. This is case-sensitive.
	Name pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Version identifier. This is case-sensitive.
	Version pulumi.StringPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (FeaturestoreEntityVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featurestoreEntityVersionArgs)(nil)).Elem()
}

type FeaturestoreEntityVersionInput interface {
	pulumi.Input

	ToFeaturestoreEntityVersionOutput() FeaturestoreEntityVersionOutput
	ToFeaturestoreEntityVersionOutputWithContext(ctx context.Context) FeaturestoreEntityVersionOutput
}

func (*FeaturestoreEntityVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturestoreEntityVersion)(nil)).Elem()
}

func (i *FeaturestoreEntityVersion) ToFeaturestoreEntityVersionOutput() FeaturestoreEntityVersionOutput {
	return i.ToFeaturestoreEntityVersionOutputWithContext(context.Background())
}

func (i *FeaturestoreEntityVersion) ToFeaturestoreEntityVersionOutputWithContext(ctx context.Context) FeaturestoreEntityVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeaturestoreEntityVersionOutput)
}

type FeaturestoreEntityVersionOutput struct{ *pulumi.OutputState }

func (FeaturestoreEntityVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeaturestoreEntityVersion)(nil)).Elem()
}

func (o FeaturestoreEntityVersionOutput) ToFeaturestoreEntityVersionOutput() FeaturestoreEntityVersionOutput {
	return o
}

func (o FeaturestoreEntityVersionOutput) ToFeaturestoreEntityVersionOutputWithContext(ctx context.Context) FeaturestoreEntityVersionOutput {
	return o
}

// The Azure API version of the resource.
func (o FeaturestoreEntityVersionOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FeaturestoreEntityVersion) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// [Required] Additional attributes of the entity.
func (o FeaturestoreEntityVersionOutput) FeaturestoreEntityVersionProperties() FeaturestoreEntityVersionResponseOutput {
	return o.ApplyT(func(v *FeaturestoreEntityVersion) FeaturestoreEntityVersionResponseOutput {
		return v.FeaturestoreEntityVersionProperties
	}).(FeaturestoreEntityVersionResponseOutput)
}

// The name of the resource
func (o FeaturestoreEntityVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FeaturestoreEntityVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FeaturestoreEntityVersionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FeaturestoreEntityVersion) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FeaturestoreEntityVersionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FeaturestoreEntityVersion) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FeaturestoreEntityVersionOutput{})
}
