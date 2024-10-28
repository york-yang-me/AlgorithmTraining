#!/bin/bash
# 读取输入 shell 脚本统计正整数的个数
# 如输入1 2 3 4 5 6 7 0应返回7
read -a arr
# 初始化正整数计数器
count=0
# 遍历输入的每个数字
for num in "${arr[@]}"; do
    # 检查是否为正整数
    if [[ $num =~ ^[1-9][0-9]*$ ]]; then
        count=$((count + 1))
    fi
done

# 输出结果
echo $count
exit 0