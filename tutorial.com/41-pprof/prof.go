package main

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

const (
	row = 10000
	col = 10000
)

func fillMatrix(m *[row][col]int) {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			m[i][j] = s.Intn(10000)
		}
	}
}

func calculate(m *[row][col]int) {
	for i := 0; i < row; i++ {
		tmp := 0
		for j := 0; j < col; j++ {
			tmp += m[i][j]
		}
	}
}

func main() {
	// ------------------- CPU -----------------------

	f, err := os.Create("cpu.prof") // 创建输出文件
	if err != nil {
		log.Fatal("could not create CPU profile:", err)
	}

	if err := pprof.StartCPUProfile(f); err != nil { // 监控 CPU
		log.Fatal("could not start CPU profile:", err)
	}

	defer pprof.StopCPUProfile()

	// 主逻辑
	x := [row][col]int{}
	fillMatrix(&x)
	calculate(&x)

	// ------------------- memory -----------------------

	f1, err := os.Create("mem.prof") // 创建输出文件
	if err != nil {
		log.Fatal("could not create memory profile:", err)
	}

	// runtime.GC()
	if err := pprof.WriteHeapProfile(f1); err != nil { // 写入内存信息
		log.Fatal("could not start memory profile:", err)
	}
	f1.Close()

	// ------------------- goroutine -----------------------

	f2, err := os.Create("goroutine.prof") // 创建输出文件
	if err != nil {
		log.Fatal("could not create goroutine profile:", err)
	}

	if gProf := pprof.Lookup("goroutine"); gProf == nil { // 写入内存信息
		log.Fatal("could not start goroutine profile:", err)
	} else {
		gProf.WriteTo(f2, 0)
	}
	f2.Close()

}

/**
```bash
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
*/
