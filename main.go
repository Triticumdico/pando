package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

//Block is the standard Block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

//BlockChain is the standard BlockChain structure
type BlockChain struct {
	blocks []*Block
}

//DeriveHash is methode from Block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

//CreateBlock is the function used to Create a block
func CreateBlock(data string, PrevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), PrevHash}
	block.DeriveHash()
	return block
}

//AddBlock is the function used to add a crearted block
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// Genesis is the Blockchain starting block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain is the function used to Init the blockchaine
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

	}

}
