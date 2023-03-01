// Code generated by rlpgen. DO NOT EDIT.

//go:build !norlpgen
// +build !norlpgen

package types

import "github.com/ethereum/go-ethereum/common"
import "github.com/ethereum/go-ethereum/core/types"
import "github.com/ethereum/go-ethereum/rlp"
import "io"

func (obj *StargazerBlock) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	_tmp0 := w.List()
	if obj.Header == nil {
		w.Write([]byte{0xC0})
	} else {
		_tmp1 := w.List()
		w.WriteBytes(obj.Header.ParentHash[:])
		w.WriteBytes(obj.Header.UncleHash[:])
		w.WriteBytes(obj.Header.Coinbase[:])
		w.WriteBytes(obj.Header.Root[:])
		w.WriteBytes(obj.Header.TxHash[:])
		w.WriteBytes(obj.Header.ReceiptHash[:])
		w.WriteBytes(obj.Header.Bloom[:])
		if obj.Header.Difficulty == nil {
			w.Write(rlp.EmptyString)
		} else {
			if obj.Header.Difficulty.Sign() == -1 {
				return rlp.ErrNegativeBigInt
			}
			w.WriteBigInt(obj.Header.Difficulty)
		}
		if obj.Header.Number == nil {
			w.Write(rlp.EmptyString)
		} else {
			if obj.Header.Number.Sign() == -1 {
				return rlp.ErrNegativeBigInt
			}
			w.WriteBigInt(obj.Header.Number)
		}
		w.WriteUint64(obj.Header.GasLimit)
		w.WriteUint64(obj.Header.GasUsed)
		w.WriteUint64(obj.Header.Time)
		w.WriteBytes(obj.Header.Extra)
		w.WriteBytes(obj.Header.MixDigest[:])
		w.WriteBytes(obj.Header.Nonce[:])
		_tmp2 := obj.Header.BaseFee != nil
		_tmp3 := obj.Header.WithdrawalsHash != nil
		if _tmp2 || _tmp3 {
			if obj.Header.BaseFee == nil {
				w.Write(rlp.EmptyString)
			} else {
				if obj.Header.BaseFee.Sign() == -1 {
					return rlp.ErrNegativeBigInt
				}
				w.WriteBigInt(obj.Header.BaseFee)
			}
		}
		if _tmp3 {
			if obj.Header.WithdrawalsHash == nil {
				w.Write([]byte{0x80})
			} else {
				w.WriteBytes(obj.Header.WithdrawalsHash[:])
			}
		}
		w.ListEnd(_tmp1)
	}
	w.ListEnd(_tmp0)
	return w.Flush()
}

func (obj *StargazerBlock) DecodeRLP(dec *rlp.Stream) error {
	var _tmp0 StargazerBlock
	{
		if _, err := dec.List(); err != nil {
			return err
		}
		// Header:
		var _tmp1 types.Header
		{
			if _, err := dec.List(); err != nil {
				return err
			}
			// ParentHash:
			var _tmp2 common.Hash
			if err := dec.ReadBytes(_tmp2[:]); err != nil {
				return err
			}
			_tmp1.ParentHash = _tmp2
			// UncleHash:
			var _tmp3 common.Hash
			if err := dec.ReadBytes(_tmp3[:]); err != nil {
				return err
			}
			_tmp1.UncleHash = _tmp3
			// Coinbase:
			var _tmp4 common.Address
			if err := dec.ReadBytes(_tmp4[:]); err != nil {
				return err
			}
			_tmp1.Coinbase = _tmp4
			// Root:
			var _tmp5 common.Hash
			if err := dec.ReadBytes(_tmp5[:]); err != nil {
				return err
			}
			_tmp1.Root = _tmp5
			// TxHash:
			var _tmp6 common.Hash
			if err := dec.ReadBytes(_tmp6[:]); err != nil {
				return err
			}
			_tmp1.TxHash = _tmp6
			// ReceiptHash:
			var _tmp7 common.Hash
			if err := dec.ReadBytes(_tmp7[:]); err != nil {
				return err
			}
			_tmp1.ReceiptHash = _tmp7
			// Bloom:
			var _tmp8 types.Bloom
			if err := dec.ReadBytes(_tmp8[:]); err != nil {
				return err
			}
			_tmp1.Bloom = _tmp8
			// Difficulty:
			_tmp9, err := dec.BigInt()
			if err != nil {
				return err
			}
			_tmp1.Difficulty = _tmp9
			// Number:
			_tmp10, err := dec.BigInt()
			if err != nil {
				return err
			}
			_tmp1.Number = _tmp10
			// GasLimit:
			_tmp11, err := dec.Uint64()
			if err != nil {
				return err
			}
			_tmp1.GasLimit = _tmp11
			// GasUsed:
			_tmp12, err := dec.Uint64()
			if err != nil {
				return err
			}
			_tmp1.GasUsed = _tmp12
			// Time:
			_tmp13, err := dec.Uint64()
			if err != nil {
				return err
			}
			_tmp1.Time = _tmp13
			// Extra:
			_tmp14, err := dec.Bytes()
			if err != nil {
				return err
			}
			_tmp1.Extra = _tmp14
			// MixDigest:
			var _tmp15 common.Hash
			if err := dec.ReadBytes(_tmp15[:]); err != nil {
				return err
			}
			_tmp1.MixDigest = _tmp15
			// Nonce:
			var _tmp16 types.BlockNonce
			if err := dec.ReadBytes(_tmp16[:]); err != nil {
				return err
			}
			_tmp1.Nonce = _tmp16
			// BaseFee:
			if dec.MoreDataInList() {
				_tmp17, err := dec.BigInt()
				if err != nil {
					return err
				}
				_tmp1.BaseFee = _tmp17
				// WithdrawalsHash:
				if dec.MoreDataInList() {
					var _tmp18 common.Hash
					if err := dec.ReadBytes(_tmp18[:]); err != nil {
						return err
					}
					_tmp1.WithdrawalsHash = &_tmp18
				}
			}
			if err := dec.ListEnd(); err != nil {
				return err
			}
		}
		_tmp0.Header = &_tmp1
		if err := dec.ListEnd(); err != nil {
			return err
		}
	}
	*obj = _tmp0
	return nil
}
