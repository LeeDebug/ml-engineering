# 出版建设

重要提示：这仍处于开发中——大部分功能已经实现，但样式表需要改进以使PDF更加美观。预计几周内完成。

此文档假设你从仓库的根目录进行操作。

## 安装要求

1. 安装构建书籍时使用的Python包
   ```sh
   pip install -r build/requirements.txt
   ```

2. 下载免费版本的[Prince XML](https://www.princexml.com/download/)。它用于生成这本书的PDF版本。

## 构建HTML

```sh
make html
```

## 构建PDF

```sh
make pdf
```

它会首先构建HTML目标，然后使用该目标来构建PDF版本。

## 检查链接和锚点

要验证所有本地链接和锚定链接的有效性，请运行：
```sh
make check-links-local
```

 若要同时检查外部链接（请谨慎使用以避免服务器被频繁请求而封禁）：
```sh
make check-links-all
```

## 移动md文件/目录并调整相对链接

例如，将 `slurm` 移到 `orchestration/slurm`
```sh
src=slurm
dst=orchestration/slurm

mkdir -p orchestration
git mv $src $dst
perl -pi -e "s|$src|$dst|" chapters-md.txt
python build/mdbook/mv-links.py $src $dst
git checkout $dst
make check-links-local
```

## 调整图片大小

当包含的图片过大时，可以稍微调整它们的大小：
```sh
mogrify -format png -resize 1024x1024\> *png
```