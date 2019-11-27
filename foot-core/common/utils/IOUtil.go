package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// 一次性读取
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	/*
	   ReadAll从r读取数据直到EOF或遇到error，返回读取的数据和遇到的错误。
	   成功的调用返回的err为nil而非EOF。
	   因为本函数定义为读取r直到EOF，它不会将读取返回的EOF视为应报告的错误。
	*/
	return ioutil.ReadAll(f)
}

// 分块读取 可在速度和内存占用之间取得很好的平衡。
func ReadBlock(filePth string, bufSize int, hookfn func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := make([]byte, bufSize) //一次读取多少个字节
	/*
	   NewReader创建一个具有默认大小缓冲、从r读取的*Reader。
	*/
	bfRd := bufio.NewReader(f)
	for {
		n, err := bfRd.Read(buf)
		hookfn(buf[:n]) // n 是成功读取字节数

		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
			if err == io.EOF {
				return nil
			}
			return err
		}
	}

	return nil
}

// 输出到控制台
func processTask(line []byte) {
	os.Stdout.Write(line)
	//fmt.Println(string(line))
}

// 逐行读取
func ReadLine(filePth string, hookfn func([]byte)) error {
	f, err := os.Open(filePth)
	if err != nil {
		return err
	}
	defer f.Close()

	bfRd := bufio.NewReader(f)
	for {
		line, err := bfRd.ReadBytes('\n')
		hookfn(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

// 文件监控
func FileMonitoring(filePth string, hookfn func([]byte)) {
	f, err := os.Open(filePth)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	f.Seek(0, 2)
	for {
		line, err := rd.ReadBytes('\n')
		// 如果是文件末尾不返回
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue;
		} else if err != nil {
			log.Fatalln(err)
		}
		go hookfn(line)
	}

}
//func main() {

	/*
	   一次性读取
	*/

	// 直接读取文件,无需打开句柄
	/*
	   ret,err:=ioutil.ReadFile("/usr/local/nginx/logs/access.log")
	   if err != nil {
	       fmt.Println(err)
	   }
	   fmt.Println(string(ret))
	*/

	// 一次性读取
	/*
	   ret, err := ReadAll("/usr/local/nginx/logs/access.log")
	   if err != nil {
	       fmt.Println(err)
	   }
	   fmt.Println(string(ret))
	*/

	// 分块读取
	/*
	   ReadBlock("/usr/local/nginx/logs/access.log", 10000, processTask)
	*/

	// 逐行读取
	/*
	   ReadLine("/usr/local/nginx/logs/access.log", processTask)
	*/

	// 示例 日志实时监控
	//FileMonitoring("/usr/local/nginx/logs/access.log", processTask)

//}
