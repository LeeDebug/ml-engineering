# 使用微调示例从头重新训练HF Hub模型

HuggingFace Transformers 提供了非常出色的微调示例：[https://github.com/huggingface/transformers/tree/main/examples/pytorch]，涵盖了几乎所有模态，并且这些示例开箱即用。

**但是，如果想要从头开始重新训练而不是进行微调呢？**

这里有一个简单的技巧可以实现这一点。

我们将使用 `facebook/opt-1.3b` 模型作为示例，并计划以 bf16 训练模式为例：

```python
cat << EOT > prep-bf16.py
from transformers import AutoConfig, AutoModel, AutoTokenizer
import torch

mname = "facebook/opt-1.3b"

config = AutoConfig.from_pretrained(mname)
model = AutoModel.from_config(config, torch_dtype=torch.bfloat16)
tokenizer = AutoTokenizer.from_pretrained(mname)

path = "opt-1.3b-bf16"

model.save_pretrained(path)
tokenizer.save_pretrained(path)
EOT
```

现在运行：

```bash
python prep-bf16.py
```

这将在当前目录下创建一个文件夹 `opt-1.3b-bf16`，其中包含用于从头开始训练模型所需的一切。换句话说，这是一个预训练的模型，只是没有进行过任何实际训练。

根据你的需要调整上面的脚本以使用 `torch.float16` 或 `torch.float32`。

现在你可以像正常微调一样继续对保存的模型进行微调：

```bash
python -m torch.distributed.run \
--nproc_per_node=1 --nnode=1 --node_rank=0 \
--master_addr=127.0.0.1 --master_port=9901 \
examples/pytorch/language-modeling/run_clm.py --bf16 \
--seed 42 --model_name_or_path opt-1.3b-bf16 \
--dataset_name wikitext --dataset_config_name wikitext-103-raw-v1 \
--per_device_train_batch_size 12 --per_device_eval_batch_size 12 \
--gradient_accumulation_steps 1 --do_train --do_eval --logging_steps 10 \
--save_steps 1000 --eval_steps 100 --weight_decay 0.1 --num_train_epochs 1 \
--adam_beta1 0.9 --adam_beta2 0.95 --learning_rate 0.0002 --lr_scheduler_type \
linear --warmup_steps 500 --report_to tensorboard --output_dir save_dir
```

关键参数是：
```bash
--model_name_or_path opt-1.3b-bf16
```

其中 `opt-1.3b-bf16` 是你在上一步中生成的本地目录。

有时，你可以找到与原模型训练时使用的相同数据集，有时则需要使用替代数据集。其余的超参数通常可以在随附的论文或文档中找到。

总结来说，这个配方允许你利用微调示例从头开始重新训练在 [HF Hub] 上发现的任何模型（https://huggingface.co/models）。