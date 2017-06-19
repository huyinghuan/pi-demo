# Raspberry PI  demo for  GO 

Go 树莓派的控制硬件demo， 附电路图

## start

```
git clone https://github.com/huyinghuan/pi-demo.git
cd pi-demo
go get
```

## compile 编译

gox or go build

gox https://github.com/mitchellh/gox

gox build:
```
 gox -osarch="linux/arm" -output="bin/xxx"
```

go build：

```
 export GOOS=linux
 export GOARCH=arm
 go build ...
```

##  Directory 

### work

demo code
示例代码

### circuite design

circuite design  电路示意图。

xx.fzz 为设计图原件 请用[fritzing] 包含电路设计
xx.png 为面包图
xx_design.png 为该面包图

## Tools 电路图工具： 

Fritzing http://fritzing.org/home/

## Book

[Raspberry Pi Sensors] author: Rushi Gajjar