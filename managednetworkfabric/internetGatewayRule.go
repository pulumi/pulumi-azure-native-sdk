// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package managednetworkfabric

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Internet Gateway Rule resource definition.
//
// Uses Azure REST API version 2023-06-15. In version 2.x of the Azure Native provider, it used API version 2023-06-15.
type InternetGatewayRule struct {
	pulumi.CustomResourceState

	// Switch configuration description.
	Annotation pulumi.StringPtrOutput `pulumi:"annotation"`
	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// List of Internet Gateway resource Id.
	InternetGatewayIds pulumi.StringArrayOutput `pulumi:"internetGatewayIds"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Rules for the InternetGateways
	RuleProperties RulePropertiesResponseOutput `pulumi:"ruleProperties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewInternetGatewayRule registers a new resource with the given unique name, arguments, and options.
func NewInternetGatewayRule(ctx *pulumi.Context,
	name string, args *InternetGatewayRuleArgs, opts ...pulumi.ResourceOption) (*InternetGatewayRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuleProperties == nil {
		return nil, errors.New("invalid value for required argument 'RuleProperties'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:managednetworkfabric/v20230615:InternetGatewayRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource InternetGatewayRule
	err := ctx.RegisterResource("azure-native:managednetworkfabric:InternetGatewayRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternetGatewayRule gets an existing InternetGatewayRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternetGatewayRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternetGatewayRuleState, opts ...pulumi.ResourceOption) (*InternetGatewayRule, error) {
	var resource InternetGatewayRule
	err := ctx.ReadResource("azure-native:managednetworkfabric:InternetGatewayRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InternetGatewayRule resources.
type internetGatewayRuleState struct {
}

type InternetGatewayRuleState struct {
}

func (InternetGatewayRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*internetGatewayRuleState)(nil)).Elem()
}

type internetGatewayRuleArgs struct {
	// Switch configuration description.
	Annotation *string `pulumi:"annotation"`
	// Name of the Internet Gateway rule.
	InternetGatewayRuleName *string `pulumi:"internetGatewayRuleName"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Rules for the InternetGateways
	RuleProperties RuleProperties `pulumi:"ruleProperties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a InternetGatewayRule resource.
type InternetGatewayRuleArgs struct {
	// Switch configuration description.
	Annotation pulumi.StringPtrInput
	// Name of the Internet Gateway rule.
	InternetGatewayRuleName pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Rules for the InternetGateways
	RuleProperties RulePropertiesInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (InternetGatewayRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internetGatewayRuleArgs)(nil)).Elem()
}

type InternetGatewayRuleInput interface {
	pulumi.Input

	ToInternetGatewayRuleOutput() InternetGatewayRuleOutput
	ToInternetGatewayRuleOutputWithContext(ctx context.Context) InternetGatewayRuleOutput
}

func (*InternetGatewayRule) ElementType() reflect.Type {
	return reflect.TypeOf((**InternetGatewayRule)(nil)).Elem()
}

func (i *InternetGatewayRule) ToInternetGatewayRuleOutput() InternetGatewayRuleOutput {
	return i.ToInternetGatewayRuleOutputWithContext(context.Background())
}

func (i *InternetGatewayRule) ToInternetGatewayRuleOutputWithContext(ctx context.Context) InternetGatewayRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternetGatewayRuleOutput)
}

type InternetGatewayRuleOutput struct{ *pulumi.OutputState }

func (InternetGatewayRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InternetGatewayRule)(nil)).Elem()
}

func (o InternetGatewayRuleOutput) ToInternetGatewayRuleOutput() InternetGatewayRuleOutput {
	return o
}

func (o InternetGatewayRuleOutput) ToInternetGatewayRuleOutputWithContext(ctx context.Context) InternetGatewayRuleOutput {
	return o
}

// Switch configuration description.
func (o InternetGatewayRuleOutput) Annotation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InternetGatewayRule) pulumi.StringPtrOutput { return v.Annotation }).(pulumi.StringPtrOutput)
}

// The Azure API version of the resource.
func (o InternetGatewayRuleOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGatewayRule) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// List of Internet Gateway resource Id.
func (o InternetGatewayRuleOutput) InternetGatewayIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InternetGatewayRule) pulumi.StringArrayOutput { return v.InternetGatewayIds }).(pulumi.StringArrayOutput)
}

// The geo-location where the resource lives
func (o InternetGatewayRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGatewayRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o InternetGatewayRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGatewayRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Provisioning state of the resource.
func (o InternetGatewayRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGatewayRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Rules for the InternetGateways
func (o InternetGatewayRuleOutput) RuleProperties() RulePropertiesResponseOutput {
	return o.ApplyT(func(v *InternetGatewayRule) RulePropertiesResponseOutput { return v.RuleProperties }).(RulePropertiesResponseOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o InternetGatewayRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *InternetGatewayRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o InternetGatewayRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *InternetGatewayRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o InternetGatewayRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InternetGatewayRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(InternetGatewayRuleOutput{})
}
