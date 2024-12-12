# 在SLURM环境中的工作

除非你很幸运，拥有一个完全由自己控制的专用集群，否则很有可能你需要使用SLURM与其他用户共享GPU资源。即使你在高性能计算（HPC）环境中进行训练，并被分配了一个专用分区，你也仍然需要使用SLURM。

SLURM的缩写代表：**Simple Linux Utility for Resource Management**——尽管现在它被称为 **Slurm Workload Manager**。它是一个用于Linux和类Unix内核的操作系统的工作负载管理调度器，在世界上许多超级计算机和计算集群中被广泛使用。

这些章节不会详尽地教你如何使用SLURM，因为有很多相关的手册可供参考，但会涵盖一些有助于训练过程的具体细节。

- [SLURM用户指南](./users.md) - 在SLURM环境中进行培训所需了解的所有内容。
- [SLURM管理指南](./admin.md) - 如果不幸需要同时管理和使用SLURM集群，则此文档中列出了许多逐步指导，可以帮助你更快地完成任务。
- [性能优化](./performance.md) - SLURM的性能细节。
- [启动脚本](./launchers) - 如何在SLURM环境中使用`torchrun`、`accelerate`、PyTorch Lightning等工具进行训练。