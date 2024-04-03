A [Pratt parser](https://en.wikipedia.org/wiki/Operator-precedence_parser) written in Go following the tutorial created by [Tyler Laceby](https://www.youtube.com/watch?v=V77J9l8N-P8).

Supports parsing of the following: 

Data Types: 
- `Null`
- `Boolean`
- `String`
- `Number` (`Int` or `Float`)

Grouping Symbols: 
- `[]`
- `()`
- `{}`

Logical Operators: 
| Operator | Meaning     |
|----------|-------------|
|`&&`| Logical `AND` |
|`\|\|`| Logical `OR` |
| `==` | Equivalence |
| `!` | Not |
| `!=` | Not equivalent |
| `>` | Greater Than |
| `>=` | Greater Than or Equal To |
| `<` | Less Than |
| `<=` | Less Than or Equal To |

Math: 
| Operator | Meaning |
|----------|-------------|
| `+` | Addition |
| `-` | Subtraction |
| `*` | Multiplication |
| `/` | Division |
| `^` | Exponential |
| `%` | Modulo |

Math Shorthand: 
| Operator | Meaning     |
|----------|-------------|
| `++` | Increment value by 1 |
| `+=` | Increment value by x |
| `--` | Decrement value by 1 |
| `-=` | Decrement value by x |
| `*=` | Multiply value by x | 
| `/=` | Divide value by x | 

Syntax: 
| Operator | Meaning     |
|----------|-------------|
| `=` | Assignment |
| `.` | TBD |
| `..` | Between (`1..10`)|
| `;` | Line termination |
| `,` | Item separation |
| `:` | TBD |
| `?` | TBD |

Keywords: 
- `LET`
- `CONST`
- `CLASS`
- `NEW`
- `TYPEOF`
- `FN`
- `IMPORT`
- `EXPORT`
- `FROM`
- `IN`
- `IF`/`ELSE`
- `FOR`/`FOREACH`
- `WHILE`
- `EOF`





