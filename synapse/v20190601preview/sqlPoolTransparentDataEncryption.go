// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Sql pool transparent data encryption configuration.
//
// Deprecated: Version 2019-06-01-preview will be removed in v2 of the provider.
type SqlPoolTransparentDataEncryption struct {
	pulumi.CustomResourceState

	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the database transparent data encryption.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlPoolTransparentDataEncryption registers a new resource with the given unique name, arguments, and options.
func NewSqlPoolTransparentDataEncryption(ctx *pulumi.Context,
	name string, args *SqlPoolTransparentDataEncryptionArgs, opts ...pulumi.ResourceOption) (*SqlPoolTransparentDataEncryption, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlPoolName == nil {
		return nil, errors.New("invalid value for required argument 'SqlPoolName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:SqlPoolTransparentDataEncryption"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:SqlPoolTransparentDataEncryption"),
		},
	})
	opts = append(opts, aliases)
	var resource SqlPoolTransparentDataEncryption
	err := ctx.RegisterResource("azure-native:synapse/v20190601preview:SqlPoolTransparentDataEncryption", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlPoolTransparentDataEncryption gets an existing SqlPoolTransparentDataEncryption resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlPoolTransparentDataEncryption(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlPoolTransparentDataEncryptionState, opts ...pulumi.ResourceOption) (*SqlPoolTransparentDataEncryption, error) {
	var resource SqlPoolTransparentDataEncryption
	err := ctx.ReadResource("azure-native:synapse/v20190601preview:SqlPoolTransparentDataEncryption", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlPoolTransparentDataEncryption resources.
type sqlPoolTransparentDataEncryptionState struct {
}

type SqlPoolTransparentDataEncryptionState struct {
}

func (SqlPoolTransparentDataEncryptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolTransparentDataEncryptionState)(nil)).Elem()
}

type sqlPoolTransparentDataEncryptionArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SQL pool name
	SqlPoolName string `pulumi:"sqlPoolName"`
	// The status of the database transparent data encryption.
	Status *string `pulumi:"status"`
	// The name of the transparent data encryption configuration.
	TransparentDataEncryptionName *string `pulumi:"transparentDataEncryptionName"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a SqlPoolTransparentDataEncryption resource.
type SqlPoolTransparentDataEncryptionArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// SQL pool name
	SqlPoolName pulumi.StringInput
	// The status of the database transparent data encryption.
	Status pulumi.StringPtrInput
	// The name of the transparent data encryption configuration.
	TransparentDataEncryptionName pulumi.StringPtrInput
	// The name of the workspace
	WorkspaceName pulumi.StringInput
}

func (SqlPoolTransparentDataEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlPoolTransparentDataEncryptionArgs)(nil)).Elem()
}

type SqlPoolTransparentDataEncryptionInput interface {
	pulumi.Input

	ToSqlPoolTransparentDataEncryptionOutput() SqlPoolTransparentDataEncryptionOutput
	ToSqlPoolTransparentDataEncryptionOutputWithContext(ctx context.Context) SqlPoolTransparentDataEncryptionOutput
}

func (*SqlPoolTransparentDataEncryption) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPoolTransparentDataEncryption)(nil)).Elem()
}

func (i *SqlPoolTransparentDataEncryption) ToSqlPoolTransparentDataEncryptionOutput() SqlPoolTransparentDataEncryptionOutput {
	return i.ToSqlPoolTransparentDataEncryptionOutputWithContext(context.Background())
}

func (i *SqlPoolTransparentDataEncryption) ToSqlPoolTransparentDataEncryptionOutputWithContext(ctx context.Context) SqlPoolTransparentDataEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlPoolTransparentDataEncryptionOutput)
}

type SqlPoolTransparentDataEncryptionOutput struct{ *pulumi.OutputState }

func (SqlPoolTransparentDataEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlPoolTransparentDataEncryption)(nil)).Elem()
}

func (o SqlPoolTransparentDataEncryptionOutput) ToSqlPoolTransparentDataEncryptionOutput() SqlPoolTransparentDataEncryptionOutput {
	return o
}

func (o SqlPoolTransparentDataEncryptionOutput) ToSqlPoolTransparentDataEncryptionOutputWithContext(ctx context.Context) SqlPoolTransparentDataEncryptionOutput {
	return o
}

// Resource location.
func (o SqlPoolTransparentDataEncryptionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPoolTransparentDataEncryption) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o SqlPoolTransparentDataEncryptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPoolTransparentDataEncryption) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of the database transparent data encryption.
func (o SqlPoolTransparentDataEncryptionOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlPoolTransparentDataEncryption) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o SqlPoolTransparentDataEncryptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SqlPoolTransparentDataEncryption) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SqlPoolTransparentDataEncryptionOutput{})
}