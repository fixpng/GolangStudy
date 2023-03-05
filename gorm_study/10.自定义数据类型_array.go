package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type Array []string

// Scan 从数据库中读取出来
func (arr *Array) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("解析失败: %v, %T\n", value, value))
	}
	*arr = strings.Split(string(data), "|")
	return nil
}

// Value 存入数据库
func (arr Array) Value() (driver.Value, error) {
	fmt.Printf("入库前,%#v,%T\n", arr, arr)
	return strings.Join(arr, "|"), nil
}

type HostModel struct {
	ID    uint
	IP    string
	Ports Array `gorm:"type:string"`
}

func main() {
	//DB.AutoMigrate(&HostModel{})

	//DB.Create(&HostModel{
	//	IP:    "192.168.100.15",
	//	Ports: []string{"8080", "22"},
	//})

	var host []HostModel
	DB.Find(&host)
	fmt.Println(host)

}
