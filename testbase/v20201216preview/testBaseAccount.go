// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201216preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Test Base Account resource.
//
// Deprecated: Version 2020-12-16-preview will be removed in v2 of the provider.
type TestBaseAccount struct {
	pulumi.CustomResourceState

	// The access level of the Test Base Account.
	AccessLevel pulumi.StringOutput `pulumi:"accessLevel"`
	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The SKU of the Test Base Account.
	Sku TestBaseAccountSKUResponseOutput `pulumi:"sku"`
	// The system metadata relating to this resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTestBaseAccount registers a new resource with the given unique name, arguments, and options.
func NewTestBaseAccount(ctx *pulumi.Context,
	name string, args *TestBaseAccountArgs, opts ...pulumi.ResourceOption) (*TestBaseAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:testbase:TestBaseAccount"),
		},
		{
			Type: pulumi.String("azure-native:testbase/v20220401preview:TestBaseAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource TestBaseAccount
	err := ctx.RegisterResource("azure-native:testbase/v20201216preview:TestBaseAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTestBaseAccount gets an existing TestBaseAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTestBaseAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TestBaseAccountState, opts ...pulumi.ResourceOption) (*TestBaseAccount, error) {
	var resource TestBaseAccount
	err := ctx.ReadResource("azure-native:testbase/v20201216preview:TestBaseAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TestBaseAccount resources.
type testBaseAccountState struct {
}

type TestBaseAccountState struct {
}

func (TestBaseAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*testBaseAccountState)(nil)).Elem()
}

type testBaseAccountArgs struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group that contains the resource.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The flag indicating if we would like to restore the Test Base Accounts which were soft deleted before.
	Restore *bool `pulumi:"restore"`
	// The SKU of the Test Base Account.
	Sku TestBaseAccountSKU `pulumi:"sku"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The resource name of the Test Base Account.
	TestBaseAccountName *string `pulumi:"testBaseAccountName"`
}

// The set of arguments for constructing a TestBaseAccount resource.
type TestBaseAccountArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group that contains the resource.
	ResourceGroupName pulumi.StringInput
	// The flag indicating if we would like to restore the Test Base Accounts which were soft deleted before.
	Restore pulumi.BoolPtrInput
	// The SKU of the Test Base Account.
	Sku TestBaseAccountSKUInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The resource name of the Test Base Account.
	TestBaseAccountName pulumi.StringPtrInput
}

func (TestBaseAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*testBaseAccountArgs)(nil)).Elem()
}

type TestBaseAccountInput interface {
	pulumi.Input

	ToTestBaseAccountOutput() TestBaseAccountOutput
	ToTestBaseAccountOutputWithContext(ctx context.Context) TestBaseAccountOutput
}

func (*TestBaseAccount) ElementType() reflect.Type {
	return reflect.TypeOf((**TestBaseAccount)(nil)).Elem()
}

func (i *TestBaseAccount) ToTestBaseAccountOutput() TestBaseAccountOutput {
	return i.ToTestBaseAccountOutputWithContext(context.Background())
}

func (i *TestBaseAccount) ToTestBaseAccountOutputWithContext(ctx context.Context) TestBaseAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TestBaseAccountOutput)
}

type TestBaseAccountOutput struct{ *pulumi.OutputState }

func (TestBaseAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TestBaseAccount)(nil)).Elem()
}

func (o TestBaseAccountOutput) ToTestBaseAccountOutput() TestBaseAccountOutput {
	return o
}

func (o TestBaseAccountOutput) ToTestBaseAccountOutputWithContext(ctx context.Context) TestBaseAccountOutput {
	return o
}

// The access level of the Test Base Account.
func (o TestBaseAccountOutput) AccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *TestBaseAccount) pulumi.StringOutput { return v.AccessLevel }).(pulumi.StringOutput)
}

// Resource Etag.
func (o TestBaseAccountOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TestBaseAccount) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o TestBaseAccountOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TestBaseAccount) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name.
func (o TestBaseAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TestBaseAccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The provisioning state of the resource.
func (o TestBaseAccountOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TestBaseAccount) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The SKU of the Test Base Account.
func (o TestBaseAccountOutput) Sku() TestBaseAccountSKUResponseOutput {
	return o.ApplyT(func(v *TestBaseAccount) TestBaseAccountSKUResponseOutput { return v.Sku }).(TestBaseAccountSKUResponseOutput)
}

// The system metadata relating to this resource
func (o TestBaseAccountOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TestBaseAccount) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags of the resource.
func (o TestBaseAccountOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TestBaseAccount) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type.
func (o TestBaseAccountOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TestBaseAccount) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TestBaseAccountOutput{})
}