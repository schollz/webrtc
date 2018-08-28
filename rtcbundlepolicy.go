package webrtc

// RTCBundlePolicy affects which media tracks are negotiated if the remote
// endpoint is not bundle-aware, and what ICE candidates are gathered. If the
// remote endpoint is bundle-aware, all media tracks and data channels are
// bundled onto the same transport.
type RTCBundlePolicy int

const (
	// RTCBundlePolicyBalanced indicates to gather ICE candidates for each
	// media type in use (audio, video, and data). If the remote endpoint is
	// not bundle-aware, negotiate only one audio and video track on separate
	// transports.
	RTCBundlePolicyBalanced RTCBundlePolicy = iota + 1

	// RTCBundlePolicyMaxCompat indicates to gather ICE candidates for each
	// track. If the remote endpoint is not bundle-aware, negotiate all media
	// tracks on separate transports.
	RTCBundlePolicyMaxCompat

	// RTCBundlePolicyMaxBundle indicates to gather ICE candidates for only
	// one track. If the remote endpoint is not bundle-aware, negotiate only
	// one media track.
	RTCBundlePolicyMaxBundle
)

// NewRTCBundlePolicy defines a proceedure for creating a new RTCBundlePolicy
// from a raw string naming the bundle policy.
func NewRTCBundlePolicy(raw string) RTCBundlePolicy {
	switch raw {
	case "balanced":
		return RTCBundlePolicyBalanced
	case "max-compat":
		return RTCBundlePolicyMaxCompat
	case "max-bundle":
		return RTCBundlePolicyMaxBundle
	default:
		return RTCBundlePolicy(Unknown)
	}
}

func (t RTCBundlePolicy) String() string {
	switch t {
	case RTCBundlePolicyBalanced:
		return "balanced"
	case RTCBundlePolicyMaxCompat:
		return "max-compat"
	case RTCBundlePolicyMaxBundle:
		return "max-bundle"
	default:
		return ErrUnknownType.Error()
	}
}
