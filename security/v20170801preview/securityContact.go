// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20170801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contact details for security issues
type SecurityContact struct {
	pulumi.CustomResourceState

	// Whether to send security alerts notifications to the security contact
	AlertNotifications pulumi.StringOutput `pulumi:"alertNotifications"`
	// Whether to send security alerts notifications to subscription admins
	AlertsToAdmins pulumi.StringOutput `pulumi:"alertsToAdmins"`
	// The email of this security contact
	Email pulumi.StringOutput `pulumi:"email"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The phone number of this security contact
	Phone pulumi.StringPtrOutput `pulumi:"phone"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSecurityContact registers a new resource with the given unique name, arguments, and options.
func NewSecurityContact(ctx *pulumi.Context,
	name string, args *SecurityContactArgs, opts ...pulumi.ResourceOption) (*SecurityContact, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertNotifications == nil {
		return nil, errors.New("invalid value for required argument 'AlertNotifications'")
	}
	if args.AlertsToAdmins == nil {
		return nil, errors.New("invalid value for required argument 'AlertsToAdmins'")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:security:SecurityContact"),
		},
		{
			Type: pulumi.String("azure-native:security/v20200101preview:SecurityContact"),
		},
	})
	opts = append(opts, aliases)
	var resource SecurityContact
	err := ctx.RegisterResource("azure-native:security/v20170801preview:SecurityContact", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityContact gets an existing SecurityContact resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityContact(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityContactState, opts ...pulumi.ResourceOption) (*SecurityContact, error) {
	var resource SecurityContact
	err := ctx.ReadResource("azure-native:security/v20170801preview:SecurityContact", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityContact resources.
type securityContactState struct {
}

type SecurityContactState struct {
}

func (SecurityContactState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityContactState)(nil)).Elem()
}

type securityContactArgs struct {
	// Whether to send security alerts notifications to the security contact
	AlertNotifications string `pulumi:"alertNotifications"`
	// Whether to send security alerts notifications to subscription admins
	AlertsToAdmins string `pulumi:"alertsToAdmins"`
	// The email of this security contact
	Email string `pulumi:"email"`
	// The phone number of this security contact
	Phone *string `pulumi:"phone"`
	// Name of the security contact object
	SecurityContactName *string `pulumi:"securityContactName"`
}

// The set of arguments for constructing a SecurityContact resource.
type SecurityContactArgs struct {
	// Whether to send security alerts notifications to the security contact
	AlertNotifications pulumi.StringInput
	// Whether to send security alerts notifications to subscription admins
	AlertsToAdmins pulumi.StringInput
	// The email of this security contact
	Email pulumi.StringInput
	// The phone number of this security contact
	Phone pulumi.StringPtrInput
	// Name of the security contact object
	SecurityContactName pulumi.StringPtrInput
}

func (SecurityContactArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityContactArgs)(nil)).Elem()
}

type SecurityContactInput interface {
	pulumi.Input

	ToSecurityContactOutput() SecurityContactOutput
	ToSecurityContactOutputWithContext(ctx context.Context) SecurityContactOutput
}

func (*SecurityContact) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContact)(nil)).Elem()
}

func (i *SecurityContact) ToSecurityContactOutput() SecurityContactOutput {
	return i.ToSecurityContactOutputWithContext(context.Background())
}

func (i *SecurityContact) ToSecurityContactOutputWithContext(ctx context.Context) SecurityContactOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityContactOutput)
}

type SecurityContactOutput struct{ *pulumi.OutputState }

func (SecurityContactOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityContact)(nil)).Elem()
}

func (o SecurityContactOutput) ToSecurityContactOutput() SecurityContactOutput {
	return o
}

func (o SecurityContactOutput) ToSecurityContactOutputWithContext(ctx context.Context) SecurityContactOutput {
	return o
}

// Whether to send security alerts notifications to the security contact
func (o SecurityContactOutput) AlertNotifications() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityContact) pulumi.StringOutput { return v.AlertNotifications }).(pulumi.StringOutput)
}

// Whether to send security alerts notifications to subscription admins
func (o SecurityContactOutput) AlertsToAdmins() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityContact) pulumi.StringOutput { return v.AlertsToAdmins }).(pulumi.StringOutput)
}

// The email of this security contact
func (o SecurityContactOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityContact) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Resource name
func (o SecurityContactOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityContact) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The phone number of this security contact
func (o SecurityContactOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityContact) pulumi.StringPtrOutput { return v.Phone }).(pulumi.StringPtrOutput)
}

// Resource type
func (o SecurityContactOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityContact) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityContactOutput{})
}