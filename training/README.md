# 训练

**子部分：**

- [模型并行](model-parallelism)

- [性能](performance)

- [容错性](fault-tolerance)

- [可重复性](reproducibility)

- [不稳定性](instabilities)

- [检查点](checkpoints)

- [训练超参数和模型初始化](hparams.md)

- [张量精度/数据类型](dtype.md)

- [仅使用单个节点模拟多节点设置](emulate-multi-node.md) - 如何使用单个节点模拟多节点设置的说明，我们在这里使用 `deepspeed` 启动器。

- [从 Hugging Face 仓库重新训练模型（使用微调示例）](re-train-hub-models.md)

**工具：**

- [printflock.py](tools/printflock.py) - 一个小型库，使您的 `print` 调用在多 GPU 环境中不会交错。

- [multi-gpu-non-interleaved-print.py](tools/multi-gpu-non-interleaved-print.py) - 基于 `flock` 的 `print` 包装器，防止当多个进程同时打印时消息被交错——这在使用多 GPU 与 `torch.distributed` 结合时是这种情况。