## go tools

```bash
# 将一个或多个 .go 文件进行编译、链接然后运行生成的可执行文件
$ go run helloworld.go 

# 编译生成一个可执行二进制文件
$ go build helloworld.go
```

```bash
# 格式化包里的所有文件
$ gofmt

# 下载包
$ go get golang.org/x/tools/cmd/goimports
```

```bash
# 单元测试 & 测试覆盖率
$ go test -v -cover

# 性能测试
$ go test -bench=. -benchmem
```

```bash
$ sudo apt-get install graphviz
$ go build prof.go
$ ./prof

$ go tool pprof prof cpu.prof
File: prof
Type: cpu
Time: Dec 6, 2021 at 1:53am (CST)
Duration: 1.31s, Total samples = 1.14s (87.29%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 1.13s, 99.12% of 1.14s total
Showing top 10 nodes out of 27
      flat  flat%   sum%        cum   cum%
     0.42s 36.84% 36.84%      0.42s 36.84%  runtime.memclrNoHeapPointers
     0.39s 34.21% 71.05%      0.52s 45.61%  math/rand.(*Rand).Int31n
     0.08s  7.02% 78.07%      0.60s 52.63%  math/rand.(*Rand).Intn
     0.06s  5.26% 83.33%      0.66s 57.89%  main.fillMatrix
     0.06s  5.26% 88.60%      0.09s  7.89%  math/rand.(*rngSource).Int63
     0.03s  2.63% 91.23%      0.12s 10.53%  math/rand.(*Rand).Int63 (inline)
     0.03s  2.63% 93.86%      0.03s  2.63%  math/rand.(*rngSource).Uint64 (inline)
     0.03s  2.63% 96.49%      0.03s  2.63%  runtime.futex
     0.02s  1.75% 98.25%      0.02s  1.75%  main.calculate
     0.01s  0.88% 99.12%      0.13s 11.40%  math/rand.(*Rand).Int31 (inline)
(pprof) list fillMatrix
Total: 1.14s
ROUTINE ======================== main.fillMatrix in /home/mark/Dev/golang-dev/tutorial.com/41-pprof/prof.go
      60ms      660ms (flat, cum) 57.89% of Total
         .          .     15:
         .          .     16:func fillMatrix(m *[row][col]int) {
         .          .     17:	s := rand.New(rand.NewSource(time.Now().UnixNano()))
         .          .     18:
         .          .     19:	for i := 0; i < row; i++ {
      20ms       20ms     20:		for j := 0; j < col; j++ {
      40ms      640ms     21:			m[i][j] = s.Intn(10000)
         .          .     22:		}
         .          .     23:	}
         .          .     24:}
         .          .     25:
         .          .     26:func calculate(m *[row][col]int) {
(pprof) svg
Generating report in profile001.svg

$ go tool pprof prof mem.prof
File: prof
Type: inuse_space
Time: Dec 6, 2021 at 1:53am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 762.95MB, 99.32% of 768.17MB total
Dropped 22 nodes (cum <= 3.84MB)
      flat  flat%   sum%        cum   cum%
  762.95MB 99.32% 99.32%   764.67MB 99.54%  main.main
         0     0% 99.32%   764.67MB 99.54%  runtime.main
(pprof) list main.main
Total: 768.17MB
ROUTINE ======================== main.main in /home/mark/Dev/golang-dev/tutorial.com/41-pprof/prof.go
  762.95MB   764.67MB (flat, cum) 99.54% of Total
         .          .     38:	f, err := os.Create("cpu.prof") // 创建输出文件
         .          .     39:	if err != nil {
         .          .     40:		log.Fatal("could not create CPU profile:", err)
         .          .     41:	}
         .          .     42:
         .     1.72MB     43:	if err := pprof.StartCPUProfile(f); err != nil { // 监控 CPU
         .          .     44:		log.Fatal("could not start CPU profile:", err)
         .          .     45:	}
         .          .     46:
         .          .     47:	defer pprof.StopCPUProfile()
         .          .     48:
         .          .     49:	// 主逻辑
  762.95MB   762.95MB     50:	x := [row][col]int{}
         .          .     51:	fillMatrix(&x)
         .          .     52:	calculate(&x)
         .          .     53:
         .          .     54:	// ------------------- memory -----------------------
         .          .     55:
(pprof)

$ go tool pprof prof goroutine.prof
File: prof
Type: goroutine
Time: Dec 6, 2021 at 1:53am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 2, 100% of 2 total
      flat  flat%   sum%        cum   cum%
         1 50.00% 50.00%          1 50.00%  runtime.gopark
         1 50.00%   100%          1 50.00%  runtime/pprof.runtime_goroutineProfileWithLabels
         0     0%   100%          1 50.00%  main.main
         0     0%   100%          1 50.00%  runtime.main
         0     0%   100%          1 50.00%  runtime/pprof.(*Profile).WriteTo
         0     0%   100%          1 50.00%  runtime/pprof.profileWriter
         0     0%   100%          1 50.00%  runtime/pprof.writeGoroutine
         0     0%   100%          1 50.00%  runtime/pprof.writeRuntimeProfile
         0     0%   100%          1 50.00%  time.Sleep
(pprof) list main.main
Total: 2
ROUTINE ======================== main.main in /home/mark/Dev/golang-dev/tutorial.com/41-pprof/prof.go
         0          1 (flat, cum) 50.00% of Total
         .          .     72:	}
         .          .     73:
         .          .     74:	if gProf := pprof.Lookup("goroutine"); gProf == nil { // 写入内存信息
         .          .     75:		log.Fatal("could not start goroutine profile:", err)
         .          .     76:	} else {
         .          1     77:		gProf.WriteTo(f2, 0)
         .          .     78:	}
         .          .     79:	f2.Close()
         .          .     80:
         .          .     81:}
         .          .     82:
(pprof)
```

## Protobuf

1. 下载 [protoc](https://github.com/protocolbuffers/protobuf) 可执行程序，解压完成之后把二进制文件添加到`~/go/bin`目录下，并把该目录添加到环境变量进行测试：

```bash
$ protoc --version
libprotoc 3.17.1
```

2. 下载 [protoc-gen-go](https://github.com/golang/protobuf)

```bash
$ go get -u -v github.com/golang/protobuf/protoc-gen-go
$ cd ~/go/pkg/mod/github.com/golang/protobuf@v1.5.2/protoc-gen-go
$ go env -w GOBIN=/home/mark/go/bin # 配置生成的二进制文件存放的目录
$ go install # 安装 protoc-gen-go 二进制可执行程序到 GOBIN 目录

# 此时会在 ~/go/bin/ 生成二进制文件
$ ls ~/go/bin/ | grep protoc
protoc
protoc-gen-go
```

3. 编译 Protobuf 文件

```bash
$ protoc --go_out=. *.proto
```