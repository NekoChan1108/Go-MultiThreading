# `sync.Cond` 三大 API

## `Wait()`

- **作用**
    - 原子地 **释放** 关联锁 (`L.Unlock`)
    - **挂起** 当前 goroutine
    - 被唤醒后 **重新加锁** (`L.Lock`) 再返回
- **典型用法**

```  go
  cond.L.Lock()
  for !condition {
      cond.Wait() // 内部：先 Unlock → 睡眠 → 醒来再 Lock
  }
  // 条件满足，继续
  cond.L.Unlock()
  ```

- **注意⚠️**
    - **必须放在 for 循环里防止虚假唤醒**
    - **调用前必须已持有 cond.L 锁**

## `Signal()`

- **作用**
    - 唤醒 **一个** 正在 `Wait()` 的 goroutine（随机选择）
- **典型用法**

``` go
  cond.L.Lock()
  condition = true
  cond.Signal() // 只叫醒一个等待者
  cond.L.Unlock()
  ```

## `Broadcast()`

- **作用**
    - 唤醒 所有 正在 `Wait()` 的 goroutines
- **典型用法**

``` go
  cond.L.Lock()
  condition = true
  cond.Broadcast() // 叫醒全部等待者
  cond.L.Unlock()
```

* Wait：先解锁 → 睡觉 → 醒来再锁
* Signal：随机叫一个起床
* Broadcast：叫所有人起床