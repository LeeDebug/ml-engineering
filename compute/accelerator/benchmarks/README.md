# 加速器基准测试

## 可达最大矩阵乘法 FLOPS 寻找器

可达最大矩阵乘法（MAMF）基准：[mamf-finder.py](./mamf-finder.py)

有关详细讨论和各种加速器的数字，请参见 [可达最大浮点运算次数](../#maximum-achievable-flops)。

虽然一些加速器制造商发布了理论 TFLOPS，但这些数值通常都无法达到。因此，在我们尝试优化软件时，没有现实的性能标准可供比较。模型 FLOPS 利用率（MFU）度量标准测量实际实现的 TFLOPS 对比理论上的 TFLOPS。通常当 MFU 大约在 50% 左右时被认为是成功的。但这并不能告诉我们距离真正的最大吞吐量有多远。

此基准会扫描各种大型矩阵乘法形状，并报告它所能记录的最高可达 TFLOPS 值。由于变压器训练和部分推理工作负载主要由大规模矩阵乘法操作主导，可以将每个加速器上测量的最佳矩阵乘法 TFLOPS 作为粗略估计来判断这是可达最大矩阵乘法 FLOPS（MAMF）。现在可以使用模型可达矩阵乘法 FLOPS 利用率（MAMFU）代替之前的 MFU。

因此现在您可以在实际数值的基础上，将您的训练或推理中测量的 TFLOPS 进行比较。由于您将更加接近 100%，因此在何时停止优化会变得更容易知道。

目前支持的高端架构：
- NVIDIA：V100, A100, H100 等
- AMD：MI250, MI300X 等
- Intel Gaudi2+

公平性说明：
- 如果您能找到一种更高效的方法来检测最佳矩阵乘法 TFLOPS，可以将每个新加速器视为黑盒，请提交包含生成的日志文件的改进 PR。
- 同时如果您知道此基准应以特殊条件运行才能显示最佳结果（例如某些内核设置等），请提交 PR 添加这些特殊的指示。例如，对于 AMD MI300X 我被告知禁用 numa_balancing 可能会有所帮助。

### 建筑特定说明：

在运行基准前，请按照以下特殊设置说明操作以获得更佳效果：
- **NVIDIA A100 + AMD MI300X**

``` 
./mamf-finder.py --m_range 0 5376 256 --n_range 0 5376 256 --k_range 0 5376 256 --output_file=$(date +"%Y-%m-%d-%H:%M:%S").txt
```

- **NVIDIA H100**

``` 
./mamf-finder.py --m_range 0 20480 256 --n_range 0 20480 256 --k_range 0 20480 256 --output_file=$(date +"%Y-%m-%d-%H:%M:%S").txt
```

要更好地了解特定加速器下哪个形状能给出最高的矩阵乘法 FLOPS，请参阅 [向量和矩阵大小的整除性](../../../training/performance/README.md#vector-and-matrix-size-divisibility)。

### 结果

截至目前我收集到的测量结果可以在 [可达最大矩阵乘法 FLOPS 对比表](../#maximum-achievable-matmul-flops-comparison-table) 中找到。当我拥有特定加速器访问权限时，我会亲自运行基准测试；而在我没有这些条件的情况下，则是热心贡献者花费时间获取了这些数字。因此我非常感激 [那些人](../../../contributors.md)。

### 示例

1. 如果您想测量由您的训练所使用的特定形状，使用确切的形状而不是范围，比如 1024x1024x1024，可以运行：

``` 
./mamf-finder.py --m 1024 --n 1024 --k 1024 --output_file=$(date +"%Y-%m-%d-%H:%M:%S").txt
```

2. 如果您想要测量特定形状，例如 512x512x512，可以运行：

``` 
./mamf-finder.py --m 512 --n 512 --k 512 --output_file=$(date +"%Y-%m-%d-%H:%M:%S").txt
```

3. 如果您想测量 1024x1024 的固定形状，可以运行：

``` 
./mamf-finder.py --m 1024 --n 1024 --k_range 512 2048 256 --output_file=$(date +"%Y-%m-%d-%H:%M:%S").txt
```

这些示例演示了如何使用 `mamf-finder.py` 脚本进行特定形状的测量。

希望这些信息对您有帮助！如果有任何其他问题或需要进一步的帮助，请随时告诉我。