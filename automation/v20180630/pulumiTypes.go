// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20180630

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the runbook property type.
type ContentHash struct {
	// Gets or sets the content hash algorithm used to hash the content.
	Algorithm string `pulumi:"algorithm"`
	// Gets or sets expected hash value of the content.
	Value string `pulumi:"value"`
}

// ContentHashInput is an input type that accepts ContentHashArgs and ContentHashOutput values.
// You can construct a concrete instance of `ContentHashInput` via:
//
//	ContentHashArgs{...}
type ContentHashInput interface {
	pulumi.Input

	ToContentHashOutput() ContentHashOutput
	ToContentHashOutputWithContext(context.Context) ContentHashOutput
}

// Definition of the runbook property type.
type ContentHashArgs struct {
	// Gets or sets the content hash algorithm used to hash the content.
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	// Gets or sets expected hash value of the content.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ContentHashArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHash)(nil)).Elem()
}

func (i ContentHashArgs) ToContentHashOutput() ContentHashOutput {
	return i.ToContentHashOutputWithContext(context.Background())
}

func (i ContentHashArgs) ToContentHashOutputWithContext(ctx context.Context) ContentHashOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashOutput)
}

func (i ContentHashArgs) ToContentHashPtrOutput() ContentHashPtrOutput {
	return i.ToContentHashPtrOutputWithContext(context.Background())
}

func (i ContentHashArgs) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashOutput).ToContentHashPtrOutputWithContext(ctx)
}

// ContentHashPtrInput is an input type that accepts ContentHashArgs, ContentHashPtr and ContentHashPtrOutput values.
// You can construct a concrete instance of `ContentHashPtrInput` via:
//
//	        ContentHashArgs{...}
//
//	or:
//
//	        nil
type ContentHashPtrInput interface {
	pulumi.Input

	ToContentHashPtrOutput() ContentHashPtrOutput
	ToContentHashPtrOutputWithContext(context.Context) ContentHashPtrOutput
}

type contentHashPtrType ContentHashArgs

func ContentHashPtr(v *ContentHashArgs) ContentHashPtrInput {
	return (*contentHashPtrType)(v)
}

func (*contentHashPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHash)(nil)).Elem()
}

func (i *contentHashPtrType) ToContentHashPtrOutput() ContentHashPtrOutput {
	return i.ToContentHashPtrOutputWithContext(context.Background())
}

func (i *contentHashPtrType) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashPtrOutput)
}

// Definition of the runbook property type.
type ContentHashOutput struct{ *pulumi.OutputState }

func (ContentHashOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHash)(nil)).Elem()
}

func (o ContentHashOutput) ToContentHashOutput() ContentHashOutput {
	return o
}

func (o ContentHashOutput) ToContentHashOutputWithContext(ctx context.Context) ContentHashOutput {
	return o
}

func (o ContentHashOutput) ToContentHashPtrOutput() ContentHashPtrOutput {
	return o.ToContentHashPtrOutputWithContext(context.Background())
}

func (o ContentHashOutput) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentHash) *ContentHash {
		return &v
	}).(ContentHashPtrOutput)
}

// Gets or sets the content hash algorithm used to hash the content.
func (o ContentHashOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHash) string { return v.Algorithm }).(pulumi.StringOutput)
}

// Gets or sets expected hash value of the content.
func (o ContentHashOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHash) string { return v.Value }).(pulumi.StringOutput)
}

type ContentHashPtrOutput struct{ *pulumi.OutputState }

func (ContentHashPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHash)(nil)).Elem()
}

func (o ContentHashPtrOutput) ToContentHashPtrOutput() ContentHashPtrOutput {
	return o
}

func (o ContentHashPtrOutput) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return o
}

func (o ContentHashPtrOutput) Elem() ContentHashOutput {
	return o.ApplyT(func(v *ContentHash) ContentHash {
		if v != nil {
			return *v
		}
		var ret ContentHash
		return ret
	}).(ContentHashOutput)
}

// Gets or sets the content hash algorithm used to hash the content.
func (o ContentHashPtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHash) *string {
		if v == nil {
			return nil
		}
		return &v.Algorithm
	}).(pulumi.StringPtrOutput)
}

// Gets or sets expected hash value of the content.
func (o ContentHashPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHash) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

// Definition of the runbook property type.
type ContentHashResponse struct {
	// Gets or sets the content hash algorithm used to hash the content.
	Algorithm string `pulumi:"algorithm"`
	// Gets or sets expected hash value of the content.
	Value string `pulumi:"value"`
}

// Definition of the runbook property type.
type ContentHashResponseOutput struct{ *pulumi.OutputState }

func (ContentHashResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHashResponse)(nil)).Elem()
}

func (o ContentHashResponseOutput) ToContentHashResponseOutput() ContentHashResponseOutput {
	return o
}

func (o ContentHashResponseOutput) ToContentHashResponseOutputWithContext(ctx context.Context) ContentHashResponseOutput {
	return o
}

// Gets or sets the content hash algorithm used to hash the content.
func (o ContentHashResponseOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHashResponse) string { return v.Algorithm }).(pulumi.StringOutput)
}

// Gets or sets expected hash value of the content.
func (o ContentHashResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHashResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ContentHashResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentHashResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHashResponse)(nil)).Elem()
}

func (o ContentHashResponsePtrOutput) ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput {
	return o
}

func (o ContentHashResponsePtrOutput) ToContentHashResponsePtrOutputWithContext(ctx context.Context) ContentHashResponsePtrOutput {
	return o
}

func (o ContentHashResponsePtrOutput) Elem() ContentHashResponseOutput {
	return o.ApplyT(func(v *ContentHashResponse) ContentHashResponse {
		if v != nil {
			return *v
		}
		var ret ContentHashResponse
		return ret
	}).(ContentHashResponseOutput)
}

// Gets or sets the content hash algorithm used to hash the content.
func (o ContentHashResponsePtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Algorithm
	}).(pulumi.StringPtrOutput)
}

// Gets or sets expected hash value of the content.
func (o ContentHashResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

// Definition of the content link.
type ContentLink struct {
	// Gets or sets the hash.
	ContentHash *ContentHash `pulumi:"contentHash"`
	// Gets or sets the uri of the runbook content.
	Uri *string `pulumi:"uri"`
	// Gets or sets the version of the content.
	Version *string `pulumi:"version"`
}

// ContentLinkInput is an input type that accepts ContentLinkArgs and ContentLinkOutput values.
// You can construct a concrete instance of `ContentLinkInput` via:
//
//	ContentLinkArgs{...}
type ContentLinkInput interface {
	pulumi.Input

	ToContentLinkOutput() ContentLinkOutput
	ToContentLinkOutputWithContext(context.Context) ContentLinkOutput
}

// Definition of the content link.
type ContentLinkArgs struct {
	// Gets or sets the hash.
	ContentHash ContentHashPtrInput `pulumi:"contentHash"`
	// Gets or sets the uri of the runbook content.
	Uri pulumi.StringPtrInput `pulumi:"uri"`
	// Gets or sets the version of the content.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (ContentLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLink)(nil)).Elem()
}

func (i ContentLinkArgs) ToContentLinkOutput() ContentLinkOutput {
	return i.ToContentLinkOutputWithContext(context.Background())
}

func (i ContentLinkArgs) ToContentLinkOutputWithContext(ctx context.Context) ContentLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkOutput)
}

func (i ContentLinkArgs) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return i.ToContentLinkPtrOutputWithContext(context.Background())
}

func (i ContentLinkArgs) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkOutput).ToContentLinkPtrOutputWithContext(ctx)
}

// ContentLinkPtrInput is an input type that accepts ContentLinkArgs, ContentLinkPtr and ContentLinkPtrOutput values.
// You can construct a concrete instance of `ContentLinkPtrInput` via:
//
//	        ContentLinkArgs{...}
//
//	or:
//
//	        nil
type ContentLinkPtrInput interface {
	pulumi.Input

	ToContentLinkPtrOutput() ContentLinkPtrOutput
	ToContentLinkPtrOutputWithContext(context.Context) ContentLinkPtrOutput
}

type contentLinkPtrType ContentLinkArgs

func ContentLinkPtr(v *ContentLinkArgs) ContentLinkPtrInput {
	return (*contentLinkPtrType)(v)
}

func (*contentLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLink)(nil)).Elem()
}

func (i *contentLinkPtrType) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return i.ToContentLinkPtrOutputWithContext(context.Background())
}

func (i *contentLinkPtrType) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkPtrOutput)
}

// Definition of the content link.
type ContentLinkOutput struct{ *pulumi.OutputState }

func (ContentLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLink)(nil)).Elem()
}

func (o ContentLinkOutput) ToContentLinkOutput() ContentLinkOutput {
	return o
}

func (o ContentLinkOutput) ToContentLinkOutputWithContext(ctx context.Context) ContentLinkOutput {
	return o
}

func (o ContentLinkOutput) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return o.ToContentLinkPtrOutputWithContext(context.Background())
}

func (o ContentLinkOutput) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentLink) *ContentLink {
		return &v
	}).(ContentLinkPtrOutput)
}

// Gets or sets the hash.
func (o ContentLinkOutput) ContentHash() ContentHashPtrOutput {
	return o.ApplyT(func(v ContentLink) *ContentHash { return v.ContentHash }).(ContentHashPtrOutput)
}

// Gets or sets the uri of the runbook content.
func (o ContentLinkOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLink) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

// Gets or sets the version of the content.
func (o ContentLinkOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLink) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ContentLinkPtrOutput struct{ *pulumi.OutputState }

func (ContentLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLink)(nil)).Elem()
}

func (o ContentLinkPtrOutput) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return o
}

func (o ContentLinkPtrOutput) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return o
}

func (o ContentLinkPtrOutput) Elem() ContentLinkOutput {
	return o.ApplyT(func(v *ContentLink) ContentLink {
		if v != nil {
			return *v
		}
		var ret ContentLink
		return ret
	}).(ContentLinkOutput)
}

// Gets or sets the hash.
func (o ContentLinkPtrOutput) ContentHash() ContentHashPtrOutput {
	return o.ApplyT(func(v *ContentLink) *ContentHash {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(ContentHashPtrOutput)
}

// Gets or sets the uri of the runbook content.
func (o ContentLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLink) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

// Gets or sets the version of the content.
func (o ContentLinkPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLink) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

// Definition of the content link.
type ContentLinkResponse struct {
	// Gets or sets the hash.
	ContentHash *ContentHashResponse `pulumi:"contentHash"`
	// Gets or sets the uri of the runbook content.
	Uri *string `pulumi:"uri"`
	// Gets or sets the version of the content.
	Version *string `pulumi:"version"`
}

// Definition of the content link.
type ContentLinkResponseOutput struct{ *pulumi.OutputState }

func (ContentLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLinkResponse)(nil)).Elem()
}

func (o ContentLinkResponseOutput) ToContentLinkResponseOutput() ContentLinkResponseOutput {
	return o
}

func (o ContentLinkResponseOutput) ToContentLinkResponseOutputWithContext(ctx context.Context) ContentLinkResponseOutput {
	return o
}

// Gets or sets the hash.
func (o ContentLinkResponseOutput) ContentHash() ContentHashResponsePtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *ContentHashResponse { return v.ContentHash }).(ContentHashResponsePtrOutput)
}

// Gets or sets the uri of the runbook content.
func (o ContentLinkResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

// Gets or sets the version of the content.
func (o ContentLinkResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ContentLinkResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentLinkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLinkResponse)(nil)).Elem()
}

func (o ContentLinkResponsePtrOutput) ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput {
	return o
}

func (o ContentLinkResponsePtrOutput) ToContentLinkResponsePtrOutputWithContext(ctx context.Context) ContentLinkResponsePtrOutput {
	return o
}

func (o ContentLinkResponsePtrOutput) Elem() ContentLinkResponseOutput {
	return o.ApplyT(func(v *ContentLinkResponse) ContentLinkResponse {
		if v != nil {
			return *v
		}
		var ret ContentLinkResponse
		return ret
	}).(ContentLinkResponseOutput)
}

// Gets or sets the hash.
func (o ContentLinkResponsePtrOutput) ContentHash() ContentHashResponsePtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *ContentHashResponse {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(ContentHashResponsePtrOutput)
}

// Gets or sets the uri of the runbook content.
func (o ContentLinkResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

// Gets or sets the version of the content.
func (o ContentLinkResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

// Definition of the module error info type.
type ModuleErrorInfoResponse struct {
	// Gets or sets the error code.
	Code *string `pulumi:"code"`
	// Gets or sets the error message.
	Message *string `pulumi:"message"`
}

// Definition of the module error info type.
type ModuleErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (ModuleErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleErrorInfoResponse)(nil)).Elem()
}

func (o ModuleErrorInfoResponseOutput) ToModuleErrorInfoResponseOutput() ModuleErrorInfoResponseOutput {
	return o
}

func (o ModuleErrorInfoResponseOutput) ToModuleErrorInfoResponseOutputWithContext(ctx context.Context) ModuleErrorInfoResponseOutput {
	return o
}

// Gets or sets the error code.
func (o ModuleErrorInfoResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModuleErrorInfoResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

// Gets or sets the error message.
func (o ModuleErrorInfoResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModuleErrorInfoResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ModuleErrorInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ModuleErrorInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleErrorInfoResponse)(nil)).Elem()
}

func (o ModuleErrorInfoResponsePtrOutput) ToModuleErrorInfoResponsePtrOutput() ModuleErrorInfoResponsePtrOutput {
	return o
}

func (o ModuleErrorInfoResponsePtrOutput) ToModuleErrorInfoResponsePtrOutputWithContext(ctx context.Context) ModuleErrorInfoResponsePtrOutput {
	return o
}

func (o ModuleErrorInfoResponsePtrOutput) Elem() ModuleErrorInfoResponseOutput {
	return o.ApplyT(func(v *ModuleErrorInfoResponse) ModuleErrorInfoResponse {
		if v != nil {
			return *v
		}
		var ret ModuleErrorInfoResponse
		return ret
	}).(ModuleErrorInfoResponseOutput)
}

// Gets or sets the error code.
func (o ModuleErrorInfoResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModuleErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

// Gets or sets the error message.
func (o ModuleErrorInfoResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModuleErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type RunbookDraft struct {
	// Gets or sets the creation time of the runbook draft.
	CreationTime *string `pulumi:"creationTime"`
	// Gets or sets the draft runbook content link.
	DraftContentLink *ContentLink `pulumi:"draftContentLink"`
	// Gets or sets whether runbook is in edit mode.
	InEdit *bool `pulumi:"inEdit"`
	// Gets or sets the last modified time of the runbook draft.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// Gets or sets the runbook output types.
	OutputTypes []string `pulumi:"outputTypes"`
	// Gets or sets the runbook draft parameters.
	Parameters map[string]RunbookParameter `pulumi:"parameters"`
}

// RunbookDraftInput is an input type that accepts RunbookDraftArgs and RunbookDraftOutput values.
// You can construct a concrete instance of `RunbookDraftInput` via:
//
//	RunbookDraftArgs{...}
type RunbookDraftInput interface {
	pulumi.Input

	ToRunbookDraftOutput() RunbookDraftOutput
	ToRunbookDraftOutputWithContext(context.Context) RunbookDraftOutput
}

type RunbookDraftArgs struct {
	// Gets or sets the creation time of the runbook draft.
	CreationTime pulumi.StringPtrInput `pulumi:"creationTime"`
	// Gets or sets the draft runbook content link.
	DraftContentLink ContentLinkPtrInput `pulumi:"draftContentLink"`
	// Gets or sets whether runbook is in edit mode.
	InEdit pulumi.BoolPtrInput `pulumi:"inEdit"`
	// Gets or sets the last modified time of the runbook draft.
	LastModifiedTime pulumi.StringPtrInput `pulumi:"lastModifiedTime"`
	// Gets or sets the runbook output types.
	OutputTypes pulumi.StringArrayInput `pulumi:"outputTypes"`
	// Gets or sets the runbook draft parameters.
	Parameters RunbookParameterMapInput `pulumi:"parameters"`
}

func (RunbookDraftArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookDraft)(nil)).Elem()
}

func (i RunbookDraftArgs) ToRunbookDraftOutput() RunbookDraftOutput {
	return i.ToRunbookDraftOutputWithContext(context.Background())
}

func (i RunbookDraftArgs) ToRunbookDraftOutputWithContext(ctx context.Context) RunbookDraftOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftOutput)
}

func (i RunbookDraftArgs) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return i.ToRunbookDraftPtrOutputWithContext(context.Background())
}

func (i RunbookDraftArgs) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftOutput).ToRunbookDraftPtrOutputWithContext(ctx)
}

// RunbookDraftPtrInput is an input type that accepts RunbookDraftArgs, RunbookDraftPtr and RunbookDraftPtrOutput values.
// You can construct a concrete instance of `RunbookDraftPtrInput` via:
//
//	        RunbookDraftArgs{...}
//
//	or:
//
//	        nil
type RunbookDraftPtrInput interface {
	pulumi.Input

	ToRunbookDraftPtrOutput() RunbookDraftPtrOutput
	ToRunbookDraftPtrOutputWithContext(context.Context) RunbookDraftPtrOutput
}

type runbookDraftPtrType RunbookDraftArgs

func RunbookDraftPtr(v *RunbookDraftArgs) RunbookDraftPtrInput {
	return (*runbookDraftPtrType)(v)
}

func (*runbookDraftPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookDraft)(nil)).Elem()
}

func (i *runbookDraftPtrType) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return i.ToRunbookDraftPtrOutputWithContext(context.Background())
}

func (i *runbookDraftPtrType) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftPtrOutput)
}

type RunbookDraftOutput struct{ *pulumi.OutputState }

func (RunbookDraftOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookDraft)(nil)).Elem()
}

func (o RunbookDraftOutput) ToRunbookDraftOutput() RunbookDraftOutput {
	return o
}

func (o RunbookDraftOutput) ToRunbookDraftOutputWithContext(ctx context.Context) RunbookDraftOutput {
	return o
}

func (o RunbookDraftOutput) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return o.ToRunbookDraftPtrOutputWithContext(context.Background())
}

func (o RunbookDraftOutput) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunbookDraft) *RunbookDraft {
		return &v
	}).(RunbookDraftPtrOutput)
}

// Gets or sets the creation time of the runbook draft.
func (o RunbookDraftOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the draft runbook content link.
func (o RunbookDraftOutput) DraftContentLink() ContentLinkPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *ContentLink { return v.DraftContentLink }).(ContentLinkPtrOutput)
}

// Gets or sets whether runbook is in edit mode.
func (o RunbookDraftOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *bool { return v.InEdit }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time of the runbook draft.
func (o RunbookDraftOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the runbook output types.
func (o RunbookDraftOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RunbookDraft) []string { return v.OutputTypes }).(pulumi.StringArrayOutput)
}

// Gets or sets the runbook draft parameters.
func (o RunbookDraftOutput) Parameters() RunbookParameterMapOutput {
	return o.ApplyT(func(v RunbookDraft) map[string]RunbookParameter { return v.Parameters }).(RunbookParameterMapOutput)
}

type RunbookDraftPtrOutput struct{ *pulumi.OutputState }

func (RunbookDraftPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookDraft)(nil)).Elem()
}

func (o RunbookDraftPtrOutput) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return o
}

func (o RunbookDraftPtrOutput) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return o
}

func (o RunbookDraftPtrOutput) Elem() RunbookDraftOutput {
	return o.ApplyT(func(v *RunbookDraft) RunbookDraft {
		if v != nil {
			return *v
		}
		var ret RunbookDraft
		return ret
	}).(RunbookDraftOutput)
}

// Gets or sets the creation time of the runbook draft.
func (o RunbookDraftPtrOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *string {
		if v == nil {
			return nil
		}
		return v.CreationTime
	}).(pulumi.StringPtrOutput)
}

// Gets or sets the draft runbook content link.
func (o RunbookDraftPtrOutput) DraftContentLink() ContentLinkPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *ContentLink {
		if v == nil {
			return nil
		}
		return v.DraftContentLink
	}).(ContentLinkPtrOutput)
}

// Gets or sets whether runbook is in edit mode.
func (o RunbookDraftPtrOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *bool {
		if v == nil {
			return nil
		}
		return v.InEdit
	}).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time of the runbook draft.
func (o RunbookDraftPtrOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedTime
	}).(pulumi.StringPtrOutput)
}

// Gets or sets the runbook output types.
func (o RunbookDraftPtrOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RunbookDraft) []string {
		if v == nil {
			return nil
		}
		return v.OutputTypes
	}).(pulumi.StringArrayOutput)
}

// Gets or sets the runbook draft parameters.
func (o RunbookDraftPtrOutput) Parameters() RunbookParameterMapOutput {
	return o.ApplyT(func(v *RunbookDraft) map[string]RunbookParameter {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(RunbookParameterMapOutput)
}

type RunbookDraftResponse struct {
	// Gets or sets the creation time of the runbook draft.
	CreationTime *string `pulumi:"creationTime"`
	// Gets or sets the draft runbook content link.
	DraftContentLink *ContentLinkResponse `pulumi:"draftContentLink"`
	// Gets or sets whether runbook is in edit mode.
	InEdit *bool `pulumi:"inEdit"`
	// Gets or sets the last modified time of the runbook draft.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// Gets or sets the runbook output types.
	OutputTypes []string `pulumi:"outputTypes"`
	// Gets or sets the runbook draft parameters.
	Parameters map[string]RunbookParameterResponse `pulumi:"parameters"`
}

type RunbookDraftResponseOutput struct{ *pulumi.OutputState }

func (RunbookDraftResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookDraftResponse)(nil)).Elem()
}

func (o RunbookDraftResponseOutput) ToRunbookDraftResponseOutput() RunbookDraftResponseOutput {
	return o
}

func (o RunbookDraftResponseOutput) ToRunbookDraftResponseOutputWithContext(ctx context.Context) RunbookDraftResponseOutput {
	return o
}

// Gets or sets the creation time of the runbook draft.
func (o RunbookDraftResponseOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the draft runbook content link.
func (o RunbookDraftResponseOutput) DraftContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *ContentLinkResponse { return v.DraftContentLink }).(ContentLinkResponsePtrOutput)
}

// Gets or sets whether runbook is in edit mode.
func (o RunbookDraftResponseOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *bool { return v.InEdit }).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time of the runbook draft.
func (o RunbookDraftResponseOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

// Gets or sets the runbook output types.
func (o RunbookDraftResponseOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RunbookDraftResponse) []string { return v.OutputTypes }).(pulumi.StringArrayOutput)
}

// Gets or sets the runbook draft parameters.
func (o RunbookDraftResponseOutput) Parameters() RunbookParameterResponseMapOutput {
	return o.ApplyT(func(v RunbookDraftResponse) map[string]RunbookParameterResponse { return v.Parameters }).(RunbookParameterResponseMapOutput)
}

type RunbookDraftResponsePtrOutput struct{ *pulumi.OutputState }

func (RunbookDraftResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookDraftResponse)(nil)).Elem()
}

func (o RunbookDraftResponsePtrOutput) ToRunbookDraftResponsePtrOutput() RunbookDraftResponsePtrOutput {
	return o
}

func (o RunbookDraftResponsePtrOutput) ToRunbookDraftResponsePtrOutputWithContext(ctx context.Context) RunbookDraftResponsePtrOutput {
	return o
}

func (o RunbookDraftResponsePtrOutput) Elem() RunbookDraftResponseOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) RunbookDraftResponse {
		if v != nil {
			return *v
		}
		var ret RunbookDraftResponse
		return ret
	}).(RunbookDraftResponseOutput)
}

// Gets or sets the creation time of the runbook draft.
func (o RunbookDraftResponsePtrOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreationTime
	}).(pulumi.StringPtrOutput)
}

// Gets or sets the draft runbook content link.
func (o RunbookDraftResponsePtrOutput) DraftContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *ContentLinkResponse {
		if v == nil {
			return nil
		}
		return v.DraftContentLink
	}).(ContentLinkResponsePtrOutput)
}

// Gets or sets whether runbook is in edit mode.
func (o RunbookDraftResponsePtrOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *bool {
		if v == nil {
			return nil
		}
		return v.InEdit
	}).(pulumi.BoolPtrOutput)
}

// Gets or sets the last modified time of the runbook draft.
func (o RunbookDraftResponsePtrOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedTime
	}).(pulumi.StringPtrOutput)
}

// Gets or sets the runbook output types.
func (o RunbookDraftResponsePtrOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) []string {
		if v == nil {
			return nil
		}
		return v.OutputTypes
	}).(pulumi.StringArrayOutput)
}

// Gets or sets the runbook draft parameters.
func (o RunbookDraftResponsePtrOutput) Parameters() RunbookParameterResponseMapOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) map[string]RunbookParameterResponse {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(RunbookParameterResponseMapOutput)
}

// Definition of the runbook parameter type.
type RunbookParameter struct {
	// Gets or sets the default value of parameter.
	DefaultValue *string `pulumi:"defaultValue"`
	// Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
	IsMandatory *bool `pulumi:"isMandatory"`
	// Get or sets the position of the parameter.
	Position *int `pulumi:"position"`
	// Gets or sets the type of the parameter.
	Type *string `pulumi:"type"`
}

// RunbookParameterInput is an input type that accepts RunbookParameterArgs and RunbookParameterOutput values.
// You can construct a concrete instance of `RunbookParameterInput` via:
//
//	RunbookParameterArgs{...}
type RunbookParameterInput interface {
	pulumi.Input

	ToRunbookParameterOutput() RunbookParameterOutput
	ToRunbookParameterOutputWithContext(context.Context) RunbookParameterOutput
}

// Definition of the runbook parameter type.
type RunbookParameterArgs struct {
	// Gets or sets the default value of parameter.
	DefaultValue pulumi.StringPtrInput `pulumi:"defaultValue"`
	// Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
	IsMandatory pulumi.BoolPtrInput `pulumi:"isMandatory"`
	// Get or sets the position of the parameter.
	Position pulumi.IntPtrInput `pulumi:"position"`
	// Gets or sets the type of the parameter.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (RunbookParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookParameter)(nil)).Elem()
}

func (i RunbookParameterArgs) ToRunbookParameterOutput() RunbookParameterOutput {
	return i.ToRunbookParameterOutputWithContext(context.Background())
}

func (i RunbookParameterArgs) ToRunbookParameterOutputWithContext(ctx context.Context) RunbookParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookParameterOutput)
}

// RunbookParameterMapInput is an input type that accepts RunbookParameterMap and RunbookParameterMapOutput values.
// You can construct a concrete instance of `RunbookParameterMapInput` via:
//
//	RunbookParameterMap{ "key": RunbookParameterArgs{...} }
type RunbookParameterMapInput interface {
	pulumi.Input

	ToRunbookParameterMapOutput() RunbookParameterMapOutput
	ToRunbookParameterMapOutputWithContext(context.Context) RunbookParameterMapOutput
}

type RunbookParameterMap map[string]RunbookParameterInput

func (RunbookParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunbookParameter)(nil)).Elem()
}

func (i RunbookParameterMap) ToRunbookParameterMapOutput() RunbookParameterMapOutput {
	return i.ToRunbookParameterMapOutputWithContext(context.Background())
}

func (i RunbookParameterMap) ToRunbookParameterMapOutputWithContext(ctx context.Context) RunbookParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookParameterMapOutput)
}

// Definition of the runbook parameter type.
type RunbookParameterOutput struct{ *pulumi.OutputState }

func (RunbookParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookParameter)(nil)).Elem()
}

func (o RunbookParameterOutput) ToRunbookParameterOutput() RunbookParameterOutput {
	return o
}

func (o RunbookParameterOutput) ToRunbookParameterOutputWithContext(ctx context.Context) RunbookParameterOutput {
	return o
}

// Gets or sets the default value of parameter.
func (o RunbookParameterOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

// Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
func (o RunbookParameterOutput) IsMandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *bool { return v.IsMandatory }).(pulumi.BoolPtrOutput)
}

// Get or sets the position of the parameter.
func (o RunbookParameterOutput) Position() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *int { return v.Position }).(pulumi.IntPtrOutput)
}

// Gets or sets the type of the parameter.
func (o RunbookParameterOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type RunbookParameterMapOutput struct{ *pulumi.OutputState }

func (RunbookParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunbookParameter)(nil)).Elem()
}

func (o RunbookParameterMapOutput) ToRunbookParameterMapOutput() RunbookParameterMapOutput {
	return o
}

func (o RunbookParameterMapOutput) ToRunbookParameterMapOutputWithContext(ctx context.Context) RunbookParameterMapOutput {
	return o
}

func (o RunbookParameterMapOutput) MapIndex(k pulumi.StringInput) RunbookParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RunbookParameter {
		return vs[0].(map[string]RunbookParameter)[vs[1].(string)]
	}).(RunbookParameterOutput)
}

// Definition of the runbook parameter type.
type RunbookParameterResponse struct {
	// Gets or sets the default value of parameter.
	DefaultValue *string `pulumi:"defaultValue"`
	// Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
	IsMandatory *bool `pulumi:"isMandatory"`
	// Get or sets the position of the parameter.
	Position *int `pulumi:"position"`
	// Gets or sets the type of the parameter.
	Type *string `pulumi:"type"`
}

// Definition of the runbook parameter type.
type RunbookParameterResponseOutput struct{ *pulumi.OutputState }

func (RunbookParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookParameterResponse)(nil)).Elem()
}

func (o RunbookParameterResponseOutput) ToRunbookParameterResponseOutput() RunbookParameterResponseOutput {
	return o
}

func (o RunbookParameterResponseOutput) ToRunbookParameterResponseOutputWithContext(ctx context.Context) RunbookParameterResponseOutput {
	return o
}

// Gets or sets the default value of parameter.
func (o RunbookParameterResponseOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

// Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
func (o RunbookParameterResponseOutput) IsMandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *bool { return v.IsMandatory }).(pulumi.BoolPtrOutput)
}

// Get or sets the position of the parameter.
func (o RunbookParameterResponseOutput) Position() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *int { return v.Position }).(pulumi.IntPtrOutput)
}

// Gets or sets the type of the parameter.
func (o RunbookParameterResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type RunbookParameterResponseMapOutput struct{ *pulumi.OutputState }

func (RunbookParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunbookParameterResponse)(nil)).Elem()
}

func (o RunbookParameterResponseMapOutput) ToRunbookParameterResponseMapOutput() RunbookParameterResponseMapOutput {
	return o
}

func (o RunbookParameterResponseMapOutput) ToRunbookParameterResponseMapOutputWithContext(ctx context.Context) RunbookParameterResponseMapOutput {
	return o
}

func (o RunbookParameterResponseMapOutput) MapIndex(k pulumi.StringInput) RunbookParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RunbookParameterResponse {
		return vs[0].(map[string]RunbookParameterResponse)[vs[1].(string)]
	}).(RunbookParameterResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ContentHashOutput{})
	pulumi.RegisterOutputType(ContentHashPtrOutput{})
	pulumi.RegisterOutputType(ContentHashResponseOutput{})
	pulumi.RegisterOutputType(ContentHashResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentLinkOutput{})
	pulumi.RegisterOutputType(ContentLinkPtrOutput{})
	pulumi.RegisterOutputType(ContentLinkResponseOutput{})
	pulumi.RegisterOutputType(ContentLinkResponsePtrOutput{})
	pulumi.RegisterOutputType(ModuleErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(ModuleErrorInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(RunbookDraftOutput{})
	pulumi.RegisterOutputType(RunbookDraftPtrOutput{})
	pulumi.RegisterOutputType(RunbookDraftResponseOutput{})
	pulumi.RegisterOutputType(RunbookDraftResponsePtrOutput{})
	pulumi.RegisterOutputType(RunbookParameterOutput{})
	pulumi.RegisterOutputType(RunbookParameterMapOutput{})
	pulumi.RegisterOutputType(RunbookParameterResponseOutput{})
	pulumi.RegisterOutputType(RunbookParameterResponseMapOutput{})
}