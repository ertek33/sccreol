// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*vMExecutedOperationMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (v VMExecutedOperation) MarshalJSON() ([]byte, error) {
	type VMExecutedOperation struct {
		Used  uint64         `json:"used"`
		Push  []*hexutil.Big `json:"push"`
		Mem   *MemoryDiff    `json:"mem"`
		Store *StorageDiff   `json:"store"`
	}
	var enc VMExecutedOperation
	enc.Used = v.Used
	if v.Push != nil {
		enc.Push = make([]*hexutil.Big, len(v.Push))
		for k, v := range v.Push {
			enc.Push[k] = (*hexutil.Big)(v)
		}
	}
	enc.Mem = v.Mem
	enc.Store = v.Store
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (v *VMExecutedOperation) UnmarshalJSON(input []byte) error {
	type VMExecutedOperation struct {
		Used  *uint64        `json:"used"`
		Push  []*hexutil.Big `json:"push"`
		Mem   *MemoryDiff    `json:"mem"`
		Store *StorageDiff   `json:"store"`
	}
	var dec VMExecutedOperation
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Used != nil {
		v.Used = *dec.Used
	}
	if dec.Push != nil {
		v.Push = make([]*big.Int, len(dec.Push))
		for k, _v := range dec.Push {
			v.Push[k] = (*big.Int)(_v)
		}
	}
	if dec.Mem != nil {
		v.Mem = dec.Mem
	}
	if dec.Store != nil {
		v.Store = dec.Store
	}
	return nil
}
