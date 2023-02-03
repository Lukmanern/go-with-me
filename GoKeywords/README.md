#### Thanks To Santosh Shrestha's [Article](https://articles.wesionary.team/know-about-25-keywords-in-go-eca109855d4d)

# ALL KEYWORDS/ Syntaxes in Go

| Declaration(6) | CompositeTypes(4) | ControlFlow(13)      | FuncModifier(2) |
| -------------- | ----------------- | -------------------- | --------------- |
| const          | chan              | break                | go              |
| var            | interface         | continue             | defer           |
| func           | map               | default `(rare)`     |                 |
| type           | struct            | if                   |                 |
| import         |                   | else                 |                 |
| package        |                   | for                  |                 |
|                |                   | range                |                 |
|                |                   | return               |                 |
|                |                   | switch               |                 |
|                |                   | case                 |                 |
|                |                   | select `(rare)`      |                 |
|                |                   | goto `(rare)`        |                 |
|                |                   | fallthrough `(rare)` |                 |

note `(rare)` referred to not common to use. Just it.

## I. Declaration

#### The keywords are used to declare all kinds of code elements in Go.

### const

The `const` keyword is used to introduce a name for a scalar value like 2 or 3.14, etc.

### var

The `var` keyword is used to create the variables in the `Go` language.

### func

The `func` keyword is used to declare a function.

### type

We can use the `type` keyword to introduce a new struct type.

### import

The `import` keyword is used to import packages.

### package

The code is grouped as a unit in a package. The `package` keyword is used to define one.

## II. Composite Types

#### The keywords are used to indicate some composite type.

### chan

The `chan` keyword is used to define a channel. In `Go`, you are allowed to run parallel pieces of code simultaneously.

### interface

The `interface` keyword is used to specify a method set. A method set is a list of methods for a type.

### map

The `map` keyword defines a map type. A map os an unordered collection of the key-value pairs.

### struct

Struct is a collection of fields. We can use the `struct` keyword followed by the field declarations.

## III. Control Flow

#### The keywords are used to control flow of code.

### break

The `break` keyword lets you terminate a loop and continue execution of the rest of the code.

### case

This is a form of a `switch` construct. We specify a variable after the switch.

### continue

The `continue` keyword allows you to go back to the beginning of the `for` loop.

### default

The `default` statement is optional but you need to use either case or default within the switch statement. The switch jumps to the default value if the case does not match the controlling expression.

### else

The `else` keyword is used with the `if` keyword to execute an alternate statement if a certain condition is false.

### for

The `for` keyword is used to begin a for loop.

### if

The `if` statement is used to check a certain condition within a loop.

### range

The `range` keyword lets you iterate items over the list items like a map or an array.

### return

Go allows you to use the return values as variables and you can use the `return` keyword for that purpose.

### select

The `select` keyword lets a goroutine to wait during the simultaneous communication operations.

### switch

The `switch` statement is used to start a loop and use the if-else logic within the block.

### fallthrough

This keyword is used in a case within a switch statement. When we use this keyword, the next case is entered.

### goto

The `goto` keyword offers a jump to a labeled statement without any condition.

## IV. Function Modifier

#### The keywords also control flow keywords, but in other specific manners. They modify function calls.

### defer

The `defer` keyword is used to defer the execution of a function until the surrounding function executes.

### go

The `go` keyword triggers a goroutine which is managed by the golang-runtime.
