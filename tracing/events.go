package tracing

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"github.com/c12o16h1/cdproto/io"
	"github.com/mailru/easyjson"
)

// EventBufferUsage [no description].
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#event-bufferUsage
type EventBufferUsage struct {
	PercentFull float64 `json:"percentFull,omitempty"` // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
	EventCount  float64 `json:"eventCount,omitempty"`  // An approximate number of events in the trace log.
	Value       float64 `json:"value,omitempty"`       // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
}

// EventDataCollected contains an bucket of collected trace events. When
// tracing is stopped collected events will be send as a sequence of
// dataCollected events followed by tracingComplete event.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#event-dataCollected
type EventDataCollected struct {
	Value []easyjson.RawMessage `json:"value"`
}

// EventTracingComplete signals that tracing is stopped and there is no trace
// buffers pending flush, all data were delivered via dataCollected events.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Tracing#event-tracingComplete
type EventTracingComplete struct {
	DataLossOccurred  bool              `json:"dataLossOccurred"`            // Indicates whether some trace data is known to have been lost, e.g. because the trace ring buffer wrapped around.
	Stream            io.StreamHandle   `json:"stream,omitempty"`            // A handle of the stream that holds resulting trace data.
	TraceFormat       StreamFormat      `json:"traceFormat,omitempty"`       // Trace data format of returned stream.
	StreamCompression StreamCompression `json:"streamCompression,omitempty"` // Compression format of returned stream.
}
