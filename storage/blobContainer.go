// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Properties of the blob container, including Id, resource name, resource type, Etag.
//
// Uses Azure REST API version 2024-01-01. In version 2.x of the Azure Native provider, it used API version 2022-09-01.
//
// Other available API versions: 2022-09-01, 2023-01-01, 2023-04-01, 2023-05-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storage [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type BlobContainer struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Default the container to use specified encryption scope for all writes.
	DefaultEncryptionScope pulumi.StringPtrOutput `pulumi:"defaultEncryptionScope"`
	// Indicates whether the blob container was deleted.
	Deleted pulumi.BoolOutput `pulumi:"deleted"`
	// Blob container deletion time.
	DeletedTime pulumi.StringOutput `pulumi:"deletedTime"`
	// Block override of encryption scope from the container default.
	DenyEncryptionScopeOverride pulumi.BoolPtrOutput `pulumi:"denyEncryptionScopeOverride"`
	// Enable NFSv3 all squash on blob container.
	EnableNfsV3AllSquash pulumi.BoolPtrOutput `pulumi:"enableNfsV3AllSquash"`
	// Enable NFSv3 root squash on blob container.
	EnableNfsV3RootSquash pulumi.BoolPtrOutput `pulumi:"enableNfsV3RootSquash"`
	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container. The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container.
	HasImmutabilityPolicy pulumi.BoolOutput `pulumi:"hasImmutabilityPolicy"`
	// The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.
	HasLegalHold pulumi.BoolOutput `pulumi:"hasLegalHold"`
	// The ImmutabilityPolicy property of the container.
	ImmutabilityPolicy ImmutabilityPolicyPropertiesResponseOutput `pulumi:"immutabilityPolicy"`
	// The object level immutability property of the container. The property is immutable and can only be set to true at the container creation time. Existing containers must undergo a migration process.
	ImmutableStorageWithVersioning ImmutableStorageWithVersioningResponsePtrOutput `pulumi:"immutableStorageWithVersioning"`
	// Returns the date and time the container was last modified.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased.
	LeaseDuration pulumi.StringOutput `pulumi:"leaseDuration"`
	// Lease state of the container.
	LeaseState pulumi.StringOutput `pulumi:"leaseState"`
	// The lease status of the container.
	LeaseStatus pulumi.StringOutput `pulumi:"leaseStatus"`
	// The LegalHold property of the container.
	LegalHold LegalHoldPropertiesResponseOutput `pulumi:"legalHold"`
	// A name-value pair to associate with the container as metadata.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies whether data in the container may be accessed publicly and the level of access.
	PublicAccess pulumi.StringPtrOutput `pulumi:"publicAccess"`
	// Remaining retention days for soft deleted blob container.
	RemainingRetentionDays pulumi.IntOutput `pulumi:"remainingRetentionDays"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The version of the deleted blob container.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewBlobContainer registers a new resource with the given unique name, arguments, and options.
func NewBlobContainer(ctx *pulumi.Context,
	name string, args *BlobContainerArgs, opts ...pulumi.ResourceOption) (*BlobContainer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storage/v20180201:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180301preview:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20180701:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20181101:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190401:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220901:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230101:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230401:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20230501:BlobContainer"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20240101:BlobContainer"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource BlobContainer
	err := ctx.RegisterResource("azure-native:storage:BlobContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlobContainer gets an existing BlobContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlobContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobContainerState, opts ...pulumi.ResourceOption) (*BlobContainer, error) {
	var resource BlobContainer
	err := ctx.ReadResource("azure-native:storage:BlobContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlobContainer resources.
type blobContainerState struct {
}

type BlobContainerState struct {
}

func (BlobContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerState)(nil)).Elem()
}

type blobContainerArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ContainerName *string `pulumi:"containerName"`
	// Default the container to use specified encryption scope for all writes.
	DefaultEncryptionScope *string `pulumi:"defaultEncryptionScope"`
	// Block override of encryption scope from the container default.
	DenyEncryptionScopeOverride *bool `pulumi:"denyEncryptionScopeOverride"`
	// Enable NFSv3 all squash on blob container.
	EnableNfsV3AllSquash *bool `pulumi:"enableNfsV3AllSquash"`
	// Enable NFSv3 root squash on blob container.
	EnableNfsV3RootSquash *bool `pulumi:"enableNfsV3RootSquash"`
	// The object level immutability property of the container. The property is immutable and can only be set to true at the container creation time. Existing containers must undergo a migration process.
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning `pulumi:"immutableStorageWithVersioning"`
	// A name-value pair to associate with the container as metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// Specifies whether data in the container may be accessed publicly and the level of access.
	PublicAccess *PublicAccess `pulumi:"publicAccess"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a BlobContainer resource.
type BlobContainerArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ContainerName pulumi.StringPtrInput
	// Default the container to use specified encryption scope for all writes.
	DefaultEncryptionScope pulumi.StringPtrInput
	// Block override of encryption scope from the container default.
	DenyEncryptionScopeOverride pulumi.BoolPtrInput
	// Enable NFSv3 all squash on blob container.
	EnableNfsV3AllSquash pulumi.BoolPtrInput
	// Enable NFSv3 root squash on blob container.
	EnableNfsV3RootSquash pulumi.BoolPtrInput
	// The object level immutability property of the container. The property is immutable and can only be set to true at the container creation time. Existing containers must undergo a migration process.
	ImmutableStorageWithVersioning ImmutableStorageWithVersioningPtrInput
	// A name-value pair to associate with the container as metadata.
	Metadata pulumi.StringMapInput
	// Specifies whether data in the container may be accessed publicly and the level of access.
	PublicAccess PublicAccessPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (BlobContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerArgs)(nil)).Elem()
}

type BlobContainerInput interface {
	pulumi.Input

	ToBlobContainerOutput() BlobContainerOutput
	ToBlobContainerOutputWithContext(ctx context.Context) BlobContainerOutput
}

func (*BlobContainer) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainer)(nil)).Elem()
}

func (i *BlobContainer) ToBlobContainerOutput() BlobContainerOutput {
	return i.ToBlobContainerOutputWithContext(context.Background())
}

func (i *BlobContainer) ToBlobContainerOutputWithContext(ctx context.Context) BlobContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobContainerOutput)
}

type BlobContainerOutput struct{ *pulumi.OutputState }

func (BlobContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobContainer)(nil)).Elem()
}

func (o BlobContainerOutput) ToBlobContainerOutput() BlobContainerOutput {
	return o
}

func (o BlobContainerOutput) ToBlobContainerOutputWithContext(ctx context.Context) BlobContainerOutput {
	return o
}

// The Azure API version of the resource.
func (o BlobContainerOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Default the container to use specified encryption scope for all writes.
func (o BlobContainerOutput) DefaultEncryptionScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringPtrOutput { return v.DefaultEncryptionScope }).(pulumi.StringPtrOutput)
}

// Indicates whether the blob container was deleted.
func (o BlobContainerOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.BoolOutput { return v.Deleted }).(pulumi.BoolOutput)
}

// Blob container deletion time.
func (o BlobContainerOutput) DeletedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.DeletedTime }).(pulumi.StringOutput)
}

// Block override of encryption scope from the container default.
func (o BlobContainerOutput) DenyEncryptionScopeOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.BoolPtrOutput { return v.DenyEncryptionScopeOverride }).(pulumi.BoolPtrOutput)
}

// Enable NFSv3 all squash on blob container.
func (o BlobContainerOutput) EnableNfsV3AllSquash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.BoolPtrOutput { return v.EnableNfsV3AllSquash }).(pulumi.BoolPtrOutput)
}

// Enable NFSv3 root squash on blob container.
func (o BlobContainerOutput) EnableNfsV3RootSquash() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.BoolPtrOutput { return v.EnableNfsV3RootSquash }).(pulumi.BoolPtrOutput)
}

// Resource Etag.
func (o BlobContainerOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The hasImmutabilityPolicy public property is set to true by SRP if ImmutabilityPolicy has been created for this container. The hasImmutabilityPolicy public property is set to false by SRP if ImmutabilityPolicy has not been created for this container.
func (o BlobContainerOutput) HasImmutabilityPolicy() pulumi.BoolOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.BoolOutput { return v.HasImmutabilityPolicy }).(pulumi.BoolOutput)
}

// The hasLegalHold public property is set to true by SRP if there are at least one existing tag. The hasLegalHold public property is set to false by SRP if all existing legal hold tags are cleared out. There can be a maximum of 1000 blob containers with hasLegalHold=true for a given account.
func (o BlobContainerOutput) HasLegalHold() pulumi.BoolOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.BoolOutput { return v.HasLegalHold }).(pulumi.BoolOutput)
}

// The ImmutabilityPolicy property of the container.
func (o BlobContainerOutput) ImmutabilityPolicy() ImmutabilityPolicyPropertiesResponseOutput {
	return o.ApplyT(func(v *BlobContainer) ImmutabilityPolicyPropertiesResponseOutput { return v.ImmutabilityPolicy }).(ImmutabilityPolicyPropertiesResponseOutput)
}

// The object level immutability property of the container. The property is immutable and can only be set to true at the container creation time. Existing containers must undergo a migration process.
func (o BlobContainerOutput) ImmutableStorageWithVersioning() ImmutableStorageWithVersioningResponsePtrOutput {
	return o.ApplyT(func(v *BlobContainer) ImmutableStorageWithVersioningResponsePtrOutput {
		return v.ImmutableStorageWithVersioning
	}).(ImmutableStorageWithVersioningResponsePtrOutput)
}

// Returns the date and time the container was last modified.
func (o BlobContainerOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// Specifies whether the lease on a container is of infinite or fixed duration, only when the container is leased.
func (o BlobContainerOutput) LeaseDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.LeaseDuration }).(pulumi.StringOutput)
}

// Lease state of the container.
func (o BlobContainerOutput) LeaseState() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.LeaseState }).(pulumi.StringOutput)
}

// The lease status of the container.
func (o BlobContainerOutput) LeaseStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.LeaseStatus }).(pulumi.StringOutput)
}

// The LegalHold property of the container.
func (o BlobContainerOutput) LegalHold() LegalHoldPropertiesResponseOutput {
	return o.ApplyT(func(v *BlobContainer) LegalHoldPropertiesResponseOutput { return v.LegalHold }).(LegalHoldPropertiesResponseOutput)
}

// A name-value pair to associate with the container as metadata.
func (o BlobContainerOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the resource
func (o BlobContainerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies whether data in the container may be accessed publicly and the level of access.
func (o BlobContainerOutput) PublicAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringPtrOutput { return v.PublicAccess }).(pulumi.StringPtrOutput)
}

// Remaining retention days for soft deleted blob container.
func (o BlobContainerOutput) RemainingRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.IntOutput { return v.RemainingRetentionDays }).(pulumi.IntOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o BlobContainerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The version of the deleted blob container.
func (o BlobContainerOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *BlobContainer) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobContainerOutput{})
}
