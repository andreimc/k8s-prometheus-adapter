package config

type MetricsDiscoveryConfig struct {
	// Rules specifies how to discover and map Prometheus metrics to
	// custom metrics API resources.  The rules are applied in
	// order.
	Rules []DiscoveryRule `yaml:"rules"`
}

// DiscoveryRule describes on set of rules for transforming Prometheus metrics to/from
// custom metrics API resources.
type DiscoveryRule struct {
	// SeriesQuery specifies which metrics this rule should consider via a Prometheus query
	// series selector query.
	SeriesQuery string `yaml:"seriesQuery"`
	// Resources specifies how associated Kubernetes resources should be discovered for
	// the given metrics.
	Resources ResourceMapping `yaml:"resources"`
	// Naming specifies how the metric name should be transformed between custom metric
	// API resources, and Prometheus metric names.
	Naming NameMapping `yaml:"naming"`
	// MetricsQuery specifies modifications to the metrics query, such as converting
	// cumulative metrics to rate metrics.  It is a template where .LabelMatchers is
	// a slice of the base label matchers and .Series is the series name, and .GroupBy
	// is the base slice of expected group-by label names.
	MetricsQuery string `yaml:"metricsQueries,omitempty"`
}

// ResourceMapping specifies how to map Kubernetes resources to Prometheus labels
type ResourceMapping struct {
	Match string `yaml:"match,omitempty"`
	// Template specifies a golang string template for converting a Kubernetes
	// group-resource to a Prometheus label.  The template object contains
	// the `.Group` and `.Resource` fields.  The `.Group` field will have
	// dots replaced with underscores.
	Template string `yaml:"template,omitempty"`
	// Overrides specifies exceptions to the above template, mapping label names
	// to group-resources
	Overrides map[string]GroupResource `yaml:"overrides,omitempty"`
}

// GroupResource represents a Kubernetes group-resource.
type GroupResource struct {
	Group    string `yaml:"group,omitempty"`
	Resource string `yaml:"resource"`
}

// NameMapping specifies how to convert Prometheus metrics
// to/from custom metrics API resources.
type NameMapping struct {
	// Prefix specifies that a prefix should be removed from
	// the Prometheus metric name.
	Prefix string `yaml:"prefix,omitempty"`
	// Suffix specifies that a suffix should be removed from
	// the Prometheus metric name.
	Suffix string `yaml:"suffix,omitempty"`
	// ConstantName fully overrides the name of the custom metrics API
	// resource.  This can only be used when SeriesQuery produces a single
	// metric family.
	ConstantName string `yaml:"constantName,omitempty"`
}
