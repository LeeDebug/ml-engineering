## 重现基于随机性的软件

在调试时，始终为所有使用的随机数生成器（RNG）设置固定种子，以便每次重新运行时获取相同的数据/代码路径。

虽然要涵盖如此多的不同系统可能会有些困难。这里尝试覆盖几个：

```python
import random, torch, numpy as np
def enforce_reproducibility(use_seed=None):
    seed = use_seed if use_seed is not None else random.randint(1, 1000000)
    print(f"使用种子: {seed}")

    random.seed(seed)    # python RNG
    np.random.seed(seed) # numpy RNG

    # pytorch RNGs
    torch.manual_seed(seed)          # cpu + cuda
    torch.cuda.manual_seed_all(seed) # 多 GPU - 可以在没有 GPU 的情况下调用
    if use_seed:  # 较慢的速度！https://pytorch.org/docs/stable/notes/randomness.html#cuda-convolution-benchmarking
        torch.backends.cudnn.deterministic = True
        torch.backends.cudnn.benchmark     = False

    return seed
```

如果使用其他子系统/框架，这里有一些其他的可能：

```python
    torch.npu.manual_seed_all(seed)
    torch.xpu.manual_seed_all(seed)
    tf.random.set_seed(seed)
```

每次再次运行相同的代码以解决某个问题时，请在代码开头设置特定的种子：
```python
enforce_reproducibility(42)
```
但请注意，这仅用于调试，因为它会激活各种有助于确定性的 torch 标志，但在生产环境中可能会降低性能。

然而，您可以使用以下方式来生产中使用该函数：
```python
enforce_reproducibility()
```
即没有显式的种子。然后它会选择一个随机的种子并记录下来！这样如果在生产中出现问题，您就可以重现相同的 RNG 设置，并且这一次不会有任何性能损失，因为仅在提供种子时才会设置 `torch.backends.cudnn` 标志。假设日志如下：
```python
使用种子: 1234
```
您只需将代码更改为：
```python
enforce_reproducibility(1234)
```
即可获得相同的 RNG 设置。

如上段所述，在系统中可能会涉及到其他 RNG，例如，如果您希望数据以相同顺序馈送到 `DataLoader` 中，请务必 [设置其种子](https://pytorch.org/docs/stable/notes/randomness.html#dataloader)。

### 重现软件和系统环境

当发现某些结果差异时（如质量或吞吐量），该方法很有用。

想法是记录启动训练（或推断）所使用的环境的关键组件，以便如果在稍后阶段需要精确地重现这些情况，则可以做到这一点。

由于使用了大量不同系统和组件，无法给出一种始终有效的方法。因此我们讨论一个可能的配方，您可以根据自己的特定环境进行调整。

将此添加到您的 slurm 启动脚本中：
```python
def enforce_reproducibility(use_seed=None):
    # ...
```

1. 如果您使用 `modules` 来加载如 cuDNN 和 CUDA 库，则在不使用 `modules` 的情况下，请删除相应的条目。

2. 我们记录环境变量。这可能非常关键，因为单个环境变量（例如 `LD_PRELOAD` 或 `LD_LIBRARY_PATH`）可能会对某些环境中的性能产生巨大影响。

3. 然后我们记录 conda 环境的包及其版本 - 这应该适用于任何虚拟 Python 环境。

4. 如果您使用 `pip install -e .` 安装了开发安装，它只知道 git 仓库的 SHA 而不了解本地修改。因此此部分会遍历未安装到 conda 环境中的所有包（我们通过查看 `site-packages/*.dist-info/direct_url.json` 找到这些包）。

一个额外有用的工具是 [conda-env-compare.pl](https://github.com/stas00/conda-tools/blob/master/conda-env-compare.md) ，它可以帮助您找出两个 conda 环境之间的确切差异。

我与我的同事在同一个云集群上以相同的代码运行训练时得到了非常不同的 TFLOPs - 实际上是从相同共享目录启动了相同的 slurm 脚本。我们首先使用 [conda-env-compare.pl](https://github.com/stas00/conda-tools/blob/master/conda-env-compare.md) 比较我们的 conda 环境，并发现了一些差异：我安装了她所有的包以匹配她的环境，但仍然表现出巨大的性能差距。然后我们比较了 `printenv` 的输出并发现我的 `LD_PRELOAD` 设置而她没有 - 这个特定的云提供商要求设置多个环境变量到自定义路径才能充分利用硬件。

这种差异对我们的结果产生了重大影响。通过这种方法，我们可以确保在生产环境中正确地再现和优化性能。