package common

import (
	"github.com/golang/protobuf/proto"

	"github.com/xuperchain/xupercore/bcs/ledger/xledger/xldgpb"
	ecom "github.com/xuperchain/xupercore/kernel/engines/xuperos/commom"
	"github.com/xuperchain/xuperos/common/xupospb/pb"
)

// 为了完全兼容老版本pb结构，转换交易结构
func TxToXledger(tx *pb.Transaction) *xldgpb.Transaction {
	if tx == nil {
		return nil
	}

	prtBuf, err := proto.Marshal(tx)
	if err != nil {
		return nil
	}

	var newTx xldgpb.Transaction
	err = proto.Unmarshal(prtBuf, &newTx)
	if err != nil {
		return nil
	}

	return &newTx
}

// 为了完全兼容老版本pb结构，转换交易结构
func TxToXchain(tx *xldgpb.Transaction) *pb.Transaction {
	if tx == nil {
		return nil
	}

	prtBuf, err := proto.Marshal(tx)
	if err != nil {
		return nil
	}

	var newTx pb.Transaction
	err = proto.Unmarshal(prtBuf, &newTx)
	if err != nil {
		return nil
	}

	return &newTx
}

// 为了完全兼容老版本pb结构，转换区块结构
func BlockToXledger(block *pb.InternalBlock) *xldgpb.InternalBlock {
	if block == nil {
		return nil
	}

	blkBuf, err := proto.Marshal(block)
	if err != nil {
		return nil
	}

	var newBlock xldgpb.InternalBlock
	err = proto.Unmarshal(blkBuf, &newBlock)
	if err != nil {
		return nil
	}

	return &newBlock
}

// 为了完全兼容老版本pb结构，转换区块结构
func BlockToXchain(block *xldgpb.InternalBlock) *pb.InternalBlock {
	if block == nil {
		return nil
	}

	blkBuf, err := proto.Marshal(block)
	if err != nil {
		return nil
	}

	var newBlock pb.InternalBlock
	err = proto.Unmarshal(blkBuf, &newBlock)
	if err != nil {
		return nil
	}

	return &newBlock
}

func ConvertInvokeReq(reqs []*InvokeRequest) ([]*protos.InvokeRequest, error) {
	if reqs == nil {
		return nil
	}

	newReqs := make([]*protos.InvokeRequest, len(reqs))
	for _, req := range reqs {
		buf, err := proto.Marshal(req)
		if err != nil {
			return nil, err
		}

		var tmp protos.InvokeRequest
		err = proto.Unmarshal(buf, &tmp)
		if err != nil {
			return nil, err
		}

		newReqs = append(newReqs, &tmp)
	}

	return newReqs, nil
}

func ConvertInvokeResp(resp *protos.InvokeResponse) *pb.InvokeResponse {
	if resp == nil {
		return nil
	}

	buf, err := proto.Marshal(resp)
	if err != nil {
		return nil
	}

	var tmp pb.InvokeResponse
	err = proto.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return &tmp
}

func UtxoToXchain(utxo *xldgpb.Utxo) *pb.Utxo {
	if utxo == nil {
		return nil
	}

	buf, err := proto.Marshal(utxo)
	if err != nil {
		return nil
	}

	var tmp pb.Utxo
	err = proto.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return &tmp
}

func UtxoToXledger(utxo *pb.Utxo) *xldgpb.Utxo {
	if utxo == nil {
		return nil
	}

	buf, err := proto.Marshal(utxo)
	if err != nil {
		return nil
	}

	var tmp xldgpb.Utxo
	err = proto.Unmarshal(buf, &tmp)
	if err != nil {
		return nil
	}

	return &tmp
}

func UtxoListToXchain(utxoList []*xldgpb.Utxo) ([]*pb.Utxo, error) {
	if utxoList == nil {
		return nil, nil
	}

	tmpList := make([]*pb.Utxo, len(utxoList))
	for _, utxo := range utxoList {
		tmp := UtxoToXchain(utxo)
		if tmp == nil {
			return nil, fmt.Errorf("convert utxo failed")
		}
		tmpList = append(tmpList, tmp)
	}

	return tmpList, nil
}