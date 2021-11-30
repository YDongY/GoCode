package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func OpenFile() {
	/** flag 选项
	const (
		// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
		O_RDONLY int = syscall.O_RDONLY // open the file read-only.
		O_WRONLY int = syscall.O_WRONLY // open the file write-only.
		O_RDWR   int = syscall.O_RDWR   // open the file read-write.
		// The remaining values may be or'ed in to control behavior.
		O_APPEND int = syscall.O_APPEND // append data to the file when writing.
		O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
		O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
		O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
		O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	)
	*/

	file, err := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func Open() {
	file, err := os.Open("file.txt") // 只读方式打开文件
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func Create() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func Read() {
	file, err := os.Open("file.txt") // 只读方式打开文件
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var buf [128]byte
	var content []byte

	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			// 读取结束
			break
		}
		if err != nil {
			fmt.Println("read file err ", err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

func ReadAt() {
	file, err := os.Open("file.txt") // 只读方式打开文件
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var buf [128]byte
	// 偏移 3 字节，从第 4 个字节开始读取数据
	if n, err := file.ReadAt(buf[:], 3); err != io.EOF {
		log.Fatal(err)
	} else {
		fmt.Println(string(buf[:n]))
	}
}

func Write() {
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := file.Write([]byte("hello world")); err != nil {
		log.Fatal(err)
	}
}

func WriteString() {
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hello Golang")
}

func WriteAt() {
	file, err := os.OpenFile("file.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := file.WriteAt([]byte("Python"), 6); err != nil {
		log.Fatal(err)
	}
}

func Remove() {
	if err := os.Remove("file.txt"); err != nil {
		log.Fatal(err)
	}
}

func Stat() {
	if fileInfo, err := os.Stat("file.txt"); err != nil {
		log.Fatal(err)
	} else {
		/**
		type FileInfo interface {
			Name() string       // base name of the file
			Size() int64        // length in bytes for regular files; system-dependent for others
			Mode() FileMode     // file mode bits
			ModTime() time.Time // modification time
			IsDir() bool        // abbreviation for Mode().IsDir()
			Sys() interface{}   // underlying data source (can return nil)
		}
		*/
		fmt.Printf("name:%s\nsize:%d\nmode:%v\nisdir:%t", fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), fileInfo.IsDir())
	}
}

func Rename() {
	if err := os.Rename("file.txt", "new_file.txt"); err != nil {
		log.Fatal(err)
	}
}
func main() {
	Rename()
}
