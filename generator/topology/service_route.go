package topology

type ServiceRoute struct {
	Route string `json:"route"`
	DownstreamCalls map[string]string `json:"downstreamCalls,omitempty"`
	MaxLatencyMillis int `json:"maxLatencyMillis"`
	TagSets []TagSet `json:"tagSets"`
}
