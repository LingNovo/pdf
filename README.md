# go-pdf

环境:
ubuntu 16.04

## 依赖
-- 安装mupdf

参照 https://www.mupdf.com/docs/building.html,本示例所用mupdf 本版为1.12.0
安装命令:
```
git clone --recursive git://git.ghostscript.com/mupdf.git
git submodule update --init
make prefix=/usr/local install
```

-- 安装c-for-go

```
go get github.com/xlab/c-for-go
```

## 其它说明
利用c-for-go 生成代码以后，有些函数报错，直接注释了。
ubuntu 环境下，有些静态库文件提示找不到，使用find 命令找到对应文件或目录，根据错误提示复制到相应目录即可。
c99 编译器： https://github.com/cznic/cc

