package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"time"
	"strconv"

	"github.com/davecgh/go-spew/spew"  // Spew 可以格式化 structs 和 slices，以便我们在console能清晰明了的看这些数据
	"github.com/joho/godotenv"         // 可以让我们读取 .env 文件里面的环境变量
	"bufio"
)

type Block struct {
	Index     int      // 区块在区块链里的位置，即索引
	Timestamp string   // 产生区块的时间戳
	Data       int     // 待写入的数据
	PrevHash  string   // 前一个区块 SHA256 的哈希值
	Hash      string   // 当前区块 SHA256 的哈希值
}
var Blockchain []Block
var bcServer chan []Block

// 主函数入口
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bcServer = make(chan []Block)

	t := time.Now()
	genesisBlock := Block{0, t.String(), 0, "", ""}
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	// 启动TCP
	server, err := net.Listen("tcp", ":"+os.Getenv("ADDR"))
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "Enter a new Data:")
	scanner := bufio.NewScanner(conn)

	go func() {
		for scanner.Scan() {
			data, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("%v not a number: %v", scanner.Text(), err)
				continue
			}
			newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], data)
			if err != nil {
				log.Println(err)
				continue
			}
			if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
				newBlockchain := append(Blockchain, newBlock)
				replaceChain(newBlockchain)
			}

			bcServer <- Blockchain
			io.WriteString(conn, "\nEnter a new Data:")
		}
	}()

	// 模拟接收广播
	go func() {
		for {
			time.Sleep(30 * time.Second)
			output, err := json.Marshal(Blockchain)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(conn, string(output))
		}
	}()

	for _ = range bcServer {
		spew.Dump(Blockchain)
	}
}

// 维持最长的链为合法链
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

// 验证新产生的区块是否合法
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// 计算区块数据的哈希值
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.Data) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

// 生成新的区块
func generateBlock(oldBlock Block, Data int) (Block, error) {
	var newBlock Block
	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = Data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}
