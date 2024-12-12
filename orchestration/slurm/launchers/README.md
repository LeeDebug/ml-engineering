# 使用 SLURM 的单节点和多节点启动脚本

以下是一些完整的 SLURM 脚本，展示了如何将各种启动器与使用 `torch.distributed` 的软件集成（这些脚本也应易于适应其他分布式环境）。

- [torchrun](torchrun-launcher.slurm) - 用于配合 [PyTorch 分布式](https://github.com/pytorch/pytorch)。
- [accelerate](accelerate-launcher.slurm) - 用于配合 [HF Accelerate](https://github.com/huggingface/accelerate)。
- [lightning](lightning-launcher.slurm) - 用于配合 [Lightning](https://lightning.ai/)（“PyTorch Lightning” 和 “Lightning Fabric”）。
- [srun](srun-launcher.slurm) - 用于配合原生的 SLURM 启动器 - 在这里，我们需要手动设置 `torch.distributed` 预期的环境变量。

所有这些脚本都使用了 [torch-distributed-gpu-test.py](../../../debug/torch-distributed-gpu-test.py) 作为示例脚本，你可以通过以下命令将其复制到这里：
```shell
cp ../../../debug/torch-distributed-gpu-test.py .
```
假设你已经克隆了这个仓库。但你可以用你需要的其他任何内容来替换它。