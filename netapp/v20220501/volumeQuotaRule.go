// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Quota Rule of a Volume
type VolumeQuotaRule struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets the status of the VolumeQuotaRule at the time the operation was called.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Size of quota
	QuotaSizeInKiBs pulumi.Float64PtrOutput `pulumi:"quotaSizeInKiBs"`
	// UserID/GroupID/SID based on the quota target type. UserID and groupID can be found by running ‘id’ or ‘getent’ command for the user or group and SID can be found by running <wmic useraccount where name='user-name' get sid>
	QuotaTarget pulumi.StringPtrOutput `pulumi:"quotaTarget"`
	// Type of quota
	QuotaType pulumi.StringPtrOutput `pulumi:"quotaType"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVolumeQuotaRule registers a new resource with the given unique name, arguments, and options.
func NewVolumeQuotaRule(ctx *pulumi.Context,
	name string, args *VolumeQuotaRuleArgs, opts ...pulumi.ResourceOption) (*VolumeQuotaRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VolumeName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:VolumeQuotaRule"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:VolumeQuotaRule"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:VolumeQuotaRule"),
		},
	})
	opts = append(opts, aliases)
	var resource VolumeQuotaRule
	err := ctx.RegisterResource("azure-native:netapp/v20220501:VolumeQuotaRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolumeQuotaRule gets an existing VolumeQuotaRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolumeQuotaRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeQuotaRuleState, opts ...pulumi.ResourceOption) (*VolumeQuotaRule, error) {
	var resource VolumeQuotaRule
	err := ctx.ReadResource("azure-native:netapp/v20220501:VolumeQuotaRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VolumeQuotaRule resources.
type volumeQuotaRuleState struct {
}

type VolumeQuotaRuleState struct {
}

func (VolumeQuotaRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeQuotaRuleState)(nil)).Elem()
}

type volumeQuotaRuleArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// Size of quota
	QuotaSizeInKiBs *float64 `pulumi:"quotaSizeInKiBs"`
	// UserID/GroupID/SID based on the quota target type. UserID and groupID can be found by running ‘id’ or ‘getent’ command for the user or group and SID can be found by running <wmic useraccount where name='user-name' get sid>
	QuotaTarget *string `pulumi:"quotaTarget"`
	// Type of quota
	QuotaType *string `pulumi:"quotaType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
	// The name of volume quota rule
	VolumeQuotaRuleName *string `pulumi:"volumeQuotaRuleName"`
}

// The set of arguments for constructing a VolumeQuotaRule resource.
type VolumeQuotaRuleArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the capacity pool
	PoolName pulumi.StringInput
	// Size of quota
	QuotaSizeInKiBs pulumi.Float64PtrInput
	// UserID/GroupID/SID based on the quota target type. UserID and groupID can be found by running ‘id’ or ‘getent’ command for the user or group and SID can be found by running <wmic useraccount where name='user-name' get sid>
	QuotaTarget pulumi.StringPtrInput
	// Type of quota
	QuotaType pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the volume
	VolumeName pulumi.StringInput
	// The name of volume quota rule
	VolumeQuotaRuleName pulumi.StringPtrInput
}

func (VolumeQuotaRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeQuotaRuleArgs)(nil)).Elem()
}

type VolumeQuotaRuleInput interface {
	pulumi.Input

	ToVolumeQuotaRuleOutput() VolumeQuotaRuleOutput
	ToVolumeQuotaRuleOutputWithContext(ctx context.Context) VolumeQuotaRuleOutput
}

func (*VolumeQuotaRule) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeQuotaRule)(nil)).Elem()
}

func (i *VolumeQuotaRule) ToVolumeQuotaRuleOutput() VolumeQuotaRuleOutput {
	return i.ToVolumeQuotaRuleOutputWithContext(context.Background())
}

func (i *VolumeQuotaRule) ToVolumeQuotaRuleOutputWithContext(ctx context.Context) VolumeQuotaRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeQuotaRuleOutput)
}

type VolumeQuotaRuleOutput struct{ *pulumi.OutputState }

func (VolumeQuotaRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeQuotaRule)(nil)).Elem()
}

func (o VolumeQuotaRuleOutput) ToVolumeQuotaRuleOutput() VolumeQuotaRuleOutput {
	return o
}

func (o VolumeQuotaRuleOutput) ToVolumeQuotaRuleOutputWithContext(ctx context.Context) VolumeQuotaRuleOutput {
	return o
}

// The geo-location where the resource lives
func (o VolumeQuotaRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o VolumeQuotaRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Gets the status of the VolumeQuotaRule at the time the operation was called.
func (o VolumeQuotaRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Size of quota
func (o VolumeQuotaRuleOutput) QuotaSizeInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.Float64PtrOutput { return v.QuotaSizeInKiBs }).(pulumi.Float64PtrOutput)
}

// UserID/GroupID/SID based on the quota target type. UserID and groupID can be found by running ‘id’ or ‘getent’ command for the user or group and SID can be found by running <wmic useraccount where name='user-name' get sid>
func (o VolumeQuotaRuleOutput) QuotaTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringPtrOutput { return v.QuotaTarget }).(pulumi.StringPtrOutput)
}

// Type of quota
func (o VolumeQuotaRuleOutput) QuotaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringPtrOutput { return v.QuotaType }).(pulumi.StringPtrOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o VolumeQuotaRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o VolumeQuotaRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o VolumeQuotaRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeQuotaRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeQuotaRuleOutput{})
}