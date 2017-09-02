# Introduction

#### 构建了一个非常简单的数据结构，这是块链数据库的本质。 而且我们可以用它们之间的链式关系向它添加块：每个块与前一个链接。 唉，我们的blockchain实现有一个很大的缺陷：添加块到链是容易和便宜的。 块链和比特币的关键之一是添加新块是一项艰巨的工作。 今天我们要解决这个缺陷。

# Proof-of-Work
#### 块链的一个关键思想是，必须执行一些努力才能将数据放入其中。这是一个艰巨的工作，使得blockchain安全和一致。此外，为了这项辛苦工作也付出了回报（这是人们获得采矿硬币的方式）。

#### 这种机制与现实生活中的机制非常相似：人们必须努力获得奖励和维持生命。在网络中，网络的一些参与者（矿工）努力维持网络，为其添加新的块，并为他们的工作获得奖励。作为其工作的结果，块以安全的方式并入到块链中，这保持了整个块链数据库的稳定性。值得注意的是，完成工作的人必须证明这一点。

#### 这个整体“努力工作和证明”机制被称为工作证明。这很难，因为它需要很多的计算能力：即使是高性能的计算机也不能很快的完成。此外，这项工作的难度不时增加，以保持新的块率每小时大约6个块。在比特币，这样的工作的目标是找到一个块的哈希，满足一些要求。这是散列，作为一个证明。因此，找到证据是实际工作。

#### 最后一件事要注意。工作证明算法必须满足要求：做工很难，但验证证明很容易。证明通常交给别人，所以对他们来说，验证它不应该花费太多时间。

# Hashing

## 在本段中，我们将讨论散列。 如果你熟悉这个概念，你可以跳过这个部分。

### 哈希是获取指定数据的哈希的过程。 哈希是对其计算的数据的唯一表示。 哈希函数是一个获取任意大小的数据并产生固定大小的哈希的函数。 以下是哈希的一些主要功能：

#####  1.原始数据无法从哈希恢复。 因此，散列不是加密。
#####  2.某些数据可以只有一个散列，散列是唯一的。
#####  3.更改输入数据中的一个字节将导致完全不同的散列。


#### 哈希函数被广泛用于检查数据的一致性。 一些软件提供商除了软件包外还发布校验和。 下载文件后，您可以将其提供给哈希函数，并将生成的哈希与软件开发人员提供的哈希进行比较。

#### 在块链中，使用散列来保证块的一致性。 哈希算法的输入数据包含前一个块的散列，从而使得链中修改块无法（或至少相当困难）：必须重新计算其中的所有块的哈希和散列。

# Hashcash

## 比特币使用Hashcash，最初是为防止电子邮件垃圾邮件而开发的工作证明算法。它可以分为以下几个步骤：

### 拿一些公开的数据（在电子邮件的情况下，它是接收方的电子邮件地址;如果是Bitcoin，它是块头）。
####    1.添加一个计数器。计数器从0开始。
####    2.获取数据+计数器组合的散列。
####    3.检查哈希符合某些要求。
#####      A.如果这样做，你就完成了。
#####      B.如果没有，增加计数器并重复步骤3和4。
### 因此，这是一个强力算法：您更改计数器，计算一个新的哈希，检查它，增加计数器，计算哈希等。这就是为什么它的计算昂贵。

### 现在让我们看看一个哈希必须满足的要求。在原来的Hashcash实现中，这个要求听起来像“哈希的前20位必须是零”。在比特币中，要求是不时地进行调整的，因为按照设计，尽管计算能力随着时间的推移而增加，越来越多的矿工加入网络，必须每10分钟生成一个块。

### 为了演示这个算法，我从前面的例子（“我喜欢甜甜圈”）中获取了数据，并发现一个以0个零字节开头的哈希：
### ca07ca是计数器的十六进制值，是十进制的13240266。

# Implementation

## 完成了理论，让我们编写代码！ 首先，我们来定义挖掘的难度：

##### const targetBits = 24

### 在比特币中，“目标位”是存储块被挖掘的困难的块头。 我们现在不会实现目标调整算法，所以我们可以将难度定义为全局常数。
### 24是一个任意数字，我们的目标是在内存中占用少于256位的目标。 而且我们希望差异足够大，但不能太大，因为差异越大，找到合适的哈希越难。

    type ProofOfWork struct {
        block  *Block
        target *big.Int
    }

    func NewProofOfWork(b *Block) *ProofOfWork {
        target := big.NewInt(1)
        target.Lsh(target, uint(256-targetBits))

        pow := &ProofOfWork{b, target}

        return pow
    }
#### 这里创建一个ProofOfWork结构，它保存一个指向一个块的指针和一个指向目标的指针。 “目标”是上一段所述要求的另一个名称。 我们使用一个大整数，因为我们将哈希与目标进行比较：我们将一个哈希转换为一个大整数，并检查它是否小于目标。

#### 在NewProofOfWork函数中，我们初始化一个值为1的big.Int，并将其左移256个 - targetBits位。 256是SHA-256哈希的长度，以比特为单位，它是我们要使用的SHA-256散列算法。 目标的十六进制表示为：

##### 0x10000000000000000000000000000000000000000000000000000000000
#### 它在内存中占用29个字节。 这里与前面的例子中的散列表进行了视觉比较：

#####    0fac49161af82ed938add1d8725835cc123a1a87b1b196488360e58d4bfb51e3
#####    0000010000000000000000000000000000000000000000000000000000000000
#####    0000008b0f41ec78bab747864db66bcb9fb89920ee75f43fdaaeb5544f7f76ca


### 第一个哈希（以“我喜欢甜甜圈”计算）大于目标，因此它不是有效的工作证明。 第二个哈希（以“我喜欢donutsca07ca”计算）比目标小，因此这是一个有效的证明。

### 您可以将目标视为范围的上限：如果数字（散列）低于边界，则它是有效的，反之亦然。 降低边界将导致有效数量减少，因此找到有效数量所需的工作更加困难。

### 现在，我们需要数据进行散列。 我们来准备一下

    func (pow *ProofOfWork) prepareData(nonce int) []byte {
        data := bytes.Join(
            [][]byte{
                pow.block.PrevBlockHash,
                pow.block.Data,
                IntToHex(pow.block.Timestamp),
                IntToHex(int64(targetBits)),
                IntToHex(int64(nonce)),
            },
            []byte{},
        )

        return data
    }
### 这件事是直截了当的：我们只是将块区与目标和随机数合并。 nonce这里是从上面的Hashcash描述的计数器，这是一个加密术语。

### 好的，所有的准备工作都完成了，我们来实现PoW算法的核心：

    func (pow *ProofOfWork) Run() (int, []byte) {
        var hashInt big.Int
        var hash [32]byte
        nonce := 0

        fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
        for nonce < maxNonce {
            data := pow.prepareData(nonce)
            hash = sha256.Sum256(data)
            fmt.Printf("\r%x", hash)
            hashInt.SetBytes(hash[:])

            if hashInt.Cmp(pow.target) == -1 {
                break
            } else {
                nonce++
            }
        }
        fmt.Print("\n\n")

        return nonce, hash[:]
    }

### 首先，我们初始化变量：hashInt是哈希的整数表示; nonce是柜台。 接下来，我们运行一个“无限”循环：它受限于maxNonce，它等于math.MaxInt64; 这样做是为了避免可能的随机数溢出。 虽然我们的PoW实现的难度太低，以至于防止溢出，但是更好的是进行此检查，以防万一。

##在循环中我们：

   #### 1.准备数据
   #### 2.用SHA-256进行哈希。
   #### 3.将散列转换为大整数。
   #### 4 将整数与目标进行比较。
### 就像前面说过的那样简单。 现在我们可以删除Block的SetHash方法并修改NewBlock函数：

    func NewBlock(data string, prevBlockHash []byte) *Block {
        block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
        pow := NewProofOfWork(block)
        nonce, hash := pow.Run()

        block.Hash = hash[:]
        block.Nonce = nonce

        return block
    }

### 在这里，您可以看到该随机数被保存为一个块属性。 这是必要的，因为需要验证证据的是nonce。 块结构现在看起来如此：

    type Block struct {
        Timestamp     int64
        Data          []byte
        PrevBlockHash []byte
        Hash          []byte
        Nonce         int
    }
### 让我们来运行程序，看看一切是否正常工作：

##### Mining the block containing "Genesis Block"
##### 00000041662c5fc2883535dc19ba8a33ac993b535da9899e593ff98e1eda56a1

##### Mining the block containing "Send 1 BTC to Ivan"
##### 00000077a856e697c69833d9effb6bdad54c730a98d674f73c0b30020cc82804

##### Mining the block containing "Send 2 more BTC to Ivan"
##### 000000b33185e927c9a989cc7d5aaaed739c56dad9fd9361dea558b9bfaf5fbe

##### Prev. hash:
##### Data: Genesis Block
##### Hash: 00000041662c5fc2883535dc19ba8a33ac993b535da9899e593ff98e1eda56a1

##### Prev. hash: 00000041662c5fc2883535dc19ba8a33ac993b535da9899e593ff98e1eda56a1
##### Data: Send 1 BTC to Ivan
##### Hash: 00000077a856e697c69833d9effb6bdad54c730a98d674f73c0b30020cc82804

##### Prev. hash: 00000077a856e697c69833d9effb6bdad54c730a98d674f73c0b30020cc82804
##### Data: Send 2 more BTC to Ivan
##### Hash: 000000b33185e927c9a989cc7d5aaaed739c56dad9fd9361dea558b9bfaf5fbe
### 可以看到，每个哈希现在以三个零字节开始，并且需要一些时间才能获得这些散列。

### 还有一件事要做：让我们可以验证作品的证明。

    func (pow *ProofOfWork) Validate() bool {
        var hashInt big.Int

        data := pow.prepareData(pow.block.Nonce)
        hash := sha256.Sum256(data)
        hashInt.SetBytes(hash[:])

        isValid := hashInt.Cmp(pow.target) == -1

        return isValid
    }
#### 这就是我们需要保存的随机数。

#### 让我们再检查一下，一切正常：

    func main() {
    ...

    for _, block := range bc.blocks {
        ...
        pow := NewProofOfWork(block)
        fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()
    }
    }


#### Output:

...

##### Prev. hash:
##### Data: Genesis Block
##### Hash: 00000093253acb814afb942e652a84a8f245069a67b5eaa709df8ac612075038
##### PoW: true

##### Prev. hash: 00000093253acb814afb942e652a84a8f245069a67b5eaa709df8ac612075038
##### Data: Send 1 BTC to Ivan
##### Hash: 0000003eeb3743ee42020e4a15262fd110a72823d804ce8e49643b5fd9d1062b
##### PoW: true

##### Prev. hash: 0000003eeb3743ee42020e4a15262fd110a72823d804ce8e49643b5fd9d1062b
##### Data: Send 2 more BTC to Ivan
##### Hash: 000000e42afddf57a3daa11b43b2e0923f23e894f96d1f24bfd9b8d2d494c57a
##### PoW: true

# Conclusion

### 我们的块链是一个更接近其实际架构的一步：添加块现在需要努力工作，因此挖掘是可能的。 但是它仍然缺乏一些关键特征：块链数据库不是持久的，没有钱包，地址，交易，没有共识机制。 所有这些我们将在以后的文章中实现的，现在，开采开采！