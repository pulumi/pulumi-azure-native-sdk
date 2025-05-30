// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudngfw

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// LocalRulestack fqdnList
//
// Uses Azure REST API version 2025-02-06-preview. In version 2.x of the Azure Native provider, it used API version 2023-09-01.
//
// Other available API versions: 2023-09-01, 2023-10-10-preview, 2024-01-19-preview, 2024-02-07-preview. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native cloudngfw [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type FqdnListLocalRulestack struct {
	pulumi.CustomResourceState

	// comment for this object
	AuditComment pulumi.StringPtrOutput `pulumi:"auditComment"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// fqdn object description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// etag info
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// fqdn list
	FqdnList pulumi.StringArrayOutput `pulumi:"fqdnList"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFqdnListLocalRulestack registers a new resource with the given unique name, arguments, and options.
func NewFqdnListLocalRulestack(ctx *pulumi.Context,
	name string, args *FqdnListLocalRulestackArgs, opts ...pulumi.ResourceOption) (*FqdnListLocalRulestack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FqdnList == nil {
		return nil, errors.New("invalid value for required argument 'FqdnList'")
	}
	if args.LocalRulestackName == nil {
		return nil, errors.New("invalid value for required argument 'LocalRulestackName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:cloudngfw/v20220829:FqdnListLocalRulestack"),
		},
		{
			Type: pulumi.String("azure-native:cloudngfw/v20220829preview:FqdnListLocalRulestack"),
		},
		{
			Type: pulumi.String("azure-native:cloudngfw/v20230901:FqdnListLocalRulestack"),
		},
		{
			Type: pulumi.String("azure-native:cloudngfw/v20230901preview:FqdnListLocalRulestack"),
		},
		{
			Type: pulumi.String("azure-native:cloudngfw/v20231010preview:FqdnListLocalRulestack"),
		},
		{
			Type: pulumi.String("azure-native:cloudngfw/v20240119preview:FqdnListLocalRulestack"),
		},
		{
			Type: pulumi.String("azure-native:cloudngfw/v20240207preview:FqdnListLocalRulestack"),
		},
		{
			Type: pulumi.String("azure-native:cloudngfw/v20250206preview:FqdnListLocalRulestack"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource FqdnListLocalRulestack
	err := ctx.RegisterResource("azure-native:cloudngfw:FqdnListLocalRulestack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFqdnListLocalRulestack gets an existing FqdnListLocalRulestack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFqdnListLocalRulestack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FqdnListLocalRulestackState, opts ...pulumi.ResourceOption) (*FqdnListLocalRulestack, error) {
	var resource FqdnListLocalRulestack
	err := ctx.ReadResource("azure-native:cloudngfw:FqdnListLocalRulestack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FqdnListLocalRulestack resources.
type fqdnListLocalRulestackState struct {
}

type FqdnListLocalRulestackState struct {
}

func (FqdnListLocalRulestackState) ElementType() reflect.Type {
	return reflect.TypeOf((*fqdnListLocalRulestackState)(nil)).Elem()
}

type fqdnListLocalRulestackArgs struct {
	// comment for this object
	AuditComment *string `pulumi:"auditComment"`
	// fqdn object description
	Description *string `pulumi:"description"`
	// fqdn list
	FqdnList []string `pulumi:"fqdnList"`
	// LocalRulestack resource name
	LocalRulestackName string `pulumi:"localRulestackName"`
	// fqdn list name
	Name *string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a FqdnListLocalRulestack resource.
type FqdnListLocalRulestackArgs struct {
	// comment for this object
	AuditComment pulumi.StringPtrInput
	// fqdn object description
	Description pulumi.StringPtrInput
	// fqdn list
	FqdnList pulumi.StringArrayInput
	// LocalRulestack resource name
	LocalRulestackName pulumi.StringInput
	// fqdn list name
	Name pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (FqdnListLocalRulestackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fqdnListLocalRulestackArgs)(nil)).Elem()
}

type FqdnListLocalRulestackInput interface {
	pulumi.Input

	ToFqdnListLocalRulestackOutput() FqdnListLocalRulestackOutput
	ToFqdnListLocalRulestackOutputWithContext(ctx context.Context) FqdnListLocalRulestackOutput
}

func (*FqdnListLocalRulestack) ElementType() reflect.Type {
	return reflect.TypeOf((**FqdnListLocalRulestack)(nil)).Elem()
}

func (i *FqdnListLocalRulestack) ToFqdnListLocalRulestackOutput() FqdnListLocalRulestackOutput {
	return i.ToFqdnListLocalRulestackOutputWithContext(context.Background())
}

func (i *FqdnListLocalRulestack) ToFqdnListLocalRulestackOutputWithContext(ctx context.Context) FqdnListLocalRulestackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FqdnListLocalRulestackOutput)
}

type FqdnListLocalRulestackOutput struct{ *pulumi.OutputState }

func (FqdnListLocalRulestackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FqdnListLocalRulestack)(nil)).Elem()
}

func (o FqdnListLocalRulestackOutput) ToFqdnListLocalRulestackOutput() FqdnListLocalRulestackOutput {
	return o
}

func (o FqdnListLocalRulestackOutput) ToFqdnListLocalRulestackOutputWithContext(ctx context.Context) FqdnListLocalRulestackOutput {
	return o
}

// comment for this object
func (o FqdnListLocalRulestackOutput) AuditComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FqdnListLocalRulestack) pulumi.StringPtrOutput { return v.AuditComment }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o FqdnListLocalRulestackOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *FqdnListLocalRulestack) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// fqdn object description
func (o FqdnListLocalRulestackOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FqdnListLocalRulestack) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// etag info
func (o FqdnListLocalRulestackOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FqdnListLocalRulestack) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

// fqdn list
func (o FqdnListLocalRulestackOutput) FqdnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FqdnListLocalRulestack) pulumi.StringArrayOutput { return v.FqdnList }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o FqdnListLocalRulestackOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FqdnListLocalRulestack) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o FqdnListLocalRulestackOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *FqdnListLocalRulestack) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o FqdnListLocalRulestackOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FqdnListLocalRulestack) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o FqdnListLocalRulestackOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FqdnListLocalRulestack) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FqdnListLocalRulestackOutput{})
}
