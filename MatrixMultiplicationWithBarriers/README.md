# 两个 Barrier 的作用

每一轮让所有 goroutine 同时开始计算并确保全部完成后再进入下一轮

# Barrier 大小为 matrixSize + 1

是因为除了 matrixSize 个 worker goroutine main 也是一个参与者进行矩阵生成的操作

# 关于main里生成矩阵的之后的两个wait

```
main          worker0        worker1 ... workerN
──gen──┐
│      wait           wait          wait
workStart.Wait() ——开闸——> compute ——> workComplete.Wait()
│                                   │
workComplete.Wait() <——开闸——< done        done
```

第一个 workStart.Wait() 把 main 阻塞住直到 所有 worker 也调用了 workStart.Wait()
此时大家同时越过起跑线；紧接着的 workComplete.Wait() 又把 main 阻塞住直到
所有 worker 调用 workComplete.Wait() 此时大家同时越过终点线

# 整体流程

![flow.png](flow.png)

