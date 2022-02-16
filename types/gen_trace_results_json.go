// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*traceResultsMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (t TraceResults) MarshalJSON() ([]byte, error) {
	type TraceResults struct {
		Output    hexutil.Bytes `json:"output"`
		Trace     []Trace       `json:"trace"`
		VmTrace   *VMTrace      `json:"vmTrace"`
		StateDiff *StateDiff    `json:"stateDiff"`
	}
	var enc TraceResults
	enc.Output = t.Output
	enc.Trace = t.Trace
	enc.VmTrace = t.VmTrace
	enc.StateDiff = t.StateDiff
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (t *TraceResults) UnmarshalJSON(input []byte) error {
	type TraceResults struct {
		Output    *hexutil.Bytes `json:"output"`
		Trace     []Trace        `json:"trace"`
		VmTrace   *VMTrace       `json:"vmTrace"`
		StateDiff *StateDiff     `json:"stateDiff"`
	}
	var dec TraceResults
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Output != nil {
		t.Output = *dec.Output
	}
	if dec.Trace != nil {
		t.Trace = dec.Trace
	}
	if dec.VmTrace != nil {
		t.VmTrace = dec.VmTrace
	}
	if dec.StateDiff != nil {
		t.StateDiff = dec.StateDiff
	}
	return nil
}
