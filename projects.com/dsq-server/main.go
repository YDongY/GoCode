package main

import (
    "dsq-server/conf"
    "dsq-server/modules/game"
    "dsq-server/modules/gate"
    "dsq-server/modules/login"
    "github.com/name5566/leaf"
    lconf "github.com/name5566/leaf/conf"
)

func main() {

    lconf.LogLevel = conf.Server.LogLevel
    lconf.LogPath = conf.Server.LogPath
    lconf.LogFlag = conf.LogFlag
    lconf.ConsolePort = conf.Server.ConsolePort
    lconf.ProfilePath = conf.Server.ProfilePath

    leaf.Run(
        game.Module,
        gate.Module,
        login.Module,
    )
}
