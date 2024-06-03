# Bamboo initial idea and thoughts

- Everything here is subject to change and most likely will change in the near future.
- These are just basic concepts and rough drafts of what the fundamentals of this language will be about.
- I will add more and make a sort of "theme" to this language later on.

```py
# Variable declarations
# let, or const
# const will be an actual constant (not like in JS)

# comments like in python using # symbol

# Variable types
let n: i32 = 1 # can be i16, i64, i128, u16, u32, u64, u128
let f: f64 = 1.0 # can be f32
let s: str = '1'
let b: bool = true
let x: nil = nil

const li: list = [1, 2, 3, ['hello', 'world'], true, nil]
const a: arr[i32] = [1, 2, 3, 4, 5, 6]
const t: tup = (1, 2, 3, 'hello', ['one', 1], (2, 'two'))

const m: map = {0: 10, 1: "10", "2": true}
const u: set[i32] = {1, 5, 3}

u = {1, 2, 3} # will return an error (const variables cannot be reassigned regardless if its annotated or not)

# Null type and union types
let nullable_variable: int | nil = nil  # Nullable variable initially assigned nil
nullable_variable = 10  # Reassigned to an integer value
nullable_variable = nil  # Reassigned to nil again

# Dynamic typing example (assigning a variable without a type annotation will give it "any" type)
let d = 10  # Variable 'a' can be dynamically typed based on subsequent assignments (under the hood it's let d: any = 10)
d = "hello"  # Valid: 'a' can be reassigned to a string
d = true  # Valid: 'a' can be reassigned to a boolean value
d = [1, 2, 3, 'hello'] # valid

# include
# Relational operators like >, < >=, <=, ==, !=
# Logical operators like: and, or, not
# Unary operators like: ++, --
# Arithemtic operators like: +, -, *, **, /, % (** is power)
# Bitwise operators like: ^, >>, <<
# Assignment operators like: +=, -=, *=, **=, /=, %=

# slicing in strings and lists
# works just like in python
let name = 'andre'
print(name[1:3]) # nd is the output

let l = [1, 2, 'andre', 'hi']
print(l[1:3]) # [2, 'andre'] is the output

# conditional statements
let is_true: bool = true
if is_true {
    print("It is true!")  # Output: It is true!
} elif not is_true {
    print("It is not true")
} else {
    print("Else")
}

let param: i32 = 5

switch param {
    case 5:
        print("five")
    case 6:
        print("six")
    case _:
        print(param)
}

# For loop example
let numbers: list = [10, 'a', 12, 'b']

for numb in numbers {
    print(numb)
}

for i, n in numbers {
    print(i, n)
}

# the range is from x to y-1
for i in (0...len(numbers)) {
    print(numbers[i])
}

# 5 is exclusive (so it goes like 0,1,2,3,4)
for i in (0...5, 1) {
    print(i)
}

for i in (5...0, -1) {
    print(i)
}

# While loop example
let count: int = 0

while count < 5 {
    print(count)
    count += 1
}

let counter: int = 0

while true {
    if counter == 5 {
        return "done"
    }
    counter += 1
}

# Function definition using fn keyword and strict typing
func greet(name: str) -> str {
    return f"Hello {name}!"
}

# Using the greet function with type annotations
result: str = greet("Andre") # outputs "Hello Andre!"

# error handling
try {
    let x: i32 = 5
    if x == "5" {
        print("x is a string")
    }
} catch {
    print("error caught")
} final {
    print("finish")
}
```
