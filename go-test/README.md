#Profiling Go

#### Go的实现是一个并行的标记和扫描垃圾收集器。 在传统的标记扫描模式中，垃圾收集器将在程序检测到不可达对象时再次清除它（即释放存储器），从而使程序停止运行（即“停止世界”）。 这是为了防止运行程序在识别/清理阶段期间最终移动参考文献的并发症。 当GC运行时，这也会导致程序用户的延迟和其他问题。 随着Go GC同时执行，所以即使GC运行，用户也不会注意到暂停或延迟。
## 分析类型

### 我们有几种可用于监控性能的方法
#### 计时器：有助于基准测试，以及比较之前和之后的修复。
#### 剖析器：用于高级验证。
## 分析步骤

## 无论您用于分析的工具如何，一般的经验法则是：

### 确定高层次的瓶颈
#### 例如，您可能会注意到长时间运行的函数调用。
### 减少操作
#### 看看花费的时间，或电话的数量，并找出一种替代方法。
#### 看看内存分配的数量，找出一种替代方法。
### 向下钻
#### 使用一个可以在较低级别提供数据的工具。
#### 考虑更多性能更高的算法或数据结构。
#### 也可能有更简单的解决方案。

