package blockchain

import (
	"DataCertPlatform/utils"
	"bytes"
	"time"
)

/*
定义区块结构体，用于表示区块
*/

type Block struct {
	Height int64 //区块的高度，第几个区块
	TimeStamp int64 //区块产生的时间戳
	PrevHash []byte //前一个区块hash
	Data []byte //数据字段
	Hash []byte //当前区块的hash值
	Version string //版本号
	Nonce int64 //区块对应的nonce值
}
/*
创建一个新区块
*/
func NewBlock(height int64,provHash []byte,data []byte) Block  {
	block:=Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  provHash,
		Data:      data,
		Version:   "0x01",
	}
	//1.将block结构体数据转换为[]byte类型
	heightBytes,_:=utils.In64ToByte(block.Height)
	timeStampBytes,_:=utils.In64ToByte(block.TimeStamp)
	versionBytes:=utils.StringToBytes(block.Version)

	var blockBytes []byte
	//bytes.join 拼接
	bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		block.PrevHash,
		block.Data,
		versionBytes,
	},[]byte{})

	block.Hash = utils.SHA256HashBlock(blockBytes)
	return block
}

/*
创建创世区块
*/

func CreateGenesisBlock() Block  {
	genesisBlock:=NewBlock(0,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},nil)
	return genesisBlock
}

