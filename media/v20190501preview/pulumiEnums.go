// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190501preview

// The encoding type for Live Event.  This value is specified at creation time and cannot be updated.
type LiveEventEncodingType string

const (
	LiveEventEncodingTypeNone         = LiveEventEncodingType("None")
	LiveEventEncodingTypeBasic        = LiveEventEncodingType("Basic")
	LiveEventEncodingTypeStandard     = LiveEventEncodingType("Standard")
	LiveEventEncodingTypePremium1080p = LiveEventEncodingType("Premium1080p")
)

// The streaming protocol for the Live Event.  This is specified at creation time and cannot be updated.
type LiveEventInputProtocol string

const (
	LiveEventInputProtocolFragmentedMP4 = LiveEventInputProtocol("FragmentedMP4")
	LiveEventInputProtocolRTMP          = LiveEventInputProtocol("RTMP")
)

type StreamOptionsFlag string

const (
	StreamOptionsFlagDefault    = StreamOptionsFlag("Default")
	StreamOptionsFlagLowLatency = StreamOptionsFlag("LowLatency")
)

func init() {
}