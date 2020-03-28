#!bin/bash
git pull
git status
git add .
time=$(date "+%Y-%m-%d %H:%M:%S")
git commit -m '提交并同步代码: $time'
git push
echo "同步完成"
