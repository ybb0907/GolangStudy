# sync介绍

## Go语言 sync.Once 源码分析

`sync.Once` 是 Go 标准库中保证**某个操作仅执行一次**的并发原语，适用于单例初始化、一次性配置加载等场景。

---

### 1. 结构体定义
```go
type Once struct {
    _    noCopy     // 通过静态检查防止结构体拷贝
    done atomic.Uint32  // 原子标志位（0未执行/1已执行）
    m    Mutex          // 互斥锁
}
```
###  2. 核心方法Do(func())
```go
func (o *Once) Do(f func()) {
    // Fast-path: 原子检查标志位
    if o.done.Load() == 0 {
        o.doSlow(f) // 未完成则进入慢速路径
    }
}

func (o *Once) doSlow(f func()) {
    o.m.Lock()
    defer o.m.Unlock()
    
    // 双重检查锁定 (Double-Check)
    if o.done.Load() == 0 {
        defer o.done.Store(1) // 延迟标记完成（即使f panic）
        f()                   // 执行目标函数
    }
}

```

### 关键逻辑

1. ​快速路径 (Fast Path)​​ 通过 atomic.Load 检查 done 标志位，避免锁竞争
   
2. 内联优化：简单逻辑可被编译器内联，减少函数调用开销
​​慢速路径 (Slow Path)​​

3. ​加锁​​：通过 Mutex 确保只有一个 goroutine 执行初始化
​
4. ​双重检查​​：防止前序 goroutine 已修改状态
​
5. ​延迟标记​​：defer o.done.Store(1) 保证 f() 完全执行后更新状态