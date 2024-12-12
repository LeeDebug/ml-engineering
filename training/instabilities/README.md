# 避免、恢复和理解不稳定性

子部分：

* [理解训练损失模式](training-loss-patterns.md) - 包括尖峰、发散、悟性时刻、恢复等类型。

## 从训练日志中学习

最好的学习是阅读[公开的LLM/VLM训练日志本](../../resources#publicly-available-training-llmvlm-logbooks)，因为在其中可以看到具体发生了什么以及是如何克服问题的。

## STD 初始化

正确初始化张量的初始分布对训练的稳定性有着巨大的影响。`std` 值不是固定的，它取决于隐藏维度大小。

这在我们预BLOOM 104B实验中证明是一个非常关键的设置：直到我们发现默认的 `--init-method-std` 在Megatron-LM中的值0.02对于我们的模型来说太大了，我们无法突破最初的几千次迭代。

我们参考了以下这两篇资料：

1. "Transformers without Tears" 研究论文 https://arxiv.org/abs/1910.05895 推荐：`sqrt(2/(NHIDDEN*5))`

2. 530B模型的训练论文 https://arxiv.org/abs/2201.11990，他们使用了更小的初始化公式: `sqrt(1/(NHIDDEN*3))`

我们选择了后者，因为它的初始值更小。

为了让这两个公式的比较更为直观，可以将它们重写为：
1. `sqrt(0.4000/NHIDDEN)`
2. `sqrt(0.3333/NHIDDEN)`

因此对于 `NHIDDEN=14336` 的情况，计算结果为 `sqrt(1/(14336*3)) = 0.00482`。这就是我们使用的值。这绝对不是BLOOM-176B训练过程中没有出现稳定性问题的唯一原因，但我认为这是非常关键的原因之一。

## 数值不稳定性

某些数学运算在处理低精度数字时可能会变得不稳定。

例如，请参阅这个非常有趣的 [PyTorch数值稳定性指南](https://pytorch.org/docs/stable/notes/numerical_accuracy.html)。

现在我们来看一个具体的例子，展示这一概念的实际应用。

在使用 fp16 混合精度训练 104B 模型的过程中，[Corby Rosset](https://github.com/corbyrosset) 提出了一项改进措施，旨在使 [自注意力机制] 更加稳定。具体内容如下：

具体来看，代码中的这一行展示了 `norm_factor` 可能在查询 * 关键矩阵乘法之后被相乘的情况。如果 Q 和 K 的维度非常大，输出可能会变得过大，从而使 `norm_factor` 无法挽救。

提议：将 `norm_factor` 内移，即在矩阵相乘前对 Q 和 K 进行缩放：
```python
        matmul_result = torch.baddbmm(
            matmul_result,
            1.0/math.sqrt(self.norm_factor) * query_layer.transpose(0, 1),   # [b * np, sq, hn]
            1.0/math.sqrt(self.norm_factor) * key_layer.transpose(0, 1).transpose(1, 2),  # [b * np, hn, sk]
            beta=0.0 if alibi is None else 1.0, C=matmul_result
        )

```

这样做的目的是在矩阵相乘之前应用 `norm_factor`。

## 数据批次和模型参数状态的“不良”组合

PaLM 团队观察到，在训练更大规模模型时，损失会在“高度不规则的时间间隔”出现尖峰。虽然他们未能追踪到根本原因，但他们通过从较早的检查点重新开始并跳过可能存在问题的数据批次来缓解了这个问题。[第5.1节 训练不稳定性](https://arxiv.org/pdf/2204.02311.pdf)

## Adam 中的时间域相关性发散

一篇名为 [关于大规模机器学习中 Adam 不稳定性的理论](https://arxiv.org/abs/2304.09871) 的研究对在高达546B参数下训练LLMs时的发散尖峰进行了严格的研究，并建议时间域相关性会导致Adam发散。这由epsilon值不够小触发，梯度估计分量变得类似于epsilon。

在第7.1节中，他们提出了实用建议，其中最有趣的一条是将epsilon设置为0，并可能处理除零条件。