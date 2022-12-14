// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Files not syncing error object
type FilesNotSyncingErrorResponse struct {
	// Error code (HResult)
	ErrorCode int `pulumi:"errorCode"`
	// Count of persistent files not syncing with the specified error code
	PersistentCount float64 `pulumi:"persistentCount"`
	// Count of transient files not syncing with the specified error code
	TransientCount float64 `pulumi:"transientCount"`
}

// Files not syncing error object
type FilesNotSyncingErrorResponseOutput struct{ *pulumi.OutputState }

func (FilesNotSyncingErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilesNotSyncingErrorResponse)(nil)).Elem()
}

func (o FilesNotSyncingErrorResponseOutput) ToFilesNotSyncingErrorResponseOutput() FilesNotSyncingErrorResponseOutput {
	return o
}

func (o FilesNotSyncingErrorResponseOutput) ToFilesNotSyncingErrorResponseOutputWithContext(ctx context.Context) FilesNotSyncingErrorResponseOutput {
	return o
}

// Error code (HResult)
func (o FilesNotSyncingErrorResponseOutput) ErrorCode() pulumi.IntOutput {
	return o.ApplyT(func(v FilesNotSyncingErrorResponse) int { return v.ErrorCode }).(pulumi.IntOutput)
}

// Count of persistent files not syncing with the specified error code
func (o FilesNotSyncingErrorResponseOutput) PersistentCount() pulumi.Float64Output {
	return o.ApplyT(func(v FilesNotSyncingErrorResponse) float64 { return v.PersistentCount }).(pulumi.Float64Output)
}

// Count of transient files not syncing with the specified error code
func (o FilesNotSyncingErrorResponseOutput) TransientCount() pulumi.Float64Output {
	return o.ApplyT(func(v FilesNotSyncingErrorResponse) float64 { return v.TransientCount }).(pulumi.Float64Output)
}

type FilesNotSyncingErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (FilesNotSyncingErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FilesNotSyncingErrorResponse)(nil)).Elem()
}

func (o FilesNotSyncingErrorResponseArrayOutput) ToFilesNotSyncingErrorResponseArrayOutput() FilesNotSyncingErrorResponseArrayOutput {
	return o
}

func (o FilesNotSyncingErrorResponseArrayOutput) ToFilesNotSyncingErrorResponseArrayOutputWithContext(ctx context.Context) FilesNotSyncingErrorResponseArrayOutput {
	return o
}

func (o FilesNotSyncingErrorResponseArrayOutput) Index(i pulumi.IntInput) FilesNotSyncingErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FilesNotSyncingErrorResponse {
		return vs[0].([]FilesNotSyncingErrorResponse)[vs[1].(int)]
	}).(FilesNotSyncingErrorResponseOutput)
}

// Server Endpoint sync status
type ServerEndpointSyncStatusResponse struct {
	// Combined Health Status.
	CombinedHealth string `pulumi:"combinedHealth"`
	// Download sync activity
	DownloadActivity SyncActivityStatusResponse `pulumi:"downloadActivity"`
	// Download Health Status.
	DownloadHealth string `pulumi:"downloadHealth"`
	// Download Status
	DownloadStatus SyncSessionStatusResponse `pulumi:"downloadStatus"`
	// Last Updated Timestamp
	LastUpdatedTimestamp string `pulumi:"lastUpdatedTimestamp"`
	// Offline Data Transfer State
	OfflineDataTransferStatus string `pulumi:"offlineDataTransferStatus"`
	// Sync activity
	SyncActivity string `pulumi:"syncActivity"`
	// Total count of persistent files not syncing (combined upload + download). Reserved for future use.
	TotalPersistentFilesNotSyncingCount float64 `pulumi:"totalPersistentFilesNotSyncingCount"`
	// Upload sync activity
	UploadActivity SyncActivityStatusResponse `pulumi:"uploadActivity"`
	// Upload Health Status.
	UploadHealth string `pulumi:"uploadHealth"`
	// Upload Status
	UploadStatus SyncSessionStatusResponse `pulumi:"uploadStatus"`
}

// Server Endpoint sync status
type ServerEndpointSyncStatusResponseOutput struct{ *pulumi.OutputState }

func (ServerEndpointSyncStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointSyncStatusResponse)(nil)).Elem()
}

func (o ServerEndpointSyncStatusResponseOutput) ToServerEndpointSyncStatusResponseOutput() ServerEndpointSyncStatusResponseOutput {
	return o
}

func (o ServerEndpointSyncStatusResponseOutput) ToServerEndpointSyncStatusResponseOutputWithContext(ctx context.Context) ServerEndpointSyncStatusResponseOutput {
	return o
}

// Combined Health Status.
func (o ServerEndpointSyncStatusResponseOutput) CombinedHealth() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.CombinedHealth }).(pulumi.StringOutput)
}

// Download sync activity
func (o ServerEndpointSyncStatusResponseOutput) DownloadActivity() SyncActivityStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) SyncActivityStatusResponse { return v.DownloadActivity }).(SyncActivityStatusResponseOutput)
}

// Download Health Status.
func (o ServerEndpointSyncStatusResponseOutput) DownloadHealth() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.DownloadHealth }).(pulumi.StringOutput)
}

// Download Status
func (o ServerEndpointSyncStatusResponseOutput) DownloadStatus() SyncSessionStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) SyncSessionStatusResponse { return v.DownloadStatus }).(SyncSessionStatusResponseOutput)
}

// Last Updated Timestamp
func (o ServerEndpointSyncStatusResponseOutput) LastUpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.LastUpdatedTimestamp }).(pulumi.StringOutput)
}

// Offline Data Transfer State
func (o ServerEndpointSyncStatusResponseOutput) OfflineDataTransferStatus() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.OfflineDataTransferStatus }).(pulumi.StringOutput)
}

// Sync activity
func (o ServerEndpointSyncStatusResponseOutput) SyncActivity() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.SyncActivity }).(pulumi.StringOutput)
}

// Total count of persistent files not syncing (combined upload + download). Reserved for future use.
func (o ServerEndpointSyncStatusResponseOutput) TotalPersistentFilesNotSyncingCount() pulumi.Float64Output {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) float64 { return v.TotalPersistentFilesNotSyncingCount }).(pulumi.Float64Output)
}

// Upload sync activity
func (o ServerEndpointSyncStatusResponseOutput) UploadActivity() SyncActivityStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) SyncActivityStatusResponse { return v.UploadActivity }).(SyncActivityStatusResponseOutput)
}

// Upload Health Status.
func (o ServerEndpointSyncStatusResponseOutput) UploadHealth() pulumi.StringOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) string { return v.UploadHealth }).(pulumi.StringOutput)
}

// Upload Status
func (o ServerEndpointSyncStatusResponseOutput) UploadStatus() SyncSessionStatusResponseOutput {
	return o.ApplyT(func(v ServerEndpointSyncStatusResponse) SyncSessionStatusResponse { return v.UploadStatus }).(SyncSessionStatusResponseOutput)
}

// Sync Session status object.
type SyncActivityStatusResponse struct {
	// Applied bytes
	AppliedBytes float64 `pulumi:"appliedBytes"`
	// Applied item count.
	AppliedItemCount float64 `pulumi:"appliedItemCount"`
	// Per item error count
	PerItemErrorCount float64 `pulumi:"perItemErrorCount"`
	// Timestamp when properties were updated
	Timestamp string `pulumi:"timestamp"`
	// Total bytes (if available)
	TotalBytes float64 `pulumi:"totalBytes"`
	// Total item count (if available)
	TotalItemCount float64 `pulumi:"totalItemCount"`
}

// Sync Session status object.
type SyncActivityStatusResponseOutput struct{ *pulumi.OutputState }

func (SyncActivityStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncActivityStatusResponse)(nil)).Elem()
}

func (o SyncActivityStatusResponseOutput) ToSyncActivityStatusResponseOutput() SyncActivityStatusResponseOutput {
	return o
}

func (o SyncActivityStatusResponseOutput) ToSyncActivityStatusResponseOutputWithContext(ctx context.Context) SyncActivityStatusResponseOutput {
	return o
}

// Applied bytes
func (o SyncActivityStatusResponseOutput) AppliedBytes() pulumi.Float64Output {
	return o.ApplyT(func(v SyncActivityStatusResponse) float64 { return v.AppliedBytes }).(pulumi.Float64Output)
}

// Applied item count.
func (o SyncActivityStatusResponseOutput) AppliedItemCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncActivityStatusResponse) float64 { return v.AppliedItemCount }).(pulumi.Float64Output)
}

// Per item error count
func (o SyncActivityStatusResponseOutput) PerItemErrorCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncActivityStatusResponse) float64 { return v.PerItemErrorCount }).(pulumi.Float64Output)
}

// Timestamp when properties were updated
func (o SyncActivityStatusResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v SyncActivityStatusResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

// Total bytes (if available)
func (o SyncActivityStatusResponseOutput) TotalBytes() pulumi.Float64Output {
	return o.ApplyT(func(v SyncActivityStatusResponse) float64 { return v.TotalBytes }).(pulumi.Float64Output)
}

// Total item count (if available)
func (o SyncActivityStatusResponseOutput) TotalItemCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncActivityStatusResponse) float64 { return v.TotalItemCount }).(pulumi.Float64Output)
}

// Sync Session status object.
type SyncSessionStatusResponse struct {
	// Array of per-item errors coming from the last sync session. Reserved for future use.
	FilesNotSyncingErrors []FilesNotSyncingErrorResponse `pulumi:"filesNotSyncingErrors"`
	// Last sync per item error count.
	LastSyncPerItemErrorCount float64 `pulumi:"lastSyncPerItemErrorCount"`
	// Last sync result (HResult)
	LastSyncResult int `pulumi:"lastSyncResult"`
	// Last sync success timestamp
	LastSyncSuccessTimestamp string `pulumi:"lastSyncSuccessTimestamp"`
	// Last sync timestamp
	LastSyncTimestamp string `pulumi:"lastSyncTimestamp"`
	// Count of persistent files not syncing. Reserved for future use.
	PersistentFilesNotSyncingCount float64 `pulumi:"persistentFilesNotSyncingCount"`
	// Count of transient files not syncing. Reserved for future use.
	TransientFilesNotSyncingCount float64 `pulumi:"transientFilesNotSyncingCount"`
}

// Sync Session status object.
type SyncSessionStatusResponseOutput struct{ *pulumi.OutputState }

func (SyncSessionStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncSessionStatusResponse)(nil)).Elem()
}

func (o SyncSessionStatusResponseOutput) ToSyncSessionStatusResponseOutput() SyncSessionStatusResponseOutput {
	return o
}

func (o SyncSessionStatusResponseOutput) ToSyncSessionStatusResponseOutputWithContext(ctx context.Context) SyncSessionStatusResponseOutput {
	return o
}

// Array of per-item errors coming from the last sync session. Reserved for future use.
func (o SyncSessionStatusResponseOutput) FilesNotSyncingErrors() FilesNotSyncingErrorResponseArrayOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) []FilesNotSyncingErrorResponse { return v.FilesNotSyncingErrors }).(FilesNotSyncingErrorResponseArrayOutput)
}

// Last sync per item error count.
func (o SyncSessionStatusResponseOutput) LastSyncPerItemErrorCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncSessionStatusResponse) float64 { return v.LastSyncPerItemErrorCount }).(pulumi.Float64Output)
}

// Last sync result (HResult)
func (o SyncSessionStatusResponseOutput) LastSyncResult() pulumi.IntOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) int { return v.LastSyncResult }).(pulumi.IntOutput)
}

// Last sync success timestamp
func (o SyncSessionStatusResponseOutput) LastSyncSuccessTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) string { return v.LastSyncSuccessTimestamp }).(pulumi.StringOutput)
}

// Last sync timestamp
func (o SyncSessionStatusResponseOutput) LastSyncTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) string { return v.LastSyncTimestamp }).(pulumi.StringOutput)
}

// Count of persistent files not syncing. Reserved for future use.
func (o SyncSessionStatusResponseOutput) PersistentFilesNotSyncingCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncSessionStatusResponse) float64 { return v.PersistentFilesNotSyncingCount }).(pulumi.Float64Output)
}

// Count of transient files not syncing. Reserved for future use.
func (o SyncSessionStatusResponseOutput) TransientFilesNotSyncingCount() pulumi.Float64Output {
	return o.ApplyT(func(v SyncSessionStatusResponse) float64 { return v.TransientFilesNotSyncingCount }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(FilesNotSyncingErrorResponseOutput{})
	pulumi.RegisterOutputType(FilesNotSyncingErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerEndpointSyncStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncActivityStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncSessionStatusResponseOutput{})
}
