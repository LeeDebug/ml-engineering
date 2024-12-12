# 关闭 NVLink 性能测试

我们比较一下在小规模维基文本样本上训练 gpt2 语言模型的结果。

结果如下：

| NVLink | 时间（秒） |
| ------ | -------: |
| Y      |    101   |
| N      |    131   |

可以看见，使用 NVLink 完成训练速度快了约 23%。在第二次测试中，我们通过设置 `NCCL_P2P_DISABLE=1` 来指示 GPU 不使用 NVLink，而是使用 PCIe。

我们将使用 [HF Transformers 示例](https://github.com/huggingface/transformers/blob/58e3d23e97078f361a533b9ec4a6a2de674ea52a/examples/pytorch/language-modeling/run_clm.py)。

以下是完整的测试代码和输出：

```bash
# 使用 NVLink 的 DDP

rm -r /tmp/test-clm; CUDA_VISIBLE_DEVICES=0,1 python -m torch.distributed.launch \
--nproc_per_node 2 examples/pytorch/language-modeling/run_clm.py --model_name_or_path gpt2 \
--dataset_name wikitext --dataset_config_name wikitext-2-raw-v1 --do_train \
--output_dir /tmp/test-clm --per_device_train_batch_size 4 --max_steps 200

{"train_runtime": 101.9003, "train_samples_per_second": 1.963, "epoch": 0.69}

# 不使用 NVLink 的 DDP

rm -r /tmp/test-clm; CUDA_VISIBLE_DEVICES=0,1 NCCL_P2P_DISABLE=1 python -m torch.distributed.launch \
--nproc_per_node 2 examples/pytorch/language-modeling/run_clm.py --model_name_or_path gpt2 \
--dataset_name wikitext --dataset_config_name wikitext-2-raw-v1 --do_train
--output_dir /tmp/test-clm --per_device_train_batch_size 4 --max_steps 200

{"train_runtime": 131.4367, "train_samples_per_second": 1.522, "epoch": 0.69}
```

硬件：2 块 TITAN RTX，每块 24GB，带有 NVLink（使用 2 条 NVLink，`nvidia-smi topo -m` 中显示为 `NV2`）
软件：`pytorch-1.8-to-be` + `cuda-11.0` / `transformers==4.3.0.dev0`