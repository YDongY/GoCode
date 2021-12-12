package conf

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

// Server 服务器配置
var Server struct {
    LogLevel    string
    LogPath     string
    WSAddr      string
    CertFile    string
    KeyFile     string
    TCPAddr     string
    MaxConnNum  int
    ConsolePort int
    ProfilePath string
}

// 获取服务器地址
func init() {
    data, err := ioutil.ReadFile("conf/server.json")
    if err != nil {
        log.Fatalf("%v\n", err)
    }
    err = json.Unmarshal(data, &Server)
    if err != nil {
        log.Fatalf("%v\n", err)
    }
}
