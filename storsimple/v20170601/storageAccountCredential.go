// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The storage account credential.
type StorageAccountCredential struct {
	pulumi.CustomResourceState

	// The details of the storage account password.
	AccessKey AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"accessKey"`
	// The storage endpoint
	EndPoint pulumi.StringOutput `pulumi:"endPoint"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// Signifies whether SSL needs to be enabled or not.
	SslStatus pulumi.StringOutput `pulumi:"sslStatus"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// The count of volumes using this storage account credential.
	VolumesCount pulumi.IntOutput `pulumi:"volumesCount"`
}

// NewStorageAccountCredential registers a new resource with the given unique name, arguments, and options.
func NewStorageAccountCredential(ctx *pulumi.Context,
	name string, args *StorageAccountCredentialArgs, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndPoint == nil {
		return nil, errors.New("invalid value for required argument 'EndPoint'")
	}
	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SslStatus == nil {
		return nil, errors.New("invalid value for required argument 'SslStatus'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storsimple:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:storsimple/v20161001:StorageAccountCredential"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccountCredential
	err := ctx.RegisterResource("azure-native:storsimple/v20170601:StorageAccountCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageAccountCredential gets an existing StorageAccountCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageAccountCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountCredentialState, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	var resource StorageAccountCredential
	err := ctx.ReadResource("azure-native:storsimple/v20170601:StorageAccountCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageAccountCredential resources.
type storageAccountCredentialState struct {
}

type StorageAccountCredentialState struct {
}

func (StorageAccountCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountCredentialState)(nil)).Elem()
}

type storageAccountCredentialArgs struct {
	// The details of the storage account password.
	AccessKey *AsymmetricEncryptedSecret `pulumi:"accessKey"`
	// The storage endpoint
	EndPoint string `pulumi:"endPoint"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *Kind `pulumi:"kind"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Signifies whether SSL needs to be enabled or not.
	SslStatus SslStatus `pulumi:"sslStatus"`
	// The storage account credential name.
	StorageAccountCredentialName *string `pulumi:"storageAccountCredentialName"`
}

// The set of arguments for constructing a StorageAccountCredential resource.
type StorageAccountCredentialArgs struct {
	// The details of the storage account password.
	AccessKey AsymmetricEncryptedSecretPtrInput
	// The storage endpoint
	EndPoint pulumi.StringInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind KindPtrInput
	// The manager name
	ManagerName pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// Signifies whether SSL needs to be enabled or not.
	SslStatus SslStatusInput
	// The storage account credential name.
	StorageAccountCredentialName pulumi.StringPtrInput
}

func (StorageAccountCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountCredentialArgs)(nil)).Elem()
}

type StorageAccountCredentialInput interface {
	pulumi.Input

	ToStorageAccountCredentialOutput() StorageAccountCredentialOutput
	ToStorageAccountCredentialOutputWithContext(ctx context.Context) StorageAccountCredentialOutput
}

func (*StorageAccountCredential) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountCredential)(nil)).Elem()
}

func (i *StorageAccountCredential) ToStorageAccountCredentialOutput() StorageAccountCredentialOutput {
	return i.ToStorageAccountCredentialOutputWithContext(context.Background())
}

func (i *StorageAccountCredential) ToStorageAccountCredentialOutputWithContext(ctx context.Context) StorageAccountCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountCredentialOutput)
}

type StorageAccountCredentialOutput struct{ *pulumi.OutputState }

func (StorageAccountCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountCredential)(nil)).Elem()
}

func (o StorageAccountCredentialOutput) ToStorageAccountCredentialOutput() StorageAccountCredentialOutput {
	return o
}

func (o StorageAccountCredentialOutput) ToStorageAccountCredentialOutputWithContext(ctx context.Context) StorageAccountCredentialOutput {
	return o
}

// The details of the storage account password.
func (o StorageAccountCredentialOutput) AccessKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) AsymmetricEncryptedSecretResponsePtrOutput { return v.AccessKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

// The storage endpoint
func (o StorageAccountCredentialOutput) EndPoint() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.EndPoint }).(pulumi.StringOutput)
}

// The Kind of the object. Currently only Series8000 is supported
func (o StorageAccountCredentialOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

// The name of the object.
func (o StorageAccountCredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Signifies whether SSL needs to be enabled or not.
func (o StorageAccountCredentialOutput) SslStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.SslStatus }).(pulumi.StringOutput)
}

// The hierarchical type of the object.
func (o StorageAccountCredentialOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The count of volumes using this storage account credential.
func (o StorageAccountCredentialOutput) VolumesCount() pulumi.IntOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.IntOutput { return v.VolumesCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageAccountCredentialOutput{})
}