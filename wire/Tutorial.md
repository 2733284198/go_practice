## Wire使用教程
通过例子学习使用Wire。这里我们要建立一个小的欢迎程序，用来了解如何使用Wire。

### 构建欢迎程序的第一步
+ 让我们创建一个小程序，它模拟一个事件，其中包含一个带有特定消息的欢迎来宾的问候语。
+ 我们创建了三种数据类型
    + 给迎宾员的信息
    + 传达信息的迎宾员
    + 以迎宾开始的活动
    ```Go
    type Message string

    type Greeter struct {
        // ... TBD
    }

    type Event struct {
        // ... TBD
    }
    ```
+ 现在，我们将创建一个简单的初始化器，总是返回一个硬编码的消息:
    ```Go
    func NewMessage() Message {
        return Message("Hi there!")
    }
    ```
+ 我们的迎宾员需要引用这条消息，我们也为欢迎者创建一个初始化器。
    ```Go
    func NewGreeter(m Message) Greeter {
        return Greeter{Message: m}
    }

    type Greeter struct {
        Message Message // <- adding a Message field
    }
    ```
+ 在初始化器中，我们给Greeter分配了一个Message字段。现在，当我们在Greeter上创建一个Greet方法时，我们可以使用这个消息:
    ```Go
    func (g Greeter) Greet() Message {
        return g.Message
    }
    ```
+ 接下来，Event需要有一个Greeter，所以我们也会为它创建一个初始化器。
    ```Go
    func NewEvent(g Greeter) Event {
        return Event{Greeter: g}
    }

    type Event struct {
        Greeter Greeter // <- adding a Greeter field
    }
    ```
+ 然后我们添加一个方法来启动事件:
    ```Go
    func (e Event) Start() {
        msg := e.Greeter.Greet()
        fmt.Println(msg)
    }
    ```
    Start方法是我们这个小应用程序的核心:它告诉欢迎者发出一个问候语，然后将该消息打印到屏幕上。
+ 现在我们已经准备好了应用程序的所有组件，让我们看看在不使用Wire的情况下初始化所有组件需要做些什么。我们的主函数是这样的:
    ```Go
    func main() {
        message := NewMessage()
        greeter := NewGreeter(message)
        event := NewEvent(greeter)

        event.Start()
    }
    ```
    首先我们创建一个消息，然后用这个消息创建一个欢迎器，最后我们用这个欢迎器创建一个事件。完成所有初始化之后，我们就可以开始事件了。
+ 我们使用依赖注入设计原则。实际上，这意味着我们传递每个组件需要的任何内容。这种设计风格有助于编写易于测试的代码，并使一种依赖关系与另一种依赖关系的交换变得容易。

### 使用Wire生成代码
依赖项注入的一个缺点是需要如此多的初始化步骤。让我们看看如何使用Wire使初始化组件的过程更顺畅。
+ 我们先简化我们的主函数：
    ```Go
    func main() {
        e := InitializeEvent()

        e.Start()
    }
    ```
+ 接下来，在一个名为wire的单独文件中。我们将定义InitializeEvent。这就是事情变得有趣的地方:
    ```Go
    // wire.go

    func InitializeEvent() Event {
        wire.Build(NewEvent, NewGreeter, NewMessage)
        return Event{}
    }
    ```
    与其依次初始化每个组件并将其传递给下一个组件，不如使用一个连接调用。构建传递我们想要使用的初始化器。在Wire中初始化器被称为“提供者”，即提供特定类型的函数。我们为事件添加一个零值作为返回值，以满足编译器的要求。注意，即使我们向事件添加值，Wire也会忽略它们。事实上，注入器的目的是提供关于使用哪些提供者来构造一个事件的信息，因此我们将在文件顶部的build约束中把它从最终的二进制代码中排除掉:
    ```Go
    //+build wireinject

    # 类似
    //+build wireinject
    // The build tag makes sure the stub is not built in the final build.
    package main
    ```
+ InitializeEvent是一个“注入器”。现在我们已经完成了注入器，可以使用wire命令行工具了。
    ```shell
    # 安装工具
    go get github.com/google/wire/cmd/wire
    ```
+ 然后在与上述代码相同的目录中，简单地运行wire。Wire将找到InitializeEvent注入器并生成一个函数，其主体由所有必要的初始化步骤填充。生成文件为wire_gen.go
    ```Go
    // wire_gen.go

    func InitializeEvent() Event {
        message := NewMessage()
        greeter := NewGreeter(message)
        event := NewEvent(greeter)
        return event
    }
    ```
    它看起来就像我们上面写的一样!现在，这是一个只有三个组件的简单示例，因此手工编写初始化器并不太困难。想象一下，对于复杂得多的组件，Wire是多么有用。当使用Wire时，我们将提交两个Wire。去wire_gen。转到源代码控制。

### 使用Wire进行更改
+ 为了演示Wire如何处理更复杂的设置，让我们重构Event的初始化器，以返回一个错误，然后看看会发生什么。
    ```Go
    func NewEvent(g Greeter) (Event, error) {
        if g.Grumpy {
            return Event{}, errors.New("could not create event: event greeter is grumpy")
        }
        return Event{Greeter: g}, nil
    }
    ```
+ 我们会说有时候一个迎宾员可能脾气暴躁，所以我们不能创建一个事件。NewGreeter的初始化现在看起来是这样的:
    ```Go
    func NewGreeter(m Message) Greeter {
        var grumpy bool
        if time.Now().Unix()%2 == 0 {
            grumpy = true
        }
        return Greeter{Message: m, Grumpy: grumpy}
    }
    ```
    我们已经向Greeter结构中添加了一个grumpy的字段，如果初始化器的调用时间与Unix时代相比是偶数秒，那么我们将创建一个暴躁的Greeter，而不是一个友好的Greeter。
    ```Go
    func (g Greeter) Greet() Message {
        if g.Grumpy {
            return Message("Go away!")
        }
        return g.Message
    }
    ```
+ 现在你明白了，一个脾气暴躁的迎宾员对一件事是多么地不利。因此NewEvent可能会失败。我们的主现在必须考虑到InitializeEvent可能会失败:
    ```Go
    func main() {
        e, err := InitializeEvent()
        if err != nil {
            fmt.Printf("failed to create event: %s\n", err)
            os.Exit(2)
        }
        e.Start()
    }
    ```
+ 我们还需要更新InitializeEvent，将错误类型添加到返回值:
    ```Go
    // wire.go

    func InitializeEvent() (Event, error) {
        wire.Build(NewEvent, NewGreeter, NewMessage)
        return Event{}, nil
    }
    ```
+  再次运行wire生成代码，如下
    ```Go
    // wire_gen.go

    func InitializeEvent(phrase string) (Event, error) {
        message := NewMessage(phrase)
        greeter := NewGreeter(message)
        event, err := NewEvent(greeter)
        if err != nil {
            return Event{}, err
        }
        return event, nil
    }
    ```
    Wire检查注入器的参数，看到我们向参数列表中添加了一个字符串(例如，phrase)，同样看到在所有提供程序中，NewMessage接受一个字符串，因此它将phrase传递给NewMessage。
