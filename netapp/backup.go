// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package netapp

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Backup of a Volume
//
// Uses Azure REST API version 2022-11-01. In version 1.x of the Azure Native provider, it used API version 2020-12-01.
//
// Other available API versions: 2022-11-01-preview, 2023-05-01-preview, 2023-07-01-preview, 2023-11-01, 2023-11-01-preview, 2024-01-01, 2024-03-01, 2024-03-01-preview, 2024-05-01, 2024-05-01-preview, 2024-07-01, 2024-07-01-preview, 2024-09-01, 2024-09-01-preview.
type Backup struct {
	pulumi.CustomResourceState

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

// NewBackup registers a new resource with the given unique name, arguments, and options.
func NewBackup(ctx *pulumi.Context,
	name string, args *BackupArgs, opts ...pulumi.ResourceOption) (*Backup, error) {
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
			Type: pulumi.String("azure-native:netapp/v20200501:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220901:Backup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101:Backup"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Backup
	err := ctx.RegisterResource("azure-native:netapp:Backup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackup gets an existing Backup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupState, opts ...pulumi.ResourceOption) (*Backup, error) {
	var resource Backup
	err := ctx.ReadResource("azure-native:netapp:Backup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Backup resources.
type backupState struct {
}

type BackupState struct {
}

func (BackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupState)(nil)).Elem()
}

type backupArgs struct {
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

// The set of arguments for constructing a Backup resource.
type BackupArgs struct {
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

func (BackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupArgs)(nil)).Elem()
}

type BackupInput interface {
	pulumi.Input

	ToBackupOutput() BackupOutput
	ToBackupOutputWithContext(ctx context.Context) BackupOutput
}

func (*Backup) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (i *Backup) ToBackupOutput() BackupOutput {
	return i.ToBackupOutputWithContext(context.Background())
}

func (i *Backup) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupOutput)
}

type BackupOutput struct{ *pulumi.OutputState }

func (BackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Backup)(nil)).Elem()
}

func (o BackupOutput) ToBackupOutput() BackupOutput {
	return o
}

func (o BackupOutput) ToBackupOutputWithContext(ctx context.Context) BackupOutput {
	return o
}

// UUID v4 used to identify the Backup
func (o BackupOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

// Type of backup Manual or Scheduled
func (o BackupOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.BackupType }).(pulumi.StringOutput)
}

// The creation date of the backup
func (o BackupOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

// Failure reason
func (o BackupOutput) FailureReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.FailureReason }).(pulumi.StringOutput)
}

// Label for backup
func (o BackupOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringPtrOutput { return v.Label }).(pulumi.StringPtrOutput)
}

// Resource location
func (o BackupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o BackupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o BackupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Size of backup
func (o BackupOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v *Backup) pulumi.Float64Output { return v.Size }).(pulumi.Float64Output)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o BackupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Backup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BackupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Manual backup an already existing snapshot. This will always be false for scheduled backups and true/false for manual backups
func (o BackupOutput) UseExistingSnapshot() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Backup) pulumi.BoolPtrOutput { return v.UseExistingSnapshot }).(pulumi.BoolPtrOutput)
}

// Volume name
func (o BackupOutput) VolumeName() pulumi.StringOutput {
	return o.ApplyT(func(v *Backup) pulumi.StringOutput { return v.VolumeName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BackupOutput{})
}
