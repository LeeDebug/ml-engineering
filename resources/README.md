# 资源

## 有用的资料汇总

- [@StellaAthena](https://github.com/StellaAthena) 创建了[Common LLM Settings 表格](https://docs.google.com/spreadsheets/d/14vbBbuRMEHoqeuMHkTfw3uiZVmyXNuoSp8s-aHvfvZk/edit#gid=0)，在你即将开始新的LLM训练时，这可以是一个超级有用的资源——因为它告诉你已经创建了多少已知的LLM训练。

- 几年前我开始整理有关模型训练数据类型的信息[（链接）](https://discuss.huggingface.co/t/model-pre-training-precision-database-fp16-fp32-bf16/5671) - 它仅包含少量模型，但如果你在进行关于数据类型的科研工作，它仍然可能很有用。我使用这些信息来尝试编写[一个模型预训练数据类型自动检测工具](https://github.com/stas00/ml-ways/blob/master/numbers/detect-model-pretrained-in-bf16-fp16-fp32.ipynb)，这里还有一个有关 [float16 vs bfloat16 数值属性比较的研究](https://github.com/stas00/ml-ways/blob/master/numbers/bfloat16-vs-float16-study.ipynb)。

## 公开可用的训练LLM/VLM日志

训练LLM/VLM的日志和纪要是了解处理训练不稳定性及选择良好超参数的最佳来源之一。

如果你知道有不在本列表中的公开LLM/VLM训练日志，请告知我或通过PR添加。感谢！

按年份分组，但不分特定顺序：

### 2021

- BigScience预BLOOM 108B训练实验（2021）：
[纪事](https://github.com/bigscience-workshop/bigscience/blob/master/train/tr8-104B-wide/chronicles.md) |
[完整规格和讨论](https://github.com/bigscience-workshop/bigscience/blob/master/train/tr8-104B-wide)
（备份：
[1](https://github.com/stas00/bigscience-backup/blob/master/train/tr8-104B-wide/chronicles.md) |
[2](https://github.com/stas00/bigscience-backup/blob/master/train/tr8-104B-wide)）

### 2022

- BigScience BLOOM-176B（2022）：
[纪事前传](https://github.com/bigscience-workshop/bigscience/blob/master/train/tr11-176B-ml/chronicles-prequel.md) |
[纪事](https://github.com/bigscience-workshop/bigscience/blob/master/train/tr11-176B-ml/chronicles.md) （备份：
[1](https://github.com/stas00/bigscience-backup/blob/master/train/tr11-176B-ml/chronicles-prequel.md) |
[2](https://github.com/stas00/bigscience-backup/blob/master/train/tr11-176B-ml/chronicles.md))

- BloombergGPT 50B LLM - 在《[BloombergGPT：一个用于金融的大型语言模型](https://arxiv.org/abs/2303.17564)》一文中的部分C

### 2023

- HuggingFace IDEFICS-80B多模态（Flamingo复现）（2023）：
[学习日志](https://github.com/huggingface/m4-logs/blob/master/memos/README.md) |
[训练纪事](https://github.com/huggingface/m4-logs/blob/master/tr-190-80b/chronicles.md) （备份：
[1](https://github.com/stas00/m4-logs-backup/blob/master/memos/README.md) |
[2](https://github.com/stas00/m4-logs-backup/blob/master/tr-190-80b/chronicles.md))

- BloombergGPT 50B LLM - 在《[BloombergGPT：一个用于金融的大型语言模型](https://arxiv.org/abs/2303.17564)》一文中的部分C

### 2024

- [MegaScale: 将大规模语言模型训练扩展到超过10,000个GPU](https://arxiv.org/abs/2402.15627) - 这篇文章涵盖了各种训练问题及其解决方案——尽管是在专有的模型上进行的，但同样具有教育和实用价值。

- Imbue 的[从裸金属到70B模型：基础设施配置与脚本](https://imbue.com/research/70b-infrastructure/) 非常详细的技术文章涵盖了他们在训练一个专用的70亿参数模型时必须克服的许多训练相关问题。

## 硬件设置日志

- Imbue 公布了他们如何搭建一个包含512节点IB胖树集群并使其正常运行的详细记录：[从裸金属到70B模型：基础设施配置与脚本](https://imbue.com/research/70b-infrastructure/)，他们在过程中还开源了自己创建的[集群工具](https://github.com/imbue-ai/cluster-health)。