# SLURM 管理

## 在多个节点上运行命令

1. 为了避免每次连接到尚未登录的新节点时被提示：
   ```plaintext
   Are you sure you want to continue connecting (yes/no/[fingerprint])?
   ```
   您可以禁用此检查，方法是将以下内容添加到 `~/.ssh/config` 文件中：
   ```plaintext
   Host *
     StrictHostKeyChecking no
   ```

当然，请确保这足够安全以满足您的需求。假设您已经位于 SLURM 集群，并且不会通过集群之外进行 SSH 连接。您可以选择不设置此操作，然后每次连接到新节点时都需要手动确认。

2. 安装 `pdsh`。

现在可以对多个节点运行所需命令了。

例如，让我们来运行 `date` 命令：

```bash
$ PDSH_RCMD_TYPE=ssh pdsh -w node-[21,23-26] date
node-25: Sat Oct 14 02:10:01 UTC 2023
node-21: Sat Oct 14 02:10:02 UTC 2023
node-23: Sat Oct 14 02:10:02 UTC 2023
node-24: Sat Oct 14 02:10:02 UTC 2023
node-26: Sat Oct 14 02:10:02 UTC 2023
```

让我们来做些更有用且更复杂的操作。假设当 SLURM 作业被取消时，我们需要杀死所有未退出的 GPU 关联进程：

首先，这个命令将提供所有占用 GPU 的进程 ID：

```bash
nvidia-smi --query-compute-apps=pid --format=csv,noheader | sort | uniq
```

因此我们可以一次性杀死这些进程：

```bash
PDSH_RCMD_TYPE=ssh pdsh -w node-[21,23-26] "nvidia-smi --query-compute-apps=pid --format=csv,noheader | sort | uniq | xargs -n1 sudo kill -9"
```

## SLURM 配置

显示 SLURM 的配置设置：

```bash
sudo scontrol show config
```

该配置文件位于控制节点上的 `/etc/slurm/slurm.conf`。

一旦更新了 `slurm.conf` 文件并重新加载配置，请从控制器节点运行以下命令：
```bash
sudo scontrol reconfigure
```

## 自动重启

如果需要安全地重启节点（例如，镜像已更新），则可以调整节点列表并运行：

```bash
scontrol reboot ASAP node-[1-64]
```

对于每个非空闲节点，此命令将在当前作业结束时等待，然后重新启动该节点并将状态恢复为 `idle`。

注意，您需要在控制器节点上的 `/etc/slurm/slurm.conf` 中设置：
```plaintext
RebootProgram = "/sbin/reboot"
```
并如果刚添加了此条目到配置文件，请重新配置 SLURM 服务。

## 更改节点的状态

更改由 `scontrol update` 执行。

示例：

将一个已准备好使用的节点从 `drain` 状态改为 `idle`：
```bash
scontrol update nodename=node-5 state=idle
```

将节点从 SLURM 的资源池中移除：
```bash
scontrol update nodename=node-5 state=drain
```

## 自动取消由于进程退出缓慢导致的节点隔离

有时当作业被取消时，进程会很慢地退出。如果 SLURM 配置为不会无限等待，它将自动对这些节点进行隔离。但是，这些节点没有理由不再可供用户使用。

因此这里是如何自动化此过程的方法：

分析 SLURM 日志文件中的事件日志：
```bash
sudo cat /var/log/slurm/slurmctld.log
```

例如，这可以帮助理解为什么某些节点的作业在规定时间之前被取消或节点完全被移除的原因。


## 修改作业的时间限制

要设置一个新的时间限制（例如2天）：

```bash
scontrol update JobID=$SLURM_JOB_ID TimeLimit=2-00:00:00
```

要在先前设置的基础上增加更多时间（例如3个小时）：

```bash
scontrol update JobID=$SLURM_JOB_ID TimeLimit=+10:00:00
```

## SLURM 出现问题时的处理

分析 SLURM 日志文件中的事件日志：
```bash
sudo cat /var/log/slurm/slurmctld.log
```

这可以例如帮助理解为什么某个节点的作业在规定时间之前被取消或节点完全被移除的原因。