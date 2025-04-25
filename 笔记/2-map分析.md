# Map

## 原理
map底层是用哈希表实现的，其中：切片、函数等不支持==运算的不能作为key
  
Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序。

## 遍历方式
采用for range方式进行遍历，和切片一样
```go
func equal(x, y map[string]int) bool {
    if len(x) != len(y) {
        return false
    }
    for k, xv := range x {
        if yv, ok := y[k]; !ok || yv != xv {
            return false
        }
    }
    return true
}
```