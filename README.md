# 开放式机器学习工程教本

这是一个开放的集合，包含方法、工具和逐步指南，旨在帮助成功训练大型语言模型和多模态模型及其推理。

这是适合LLM/VLM训练工程师和技术员的技术资料。因此，这里的相关内容包含了大量脚本和复制粘贴命令，以便您能够快速解决您的需求。

这个仓库是我对大规模语言模型（LLM）（以及VLMs）培训经验的持续记录；我在2022年培训了开源的[BLOOM-176B](https://huggingface.co/bigscience/bloom) 模型，在2023年培训了多模态[IDEFICS-80B](https://huggingface.co/HuggingFaceM4/idefics-80b-instruct)模型，以及在2024年在[Contextual.AI](https://contextual.ai/)上培训RAG模型的经验总结。

我主要整理这些信息是为了自己方便快速找到过去已经研究并有效的工作解决方案。不过，一如既往地，我很乐意与更广泛的机器学习社区分享这些笔记。

## 目录

**第1部分. 洞察**

1. **[AI战场工程](./insights/ai-battlefield.md)** - 为了成功所需要了解的内容

**第2部分. 硬件**

1. **[计算资源](compute)** - 加速器、CPU和CPU内存。
1. **[存储](storage)** - 本地、分布式及共享文件系统。
1. **[网络](network)** - 内节点和外节点间的通信。

**第3部分. 实验**

1. **[公开可用的训练日志](#publicly-available-training-llmvlm-logbooks)**

## 致谢

这一切都不可能没有我被赋予特定LLM/VLM培训任务的机会。由于租用大型机器学习计算集群的成本高昂，只有少数人能够享受到这一特权。因此，希望其他机器学习社区能从这些笔记中间接学到东西。

特别感谢[Thom Wolf](https://github.com/thomwolf)，他建议我在开始大规模训练之前就开始BLOOM-176B的培训。这项目将我推入了紧张的学习过程之中。当然还有HuggingFace，给我提供了全职参与BLOOM-176B和后来IDEFICS-80B培训的机会。

最近，在[Contextual.AI](https://contextual.ai/)上继续扩展我的知识与经验的过程中，我非常感激Aman和Douwe提供的这一机会。我也要感谢为使这段文字变得精彩且无误而贡献的众多[贡献者](contributors.md)。

## 贡献

如果您发现了一个错误、拼写错误或希望提出改进意见，请不要犹豫，可以打开一个[Issue](https://github.com/stas00/ml-engineering/issues)，也可以提交一个PR。

## 许可证

此站点的内容根据[ Attribution-ShareAlike 4.0 International ](LICENSE-CC-BY-SA) 发行。

## 引用

```bibtex
@misc{bekman2024mlengineering,
  author = {Bekman, Stas},
  title = {Machine Learning Engineering Open Book},
  year = {2023-2024},
  publisher = {Stasosphere Online Inc.},
  journal = {GitHub repository},
  url = {https://github.com/stas00/ml-engineering}
}
```

## 我的仓库地图

✔ **机器学习：**
 - [开放机器学习教本](https://github.com/stas00/ml-engineering) |
 - [ML方法](https://github.com/stas00/ml-ways) |
 - [移植](https://github.com/stas00/porting)

✔ **指南：**
 - [调试的艺术](https://github.com/stas00/the-art-of-debugging)

✔ **应用：**
 - [ipy实验](https://github.com/stas00/ipyexperiments)

✔ **工具和速查表：**
 - [Bash](https://github.com/stas00/bash-tools) |
 - [Conda](https://github.com/stas00/conda-tools) |
 - [Git](https://github.com/stas00/git-tools) |
 - [Jupyter Notebook](https://github.com/stas00/jupyter-notebook-tools) |
 - [Make](https://github.com/stas00/make-tools) |
 - [Python](https://github.com/stas00/python-tools) |
 - [TensorBoard](https://github.com/stas00/tensorboard-tools) |
 - [Unix](https://github.com/stas00/unix-tools)