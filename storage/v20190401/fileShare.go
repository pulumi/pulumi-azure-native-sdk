// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Properties of the file share, including Id, resource name, resource type, Etag.
//
// Deprecated: Version 2019-04-01 will be removed in v2 of the provider.
type FileShare struct {
	pulumi.CustomResourceState

	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Returns the date and time the share was last modified.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// A name-value pair to associate with the share as metadata.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120).
	ShareQuota pulumi.IntPtrOutput `pulumi:"shareQuota"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFileShare registers a new resource with the given unique name, arguments, and options.
func NewFileShare(ctx *pulumi.Context,
	name string, args *FileShareArgs, opts ...pulumi.ResourceOption) (*FileShare, error) {
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
			Type: pulumi.String("azure-native:storage:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20190601:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20200801preview:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210101:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210201:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210401:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210601:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210801:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20210901:FileShare"),
		},
		{
			Type: pulumi.String("azure-native:storage/v20220501:FileShare"),
		},
	})
	opts = append(opts, aliases)
	var resource FileShare
	err := ctx.RegisterResource("azure-native:storage/v20190401:FileShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileShare gets an existing FileShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileShareState, opts ...pulumi.ResourceOption) (*FileShare, error) {
	var resource FileShare
	err := ctx.ReadResource("azure-native:storage/v20190401:FileShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileShare resources.
type fileShareState struct {
}

type FileShareState struct {
}

func (FileShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareState)(nil)).Elem()
}

type fileShareArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// A name-value pair to associate with the share as metadata.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ShareName *string `pulumi:"shareName"`
	// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120).
	ShareQuota *int `pulumi:"shareQuota"`
}

// The set of arguments for constructing a FileShare resource.
type FileShareArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// A name-value pair to associate with the share as metadata.
	Metadata pulumi.StringMapInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the file share within the specified storage account. File share names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ShareName pulumi.StringPtrInput
	// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120).
	ShareQuota pulumi.IntPtrInput
}

func (FileShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareArgs)(nil)).Elem()
}

type FileShareInput interface {
	pulumi.Input

	ToFileShareOutput() FileShareOutput
	ToFileShareOutputWithContext(ctx context.Context) FileShareOutput
}

func (*FileShare) ElementType() reflect.Type {
	return reflect.TypeOf((**FileShare)(nil)).Elem()
}

func (i *FileShare) ToFileShareOutput() FileShareOutput {
	return i.ToFileShareOutputWithContext(context.Background())
}

func (i *FileShare) ToFileShareOutputWithContext(ctx context.Context) FileShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileShareOutput)
}

type FileShareOutput struct{ *pulumi.OutputState }

func (FileShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileShare)(nil)).Elem()
}

func (o FileShareOutput) ToFileShareOutput() FileShareOutput {
	return o
}

func (o FileShareOutput) ToFileShareOutputWithContext(ctx context.Context) FileShareOutput {
	return o
}

// Resource Etag.
func (o FileShareOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Returns the date and time the share was last modified.
func (o FileShareOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// A name-value pair to associate with the share as metadata.
func (o FileShareOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the resource
func (o FileShareOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The maximum size of the share, in gigabytes. Must be greater than 0, and less than or equal to 5TB (5120).
func (o FileShareOutput) ShareQuota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileShare) pulumi.IntPtrOutput { return v.ShareQuota }).(pulumi.IntPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FileShareOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FileShare) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FileShareOutput{})
}