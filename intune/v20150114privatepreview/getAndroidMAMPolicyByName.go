// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20150114privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Android Policy entity for Intune MAM.
func LookupAndroidMAMPolicyByName(ctx *pulumi.Context, args *LookupAndroidMAMPolicyByNameArgs, opts ...pulumi.InvokeOption) (*LookupAndroidMAMPolicyByNameResult, error) {
	var rv LookupAndroidMAMPolicyByNameResult
	err := ctx.Invoke("azure-native:intune/v20150114privatepreview:getAndroidMAMPolicyByName", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAndroidMAMPolicyByNameArgs struct {
	HostName   string  `pulumi:"hostName"`
	PolicyName string  `pulumi:"policyName"`
	Select     *string `pulumi:"select"`
}

// Android Policy entity for Intune MAM.
type LookupAndroidMAMPolicyByNameResult struct {
	AccessRecheckOfflineTimeout *string           `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  *string           `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         *string           `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           *string           `pulumi:"appSharingToLevel"`
	Authentication              *string           `pulumi:"authentication"`
	ClipboardSharingLevel       *string           `pulumi:"clipboardSharingLevel"`
	DataBackup                  *string           `pulumi:"dataBackup"`
	Description                 *string           `pulumi:"description"`
	DeviceCompliance            *string           `pulumi:"deviceCompliance"`
	FileEncryption              *string           `pulumi:"fileEncryption"`
	FileSharingSaveAs           *string           `pulumi:"fileSharingSaveAs"`
	FriendlyName                string            `pulumi:"friendlyName"`
	GroupStatus                 string            `pulumi:"groupStatus"`
	Id                          string            `pulumi:"id"`
	LastModifiedTime            string            `pulumi:"lastModifiedTime"`
	Location                    *string           `pulumi:"location"`
	ManagedBrowser              *string           `pulumi:"managedBrowser"`
	Name                        string            `pulumi:"name"`
	NumOfApps                   int               `pulumi:"numOfApps"`
	OfflineWipeTimeout          *string           `pulumi:"offlineWipeTimeout"`
	Pin                         *string           `pulumi:"pin"`
	PinNumRetry                 *int              `pulumi:"pinNumRetry"`
	ScreenCapture               *string           `pulumi:"screenCapture"`
	Tags                        map[string]string `pulumi:"tags"`
	Type                        string            `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupAndroidMAMPolicyByNameResult
func (val *LookupAndroidMAMPolicyByNameResult) Defaults() *LookupAndroidMAMPolicyByNameResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AppSharingFromLevel) {
		appSharingFromLevel_ := "none"
		tmp.AppSharingFromLevel = &appSharingFromLevel_
	}
	if isZero(tmp.AppSharingToLevel) {
		appSharingToLevel_ := "none"
		tmp.AppSharingToLevel = &appSharingToLevel_
	}
	if isZero(tmp.Authentication) {
		authentication_ := "required"
		tmp.Authentication = &authentication_
	}
	if isZero(tmp.ClipboardSharingLevel) {
		clipboardSharingLevel_ := "blocked"
		tmp.ClipboardSharingLevel = &clipboardSharingLevel_
	}
	if isZero(tmp.DataBackup) {
		dataBackup_ := "allow"
		tmp.DataBackup = &dataBackup_
	}
	if isZero(tmp.DeviceCompliance) {
		deviceCompliance_ := "enable"
		tmp.DeviceCompliance = &deviceCompliance_
	}
	if isZero(tmp.FileEncryption) {
		fileEncryption_ := "required"
		tmp.FileEncryption = &fileEncryption_
	}
	if isZero(tmp.FileSharingSaveAs) {
		fileSharingSaveAs_ := "allow"
		tmp.FileSharingSaveAs = &fileSharingSaveAs_
	}
	if isZero(tmp.GroupStatus) {
		tmp.GroupStatus = "notTargeted"
	}
	if isZero(tmp.ManagedBrowser) {
		managedBrowser_ := "required"
		tmp.ManagedBrowser = &managedBrowser_
	}
	if isZero(tmp.Pin) {
		pin_ := "required"
		tmp.Pin = &pin_
	}
	if isZero(tmp.ScreenCapture) {
		screenCapture_ := "allow"
		tmp.ScreenCapture = &screenCapture_
	}
	return &tmp
}

func LookupAndroidMAMPolicyByNameOutput(ctx *pulumi.Context, args LookupAndroidMAMPolicyByNameOutputArgs, opts ...pulumi.InvokeOption) LookupAndroidMAMPolicyByNameResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAndroidMAMPolicyByNameResult, error) {
			args := v.(LookupAndroidMAMPolicyByNameArgs)
			r, err := LookupAndroidMAMPolicyByName(ctx, &args, opts...)
			var s LookupAndroidMAMPolicyByNameResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAndroidMAMPolicyByNameResultOutput)
}

type LookupAndroidMAMPolicyByNameOutputArgs struct {
	HostName   pulumi.StringInput    `pulumi:"hostName"`
	PolicyName pulumi.StringInput    `pulumi:"policyName"`
	Select     pulumi.StringPtrInput `pulumi:"select"`
}

func (LookupAndroidMAMPolicyByNameOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAndroidMAMPolicyByNameArgs)(nil)).Elem()
}

// Android Policy entity for Intune MAM.
type LookupAndroidMAMPolicyByNameResultOutput struct{ *pulumi.OutputState }

func (LookupAndroidMAMPolicyByNameResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAndroidMAMPolicyByNameResult)(nil)).Elem()
}

func (o LookupAndroidMAMPolicyByNameResultOutput) ToLookupAndroidMAMPolicyByNameResultOutput() LookupAndroidMAMPolicyByNameResultOutput {
	return o
}

func (o LookupAndroidMAMPolicyByNameResultOutput) ToLookupAndroidMAMPolicyByNameResultOutputWithContext(ctx context.Context) LookupAndroidMAMPolicyByNameResultOutput {
	return o
}

func (o LookupAndroidMAMPolicyByNameResultOutput) AccessRecheckOfflineTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.AccessRecheckOfflineTimeout }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) AccessRecheckOnlineTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.AccessRecheckOnlineTimeout }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) AppSharingFromLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.AppSharingFromLevel }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) AppSharingToLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.AppSharingToLevel }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) Authentication() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.Authentication }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) ClipboardSharingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.ClipboardSharingLevel }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) DataBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.DataBackup }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) DeviceCompliance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.DeviceCompliance }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) FileEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.FileEncryption }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) FileSharingSaveAs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.FileSharingSaveAs }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) string { return v.FriendlyName }).(pulumi.StringOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) GroupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) string { return v.GroupStatus }).(pulumi.StringOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) string { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) ManagedBrowser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.ManagedBrowser }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) NumOfApps() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) int { return v.NumOfApps }).(pulumi.IntOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) OfflineWipeTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.OfflineWipeTimeout }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) Pin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.Pin }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) PinNumRetry() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *int { return v.PinNumRetry }).(pulumi.IntPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) ScreenCapture() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) *string { return v.ScreenCapture }).(pulumi.StringPtrOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupAndroidMAMPolicyByNameResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAndroidMAMPolicyByNameResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAndroidMAMPolicyByNameResultOutput{})
}