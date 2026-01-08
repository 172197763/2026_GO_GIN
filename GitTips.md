## 合并发现冲突，想要撤回本次合并时
命令：git merge --abort<br>
效果：直接中止合并，回到合并前的状态

## cherry-pick发现冲突，想要撤回本次合并时
命令：git cherry-pick --abort<br>
效果：直接中止cherry-pick，回到cherry-pick前的状态

## 单独合并某一次提交
命令：git cherry-pick -m 1 提交hash（前4位即可）<br>
效果：只合并单次提交内容