package blockchain

//Block is the standard Block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

//BlockChain is the standard BlockChain structure
type BlockChain struct {
	Blocks []*Block
}

//CreateBlock is the function used to Create a block
func CreateBlock(data string, PrevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), PrevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//AddBlock is the function used to add a crearted block
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis is the Blockchain starting block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain is the function used to Init the blockchaine
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
