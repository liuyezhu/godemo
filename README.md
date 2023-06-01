项目配置开启go mod
go 1.11和1.12版本

将下面两个设置添加到系统的环境变量中
GO111MODULE=on
GOPROXY=https://goproxy.io

1.13版本之后
需要注意的是这种方式并不会覆盖之前的配置，有点坑，你需要先把系统的环境变量里面的给删掉再设置
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

热启动
https://github.com/cosmtrek/air/releases