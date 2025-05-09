// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exemplar

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// Filter determines if a measurement should be offered.
//
// The passed ctx needs to contain any baggage or span that were active
// when the measurement was made. This information may be used by the
// Reservoir in making a sampling decision.
type Filter func(context.Context) bool

// SampledFilter is a [Filter] that will only offer measurements
// if the passed context associated with the measurement contains a sampled
// [go.opentelemetry.io/otel/trace.SpanContext].
func SampledFilter(ctx context.Context) bool {
	return trace.SpanContextFromContext(ctx).IsSampled()
}

// AlwaysOnFilter is a [Filter] that always offers measurements.
func AlwaysOnFilter(_ context.Context) bool {
	return true
}
