// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240701preview

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Capacity pool resource
type Pool struct {
	pulumi.CustomResourceState

	// If enabled (true) the pool can contain cool Access enabled volumes.
	CoolAccess            pulumi.BoolPtrOutput    `pulumi:"coolAccess"`
	CustomThroughputMibps pulumi.Float64PtrOutput `pulumi:"customThroughputMibps"`
	// Encryption type of the capacity pool, set encryption type for data at rest for this pool and all volumes in it. This value can only be set when creating new pool.
	EncryptionType pulumi.StringPtrOutput `pulumi:"encryptionType"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// UUID v4 used to identify the Pool
	PoolId pulumi.StringOutput `pulumi:"poolId"`
	// Azure lifecycle management
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The qos type of the pool
	QosType pulumi.StringPtrOutput `pulumi:"qosType"`
	// The service level of the file system
	ServiceLevel pulumi.StringOutput `pulumi:"serviceLevel"`
	// Provisioned size of the pool (in bytes). Allowed values are in 1TiB chunks (value must be multiple of 1099511627776).
	Size pulumi.Float64Output `pulumi:"size"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Total throughput of pool in MiB/s
	TotalThroughputMibps pulumi.Float64Output `pulumi:"totalThroughputMibps"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Utilized throughput of pool in MiB/s
	UtilizedThroughputMibps pulumi.Float64Output `pulumi:"utilizedThroughputMibps"`
}

// NewPool registers a new resource with the given unique name, arguments, and options.
func NewPool(ctx *pulumi.Context,
	name string, args *PoolArgs, opts ...pulumi.ResourceOption) (*Pool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceLevel == nil {
		return nil, errors.New("invalid value for required argument 'ServiceLevel'")
	}
	if args.CoolAccess == nil {
		args.CoolAccess = pulumi.BoolPtr(false)
	}
	if args.EncryptionType == nil {
		args.EncryptionType = pulumi.StringPtr("Single")
	}
	if args.QosType == nil {
		args.QosType = pulumi.StringPtr("Auto")
	}
	if args.Size == nil {
		args.Size = pulumi.Float64(4398046511104.0)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp/v20170815:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190501:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190601:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190701:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190801:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191001:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191101:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200201:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200301:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220901:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20221101preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230501:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230501preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230701:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20230701preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20231101:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20231101preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20240101:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20240301:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20240301preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20240501:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20240501preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20240701:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20240901:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20240901preview:Pool"),
		},
		{
			Type: pulumi.String("azure-native:netapp:Pool"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Pool
	err := ctx.RegisterResource("azure-native:netapp/v20240701preview:Pool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPool gets an existing Pool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PoolState, opts ...pulumi.ResourceOption) (*Pool, error) {
	var resource Pool
	err := ctx.ReadResource("azure-native:netapp/v20240701preview:Pool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pool resources.
type poolState struct {
}

type PoolState struct {
}

func (PoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*poolState)(nil)).Elem()
}

type poolArgs struct {
	// The name of the NetApp account
	AccountName string `pulumi:"accountName"`
	// If enabled (true) the pool can contain cool Access enabled volumes.
	CoolAccess            *bool    `pulumi:"coolAccess"`
	CustomThroughputMibps *float64 `pulumi:"customThroughputMibps"`
	// Encryption type of the capacity pool, set encryption type for data at rest for this pool and all volumes in it. This value can only be set when creating new pool.
	EncryptionType *string `pulumi:"encryptionType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the capacity pool
	PoolName *string `pulumi:"poolName"`
	// The qos type of the pool
	QosType *string `pulumi:"qosType"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The service level of the file system
	ServiceLevel string `pulumi:"serviceLevel"`
	// Provisioned size of the pool (in bytes). Allowed values are in 1TiB chunks (value must be multiple of 1099511627776).
	Size float64 `pulumi:"size"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Pool resource.
type PoolArgs struct {
	// The name of the NetApp account
	AccountName pulumi.StringInput
	// If enabled (true) the pool can contain cool Access enabled volumes.
	CoolAccess            pulumi.BoolPtrInput
	CustomThroughputMibps pulumi.Float64PtrInput
	// Encryption type of the capacity pool, set encryption type for data at rest for this pool and all volumes in it. This value can only be set when creating new pool.
	EncryptionType pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the capacity pool
	PoolName pulumi.StringPtrInput
	// The qos type of the pool
	QosType pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The service level of the file system
	ServiceLevel pulumi.StringInput
	// Provisioned size of the pool (in bytes). Allowed values are in 1TiB chunks (value must be multiple of 1099511627776).
	Size pulumi.Float64Input
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (PoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*poolArgs)(nil)).Elem()
}

type PoolInput interface {
	pulumi.Input

	ToPoolOutput() PoolOutput
	ToPoolOutputWithContext(ctx context.Context) PoolOutput
}

func (*Pool) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (i *Pool) ToPoolOutput() PoolOutput {
	return i.ToPoolOutputWithContext(context.Background())
}

func (i *Pool) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoolOutput)
}

type PoolOutput struct{ *pulumi.OutputState }

func (PoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pool)(nil)).Elem()
}

func (o PoolOutput) ToPoolOutput() PoolOutput {
	return o
}

func (o PoolOutput) ToPoolOutputWithContext(ctx context.Context) PoolOutput {
	return o
}

// If enabled (true) the pool can contain cool Access enabled volumes.
func (o PoolOutput) CoolAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.BoolPtrOutput { return v.CoolAccess }).(pulumi.BoolPtrOutput)
}

func (o PoolOutput) CustomThroughputMibps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.Float64PtrOutput { return v.CustomThroughputMibps }).(pulumi.Float64PtrOutput)
}

// Encryption type of the capacity pool, set encryption type for data at rest for this pool and all volumes in it. This value can only be set when creating new pool.
func (o PoolOutput) EncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.EncryptionType }).(pulumi.StringPtrOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o PoolOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o PoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o PoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// UUID v4 used to identify the Pool
func (o PoolOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.PoolId }).(pulumi.StringOutput)
}

// Azure lifecycle management
func (o PoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The qos type of the pool
func (o PoolOutput) QosType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringPtrOutput { return v.QosType }).(pulumi.StringPtrOutput)
}

// The service level of the file system
func (o PoolOutput) ServiceLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.ServiceLevel }).(pulumi.StringOutput)
}

// Provisioned size of the pool (in bytes). Allowed values are in 1TiB chunks (value must be multiple of 1099511627776).
func (o PoolOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v *Pool) pulumi.Float64Output { return v.Size }).(pulumi.Float64Output)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o PoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Pool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o PoolOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Total throughput of pool in MiB/s
func (o PoolOutput) TotalThroughputMibps() pulumi.Float64Output {
	return o.ApplyT(func(v *Pool) pulumi.Float64Output { return v.TotalThroughputMibps }).(pulumi.Float64Output)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o PoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Pool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Utilized throughput of pool in MiB/s
func (o PoolOutput) UtilizedThroughputMibps() pulumi.Float64Output {
	return o.ApplyT(func(v *Pool) pulumi.Float64Output { return v.UtilizedThroughputMibps }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(PoolOutput{})
}
