# 调试与故障排除

## 指南

- [调试 PyTorch 程序](./pytorch.md)

- [诊断多节点多GPU Python 程序中的挂起和死锁问题](./torch-distributed-hanging-solutions.md)

- [网络调试](../network/debug/)

- [排查 NVIDIA GPU 问题](../compute/accelerator/nvidia/debug.md)

- [检测下溢与上溢](./underflow_overflow.md)


## 工具

- [调试工具](./tools.md)

- [torch-distributed-gpu-test.py](./torch-distributed-gpu-test.py) - 这是一个用于诊断的 `torch.distributed` 脚本，可以检查集群中的所有 GPU（一个或多个节点）是否能够互相通信并分配 GPU 内存。

- [NicerTrace](./NicerTrace.py) - 这是一个改进过的 `trace` Python 模块，构造函数中添加了多个额外标志，并且输出更加有用。