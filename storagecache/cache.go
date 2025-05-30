// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storagecache

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A cache instance. Follows Azure Resource Manager standards: https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md
//
// Uses Azure REST API version 2024-03-01. In version 2.x of the Azure Native provider, it used API version 2023-05-01.
//
// Other available API versions: 2023-05-01, 2023-11-01-preview, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native storagecache [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type Cache struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// The size of this Cache, in GB.
	CacheSizeGB pulumi.IntPtrOutput `pulumi:"cacheSizeGB"`
	// Specifies Directory Services settings of the cache.
	DirectoryServicesSettings CacheDirectorySettingsResponsePtrOutput `pulumi:"directoryServicesSettings"`
	// Specifies encryption settings of the cache.
	EncryptionSettings CacheEncryptionSettingsResponsePtrOutput `pulumi:"encryptionSettings"`
	// Health of the cache.
	Health CacheHealthResponseOutput `pulumi:"health"`
	// The identity of the cache, if configured.
	Identity CacheIdentityResponsePtrOutput `pulumi:"identity"`
	// Region name string.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Array of IPv4 addresses that can be used by clients mounting this cache.
	MountAddresses pulumi.StringArrayOutput `pulumi:"mountAddresses"`
	// Name of cache.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies network settings of the cache.
	NetworkSettings CacheNetworkSettingsResponsePtrOutput `pulumi:"networkSettings"`
	// Specifies the priming jobs defined in the cache.
	PrimingJobs PrimingJobResponseArrayOutput `pulumi:"primingJobs"`
	// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Specifies security settings of the cache.
	SecuritySettings CacheSecuritySettingsResponsePtrOutput `pulumi:"securitySettings"`
	// SKU for the cache.
	Sku CacheResponseSkuPtrOutput `pulumi:"sku"`
	// Specifies the space allocation percentage for each storage target in the cache.
	SpaceAllocation StorageTargetSpaceAllocationResponseArrayOutput `pulumi:"spaceAllocation"`
	// Subnet used for the cache.
	Subnet pulumi.StringPtrOutput `pulumi:"subnet"`
	// The system meta data relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the cache; Microsoft.StorageCache/Cache
	Type pulumi.StringOutput `pulumi:"type"`
	// Upgrade settings of the cache.
	UpgradeSettings CacheUpgradeSettingsResponsePtrOutput `pulumi:"upgradeSettings"`
	// Upgrade status of the cache.
	UpgradeStatus CacheUpgradeStatusResponseOutput `pulumi:"upgradeStatus"`
	// Availability zones for resources. This field should only contain a single element in the array.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewCache registers a new resource with the given unique name, arguments, and options.
func NewCache(ctx *pulumi.Context,
	name string, args *CacheArgs, opts ...pulumi.ResourceOption) (*Cache, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.DirectoryServicesSettings != nil {
		args.DirectoryServicesSettings = args.DirectoryServicesSettings.ToCacheDirectorySettingsPtrOutput().ApplyT(func(v *CacheDirectorySettings) *CacheDirectorySettings { return v.Defaults() }).(CacheDirectorySettingsPtrOutput)
	}
	if args.NetworkSettings != nil {
		args.NetworkSettings = args.NetworkSettings.ToCacheNetworkSettingsPtrOutput().ApplyT(func(v *CacheNetworkSettings) *CacheNetworkSettings { return v.Defaults() }).(CacheNetworkSettingsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:storagecache/v20190801preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20191101:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20200301:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20201001:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210301:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210501:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20210901:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20220101:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20220501:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20230101:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20230301preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20230501:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20231101preview:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20240301:Cache"),
		},
		{
			Type: pulumi.String("azure-native:storagecache/v20240701:Cache"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Cache
	err := ctx.RegisterResource("azure-native:storagecache:Cache", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCache gets an existing Cache resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCache(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CacheState, opts ...pulumi.ResourceOption) (*Cache, error) {
	var resource Cache
	err := ctx.ReadResource("azure-native:storagecache:Cache", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cache resources.
type cacheState struct {
}

type CacheState struct {
}

func (CacheState) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheState)(nil)).Elem()
}

type cacheArgs struct {
	// Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
	CacheName *string `pulumi:"cacheName"`
	// The size of this Cache, in GB.
	CacheSizeGB *int `pulumi:"cacheSizeGB"`
	// Specifies Directory Services settings of the cache.
	DirectoryServicesSettings *CacheDirectorySettings `pulumi:"directoryServicesSettings"`
	// Specifies encryption settings of the cache.
	EncryptionSettings *CacheEncryptionSettings `pulumi:"encryptionSettings"`
	// The identity of the cache, if configured.
	Identity *CacheIdentity `pulumi:"identity"`
	// Region name string.
	Location *string `pulumi:"location"`
	// Specifies network settings of the cache.
	NetworkSettings *CacheNetworkSettings `pulumi:"networkSettings"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies security settings of the cache.
	SecuritySettings *CacheSecuritySettings `pulumi:"securitySettings"`
	// SKU for the cache.
	Sku *CacheSku `pulumi:"sku"`
	// Subnet used for the cache.
	Subnet *string `pulumi:"subnet"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Upgrade settings of the cache.
	UpgradeSettings *CacheUpgradeSettings `pulumi:"upgradeSettings"`
	// Availability zones for resources. This field should only contain a single element in the array.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a Cache resource.
type CacheArgs struct {
	// Name of cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
	CacheName pulumi.StringPtrInput
	// The size of this Cache, in GB.
	CacheSizeGB pulumi.IntPtrInput
	// Specifies Directory Services settings of the cache.
	DirectoryServicesSettings CacheDirectorySettingsPtrInput
	// Specifies encryption settings of the cache.
	EncryptionSettings CacheEncryptionSettingsPtrInput
	// The identity of the cache, if configured.
	Identity CacheIdentityPtrInput
	// Region name string.
	Location pulumi.StringPtrInput
	// Specifies network settings of the cache.
	NetworkSettings CacheNetworkSettingsPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Specifies security settings of the cache.
	SecuritySettings CacheSecuritySettingsPtrInput
	// SKU for the cache.
	Sku CacheSkuPtrInput
	// Subnet used for the cache.
	Subnet pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Upgrade settings of the cache.
	UpgradeSettings CacheUpgradeSettingsPtrInput
	// Availability zones for resources. This field should only contain a single element in the array.
	Zones pulumi.StringArrayInput
}

func (CacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheArgs)(nil)).Elem()
}

type CacheInput interface {
	pulumi.Input

	ToCacheOutput() CacheOutput
	ToCacheOutputWithContext(ctx context.Context) CacheOutput
}

func (*Cache) ElementType() reflect.Type {
	return reflect.TypeOf((**Cache)(nil)).Elem()
}

func (i *Cache) ToCacheOutput() CacheOutput {
	return i.ToCacheOutputWithContext(context.Background())
}

func (i *Cache) ToCacheOutputWithContext(ctx context.Context) CacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheOutput)
}

type CacheOutput struct{ *pulumi.OutputState }

func (CacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cache)(nil)).Elem()
}

func (o CacheOutput) ToCacheOutput() CacheOutput {
	return o
}

func (o CacheOutput) ToCacheOutputWithContext(ctx context.Context) CacheOutput {
	return o
}

// The Azure API version of the resource.
func (o CacheOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// The size of this Cache, in GB.
func (o CacheOutput) CacheSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cache) pulumi.IntPtrOutput { return v.CacheSizeGB }).(pulumi.IntPtrOutput)
}

// Specifies Directory Services settings of the cache.
func (o CacheOutput) DirectoryServicesSettings() CacheDirectorySettingsResponsePtrOutput {
	return o.ApplyT(func(v *Cache) CacheDirectorySettingsResponsePtrOutput { return v.DirectoryServicesSettings }).(CacheDirectorySettingsResponsePtrOutput)
}

// Specifies encryption settings of the cache.
func (o CacheOutput) EncryptionSettings() CacheEncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Cache) CacheEncryptionSettingsResponsePtrOutput { return v.EncryptionSettings }).(CacheEncryptionSettingsResponsePtrOutput)
}

// Health of the cache.
func (o CacheOutput) Health() CacheHealthResponseOutput {
	return o.ApplyT(func(v *Cache) CacheHealthResponseOutput { return v.Health }).(CacheHealthResponseOutput)
}

// The identity of the cache, if configured.
func (o CacheOutput) Identity() CacheIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Cache) CacheIdentityResponsePtrOutput { return v.Identity }).(CacheIdentityResponsePtrOutput)
}

// Region name string.
func (o CacheOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Array of IPv4 addresses that can be used by clients mounting this cache.
func (o CacheOutput) MountAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringArrayOutput { return v.MountAddresses }).(pulumi.StringArrayOutput)
}

// Name of cache.
func (o CacheOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies network settings of the cache.
func (o CacheOutput) NetworkSettings() CacheNetworkSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Cache) CacheNetworkSettingsResponsePtrOutput { return v.NetworkSettings }).(CacheNetworkSettingsResponsePtrOutput)
}

// Specifies the priming jobs defined in the cache.
func (o CacheOutput) PrimingJobs() PrimingJobResponseArrayOutput {
	return o.ApplyT(func(v *Cache) PrimingJobResponseArrayOutput { return v.PrimingJobs }).(PrimingJobResponseArrayOutput)
}

// ARM provisioning state, see https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/Addendum.md#provisioningstate-property
func (o CacheOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Specifies security settings of the cache.
func (o CacheOutput) SecuritySettings() CacheSecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v *Cache) CacheSecuritySettingsResponsePtrOutput { return v.SecuritySettings }).(CacheSecuritySettingsResponsePtrOutput)
}

// SKU for the cache.
func (o CacheOutput) Sku() CacheResponseSkuPtrOutput {
	return o.ApplyT(func(v *Cache) CacheResponseSkuPtrOutput { return v.Sku }).(CacheResponseSkuPtrOutput)
}

// Specifies the space allocation percentage for each storage target in the cache.
func (o CacheOutput) SpaceAllocation() StorageTargetSpaceAllocationResponseArrayOutput {
	return o.ApplyT(func(v *Cache) StorageTargetSpaceAllocationResponseArrayOutput { return v.SpaceAllocation }).(StorageTargetSpaceAllocationResponseArrayOutput)
}

// Subnet used for the cache.
func (o CacheOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringPtrOutput { return v.Subnet }).(pulumi.StringPtrOutput)
}

// The system meta data relating to this resource.
func (o CacheOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Cache) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o CacheOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of the cache; Microsoft.StorageCache/Cache
func (o CacheOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Upgrade settings of the cache.
func (o CacheOutput) UpgradeSettings() CacheUpgradeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Cache) CacheUpgradeSettingsResponsePtrOutput { return v.UpgradeSettings }).(CacheUpgradeSettingsResponsePtrOutput)
}

// Upgrade status of the cache.
func (o CacheOutput) UpgradeStatus() CacheUpgradeStatusResponseOutput {
	return o.ApplyT(func(v *Cache) CacheUpgradeStatusResponseOutput { return v.UpgradeStatus }).(CacheUpgradeStatusResponseOutput)
}

// Availability zones for resources. This field should only contain a single element in the array.
func (o CacheOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cache) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(CacheOutput{})
}
