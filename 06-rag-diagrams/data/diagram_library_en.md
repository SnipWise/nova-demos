# ASCII Diagram Library for RAG

## SECTION 1: FLOWCHARTS

### 001 - Simple linear flow (3 steps)
```
┌─────────┐     ┌─────────┐     ┌─────────┐
│ Input   │────▶│ Process │────▶│ Output  │
└─────────┘     └─────────┘     └─────────┘
```

### 002 - Vertical linear flow (4 steps)
```
┌─────────────┐
│   Start     │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  Step 1     │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  Step 2     │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│    End      │
└─────────────┘
```

### 003 - Flow with binary decision (if/else)
```
┌─────────────┐
│   Input     │
└──────┬──────┘
       │
       ▼
   ┌───────┐
  ╱         ╲
 ╱ Condition ╲
 ╲     ?     ╱
  ╲         ╱
   └───┬───┘
   Yes │ No
   ┌───┴───┐
   │       │
   ▼       ▼
┌─────┐ ┌─────┐
│ A   │ │ B   │
└──┬──┘ └──┬──┘
   │       │
   └───┬───┘
       │
       ▼
┌─────────────┐
│   Output    │
└─────────────┘
```

### 004 - Flow with loop (while/for)
```
┌─────────────┐
│   Start     │
└──────┬──────┘
       │
       ▼
┌─────────────┐◀──────┐
│  Processing │       │
└──────┬──────┘       │
       │              │
       ▼              │
   ┌───────┐          │
  ╱ Again?  ╲   Yes   │
  ╲         ╱─────────┘
   └───┬───┘
       │ No
       ▼
┌─────────────┐
│    End      │
└─────────────┘
```

### 005 - Flow with multiple branches (switch/case)
```
┌─────────────────┐
│     Input       │
└────────┬────────┘
         │
         ▼
    ┌─────────┐
   ╱   Type?   ╲
  ╱             ╲
 ╱───────────────╲
    │   │   │
    A   B   C
    │   │   │
    ▼   ▼   ▼
┌───┐ ┌───┐ ┌───┐
│ 1 │ │ 2 │ │ 3 │
└─┬─┘ └─┬─┘ └─┬─┘
  │     │     │
  └─────┼─────┘
        │
        ▼
┌─────────────────┐
│     Output      │
└─────────────────┘
```

### 006 - Parallel flow (fork/join)
```
┌─────────────────┐
│     Start       │
└────────┬────────┘
         │
    ═════╧═════    Fork
    │    │    │
    ▼    ▼    ▼
┌─────┐┌─────┐┌─────┐
│ T1  ││ T2  ││ T3  │
└──┬──┘└──┬──┘└──┬──┘
   │      │      │
   ═══════╪══════╛    Join
          │
          ▼
┌─────────────────┐
│      End        │
└─────────────────┘
```

### 007 - Flow with subprocess
```
┌─────────────────┐
│     Input       │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ ┌─────────────┐ │
│ │ Subprocess 1│ │
│ └──────┬──────┘ │
│        │        │
│        ▼        │
│ ┌─────────────┐ │
│ │ Subprocess 2│ │
│ └─────────────┘ │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     Output      │
└─────────────────┘
```

### 008 - Flow with error handling (try/catch)
```
┌─────────────┐
│   Start     │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│    Try      │
│ ┌─────────┐ │
│ │ Action  │─┼──────┐
│ └─────────┘ │      │
└──────┬──────┘      │ Error
       │             │
       │ OK          ▼
       │      ┌─────────────┐
       │      │   Catch     │
       │      │ ┌─────────┐ │
       │      │ │ Handler │ │
       │      │ └─────────┘ │
       │      └──────┬──────┘
       │             │
       └──────┬──────┘
              │
              ▼
┌─────────────────────┐
│       End           │
└─────────────────────┘
```

### 009 - Multi-step validation flow
```
┌───────────────┐
│    Input      │
└───────┬───────┘
        │
        ▼
┌───────────────┐     ┌─────────┐
│  Validation 1 │────▶│ Error   │
└───────┬───────┘ No  └─────────┘
        │ OK
        ▼
┌───────────────┐     ┌─────────┐
│  Validation 2 │────▶│ Error   │
└───────┬───────┘ No  └─────────┘
        │ OK
        ▼
┌───────────────┐     ┌─────────┐
│  Validation 3 │────▶│ Error   │
└───────┬───────┘ No  └─────────┘
        │ OK
        ▼
┌───────────────┐
│   Success     │
└───────────────┘
```

### 010 - Data pipeline
```
┌────────┐   ┌────────┐   ┌────────┐   ┌────────┐
│ Source │──▶│Extract │──▶│Transform──▶│  Load  │
└────────┘   └────────┘   └────────┘   └────────┘
    │            │            │            │
    ▼            ▼            ▼            ▼
┌────────┐   ┌────────┐   ┌────────┐   ┌────────┐
│  DB    │   │  Raw   │   │ Clean  │   │  DWH   │
└────────┘   └────────┘   └────────┘   └────────┘
```

---

## SECTION 9: SIMPLIFIED UML DIAGRAMS

### 054 - Simple class
```
┌────────────────────────┐
│        Person          │
├────────────────────────┤
│ - name: string         │
│ - age: int             │
│ - email: string        │
├────────────────────────┤
│ + getName(): string    │
│ + setName(s: string)   │
│ + getAge(): int        │
│ + setAge(a: int)       │
└────────────────────────┘
```

### 055 - Inheritance (two children)
```
         ┌───────────────┐
         │    Animal     │
         ├───────────────┤
         │ - name        │
         │ - age         │
         ├───────────────┤
         │ + eat()       │
         │ + sleep()     │
         └───────┬───────┘
                 △
        ┌────────┴────────┐
        │                 │
┌───────┴───────┐ ┌───────┴───────┐
│      Dog      │ │      Cat      │
├───────────────┤ ├───────────────┤
│ - breed       │ │ - indoor      │
├───────────────┤ ├───────────────┤
│ + bark()      │ │ + meow()      │
│ + fetch()     │ │ + scratch()   │
└───────────────┘ └───────────────┘
```

### 055b - Inheritance with THREE child classes (Lion, Owl, Fox from Animal)
```
UML class diagram showing THREE children inheriting from Character parent class.
This example shows Lion, Owl, and Fox classes ALL inheriting from Animal.
CRITICAL: The horizontal connector has THREE connection points (┌────┼────┐).
CRITICAL: THREE vertical lines (│) connect ALL three children (Lion, Owl, Fox).

         ┌───────────────┐
         │    Animal     │
         ├───────────────┤
         │ - name        │
         │ - age         │
         ├───────────────┤
         │ + move()      │
         └───────┬───────┘
                 △
        ┌────────┴────────┬─────────────────┐
        │                 │                 │
┌───────┴───────┐ ┌───────┴───────┐ ┌───────┴───────┐
│     Lion      │ │      Owl      │ │      Fox      │
├───────────────┤ ├───────────────┤ ├───────────────┤
│ - strength    │ │ - wingspan    │ │ - agility     │
│ - mane        │ │ - vision      │ │ - furColor    │
├───────────────┤ ├───────────────┤ ├───────────────┤
│ + roar()      │ │ + fly()       │ │ + sneak()     │
│ + hunt()      │ │ + hunt()      │ │ + dig()       │
└───────────────┘ └───────────────┘ └───────────────┘
```

### 056 - Composition and aggregation
```
┌───────────────┐      ┌───────────────┐
│    Company    │◆─────│  Department   │
└───────────────┘  1:n └───────────────┘
                              │
                              │ 1:n
                              ◇
                       ┌──────┴──────┐
                       │  Employee   │
                       └─────────────┘

◆ = Composition (strong)
◇ = Aggregation (weak)
```

### 057 - Interface and implementation
```
       ┌ ─ ─ ─ ─ ─ ─ ─ ┐
             <<interface>>
       │   Drawable    │
       ├───────────────┤
       │ + draw()      │
       │ + resize()    │
       └ ─ ─ ─ ┬ ─ ─ ─ ┘
               △
       ┌───────┴───────┐
       │               │
┌──────┴──────┐ ┌──────┴──────┐
│   Circle    │ │  Rectangle  │
├─────────────┤ ├─────────────┤
│ - radius    │ │ - width     │
│             │ │ - height    │
├─────────────┤ ├─────────────┤
│ + draw()    │ │ + draw()    │
│ + resize()  │ │ + resize()  │
└─────────────┘ └─────────────┘
```

### 058 - Association (bidirectional)
```
┌───────────────────┐          ┌───────────────────┐
│     Student       │          │      Course       │
├───────────────────┤          ├───────────────────┤
│ - studentId       │          │ - courseId        │
│ - name            │          │ - title           │
│ - courses[]       │◇────────◇│ - students[]      │
├───────────────────┤   n:n    ├───────────────────┤
│ + enroll()        │          │ + addStudent()    │
│ + drop()          │          │ + removeStudent() │
└───────────────────┘          └───────────────────┘
```

### 059 - Association (unidirectional)
```
┌───────────────────┐          ┌───────────────────┐
│      Order        │          │     Customer      │
├───────────────────┤          ├───────────────────┤
│ - orderId         │          │ - customerId      │
│ - date            │          │ - name            │
│ - customer        │─────────▶│ - email           │
├───────────────────┤   1:1    ├───────────────────┤
│ + getTotal()      │          │ + getOrders()     │
│ + addItem()       │          │ + updateInfo()    │
└───────────────────┘          └───────────────────┘
```

### 060 - Abstract class
```
         ┌───────────────────┐
         │   <<abstract>>    │
         │      Shape        │
         ├───────────────────┤
         │ # color: string   │
         │ # x: int          │
         │ # y: int          │
         ├───────────────────┤
         │ + move(x, y)      │
         │ + getArea()*      │
         │ + getPerimeter()* │
         └─────────┬─────────┘
                   △
         ┌─────────┴─────────┐
         │                   │
┌────────┴────────┐  ┌───────┴────────┐
│    Triangle     │  │    Square      │
├─────────────────┤  ├────────────────┤
│ - base          │  │ - side         │
│ - height        │  │                │
├─────────────────┤  ├────────────────┤
│ + getArea()     │  │ + getArea()    │
│ + getPerimeter()│  │ + getPerimeter()│
└─────────────────┘  └────────────────┘

* = abstract method
# = protected
```

### 061 - Multiple inheritance via interfaces
```
   ┌ ─ ─ ─ ─ ─ ─ ┐  ┌ ─ ─ ─ ─ ─ ─ ┐
     <<interface>>      <<interface>>
   │  Flyable    │  │  Swimmable  │
   ├─────────────┤  ├─────────────┤
   │ + fly()     │  │ + swim()    │
   └ ─ ─ ┬ ─ ─ ─ ┘  └ ─ ─ ┬ ─ ─ ─ ┘
         △                △
         └────────┬───────┘
                  │
         ┌────────┴────────┐
         │      Duck       │
         ├─────────────────┤
         │ - name          │
         │ - wingspan      │
         ├─────────────────┤
         │ + fly()         │
         │ + swim()        │
         │ + quack()       │
         └─────────────────┘
```

### 062 - Dependency relationship
```
┌────────────────┐
│   Controller   │
├────────────────┤         ┌────────────────┐
│ - service      │- - - - ▶│    Service     │
├────────────────┤  uses   ├────────────────┤
│ + handleReq()  │         │ + process()    │
└────────────────┘         └────────────────┘

- - - ▶ = dependency (uses)
```

### 063 - Singleton pattern
```
┌─────────────────────────────┐
│      DatabaseConnection     │
├─────────────────────────────┤
│ - instance: static          │
│ - connection: Connection    │
├─────────────────────────────┤
│ - DatabaseConnection()      │  ← private constructor
│ + getInstance(): static     │
│ + query(sql: string)        │
│ + close()                   │
└─────────────────────────────┘
```

### 064 - Observer pattern
```
       ┌ ─ ─ ─ ─ ─ ─ ─ ┐
             <<interface>>
       │   Observer    │
       ├───────────────┤
       │ + update()    │
       └ ─ ─ ─ △ ─ ─ ─ ┘
               │
               │ implements
       ┌───────┴────────────┐
       │                    │
┌──────┴──────┐   ┌─────────┴────────┐
│ConcreteObsA │   │ ConcreteObsB     │
├─────────────┤   ├──────────────────┤
│ + update()  │   │ + update()       │
└─────────────┘   └──────────────────┘
       △                   △
       │                   │
       │      notifies     │
       └──────────┬────────┘
                  │
         ┌────────┴────────┐
         │    Subject      │
         ├─────────────────┤
         │ - observers[]   │
         ├─────────────────┤
         │ + attach(obs)   │
         │ + detach(obs)   │
         │ + notify()      │
         └─────────────────┘
```

### 065 - Factory pattern
```
       ┌ ─ ─ ─ ─ ─ ─ ─ ┐
             <<interface>>
       │    Product    │
       ├───────────────┤
       │ + operation() │
       └ ─ ─ ─ △ ─ ─ ─ ┘
               │
       ┌───────┴────────┐
       │                │
┌──────┴──────┐  ┌──────┴──────┐
│  ProductA   │  │  ProductB   │
├─────────────┤  ├─────────────┤
│+operation() │  │+operation() │
└─────────────┘  └─────────────┘
                        △
                        │ creates
                        │
                ┌───────┴────────┐
                │    Factory     │
                ├────────────────┤
                │+createProduct()│
                │  (type): Prod  │
                └────────────────┘
```

### 066 - Strategy pattern
```
       ┌ ─ ─ ─ ─ ─ ─ ─ ─ ┐
             <<interface>>
       │    Strategy     │
       ├─────────────────┤
       │ + execute()     │
       └ ─ ─ ─ △ ─ ─ ─ ─ ┘
               │
       ┌───────┴────────┬───────────────┐
       │                │               │
┌──────┴──────┐  ┌──────┴──────┐ ┌──────┴──────┐
│StrategyA    │  │StrategyB    │ │StrategyC    │
├─────────────┤  ├─────────────┤ ├─────────────┤
│ + execute() │  │ + execute() │ │ + execute() │
└─────────────┘  └─────────────┘ └─────────────┘
                         △
                         │ uses
                         │
                 ┌───────┴────────┐
                 │    Context     │
                 ├────────────────┤
                 │ - strategy     │
                 ├────────────────┤
                 │ + setStrategy()│
                 │ + doWork()     │
                 └────────────────┘
```

### 067 - Decorator pattern
```
       ┌ ─ ─ ─ ─ ─ ─ ─ ─ ┐
             <<interface>>
       │   Component     │
       ├─────────────────┤
       │ + operation()   │
       └ ─ ─ ─ △ ─ ─ ─ ─ ┘
               │
       ┌───────┴────────────────┐
       │                        │
┌──────┴───────┐       ┌────────┴────────┐
│ConcreteComp  │       │   Decorator     │
├──────────────┤       ├─────────────────┤
│+operation()  │       │ - component     │
└──────────────┘       │ + operation()   │
                       └────────△────────┘
                                │
                 ┌──────────────┴──────────────┐
                 │                             │
          ┌──────┴──────┐            ┌─────────┴────────┐
          │DecoratorA   │            │  DecoratorB      │
          ├─────────────┤            ├──────────────────┤
          │+operation() │            │ + operation()    │
          └─────────────┘            └──────────────────┘
```

### 068 - Adapter pattern
```
┌──────────────┐          ┌ ─ ─ ─ ─ ─ ─ ─ ┐
│    Client    │                Target
├──────────────┤          │  <<interface>>│
│              │──────────▶ ├─────────────┤
└──────────────┘          │ + request()  │
                          └ ─ ─ △ ─ ─ ─ ─ ┘
                                │
                                │ implements
                          ┌─────┴────────┐
                          │   Adapter    │
                          ├──────────────┤
                          │ - adaptee    │─────────┐
                          ├──────────────┤         │
                          │ + request()  │         │
                          └──────────────┘         │
                                                   │
                                                   ▼
                                          ┌────────────────┐
                                          │    Adaptee     │
                                          ├────────────────┤
                                          │+specificRequest│
                                          └────────────────┘
```

### 069 - Composite pattern
```
       ┌ ─ ─ ─ ─ ─ ─ ─ ─ ┐
             <<interface>>
       │   Component     │
       ├─────────────────┤
       │ + operation()   │
       │ + add(c)        │
       │ + remove(c)     │
       └ ─ ─ ─ △ ─ ─ ─ ─ ┘
               │
       ┌───────┴────────────────┐
       │                        │
┌──────┴───────┐       ┌────────┴────────┐
│    Leaf      │       │   Composite     │
├──────────────┤       ├─────────────────┤
│+operation()  │       │ - children[]    │
└──────────────┘       │ + operation()   │
                       │ + add(c)        │
                       │ + remove(c)     │
                       │ + getChild(i)   │
                       └─────────────────┘
```

### 070 - Enumeration in class
```
┌────────────────────────┐
│      <<enum>>          │
│       Status           │
├────────────────────────┤
│ PENDING                │
│ IN_PROGRESS            │
│ COMPLETED              │
│ CANCELLED              │
└────────────────────────┘
          △
          │ uses
          │
┌─────────┴──────────────┐
│       Task             │
├────────────────────────┤
│ - id: int              │
│ - title: string        │
│ - status: Status       │
├────────────────────────┤
│ + updateStatus(s)      │
│ + isCompleted(): bool  │
└────────────────────────┘
```

### 071 - Generic class (template)
```
┌──────────────────────────┐
│   List<T>                │
├──────────────────────────┤
│ - items: T[]             │
│ - size: int              │
├──────────────────────────┤
│ + add(item: T)           │
│ + get(index): T          │
│ + remove(index)          │
│ + size(): int            │
└──────────────────────────┘
          △
          │
  ┌───────┴────────┐
  │                │
┌─┴────────────┐ ┌─┴─────────────┐
│List<String>  │ │ List<Integer> │
└──────────────┘ └───────────────┘
```

### 072 - Inner/Nested class
```
┌─────────────────────────────────────┐
│          OuterClass                 │
├─────────────────────────────────────┤
│ - outerField: string                │
├─────────────────────────────────────┤
│ + outerMethod()                     │
│ ┌─────────────────────────────────┐ │
│ │       InnerClass                │ │
│ ├─────────────────────────────────┤ │
│ │ - innerField: int               │ │
│ ├─────────────────────────────────┤ │
│ │ + innerMethod()                 │ │
│ │ + accessOuter()                 │ │
│ └─────────────────────────────────┘ │
└─────────────────────────────────────┘
```

### 073 - Association class
```
┌──────────────┐                 ┌──────────────┐
│   Student    │                 │   Course     │
├──────────────┤                 ├──────────────┤
│ - studentId  │◆───────────────◆│ - courseId   │
│ - name       │                 │ - title      │
└──────────────┘                 └──────────────┘
       │                                │
       │           ┌─────────────────┐  │
       └──────────▶│  Enrollment     │◀─┘
                   ├─────────────────┤
                   │ - date          │
                   │ - grade         │
                   │ - semester      │
                   ├─────────────────┤
                   │ + getGrade()    │
                   │ + setGrade(g)   │
                   └─────────────────┘
```

### 074 - Static members
```
┌────────────────────────────────┐
│         MathUtils              │
├────────────────────────────────┤
│ + PI: static final = 3.14159   │
│ + E: static final = 2.71828    │
├────────────────────────────────┤
│ + max(a, b): static            │
│ + min(a, b): static            │
│ + abs(x): static               │
│ + sqrt(x): static              │
└────────────────────────────────┘

Underlined = static
```

### 075 - Multiple associations from one class
```
                  ┌──────────────┐
                  │   Customer   │
                  ├──────────────┤
                  │ - customerId │
                  │ - name       │
                  └───┬──────┬───┘
                      │      │
           ┌──────────┘      └──────────┐
           │ 1:n                      1:1│
           ▼                             ▼
    ┌──────────────┐            ┌────────────────┐
    │    Order     │            │    Address     │
    ├──────────────┤            ├────────────────┤
    │ - orderId    │            │ - street       │
    │ - date       │            │ - city         │
    └──────────────┘            │ - zipCode      │
                                └────────────────┘
```

### 076 - Inheritance with four children
```
This example shows FOUR children inheriting from Vehicle parent class.
CRITICAL: The horizontal connector has FOUR connection points.
CRITICAL: FOUR vertical lines (│) connect ALL four children.

              ┌───────────────┐
              │    Vehicle    │
              ├───────────────┤
              │ - speed       │
              │ - fuel        │
              ├───────────────┤
              │ + start()     │
              │ + stop()      │
              └───────┬───────┘
                      △
     ┌────────────────┼────────────────┬────────────────┐
     │                │                │                │
┌────┴─────┐  ┌───────┴──────┐  ┌──────┴─────┐  ┌──────┴─────┐
│   Car    │  │  Motorcycle  │  │   Truck    │  │    Bus     │
├──────────┤  ├──────────────┤  ├────────────┤  ├────────────┤
│ - doors  │  │ - hasSidecar │  │ - capacity │  │ - seats    │
├──────────┤  ├──────────────┤  ├────────────┤  ├────────────┤
│ + park() │  │ + wheelie()  │  │ + load()   │  │ + board()  │
└──────────┘  └──────────────┘  └────────────┘  └────────────┘
```

### 077 - Interface with multiple interfaces (multiple inheritance)
```
┌ ─ ─ ─ ─ ─ ─ ┐  ┌ ─ ─ ─ ─ ─ ─ ┐  ┌ ─ ─ ─ ─ ─ ─ ┐
  <<interface>>    <<interface>>    <<interface>>
│  Readable   │  │  Writable   │  │  Closeable  │
├─────────────┤  ├─────────────┤  ├─────────────┤
│ + read()    │  │ + write()   │  │ + close()   │
└ ─ ─ △ ─ ─ ─ ┘  └ ─ ─ △ ─ ─ ─ ┘  └ ─ ─ △ ─ ─ ─ ┘
      │                │                │
      └────────────────┼────────────────┘
                       │
              ┌────────┴────────┐
              │      File       │
              ├─────────────────┤
              │ - path: string  │
              │ - mode: string  │
              ├─────────────────┤
              │ + read()        │
              │ + write()       │
              │ + close()       │
              └─────────────────┘
```

### 078 - Repository pattern (DAO)
```
       ┌ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ┐
             <<interface>>
       │    Repository<T>    │
       ├─────────────────────┤
       │ + findById(id): T   │
       │ + findAll(): T[]    │
       │ + save(entity: T)   │
       │ + delete(id)        │
       └ ─ ─ ─ ─ △ ─ ─ ─ ─ ─ ┘
                 │
                 │ implements
       ┌─────────┴──────────┐
       │  UserRepository    │
       ├────────────────────┤
       │ - db: Database     │───────┐
       ├────────────────────┤       │
       │ + findById(id)     │       │
       │ + findAll()        │       │ uses
       │ + save(user)       │       │
       │ + delete(id)       │       │
       │ + findByEmail(e)   │       ▼
       └────────────────────┘  ┌──────────┐
                  │            │   User   │
                  │ manages    ├──────────┤
                  └───────────▶│ - id     │
                               │ - name   │
                               │ - email  │
                               └──────────┘
```

### 079 - MVC pattern
```
┌──────────────┐         ┌──────────────┐
│    View      │         │  Controller  │
├──────────────┤         ├──────────────┤
│ + render()   │◀────────│ - model      │
│ + update()   │ updates │ - view       │
└──────┬───────┘         ├──────────────┤
       │                 │ + handleReq()│
       │ user input      │ + updateView │
       └────────────────▶└──────┬───────┘
                                │
                                │ updates
                                ▼
                         ┌──────────────┐
                         │    Model     │
                         ├──────────────┤
                         │ - data       │
                         ├──────────────┤
                         │ + getData()  │
                         │ + setData()  │
                         │ + notify()   │
                         └──────────────┘
```

### 080 - Class with constants and enums
```
┌────────────────────────────────┐
│         Configuration          │
├────────────────────────────────┤
│ + MAX_RETRIES: final = 3       │
│ + TIMEOUT: final = 5000        │
│ + DEFAULT_PORT: final = 8080   │
├────────────────────────────────┤
│ + getConfig(): Map             │
│ + validate(): boolean          │
└────────────────────────────────┘
          │
          │ has
          ▼
┌────────────────────────────────┐
│      <<enum>>                  │
│      Environment               │
├────────────────────────────────┤
│ DEVELOPMENT                    │
│ STAGING                        │
│ PRODUCTION                     │
├────────────────────────────────┤
│ + toString(): string           │
└────────────────────────────────┘
```

---

# END OF DIAGRAM LIBRARY

Total: 80 diagrams covering:
- Flows and processes (10)
- System architecture (10)
- Sequences (8)
- Data structures (10)
- Network (4)
- CI/CD (4)
- Machine Learning (4)
- State machines (3)
- UML - Class diagrams (27)
  * Basic class structures (054-057)
  * Associations (058-059, 073, 075)
  * Inheritance (055, 055b, 060, 076)
  * Interfaces (057, 061, 077)
  * Design patterns (063-069, 078-079)
  * Advanced features (070-072, 074, 080)
- Database (3)
- Kubernetes/Containers (4)
- Agent/LLM (5)
- Miscellaneous (6)
