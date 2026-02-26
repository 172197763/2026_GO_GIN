# 菜鸟要知道的「GO interface」

> 🤫 本文基于 Go 技术栈进行解释相关概念及部分源码展示~

## ✨ interface 是什么东西？

和其他语言一样，**接口**就是一组方法声明。

我们可以看看以下代码：

```go
func main() {
    show(&Duck{})
    show(&Chicken{})
}

func show(b Bird) {
    b.Eat()
    b.Sound()
}

type Bird interface {
    Eat()
    Sound()
}

type Duck struct{}

func (d *Duck) Eat() {
    println("Duck eat")
}

func (d *Duck) Sound() {
    println("Duck sound")
}

type Chicken struct{}

func (c *Chicken) Eat() {
    println("Chicken eat")
}

func (c *Chicken) Sound() {
    println("Chicken sound")
}
```

### 输出结果：
```
Duck eat
Duck sound
Chicken eat
Chicken sound
```

在上述代码中，`Duck` 和 `Chicken` 都实现了 `Eat()` 和 `Sound()` 方法。因此，在任何需要 `Bird` 类型的地方，我们都可以直接传入 `Duck` 或 `Chicken` 实例。这种行为称为 **隐式实现**。

---

## ✨ 隐式实现

Go 中没有 `implement` 关键字，而是采用 **隐式实现** 来表示一个类型是否实现了某个接口（无需显式声明）。这与其他面向对象语言不同，后者通常要求显式声明类实现了某个接口。

Go 的设计哲学源自 **鸭子类型**：

> “如果它看起来像鸭子，游泳像鸭子，叫声像鸭子，那么它就是鸭子。”

在 Go 的上下文中，这意味着只要一个类型实现了接口的所有方法，就认为该类型实现了该接口。

---

## ✨ 包含方法的 interface 如何实现？

包含方法的 interface 被称为 `iface`，其数据结构如下：

```go
type iface struct {
    tab  *itab          // 接口方法表（interface table）
    data unsafe.Pointer // 指向实际数据的指针
}

type itab struct {
    inter *interfacetype // 接口的类型信息（如 io.Reader）
    _type *_type        // 具体类型的类型信息（如 *os.File）
    hash  uint32        // 接口 + 类型的哈希值，用于快速查找
    fun   []uintptr    // 方法地址数组（变长，实际大小由接口方法数决定）
}
```

### 示例代码讲解 iface 底层原理：

```go
type Bird interface {
    Eat()
    Sound()
}

type Duck struct {
    Name string
}

func (d *Duck) Iam() {
    println("Duck my name is", d.Name)
}

func (d *Duck) Eat() {
    println("Duck eat")
}

func (d *Duck) Sound() {
    println("Duck sound")
}

func main() {
    duck := &Duck{Name: "小鸭子"}
    show(duck)
    duck.Iam() // 调用的是 Duck 的 Iam 方法
}

func show(b Bird) {
    b.Eat()   // 根据 iface 中 fun 找到 Eat 方法地址并调用
    b.Sound() // 根据 iface 中 fun 找到 Sound 方法地址并调用
}
```

#### 执行流程说明：
1. 初始化 `duck` 变量时，Go 会根据 `Duck`结构创建一个实例。
2. 调用 `show`函数时，由于参数类型为 `Bird`，此时会生成一个 `Duck` 对应的 `Bird`接口实例。
3. 在 `show`函数内部，会通过 `iface` 中的 `fun` 数组找到 `Eat`和 `Sound` 方法的地址并调用。
4. 调用 `duck.Iam()` 时，仅会在 `duck` 自身中查找 `Iam()`方法，而不会经过`Bird`接口实例。

---

### ❓ 常见问题解答

#### Q1：什么时候会产生 `iface` 的实例？
只有当发生 **从具体类型 → 接口类型** 的转换时，编译器才会检查是否实现接口。如果不实现接口中的方法，则编译器会报错。

#### Q2：为什么 `itab` 中要存储 `fun` 数组？`_type` 中不是已经有方法地址了吗？
因为 `Duck` 还有其他方法（例如 `Iam()`），所以 `iface` 中的 `fun` 数组会根据 `Bird`接口中声明的方法，动态获取 `Duck` 中对应方法的地址。

---

### 🔍 额外科普

Go 中每个结构的方法只会存在一份。当方法被调用时，Go 会将实例隐式传递给方法，然后执行方法逻辑。

---

## ✨ `interface{}` —— 空接口

`interface{}` 是没有任何方法的接口，也称为空接口（`eface`）。它可以存储任意类型的数据。

由于 `interface{}` 没有声明任何方法，因此所有类型都可以被视为实现了 `interface{}`。

所以在日常开发中，经常可以看到这样的参数定义：

```go
func processData(data interface{}) {
    // 处理任意类型的数据
}
```

### `eface` 数据结构：

```go
type eface struct {
    _type *_type       // 描述当前接口的类型信息
    data  unsafe.Pointer // 指向实际数据的指针
}
```

### 注意事项：

- 当 `_type` 为 `nil` 且 `data`也为 `nil` 时，`eface` 实例才被认为是 `nil`。

#### 示例代码：

```go
var a interface{}
if a == nil {
    println("a is nil") // 输出
}

var b *int = nil
var data interface{} = b
if data == nil {
    println("data is nil") // 不输出
}
```

## ✨ 实际应用
### 1.嵌套与组合
示例代码：
```go
type Eater interface {
    Eat()
}

type Sounder interface {
    Sound()
}

// 组合接口
type Animal interface {
    Eater
    Sounder
}

type Dog struct{}

func (d *Dog) Eat() {
    println("Dog eat")
}

func (d *Dog) Sound() {
    println("Dog bark")
}

func main() {
    var animal Animal = &Dog{}
    animal.Eat()
    animal.Sound()
}
```

解释：
- Animal 接口通过嵌套 Eater 和 Sounder 接口，自动获得了这两个接口的所有方法。
- 如果一个类型实现了 Eater 和 Sounder 的所有方法，就等于实现了 Animal 接口。
### 2.类型断言
示例代码：
```
func process(i interface{}) {
    if s, ok := i.(string); ok {
        println("String value:", s)
    } else if n, ok := i.(int); ok {
        println("Int value:", n)
    } else {
        println("Unknown type")
    }
}

func main() {
    process("hello")
    process(42)
    process(true)
}
//结果
//String value: hello
//Int value: 42
//Unknown type
```
### 3.反射
示例代码：
```
import (
    "fmt"
    "reflect"
)

func inspect(i interface{}) {
    t := reflect.TypeOf(i)
    v := reflect.ValueOf(i)
    fmt.Printf("Type: %v, Value: %v\n", t, v)
}

func main() {
    inspect("hello")
    inspect(42)
    inspect([]int{1, 2, 3})
}
```