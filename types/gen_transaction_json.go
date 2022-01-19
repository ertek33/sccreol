// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ = (*transactionMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (t Transaction) MarshalJSON() ([]byte, error) {
	type Transaction struct {
		Accesses             types.AccessList `json:"accessList,omitempty"`
		BlockHash            *common.Hash     `json:"blockHash"`
		BlockNumber          *hexutil.Big     `json:"blockNumber"`
		ChainID              *hexutil.Big     `json:"chainId,omitempty"`
		From                 common.Address   `json:"from"`
		Gas                  hexutil.Uint64   `json:"gas"                            gencodec:"required"`
		GasPrice             *hexutil.Big     `json:"gasPrice"`
		Hash                 common.Hash      `json:"hash"`
		Input                hexutil.Bytes    `json:"input"                          gencodec:"required"`
		MaxFeePerGas         *hexutil.Big     `json:"maxFeePerGas,omitempty"`
		MaxPriorityFeePerGas *hexutil.Big     `json:"maxPriorityFeePerGas,omitempty"`
		Nonce                hexutil.Uint64   `json:"nonce"                          gencodec:"required"`
		R                    *hexutil.Big     `json:"r"                              gencodec:"required"`
		S                    *hexutil.Big     `json:"s"                              gencodec:"required"`
		To                   *common.Address  `json:"to" rlp:"nil"`
		TransactionIndex     *hexutil.Uint64  `json:"transactionIndex"`
		Type                 hexutil.Uint64   `json:"type"`
		V                    *hexutil.Big     `json:"v"                              gencodec:"required"`
		Value                *hexutil.Big     `json:"value"                          gencodec:"required"`
	}
	var enc Transaction
	enc.Accesses = t.Accesses
	enc.BlockHash = t.BlockHash
	enc.BlockNumber = (*hexutil.Big)(t.BlockNumber)
	enc.ChainID = (*hexutil.Big)(t.ChainID)
	enc.From = t.From
	enc.Gas = hexutil.Uint64(t.Gas)
	enc.GasPrice = (*hexutil.Big)(t.GasPrice)
	enc.Hash = t.Hash
	enc.Input = t.Input
	enc.MaxFeePerGas = (*hexutil.Big)(t.MaxFeePerGas)
	enc.MaxPriorityFeePerGas = (*hexutil.Big)(t.MaxPriorityFeePerGas)
	enc.Nonce = hexutil.Uint64(t.Nonce)
	enc.R = (*hexutil.Big)(t.R)
	enc.S = (*hexutil.Big)(t.S)
	enc.To = t.To
	enc.TransactionIndex = (*hexutil.Uint64)(t.TransactionIndex)
	enc.Type = hexutil.Uint64(t.Type)
	enc.V = (*hexutil.Big)(t.V)
	enc.Value = (*hexutil.Big)(t.Value)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (t *Transaction) UnmarshalJSON(input []byte) error {
	type Transaction struct {
		Accesses             *types.AccessList `json:"accessList,omitempty"`
		BlockHash            *common.Hash      `json:"blockHash"`
		BlockNumber          *hexutil.Big      `json:"blockNumber"`
		ChainID              *hexutil.Big      `json:"chainId,omitempty"`
		From                 *common.Address   `json:"from"`
		Gas                  *hexutil.Uint64   `json:"gas"                            gencodec:"required"`
		GasPrice             *hexutil.Big      `json:"gasPrice"`
		Hash                 *common.Hash      `json:"hash"`
		Input                *hexutil.Bytes    `json:"input"                          gencodec:"required"`
		MaxFeePerGas         *hexutil.Big      `json:"maxFeePerGas,omitempty"`
		MaxPriorityFeePerGas *hexutil.Big      `json:"maxPriorityFeePerGas,omitempty"`
		Nonce                *hexutil.Uint64   `json:"nonce"                          gencodec:"required"`
		R                    *hexutil.Big      `json:"r"                              gencodec:"required"`
		S                    *hexutil.Big      `json:"s"                              gencodec:"required"`
		To                   *common.Address   `json:"to" rlp:"nil"`
		TransactionIndex     *hexutil.Uint64   `json:"transactionIndex"`
		Type                 *hexutil.Uint64   `json:"type"`
		V                    *hexutil.Big      `json:"v"                              gencodec:"required"`
		Value                *hexutil.Big      `json:"value"                          gencodec:"required"`
	}
	var dec Transaction
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Accesses != nil {
		t.Accesses = *dec.Accesses
	}
	if dec.BlockHash != nil {
		t.BlockHash = dec.BlockHash
	}
	if dec.BlockNumber != nil {
		t.BlockNumber = (*big.Int)(dec.BlockNumber)
	}
	if dec.ChainID != nil {
		t.ChainID = (*big.Int)(dec.ChainID)
	}
	if dec.From != nil {
		t.From = *dec.From
	}
	if dec.Gas == nil {
		return errors.New("missing required field 'gas' for Transaction")
	}
	t.Gas = uint64(*dec.Gas)
	if dec.GasPrice != nil {
		t.GasPrice = (*big.Int)(dec.GasPrice)
	}
	if dec.Hash != nil {
		t.Hash = *dec.Hash
	}
	if dec.Input == nil {
		return errors.New("missing required field 'input' for Transaction")
	}
	t.Input = *dec.Input
	if dec.MaxFeePerGas != nil {
		t.MaxFeePerGas = (*big.Int)(dec.MaxFeePerGas)
	}
	if dec.MaxPriorityFeePerGas != nil {
		t.MaxPriorityFeePerGas = (*big.Int)(dec.MaxPriorityFeePerGas)
	}
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' for Transaction")
	}
	t.Nonce = uint64(*dec.Nonce)
	if dec.R == nil {
		return errors.New("missing required field 'r' for Transaction")
	}
	t.R = (*big.Int)(dec.R)
	if dec.S == nil {
		return errors.New("missing required field 's' for Transaction")
	}
	t.S = (*big.Int)(dec.S)
	if dec.To != nil {
		t.To = dec.To
	}
	if dec.TransactionIndex != nil {
		t.TransactionIndex = (*uint64)(dec.TransactionIndex)
	}
	if dec.Type != nil {
		t.Type = uint64(*dec.Type)
	}
	if dec.V == nil {
		return errors.New("missing required field 'v' for Transaction")
	}
	t.V = (*big.Int)(dec.V)
	if dec.Value == nil {
		return errors.New("missing required field 'value' for Transaction")
	}
	t.Value = (*big.Int)(dec.Value)
	return nil
}
