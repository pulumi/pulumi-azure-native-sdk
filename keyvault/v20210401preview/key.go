// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The key resource.
type Key struct {
	pulumi.CustomResourceState

	// The attributes of the key.
	Attributes KeyAttributesResponsePtrOutput `pulumi:"attributes"`
	// The elliptic curve name. For valid values, see JsonWebKeyCurveName.
	CurveName pulumi.StringPtrOutput   `pulumi:"curveName"`
	KeyOps    pulumi.StringArrayOutput `pulumi:"keyOps"`
	// The key size in bits. For example: 2048, 3072, or 4096 for RSA.
	KeySize pulumi.IntPtrOutput `pulumi:"keySize"`
	// The URI to retrieve the current version of the key.
	KeyUri pulumi.StringOutput `pulumi:"keyUri"`
	// The URI to retrieve the specific version of the key.
	KeyUriWithVersion pulumi.StringOutput `pulumi:"keyUriWithVersion"`
	// The type of the key. For valid values, see JsonWebKeyType.
	Kty pulumi.StringPtrOutput `pulumi:"kty"`
	// Azure location of the key vault resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the key vault resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key rotation policy in response. It will be used for both output and input. Omitted if empty
	RotationPolicy RotationPolicyResponsePtrOutput `pulumi:"rotationPolicy"`
	// Tags assigned to the key vault resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type of the key vault resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:keyvault:Key"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20190901:Key"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20200401preview:Key"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20210601preview:Key"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20211001:Key"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20211101preview:Key"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20220701:Key"),
		},
	})
	opts = append(opts, aliases)
	var resource Key
	err := ctx.RegisterResource("azure-native:keyvault/v20210401preview:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("azure-native:keyvault/v20210401preview:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
}

type KeyState struct {
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// The name of the key to be created.
	KeyName *string `pulumi:"keyName"`
	// The properties of the key to be created.
	Properties KeyProperties `pulumi:"properties"`
	// The name of the resource group which contains the specified key vault.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The tags that will be assigned to the key.
	Tags map[string]string `pulumi:"tags"`
	// The name of the key vault which contains the key to be created.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// The name of the key to be created.
	KeyName pulumi.StringPtrInput
	// The properties of the key to be created.
	Properties KeyPropertiesInput
	// The name of the resource group which contains the specified key vault.
	ResourceGroupName pulumi.StringInput
	// The tags that will be assigned to the key.
	Tags pulumi.StringMapInput
	// The name of the key vault which contains the key to be created.
	VaultName pulumi.StringInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

// The attributes of the key.
func (o KeyOutput) Attributes() KeyAttributesResponsePtrOutput {
	return o.ApplyT(func(v *Key) KeyAttributesResponsePtrOutput { return v.Attributes }).(KeyAttributesResponsePtrOutput)
}

// The elliptic curve name. For valid values, see JsonWebKeyCurveName.
func (o KeyOutput) CurveName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.CurveName }).(pulumi.StringPtrOutput)
}

func (o KeyOutput) KeyOps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Key) pulumi.StringArrayOutput { return v.KeyOps }).(pulumi.StringArrayOutput)
}

// The key size in bits. For example: 2048, 3072, or 4096 for RSA.
func (o KeyOutput) KeySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.IntPtrOutput { return v.KeySize }).(pulumi.IntPtrOutput)
}

// The URI to retrieve the current version of the key.
func (o KeyOutput) KeyUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.KeyUri }).(pulumi.StringOutput)
}

// The URI to retrieve the specific version of the key.
func (o KeyOutput) KeyUriWithVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.KeyUriWithVersion }).(pulumi.StringOutput)
}

// The type of the key. For valid values, see JsonWebKeyType.
func (o KeyOutput) Kty() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Key) pulumi.StringPtrOutput { return v.Kty }).(pulumi.StringPtrOutput)
}

// Azure location of the key vault resource.
func (o KeyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the key vault resource.
func (o KeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Key rotation policy in response. It will be used for both output and input. Omitted if empty
func (o KeyOutput) RotationPolicy() RotationPolicyResponsePtrOutput {
	return o.ApplyT(func(v *Key) RotationPolicyResponsePtrOutput { return v.RotationPolicy }).(RotationPolicyResponsePtrOutput)
}

// Tags assigned to the key vault resource.
func (o KeyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Key) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type of the key vault resource.
func (o KeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Key) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyOutput{})
}