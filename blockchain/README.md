此程序尝试使用Go语言实现一个简易的区块链，一是为了更好的理解区块链，二是为了更好的掌握Go语言的实现，仅供参考。

我们假设你的设备已经具备了Go语言的开发环境。

    go get github.com/davecgh/go-spew/spew
Spew 可以格式化 structs 和 slices，以便我们在console能清晰明了的看这些数据。

    go get github.com/gorilla/mux
Gorilla/mux 是一个很流行的处理http请求的包

    go get github.com/joho/godotenv
Godotenv 可以让我们读取 .env 文件里面的环境变量。所以我们在根目录创建一个 .env 文件，写入下面内容:PORT=8080

在第一个终端，运行程序

    go run main.go

新开第二个终端，运行：

    nc localhost 9000
输入欲传递的数，第一个终端便显示第二个交易，一段时间（30s）此交易会广播到其他终端中。