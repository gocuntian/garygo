# Introduction

#### 块链是二十一世纪最革命的技术之一，它还在成熟，哪些潜力尚未完全实现。 在本质上，块链只是一个分布式的记录数据库。 但是什么使它独一无二的是它不是一个私人数据库，而是一个公共的数据库，即每个使用它的人都有一个完整的或部分的副本。 只有经过其他数据库管理员的同意，才能添加新记录。 此外，它是使密码币和智能合同成为可能的块。

#### 在本系列文章中，我们将基于简单的块链实现来构建简化的加密机制。

# Block
#### 我们先从“blockchain”的“block”部分开始。 在块链中，它是存储有价值信息的块。 例如，比特币块存储事务，这是任何加密的实质。 此外，块包含一些技术信息，如其版本，当前时间戳和上一个块的散列。
#### 在本文中，我们不会像块链或比特币规范中所描述的那样实现块，而是使用它的简化版本，它只包含重要的信息。 这是它的样子：

    type Block struct {
        Timestamp     int64
        Data          []byte
        PrevBlockHash []byte
        Hash          []byte
    }

#### Timestamp 是当前的时间戳（创建块时）, Data 是块中包含的实际有价值的信息, PrevBlockHash 存储上一个块的散列, and Hash 是块的散列. 在比特币规格 Timestamp, PrevBlockHash, and Hash 是块头,它们形成一个单独的数据结构，而事务（Data在我们的例子中）是一个单独的数据结构。 所以我们在这里混合它们是为了简单。

#### 那么我们如何计算哈希呢？ 散列方式的计算是块链的非常重要的特征，而且这个功能使得blockchain安全。 事实是，计算散列是一个计算上困难的操作，即使在快速的计算机上也需要一些时间（这就是为什么人们购买强大的GPU来挖掘比特币）。 这是一个有意的建筑设计，这使得添加新的块变得困难，从而在添加新块之后阻止它们进行修改。 我们将在以后的文章中讨论并实施这一机制。

#### 现在，我们只需要采用块字段，连接它们，并在并置组合上计算一个SHA-256哈希值。 我们在SetHash方法中这样做：

    func (b *Block) SetHash() {
        timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
        headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
        hash := sha256.Sum256(headers)

        b.Hash = hash[:]
    }

#### 接下来，在Golang大会之后，我们将实现一个简化块创建的函数：

    func NewBlock(data string, prevBlockHash []byte) *Block {
        block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
        block.SetHash()
        return block
    }

# Blockchain

#### 现在我们来实现一个blockchain。 在本质上，块链只是一个具有一定结构的数据库：它是一个有序的，反向链接的列表。 这意味着块以插入顺序存储，并且每个块都链接到前一个块。 该结构允许快速获取链中的最新块，并通过其散列（有效地）获取块。

#### 在Golang中，这个结构可以通过使用数组和map来实现：数组将保持有序哈希（数组在Go中排序），映射将保持哈希→块对（映射无序）。 但是对于我们的blockchain原型，我们只需要使用一个数组，因为现在我们不需要通过hash来获取数据块

    type Blockchain struct {
        blocks []*Block
    }
#### 这是我们的第一个blockchain！ 我从来没有想过会这么容易😉

#### 现在让我们可以添加块：

    func (bc *Blockchain) AddBlock(data string) {
        prevBlock := bc.blocks[len(bc.blocks)-1]
        newBlock := NewBlock(data, prevBlock.Hash)
        bc.blocks = append(bc.blocks, newBlock)
    }

#### 要添加一个新块，我们需要一个现有的块，但是我们的块中没有块！ 所以，在任何块链中，必须至少有一个块，而这个块中的第一个被称为成因块。 我们来实现一个创建这样一个块的方法：

    func NewGenesisBlock() *Block {
        return NewBlock("Genesis Block", []byte{})
    }
#### 现在，我们可以实现一个创建一个具有起源阻塞的块的功能：

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

#### 让我们检查一下blockchain是否正常工作：

    func main() {
        bc := NewBlockchain()

        bc.AddBlock("Send 1 BTC to Ivan")
        bc.AddBlock("Send 2 more BTC to Ivan")

        for _, block := range bc.blocks {
            fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
            fmt.Printf("Data: %s\n", block.Data)
            fmt.Printf("Hash: %x\n", block.Hash)
            fmt.Println()
        }
    }

# Conclusion

#### 构建了一个非常简单的块链原型：它只是一个块数组，每个块都有一个到前一个的连接。 实际的块链更复杂。 在我们的块链中，添加新的块是容易和快速的，但是在实际的块链中添加新的块需要一些工作：在获得添加块的权限之前，必须执行一些重的计算（这个机制被称为工作证明）。 此外，块链是一个没有单个决策者的分布式数据库。 因此，一个新的块必须经过网络的其他参与者的确认和批准（这种机制称为共识）。 而且我们的块链还没有交易！