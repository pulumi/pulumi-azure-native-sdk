// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package avs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NSX Port Mirroring
//
// Uses Azure REST API version 2023-09-01. In version 2.x of the Azure Native provider, it used API version 2022-05-01.
//
// Other available API versions: 2022-05-01, 2023-03-01, 2024-09-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native avs [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type WorkloadNetworkPortMirroring struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Destination VM Group.
	Destination pulumi.StringPtrOutput `pulumi:"destination"`
	// Direction of port mirroring profile.
	Direction pulumi.StringPtrOutput `pulumi:"direction"`
	// Display name of the port mirroring profile.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// NSX revision number.
	Revision pulumi.Float64PtrOutput `pulumi:"revision"`
	// Source VM Group.
	Source pulumi.StringPtrOutput `pulumi:"source"`
	// Port Mirroring Status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadNetworkPortMirroring registers a new resource with the given unique name, arguments, and options.
func NewWorkloadNetworkPortMirroring(ctx *pulumi.Context,
	name string, args *WorkloadNetworkPortMirroringArgs, opts ...pulumi.ResourceOption) (*WorkloadNetworkPortMirroring, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs/v20200717preview:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210101preview:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20210601:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20220501:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230301:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20230901:WorkloadNetworkPortMirroring"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20240901:WorkloadNetworkPortMirroring"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource WorkloadNetworkPortMirroring
	err := ctx.RegisterResource("azure-native:avs:WorkloadNetworkPortMirroring", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadNetworkPortMirroring gets an existing WorkloadNetworkPortMirroring resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadNetworkPortMirroring(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadNetworkPortMirroringState, opts ...pulumi.ResourceOption) (*WorkloadNetworkPortMirroring, error) {
	var resource WorkloadNetworkPortMirroring
	err := ctx.ReadResource("azure-native:avs:WorkloadNetworkPortMirroring", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadNetworkPortMirroring resources.
type workloadNetworkPortMirroringState struct {
}

type WorkloadNetworkPortMirroringState struct {
}

func (WorkloadNetworkPortMirroringState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkPortMirroringState)(nil)).Elem()
}

type workloadNetworkPortMirroringArgs struct {
	// Destination VM Group.
	Destination *string `pulumi:"destination"`
	// Direction of port mirroring profile.
	Direction *string `pulumi:"direction"`
	// Display name of the port mirroring profile.
	DisplayName *string `pulumi:"displayName"`
	// ID of the NSX port mirroring profile.
	PortMirroringId *string `pulumi:"portMirroringId"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// NSX revision number.
	Revision *float64 `pulumi:"revision"`
	// Source VM Group.
	Source *string `pulumi:"source"`
}

// The set of arguments for constructing a WorkloadNetworkPortMirroring resource.
type WorkloadNetworkPortMirroringArgs struct {
	// Destination VM Group.
	Destination pulumi.StringPtrInput
	// Direction of port mirroring profile.
	Direction pulumi.StringPtrInput
	// Display name of the port mirroring profile.
	DisplayName pulumi.StringPtrInput
	// ID of the NSX port mirroring profile.
	PortMirroringId pulumi.StringPtrInput
	// Name of the private cloud
	PrivateCloudName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// NSX revision number.
	Revision pulumi.Float64PtrInput
	// Source VM Group.
	Source pulumi.StringPtrInput
}

func (WorkloadNetworkPortMirroringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadNetworkPortMirroringArgs)(nil)).Elem()
}

type WorkloadNetworkPortMirroringInput interface {
	pulumi.Input

	ToWorkloadNetworkPortMirroringOutput() WorkloadNetworkPortMirroringOutput
	ToWorkloadNetworkPortMirroringOutputWithContext(ctx context.Context) WorkloadNetworkPortMirroringOutput
}

func (*WorkloadNetworkPortMirroring) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkPortMirroring)(nil)).Elem()
}

func (i *WorkloadNetworkPortMirroring) ToWorkloadNetworkPortMirroringOutput() WorkloadNetworkPortMirroringOutput {
	return i.ToWorkloadNetworkPortMirroringOutputWithContext(context.Background())
}

func (i *WorkloadNetworkPortMirroring) ToWorkloadNetworkPortMirroringOutputWithContext(ctx context.Context) WorkloadNetworkPortMirroringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadNetworkPortMirroringOutput)
}

type WorkloadNetworkPortMirroringOutput struct{ *pulumi.OutputState }

func (WorkloadNetworkPortMirroringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadNetworkPortMirroring)(nil)).Elem()
}

func (o WorkloadNetworkPortMirroringOutput) ToWorkloadNetworkPortMirroringOutput() WorkloadNetworkPortMirroringOutput {
	return o
}

func (o WorkloadNetworkPortMirroringOutput) ToWorkloadNetworkPortMirroringOutputWithContext(ctx context.Context) WorkloadNetworkPortMirroringOutput {
	return o
}

// The Azure API version of the resource.
func (o WorkloadNetworkPortMirroringOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Destination VM Group.
func (o WorkloadNetworkPortMirroringOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.Destination }).(pulumi.StringPtrOutput)
}

// Direction of port mirroring profile.
func (o WorkloadNetworkPortMirroringOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.Direction }).(pulumi.StringPtrOutput)
}

// Display name of the port mirroring profile.
func (o WorkloadNetworkPortMirroringOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The name of the resource
func (o WorkloadNetworkPortMirroringOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state
func (o WorkloadNetworkPortMirroringOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// NSX revision number.
func (o WorkloadNetworkPortMirroringOutput) Revision() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.Float64PtrOutput { return v.Revision }).(pulumi.Float64PtrOutput)
}

// Source VM Group.
func (o WorkloadNetworkPortMirroringOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

// Port Mirroring Status.
func (o WorkloadNetworkPortMirroringOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WorkloadNetworkPortMirroringOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkloadNetworkPortMirroringOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkloadNetworkPortMirroring) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkloadNetworkPortMirroringOutput{})
}
