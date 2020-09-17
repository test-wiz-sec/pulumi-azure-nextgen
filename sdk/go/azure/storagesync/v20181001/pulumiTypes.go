// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ServerEndpoint Health object.
type ServerEndpointHealthResponse struct {
	// Combined Health Status.
	CombinedHealth *string `pulumi:"combinedHealth"`
	// Current progress
	CurrentProgress *SyncProgressStatusResponse `pulumi:"currentProgress"`
	// Download Health Status.
	DownloadHealth *string `pulumi:"downloadHealth"`
	// Download Status
	DownloadStatus *SyncSessionStatusResponse `pulumi:"downloadStatus"`
	// Last Updated Timestamp
	LastUpdatedTimestamp *string `pulumi:"lastUpdatedTimestamp"`
	// Offline Data Transfer State
	OfflineDataTransferStatus *string `pulumi:"offlineDataTransferStatus"`
	// Upload Health Status.
	UploadHealth *string `pulumi:"uploadHealth"`
	// Upload Status
	UploadStatus *SyncSessionStatusResponse `pulumi:"uploadStatus"`
}

// ServerEndpointHealthResponseInput is an input type that accepts ServerEndpointHealthResponseArgs and ServerEndpointHealthResponseOutput values.
// You can construct a concrete instance of `ServerEndpointHealthResponseInput` via:
//
//          ServerEndpointHealthResponseArgs{...}
type ServerEndpointHealthResponseInput interface {
	pulumi.Input

	ToServerEndpointHealthResponseOutput() ServerEndpointHealthResponseOutput
	ToServerEndpointHealthResponseOutputWithContext(context.Context) ServerEndpointHealthResponseOutput
}

// ServerEndpoint Health object.
type ServerEndpointHealthResponseArgs struct {
	// Combined Health Status.
	CombinedHealth pulumi.StringPtrInput `pulumi:"combinedHealth"`
	// Current progress
	CurrentProgress SyncProgressStatusResponsePtrInput `pulumi:"currentProgress"`
	// Download Health Status.
	DownloadHealth pulumi.StringPtrInput `pulumi:"downloadHealth"`
	// Download Status
	DownloadStatus SyncSessionStatusResponsePtrInput `pulumi:"downloadStatus"`
	// Last Updated Timestamp
	LastUpdatedTimestamp pulumi.StringPtrInput `pulumi:"lastUpdatedTimestamp"`
	// Offline Data Transfer State
	OfflineDataTransferStatus pulumi.StringPtrInput `pulumi:"offlineDataTransferStatus"`
	// Upload Health Status.
	UploadHealth pulumi.StringPtrInput `pulumi:"uploadHealth"`
	// Upload Status
	UploadStatus SyncSessionStatusResponsePtrInput `pulumi:"uploadStatus"`
}

func (ServerEndpointHealthResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointHealthResponse)(nil)).Elem()
}

func (i ServerEndpointHealthResponseArgs) ToServerEndpointHealthResponseOutput() ServerEndpointHealthResponseOutput {
	return i.ToServerEndpointHealthResponseOutputWithContext(context.Background())
}

func (i ServerEndpointHealthResponseArgs) ToServerEndpointHealthResponseOutputWithContext(ctx context.Context) ServerEndpointHealthResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointHealthResponseOutput)
}

func (i ServerEndpointHealthResponseArgs) ToServerEndpointHealthResponsePtrOutput() ServerEndpointHealthResponsePtrOutput {
	return i.ToServerEndpointHealthResponsePtrOutputWithContext(context.Background())
}

func (i ServerEndpointHealthResponseArgs) ToServerEndpointHealthResponsePtrOutputWithContext(ctx context.Context) ServerEndpointHealthResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointHealthResponseOutput).ToServerEndpointHealthResponsePtrOutputWithContext(ctx)
}

// ServerEndpointHealthResponsePtrInput is an input type that accepts ServerEndpointHealthResponseArgs, ServerEndpointHealthResponsePtr and ServerEndpointHealthResponsePtrOutput values.
// You can construct a concrete instance of `ServerEndpointHealthResponsePtrInput` via:
//
//          ServerEndpointHealthResponseArgs{...}
//
//  or:
//
//          nil
type ServerEndpointHealthResponsePtrInput interface {
	pulumi.Input

	ToServerEndpointHealthResponsePtrOutput() ServerEndpointHealthResponsePtrOutput
	ToServerEndpointHealthResponsePtrOutputWithContext(context.Context) ServerEndpointHealthResponsePtrOutput
}

type serverEndpointHealthResponsePtrType ServerEndpointHealthResponseArgs

func ServerEndpointHealthResponsePtr(v *ServerEndpointHealthResponseArgs) ServerEndpointHealthResponsePtrInput {
	return (*serverEndpointHealthResponsePtrType)(v)
}

func (*serverEndpointHealthResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointHealthResponse)(nil)).Elem()
}

func (i *serverEndpointHealthResponsePtrType) ToServerEndpointHealthResponsePtrOutput() ServerEndpointHealthResponsePtrOutput {
	return i.ToServerEndpointHealthResponsePtrOutputWithContext(context.Background())
}

func (i *serverEndpointHealthResponsePtrType) ToServerEndpointHealthResponsePtrOutputWithContext(ctx context.Context) ServerEndpointHealthResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerEndpointHealthResponsePtrOutput)
}

// ServerEndpoint Health object.
type ServerEndpointHealthResponseOutput struct{ *pulumi.OutputState }

func (ServerEndpointHealthResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerEndpointHealthResponse)(nil)).Elem()
}

func (o ServerEndpointHealthResponseOutput) ToServerEndpointHealthResponseOutput() ServerEndpointHealthResponseOutput {
	return o
}

func (o ServerEndpointHealthResponseOutput) ToServerEndpointHealthResponseOutputWithContext(ctx context.Context) ServerEndpointHealthResponseOutput {
	return o
}

func (o ServerEndpointHealthResponseOutput) ToServerEndpointHealthResponsePtrOutput() ServerEndpointHealthResponsePtrOutput {
	return o.ToServerEndpointHealthResponsePtrOutputWithContext(context.Background())
}

func (o ServerEndpointHealthResponseOutput) ToServerEndpointHealthResponsePtrOutputWithContext(ctx context.Context) ServerEndpointHealthResponsePtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *ServerEndpointHealthResponse {
		return &v
	}).(ServerEndpointHealthResponsePtrOutput)
}

// Combined Health Status.
func (o ServerEndpointHealthResponseOutput) CombinedHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *string { return v.CombinedHealth }).(pulumi.StringPtrOutput)
}

// Current progress
func (o ServerEndpointHealthResponseOutput) CurrentProgress() SyncProgressStatusResponsePtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *SyncProgressStatusResponse { return v.CurrentProgress }).(SyncProgressStatusResponsePtrOutput)
}

// Download Health Status.
func (o ServerEndpointHealthResponseOutput) DownloadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *string { return v.DownloadHealth }).(pulumi.StringPtrOutput)
}

// Download Status
func (o ServerEndpointHealthResponseOutput) DownloadStatus() SyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *SyncSessionStatusResponse { return v.DownloadStatus }).(SyncSessionStatusResponsePtrOutput)
}

// Last Updated Timestamp
func (o ServerEndpointHealthResponseOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *string { return v.LastUpdatedTimestamp }).(pulumi.StringPtrOutput)
}

// Offline Data Transfer State
func (o ServerEndpointHealthResponseOutput) OfflineDataTransferStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *string { return v.OfflineDataTransferStatus }).(pulumi.StringPtrOutput)
}

// Upload Health Status.
func (o ServerEndpointHealthResponseOutput) UploadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *string { return v.UploadHealth }).(pulumi.StringPtrOutput)
}

// Upload Status
func (o ServerEndpointHealthResponseOutput) UploadStatus() SyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v ServerEndpointHealthResponse) *SyncSessionStatusResponse { return v.UploadStatus }).(SyncSessionStatusResponsePtrOutput)
}

type ServerEndpointHealthResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerEndpointHealthResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerEndpointHealthResponse)(nil)).Elem()
}

func (o ServerEndpointHealthResponsePtrOutput) ToServerEndpointHealthResponsePtrOutput() ServerEndpointHealthResponsePtrOutput {
	return o
}

func (o ServerEndpointHealthResponsePtrOutput) ToServerEndpointHealthResponsePtrOutputWithContext(ctx context.Context) ServerEndpointHealthResponsePtrOutput {
	return o
}

func (o ServerEndpointHealthResponsePtrOutput) Elem() ServerEndpointHealthResponseOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) ServerEndpointHealthResponse { return *v }).(ServerEndpointHealthResponseOutput)
}

// Combined Health Status.
func (o ServerEndpointHealthResponsePtrOutput) CombinedHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.CombinedHealth
	}).(pulumi.StringPtrOutput)
}

// Current progress
func (o ServerEndpointHealthResponsePtrOutput) CurrentProgress() SyncProgressStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *SyncProgressStatusResponse {
		if v == nil {
			return nil
		}
		return v.CurrentProgress
	}).(SyncProgressStatusResponsePtrOutput)
}

// Download Health Status.
func (o ServerEndpointHealthResponsePtrOutput) DownloadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.DownloadHealth
	}).(pulumi.StringPtrOutput)
}

// Download Status
func (o ServerEndpointHealthResponsePtrOutput) DownloadStatus() SyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *SyncSessionStatusResponse {
		if v == nil {
			return nil
		}
		return v.DownloadStatus
	}).(SyncSessionStatusResponsePtrOutput)
}

// Last Updated Timestamp
func (o ServerEndpointHealthResponsePtrOutput) LastUpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastUpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

// Offline Data Transfer State
func (o ServerEndpointHealthResponsePtrOutput) OfflineDataTransferStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.OfflineDataTransferStatus
	}).(pulumi.StringPtrOutput)
}

// Upload Health Status.
func (o ServerEndpointHealthResponsePtrOutput) UploadHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *string {
		if v == nil {
			return nil
		}
		return v.UploadHealth
	}).(pulumi.StringPtrOutput)
}

// Upload Status
func (o ServerEndpointHealthResponsePtrOutput) UploadStatus() SyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v *ServerEndpointHealthResponse) *SyncSessionStatusResponse {
		if v == nil {
			return nil
		}
		return v.UploadStatus
	}).(SyncSessionStatusResponsePtrOutput)
}

// Sync Session status object.
type SyncProgressStatusResponse struct {
	// Applied bytes
	AppliedBytes *int `pulumi:"appliedBytes"`
	// Applied item count.
	AppliedItemCount *int `pulumi:"appliedItemCount"`
	// Per item error count
	PerItemErrorCount *int `pulumi:"perItemErrorCount"`
	// Progress timestamp
	ProgressTimestamp *string `pulumi:"progressTimestamp"`
	// Sync direction.
	SyncDirection *string `pulumi:"syncDirection"`
	// Total bytes
	TotalBytes *int `pulumi:"totalBytes"`
	// Total item count
	TotalItemCount *int `pulumi:"totalItemCount"`
}

// SyncProgressStatusResponseInput is an input type that accepts SyncProgressStatusResponseArgs and SyncProgressStatusResponseOutput values.
// You can construct a concrete instance of `SyncProgressStatusResponseInput` via:
//
//          SyncProgressStatusResponseArgs{...}
type SyncProgressStatusResponseInput interface {
	pulumi.Input

	ToSyncProgressStatusResponseOutput() SyncProgressStatusResponseOutput
	ToSyncProgressStatusResponseOutputWithContext(context.Context) SyncProgressStatusResponseOutput
}

// Sync Session status object.
type SyncProgressStatusResponseArgs struct {
	// Applied bytes
	AppliedBytes pulumi.IntPtrInput `pulumi:"appliedBytes"`
	// Applied item count.
	AppliedItemCount pulumi.IntPtrInput `pulumi:"appliedItemCount"`
	// Per item error count
	PerItemErrorCount pulumi.IntPtrInput `pulumi:"perItemErrorCount"`
	// Progress timestamp
	ProgressTimestamp pulumi.StringPtrInput `pulumi:"progressTimestamp"`
	// Sync direction.
	SyncDirection pulumi.StringPtrInput `pulumi:"syncDirection"`
	// Total bytes
	TotalBytes pulumi.IntPtrInput `pulumi:"totalBytes"`
	// Total item count
	TotalItemCount pulumi.IntPtrInput `pulumi:"totalItemCount"`
}

func (SyncProgressStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncProgressStatusResponse)(nil)).Elem()
}

func (i SyncProgressStatusResponseArgs) ToSyncProgressStatusResponseOutput() SyncProgressStatusResponseOutput {
	return i.ToSyncProgressStatusResponseOutputWithContext(context.Background())
}

func (i SyncProgressStatusResponseArgs) ToSyncProgressStatusResponseOutputWithContext(ctx context.Context) SyncProgressStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncProgressStatusResponseOutput)
}

func (i SyncProgressStatusResponseArgs) ToSyncProgressStatusResponsePtrOutput() SyncProgressStatusResponsePtrOutput {
	return i.ToSyncProgressStatusResponsePtrOutputWithContext(context.Background())
}

func (i SyncProgressStatusResponseArgs) ToSyncProgressStatusResponsePtrOutputWithContext(ctx context.Context) SyncProgressStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncProgressStatusResponseOutput).ToSyncProgressStatusResponsePtrOutputWithContext(ctx)
}

// SyncProgressStatusResponsePtrInput is an input type that accepts SyncProgressStatusResponseArgs, SyncProgressStatusResponsePtr and SyncProgressStatusResponsePtrOutput values.
// You can construct a concrete instance of `SyncProgressStatusResponsePtrInput` via:
//
//          SyncProgressStatusResponseArgs{...}
//
//  or:
//
//          nil
type SyncProgressStatusResponsePtrInput interface {
	pulumi.Input

	ToSyncProgressStatusResponsePtrOutput() SyncProgressStatusResponsePtrOutput
	ToSyncProgressStatusResponsePtrOutputWithContext(context.Context) SyncProgressStatusResponsePtrOutput
}

type syncProgressStatusResponsePtrType SyncProgressStatusResponseArgs

func SyncProgressStatusResponsePtr(v *SyncProgressStatusResponseArgs) SyncProgressStatusResponsePtrInput {
	return (*syncProgressStatusResponsePtrType)(v)
}

func (*syncProgressStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncProgressStatusResponse)(nil)).Elem()
}

func (i *syncProgressStatusResponsePtrType) ToSyncProgressStatusResponsePtrOutput() SyncProgressStatusResponsePtrOutput {
	return i.ToSyncProgressStatusResponsePtrOutputWithContext(context.Background())
}

func (i *syncProgressStatusResponsePtrType) ToSyncProgressStatusResponsePtrOutputWithContext(ctx context.Context) SyncProgressStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncProgressStatusResponsePtrOutput)
}

// Sync Session status object.
type SyncProgressStatusResponseOutput struct{ *pulumi.OutputState }

func (SyncProgressStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncProgressStatusResponse)(nil)).Elem()
}

func (o SyncProgressStatusResponseOutput) ToSyncProgressStatusResponseOutput() SyncProgressStatusResponseOutput {
	return o
}

func (o SyncProgressStatusResponseOutput) ToSyncProgressStatusResponseOutputWithContext(ctx context.Context) SyncProgressStatusResponseOutput {
	return o
}

func (o SyncProgressStatusResponseOutput) ToSyncProgressStatusResponsePtrOutput() SyncProgressStatusResponsePtrOutput {
	return o.ToSyncProgressStatusResponsePtrOutputWithContext(context.Background())
}

func (o SyncProgressStatusResponseOutput) ToSyncProgressStatusResponsePtrOutputWithContext(ctx context.Context) SyncProgressStatusResponsePtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *SyncProgressStatusResponse {
		return &v
	}).(SyncProgressStatusResponsePtrOutput)
}

// Applied bytes
func (o SyncProgressStatusResponseOutput) AppliedBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *int { return v.AppliedBytes }).(pulumi.IntPtrOutput)
}

// Applied item count.
func (o SyncProgressStatusResponseOutput) AppliedItemCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *int { return v.AppliedItemCount }).(pulumi.IntPtrOutput)
}

// Per item error count
func (o SyncProgressStatusResponseOutput) PerItemErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *int { return v.PerItemErrorCount }).(pulumi.IntPtrOutput)
}

// Progress timestamp
func (o SyncProgressStatusResponseOutput) ProgressTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *string { return v.ProgressTimestamp }).(pulumi.StringPtrOutput)
}

// Sync direction.
func (o SyncProgressStatusResponseOutput) SyncDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *string { return v.SyncDirection }).(pulumi.StringPtrOutput)
}

// Total bytes
func (o SyncProgressStatusResponseOutput) TotalBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *int { return v.TotalBytes }).(pulumi.IntPtrOutput)
}

// Total item count
func (o SyncProgressStatusResponseOutput) TotalItemCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncProgressStatusResponse) *int { return v.TotalItemCount }).(pulumi.IntPtrOutput)
}

type SyncProgressStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (SyncProgressStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncProgressStatusResponse)(nil)).Elem()
}

func (o SyncProgressStatusResponsePtrOutput) ToSyncProgressStatusResponsePtrOutput() SyncProgressStatusResponsePtrOutput {
	return o
}

func (o SyncProgressStatusResponsePtrOutput) ToSyncProgressStatusResponsePtrOutputWithContext(ctx context.Context) SyncProgressStatusResponsePtrOutput {
	return o
}

func (o SyncProgressStatusResponsePtrOutput) Elem() SyncProgressStatusResponseOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) SyncProgressStatusResponse { return *v }).(SyncProgressStatusResponseOutput)
}

// Applied bytes
func (o SyncProgressStatusResponsePtrOutput) AppliedBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.AppliedBytes
	}).(pulumi.IntPtrOutput)
}

// Applied item count.
func (o SyncProgressStatusResponsePtrOutput) AppliedItemCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.AppliedItemCount
	}).(pulumi.IntPtrOutput)
}

// Per item error count
func (o SyncProgressStatusResponsePtrOutput) PerItemErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.PerItemErrorCount
	}).(pulumi.IntPtrOutput)
}

// Progress timestamp
func (o SyncProgressStatusResponsePtrOutput) ProgressTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProgressTimestamp
	}).(pulumi.StringPtrOutput)
}

// Sync direction.
func (o SyncProgressStatusResponsePtrOutput) SyncDirection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.SyncDirection
	}).(pulumi.StringPtrOutput)
}

// Total bytes
func (o SyncProgressStatusResponsePtrOutput) TotalBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.TotalBytes
	}).(pulumi.IntPtrOutput)
}

// Total item count
func (o SyncProgressStatusResponsePtrOutput) TotalItemCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncProgressStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.TotalItemCount
	}).(pulumi.IntPtrOutput)
}

// Sync Session status object.
type SyncSessionStatusResponse struct {
	// Last sync per item error count.
	LastSyncPerItemErrorCount *int `pulumi:"lastSyncPerItemErrorCount"`
	// Last sync status
	LastSyncResult *int `pulumi:"lastSyncResult"`
	// Last sync success timestamp
	LastSyncSuccessTimestamp *string `pulumi:"lastSyncSuccessTimestamp"`
	// Last sync timestamp
	LastSyncTimestamp *string `pulumi:"lastSyncTimestamp"`
}

// SyncSessionStatusResponseInput is an input type that accepts SyncSessionStatusResponseArgs and SyncSessionStatusResponseOutput values.
// You can construct a concrete instance of `SyncSessionStatusResponseInput` via:
//
//          SyncSessionStatusResponseArgs{...}
type SyncSessionStatusResponseInput interface {
	pulumi.Input

	ToSyncSessionStatusResponseOutput() SyncSessionStatusResponseOutput
	ToSyncSessionStatusResponseOutputWithContext(context.Context) SyncSessionStatusResponseOutput
}

// Sync Session status object.
type SyncSessionStatusResponseArgs struct {
	// Last sync per item error count.
	LastSyncPerItemErrorCount pulumi.IntPtrInput `pulumi:"lastSyncPerItemErrorCount"`
	// Last sync status
	LastSyncResult pulumi.IntPtrInput `pulumi:"lastSyncResult"`
	// Last sync success timestamp
	LastSyncSuccessTimestamp pulumi.StringPtrInput `pulumi:"lastSyncSuccessTimestamp"`
	// Last sync timestamp
	LastSyncTimestamp pulumi.StringPtrInput `pulumi:"lastSyncTimestamp"`
}

func (SyncSessionStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncSessionStatusResponse)(nil)).Elem()
}

func (i SyncSessionStatusResponseArgs) ToSyncSessionStatusResponseOutput() SyncSessionStatusResponseOutput {
	return i.ToSyncSessionStatusResponseOutputWithContext(context.Background())
}

func (i SyncSessionStatusResponseArgs) ToSyncSessionStatusResponseOutputWithContext(ctx context.Context) SyncSessionStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncSessionStatusResponseOutput)
}

func (i SyncSessionStatusResponseArgs) ToSyncSessionStatusResponsePtrOutput() SyncSessionStatusResponsePtrOutput {
	return i.ToSyncSessionStatusResponsePtrOutputWithContext(context.Background())
}

func (i SyncSessionStatusResponseArgs) ToSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) SyncSessionStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncSessionStatusResponseOutput).ToSyncSessionStatusResponsePtrOutputWithContext(ctx)
}

// SyncSessionStatusResponsePtrInput is an input type that accepts SyncSessionStatusResponseArgs, SyncSessionStatusResponsePtr and SyncSessionStatusResponsePtrOutput values.
// You can construct a concrete instance of `SyncSessionStatusResponsePtrInput` via:
//
//          SyncSessionStatusResponseArgs{...}
//
//  or:
//
//          nil
type SyncSessionStatusResponsePtrInput interface {
	pulumi.Input

	ToSyncSessionStatusResponsePtrOutput() SyncSessionStatusResponsePtrOutput
	ToSyncSessionStatusResponsePtrOutputWithContext(context.Context) SyncSessionStatusResponsePtrOutput
}

type syncSessionStatusResponsePtrType SyncSessionStatusResponseArgs

func SyncSessionStatusResponsePtr(v *SyncSessionStatusResponseArgs) SyncSessionStatusResponsePtrInput {
	return (*syncSessionStatusResponsePtrType)(v)
}

func (*syncSessionStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncSessionStatusResponse)(nil)).Elem()
}

func (i *syncSessionStatusResponsePtrType) ToSyncSessionStatusResponsePtrOutput() SyncSessionStatusResponsePtrOutput {
	return i.ToSyncSessionStatusResponsePtrOutputWithContext(context.Background())
}

func (i *syncSessionStatusResponsePtrType) ToSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) SyncSessionStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncSessionStatusResponsePtrOutput)
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

func (o SyncSessionStatusResponseOutput) ToSyncSessionStatusResponsePtrOutput() SyncSessionStatusResponsePtrOutput {
	return o.ToSyncSessionStatusResponsePtrOutputWithContext(context.Background())
}

func (o SyncSessionStatusResponseOutput) ToSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) SyncSessionStatusResponsePtrOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) *SyncSessionStatusResponse {
		return &v
	}).(SyncSessionStatusResponsePtrOutput)
}

// Last sync per item error count.
func (o SyncSessionStatusResponseOutput) LastSyncPerItemErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) *int { return v.LastSyncPerItemErrorCount }).(pulumi.IntPtrOutput)
}

// Last sync status
func (o SyncSessionStatusResponseOutput) LastSyncResult() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) *int { return v.LastSyncResult }).(pulumi.IntPtrOutput)
}

// Last sync success timestamp
func (o SyncSessionStatusResponseOutput) LastSyncSuccessTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) *string { return v.LastSyncSuccessTimestamp }).(pulumi.StringPtrOutput)
}

// Last sync timestamp
func (o SyncSessionStatusResponseOutput) LastSyncTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncSessionStatusResponse) *string { return v.LastSyncTimestamp }).(pulumi.StringPtrOutput)
}

type SyncSessionStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (SyncSessionStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncSessionStatusResponse)(nil)).Elem()
}

func (o SyncSessionStatusResponsePtrOutput) ToSyncSessionStatusResponsePtrOutput() SyncSessionStatusResponsePtrOutput {
	return o
}

func (o SyncSessionStatusResponsePtrOutput) ToSyncSessionStatusResponsePtrOutputWithContext(ctx context.Context) SyncSessionStatusResponsePtrOutput {
	return o
}

func (o SyncSessionStatusResponsePtrOutput) Elem() SyncSessionStatusResponseOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) SyncSessionStatusResponse { return *v }).(SyncSessionStatusResponseOutput)
}

// Last sync per item error count.
func (o SyncSessionStatusResponsePtrOutput) LastSyncPerItemErrorCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.LastSyncPerItemErrorCount
	}).(pulumi.IntPtrOutput)
}

// Last sync status
func (o SyncSessionStatusResponsePtrOutput) LastSyncResult() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *int {
		if v == nil {
			return nil
		}
		return v.LastSyncResult
	}).(pulumi.IntPtrOutput)
}

// Last sync success timestamp
func (o SyncSessionStatusResponsePtrOutput) LastSyncSuccessTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastSyncSuccessTimestamp
	}).(pulumi.StringPtrOutput)
}

// Last sync timestamp
func (o SyncSessionStatusResponsePtrOutput) LastSyncTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncSessionStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastSyncTimestamp
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerEndpointHealthResponseOutput{})
	pulumi.RegisterOutputType(ServerEndpointHealthResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncProgressStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncProgressStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncSessionStatusResponseOutput{})
	pulumi.RegisterOutputType(SyncSessionStatusResponsePtrOutput{})
}