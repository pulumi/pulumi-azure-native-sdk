// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datalakeanalytics

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data Lake Analytics firewall rule information.
// API Version: 2016-11-01.
type FirewallRule struct {
	pulumi.CustomResourceState

	// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	EndIpAddress pulumi.StringOutput `pulumi:"endIpAddress"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	StartIpAddress pulumi.StringOutput `pulumi:"startIpAddress"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.EndIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'EndIpAddress'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'StartIpAddress'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20151001preview:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20161101:FirewallRule"),
		},
		{
			Type: pulumi.String("azure-native:datalakeanalytics/v20191101preview:FirewallRule"),
		},
	})
	opts = append(opts, aliases)
	var resource FirewallRule
	err := ctx.RegisterResource("azure-native:datalakeanalytics:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("azure-native:datalakeanalytics:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
}

type FirewallRuleState struct {
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName string `pulumi:"accountName"`
	// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	EndIpAddress string `pulumi:"endIpAddress"`
	// The name of the firewall rule to create or update.
	FirewallRuleName *string `pulumi:"firewallRuleName"`
	// The name of the Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	StartIpAddress string `pulumi:"startIpAddress"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName pulumi.StringInput
	// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	EndIpAddress pulumi.StringInput
	// The name of the firewall rule to create or update.
	FirewallRuleName pulumi.StringPtrInput
	// The name of the Azure resource group.
	ResourceGroupName pulumi.StringInput
	// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
	StartIpAddress pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}

type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput
}

func (*FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (i *FirewallRule) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i *FirewallRule) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}

type FirewallRuleOutput struct{ *pulumi.OutputState }

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

// The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
func (o FirewallRuleOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.EndIpAddress }).(pulumi.StringOutput)
}

// The resource name.
func (o FirewallRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
func (o FirewallRuleOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.StartIpAddress }).(pulumi.StringOutput)
}

// The resource type.
func (o FirewallRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallRuleOutput{})
}