// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databoxedge

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The storage account credential.
//
// Uses Azure REST API version 2023-07-01. In version 2.x of the Azure Native provider, it used API version 2022-03-01.
//
// Other available API versions: 2022-03-01, 2022-04-01-preview, 2022-12-01-preview, 2023-01-01-preview, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native databoxedge [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type StorageAccountCredential struct {
	pulumi.CustomResourceState

	// Encrypted storage key.
	AccountKey AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"accountKey"`
	// Type of storage accessed on the storage account.
	AccountType pulumi.StringOutput `pulumi:"accountType"`
	// Alias for the storage account.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// Blob end point for private clouds.
	BlobDomainName pulumi.StringPtrOutput `pulumi:"blobDomainName"`
	// Connection string for the storage account. Use this string if username and account key are not specified.
	ConnectionString pulumi.StringPtrOutput `pulumi:"connectionString"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Signifies whether SSL needs to be enabled or not.
	SslStatus pulumi.StringOutput `pulumi:"sslStatus"`
	// Id of the storage account.
	StorageAccountId pulumi.StringPtrOutput `pulumi:"storageAccountId"`
	// Metadata pertaining to creation and last modification of StorageAccountCredential
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// Username for the storage account.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
}

// NewStorageAccountCredential registers a new resource with the given unique name, arguments, and options.
func NewStorageAccountCredential(ctx *pulumi.Context,
	name string, args *StorageAccountCredentialArgs, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountType == nil {
		return nil, errors.New("invalid value for required argument 'AccountType'")
	}
	if args.Alias == nil {
		return nil, errors.New("invalid value for required argument 'Alias'")
	}
	if args.DeviceName == nil {
		return nil, errors.New("invalid value for required argument 'DeviceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SslStatus == nil {
		return nil, errors.New("invalid value for required argument 'SslStatus'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databoxedge/v20190301:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190701:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20190801:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200501preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20200901preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20201201:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210201preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20210601preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220301:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20220401preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20221201preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230101preview:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20230701:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azure-native:databoxedge/v20231201:StorageAccountCredential"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource StorageAccountCredential
	err := ctx.RegisterResource("azure-native:databoxedge:StorageAccountCredential", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:databoxedge:StorageAccountCredential", name, id, state, &resource, opts...)
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
	// Encrypted storage key.
	AccountKey *AsymmetricEncryptedSecret `pulumi:"accountKey"`
	// Type of storage accessed on the storage account.
	AccountType string `pulumi:"accountType"`
	// Alias for the storage account.
	Alias string `pulumi:"alias"`
	// Blob end point for private clouds.
	BlobDomainName *string `pulumi:"blobDomainName"`
	// Connection string for the storage account. Use this string if username and account key are not specified.
	ConnectionString *string `pulumi:"connectionString"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The storage account credential name.
	Name *string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Signifies whether SSL needs to be enabled or not.
	SslStatus string `pulumi:"sslStatus"`
	// Id of the storage account.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// Username for the storage account.
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a StorageAccountCredential resource.
type StorageAccountCredentialArgs struct {
	// Encrypted storage key.
	AccountKey AsymmetricEncryptedSecretPtrInput
	// Type of storage accessed on the storage account.
	AccountType pulumi.StringInput
	// Alias for the storage account.
	Alias pulumi.StringInput
	// Blob end point for private clouds.
	BlobDomainName pulumi.StringPtrInput
	// Connection string for the storage account. Use this string if username and account key are not specified.
	ConnectionString pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// The storage account credential name.
	Name pulumi.StringPtrInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// Signifies whether SSL needs to be enabled or not.
	SslStatus pulumi.StringInput
	// Id of the storage account.
	StorageAccountId pulumi.StringPtrInput
	// Username for the storage account.
	UserName pulumi.StringPtrInput
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

// Encrypted storage key.
func (o StorageAccountCredentialOutput) AccountKey() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) AsymmetricEncryptedSecretResponsePtrOutput { return v.AccountKey }).(AsymmetricEncryptedSecretResponsePtrOutput)
}

// Type of storage accessed on the storage account.
func (o StorageAccountCredentialOutput) AccountType() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.AccountType }).(pulumi.StringOutput)
}

// Alias for the storage account.
func (o StorageAccountCredentialOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// The Azure API version of the resource.
func (o StorageAccountCredentialOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Blob end point for private clouds.
func (o StorageAccountCredentialOutput) BlobDomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringPtrOutput { return v.BlobDomainName }).(pulumi.StringPtrOutput)
}

// Connection string for the storage account. Use this string if username and account key are not specified.
func (o StorageAccountCredentialOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringPtrOutput { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

// The object name.
func (o StorageAccountCredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Signifies whether SSL needs to be enabled or not.
func (o StorageAccountCredentialOutput) SslStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.SslStatus }).(pulumi.StringOutput)
}

// Id of the storage account.
func (o StorageAccountCredentialOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringPtrOutput { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

// Metadata pertaining to creation and last modification of StorageAccountCredential
func (o StorageAccountCredentialOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *StorageAccountCredential) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The hierarchical type of the object.
func (o StorageAccountCredentialOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Username for the storage account.
func (o StorageAccountCredentialOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountCredential) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageAccountCredentialOutput{})
}
