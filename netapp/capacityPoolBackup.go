// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Backup of a Volume
//
// Uses Azure REST API version 2022-11-01.
type CapacityPoolBackup struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// UUID v4 used to identify the Backup
	BackupId pulumi.StringOutput `pulumi:"backupId"`
	// Type of backup Manual or Scheduled
	BackupType pulumi.StringOutput `pulumi:"backupType"`
	// The creation date of the backup
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// Failure reason
	FailureReason pulumi.StringOutput `pulumi:"failureReason"`
	// Label for backup
	Label pulumi.StringPtrOutput `pulumi:"label"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Size of backup
	Size pulumi.Float64Output `pulumi:"size"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Manual backup an already existing snapshot. This will always be false for scheduled backups and true/false for manual backups
	UseExistingSnapshot pulumi.BoolPtrOutput `pulumi:"useExistingSnapshot"`
	// Volume name
	VolumeName pulumi.StringOutput `pulumi:"volumeName"`
}

// NewCapacityPoolBackup registers a new resource with the given unique name, arguments, and options.
func NewCapacityPoolBackup(ctx *pulumi.Context,
	name string, args *CapacityPoolBackupArgs, opts ...pulumi.ResourceOption) (*CapacityPoolBackup, error) {
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
	if args.UseExistingSnapshot == nil {
		args.UseExistingSnapshot = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp/v20200501:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220901:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101:CapacityPoolBackup"),
		},
		{
			Type: pulumi.String("azure-native:netapp:Backup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource CapacityPoolBackup
	err := ctx.RegisterResource("azure-native:netapp:CapacityPoolBackup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCapacityPoolBackup gets an existing CapacityPoolBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCapacityPoolBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityPoolBackupState, opts ...pulumi.ResourceOption) (*CapacityPoolBackup, error) {
	var resource CapacityPoolBackup
	err := ctx.ReadResource("azure-native:netapp:CapacityPoolBackup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CapacityPoolBackup resources.
type capacityPoolBackupState struct {
}

type CapacityPoolBackupState struct {
}

func (CapacityPoolBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityPoolBackupState)(nil)).Elem()
}

type capacityPoolBackupArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// The name of the backup
	BackupName *string `pulumi:"backupName"`
	// Label for backup
	Label *string `pulumi:"label"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the capacity pool
	PoolName string `pulumi:"poolName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Manual backup an already existing snapshot. This will always be false for scheduled backups and true/false for manual backups
	UseExistingSnapshot *bool `pulumi:"useExistingSnapshot"`
	// The name of the volume
	VolumeName string `pulumi:"volumeName"`
}

// The set of arguments for constructing a CapacityPoolBackup resource.
type CapacityPoolBackupArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// The name of the backup
	BackupName pulumi.StringPtrInput
	// Label for backup
	Label pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the capacity pool
	PoolName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Manual backup an already existing snapshot. This will always be false for scheduled backups and true/false for manual backups
	UseExistingSnapshot pulumi.BoolPtrInput
	// The name of the volume
	VolumeName pulumi.StringInput
}

func (CapacityPoolBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityPoolBackupArgs)(nil)).Elem()
}

type CapacityPoolBackupInput interface {
	pulumi.Input

	ToCapacityPoolBackupOutput() CapacityPoolBackupOutput
	ToCapacityPoolBackupOutputWithContext(ctx context.Context) CapacityPoolBackupOutput
}

func (*CapacityPoolBackup) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityPoolBackup)(nil)).Elem()
}

func (i *CapacityPoolBackup) ToCapacityPoolBackupOutput() CapacityPoolBackupOutput {
	return i.ToCapacityPoolBackupOutputWithContext(context.Background())
}

func (i *CapacityPoolBackup) ToCapacityPoolBackupOutputWithContext(ctx context.Context) CapacityPoolBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityPoolBackupOutput)
}

type CapacityPoolBackupOutput struct{ *pulumi.OutputState }

func (CapacityPoolBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityPoolBackup)(nil)).Elem()
}

func (o CapacityPoolBackupOutput) ToCapacityPoolBackupOutput() CapacityPoolBackupOutput {
	return o
}

func (o CapacityPoolBackupOutput) ToCapacityPoolBackupOutputWithContext(ctx context.Context) CapacityPoolBackupOutput {
	return o
}

// The Azure API version of the resource.
func (o CapacityPoolBackupOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// UUID v4 used to identify the Backup
func (o CapacityPoolBackupOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

// Type of backup Manual or Scheduled
func (o CapacityPoolBackupOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.StringOutput { return v.BackupType }).(pulumi.StringOutput)
}

// The creation date of the backup
func (o CapacityPoolBackupOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Failure reason
func (o CapacityPoolBackupOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.StringOutput { return v.FailureReason }).(pulumi.StringOutput)
}

// Label for backup
func (o CapacityPoolBackupOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

// Resource location
func (o CapacityPoolBackupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o CapacityPoolBackupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o CapacityPoolBackupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Size of backup
func (o CapacityPoolBackupOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.Float64Output { return v.Size }).(pulumi.Float64Output)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o CapacityPoolBackupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o CapacityPoolBackupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Manual backup an already existing snapshot. This will always be false for scheduled backups and true/false for manual backups
func (o CapacityPoolBackupOutput) UseExistingSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.BoolPtrOutput { return v.UseExistingSnapshot }).(pulumi.BoolPtrOutput)
}

// Volume name
func (o CapacityPoolBackupOutput) VolumeName() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityPoolBackup) pulumi.StringOutput { return v.VolumeName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CapacityPoolBackupOutput{})
}
