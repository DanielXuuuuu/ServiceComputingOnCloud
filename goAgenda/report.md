# goAgenda

#### 安装cobra

使用以下命令安装

```/
go get -v github.com/spf13/cobra/cobra
```

会出现以下错误信息：

> Fetching https://golang.org/x/sys/unix?go-get=1
>
> https fetch failed: Get https://golang.org/x/sys/unix?go-get=1: dial tcp 216.239.37.1:443: i/o timeout

解决办法是：进入`$GOPATH/src/golang.org/x`目录下，执行：

```
git clone https://github.com/golang/sys.git
git clone https://github.com/golang/text.git
```

然后执行`go install github.com/spf13/cobra/cobra`完成安装。可以看到`$GOBIN`下出现了cobra的可执行程序。

#### 项目创建

使用命令`cobra init goAgenda --pkg-name goAgenda`创建项目。

#### 项目设计



#### 持久化

