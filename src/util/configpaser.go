package util

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	config1 = "/etc/private_website/config/"
	config2 = "../config/"
)

var (
	Mysql  map[string]string
	Redis  map[string]string
	Net    map[string]string
	LogDir string
)

func init() {
	if _, err := ioutil.ReadDir(config1); err != nil {
		Log(WARN, "config1 文件夹不存在")
	} else {
		Log(DEBUG, "Use config1 to read config!")
		read_config(config1)
		return
	}
	if _, err := ioutil.ReadDir(config2); err != nil {
		Log(WARN, "config2 文件夹不存在")
	} else {
		Log(DEBUG, "Use config2 to read config!")
		read_config(config1)
		return
	}
}

func read_config(dir string) {
	fileInfo, _ := ioutil.ReadDir(dir)
	bytes := make([]byte, 10)
	for _, v := range fileInfo {
		f, err := os.Open(dir + v.Name())
		defer f.Close()
		if err != nil {
			Log(ERROR, "config file open err:", err)
			panic(err)
		}
		for {
			byt := make([]byte, 10)
			f.Read(byt)
			for _, i := range byt {
				bytes = append(bytes, i)
			}
		}
	}
	str := string(bytes)
	splitStr := strings.Split(str, "&")
	for _, i := range splitStr {
		switch {
		case strings.Contains(strings.Split(i, "\r\n")[0], "mysql"):
			mysql_init(strings.Split(i, "\r\n"))
			break
		case strings.Contains(strings.Split(i, "\r\n")[0], "#"):
			break
		case strings.Contains(strings.Split(i, "\r\n")[0], "redis"):
			redis_init(strings.Split(i, "\r\n"))
			break
		case strings.Contains(strings.Split(i, "\r\n")[0], "net"):
			net_init(strings.Split(i, "\r\n"))
			break
		default:
			break
		}

	}

}

func mysql_init(str []string) {

}

func redis_init(str []string) {

}

func net_init(str []string) {

}
