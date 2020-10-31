package blockchain

import (
	"DataCertPlatform/utils"
	"bytes"
	"math/big"
)

/*
工作量证明算法结构体
*/
const DIFFICULTY  = 6

type ProofOfWork struct {
	Target *big.Int //系统的目标值
	Block Block //要找的nonce值对应的区块
}

func NewPow(block Block) ProofOfWork{
	t:=big.NewInt(1)
	t=t.Lsh(t,255-DIFFICULTY)
	pow:=ProofOfWork{
		Target:t,
		Block:block,
	}
	return pow
}

func (p ProofOfWork) Run() ([]byte,int64){
	var nonce int64
	nonce=0
	var blockHash []byte
	for  {
	block:=p.Block
	heightBytes,_:=utils.In64ToByte(block.Height)
	timeStampBytes,_:=utils.In64ToByte(block.TimeStamp)
	versionBytes:=utils.StringToBytes(block.Version)
	nonceBytes,_:=utils.In64ToByte(nonce)
	blockBytes:=bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		block.PrevHash,
		block.Data,
		versionBytes,
		nonceBytes,
	},[]byte{})

	blockHash =utils.SHA256HashBlock(blockBytes)
	target:=p.Target
	var hashBig *big.Int
	hashBig=new(big.Int)
	hashBig=hashBig.SetBytes(blockHash)
	//fmt.Println("当前尝试的Nonce值：",nonce)
	if hashBig.Cmp(target)==-1 {
		break
	}
	nonce++
 }
  return blockHash,nonce
}
