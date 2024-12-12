# 检查点

- [torch-checkpoint-convert-to-bf16](./torch-checkpoint-convert-to-bf16) - 将现有的 fp32 PyTorch 检查点转换为 bf16。如果找到 safetensors（[链接](https://github.com/huggingface/safetensors/)），也会将其进行转换。可以轻松适应其他类似用例。

- [torch-checkpoint-shrink.py](./torch-checkpoint-shrink.py) - 此脚本修复了一些原因导致存储在检查点中的张量大小超过其保存时视图大小的情况。它会克隆当前视图并仅使用当前视图的存储重新保存这些张量。