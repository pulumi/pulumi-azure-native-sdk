// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package containerregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The action of virtual network rule.
type Action string

const (
	ActionAllow = Action("Allow")
)

// A message indicating if changes on the service provider require any updates on the consumer.
type ActionsRequired string

const (
	ActionsRequiredNone     = ActionsRequired("None")
	ActionsRequiredRecreate = ActionsRequired("Recreate")
)

// The OS architecture.
type Architecture string

const (
	ArchitectureAmd64 = Architecture("amd64")
	ArchitectureX86   = Architecture("x86")
	Architecture_386  = Architecture("386")
	ArchitectureArm   = Architecture("arm")
	ArchitectureArm64 = Architecture("arm64")
)

// Indicates whether audit logs are enabled on the connected registry.
type AuditLogStatus string

const (
	AuditLogStatusEnabled  = AuditLogStatus("Enabled")
	AuditLogStatusDisabled = AuditLogStatus("Disabled")
)

// The type of the auto trigger for base image dependency updates.
type BaseImageTriggerType string

const (
	BaseImageTriggerTypeAll     = BaseImageTriggerType("All")
	BaseImageTriggerTypeRuntime = BaseImageTriggerType("Runtime")
)

// The mode of the connected registry resource that indicates the permissions of the registry.
type ConnectedRegistryMode string

const (
	ConnectedRegistryModeRegistry = ConnectedRegistryMode("Registry")
	ConnectedRegistryModeMirror   = ConnectedRegistryMode("Mirror")
)

// The private link service connection status.
type ConnectionStatus string

const (
	ConnectionStatusApproved     = ConnectionStatus("Approved")
	ConnectionStatusPending      = ConnectionStatus("Pending")
	ConnectionStatusRejected     = ConnectionStatus("Rejected")
	ConnectionStatusDisconnected = ConnectionStatus("Disconnected")
)

// The default action of allow or deny when no other rules match.
type DefaultAction string

const (
	DefaultActionAllow = DefaultAction("Allow")
	DefaultActionDeny  = DefaultAction("Deny")
)

// The verbosity of logs persisted on the connected registry.
type LogLevel string

const (
	LogLevelDebug       = LogLevel("Debug")
	LogLevelInformation = LogLevel("Information")
	LogLevelWarning     = LogLevel("Warning")
	LogLevelError       = LogLevel("Error")
	LogLevelNone        = LogLevel("None")
)

// The operating system type required for the run.
type OS string

const (
	OSWindows = OS("Windows")
	OSLinux   = OS("Linux")
)

type PipelineOptions string

const (
	PipelineOptionsOverwriteTags             = PipelineOptions("OverwriteTags")
	PipelineOptionsOverwriteBlobs            = PipelineOptions("OverwriteBlobs")
	PipelineOptionsDeleteSourceBlobOnSuccess = PipelineOptions("DeleteSourceBlobOnSuccess")
	PipelineOptionsContinueOnErrors          = PipelineOptions("ContinueOnErrors")
)

// The type of the source.
type PipelineRunSourceType string

const (
	PipelineRunSourceTypeAzureStorageBlob = PipelineRunSourceType("AzureStorageBlob")
)

// The type of the target.
type PipelineRunTargetType string

const (
	PipelineRunTargetTypeAzureStorageBlob = PipelineRunTargetType("AzureStorageBlob")
)

// The type of source for the import pipeline.
type PipelineSourceType string

const (
	PipelineSourceTypeAzureStorageBlobContainer = PipelineSourceType("AzureStorageBlobContainer")
)

// The value that indicates whether the policy is enabled or not.
type PolicyStatus string

const (
	PolicyStatusEnabled  = PolicyStatus("enabled")
	PolicyStatusDisabled = PolicyStatus("disabled")
)

// The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ResourceIdentityTypeInput is an input type that accepts ResourceIdentityTypeArgs and ResourceIdentityTypeOutput values.
// You can construct a concrete instance of `ResourceIdentityTypeInput` via:
//
//	ResourceIdentityTypeArgs{...}
type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

// The type of the secret object which determines how the value of the secret object has to be
// interpreted.
type SecretObjectType string

const (
	SecretObjectTypeOpaque      = SecretObjectType("Opaque")
	SecretObjectTypeVaultsecret = SecretObjectType("Vaultsecret")
)

// The SKU name of the container registry. Required for registry creation.
type SkuName string

const (
	SkuNameClassic  = SkuName("Classic")
	SkuNameBasic    = SkuName("Basic")
	SkuNameStandard = SkuName("Standard")
	SkuNamePremium  = SkuName("Premium")
)

// The type of source control service.
type SourceControlType string

const (
	SourceControlTypeGithub                  = SourceControlType("Github")
	SourceControlTypeVisualStudioTeamService = SourceControlType("VisualStudioTeamService")
)

// The authentication mode which determines the source registry login scope. The credentials for the source registry
// will be generated using the given scope. These credentials will be used to login to
// the source registry during the run.
type SourceRegistryLoginMode string

const (
	SourceRegistryLoginModeNone    = SourceRegistryLoginMode("None")
	SourceRegistryLoginModeDefault = SourceRegistryLoginMode("Default")
)

type SourceTriggerEvent string

const (
	SourceTriggerEventCommit      = SourceTriggerEvent("commit")
	SourceTriggerEventPullrequest = SourceTriggerEvent("pullrequest")
)

// The type of the step.
type StepType string

const (
	StepTypeDocker      = StepType("Docker")
	StepTypeFileTask    = StepType("FileTask")
	StepTypeEncodedTask = StepType("EncodedTask")
)

// The current status of task.
type TaskStatus string

const (
	TaskStatusDisabled = TaskStatus("Disabled")
	TaskStatusEnabled  = TaskStatus("Enabled")
)

type TokenCertificateName string

const (
	TokenCertificateNameCertificate1 = TokenCertificateName("certificate1")
	TokenCertificateNameCertificate2 = TokenCertificateName("certificate2")
)

// The password name "password1" or "password2"
type TokenPasswordName string

const (
	TokenPasswordNamePassword1 = TokenPasswordName("password1")
	TokenPasswordNamePassword2 = TokenPasswordName("password2")
)

// The status of the token example enabled or disabled.
type TokenStatus string

const (
	TokenStatusEnabled  = TokenStatus("enabled")
	TokenStatusDisabled = TokenStatus("disabled")
)

// The type of Auth token.
type TokenType string

const (
	TokenTypePAT   = TokenType("PAT")
	TokenTypeOAuth = TokenType("OAuth")
)

// The current status of trigger.
type TriggerStatus string

const (
	TriggerStatusDisabled = TriggerStatus("Disabled")
	TriggerStatusEnabled  = TriggerStatus("Enabled")
)

// The type of trust policy.
type TrustPolicyType string

const (
	TrustPolicyTypeNotary = TrustPolicyType("Notary")
)

// Type of Payload body for Base image update triggers.
type UpdateTriggerPayloadType string

const (
	UpdateTriggerPayloadTypeDefault = UpdateTriggerPayloadType("Default")
	UpdateTriggerPayloadTypeToken   = UpdateTriggerPayloadType("Token")
)

// Variant of the CPU.
type Variant string

const (
	VariantV6 = Variant("v6")
	VariantV7 = Variant("v7")
	VariantV8 = Variant("v8")
)

type WebhookAction string

const (
	WebhookActionPush          = WebhookAction("push")
	WebhookActionDelete        = WebhookAction("delete")
	WebhookActionQuarantine    = WebhookAction("quarantine")
	WebhookAction_Chart_push   = WebhookAction("chart_push")
	WebhookAction_Chart_delete = WebhookAction("chart_delete")
)

// The status of the webhook at the time the operation was called.
type WebhookStatus string

const (
	WebhookStatusEnabled  = WebhookStatus("enabled")
	WebhookStatusDisabled = WebhookStatus("disabled")
)

func init() {
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}