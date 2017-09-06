# Introduction

#### 我们已经建立了一个带有工作证明的系统，这使得挖掘成为可能。 我们的实现越来越接近一个功能完整的块，但它仍然缺乏一些重要的功能。 今天将开始将数据库存储在数据库中，之后我们将使用一个简单的命令行界面来执行与块链的操作。 其实质是块分布式数据库。 我们现在将忽略“分布式”部分，并专注于“数据库”部分。

# Database Choice
#### 目前，我们的实施没有数据库; 相反，我们每次运行程序并将它们存储在内存中时都会创建块。 我们不能重用块链接，我们不能与他人共享，因此我们需要将其存储在磁盘上。

#### 我们需要哪个数据库？ 其实，其中任何一个。 在原始的比特币纸上，没有任何关于使用某个数据库的说法，所以由开发人员使用哪个数据库。 Bitcoin Core最初由Satoshi Nakamoto出版，目前是Bitcoin的参考实施，使用LevelDB（尽管仅在2012年被引入客户端）。 我们会使用...



# BoltDB
#####  1.简单而简约。
#####  2.它在Go中实现。
#####  3.它不需要运行服务器。
#####  4.它允许构建我们想要的数据结构。
## https://github.com/boltdb/bolt

#### Bolt是一个纯粹的Go钥匙/价值商店，灵感来自于Howard Chu的LMDB项目。 该项目的目标是为不需要完整数据库服务器（如Postgres或MySQL）的项目提供一个简单，快速，可靠的数据库。由于Bolt意在被用作这样一个低级别的功能，简单性就是关键。 API会很小，只关注获取值和设置值。 而已。

#### BoltDB是一个键/值存储，这意味着没有像SQL RDBMS（MySQL，PostgreSQL等）中的表，没有行，没有列。 相反，数据作为键值对存储（如在Golang地图中）。 键值对存储在桶中，用于对相似的对进行分组（这与RDBMS中的表类似）。 因此，为了获得价值，您需要知道一个桶和一个钥匙。

#### BoltDB的一个重要事情是没有数据类型：键和值是字节数组。 由于我们将存储Go结构（特别是Block），我们需要对它们进行序列化，即实现将Go结构转换为字节数组并将其从字节数组恢复的机制。 我们将使用encoding / gob，但也可以使用JSON，XML，协议缓冲区等。 我们使用编码/ gob，因为它很简单，是标准Go库的一部分。



# Database Structure

## 在开始执行持久性逻辑之前，我们首先需要决定如何在数据库中存储数据。 为此，我们将参考比特币核心的方式。

## 简单来说，Bitcoin Core使用两个“桶”来存储数据：
####    1.块存储描述链中所有块的元数据。
####    2.chainstate存储一个链的状态，这是目前所有的未用的交易输出和一些元数据。

### 另外，块作为单独的文件存储在磁盘上。 这是为了表现目的：读取单个块不需要将所有（或其中一些）加载到内存中。 我们不会实现这一点。




## 在块中，键 - >值对是：

####    1. 'b' + 32-byte block hash -> 块索引记录
####    2. 'f' + 4-byte file number -> 文件信息记录
####    3. 'l' -> 4-byte file number: 最后使用的块文件号
####    4. 'R' -> 1-byte boolean: 我们是否正在重建索引
####    5. 'F' + 1-byte flag name length + flag name string -> 1 byte boolean: 可以打开或关闭的各种标志
####    6. 't' + 32-byte transaction hash ->事务索引记录
## 在chainstate中，键 - >值对是：

####    1.'c' + 32-byte transaction hash -> 该事务的未经销事务输出记录
####    2.'B' -> 32-byte block hash: 数据库表示未用量事务输出的块哈希
## https://en.bitcoin.it/wiki/Bitcoin_Core_0.11_(ch_2):_Data_Storage

## 由于我们还没有交易，所以我们只有块数据。 另外，如上所述，我们将整个数据库存储为单个文件，而不将块存储在单独的文件中。 所以我们不需要任何与文件号相关的任何东西。 所以这些是我们将使用的关键 - >值对：
####    1. 32-byte block-hash -> 块结构（序列化）
####    2.'l' -> 链中最后一个块的哈希值
## 我们需要知道的是开始实现持久性机制。

# Serialization

## 如前所述，BoltDB中的值只能是[]字节类型，我们要将块结构存储在数据库中。 我们将使用encoding / gob来序列化结构体。

## 我们来实现Block的Serialize方法（为简洁起见，省略错误处理）：
    func (b *Block) Serialize() []byte {
        var result bytes.Buffer
        encoder := gob.NewEncoder(&result)

        err := encoder.Encode(b)

        return result.Bytes()
    }

####这一切很简单：首先，我们声明一个缓冲区，将存储序列化数据; 然后我们初始化一个gob编码器并对该块进行编码; 结果作为字节数组返回。接下来，我们需要一个反序列化函数，它将接收一个字节数组作为输入并返回一个块。 这不是一个方法，而是一个独立的功能：

    func DeserializeBlock(d []byte) *Block {
        var block Block

        decoder := gob.NewDecoder(bytes.NewReader(d))
        err := decoder.Decode(&block)

        return &block
    }

#### 这就是序列化！

## Persistence

### 我们从NewBlockchain函数开始。 目前，它创建了一个新的Blockchain实例，并向其添加成因块。 我们想要做的是：

####    1.打开一个DB文件。
####    2.检查是否存在一个块链。
####    3.如果有一个blockchain：
#####     a.创建一个新的Blockchain实例。
#####     b.将Blockchain实例的提示设置为DB中存储的最后一个块哈希。
####    4.如果没有现有的blockchain：
#####     a.创造起源障碍。
#####     b.存储在数据库中。
#####     c.将起源块的哈希保存为最后一个块哈希。
#####     d.创建一个新的Blockchain实例，其提示指向起源块。
#####     e.在代码中，它看起来像这样：

    func NewBlockchain() *Blockchain {
        var tip []byte
        db, err := bolt.Open(dbFile, 0600, nil)

        err = db.Update(func(tx *bolt.Tx) error {
            b := tx.Bucket([]byte(blocksBucket))

            if b == nil {
                genesis := NewGenesisBlock()
                b, err := tx.CreateBucket([]byte(blocksBucket))
                err = b.Put(genesis.Hash, genesis.Serialize())
                err = b.Put([]byte("l"), genesis.Hash)
                tip = genesis.Hash
            } else {
                tip = b.Get([]byte("l"))
            }

            return nil
        })

        bc := Blockchain{tip, db}

        return &bc
    }

## 审查一下。
    db, err := bolt.Open(dbFile, 0600, nil)
## 这是打开BoltDB文件的标准方式。 请注意，如果没有这样的文件，它不会返回错误。
    err = db.Update(func(tx *bolt.Tx) error {
    ...
    })
## 在BoltDB中，具有数据库的操作在事务中运行。 并且有两种类型的事务：只读和读写。 在这里，我们打开一个读写事务（db.Update（...）），因为我们期望把数据块中的成因块。
    
    b := tx.Bucket([]byte(blocksBucket))

    if b == nil {
        genesis := NewGenesisBlock()
        b, err := tx.CreateBucket([]byte(blocksBucket))
        err = b.Put(genesis.Hash, genesis.Serialize())
        err = b.Put([]byte("l"), genesis.Hash)
        tip = genesis.Hash
    } else {
        tip = b.Get([]byte("l"))
    }

## 这是功能的核心。 在这里，我们获取存储我们的块的存储桶：如果存在，我们从它读取l个密钥; 如果不存在，我们生成成因块，创建存储块，将块保存到其中，并更新存储链的最后块哈希的l密钥。

## 另外，请注意创建Blockchain的新方法：

    bc := Blockchain{tip, db}

## 我们不存储所有的块，而只存储链的末端。 另外，我们存储一个数据库连接，因为我们想打开它一次，并在程序运行时保持打开。 因此，Blockchain结构现在看起来像这样：

    type Blockchain struct {
        tip []byte
        db  *bolt.DB
    }
## 接下来我们要更新的是AdBlock方法：将块添加到链中并不像将数组添加到数组那样简单。 从现在开始，我们将在DB中存储块：

    func (bc *Blockchain) AddBlock(data string) {
        var lastHash []byte

        err := bc.db.View(func(tx *bolt.Tx) error {
            b := tx.Bucket([]byte(blocksBucket))
            lastHash = b.Get([]byte("l"))

            return nil
        })

        newBlock := NewBlock(data, lastHash)

        err = bc.db.Update(func(tx *bolt.Tx) error {
            b := tx.Bucket([]byte(blocksBucket))
            err := b.Put(newBlock.Hash, newBlock.Serialize())
            err = b.Put([]byte("l"), newBlock.Hash)
            bc.tip = newBlock.Hash

            return nil
        })
    }


    err := bc.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(blocksBucket))
        lastHash = b.Get([]byte("l"))

        return nil
    })

## 这是其他（只读）类型的BoltDB事务。 这里我们从DB获取最后一个块哈希值，以使用它来挖掘一个新的块哈希。
    newBlock := NewBlock(data, lastHash)
    b := tx.Bucket([]byte(blocksBucket))
    err := b.Put(newBlock.Hash, newBlock.Serialize())
    err = b.Put([]byte("l"), newBlock.Hash)
    bc.tip = newBlock.Hash
## 在挖掘一个新的块后，我们将其序列化表示保存到DB中，并更新l密钥，该密钥现在存储新块的哈希值。

# Inspecting Blockchain

### 所有新的块现在都保存在数据库中，所以我们可以重新打开一个块，并添加一个新的块。 但是在实现之后，我们失去了一个很好的功能：我们不能再打印出块链接块了，因为我们不再在数组中存储块了。 我们来解决这个缺陷！

### BoltDB允许遍历桶中的所有密钥，但密钥按字节排序的顺序存储，我们希望以按照块的顺序打印块。 另外，因为我们不想将所有的块加载到内存中（我们的blockchain数据库可能是巨大的...或者让我们假装可以），我们将逐个读取它们。 为此，我们需要一个blockchain迭代器：

    type BlockchainIterator struct {
        currentHash []byte
        db          *bolt.DB
    }

### 每次我们想要迭代块中的块时，将创建一个迭代器，它将存储当前迭代的块哈希值和与DB的连接。 由于后者，迭代器在逻辑上附加到块链（它是一个存储数据库连接的Blockchain实例），因此，以Blockchain方法创建：
    func (bc *Blockchain) Iterator() *BlockchainIterator {
        bci := &BlockchainIterator{bc.tip, bc.db}

        return bci
    }
### 请注意，迭代器最初指向块链的末尾，因此块将从上到下从最新到最后。 事实上，选择一个提示意味着“投票”为一个块。 一个块链可以有多个分支，它们是最长的，被认为是主要的。 得到一个提示（它可以是块链中的任何块），我们可以重构整个块链，并找到其长度和构建它所需的工作。 这个事实也意味着提示是一种块链的标识符。
### BlockchainIterator只会做一件事：它会从块链中返回下一个块。
    func (i *BlockchainIterator) Next() *Block {
        var block *Block

        err := i.db.View(func(tx *bolt.Tx) error {
            b := tx.Bucket([]byte(blocksBucket))
            encodedBlock := b.Get(i.currentHash)
            block = DeserializeBlock(encodedBlock)

            return nil
        })

        i.currentHash = block.PrevBlockHash

        return block
    }

# CLI
#### 到目前为止，我们的实现没有提供任何接口来与程序交互：我们在主函数中简单地执行了NewBlockchain，bc.AddBlock。 时间来改善这个！ 我们想要这些命令：
### blockchain go adblock“支付0.031337一杯咖啡”
### 块链打印链
### 所有命令行相关操作将由CLI结构处理：

    type CLI struct {
        bc *Blockchain
    }

###  它的“entrypoint”是运行功能：

    func (cli *CLI) Run() {
        cli.validateArgs()

        addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
        printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

        addBlockData := addBlockCmd.String("data", "", "Block data")

        switch os.Args[1] {
        case "addblock":
            err := addBlockCmd.Parse(os.Args[2:])
        case "printchain":
            err := printChainCmd.Parse(os.Args[2:])
        default:
            cli.printUsage()
            os.Exit(1)
        }

        if addBlockCmd.Parsed() {
            if *addBlockData == "" {
                addBlockCmd.Usage()
                os.Exit(1)
            }
            cli.addBlock(*addBlockData)
        }

        if printChainCmd.Parsed() {
            cli.printChain()
        }
    }
### 我们使用标准标记包来解析命令行参数。

    addBlockCmd：= flag.NewFlagSet（“addblock”，flag.ExitOnError）
    printChainCmd：= flag.NewFlagSet（“printchain”，flag.ExitOnError）
    addBlockData：= addBlockCmd.String（“data”，“”，“Block data”）

### 首先，我们创建两个子命令，addblock和printchain，然后我们向前者添加-data标志。 printchain不会有任何标志。

    switch os.Args[1] {
    case "addblock":
        err := addBlockCmd.Parse(os.Args[2:])
    case "printchain":
        err := printChainCmd.Parse(os.Args[2:])
    default:
        cli.printUsage()
        os.Exit(1)
    }

### 接下来我们检查用户提供的命令和解析相关的标志子命令。

    if addBlockCmd.Parsed() {
        if *addBlockData == "" {
            addBlockCmd.Usage()
            os.Exit(1)
        }
        cli.addBlock(*addBlockData)
    }

    if printChainCmd.Parsed() {
        cli.printChain()
    }

### 接下来，我们检查哪些子命令已被解析并运行相关函数。

    func (cli *CLI) addBlock(data string) {
        cli.bc.AddBlock(data)
        fmt.Println("Success!")
    }

    func (cli *CLI) printChain() {
        bci := cli.bc.Iterator()

        for {
            block := bci.Next()

            fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
            fmt.Printf("Data: %s\n", block.Data)
            fmt.Printf("Hash: %x\n", block.Hash)
            pow := NewProofOfWork(block)
            fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
            fmt.Println()

            if len(block.PrevBlockHash) == 0 {
                break
            }
        }
    }

### 这件事与我们以前一样非常相似。 唯一的区别是我们现在正在使用BlockchainIterator来迭代块中的块。
### 也不要忘了修改主要功能：

    func main() {
        bc := NewBlockchain()
        defer bc.db.Close()

        cli := CLI{bc}
        cli.Run()
    }
### 请注意，无论提供什么命令行参数，都会创建一个新的Blockchain。

### 就是这样！ 让我们检查一切是否符合预期：

##### $ blockchain_go printchain
##### No existing blockchain found. Creating a new one...
##### Mining the block containing "Genesis Block"
##### 000000edc4a82659cebf087adee1ea353bd57fcd59927662cd5ff1c4f618109b

##### Prev. hash:
##### Data: Genesis Block
##### Hash: 000000edc4a82659cebf087adee1ea353bd57fcd59927662cd5ff1c4f618109b
##### PoW: true

##### $ blockchain_go addblock -data "Send 1 BTC to Ivan"
##### Mining the block containing "Send 1 BTC to Ivan"
##### 000000d7b0c76e1001cdc1fc866b95a481d23f3027d86901eaeb77ae6d002b13

##### Success!

##### $ blockchain_go addblock -data "Pay 0.31337 BTC for a coffee"
##### Mining the block containing "Pay 0.31337 BTC for a coffee"
##### 000000aa0748da7367dec6b9de5027f4fae0963df89ff39d8f20fd7299307148

##### Success!

##### $ blockchain_go printchain
##### Prev. hash: 000000d7b0c76e1001cdc1fc866b95a481d23f3027d86901eaeb77ae6d002b13
##### Data: Pay 0.31337 BTC for a coffee
##### Hash: 000000aa0748da7367dec6b9de5027f4fae0963df89ff39d8f20fd7299307148
##### PoW: true

##### Prev. hash: 000000edc4a82659cebf087adee1ea353bd57fcd59927662cd5ff1c4f618109b
##### Data: Send 1 BTC to Ivan
##### Hash: 000000d7b0c76e1001cdc1fc866b95a481d23f3027d86901eaeb77ae6d002b13
##### PoW: true

##### Prev. hash:
##### Data: Genesis Block
##### Hash: 000000edc4a82659cebf087adee1ea353bd57fcd59927662cd5ff1c4f618109b
##### PoW: true
##### (sound of a beer can opening)