package main

import (
	"encoding/csv"
	"fmt"
	"github.com/axgle/mahonia"
	"os"
)

func main() {
	//WriteCsv()
	ReadCsv()
}

func WriteCsv() {
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	_, wErr := f.WriteString("\xEF\xBB\xBF")
	if wErr != nil {
		return
	}
	//创建一个新的写入文件流
	w := csv.NewWriter(f)
	data := [][]string{

		{"1", "刘备", "23"},
		{"2", "张飞", "23"},
		{"3", "关羽", "23"},
		{"4", "赵云", "23"},
		{"5", "黄忠", "23"},
		{"6", "马超", "23"},
	}
	fmt.Println("成功写入")
	er := w.WriteAll(data)
	if er != nil {
		return
	}
	w.Flush()
}

func ReadCsv() {
	//filename := "test.csv"
	//fs, err := os.Open(filename)
	//if err != nil {
	//	log.Fatalf("can not open the , err is %+v\n", err)
	//}
	//defer func(fs *os.File) {
	//	errq := fs.Close()
	//	if errq != nil {
	//		fmt.Println(errq)
	//	}
	//}(fs)
	//
	//

	// 打开这个 csv 文件
	file, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 初始化一个 csv reader，并通过这个 reader 从 csv 文件读取数据
	reader := csv.NewReader(file)
	// 设置返回记录中每行数据期望的字段数，-1 表示返回所有字段
	reader.FieldsPerRecord = -1
	// 通过 readAll 方法返回 csv 文件中的所有内容
	record, err := reader.ReadAll()

	if err != nil {
		panic(err)
	}

	decoder := mahonia.NewDecoder("gbk") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	fmt.Println(decoder)
	// 遍历从 csv 文件中读取的所有内容，并将其追加到 tutorials2 切片中
	for _, item := range record {

		fmt.Println(item[0])
		fmt.Println(item[1])

		fmt.Println(item[2])
	}

	// 打印 tutorials2 的第一个元素验证 csv 文件写入/读取是否成功
	//fmt.Println(tutorials2[0].Id)
	//fmt.Println(tutorials2[0].Title)
	//fmt.Println(tutorials2[0].Summary)
	//fmt.Println(tutorials2[0].Author)

	//decoder := mahonia.NewDecoder("gbk") // 把原来ANSI格式的文本文件里的字符，用gbk进行解码。
	//f, err := ioutil.ReadAll(decoder.NewReader(fs))
	//reader := transform.NewReader(bytes.NewReader(f), simplifiedchinese.GBK.NewDecoder())
	//fmt.Println(string(f))

	//for _, row := range f {
	//	fmt.Println(string(row[2]))
	//}
	////fmt.Println(f)

	//r := csv.NewReader(fs)
	//针对大文，一行一行的读取文件
	//i := 0
	//for i < 10 {
	//	row, errw := r.Read()
	//
	//	if errw != nil && errw != io.EOF {
	//
	//		if errw == io.EOF {
	//			break
	//		}
	//	}
	//	i += 1
	//	fmt.Println(row)
	//}
	//fmt.Printf("\n------------------\n")
	//fs1, erro := os.Open(filename)
	//if erro != nil {
	//	log.Fatalf("error %+v", erro)
	//}
	//r2 := csv.NewReader(fs1)
	//content, err1 := r2.ReadAll()
	//
	//if err1 != nil {
	//	log.Fatalf("can not readall, err is %+v", err1)
	//}
	//for _, row := range content {
	//	fmt.Println(row)
	//	//reader := transform.NewReader(bytes.NewReader(), simplifiedchinese.GBK.NewDecoder())
	//	//name, _ := iconv.ConvertString(row[2], "gb2312", "utf-8")
	//	//fmt.Println(name)
	//}
}
