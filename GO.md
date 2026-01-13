# GO相关知识
## 守护进程执行
文件：console\cmd\daemon.go<br>
效果：常驻后台进程方式执行任务
## 定时器：Timer 和 Ticker
### time.Timer
文件：console\cmd\time.go<br>
简述：`time.Timer`是一次性定时器，用于在**指定时间后**执行一个函数。<br>

对比sleep
|    | sleep |     Timer |
| :----- | :--: | -------: |
| 阻塞 |  是  | 否（配合select使用） |
| 可中断 |  否  | timer.Stop() |
| 复用性 |  一次性使用  | 可通过 Reset() 复用 |

### time.Ticker
 `time.Ticker`是周期性定时器，用于在**指定时间间隔**内执行一个函数。