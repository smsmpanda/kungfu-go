一个并发程序可以在一个处理器或者内核上使用多个线程来执行任务，但是<e style="font-weight:600;color:red">只有同一个程序在某个时间点同时运行在多核或者多处理器上才是真正的并行</e>。


不要使用全局变量或者共享内存，它们会给你的代码在并发运算的时候带来危险。



#### 在 Go 中
应用程序并发处理的部分被称作 goroutines（协程），它可以进行更有效的并发运算。在协程和操作系统线程之间并无一对一的关系：协程是根据一个或多个线程的可用性，映射（多路复用，执行于）在他们之上的；协程调度器在 Go 运行时很好的完成了这个工作。


通道实际上是类型化消息的队列：使数据得以传输。它是先进先出(FIFO) 的结构所以可以保证发送给他们的元素的顺序（有些人知道，通道可以比作 Unix shells 中的双向管道 (two-way pipe) ）。通道也是引用类型，所以我们使用 make() 函数来给它分配内存。这里先声明了一个字符串通道 ch1，然后创建了它（实例化）：

> var ch1 chan string
> ch1 = make(chan string)

##### 通道阻塞

默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束。可以想象一个无缓冲的通道在没有空间来保存数据的时候：必须要一个接收者准备好接收通道的数据然后发送者可以直接把数据发送给接收者。所以通道的发送/接收操作在对方准备好之前是阻塞的：

1）对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：如果 ch 中的数据无人接收，就无法再给通道传入其他数据：新的输入无法在通道非空的情况下传入。所以发送操作会等待 ch 再次变为可用状态：就是通道值被接收时（可以传入变量）。

2）对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。

尽管这看上去是非常严格的约束，实际在大部分情况下工作的很不错。


#### 通道的方向
通道类型可以用注解来表示它只发送或者只接收：
> var send_only chan<- int 		// channel can only receive data
> var recv_only <-chan int		// channel can only send data