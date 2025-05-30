// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workloads

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A provider instance associated with SAP monitor.
//
// Uses Azure REST API version 2024-02-01-preview. In version 2.x of the Azure Native provider, it used API version 2023-04-01.
//
// Other available API versions: 2023-04-01, 2023-10-01-preview, 2023-12-01-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native workloads [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type ProviderInstance struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Defines the provider instance errors.
	Errors ErrorDetailResponseOutput `pulumi:"errors"`
	// Resource health details
	Health HealthResponseOutput `pulumi:"health"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Defines the provider specific properties.
	ProviderSettings pulumi.AnyOutput `pulumi:"providerSettings"`
	// State of provisioning of the provider instance
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProviderInstance registers a new resource with the given unique name, arguments, and options.
func NewProviderInstance(ctx *pulumi.Context,
	name string, args *ProviderInstanceArgs, opts ...pulumi.ResourceOption) (*ProviderInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:workloads/v20211201preview:ProviderInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20221101preview:ProviderInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20230401:ProviderInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20231001preview:ProviderInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20231201preview:ProviderInstance"),
		},
		{
			Type: pulumi.String("azure-native:workloads/v20240201preview:ProviderInstance"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource ProviderInstance
	err := ctx.RegisterResource("azure-native:workloads:ProviderInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderInstance gets an existing ProviderInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderInstanceState, opts ...pulumi.ResourceOption) (*ProviderInstance, error) {
	var resource ProviderInstance
	err := ctx.ReadResource("azure-native:workloads:ProviderInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderInstance resources.
type providerInstanceState struct {
}

type ProviderInstanceState struct {
}

func (ProviderInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerInstanceState)(nil)).Elem()
}

type providerInstanceArgs struct {
	// Name of the SAP monitor resource.
	MonitorName string `pulumi:"monitorName"`
	// Name of the provider instance.
	ProviderInstanceName *string `pulumi:"providerInstanceName"`
	// Defines the provider specific properties.
	ProviderSettings interface{} `pulumi:"providerSettings"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ProviderInstance resource.
type ProviderInstanceArgs struct {
	// Name of the SAP monitor resource.
	MonitorName pulumi.StringInput
	// Name of the provider instance.
	ProviderInstanceName pulumi.StringPtrInput
	// Defines the provider specific properties.
	ProviderSettings pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ProviderInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerInstanceArgs)(nil)).Elem()
}

type ProviderInstanceInput interface {
	pulumi.Input

	ToProviderInstanceOutput() ProviderInstanceOutput
	ToProviderInstanceOutputWithContext(ctx context.Context) ProviderInstanceOutput
}

func (*ProviderInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderInstance)(nil)).Elem()
}

func (i *ProviderInstance) ToProviderInstanceOutput() ProviderInstanceOutput {
	return i.ToProviderInstanceOutputWithContext(context.Background())
}

func (i *ProviderInstance) ToProviderInstanceOutputWithContext(ctx context.Context) ProviderInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderInstanceOutput)
}

type ProviderInstanceOutput struct{ *pulumi.OutputState }

func (ProviderInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderInstance)(nil)).Elem()
}

func (o ProviderInstanceOutput) ToProviderInstanceOutput() ProviderInstanceOutput {
	return o
}

func (o ProviderInstanceOutput) ToProviderInstanceOutputWithContext(ctx context.Context) ProviderInstanceOutput {
	return o
}

// The Azure API version of the resource.
func (o ProviderInstanceOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderInstance) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Defines the provider instance errors.
func (o ProviderInstanceOutput) Errors() ErrorDetailResponseOutput {
	return o.ApplyT(func(v *ProviderInstance) ErrorDetailResponseOutput { return v.Errors }).(ErrorDetailResponseOutput)
}

// Resource health details
func (o ProviderInstanceOutput) Health() HealthResponseOutput {
	return o.ApplyT(func(v *ProviderInstance) HealthResponseOutput { return v.Health }).(HealthResponseOutput)
}

// The name of the resource
func (o ProviderInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Defines the provider specific properties.
func (o ProviderInstanceOutput) ProviderSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *ProviderInstance) pulumi.AnyOutput { return v.ProviderSettings }).(pulumi.AnyOutput)
}

// State of provisioning of the provider instance
func (o ProviderInstanceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderInstance) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ProviderInstanceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ProviderInstance) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ProviderInstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderInstance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ProviderInstanceOutput{})
}
