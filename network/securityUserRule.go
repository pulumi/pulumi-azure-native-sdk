// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Network security user rule.
//
// Uses Azure REST API version 2024-05-01. In version 2.x of the Azure Native provider, it used API version 2024-03-01.
//
// Other available API versions: 2024-03-01, 2024-07-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native network [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
type SecurityUserRule struct {
	pulumi.CustomResourceState

	// The Azure API version of the resource.
	AzureApiVersion pulumi.StringOutput `pulumi:"azureApiVersion"`
	// A description for this rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination port ranges.
	DestinationPortRanges pulumi.StringArrayOutput `pulumi:"destinationPortRanges"`
	// The destination address prefixes. CIDR or destination IP ranges.
	Destinations AddressPrefixItemResponseArrayOutput `pulumi:"destinations"`
	// Indicates if the traffic matched against the rule in inbound or outbound.
	Direction pulumi.StringOutput `pulumi:"direction"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network protocol this rule applies to.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The provisioning state of the security configuration user rule resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Unique identifier for this resource.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// The source port ranges.
	SourcePortRanges pulumi.StringArrayOutput `pulumi:"sourcePortRanges"`
	// The CIDR or source IP ranges.
	Sources AddressPrefixItemResponseArrayOutput `pulumi:"sources"`
	// The system metadata related to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSecurityUserRule registers a new resource with the given unique name, arguments, and options.
func NewSecurityUserRule(ctx *pulumi.Context,
	name string, args *SecurityUserRuleArgs, opts ...pulumi.ResourceOption) (*SecurityUserRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'ConfigurationName'")
	}
	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.NetworkManagerName == nil {
		return nil, errors.New("invalid value for required argument 'NetworkManagerName'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RuleCollectionName == nil {
		return nil, errors.New("invalid value for required argument 'RuleCollectionName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20210201preview:SecurityUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:DefaultUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:SecurityUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501preview:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220201preview:SecurityUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:DefaultUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:SecurityUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220401preview:UserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240301:SecurityUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240501:SecurityUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network/v20240701:SecurityUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network:DefaultUserRule"),
		},
		{
			Type: pulumi.String("azure-native:network:UserRule"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource SecurityUserRule
	err := ctx.RegisterResource("azure-native:network:SecurityUserRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityUserRule gets an existing SecurityUserRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityUserRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityUserRuleState, opts ...pulumi.ResourceOption) (*SecurityUserRule, error) {
	var resource SecurityUserRule
	err := ctx.ReadResource("azure-native:network:SecurityUserRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityUserRule resources.
type securityUserRuleState struct {
}

type SecurityUserRuleState struct {
}

func (SecurityUserRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityUserRuleState)(nil)).Elem()
}

type securityUserRuleArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName string `pulumi:"configurationName"`
	// A description for this rule.
	Description *string `pulumi:"description"`
	// The destination port ranges.
	DestinationPortRanges []string `pulumi:"destinationPortRanges"`
	// The destination address prefixes. CIDR or destination IP ranges.
	Destinations []AddressPrefixItem `pulumi:"destinations"`
	// Indicates if the traffic matched against the rule in inbound or outbound.
	Direction string `pulumi:"direction"`
	// The name of the network manager.
	NetworkManagerName string `pulumi:"networkManagerName"`
	// Network protocol this rule applies to.
	Protocol string `pulumi:"protocol"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName string `pulumi:"ruleCollectionName"`
	// The name of the rule.
	RuleName *string `pulumi:"ruleName"`
	// The source port ranges.
	SourcePortRanges []string `pulumi:"sourcePortRanges"`
	// The CIDR or source IP ranges.
	Sources []AddressPrefixItem `pulumi:"sources"`
}

// The set of arguments for constructing a SecurityUserRule resource.
type SecurityUserRuleArgs struct {
	// The name of the network manager Security Configuration.
	ConfigurationName pulumi.StringInput
	// A description for this rule.
	Description pulumi.StringPtrInput
	// The destination port ranges.
	DestinationPortRanges pulumi.StringArrayInput
	// The destination address prefixes. CIDR or destination IP ranges.
	Destinations AddressPrefixItemArrayInput
	// Indicates if the traffic matched against the rule in inbound or outbound.
	Direction pulumi.StringInput
	// The name of the network manager.
	NetworkManagerName pulumi.StringInput
	// Network protocol this rule applies to.
	Protocol pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the network manager security Configuration rule collection.
	RuleCollectionName pulumi.StringInput
	// The name of the rule.
	RuleName pulumi.StringPtrInput
	// The source port ranges.
	SourcePortRanges pulumi.StringArrayInput
	// The CIDR or source IP ranges.
	Sources AddressPrefixItemArrayInput
}

func (SecurityUserRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityUserRuleArgs)(nil)).Elem()
}

type SecurityUserRuleInput interface {
	pulumi.Input

	ToSecurityUserRuleOutput() SecurityUserRuleOutput
	ToSecurityUserRuleOutputWithContext(ctx context.Context) SecurityUserRuleOutput
}

func (*SecurityUserRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityUserRule)(nil)).Elem()
}

func (i *SecurityUserRule) ToSecurityUserRuleOutput() SecurityUserRuleOutput {
	return i.ToSecurityUserRuleOutputWithContext(context.Background())
}

func (i *SecurityUserRule) ToSecurityUserRuleOutputWithContext(ctx context.Context) SecurityUserRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityUserRuleOutput)
}

type SecurityUserRuleOutput struct{ *pulumi.OutputState }

func (SecurityUserRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityUserRule)(nil)).Elem()
}

func (o SecurityUserRuleOutput) ToSecurityUserRuleOutput() SecurityUserRuleOutput {
	return o
}

func (o SecurityUserRuleOutput) ToSecurityUserRuleOutputWithContext(ctx context.Context) SecurityUserRuleOutput {
	return o
}

// The Azure API version of the resource.
func (o SecurityUserRuleOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserRule) pulumi.StringOutput { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// A description for this rule.
func (o SecurityUserRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityUserRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination port ranges.
func (o SecurityUserRuleOutput) DestinationPortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityUserRule) pulumi.StringArrayOutput { return v.DestinationPortRanges }).(pulumi.StringArrayOutput)
}

// The destination address prefixes. CIDR or destination IP ranges.
func (o SecurityUserRuleOutput) Destinations() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v *SecurityUserRule) AddressPrefixItemResponseArrayOutput { return v.Destinations }).(AddressPrefixItemResponseArrayOutput)
}

// Indicates if the traffic matched against the rule in inbound or outbound.
func (o SecurityUserRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// A unique read-only string that changes whenever the resource is updated.
func (o SecurityUserRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Resource name.
func (o SecurityUserRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network protocol this rule applies to.
func (o SecurityUserRuleOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserRule) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// The provisioning state of the security configuration user rule resource.
func (o SecurityUserRuleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserRule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Unique identifier for this resource.
func (o SecurityUserRuleOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserRule) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

// The source port ranges.
func (o SecurityUserRuleOutput) SourcePortRanges() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityUserRule) pulumi.StringArrayOutput { return v.SourcePortRanges }).(pulumi.StringArrayOutput)
}

// The CIDR or source IP ranges.
func (o SecurityUserRuleOutput) Sources() AddressPrefixItemResponseArrayOutput {
	return o.ApplyT(func(v *SecurityUserRule) AddressPrefixItemResponseArrayOutput { return v.Sources }).(AddressPrefixItemResponseArrayOutput)
}

// The system metadata related to this resource.
func (o SecurityUserRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SecurityUserRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource type.
func (o SecurityUserRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityUserRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityUserRuleOutput{})
}
