# 张量精度 / 数据类型

以下是截至当前在机器学习（ML）中常用的常见数据类型，通常称为 `dtype`：

浮点格式：
- fp32 - 32位
- tf32 - 19位（NVIDIA Ampere+）
- fp16 - 16位
- bf16 - 16位
- fp8 - 8位（E4M3 和 E5M2 格式）

为了进行视觉对比，请参阅以下表示：

![fp32-tf32-fp16-bf16](images/fp32-tf32-fp16-bf16.png)

([来源](https://developer.nvidia.com/blog/accelerating-ai-training-with-tf32-tensor-cores/))

![fp16-bf16-fp8](images/fp16-bf16-fp8.png)

([来源](https://docs.nvidia.com/deeplearning/transformer-engine/user-guide/examples/fp8_primer.html))

用于量化时使用的整数格式：

- int8 - 8位
- int4 - 4位
- int1 - 1位

## ML 数据类型演进

最初，ML 使用 fp32，但这非常缓慢。

接下来发明了混合精度技术，即使用 fp16 和 fp32 的组合，大大加快了训练速度。

![fp32/fp16 混合精度](images/mixed-precision-fp16.png)

([来源](https://developer.nvidia.com/blog/video-mixed-precision-techniques-tensor-cores-deep-learning/))

但 fp16 证明不够稳定，训练大规模语言模型（LLM）极其困难。

幸运的是 bf16 出现并使用相同的混合精度协议取代了 fp16。这使得 LLM 训练更加稳定。

然后出现了 fp8，并且混合精度切换到了 [该方案](https://docs.nvidia.com/deeplearning/transformer-engine/user-guide/examples/fp8_primer.html)，使其训练速度更快。请参阅论文：[FP8 格式在深度学习中的应用](https://arxiv.org/abs/2305.16479)。

## 训练后改变精度

有时可以在模型训练完成后更改精度。

- 使用预训练的 bf16 模型在 fp16 环境中通常会失败 - 因为 fp16 能表示的最大数值是 64k，可能会导致溢出。有关详细讨论和可能解决方案，请参阅此 [PR](https://github.com/huggingface/transformers/pull/10956)。

- 使用预训练的 fp16 模型在 bf16 环境中通常会工作 - 转换时可能会损失一些性能，但应该可以正常工作 - 最好在其上进行微调再使用。

## 低精度数据类型后的精度变化

有时可以在模型训练完成后改变精度。

- 使用预训练的 bf16 模型在 fp16 环境中通常会失败 - 因为 fp16 能表示的最大数值是 64k，可能会导致溢出。有关详细讨论和可能解决方案，请参阅此 [PR](https://github.com/huggingface/transformers/pull/10956)。

- 使用预训练的 fp16 模型在 bf16 环境中通常会工作 - 转换时可能会损失一些性能，但应该可以正常工作 - 最好在其上进行微调再使用。


### 低精度下的注意事项

每当使用低精度数据类型时，必须小心不要将中间结果保留在该数据类型中。

- 层归一化（LayerNorm）等操作不应在半精度下执行，否则可能会丢失大量数据。因此当这些操作实现正确时，它们会在输入数据类型的 dtype 中高效地进行内部工作，但使用 fp32 积累寄存器，然后将输出降级为输入的精度。

一般而言，只是积聚运算在 fp32 中完成，因为否则将许多低精度数字相加会非常损失精度。

以下是一些示例：

1. 归约集体操作

   - fp16: 如果有缩放因子可以接受在 fp16 下进行
   - bf16: 只能在 fp32 下进行

2. 梯度累积

   - 对于 fp16 和 bf16，最好在 fp32 下进行，但对于 bf16 必须如此。

3. 优化器步骤 / 消失的梯度问题

   - 当向大数值添加一个很小的梯度时，这个加法通常会被抵消。因此通常使用 fp32 主权重和 fp32 优化状态。
   
   - 可以在使用 [卡曼求和](https://en.wikipedia.org/wiki/Kahan_summation_algorithm) 或 [随机舍入](https://en.wikipedia.org/wiki/Rounding)（参见论文 [重访 BFloat16 训练](https://arxiv.org/abs/2010.06192)）的情况下使用 fp16 主权重和优化状态。

对于后者的一个示例，请参阅：[AnyPrecision 优化器](https://github.com/pytorch/torchdistx/pull/52)，最新版本可在此找到 [这里](https://github.com/facebookresearch/multimodal/blob/6bf3779a064dc72cde48793521a5be151695fc62/torchmultimodal/modules/optimizers/anyprecision.py#L17)。

## 训练后更改精度

有时可以在训练完成后更改精度：

- 使用预训练的 bf16 模型在 fp16 环境中通常会失败 - 因为 fp16 能表示的最大数值是 64k，可能会导致溢出。有关详细讨论和可能解决方案，请参阅此 [PR](https://github.com/huggingface/transformers/pull/10956)。

- 使用预训练的 fp16 模型在 bf16 环境中通常会工作 - 转换时可能会损失一些性能，但应该可以正常工作 - 最好在其上进行微调再使用。