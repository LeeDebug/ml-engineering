# SLURM 性能

这里将讨论影响 SLURM 性能的一些特定设置。

## `srun` 的 `--cpus-per-task` 可能需要明确指定

您需要确保由 `srun` 启动的程序能够获得预期的数量的 CPU 核心。例如，在典型的机器学习训练程序中，每个 GPU 需要至少一个 CPU 核心来驱动其进程，并且还需要几颗额外的核心用于 `DataLoader`。您需要多个核心以使各个任务能够并行执行。假设您有 8 块 GPU，并且每块 GPU 具有两个 `DataLoader` 工作线程，那么每个节点至少需要 `3 * 8 = 24` 颗 CPU 核心。

每项任务的 CPU 核心数由 `--cpus-per-task` 定义，该参数被传递给 `sbatch` 或 `salloc`。根据以往的做法，`srun` 会继承这些设置。然而，最近这种行为有所改变：

来自 `sbatch` 手册页的一段话如下：

> 注意：从 22.05 版本开始，`srun` 将不再继承由 `salloc` 或 `sbatch` 请求的 `--cpus-per-task` 值。若要为任务请求这些设置，需要在调用 `srun` 时再次指定该值或通过 `SRUN_CPUS_PER_TASK` 环境变量进行设置。

这意味着，如果您过去的 SLURM 脚本如下所示：

```shell
#SBATCH --cpus-per-task=48
[...]
srun myprogram
```

那么程序将接收到 48 颗 CPU 核心，因为 `srun` 会继承来自 `sbatch` 或 `salloc` 的 `--cpus-per-task=48` 设置。根据手册页的说明，自 SLURM 22.05 版本起，这种行为不再有效。

注：我在使用 SLURM@22.05.09 测试时发现旧的行为仍然有效，但 23.x 系列确实如此。因此，这一改变可能发生在较晚的 22.05 系列中。

因此，如果您保持现状不变，现在程序将仅接收到一颗 CPU 核心（除非 `srun` 的默认设置已经被修改）。

您可以轻松测试您的 SLURM 配置是否受到影响，使用 `os.sched_getaffinity(0)`，因为它会显示当前进程可以使用的 CPU 核心。因此，通过 `len(os.sched_getaffinity(0))` 计算这些核心的数量应该很容易实现。

以下是测试的方法：
```shell
$ cat test.slurm
#!/bin/bash
#SBATCH --job-name=test-cpu-cores-per-task
#SBATCH --nodes=1
#SBATCH --ntasks-per-node=1
#SBATCH --cpus-per-task=48   # 根据您的环境修改此值，如果您拥有的 CPU 核心少于 48 颗
#SBATCH --time=0:10:00
#SBATCH --partition=x        # 根据您的环境修改到正确的分区名称
#SBATCH --output=%x-%j.out

srun python -c 'import os; print(f"visible cpu cores: {len(os.sched_getaffinity(0))}")'
```

如果输出如下：
```shell
visible cpu cores: 48
```
则您不需要做任何事情；但如果您看到以下结果或其他小于 48 的值：
```shell
visible cpu cores: 1
```
那么您需要采取行动。

要解决这个问题，您可以更改 SLURM 脚本为以下内容之一：

```shell
#SBATCH --cpus-per-task=48
[...]
srun --cpus-per-task=48 myprogram
```

或：
```shell
#SBATCH --cpus-per-task=48
[...]
SRUN_CPUS_PER_TASK=48
srun myprogram
```

或者，您可以使用一次写入即可忽略的自动设置方法：

```shell
#SBATCH --cpus-per-task=48
[...]
SRUN_CPUS_PER_TASK=$SLURM_CPUS_PER_TASK
srun myprogram
```

## 启用超线程还是不启用

如在 [超线程](users.md#hyper-threads) 部分所解释的那样，如果您的 CPU 支持超线程，则可以将可用的 CPU 核心数量加倍。对于某些工作负载来说，这可能会带来整体性能更快的提升。

然而，您需要测试启用或不启用 HT 的情况下的性能，并进行比较以选择表现最佳的设置。

案例研究：在 AWS p4 节点上，我们发现如果启用了超线程，则程序的运行时间会更长。因此，在这种情况下，我们应该禁用超线程。

## 启用还是不启用超线程

在 AWS p4 节点上进行实验时，我们注意到如果启用超线程，则程序的执行时间会增加。因此，在这种情况下，我们应该禁用超线程以获得最佳性能表现。具体来说，您可以使用 `taskset` 工具来限制进程只能使用特定的 CPU 核心集。例如：

```shell
srun --cpus-per-task=4 -c 0-3 bash your_script.sh
```

这样可以确保您的任务仅在指定的核心上运行，并且不会受到影响超线程带来的潜在问题。

## 结论

通过明确指定 `--cpus-per-task` 和适当管理超线程，您可以更好地控制 SLURM 中程序的执行环境，从而获得更好的性能表现。希望这些指南对您有所帮助！如果有任何疑问，请随时提问。祝您使用顺利！