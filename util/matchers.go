package util

import (
	"fmt"

	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/storage/metric"
)

// SplitFiltersAndMatchers splits empty matchers off, which are treated as filters, see #220
func SplitFiltersAndMatchers(allMatchers []*metric.LabelMatcher) (filters, matchers []*metric.LabelMatcher) {
	for _, matcher := range allMatchers {
		if matcher.Match("") {
			filters = append(filters, matcher)
		} else {
			matchers = append(matchers, matcher)
		}
	}
	return
}

// ExtractMetricNameFromMetric extract the metric name from a model.Metric
func ExtractMetricNameFromMetric(m model.Metric) (model.LabelValue, error) {
	for name, value := range m {
		if name == model.MetricNameLabel {
			return value, nil
		}
	}
	return "", fmt.Errorf("no MetricNameLabel for chunk")
}