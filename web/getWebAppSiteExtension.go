// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v3/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Description for Get site extension information by its ID for a web site, or a deployment slot.
//
// Uses Azure REST API version 2024-04-01.
//
// Other available API versions: 2016-08-01, 2018-02-01, 2018-11-01, 2019-08-01, 2020-06-01, 2020-09-01, 2020-10-01, 2020-12-01, 2021-01-01, 2021-01-15, 2021-02-01, 2021-03-01, 2022-03-01, 2022-09-01, 2023-01-01, 2023-12-01. These can be accessed by generating a local SDK package using the CLI command `pulumi package add azure-native web [ApiVersion]`. See the [version guide](../../../version-guide/#accessing-any-api-version-via-local-packages) for details.
func LookupWebAppSiteExtension(ctx *pulumi.Context, args *LookupWebAppSiteExtensionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSiteExtensionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWebAppSiteExtensionResult
	err := ctx.Invoke("azure-native:web:getWebAppSiteExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSiteExtensionArgs struct {
	// Site name.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Site extension name.
	SiteExtensionId string `pulumi:"siteExtensionId"`
}

// Site Extension Information.
type LookupWebAppSiteExtensionResult struct {
	// List of authors.
	Authors []string `pulumi:"authors"`
	// The Azure API version of the resource.
	AzureApiVersion string `pulumi:"azureApiVersion"`
	// Site Extension comment.
	Comment *string `pulumi:"comment"`
	// Detailed description.
	Description *string `pulumi:"description"`
	// Count of downloads.
	DownloadCount *int `pulumi:"downloadCount"`
	// Site extension ID.
	ExtensionId *string `pulumi:"extensionId"`
	// Site extension type.
	ExtensionType *string `pulumi:"extensionType"`
	// Extension URL.
	ExtensionUrl *string `pulumi:"extensionUrl"`
	// Feed URL.
	FeedUrl *string `pulumi:"feedUrl"`
	// Icon URL.
	IconUrl *string `pulumi:"iconUrl"`
	// Resource Id.
	Id string `pulumi:"id"`
	// Installed timestamp.
	InstalledDateTime *string `pulumi:"installedDateTime"`
	// Installer command line parameters.
	InstallerCommandLineParams *string `pulumi:"installerCommandLineParams"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// License URL.
	LicenseUrl *string `pulumi:"licenseUrl"`
	// <code>true</code> if the local version is the latest version; <code>false</code> otherwise.
	LocalIsLatestVersion *bool `pulumi:"localIsLatestVersion"`
	// Local path.
	LocalPath *string `pulumi:"localPath"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Project URL.
	ProjectUrl *string `pulumi:"projectUrl"`
	// Provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Published timestamp.
	PublishedDateTime *string `pulumi:"publishedDateTime"`
	// Summary description.
	Summary *string `pulumi:"summary"`
	Title   *string `pulumi:"title"`
	// Resource type.
	Type string `pulumi:"type"`
	// Version information.
	Version *string `pulumi:"version"`
}

func LookupWebAppSiteExtensionOutput(ctx *pulumi.Context, args LookupWebAppSiteExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppSiteExtensionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupWebAppSiteExtensionResultOutput, error) {
			args := v.(LookupWebAppSiteExtensionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: utilities.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("azure-native:web:getWebAppSiteExtension", args, LookupWebAppSiteExtensionResultOutput{}, options).(LookupWebAppSiteExtensionResultOutput), nil
		}).(LookupWebAppSiteExtensionResultOutput)
}

type LookupWebAppSiteExtensionOutputArgs struct {
	// Site name.
	Name pulumi.StringInput `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Site extension name.
	SiteExtensionId pulumi.StringInput `pulumi:"siteExtensionId"`
}

func (LookupWebAppSiteExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSiteExtensionArgs)(nil)).Elem()
}

// Site Extension Information.
type LookupWebAppSiteExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppSiteExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppSiteExtensionResult)(nil)).Elem()
}

func (o LookupWebAppSiteExtensionResultOutput) ToLookupWebAppSiteExtensionResultOutput() LookupWebAppSiteExtensionResultOutput {
	return o
}

func (o LookupWebAppSiteExtensionResultOutput) ToLookupWebAppSiteExtensionResultOutputWithContext(ctx context.Context) LookupWebAppSiteExtensionResultOutput {
	return o
}

// List of authors.
func (o LookupWebAppSiteExtensionResultOutput) Authors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) []string { return v.Authors }).(pulumi.StringArrayOutput)
}

// The Azure API version of the resource.
func (o LookupWebAppSiteExtensionResultOutput) AzureApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) string { return v.AzureApiVersion }).(pulumi.StringOutput)
}

// Site Extension comment.
func (o LookupWebAppSiteExtensionResultOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Comment }).(pulumi.StringPtrOutput)
}

// Detailed description.
func (o LookupWebAppSiteExtensionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Count of downloads.
func (o LookupWebAppSiteExtensionResultOutput) DownloadCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *int { return v.DownloadCount }).(pulumi.IntPtrOutput)
}

// Site extension ID.
func (o LookupWebAppSiteExtensionResultOutput) ExtensionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ExtensionId }).(pulumi.StringPtrOutput)
}

// Site extension type.
func (o LookupWebAppSiteExtensionResultOutput) ExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ExtensionType }).(pulumi.StringPtrOutput)
}

// Extension URL.
func (o LookupWebAppSiteExtensionResultOutput) ExtensionUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ExtensionUrl }).(pulumi.StringPtrOutput)
}

// Feed URL.
func (o LookupWebAppSiteExtensionResultOutput) FeedUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.FeedUrl }).(pulumi.StringPtrOutput)
}

// Icon URL.
func (o LookupWebAppSiteExtensionResultOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.IconUrl }).(pulumi.StringPtrOutput)
}

// Resource Id.
func (o LookupWebAppSiteExtensionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Installed timestamp.
func (o LookupWebAppSiteExtensionResultOutput) InstalledDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.InstalledDateTime }).(pulumi.StringPtrOutput)
}

// Installer command line parameters.
func (o LookupWebAppSiteExtensionResultOutput) InstallerCommandLineParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.InstallerCommandLineParams }).(pulumi.StringPtrOutput)
}

// Kind of resource.
func (o LookupWebAppSiteExtensionResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// License URL.
func (o LookupWebAppSiteExtensionResultOutput) LicenseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.LicenseUrl }).(pulumi.StringPtrOutput)
}

// <code>true</code> if the local version is the latest version; <code>false</code> otherwise.
func (o LookupWebAppSiteExtensionResultOutput) LocalIsLatestVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *bool { return v.LocalIsLatestVersion }).(pulumi.BoolPtrOutput)
}

// Local path.
func (o LookupWebAppSiteExtensionResultOutput) LocalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.LocalPath }).(pulumi.StringPtrOutput)
}

// Resource Name.
func (o LookupWebAppSiteExtensionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Project URL.
func (o LookupWebAppSiteExtensionResultOutput) ProjectUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ProjectUrl }).(pulumi.StringPtrOutput)
}

// Provisioning state.
func (o LookupWebAppSiteExtensionResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

// Published timestamp.
func (o LookupWebAppSiteExtensionResultOutput) PublishedDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.PublishedDateTime }).(pulumi.StringPtrOutput)
}

// Summary description.
func (o LookupWebAppSiteExtensionResultOutput) Summary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Summary }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppSiteExtensionResultOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Title }).(pulumi.StringPtrOutput)
}

// Resource type.
func (o LookupWebAppSiteExtensionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) string { return v.Type }).(pulumi.StringOutput)
}

// Version information.
func (o LookupWebAppSiteExtensionResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppSiteExtensionResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppSiteExtensionResultOutput{})
}
