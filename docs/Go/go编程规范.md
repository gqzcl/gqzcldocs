#  The Go Programming Language Specification 编程语言规范

## Version of Jan 14, 2020

## 2020年1月14日

| [Introduction 引言](http://docscn.studygolang.com/ref/spec#Introduction)[Notation 记号法](http://docscn.studygolang.com/ref/spec#Notation)[Source code representation 源代码表示](http://docscn.studygolang.com/ref/spec#Source_code_representation)[Characters 人物](http://docscn.studygolang.com/ref/spec#Characters)[Letters and digits 字母和数字](http://docscn.studygolang.com/ref/spec#Letters_and_digits)[Lexical elements 词汇元素](http://docscn.studygolang.com/ref/spec#Lexical_elements)[Comments 评论](http://docscn.studygolang.com/ref/spec#Comments)[Tokens 代币](http://docscn.studygolang.com/ref/spec#Tokens)[Semicolons 分号](http://docscn.studygolang.com/ref/spec#Semicolons)[Identifiers 标识符](http://docscn.studygolang.com/ref/spec#Identifiers)[Keywords 关键词](http://docscn.studygolang.com/ref/spec#Keywords)[Operators and punctuation 操作符和标点符号](http://docscn.studygolang.com/ref/spec#Operators_and_punctuation)[Integer literals 整数字面值](http://docscn.studygolang.com/ref/spec#Integer_literals)[Floating-point literals 浮点文字](http://docscn.studygolang.com/ref/spec#Floating-point_literals)[Imaginary literals 虚构的文字](http://docscn.studygolang.com/ref/spec#Imaginary_literals)[Rune literals 符文字面值](http://docscn.studygolang.com/ref/spec#Rune_literals)[String literals 字符串文字](http://docscn.studygolang.com/ref/spec#String_literals)[Constants 常量](http://docscn.studygolang.com/ref/spec#Constants)[Variables 变量](http://docscn.studygolang.com/ref/spec#Variables)[Types 类型](http://docscn.studygolang.com/ref/spec#Types)[Method sets 方法集](http://docscn.studygolang.com/ref/spec#Method_sets)[Boolean types 布尔类型](http://docscn.studygolang.com/ref/spec#Boolean_types)[Numeric types 数字类型](http://docscn.studygolang.com/ref/spec#Numeric_types)[String types 字符串类型](http://docscn.studygolang.com/ref/spec#String_types)[Array types 数组类型](http://docscn.studygolang.com/ref/spec#Array_types)[Slice types 切片类型](http://docscn.studygolang.com/ref/spec#Slice_types)[Struct types 结构类型](http://docscn.studygolang.com/ref/spec#Struct_types)[Pointer types 指针类型](http://docscn.studygolang.com/ref/spec#Pointer_types)[Function types 功能类型](http://docscn.studygolang.com/ref/spec#Function_types)[Interface types 接口类型](http://docscn.studygolang.com/ref/spec#Interface_types)[Map types 地图类型](http://docscn.studygolang.com/ref/spec#Map_types)[Channel types 通道类型](http://docscn.studygolang.com/ref/spec#Channel_types)[Properties of types and values 类型和值的属性](http://docscn.studygolang.com/ref/spec#Properties_of_types_and_values)[Type identity 类型标识](http://docscn.studygolang.com/ref/spec#Type_identity)[Assignability 可转让性](http://docscn.studygolang.com/ref/spec#Assignability)[Representability 具有代表性](http://docscn.studygolang.com/ref/spec#Representability)[Blocks 积木](http://docscn.studygolang.com/ref/spec#Blocks)[Declarations and scope 声明和范围](http://docscn.studygolang.com/ref/spec#Declarations_and_scope)[Label scopes 标签范围](http://docscn.studygolang.com/ref/spec#Label_scopes)[Blank identifier 空白标识符](http://docscn.studygolang.com/ref/spec#Blank_identifier)[Predeclared identifiers 预先声明的标识符](http://docscn.studygolang.com/ref/spec#Predeclared_identifiers)[Exported identifiers 导出标识符](http://docscn.studygolang.com/ref/spec#Exported_identifiers)[Uniqueness of identifiers 标识符的唯一性](http://docscn.studygolang.com/ref/spec#Uniqueness_of_identifiers)[Constant declarations 不断的声明](http://docscn.studygolang.com/ref/spec#Constant_declarations)[Iota 女名女子名](http://docscn.studygolang.com/ref/spec#Iota)[Type declarations 类型声明](http://docscn.studygolang.com/ref/spec#Type_declarations)[Variable declarations 变量声明](http://docscn.studygolang.com/ref/spec#Variable_declarations)[Short variable declarations 短变量声明](http://docscn.studygolang.com/ref/spec#Short_variable_declarations)[Function declarations 函数声明](http://docscn.studygolang.com/ref/spec#Function_declarations)[Method declarations 方法声明](http://docscn.studygolang.com/ref/spec#Method_declarations)[Expressions 表达方式](http://docscn.studygolang.com/ref/spec#Expressions)[Operands 操作数](http://docscn.studygolang.com/ref/spec#Operands)[Qualified identifiers 限定标识符](http://docscn.studygolang.com/ref/spec#Qualified_identifiers)[Composite literals 合成文字](http://docscn.studygolang.com/ref/spec#Composite_literals)[Function literals 函数文字](http://docscn.studygolang.com/ref/spec#Function_literals)[Primary expressions 主要表达式](http://docscn.studygolang.com/ref/spec#Primary_expressions)[Selectors 选择器](http://docscn.studygolang.com/ref/spec#Selectors)[Method expressions 方法表达式](http://docscn.studygolang.com/ref/spec#Method_expressions)[Method values 方法值](http://docscn.studygolang.com/ref/spec#Method_values) | [Index expressions 索引表达式](http://docscn.studygolang.com/ref/spec#Index_expressions)[Slice expressions 片段表达式](http://docscn.studygolang.com/ref/spec#Slice_expressions)[Type assertions 输入断言](http://docscn.studygolang.com/ref/spec#Type_assertions)[Calls 电话](http://docscn.studygolang.com/ref/spec#Calls)[Passing arguments to ... parameters 将参数传递给... 参数](http://docscn.studygolang.com/ref/spec#Passing_arguments_to_..._parameters)[Operators 操作员](http://docscn.studygolang.com/ref/spec#Operators)[Arithmetic operators 算术运算符](http://docscn.studygolang.com/ref/spec#Arithmetic_operators)[Comparison operators 比较操作符](http://docscn.studygolang.com/ref/spec#Comparison_operators)[Logical operators 逻辑运算符](http://docscn.studygolang.com/ref/spec#Logical_operators)[Address operators 地址操作符](http://docscn.studygolang.com/ref/spec#Address_operators)[Receive operator 接收操作员](http://docscn.studygolang.com/ref/spec#Receive_operator)[Conversions 转化率](http://docscn.studygolang.com/ref/spec#Conversions)[Constant expressions 常量表达式](http://docscn.studygolang.com/ref/spec#Constant_expressions)[Order of evaluation 评估顺序](http://docscn.studygolang.com/ref/spec#Order_of_evaluation)[Statements 声明](http://docscn.studygolang.com/ref/spec#Statements)[Terminating statements 终止语句](http://docscn.studygolang.com/ref/spec#Terminating_statements)[Empty statements 空洞的陈述](http://docscn.studygolang.com/ref/spec#Empty_statements)[Labeled statements 标记语句](http://docscn.studygolang.com/ref/spec#Labeled_statements)[Expression statements 表达式语句](http://docscn.studygolang.com/ref/spec#Expression_statements)[Send statements 发送声明](http://docscn.studygolang.com/ref/spec#Send_statements)[IncDec statements IncDec 声明](http://docscn.studygolang.com/ref/spec#IncDec_statements)[Assignments 工作分配](http://docscn.studygolang.com/ref/spec#Assignments)[If statements 如果语句](http://docscn.studygolang.com/ref/spec#If_statements)[Switch statements 开关语句](http://docscn.studygolang.com/ref/spec#Switch_statements)[For statements 对于声明](http://docscn.studygolang.com/ref/spec#For_statements)[Go statements 发布声明](http://docscn.studygolang.com/ref/spec#Go_statements)[Select statements 选择语句](http://docscn.studygolang.com/ref/spec#Select_statements)[Return statements 返回语句](http://docscn.studygolang.com/ref/spec#Return_statements)[Break statements 打破陈述](http://docscn.studygolang.com/ref/spec#Break_statements)[Continue statements 继续陈述](http://docscn.studygolang.com/ref/spec#Continue_statements)[Goto statements 后藤语句](http://docscn.studygolang.com/ref/spec#Goto_statements)[Fallthrough statements 失败的陈述](http://docscn.studygolang.com/ref/spec#Fallthrough_statements)[Defer statements 推迟语句](http://docscn.studygolang.com/ref/spec#Defer_statements)[Built-in functions 内置功能](http://docscn.studygolang.com/ref/spec#Built-in_functions)[Close 关闭](http://docscn.studygolang.com/ref/spec#Close)[Length and capacity 长度和容量](http://docscn.studygolang.com/ref/spec#Length_and_capacity)[Allocation 分配](http://docscn.studygolang.com/ref/spec#Allocation)[Making slices, maps and channels 制作切片、地图和通道](http://docscn.studygolang.com/ref/spec#Making_slices_maps_and_channels)[Appending to and copying slices 追加并复制切片](http://docscn.studygolang.com/ref/spec#Appending_and_copying_slices)[Deletion of map elements 删除地图元素](http://docscn.studygolang.com/ref/spec#Deletion_of_map_elements)[Manipulating complex numbers 操纵复数](http://docscn.studygolang.com/ref/spec#Complex_numbers)[Handling panics 处理恐慌](http://docscn.studygolang.com/ref/spec#Handling_panics)[Bootstrapping 自力更生](http://docscn.studygolang.com/ref/spec#Bootstrapping)[Packages 包装](http://docscn.studygolang.com/ref/spec#Packages)[Source file organization 源文件组织](http://docscn.studygolang.com/ref/spec#Source_file_organization)[Package clause 包装条款](http://docscn.studygolang.com/ref/spec#Package_clause)[Import declarations 进口报关单](http://docscn.studygolang.com/ref/spec#Import_declarations)[An example package 一个示例包](http://docscn.studygolang.com/ref/spec#An_example_package)[Program initialization and execution 程序的初始化和执行](http://docscn.studygolang.com/ref/spec#Program_initialization_and_execution)[The zero value 零值](http://docscn.studygolang.com/ref/spec#The_zero_value)[Package initialization 包初始化](http://docscn.studygolang.com/ref/spec#Package_initialization)[Program execution 程序执行](http://docscn.studygolang.com/ref/spec#Program_execution)[Errors 错误](http://docscn.studygolang.com/ref/spec#Errors)[Run-time panics 运行时恐慌](http://docscn.studygolang.com/ref/spec#Run_time_panics)[System considerations 系统方面的考虑](http://docscn.studygolang.com/ref/spec#System_considerations)[Package unsafe 包装不安全](http://docscn.studygolang.com/ref/spec#Package_unsafe)[Size and alignment guarantees 尺寸和对齐保证](http://docscn.studygolang.com/ref/spec#Size_and_alignment_guarantees) |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
|                                                              |                                                              |

## Introduction 引言

This is a reference manual for the Go programming language. For more information and other documents, see [golang.org](http://docscn.studygolang.com/).

这是 Go 编程语言的参考手册。有关更多信息和其他文档，请参阅 golang. org。

Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. Programs are constructed from *packages*, whose properties allow efficient management of dependencies.

Go 是一种通用语言，设计时考虑到了系统编程。它是强类型和垃圾回收的，并且对并发编程有明确的支持。程序是由包构成的，其属性允许对依赖关系进行有效的管理。

The grammar is compact and simple to parse, allowing for easy analysis by automatic tools such as integrated development environments.

该文法简洁且易于解析，允许使用诸如集成开发环境之类的自动工具进行简单的分析。

## Notation 记号法

The syntax is specified using Extended Backus-Naur Form (EBNF):

语法是使用 Extended Backus-Naur Form (EBNF)指定的:

```
Production  = production_name "=" [ Expression ] "." .
Expression  = Alternative { "|" Alternative } .
Alternative = Term { Term } .
Term        = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

Productions are expressions constructed from terms and the following operators, in increasing precedence:

结果是由术语和以下运算符构造的表达式，优先级依次递增:

```
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

Lower-case production names are used to identify lexical tokens. Non-terminals are in CamelCase. Lexical tokens are enclosed in double quotes `""` or back quotes ````.

小写生产名称用于标识词法标记。非终端是在 CamelCase。词法标记用双引号“”或反引号“括起来。

The form `a … b` represents the set of characters from `a` through `b` as alternatives. The horizontal ellipsis `…` is also used elsewhere in the spec to informally denote various enumerations or code snippets that are not further specified. The character `…` (as opposed to the three characters `...`) is not a token of the Go language.

形式 a... b 表示从 a 到 b 的字符集作为替换。规范中的其他地方也使用了水平省略号... 来非正式地表示未进一步指定的各种枚举或代码段。字符... (相对于三个字符...)不是围棋语言的标志。

## Source code representation 源代码表示

Source code is Unicode text encoded in [UTF-8](https://en.wikipedia.org/wiki/UTF-8). The text is not canonicalized, so a single accented code point is distinct from the same character constructed from combining an accent and a letter; those are treated as two code points. For simplicity, this document will use the unqualified term *character* to refer to a Unicode code point in the source text.

源代码是采用 UTF-8编码的 Unicode 文本。文本没有规范化，因此单个重音代码点不同于由重音符和字母组合构成的相同字符; 这些字符被视为两个代码点。为了简单起见，本文将使用非限定术语字符来指代源文本中的 Unicode字符。

Each code point is distinct; for instance, upper and lower case letters are different characters.

每个编码点都是不同的; 例如，大小写字母是不同的字符。

Implementation restriction: For compatibility with other tools, a compiler may disallow the NUL character (U+0000) in the source text.

实现限制: 为了与其他工具兼容，编译器可能不允许源文本中的 NUL 字符(u + 0000)。

Implementation restriction: For compatibility with other tools, a compiler may ignore a UTF-8-encoded byte order mark (U+FEFF) if it is the first Unicode code point in the source text. A byte order mark may be disallowed anywhere else in the source.

实现限制: 为了与其他工具兼容，编译器可能会忽略 utf-8编码的字节顺序标记(u + feff) ，如果它是源文本中的第一个 Unicode字符。字节顺序标记可能在源代码的其他任何地方被禁用。

### Characters 人物

The following terms are used to denote specific Unicode character classes:

以下术语用于表示特定的 Unicode字符类别:

```
newline        = /* the Unicode code point U+000A */ .
unicode_char   = /* an arbitrary Unicode code point except newline */ .
unicode_letter = /* a Unicode code point classified as "Letter" */ .
unicode_digit  = /* a Unicode code point classified as "Number, decimal digit" */ .
```

In [The Unicode Standard 8.0](https://www.unicode.org/versions/Unicode8.0.0/), Section 4.5 "General Category" defines a set of character categories. Go treats all characters in any of the Letter categories Lu, Ll, Lt, Lm, or Lo as Unicode letters, and those in the Number category Nd as Unicode digits.

在 Unicode 标准8.0中，第4.5节“一般类别”定义了一组字符类别。Go 将任何字母类别中的所有字符都视为 Unicode 字母，而数字类别中的 Nd 则视为 Unicode 数字。

### Letters and digits 字母和数字

The underscore character `_` (U+005F) is considered a letter.

下划线字符 _ (u + 005f)被认为是字母。

```
letter        = unicode_letter | "_" .
decimal_digit = "0" … "9" .
binary_digit  = "0" | "1" .
octal_digit   = "0" … "7" .
hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
```

## Lexical elements 词汇元素

### Comments 评论

Comments serve as program documentation. There are two forms:

注释作为程序文档，有两种形式:

1. *Line comments 行注释* start with the character sequence 从字符序列开始`//` and stop at the end of the line. 然后在终点站停下
2. *General comments 一般评论* start with the character sequence 从字符序列开始`/*` and stop with the first subsequent character sequence 然后停止第一个后续字符序列`*/`.

A comment cannot start inside a [rune](http://docscn.studygolang.com/ref/spec#Rune_literals) or [string literal](http://docscn.studygolang.com/ref/spec#String_literals), or inside a comment. A general comment containing no newlines acts like a space. Any other comment acts like a newline.

注释不能在符文、字符串或注释中启动。不包含换行符的一般性注释就像一个空格。其他任何注释都像换行符一样。

### Tokens 代币

Tokens form the vocabulary of the Go language. There are four classes: *identifiers*, *keywords*, *operators and punctuation*, and *literals*. *White space*, formed from spaces (U+0020), horizontal tabs (U+0009), carriage returns (U+000D), and newlines (U+000A), is ignored except as it separates tokens that would otherwise combine into a single token. Also, a newline or end of file may trigger the insertion of a [semicolon](http://docscn.studygolang.com/ref/spec#Semicolons). While breaking the input into tokens, the next token is the longest sequence of characters that form a valid token.

标记形成了围棋语言的词汇表。有四个类: 标识符、关键字、运算符和标点符号以及文本。由空格(u + 0020)、水平制表符(u + 0009)、回车(u + 000d)和换行符(u + 000a)组成的空白被忽略，除非它分离标记，否则这些标记将合并为单个标记。此外，换行符或文件末尾可能会触发插入分号。在将输入分解为标记时，下一个标记是形成有效标记的最长字符序列。

### Semicolons 分号

The formal grammar uses semicolons `";"` as terminators in a number of productions. Go programs may omit most of these semicolons using the following two rules:

形式语法在许多结果中使用分号“ ;”作为结束符。程序可以用以下两个规则省略大部分分号:

1. When the input is broken into tokens, a semicolon is automatically inserted into the token stream immediately after a line's final token if that token is

    

   当输入被分解为令牌时，如果一行的最终令牌是，则分号会自动插入到令牌流中

   - an 一个[identifier 标识符](http://docscn.studygolang.com/ref/spec#Identifiers)
   - an 一个[integer 整数](http://docscn.studygolang.com/ref/spec#Integer_literals), [floating-point 浮点数](http://docscn.studygolang.com/ref/spec#Floating-point_literals), [imaginary 想象的](http://docscn.studygolang.com/ref/spec#Imaginary_literals), [rune 符文](http://docscn.studygolang.com/ref/spec#Rune_literals), or ，或[string 弦](http://docscn.studygolang.com/ref/spec#String_literals) literal 文字的
   - one of the 其中一个[keywords 关键字](http://docscn.studygolang.com/ref/spec#Keywords) `break`, `continue`, `fallthrough`, or ，或`return`
   - one of the 其中一个[operators and punctuation 操作符和标点符号](http://docscn.studygolang.com/ref/spec#Operators_and_punctuation) `++`, `--`, `)`, `]`, or ，或`}`

2. To allow complex statements to occupy a single line, a semicolon may be omitted before a closing 为了使复杂语句占据一行，结束语前可以省略分号`")"` or 或`"}"`.

To reflect idiomatic use, code examples in this document elide semicolons using these rules.

为了反映惯用用法，本文档中的代码示例使用这些规则省略了分号。

### Identifiers 标识符

Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.

标识符命名程序实体，如变量和类型。标识符是由一个或多个字母和数字组成的序列。标识符的第一个字符必须是字母。

```
identifier = letter { letter | unicode_digit } .
a
_x9
ThisVariableIsExported
αβ
```

Some identifiers are [predeclared](http://docscn.studygolang.com/ref/spec#Predeclared_identifiers).

有些标识符是预声明的。

### Keywords 关键词

The following keywords are reserved and may not be used as identifiers.

下列关键字是保留的，不能用作标识符。

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### Operators and punctuation 操作符和标点符号

The following character sequences represent [operators](http://docscn.studygolang.com/ref/spec#Operators) (including [assignment operators](http://docscn.studygolang.com/ref/spec#assign_op)) and punctuation:

下面的字符序列表示运算符(包括赋值运算符)和标点符号:

```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```

### Integer literals 整数字面值

An integer literal is a sequence of digits representing an [integer constant](http://docscn.studygolang.com/ref/spec#Constants). An optional prefix sets a non-decimal base: `0b` or `0B` for binary, `0`, `0o`, or `0O` for octal, and `0x` or `0X` for hexadecimal. A single `0` is considered a decimal zero. In hexadecimal literals, letters `a` through `f` and `A` through `F` represent values 10 through 15.

整数文字量是表示整数常量的数字序列。一个可选的前缀设置一个非十进制的基数: 0B 或0B 表示二进制，0,0O 或0O 表示八进制，0X 或0X 表示十六进制。一个0被认为是一个十进制零。在十六进制字面值中，字母 a 到 f 和 a 到 f 表示值10到15。

For readability, an underscore character `_` may appear after a base prefix or between successive digits; such underscores do not change the literal's value.

为了便于阅读，下划线字符 _ 可以出现在基前缀之后或连续数字之间; 这样的下划线不会改变字面值。

```
int_lit        = decimal_lit | binary_lit | octal_lit | hex_lit .
decimal_lit    = "0" | ( "1" … "9" ) [ [ "_" ] decimal_digits ] .
binary_lit     = "0" ( "b" | "B" ) [ "_" ] binary_digits .
octal_lit      = "0" [ "o" | "O" ] [ "_" ] octal_digits .
hex_lit        = "0" ( "x" | "X" ) [ "_" ] hex_digits .

decimal_digits = decimal_digit { [ "_" ] decimal_digit } .
binary_digits  = binary_digit { [ "_" ] binary_digit } .
octal_digits   = octal_digit { [ "_" ] octal_digit } .
hex_digits     = hex_digit { [ "_" ] hex_digit } .
42
4_2
0600
0_600
0o600
0O600       // second character is capital letter 'O'
0xBadFace
0xBad_Face
0x_67_7a_2f_cc_40_c6
170141183460469231731687303715884105727
170_141183_460469_231731_687303_715884_105727

_42         // an identifier, not an integer literal
42_         // invalid: _ must separate successive digits
4__2        // invalid: only one _ at a time
0_xBadFace  // invalid: _ must separate successive digits
```

### Floating-point literals 浮点文字

A floating-point literal is a decimal or hexadecimal representation of a [floating-point constant](http://docscn.studygolang.com/ref/spec#Constants).

浮点文字是浮点常数的十进制或十六进制表示形式。

A decimal floating-point literal consists of an integer part (decimal digits), a decimal point, a fractional part (decimal digits), and an exponent part (`e` or `E` followed by an optional sign and decimal digits). One of the integer part or the fractional part may be elided; one of the decimal point or the exponent part may be elided. An exponent value exp scales the mantissa (integer and fractional part) by 10exp.

十进制浮点数字面值由一个整数部分(小数位)、一个小数部分(小数位)和一个指数部分(e 或 e 后面跟随一个可选的符号和小数位)组成。一个整数部分或小数部分可以省略，一个小数点或指数部分可以省略。指数值 exp 将尾数(整数部分和小数部分)缩放10exp。

A hexadecimal floating-point literal consists of a `0x` or `0X` prefix, an integer part (hexadecimal digits), a radix point, a fractional part (hexadecimal digits), and an exponent part (`p` or `P` followed by an optional sign and decimal digits). One of the integer part or the fractional part may be elided; the radix point may be elided as well, but the exponent part is required. (This syntax matches the one given in IEEE 754-2008 §5.12.3.) An exponent value exp scales the mantissa (integer and fractional part) by 2exp.

十六进制浮点数字面值由一个0 x 或0 x 前缀、一个整数部分(十六进制数字)、一个基数部分、一个小数部分(十六进制数字)和一个指数部分(p 或 p 后跟一个可选的符号和十进制数字)组成。可以省略整数部分或小数部分之一; 也可以省略基数点，但需要省略指数部分。(此语法与 IEEE 754-20085.12.3中给出的语法相匹配。)指数值 exp 将尾数(整数部分和小数部分)缩放2 exp。

For readability, an underscore character `_` may appear after a base prefix or between successive digits; such underscores do not change the literal value.

为了便于阅读，下划线字符 _ 可以出现在基前缀之后或连续数字之间; 这样的下划线不会改变原义值。

```
float_lit         = decimal_float_lit | hex_float_lit .

decimal_float_lit = decimal_digits "." [ decimal_digits ] [ decimal_exponent ] |
                    decimal_digits decimal_exponent |
                    "." decimal_digits [ decimal_exponent ] .
decimal_exponent  = ( "e" | "E" ) [ "+" | "-" ] decimal_digits .

hex_float_lit     = "0" ( "x" | "X" ) hex_mantissa hex_exponent .
hex_mantissa      = [ "_" ] hex_digits "." [ hex_digits ] |
                    [ "_" ] hex_digits |
                    "." hex_digits .
hex_exponent      = ( "p" | "P" ) [ "+" | "-" ] decimal_digits .
0.
72.40
072.40       // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
1_5.         // == 15.0
0.15e+0_2    // == 15.0

0x1p-2       // == 0.25
0x2.p10      // == 2048.0
0x1.Fp+0     // == 1.9375
0X.8p-0      // == 0.5
0X_1FFFP-16  // == 0.1249847412109375
0x15e-2      // == 0x15e - 2 (integer subtraction)

0x.p1        // invalid: mantissa has no digits
1p-2         // invalid: p exponent requires hexadecimal mantissa
0x1.5e-2     // invalid: hexadecimal mantissa requires p exponent
1_.5         // invalid: _ must separate successive digits
1._5         // invalid: _ must separate successive digits
1.5_e1       // invalid: _ must separate successive digits
1.5e_1       // invalid: _ must separate successive digits
1.5e1_       // invalid: _ must separate successive digits
```

### Imaginary literals 虚构的文字

An imaginary literal represents the imaginary part of a [complex constant](http://docscn.studygolang.com/ref/spec#Constants). It consists of an [integer](http://docscn.studygolang.com/ref/spec#Integer_literals) or [floating-point](http://docscn.studygolang.com/ref/spec#Floating-point_literals) literal followed by the lower-case letter `i`. The value of an imaginary literal is the value of the respective integer or floating-point literal multiplied by the imaginary unit *i*.

假想的字面量表示复常数的假想部分。它由一个整数或浮点字面值后跟小写字母 i 组成。虚值的值是相应的整数或浮点值乘以虚值单位 i 的值。

```
imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
```

For backward compatibility, an imaginary literal's integer part consisting entirely of decimal digits (and possibly underscores) is considered a decimal integer, even if it starts with a leading `0`.

对于向下兼容来说，一个虚构的整数部分完全由十进制数字(可能还有下划线)组成，即使以0开头，也被认为是一个十进制整数。

```
0i
0123i         // == 123i for backward-compatibility
0o123i        // == 0o123 * 1i == 83i
0xabci        // == 0xabc * 1i == 2748i
0.i
2.71828i
1.e+0i
6.67428e-11i
1E6i
.25i
.12345E+5i
0x1p-2i       // == 0x1p-2 * 1i == 0.25i
```

### Rune literals 符文字面值

A rune literal represents a [rune constant](http://docscn.studygolang.com/ref/spec#Constants), an integer value identifying a Unicode code point. A rune literal is expressed as one or more characters enclosed in single quotes, as in `'x'` or `'\n'`. Within the quotes, any character may appear except newline and unescaped single quote. A single quoted character represents the Unicode value of the character itself, while multi-character sequences beginning with a backslash encode values in various formats.

一个 rune literal 表示一个 rune 常量，一个整数值表示一个 rune Unicode字符。一个 rune literal 是用单引号括起来的一个或多个字符来表示的，比如在‘ x’或者‘ n’中。在引号中，除了换行符和未转义的单引号外，任何字符都可以出现。单引号字符表示字符本身的 Unicode 值，而多字符序列以各种格式的反斜杠编码值开头。

The simplest form represents the single character within the quotes; since Go source text is Unicode characters encoded in UTF-8, multiple UTF-8-encoded bytes may represent a single integer value. For instance, the literal `'a'` holds a single byte representing a literal `a`, Unicode U+0061, value `0x61`, while `'ä'` holds two bytes (`0xc3` `0xa4`) representing a literal `a`-dieresis, U+00E4, value `0xe4`.

最简单的形式表示引号中的单个字符; 因为 Go 源文本是 UTF-8编码的 Unicode 字符，所以多个 UTF-8编码的字节可以表示单个整数值。例如，文本‘ a’包含一个字节，表示文本 a，Unicode + 0061，value 0x61，而‘ ä’包含两个字节(0xc30xa4) ，表示文本 a-dieresis，u + 00 E4，value 0x4。

Several backslash escapes allow arbitrary values to be encoded as ASCII text. There are four ways to represent the integer value as a numeric constant: `\x` followed by exactly two hexadecimal digits; `\u` followed by exactly four hexadecimal digits; `\U` followed by exactly eight hexadecimal digits, and a plain backslash `\` followed by exactly three octal digits. In each case the value of the literal is the value represented by the digits in the corresponding base.

多个反斜杠转义允许将任意值编码为 ASCII 文本。有四种方法可以将整数值表示为一个数字常量: x 后跟正好两个十六进制数字; u 后跟正好四个十六进制数字; u 后跟正好八个十六进制数字，纯斜杠后跟正好三个八进制数字。在每种情况下，文字的值都是由对应基数中的数字表示的值。

Although these representations all result in an integer, they have different valid ranges. Octal escapes must represent a value between 0 and 255 inclusive. Hexadecimal escapes satisfy this condition by construction. The escapes `\u` and `\U` represent Unicode code points so within them some values are illegal, in particular those above `0x10FFFF` and surrogate halves.

尽管这些表示都以整数形式出现，但它们有不同的有效范围。八进制转义必须表示一个介于0和255之间的值。十六进制转义通过构造来满足这一条件。Escapes u 和 u 表示 Unicode 代码点，因此它们中的一些值是非法的，特别是上面的0x10f 和代理 half。

After a backslash, certain single-character escapes represent special values:

在反斜杠之后，某些单字符转义表示特殊的值:

```
\a   U+0007 alert or bell
\b   U+0008 backspace
\f   U+000C form feed
\n   U+000A line feed or newline
\r   U+000D carriage return
\t   U+0009 horizontal tab
\v   U+000b vertical tab
\\   U+005c backslash
\'   U+0027 single quote  (valid escape only within rune literals)
\"   U+0022 double quote  (valid escape only within string literals)
```

All other sequences starting with a backslash are illegal inside rune literals.

所有其他以反斜杠开头的序列在符文文本中都是非法的。

```
rune_lit         = "'" ( unicode_value | byte_value ) "'" .
unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
byte_value       = octal_byte_value | hex_byte_value .
octal_byte_value = `\` octal_digit octal_digit octal_digit .
hex_byte_value   = `\` "x" hex_digit hex_digit .
little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
                           hex_digit hex_digit hex_digit hex_digit .
escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'\''         // rune literal containing single quote character
'aa'         // illegal: too many characters
'\xa'        // illegal: too few hexadecimal digits
'\0'         // illegal: too few octal digits
'\uDFFF'     // illegal: surrogate half
'\U00110000' // illegal: invalid Unicode code point
```

### String literals 字符串文字

A string literal represents a [string constant](http://docscn.studygolang.com/ref/spec#Constants) obtained from concatenating a sequence of characters. There are two forms: raw string literals and interpreted string literals.

字符串字面值表示串联一个字符序列而获得的字符串常量。有两种形式: 原始字符串和解释字符串。

Raw string literals are character sequences between back quotes, as in ``foo``. Within the quotes, any character may appear except back quote. The value of a raw string literal is the string composed of the uninterpreted (implicitly UTF-8-encoded) characters between the quotes; in particular, backslashes have no special meaning and the string may contain newlines. Carriage return characters ('\r') inside raw string literals are discarded from the raw string value.

原始字符串文字是反引号之间的字符序列，如‘ foo’中的字符序列。在引号中，除了反引号外，任何字符都可以出现。原始字符串的值是由引号之间未解释的(隐式 utf-8编码的)字符组成的字符串; 特别是反斜杠没有特殊含义，字符串可能包含换行符。原始字符串中的回车字符(’r’)将从原始字符串值中丢弃。

Interpreted string literals are character sequences between double quotes, as in `"bar"`. Within the quotes, any character may appear except newline and unescaped double quote. The text between the quotes forms the value of the literal, with backslash escapes interpreted as they are in [rune literals](http://docscn.studygolang.com/ref/spec#Rune_literals) (except that `\'` is illegal and `\"` is legal), with the same restrictions. The three-digit octal (`\`*nnn*) and two-digit hexadecimal (`\x`*nn*) escapes represent individual *bytes* of the resulting string; all other escapes represent the (possibly multi-byte) UTF-8 encoding of individual *characters*. Thus inside a string literal `\377` and `\xFF` represent a single byte of value `0xFF`=255, while `ÿ`, `\u00FF`, `\U000000FF` and `\xc3\xbf` represent the two bytes `0xc3` `0xbf` of the UTF-8 encoding of character U+00FF.

解释的字符串文字是双引号之间的字符序列，如“ bar”中所示。在引号中，除了换行符和未转义的双引号之外，任何字符都可以出现。双引号之间的文本构成了文字的值，反斜杠转义解释为它们是符文文本(除非’是非法的和’是合法的) ，具有相同的限制。三位八进制(nnn)和两位十六进制(xnn)转义表示结果字符串的单个字节; 所有其他转义表示单个字符的(可能是多字节的) UTF-8编码。因此，字符串 literal 377和 xFF 表示值0xFF = 255的一个字节，而 ÿ、 u00FF、 U000000FF 和 xc3 xbf 表示字符 u + 00ff 的 UTF-8编码的两个字节0xc30xbf。

```
string_lit             = raw_string_lit | interpreted_string_lit .
raw_string_lit         = "`" { unicode_char | newline } "`" .
interpreted_string_lit = `"` { unicode_value | byte_value } `"` .
`abc`                // same as "abc"
`\n
\n`                  // same as "\\n\n\\n"
"\n"
"\""                 // same as `"`
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"             // illegal: surrogate half
"\U00110000"         // illegal: invalid Unicode code point
```

These examples all represent the same string:

这些例子都代表同一个字符串:

```
"日本語"                                 // UTF-8 input text
`日本語`                                 // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
```

If the source code represents a character as two code points, such as a combining form involving an accent and a letter, the result will be an error if placed in a rune literal (it is not a single code point), and will appear as two code points if placed in a string literal.

如果源代码将一个字符表示为两个代码点，例如一个包含重音和字母的组合表单，那么如果将其放在一个 rune literal 中，结果将是一个错误(它不是单个代码点) ，如果将其放在一个字符串 literal 中，结果将显示为两个代码点。

## Constants 常量

There are *boolean constants*, *rune constants*, *integer constants*, *floating-point constants*, *complex constants*, and *string constants*. Rune, integer, floating-point, and complex constants are collectively called *numeric constants*.

有布尔常量、符文常量、整数常量、浮点常量、复合常量和字符串常量。Rune、 integer、浮点数和复数常量统称为数值型常量。

A constant value is represented by a [rune](http://docscn.studygolang.com/ref/spec#Rune_literals), [integer](http://docscn.studygolang.com/ref/spec#Integer_literals), [floating-point](http://docscn.studygolang.com/ref/spec#Floating-point_literals), [imaginary](http://docscn.studygolang.com/ref/spec#Imaginary_literals), or [string](http://docscn.studygolang.com/ref/spec#String_literals) literal, an identifier denoting a constant, a [constant expression](http://docscn.studygolang.com/ref/spec#Constant_expressions), a [conversion](http://docscn.studygolang.com/ref/spec#Conversions) with a result that is a constant, or the result value of some built-in functions such as `unsafe.Sizeof` applied to any value, `cap` or `len` applied to [some expressions](http://docscn.studygolang.com/ref/spec#Length_and_capacity), `real` and `imag` applied to a complex constant and `complex` applied to numeric constants. The boolean truth values are represented by the predeclared constants `true` and `false`. The predeclared identifier [iota](http://docscn.studygolang.com/ref/spec#Iota) denotes an integer constant.

常量值由符文、整数、浮点数、虚数或字符串文字、表示常量的标识符、常量表达式、结果为常量的转换或某些内置函数(如 unsafe)的结果值表示。Sizeof 应用于任何值，cap 或 len 应用于某些表达式，real 和 imag 应用于复杂的常量和复杂的数字常量。布尔真值由预先声明的常量 true 和 false 表示。预声明的标识符 iota 表示整数常量。

In general, complex constants are a form of [constant expression](http://docscn.studygolang.com/ref/spec#Constant_expressions) and are discussed in that section.

一般来说，复常数是常数表达式的一种形式，本节将对其进行讨论。

Numeric constants represent exact values of arbitrary precision and do not overflow. Consequently, there are no constants denoting the IEEE-754 negative zero, infinity, and not-a-number values.

数值常量表示任意精度的精确值，不会溢出。因此，没有表示 IEEE-754负零、无穷和非数值的常数。

Constants may be [typed](http://docscn.studygolang.com/ref/spec#Types) or *untyped*. Literal constants, `true`, `false`, `iota`, and certain [constant expressions](http://docscn.studygolang.com/ref/spec#Constant_expressions) containing only untyped constant operands are untyped.

常量可以是类型化的，也可以是非类型化的。文字常量、 true、 false、 iota 和某些只包含非类型常量操作数的常量表达式是无类型的。

A constant may be given a type explicitly by a [constant declaration](http://docscn.studygolang.com/ref/spec#Constant_declarations) or [conversion](http://docscn.studygolang.com/ref/spec#Conversions), or implicitly when used in a [variable declaration](http://docscn.studygolang.com/ref/spec#Variable_declarations) or an [assignment](http://docscn.studygolang.com/ref/spec#Assignments) or as an operand in an [expression](http://docscn.studygolang.com/ref/spec#Expressions). It is an error if the constant value cannot be [represented](http://docscn.studygolang.com/ref/spec#Representability) as a value of the respective type.

常量可以通过常量声明或转换显式地给定类型，或者隐式地用于变量声明或赋值，或者用作表达式中的操作数。如果常量值不能表示为相应类型的值，则为错误。

An untyped constant has a *default type* which is the type to which the constant is implicitly converted in contexts where a typed value is required, for instance, in a [short variable declaration](http://docscn.studygolang.com/ref/spec#Short_variable_declarations) such as `i := 0` where there is no explicit type. The default type of an untyped constant is `bool`, `rune`, `int`, `float64`, `complex128` or `string` respectively, depending on whether it is a boolean, rune, integer, floating-point, complex, or string constant.

非类型化常量具有默认类型，该类型是在需要类型化值的上下文中隐式转换该常量的类型，例如，在没有显式类型的简短变量声明中，如 i: = 0。非类型化常量的默认类型分别为 bool、 rune、 int、 float64、 complex128或 string，具体取决于它是布尔型、 rune 型、整数型、浮点型、复数型还是字符串常量。

Implementation restriction: Although numeric constants have arbitrary precision in the language, a compiler may implement them using an internal representation with limited precision. That said, every implementation must:

实现限制: 尽管数值常量在语言中具有任意精度，但编译器可以使用有限精度的内部表示来实现它们。尽管如此，每一项实施都必须:

- Represent integer constants with at least 256 bits. 用至少256位表示整数常量
- Represent floating-point constants, including the parts of a complex constant, with a mantissa of at least 256 bits and a signed binary exponent of at least 16 bits. 表示浮点常数，包括复常数的部分，尾数至少为256位，有符号二进制指数至少为16位
- Give an error if unable to represent an integer constant precisely. 如果无法精确表示整数常数，则给出错误
- Give an error if unable to represent a floating-point or complex constant due to overflow. 如果由于溢出而无法表示浮点或复常数，则给出错误
- Round to the nearest representable constant if unable to represent a floating-point or complex constant due to limits on precision. 如果由于精度限制而无法表示浮点数或复数常数，则舍入到最接近的可表示常数

These requirements apply both to literal constants and to the result of evaluating [constant expressions](http://docscn.studygolang.com/ref/spec#Constant_expressions).

这些要求既适用于文字常量，也适用于计算常量表达式的结果。

## Variables 变量

A variable is a storage location for holding a *value*. The set of permissible values is determined by the variable's *[type](http://docscn.studygolang.com/ref/spec#Types)*.

变量是保存值的存储位置。允许值的集合由变量的类型确定。

A [variable declaration](http://docscn.studygolang.com/ref/spec#Variable_declarations) or, for function parameters and results, the signature of a [function declaration](http://docscn.studygolang.com/ref/spec#Function_declarations) or [function literal](http://docscn.studygolang.com/ref/spec#Function_literals) reserves storage for a named variable. Calling the built-in function [`new`](http://docscn.studygolang.com/ref/spec#Allocation) or taking the address of a [composite literal](http://docscn.studygolang.com/ref/spec#Composite_literals) allocates storage for a variable at run time. Such an anonymous variable is referred to via a (possibly implicit) [pointer indirection](http://docscn.studygolang.com/ref/spec#Address_operators).

变量声明，或者，对于函数参数和结果，函数声明的签名或者匿名函数声明为指定的变量保留存储空间。在运行时调用内置函数 new 或获取复合文本的地址为变量分配存储空间。这样的匿名变量通过一个(可能是隐式的)指针间接引用。

*Structured* variables of [array](http://docscn.studygolang.com/ref/spec#Array_types), [slice](http://docscn.studygolang.com/ref/spec#Slice_types), and [struct](http://docscn.studygolang.com/ref/spec#Struct_types) types have elements and fields that may be [addressed](http://docscn.studygolang.com/ref/spec#Address_operators) individually. Each such element acts like a variable.

数组、片和结构类型的结构化变量具有可单独寻址的元素和字段。每个这样的元素都像一个变量。

The *static type* (or just *type*) of a variable is the type given in its declaration, the type provided in the `new` call or composite literal, or the type of an element of a structured variable. Variables of interface type also have a distinct *dynamic type*, which is the concrete type of the value assigned to the variable at run time (unless the value is the predeclared identifier `nil`, which has no type). The dynamic type may vary during execution but values stored in interface variables are always [assignable](http://docscn.studygolang.com/ref/spec#Assignability) to the static type of the variable.

变量的静态类型(或仅仅类型)是其声明中给出的类型、新调用或复合文本中提供的类型或结构化变量的元素类型。接口类型的变量还有一个不同的动态类型，这是在运行时分配给变量的值的具体类型(除非该值是预声明的标识符 nil，它没有类型)。在执行过程中，动态类型可能会发生变化，但是存储在接口变量中的值总是可以分配给变量的静态类型。

```
var x interface{}  // x is nil and has static type interface{}
var v *T           // v has value nil, static type *T
x = 42             // x has value 42 and dynamic type int
x = v              // x has value (*T)(nil) and dynamic type *T
```

A variable's value is retrieved by referring to the variable in an [expression](http://docscn.studygolang.com/ref/spec#Expressions); it is the most recent value [assigned](http://docscn.studygolang.com/ref/spec#Assignments) to the variable. If a variable has not yet been assigned a value, its value is the [zero value](http://docscn.studygolang.com/ref/spec#The_zero_value) for its type.

通过引用表达式中的变量来检索变量的值; 它是分配给该变量的最新值。如果一个变量还没有被赋值，那么它的值就是它的类型的零值。

## Types 类型

A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a *type name*, if it has one, or specified using a *type literal*, which composes a type from existing types.

类型确定一组值以及特定于这些值的操作和方法。类型可以用类型名称表示(如果它有类型名称的话) ，或者使用类型文字(由现有类型组成的类型)指定。

```
Type      = TypeName | TypeLit | "(" Type ")" .
TypeName  = identifier | QualifiedIdent .
TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType |
	    SliceType | MapType | ChannelType .
```

The language [predeclares](http://docscn.studygolang.com/ref/spec#Predeclared_identifiers) certain type names. Others are introduced with [type declarations](http://docscn.studygolang.com/ref/spec#Type_declarations). *Composite types*—array, struct, pointer, function, interface, slice, map, and channel types—may be constructed using type literals.

该语言预先声明了某些类型名称。另外一些是通过类型声明引入的。可以使用类型文本构造复合类型ー array、 struct、 pointer、 function、 interface、 slice、 map 和 channel 类型ー。

Each type `T` has an *underlying type*: If `T` is one of the predeclared boolean, numeric, or string types, or a type literal, the corresponding underlying type is `T` itself. Otherwise, `T`'s underlying type is the underlying type of the type to which `T` refers in its [type declaration](http://docscn.studygolang.com/ref/spec#Type_declarations).

每个类型 t 都有一个基础类型: 如果 t 是预先声明的布尔类型、数值类型或字符串类型或类型文字之一，那么对应的基础类型就是 t 本身。否则，t 的基础类型是 t 在其类型声明中引用的类型的基础类型。

```
type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)
```

The underlying type of `string`, `A1`, `A2`, `B1`, and `B2` is `string`. The underlying type of `[]B1`, `B3`, and `B4` is `[]B1`.

字符串的基本类型是字符串，A1、 A2、 B1和 B2。[] B1，B3和 B4的基本类型是[] B1。

### Method sets 方法集

A type may have a *method set* associated with it. The method set of an [interface type](http://docscn.studygolang.com/ref/spec#Interface_types) is its interface. The method set of any other type `T` consists of all [methods](http://docscn.studygolang.com/ref/spec#Method_declarations) declared with receiver type `T`. The method set of the corresponding [pointer type](http://docscn.studygolang.com/ref/spec#Pointer_types) `*T` is the set of all methods declared with receiver `*T` or `T` (that is, it also contains the method set of `T`). Further rules apply to structs containing embedded fields, as described in the section on [struct types](http://docscn.studygolang.com/ref/spec#Struct_types). Any other type has an empty method set. In a method set, each method must have a [unique](http://docscn.studygolang.com/ref/spec#Uniqueness_of_identifiers) non-[blank](http://docscn.studygolang.com/ref/spec#Blank_identifier) [method name](http://docscn.studygolang.com/ref/spec#MethodName).

类型可以有一个与之关联的方法集。接口类型的方法集是它的接口。任何其他类型 t 的方法集都包含用接收器类型 t 声明的所有方法。相应的指针类型 * t 的方法集是用接收器 * t 或 t 声明的所有方法的集合(也就是说，它还包含 t 的方法集)。进一步的规则适用于包含嵌入字段的结构，如关于结构类型的部分所述。任何其他类型都有一个空方法集。在方法集中，每个方法必须有唯一的非空方法名称。

The method set of a type determines the interfaces that the type [implements](http://docscn.studygolang.com/ref/spec#Interface_types) and the methods that can be [called](http://docscn.studygolang.com/ref/spec#Calls) using a receiver of that type.

类型的方法集确定该类型实现的接口以及可以使用该类型的接收方调用的方法。

### Boolean types 布尔类型

A *boolean type* represents the set of Boolean truth values denoted by the predeclared constants `true` and `false`. The predeclared boolean type is `bool`; it is a [defined type](http://docscn.studygolang.com/ref/spec#Type_definitions).

布尔类型表示由预先声明的常量 true 和 false 表示的布尔真值集。预先声明的布尔类型是 bool; 它是已定义的类型。

### Numeric types 数字类型

A *numeric type* represents sets of integer or floating-point values. The predeclared architecture-independent numeric types are:

数值类型表示整数或浮点值的集合。预声明的与体系结构无关的数值类型是:

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

The value of an *n*-bit integer is *n* bits wide and represented using [two's complement arithmetic](https://en.wikipedia.org/wiki/Two's_complement).

一个 n 位整数的值是 n 位宽的，并且使用两个补数运算来表示。

There is also a set of predeclared numeric types with implementation-specific sizes:

还有一组预先声明的具有特定实现大小的数值类型:

```
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

To avoid portability issues all numeric types are [defined types](http://docscn.studygolang.com/ref/spec#Type_definitions) and thus distinct except `byte`, which is an [alias](http://docscn.studygolang.com/ref/spec#Alias_declarations) for `uint8`, and `rune`, which is an alias for `int32`. Explicit conversions are required when different numeric types are mixed in an expression or assignment. For instance, `int32` and `int` are not the same type even though they may have the same size on a particular architecture.

为了避免可移植性问题，所有数值类型都是已定义类型，因此除了 byte 和 rune，前者是 uint8的别名，后者是 int32的别名。当表达式或赋值中混合了不同的数值类型时，需要进行显式转换。例如，int32和 int 不是相同的类型，即使它们在特定的体系结构上具有相同的大小。

### String types 字符串类型

A *string type* represents the set of string values. A string value is a (possibly empty) sequence of bytes. The number of bytes is called the length of the string and is never negative. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is `string`; it is a [defined type](http://docscn.studygolang.com/ref/spec#Type_definitions).

字符串类型表示字符串值的集合。字符串值是一个(可能是空的)字节序列。字节数称为字符串的长度，从来不是负数。字符串是不可变的: 一旦创建，就不可能改变字符串的内容。预声明的字符串类型是 string; 它是已定义的类型。

The length of a string `s` can be discovered using the built-in function [`len`](http://docscn.studygolang.com/ref/spec#Length_and_capacity). The length is a compile-time constant if the string is a constant. A string's bytes can be accessed by integer [indices](http://docscn.studygolang.com/ref/spec#Index_expressions) 0 through `len(s)-1`. It is illegal to take the address of such an element; if `s[i]` is the `i`'th byte of a string, `&s[i]` is invalid.

字符串 s 的长度可以使用内置函数 len 来发现。如果字符串是常量，则长度为编译时常量。整数索引0到 len (s)-1可以访问字符串的字节。获取这种元素的地址是非法的; 如果 s [ i ]是字符串的 i‘ th 字节，& s [ i ]是无效的。

### Array types 数组类型

An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the array and is never negative.

数组是单一类型的元素的编号序列，称为元素类型。元素的个数称为数组的长度，永远不会为负数。

```
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
```

The length is part of the array's type; it must evaluate to a non-negative [constant](http://docscn.studygolang.com/ref/spec#Constants) [representable](http://docscn.studygolang.com/ref/spec#Representability) by a value of type `int`. The length of array `a` can be discovered using the built-in function [`len`](http://docscn.studygolang.com/ref/spec#Length_and_capacity). The elements can be addressed by integer [indices](http://docscn.studygolang.com/ref/spec#Index_expressions) 0 through `len(a)-1`. Array types are always one-dimensional but may be composed to form multi-dimensional types.

长度是数组类型的一部分; 它必须计算为一个非负常数，该常数可由 int 类型的值表示。数组 a 的长度可以通过使用内置函数 len 来发现。元素可以通过整数索引0到 len (a)-1来寻址。数组类型始终是一维的，但可以组合成多维类型。

```
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

### Slice types 切片类型

A slice is a descriptor for a contiguous segment of an *underlying array* and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The number of elements is called the length of the slice and is never negative. The value of an uninitialized slice is `nil`.

片是基础数组的连续段的描述符，并提供对该数组中元素序列的编号访问。片类型表示其元素类型的所有数组片的集合。元素的数量称为切片的长度，从来不是负数。未初始化的片的值为 nil。

```
SliceType = "[" "]" ElementType .
```

The length of a slice `s` can be discovered by the built-in function [`len`](http://docscn.studygolang.com/ref/spec#Length_and_capacity); unlike with arrays it may change during execution. The elements can be addressed by integer [indices](http://docscn.studygolang.com/ref/spec#Index_expressions) 0 through `len(s)-1`. The slice index of a given element may be less than the index of the same element in the underlying array.

片的长度可以通过内置的函数 len 来发现; 与数组不同，它可以在执行过程中改变。元素可以通过整数索引0到 len (s)-1来寻址。给定元素的片索引可以小于基础数组中同一元素的索引。

A slice, once initialized, is always associated with an underlying array that holds its elements. A slice therefore shares storage with its array and with other slices of the same array; by contrast, distinct arrays always represent distinct storage.

一个片一旦初始化，总是与包含其元素的基础数组关联。因此，切片与其数组以及同一数组的其他切片共享存储; 相比之下，不同的数组始终表示不同的存储。

The array underlying a slice may extend past the end of the slice. The *capacity* is a measure of that extent: it is the sum of the length of the slice and the length of the array beyond the slice; a slice of length up to that capacity can be created by [*slicing*](http://docscn.studygolang.com/ref/spec#Slice_expressions) a new one from the original slice. The capacity of a slice `a` can be discovered using the built-in function [`cap(a)`](http://docscn.studygolang.com/ref/spec#Length_and_capacity).

切片的基础数组可以延伸到切片的末端。容量就是这种程度的度量: 它是切片长度和切片之外的数组长度的总和; 通过从原始切片切片创建一个新的切片，可以创建一个达到这个容量的切片。可以使用内置的功能盖(a)发现片 a 的容量。

A new, initialized slice value for a given element type `T` is made using the built-in function [`make`](http://docscn.studygolang.com/ref/spec#Making_slices_maps_and_channels), which takes a slice type and parameters specifying the length and optionally the capacity. A slice created with `make` always allocates a new, hidden array to which the returned slice value refers. That is, executing

使用内置函数 make 为给定元素类型 t 创建一个新的初始化切片值，该值采用切片类型和参数，指定长度和可选的容量。使用 make 创建的片总是分配返回的片值所引用的新的隐藏数组。也就是说，执行

```
make([]T, length, capacity)
```

produces the same slice as allocating an array and [slicing](http://docscn.studygolang.com/ref/spec#Slice_expressions) it, so these two expressions are equivalent:

生成与分配数组和切片相同的切片，所以这两个表达式是等价的:

```
make([]int, 50, 100)
new([100]int)[0:50]
```

Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. With arrays of arrays, the inner arrays are, by construction, always the same length; however with slices of slices (or arrays of slices), the inner lengths may vary dynamically. Moreover, the inner slices must be initialized individually.

像数组一样，切片始终是一维的，但可以组合为构造更高维的对象。对于数组数组，根据结构，内部数组总是相同的长度; 但是对于切片(或切片数组) ，内部长度可以动态变化。此外，内部片必须单独初始化。

### Struct types 结构类型

A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-[blank](http://docscn.studygolang.com/ref/spec#Blank_identifier) field names must be [unique](http://docscn.studygolang.com/ref/spec#Uniqueness_of_identifiers).

Struct 是一个命名元素序列，称为字段，每个字段都有一个名称和一个类型。字段名可以显式指定(IdentifierList)或隐式指定(EmbeddedField)。在结构中，非空字段名称必须是唯一的。

```
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName .
Tag           = string_lit .
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
```

A field declared with a type but no explicit field name is called an *embedded field*. An embedded field must be specified as a type name `T` or as a pointer to a non-interface type name `*T`, and `T` itself may not be a pointer type. The unqualified type name acts as the field name.

用类型声明但没有显式字段名称的字段称为嵌入字段。嵌入字段必须指定为类型名 t 或指向非接口类型名 * t 的指针，而且 t 本身可能不是指针类型。未限定的类型名充当字段名。

```
// A struct with four embedded fields of types T1, *T2, P.T3 and *P.T4
struct {
	T1        // field name is T1
	*T2       // field name is T2
	P.T3      // field name is T3
	*P.T4     // field name is T4
	x, y int  // field names are x and y
}
```

The following declaration is illegal because field names must be unique in a struct type:

下面的声明是非法的，因为字段名在结构类型中必须是唯一的:

```
struct {
	T     // conflicts with embedded field *T and *P.T
	*T    // conflicts with embedded field T and *P.T
	*P.T  // conflicts with embedded field T and *T
}
```

A field or [method](http://docscn.studygolang.com/ref/spec#Method_declarations) `f` of an embedded field in a struct `x` is called *promoted* if `x.f` is a legal [selector](http://docscn.studygolang.com/ref/spec#Selectors) that denotes that field or method `f`.

如果 x.f 是表示字段或方法 f 的合法选择器，则调用 structx 中嵌入字段的字段或方法 f。

Promoted fields act like ordinary fields of a struct except that they cannot be used as field names in [composite literals](http://docscn.studygolang.com/ref/spec#Composite_literals) of the struct.

提升字段与结构的普通字段相似，只是不能用作结构的复合文本中的字段名。

Given a struct type `S` and a [defined type](http://docscn.studygolang.com/ref/spec#Type_definitions) `T`, promoted methods are included in the method set of the struct as follows:

给定结构类型 s 和定义类型 t，提升方法包含在结构的方法集中，如下所示:

- If 如果`S` contains an embedded field 包含一个嵌入字段`T`, the 、预防退伍军人病委员会[method sets 方法集](http://docscn.studygolang.com/ref/spec#Method_sets) of 的`S` and 及`*S` both include promoted methods with receiver 两者都包括接收器的升级方法`T`. The method set of 。方法集`*S` also includes promoted methods with receiver 还包括用接收器提升的方法`*T`.
- If 如果`S` contains an embedded field 包含一个嵌入字段`*T`, the method sets of ，方法集的`S` and 及`*S` both include promoted methods with receiver 两者都包括接收器的升级方法`T` or 或`*T`.

A field declaration may be followed by an optional string literal *tag*, which becomes an attribute for all the fields in the corresponding field declaration. An empty tag string is equivalent to an absent tag. The tags are made visible through a [reflection interface](http://docscn.studygolang.com/pkg/reflect/#StructTag) and take part in [type identity](http://docscn.studygolang.com/ref/spec#Type_identity) for structs but are otherwise ignored.

字段声明之后可以跟随一个可选的字符串文本标记，该标记将成为相应字段声明中所有字段的属性。一个空的标记字符串等价于一个缺失的标记。标记通过反射接口可见，并参与结构的类型标识，但在其他情况下被忽略。

```
struct {
	x, y float64 ""  // an empty tag string is like an absent tag
	name string  "any string is permitted as a tag"
	_    [4]byte "ceci n'est pas un champ de structure"
}

// A struct corresponding to a TimeStamp protocol buffer.
// The tag strings define the protocol buffer field numbers;
// they follow the convention outlined by the reflect package.
struct {
	microsec  uint64 `protobuf:"1"`
	serverIP6 uint64 `protobuf:"2"`
}
```

### Pointer types 指针类型

A pointer type denotes the set of all pointers to [variables](http://docscn.studygolang.com/ref/spec#Variables) of a given type, called the *base type* of the pointer. The value of an uninitialized pointer is `nil`.

指针类型表示指向给定类型变量(称为指针的基类型)的所有指针的集合。未初始化的指针的值为 nil。

```
PointerType = "*" BaseType .
BaseType    = Type .
*Point
*[4]int
```

### Function types 功能类型

A function type denotes the set of all functions with the same parameter and result types. The value of an uninitialized variable of function type is `nil`.

函数类型表示具有相同参数和结果类型的所有函数的集合。函数类型的未初始化变量的值为 nil。

```
FunctionType   = "func" Signature .
Signature      = Parameters [ Result ] .
Result         = Parameters | Type .
Parameters     = "(" [ ParameterList [ "," ] ] ")" .
ParameterList  = ParameterDecl { "," ParameterDecl } .
ParameterDecl  = [ IdentifierList ] [ "..." ] Type .
```

Within a list of parameters or results, the names (IdentifierList) must either all be present or all be absent. If present, each name stands for one item (parameter or result) of the specified type and all non-[blank](http://docscn.studygolang.com/ref/spec#Blank_identifier) names in the signature must be [unique](http://docscn.studygolang.com/ref/spec#Uniqueness_of_identifiers). If absent, each type stands for one item of that type. Parameter and result lists are always parenthesized except that if there is exactly one unnamed result it may be written as an unparenthesized type.

在参数或结果列表中，名称(IdentifierList)必须全部出现或全部缺失。如果存在，则每个名称代表指定类型的一个项(参数或结果) ，签名中的所有非空名称都必须是唯一的。如果没有，每个类型表示该类型的一个项。参数和结果列表总是用圆括号括起来，除非只有一个未命名的结果，否则可以将其写为未用圆括号括起来的类型。

The final incoming parameter in a function signature may have a type prefixed with `...`. A function with such a parameter is called *variadic* and may be invoked with zero or more arguments for that parameter.

函数签名中的最终传入参数可以有一个前缀为... 的类型。具有这样一个参数的函数称为可变参数，并且可以使用该参数的零个或多个参数进行调用。

```
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```

### Interface types 接口类型

An interface type specifies a [method set](http://docscn.studygolang.com/ref/spec#Method_sets) called its *interface*. A variable of interface type can store a value of any type with a method set that is any superset of the interface. Such a type is said to *implement the interface*. The value of an uninitialized variable of interface type is `nil`.

接口类型指定称为其接口的方法集。接口类型的变量可以用方法集存储任何类型的值，该方法集是接口的任何超集。这种类型被称为实现接口。接口类型的未初始化变量的值为 nil。

```
InterfaceType      = "interface" "{" { ( MethodSpec | InterfaceTypeName ) ";" } "}" .
MethodSpec         = MethodName Signature .
MethodName         = identifier .
InterfaceTypeName  = TypeName .
```

An interface type may specify methods *explicitly* through method specifications, or it may *embed* methods of other interfaces through interface type names.

接口类型可以通过方法规范显式地指定方法，也可以通过接口类型名称嵌入其他接口的方法。

```
// A simple File interface.
interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}
```

The name of each explicitly specified method must be [unique](http://docscn.studygolang.com/ref/spec#Uniqueness_of_identifiers) and not [blank](http://docscn.studygolang.com/ref/spec#Blank_identifier).

每个显式指定的方法的名称必须是唯一的，而不是空白的。

```
interface {
	String() string
	String() string  // illegal: String not unique
	_(x int)         // illegal: method must have non-blank name
}
```

More than one type may implement an interface. For instance, if two types `S1` and `S2` have the method set

多个类型可以实现一个接口。例如，如果两个类型 S1和 S2已经设置了该方法

```
func (p T) Read(p []byte) (n int, err error)
func (p T) Write(p []byte) (n int, err error)
func (p T) Close() error
```

(where `T` stands for either `S1` or `S2`) then the `File` interface is implemented by both `S1` and `S2`, regardless of what other methods `S1` and `S2` may have or share.

(t 代表 S1或 S2)那么文件接口是由 S1和 S2实现的，而不管 S1和 S2可能拥有或共享什么其他方法。

A type implements any interface comprising any subset of its methods and may therefore implement several distinct interfaces. For instance, all types implement the *empty interface*:

类型实现包含其方法的任何子集的任何接口，因此可以实现几个不同的接口。例如，所有类型都实现空接口:

```
interface{}
```

Similarly, consider this interface specification, which appears within a [type declaration](http://docscn.studygolang.com/ref/spec#Type_declarations) to define an interface called `Locker`:

类似地，考虑一下这个接口规范，它出现在一个类型声明中来定义一个名为 Locker 的接口:

```
type Locker interface {
	Lock()
	Unlock()
}
```

If `S1` and `S2` also implement

如中一及中二同时实施

```
func (p T) Lock() { … }
func (p T) Unlock() { … }
```

they implement the `Locker` interface as well as the `File` interface.

它们实现了 Locker 接口和 File 接口。

An interface `T` may use a (possibly qualified) interface type name `E` in place of a method specification. This is called *embedding* interface `E` in `T`. The [method set](http://docscn.studygolang.com/ref/spec#Method_sets) of `T` is the *union* of the method sets of `T`’s explicitly declared methods and of `T`’s embedded interfaces.

接口 t 可以使用(可能限定的)接口类型名 e 来代替方法规范。这就是所谓的嵌入界面 e 在 t。T 的方法集是 t 的显式声明方法集和 t 的嵌入接口的方法集的并集。

```
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter's methods are Read, Write, and Close.
type ReadWriter interface {
	Reader  // includes methods of Reader in ReadWriter's method set
	Writer  // includes methods of Writer in ReadWriter's method set
}
```

A *union* of method sets contains the (exported and non-exported) methods of each method set exactly once, and methods with the [same](http://docscn.studygolang.com/ref/spec#Uniqueness_of_identifiers) names must have [identical](http://docscn.studygolang.com/ref/spec#Type_identity) signatures.

方法集的并集包含每个方法集的(导出和非导出)方法，并且具有相同名称的方法必须具有相同的签名。

```
type ReadCloser interface {
	Reader   // includes methods of Reader in ReadCloser's method set
	Close()  // illegal: signatures of Reader.Close and Close are different
}
```

An interface type `T` may not embed itself or any interface type that embeds `T`, recursively.

接口类型 t 不能嵌入自身或者递归嵌入 t 的任何接口类型。

```
// illegal: Bad cannot embed itself
type Bad interface {
	Bad
}

// illegal: Bad1 cannot embed itself using Bad2
type Bad1 interface {
	Bad2
}
type Bad2 interface {
	Bad1
}
```

### Map types 地图类型

A map is an unordered group of elements of one type, called the element type, indexed by a set of unique *keys* of another type, called the key type. The value of an uninitialized map is `nil`.

Map 是一种类型的无序元素组，称为元素类型，由另一种类型的一组唯一键(称为键类型)进行索引。未初始化映射的值为 nil。

```
MapType     = "map" "[" KeyType "]" ElementType .
KeyType     = Type .
```

The [comparison operators](http://docscn.studygolang.com/ref/spec#Comparison_operators) `==` and `!=` must be fully defined for operands of the key type; thus the key type must not be a function, map, or slice. If the key type is an interface type, these comparison operators must be defined for the dynamic key values; failure will cause a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics).

比较运算符 = = 和！必须为键类型的操作数完全定义; 因此键类型不能是函数、映射或片。如果键类型是接口类型，则必须为动态键值定义这些比较运算符; 失败将导致运行时恐慌。

```
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
```

The number of map elements is called its length. For a map `m`, it can be discovered using the built-in function [`len`](http://docscn.studygolang.com/ref/spec#Length_and_capacity) and may change during execution. Elements may be added during execution using [assignments](http://docscn.studygolang.com/ref/spec#Assignments) and retrieved with [index expressions](http://docscn.studygolang.com/ref/spec#Index_expressions); they may be removed with the [`delete`](http://docscn.studygolang.com/ref/spec#Deletion_of_map_elements) built-in function.

映射元素的个数称为它的长度。对于 map m，可以使用内置函数 len 发现它，并且在执行过程中可能会发生更改。可以在执行期间使用赋值添加元素，并使用索引表达式检索元素; 可以使用 delete 内置函数删除元素。

A new, empty map value is made using the built-in function [`make`](http://docscn.studygolang.com/ref/spec#Making_slices_maps_and_channels), which takes the map type and an optional capacity hint as arguments:

使用内置函数 make 生成一个新的空 map 值，它将 map 类型和一个可选的容量提示作为参数:

```
make(map[string]int)
make(map[string]int, 100)
```

The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them, with the exception of `nil` maps. A `nil` map is equivalent to an empty map except that no elements may be added.

初始容量没有限制其大小: 映射增长以容纳存储在其中的项目数，空映射除外。Nil 映射相当于空映射，只是不能添加任何元素。

### Channel types 通道类型

A channel provides a mechanism for [concurrently executing functions](http://docscn.studygolang.com/ref/spec#Go_statements) to communicate by [sending](http://docscn.studygolang.com/ref/spec#Send_statements) and [receiving](http://docscn.studygolang.com/ref/spec#Receive_operator) values of a specified element type. The value of an uninitialized channel is `nil`.

通道提供了一种机制，用于同时执行的函数通过发送和接收指定元素类型的值进行通信。未初始化通道的值为 nil。

```
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .
```

The optional `<-` operator specifies the channel *direction*, *send* or *receive*. If no direction is given, the channel is *bidirectional*. A channel may be constrained only to send or only to receive by [assignment](http://docscn.studygolang.com/ref/spec#Assignments) or explicit [conversion](http://docscn.studygolang.com/ref/spec#Conversions).

可选的 <-operator 指定通道方向，发送或接收。如果没有给出方向，则通道是双向的。通道可能仅限于通过赋值或显式转换发送或接收。

```
chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints
```

The `<-` operator associates with the leftmost `chan` possible:

<-operator 与最左边的 chan 关联:

```
chan<- chan int    // same as chan<- (chan int)
chan<- <-chan int  // same as chan<- (<-chan int)
<-chan <-chan int  // same as <-chan (<-chan int)
chan (<-chan int)
```

A new, initialized channel value can be made using the built-in function [`make`](http://docscn.studygolang.com/ref/spec#Making_slices_maps_and_channels), which takes the channel type and an optional *capacity* as arguments:

可以使用内置函数 make 创建一个新的初始化通道值，该函数将通道类型和可选容量作为参数:

```
make(chan int, 100)
```

The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready. Otherwise, the channel is buffered and communication succeeds without blocking if the buffer is not full (sends) or not empty (receives). A `nil` channel is never ready for communication.

容量(以元素数目表示)设置通道中缓冲区的大小。如果容量为零或不存在，则信道不缓冲，通信只有在发送方和接收方都准备好时才能成功。否则，如果缓冲区没有满(发送)或者没有空(接收) ，通道将被缓冲并且通信成功而不会阻塞。无通道永远不会准备好进行通信。

A channel may be closed with the built-in function [`close`](http://docscn.studygolang.com/ref/spec#Close). The multi-valued assignment form of the [receive operator](http://docscn.studygolang.com/ref/spec#Receive_operator) reports whether a received value was sent before the channel was closed.

通道可以关闭，内置函数关闭。接收运算符的多值赋值表单报告接收值是否在通道关闭之前发送。

A single channel may be used in [send statements](http://docscn.studygolang.com/ref/spec#Send_statements), [receive operations](http://docscn.studygolang.com/ref/spec#Receive_operator), and calls to the built-in functions [`cap`](http://docscn.studygolang.com/ref/spec#Length_and_capacity) and [`len`](http://docscn.studygolang.com/ref/spec#Length_and_capacity) by any number of goroutines without further synchronization. Channels act as first-in-first-out queues. For example, if one goroutine sends values on a channel and a second goroutine receives them, the values are received in the order sent.

单个通道可以用于发送语句、接收操作以及任意数量的 goroutine 对内置函数的调用，而无需进一步同步。通道充当先进先出队列。例如，如果一个 goroutine 在通道上发送值，而另一个 goroutine 接收值，则按照发送的顺序接收值。

## Properties of types and values 类型和值的属性

### Type identity 类型标识

Two types are either *identical* or *different*.

两种类型要么相同，要么不同。

A [defined type](http://docscn.studygolang.com/ref/spec#Type_definitions) is always different from any other type. Otherwise, two types are identical if their [underlying](http://docscn.studygolang.com/ref/spec#Types) type literals are structurally equivalent; that is, they have the same literal structure and corresponding components have identical types. In detail:

已定义的类型始终不同于任何其他类型。否则，如果两个类型的基础类型文字在结构上是等价的，那么它们是相同的; 也就是说，它们具有相同的文字结构，而相应的组件具有相同的类型。详细内容:

- Two array types are identical if they have identical element types and the same array length. 如果两个数组具有相同的元素类型和相同的数组长度，那么它们是相同的
- Two slice types are identical if they have identical element types. 如果两个切片类型具有相同的元素类型，那么它们是相同的
- Two struct types are identical if they have the same sequence of fields, and if corresponding fields have the same names, and identical types, and identical tags. 如果两个结构类型具有相同的字段序列，并且相应的字段具有相同的名称、相同的类型和相同的标记，则它们是相同的[Non-exported 非输出](http://docscn.studygolang.com/ref/spec#Exported_identifiers) field names from different packages are always different. 不同软件包的字段名总是不同的
- Two pointer types are identical if they have identical base types. 如果两个指针类型具有相同的基类型，则它们是相同的
- Two function types are identical if they have the same number of parameters and result values, corresponding parameter and result types are identical, and either both functions are variadic or neither is. Parameter and result names are not required to match. 如果两个函数类型具有相同数目的参数和结果值，那么它们是相同的，相应的参数和结果类型是相同的，并且两个函数要么是可变参数，要么都不是。参数和结果名称不需要匹配
- Two interface types are identical if they have the same set of methods with the same names and identical function types. 如果两个接口类型具有相同的方法集，且具有相同的名称和相同的函数类型，则它们是相同的[Non-exported 非输出](http://docscn.studygolang.com/ref/spec#Exported_identifiers) method names from different packages are always different. The order of the methods is irrelevant. 来自不同包的方法名称总是不同的。方法的顺序是无关的
- Two map types are identical if they have identical key and element types. 如果两个映射类型具有相同的键和元素类型，那么它们是相同的
- Two channel types are identical if they have identical element types and the same direction. 如果两个通道具有相同的元素类型和相同的方向，则它们是相同的

Given the declarations

考虑到这些声明

```
type (
	A0 = []string
	A1 = A0
	A2 = struct{ a, b int }
	A3 = int
	A4 = func(A3, float64) *A0
	A5 = func(x int, _ float64) *[]string
)

type (
	B0 A0
	B1 []string
	B2 struct{ a, b int }
	B3 struct{ a, c int }
	B4 func(int, float64) *B0
	B5 func(x int, y float64) *A1
)

type	C0 = B0
```

these types are identical:

这些类型是一样的:

```
A0, A1, and []string
A2 and struct{ a, b int }
A3 and int
A4, func(int, float64) *[]string, and A5

B0 and C0
[]int and []int
struct{ a, b *T5 } and struct{ a, b *T5 }
func(x int, y float64) *[]string, func(int, float64) (result *[]string), and A5
```

`B0` and `B1` are different because they are new types created by distinct [type definitions](http://docscn.studygolang.com/ref/spec#Type_definitions); `func(int, float64) *B0` and `func(x int, y float64) *[]string` are different because `B0` is different from `[]string`.

B0和 B1是不同的，因为它们是由不同的类型定义创建的新类型; func (int，float64) * B0和 func (x int，y float64) * [] string 是不同的，因为 B0不同于[] string。

### Assignability 可转让性

A value `x` is *assignable* to a [variable](http://docscn.studygolang.com/ref/spec#Variables) of type `T` ("`x` is assignable to `T`") if one of the following conditions applies:

如果下列条件之一适用，则值 x 可分配给类型为 t 的变量(“ x 可分配给 t”) :

- `x`'s type is identical to 的类型与`T`.
- `x`'s type 的类型`V` and 及`T` have identical 有相同的[underlying types 基本类型](http://docscn.studygolang.com/ref/spec#Types) and at least one of 至少有一个`V` or 或`T` is not a 不是一个[defined 定义](http://docscn.studygolang.com/ref/spec#Type_definitions) type. 类型
- `T` is an interface type and 是一种接口类型`x` [implements 工具](http://docscn.studygolang.com/ref/spec#Interface_types) `T`.
- `x` is a bidirectional channel value, 是一个双向信道值,`T` is a channel type, 是一种通道类型,`x`'s type 的类型`V` and 及`T` have identical element types, and at least one of 有相同的元素类型，并且至少有一个`V` or 或`T` is not a defined type. 不是定义类型
- `x` is the predeclared identifier 是预先声明的标识符`nil` and 及`T` is a pointer, function, slice, map, channel, or interface type. 是指针、函数、片段、映射、通道或接口类型
- `x` is an untyped 是一个未打字的[constant 常数](http://docscn.studygolang.com/ref/spec#Constants) [representable 可代表的](http://docscn.studygolang.com/ref/spec#Representability) by a value of type 类型的值`T`.

### Representability 具有代表性

A [constant](http://docscn.studygolang.com/ref/spec#Constants) `x` is *representable* by a value of type `T` if one of the following conditions applies:

如果符合下列条件之一，常数 x 可以用类型 t 的值表示:

- `x` is in the set of values 是在一系列的值中[determined 坚定的](http://docscn.studygolang.com/ref/spec#Types) by 作者`T`.
- `T` is a floating-point type and 是浮点类型`x` can be rounded to 可舍入至`T`'s precision without overflow. Rounding uses IEEE 754 round-to-even rules but with an IEEE negative zero further simplified to an unsigned zero. Note that constant values never result in an IEEE negative zero, NaN, or infinity. 没有溢出的精度。四舍五入使用 IEEE 754四舍五入规则，但使用 IEEE 负零进一步简化为无符号零。请注意，常量值不会导致 IEEE 负零、 NaN 或无穷大
- `T` is a complex type, and 是一种复杂的类型`x`'s 是的[components 组件](http://docscn.studygolang.com/ref/spec#Complex_numbers) `real(x)` and 及`imag(x)` are representable by values of 可以通过值来表示`T`'s component type ( 的组件类型(`float32` or 或`float64`).

```
x                   T           x is representable by a value of T because

'a'                 byte        97 is in the set of byte values
97                  rune        rune is an alias for int32, and 97 is in the set of 32-bit integers
"foo"               string      "foo" is in the set of string values
1024                int16       1024 is in the set of 16-bit integers
42.0                byte        42 is in the set of unsigned 8-bit integers
1e10                uint64      10000000000 is in the set of unsigned 64-bit integers
2.718281828459045   float32     2.718281828459045 rounds to 2.7182817 which is in the set of float32 values
-1e-1000            float64     -1e-1000 rounds to IEEE -0.0 which is further simplified to 0.0
0i                  int         0 is an integer value
(42 + 0i)           float32     42.0 (with zero imaginary part) is in the set of float32 values
x                   T           x is not representable by a value of T because

0                   bool        0 is not in the set of boolean values
'a'                 string      'a' is a rune, it is not in the set of string values
1024                byte        1024 is not in the set of unsigned 8-bit integers
-1                  uint16      -1 is not in the set of unsigned 16-bit integers
1.1                 int         1.1 is not an integer value
42i                 float32     (0 + 42i) is not in the set of float32 values
1e1000              float64     1e1000 overflows to IEEE +Inf after rounding
```

## Blocks 积木

A *block* is a possibly empty sequence of declarations and statements within matching brace brackets.

块可能是匹配大括号中的声明和语句的空序列。

```
Block = "{" StatementList "}" .
StatementList = { Statement ";" } .
```

In addition to explicit blocks in the source code, there are implicit blocks:

除了源代码中的显式块，还有隐式块:

1. The 这个*universe block 宇宙块* encompasses all Go source text. 包含所有 Go 源文本
2. Each 每个人[package 包装](http://docscn.studygolang.com/ref/spec#Packages) has a 有*package block 封装块* containing all Go source text for that package. 包含该包的所有 Go 源文本
3. Each file has a 每个文件都有一个*file block 文件块* containing all Go source text in that file. 包含该文件中的所有 Go 源文本
4. Each 每个人["if" “如果”](http://docscn.studygolang.com/ref/spec#If_statements), ["for" “ for”](http://docscn.studygolang.com/ref/spec#For_statements), and ，及["switch" ”开关”](http://docscn.studygolang.com/ref/spec#Switch_statements) statement is considered to be in its own implicit block. 语句被认为在它自己的隐式块中
5. Each clause in a 中的每个子句["switch" ”开关”](http://docscn.studygolang.com/ref/spec#Switch_statements) or 或["select" ”选择”](http://docscn.studygolang.com/ref/spec#Select_statements) statement acts as an implicit block. 语句充当隐式块

Blocks nest and influence [scoping](http://docscn.studygolang.com/ref/spec#Declarations_and_scope).

块嵌套和影响范围。

## Declarations and scope 声明和范围

A *declaration* binds a non-[blank](http://docscn.studygolang.com/ref/spec#Blank_identifier) identifier to a [constant](http://docscn.studygolang.com/ref/spec#Constant_declarations), [type](http://docscn.studygolang.com/ref/spec#Type_declarations), [variable](http://docscn.studygolang.com/ref/spec#Variable_declarations), [function](http://docscn.studygolang.com/ref/spec#Function_declarations), [label](http://docscn.studygolang.com/ref/spec#Labeled_statements), or [package](http://docscn.studygolang.com/ref/spec#Import_declarations). Every identifier in a program must be declared. No identifier may be declared twice in the same block, and no identifier may be declared in both the file and package block.

声明将非空标识符绑定到常量、类型、变量、函数、标签或包。程序中的每个标识符都必须声明。不能在同一块中声明任何标识符两次，也不能在文件块和包块中声明任何标识符。

The [blank identifier](http://docscn.studygolang.com/ref/spec#Blank_identifier) may be used like any other identifier in a declaration, but it does not introduce a binding and thus is not declared. In the package block, the identifier `init` may only be used for [`init` function](http://docscn.studygolang.com/ref/spec#Package_initialization) declarations, and like the blank identifier it does not introduce a new binding.

空白标识符可以像声明中的任何其他标识符一样使用，但它不引入绑定，因此不声明。在包块中，标识符 init 只能用于 init 函数声明，并且与空白标识符一样，它不引入新的绑定。

```
Declaration   = ConstDecl | TypeDecl | VarDecl .
TopLevelDecl  = Declaration | FunctionDecl | MethodDecl .
```

The *scope* of a declared identifier is the extent of source text in which the identifier denotes the specified constant, type, variable, function, label, or package.

声明标识符的范围是源文本的范围，其中标识符表示指定的常量、类型、变量、函数、标签或包。

Go is lexically scoped using [blocks](http://docscn.studygolang.com/ref/spec#Blocks):

使用块来定义字典范围:

1. The scope of a 的范围[predeclared identifier 预先声明的标识符](http://docscn.studygolang.com/ref/spec#Predeclared_identifiers) is the universe block. 就是宇宙块
2. The scope of an identifier denoting a constant, type, variable, or function (but not method) declared at top level (outside any function) is the package block. 标识符的作用域表示在顶层(任何函数之外)声明的常量、类型、变量或函数(但不包括方法) ，即包块
3. The scope of the package name of an imported package is the file block of the file containing the import declaration. 导入包的包名称的范围是包含导入声明的文件的文件块
4. The scope of an identifier denoting a method receiver, function parameter, or result variable is the function body. 表示方法接收者、函数参数或结果变量的标识符的作用域是函数体
5. The scope of a constant or variable identifier declared inside a function begins at the end of the ConstSpec or VarSpec (ShortVarDecl for short variable declarations) and ends at the end of the innermost containing block. 函数内声明的常量或变量标识符的作用域从 constanspec 或 VarSpec (短变量声明的 ShortVarDecl)的结尾开始，到最内层包含块的结尾结束
6. The scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block. 函数内声明的类型标识符的作用域从 TypeSpec 中的标识符开始，到最里面的包含块结束

An identifier declared in a block may be redeclared in an inner block. While the identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration.

在块中声明的标识符可以在内部块中重新声明。当内部声明的标识符在范围内时，它表示由内部声明声明的实体。

The [package clause](http://docscn.studygolang.com/ref/spec#Package_clause) is not a declaration; the package name does not appear in any scope. Its purpose is to identify the files belonging to the same [package](http://docscn.studygolang.com/ref/spec#Packages) and to specify the default package name for import declarations.

包子句不是声明; 包名称不出现在任何范围中。其目的是标识属于同一个包的文件，并为导入声明指定默认的包名。

### Label scopes 标签范围

Labels are declared by [labeled statements](http://docscn.studygolang.com/ref/spec#Labeled_statements) and are used in the ["break"](http://docscn.studygolang.com/ref/spec#Break_statements), ["continue"](http://docscn.studygolang.com/ref/spec#Continue_statements), and ["goto"](http://docscn.studygolang.com/ref/spec#Goto_statements) statements. It is illegal to define a label that is never used. In contrast to other identifiers, labels are not block scoped and do not conflict with identifiers that are not labels. The scope of a label is the body of the function in which it is declared and excludes the body of any nested function.

标签由带标签的语句声明，用于“ break”、“ continue”和“ goto”语句。定义一个从未使用过的标签是非法的。与其他标识符不同，标识符不阻塞作用域，并且与不是标识符的标识符不冲突。标签的作用域是声明它的函数体，并排除任何嵌套函数体。

### Blank identifier 空白标识符

The *blank identifier* is represented by the underscore character `_`. It serves as an anonymous placeholder instead of a regular (non-blank) identifier and has special meaning in [declarations](http://docscn.studygolang.com/ref/spec#Declarations_and_scope), as an [operand](http://docscn.studygolang.com/ref/spec#Operands), and in [assignments](http://docscn.studygolang.com/ref/spec#Assignments).

空白标识符由下划线字符 _ 表示。它充当匿名占位符，而不是常规(非空白)标识符，在声明、操作数和赋值中具有特殊意义。

### Predeclared identifiers 预先声明的标识符

The following identifiers are implicitly declared in the [universe block](http://docscn.studygolang.com/ref/spec#Blocks):

在 universe 块中隐式声明了以下标识符:

```
Types:
	bool byte complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap close complex copy delete imag len
	make new panic print println real recover
```

### Exported identifiers 导出标识符

An identifier may be *exported* to permit access to it from another package. An identifier is exported if both:

可以导出一个标识符，以允许从另一个包访问该标识符。如果两者兼有，则导出标识符:

1. the first character of the identifier's name is a Unicode upper case letter (Unicode class "Lu"); and 标识符名称的第一个字符是 Unicode 大写字母(Unicode 类“ Lu”) ; 以及
2. the identifier is declared in the 标识符在[package block 封装块](http://docscn.studygolang.com/ref/spec#Blocks) or it is a 或者它是一个[field name 字段名](http://docscn.studygolang.com/ref/spec#Struct_types) or 或[method name 方法名](http://docscn.studygolang.com/ref/spec#MethodName).

All other identifiers are not exported.

不导出所有其他标识符。

### Uniqueness of identifiers 标识符的唯一性

Given a set of identifiers, an identifier is called *unique* if it is *different* from every other in the set. Two identifiers are different if they are spelled differently, or if they appear in different [packages](http://docscn.studygolang.com/ref/spec#Packages) and are not [exported](http://docscn.studygolang.com/ref/spec#Exported_identifiers). Otherwise, they are the same.

给定一组标识符，如果标识符与集合中的其他标识符不同，则称其为唯一标识符。如果两个标识符的拼写不同，或者它们出现在不同的包中并且没有导出，则它们是不同的。否则，它们是一样的。

### Constant declarations 不断的声明

A constant declaration binds a list of identifiers (the names of the constants) to the values of a list of [constant expressions](http://docscn.studygolang.com/ref/spec#Constant_expressions). The number of identifiers must be equal to the number of expressions, and the *n*th identifier on the left is bound to the value of the *n*th expression on the right.

常量声明将标识符列表(常量的名称)绑定到常量表达式列表的值。标识符的数量必须等于表达式的数量，左侧的第 n 个标识符与右侧的第 n 个表达式的值绑定。

```
ConstDecl      = "const" ( ConstSpec | "(" { ConstSpec ";" } ")" ) .
ConstSpec      = IdentifierList [ [ Type ] "=" ExpressionList ] .

IdentifierList = identifier { "," identifier } .
ExpressionList = Expression { "," Expression } .
```

If the type is present, all constants take the type specified, and the expressions must be [assignable](http://docscn.studygolang.com/ref/spec#Assignability) to that type. If the type is omitted, the constants take the individual types of the corresponding expressions. If the expression values are untyped [constants](http://docscn.studygolang.com/ref/spec#Constants), the declared constants remain untyped and the constant identifiers denote the constant values. For instance, if the expression is a floating-point literal, the constant identifier denotes a floating-point constant, even if the literal's fractional part is zero.

如果存在类型，则所有常量都采用指定的类型，表达式必须可分配给该类型。如果省略类型，则常量取相应表达式的单个类型。如果表达式值是非类型化常量，则声明的常量仍然是非类型化的，常量标识符表示常量值。例如，如果表达式是浮点文字，则常量标识符表示浮点常量，即使该文字的小数部分为零。

```
const Pi float64 = 3.14159265358979323846
const zero = 0.0         // untyped floating-point constant
const (
	size int64 = 1024
	eof        = -1  // untyped integer constant
)
const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", untyped integer and string constants
const u, v float32 = 0, 3    // u = 0.0, v = 3.0
```

Within a parenthesized `const` declaration list the expression list may be omitted from any but the first ConstSpec. Such an empty list is equivalent to the textual substitution of the first preceding non-empty expression list and its type if any. Omitting the list of expressions is therefore equivalent to repeating the previous list. The number of identifiers must be equal to the number of expressions in the previous list. Together with the [`iota` constant generator](http://docscn.studygolang.com/ref/spec#Iota) this mechanism permits light-weight declaration of sequential values:

在带括号的 const 声明列表中，除了第一个 ConstSpec 之外，表达式列表可以省略。这样一个空列表等价于前面第一个非空表达式列表的文本替换，如果有的话，还等价于它的类型。因此，省略表达式列表等同于重复前面的列表。标识符的数目必须等于前一个列表中的表达式的数目。该机制与 iota 常数生成器一起，允许轻量级声明顺序值:

```
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays  // this constant is not exported
)
```

### Iota 女名女子名

Within a [constant declaration](http://docscn.studygolang.com/ref/spec#Constant_declarations), the predeclared identifier `iota` represents successive untyped integer [constants](http://docscn.studygolang.com/ref/spec#Constants). Its value is the index of the respective [ConstSpec](http://docscn.studygolang.com/ref/spec#ConstSpec) in that constant declaration, starting at zero. It can be used to construct a set of related constants:

在一个常量声明中，预声明的标识符 iota 表示连续的非类型整数常量。它的值是该常量声明中各个 constanspec 的索引，从零开始。它可以用来构造一组相关的常量:

```
const (
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)

const (
	a = 1 << iota  // a == 1  (iota == 0)
	b = 1 << iota  // b == 2  (iota == 1)
	c = 3          // c == 3  (iota == 2, unused)
	d = 1 << iota  // d == 8  (iota == 3)
)

const (
	u         = iota * 42  // u == 0     (untyped integer constant)
	v float64 = iota * 42  // v == 42.0  (float64 constant)
	w         = iota * 42  // w == 84    (untyped integer constant)
)

const x = iota  // x == 0
const y = iota  // y == 0
```

By definition, multiple uses of `iota` in the same ConstSpec all have the same value:

根据定义，在相同的 constanspec 中多次使用 iota 都具有相同的值:

```
const (
	bit0, mask0 = 1 << iota, 1<<iota - 1  // bit0 == 1, mask0 == 0  (iota == 0)
	bit1, mask1                           // bit1 == 2, mask1 == 1  (iota == 1)
	_, _                                  //                        (iota == 2, unused)
	bit3, mask3                           // bit3 == 8, mask3 == 7  (iota == 3)
)
```

This last example exploits the [implicit repetition](http://docscn.studygolang.com/ref/spec#Constant_declarations) of the last non-empty expression list.

最后一个示例利用了最后一个非空表达式列表的隐式重复。

### Type declarations 类型声明

A type declaration binds an identifier, the *type name*, to a [type](http://docscn.studygolang.com/ref/spec#Types). Type declarations come in two forms: alias declarations and type definitions.

类型声明将标识符(类型名称)绑定到类型。类型声明有两种形式: 别名声明和类型定义。

```
TypeDecl = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) .
TypeSpec = AliasDecl | TypeDef .
```

#### Alias declarations

#### 别名声明

An alias declaration binds an identifier to the given type.

别名声明将标识符绑定到给定的类型。

```
AliasDecl = identifier "=" Type .
```

Within the [scope](http://docscn.studygolang.com/ref/spec#Declarations_and_scope) of the identifier, it serves as an *alias* for the type.

在标识符的范围内，它充当类型的别名。

```
type (
	nodeList = []*Node  // nodeList and []*Node are identical types
	Polar    = polar    // Polar and polar denote identical types
)
```

#### Type definitions

#### 类型定义

A type definition creates a new, distinct type with the same [underlying type](http://docscn.studygolang.com/ref/spec#Types) and operations as the given type, and binds an identifier to it.

类型定义创建具有与给定类型相同的基础类型和操作的新的不同类型，并将标识符绑定到该类型。

```
TypeDef = identifier Type .
```

The new type is called a *defined type*. It is [different](http://docscn.studygolang.com/ref/spec#Type_identity) from any other type, including the type it is created from.

新类型称为已定义类型。它不同于任何其他类型，包括它所创建的类型。

```
type (
	Point struct{ x, y float64 }  // Point and struct{ x, y float64 } are different types
	polar Point                   // polar and Point denote different types
)

type TreeNode struct {
	left, right *TreeNode
	value *Comparable
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}
```

A defined type may have [methods](http://docscn.studygolang.com/ref/spec#Method_declarations) associated with it. It does not inherit any methods bound to the given type, but the [method set](http://docscn.studygolang.com/ref/spec#Method_sets) of an interface type or of elements of a composite type remains unchanged:

定义的类型可能具有与其关联的方法。它不继承任何绑定到给定类型的方法，但接口类型或复合类型元素的方法集保持不变:

```
// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct         { /* Mutex fields */ }
func (m *Mutex) Lock()    { /* Lock implementation */ }
func (m *Mutex) Unlock()  { /* Unlock implementation */ }

// NewMutex has the same composition as Mutex but its method set is empty.
type NewMutex Mutex

// The method set of PtrMutex's underlying type *Mutex remains unchanged,
// but the method set of PtrMutex is empty.
type PtrMutex *Mutex

// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its embedded field Mutex.
type PrintableMutex struct {
	Mutex
}

// MyBlock is an interface type that has the same method set as Block.
type MyBlock Block
```

Type definitions may be used to define different boolean, numeric, or string types and associate methods with them:

类型定义可用于定义不同的布尔类型、数值类型或字符串类型，并将方法与它们关联:

```
type TimeZone int

const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

func (tz TimeZone) String() string {
	return fmt.Sprintf("GMT%+dh", tz)
}
```

### Variable declarations 变量声明

A variable declaration creates one or more [variables](http://docscn.studygolang.com/ref/spec#Variables), binds corresponding identifiers to them, and gives each a type and an initial value.

变量声明创建一个或多个变量，将相应的标识符绑定到这些变量，并为每个变量提供一个类型和初始值。

```
VarDecl     = "var" ( VarSpec | "(" { VarSpec ";" } ")" ) .
VarSpec     = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .
var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
	i       int
	u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]  // map lookup; only interested in "found"
```

If a list of expressions is given, the variables are initialized with the expressions following the rules for [assignments](http://docscn.studygolang.com/ref/spec#Assignments). Otherwise, each variable is initialized to its [zero value](http://docscn.studygolang.com/ref/spec#The_zero_value).

如果给定了表达式列表，则按照赋值规则使用表达式初始化变量。否则，将每个变量初始化为其零值。

If a type is present, each variable is given that type. Otherwise, each variable is given the type of the corresponding initialization value in the assignment. If that value is an untyped constant, it is first implicitly [converted](http://docscn.studygolang.com/ref/spec#Conversions) to its [default type](http://docscn.studygolang.com/ref/spec#Constants); if it is an untyped boolean value, it is first implicitly converted to type `bool`. The predeclared value `nil` cannot be used to initialize a variable with no explicit type.

如果存在类型，则为每个变量提供该类型。否则，每个变量都会在赋值中给出相应初始化值的类型。如果该值是非类型化常量，则首先将其隐式转换为默认类型; 如果该值是非类型化布尔值，则首先将其隐式转换为 bool 类型。预声明的值 nil 不能用于初始化没有显式类型的变量。

```
var d = math.Sin(0.5)  // d is float64
var i = 42             // i is int
var t, ok = x.(T)      // t is T, ok is bool
var n = nil            // illegal
```

Implementation restriction: A compiler may make it illegal to declare a variable inside a [function body](http://docscn.studygolang.com/ref/spec#Function_declarations) if the variable is never used.

实现限制: 如果一个变量从未被使用，编译器可能会使在函数体中声明该变量是非法的。

### Short variable declarations 短变量声明

A *short variable declaration* uses the syntax:

一个简短的变量声明使用了以下语法:

```
ShortVarDecl = IdentifierList ":=" ExpressionList .
```

It is shorthand for a regular [variable declaration](http://docscn.studygolang.com/ref/spec#Variable_declarations) with initializer expressions but no types:

它简化了带有初始化器表达式的正则变量声明，但没有类型:

```
"var" IdentifierList = ExpressionList .
i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w, _ := os.Pipe()  // os.Pipe() returns a connected pair of Files and an error, if any
_, y, _ := coord(p)   // coord() returns three values; only interested in y coordinate
```

Unlike regular variable declarations, a short variable declaration may *redeclare* variables provided they were originally declared earlier in the same block (or the parameter lists if the block is the function body) with the same type, and at least one of the non-[blank](http://docscn.studygolang.com/ref/spec#Blank_identifier) variables is new. As a consequence, redeclaration can only appear in a multi-variable short declaration. Redeclaration does not introduce a new variable; it just assigns a new value to the original.

与常规变量声明不同，短变量声明可以重新声明变量，前提是它们最初是在同一块中声明的(或者如果块是函数体，参数列表) ，并且具有相同的类型，而且至少有一个非空变量是新的。因此，重新声明只能出现在多变量短声明中。重新声明不会引入新的变量; 它只是将一个新值赋给原始变量。

```
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // redeclares offset
a, a := 1, 2                              // illegal: double declaration of a or no new variable if a was declared elsewhere
```

Short variable declarations may appear only inside functions. In some contexts such as the initializers for ["if"](http://docscn.studygolang.com/ref/spec#If_statements), ["for"](http://docscn.studygolang.com/ref/spec#For_statements), or ["switch"](http://docscn.studygolang.com/ref/spec#Switch_statements) statements, they can be used to declare local temporary variables.

短变量声明可能只出现在函数内部。在某些上下文中，例如“ if”、“ for”或“ switch”语句的初始化器，可以使用它们声明局部临时变量。

### Function declarations 函数声明

A function declaration binds an identifier, the *function name*, to a function.

函数声明将一个标识符(函数名)绑定到一个函数。

```
FunctionDecl = "func" FunctionName Signature [ FunctionBody ] .
FunctionName = identifier .
FunctionBody = Block .
```

If the function's [signature](http://docscn.studygolang.com/ref/spec#Function_types) declares result parameters, the function body's statement list must end in a [terminating statement](http://docscn.studygolang.com/ref/spec#Terminating_statements).

如果函数的签名声明结果参数，函数体的语句列表必须以终止语句结束。

```
func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	// invalid: missing return statement
}
```

A function declaration may omit the body. Such a declaration provides the signature for a function implemented outside Go, such as an assembly routine.

函数声明可能省略正文。这样的声明为在围棋之外实现的函数(如程序集例程)提供签名。

```
func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func flushICache(begin, end uintptr)  // implemented externally
```

### Method declarations 方法声明

A method is a [function](http://docscn.studygolang.com/ref/spec#Function_declarations) with a *receiver*. A method declaration binds an identifier, the *method name*, to a method, and associates the method with the receiver's *base type*.

方法是带有接收器的函数。方法声明将标识符(方法名)绑定到方法，并将该方法与接收方的基类型关联。

```
MethodDecl = "func" Receiver MethodName Signature [ FunctionBody ] .
Receiver   = Parameters .
```

The receiver is specified via an extra parameter section preceding the method name. That parameter section must declare a single non-variadic parameter, the receiver. Its type must be a [defined](http://docscn.studygolang.com/ref/spec#Type_definitions) type `T` or a pointer to a defined type `T`. `T` is called the receiver *base type*. A receiver base type cannot be a pointer or interface type and it must be defined in the same package as the method. The method is said to be *bound* to its receiver base type and the method name is visible only within [selectors](http://docscn.studygolang.com/ref/spec#Selectors) for type `T` or `*T`.

接收器是通过方法名称前面的一个额外参数部分指定的。参数部分必须声明一个非可变参数，即接收器。它的类型必须是已定义的类型 t，或者指向已定义类型 t 的指针 t 称为接收器基类型。接收器基类型不能是指针或接口类型，它必须在与方法相同的包中定义。该方法被称为绑定到其接收方基类型，并且方法名称仅在类型 t 或 * t 的选择器中可见。

A non-[blank](http://docscn.studygolang.com/ref/spec#Blank_identifier) receiver identifier must be [unique](http://docscn.studygolang.com/ref/spec#Uniqueness_of_identifiers) in the method signature. If the receiver's value is not referenced inside the body of the method, its identifier may be omitted in the declaration. The same applies in general to parameters of functions and methods.

非空接收方标识符在方法签名中必须是唯一的。如果方法体内没有引用接收方的值，则可以在声明中省略其标识符。这同样适用于函数和方法的参数。

For a base type, the non-blank names of methods bound to it must be unique. If the base type is a [struct type](http://docscn.studygolang.com/ref/spec#Struct_types), the non-blank method and field names must be distinct.

对于基类型，绑定到它的方法的非空名称必须是唯一的。如果基类型是结构类型，则非空方法和字段名称必须不同。

Given defined type `Point`, the declarations

给定定义的类型 Point，声明

```
func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
```

bind the methods `Length` and `Scale`, with receiver type `*Point`, to the base type `Point`.

将方法 Length 和 Scale 与 receiver type * Point 绑定到基类型 Point。

The type of a method is the type of a function with the receiver as first argument. For instance, the method `Scale` has type

方法的类型是以接收方为第一个参数的函数类型。例如，Scale 方法具有

```
func(p *Point, factor float64)
```

However, a function declared this way is not a method.

但是，以这种方式声明的函数不是方法。

## Expressions 表达方式

An expression specifies the computation of a value by applying operators and functions to operands.

表达式通过对操作数应用运算符和函数来指定值的计算。

### Operands 操作数

Operands denote the elementary values in an expression. An operand may be a literal, a (possibly [qualified](http://docscn.studygolang.com/ref/spec#Qualified_identifiers)) non-[blank](http://docscn.studygolang.com/ref/spec#Blank_identifier) identifier denoting a [constant](http://docscn.studygolang.com/ref/spec#Constant_declarations), [variable](http://docscn.studygolang.com/ref/spec#Variable_declarations), or [function](http://docscn.studygolang.com/ref/spec#Function_declarations), or a parenthesized expression.

操作数表示表达式中的基本值。操作数可以是文字、表示常量、变量或函数的(可能限定的)非空标识符或括号表达式。

The [blank identifier](http://docscn.studygolang.com/ref/spec#Blank_identifier) may appear as an operand only on the left-hand side of an [assignment](http://docscn.studygolang.com/ref/spec#Assignments).

空白标识符只能作为操作数出现在赋值的左侧。

```
Operand     = Literal | OperandName | "(" Expression ")" .
Literal     = BasicLit | CompositeLit | FunctionLit .
BasicLit    = int_lit | float_lit | imaginary_lit | rune_lit | string_lit .
OperandName = identifier | QualifiedIdent .
```

### Qualified identifiers 限定标识符

A qualified identifier is an identifier qualified with a package name prefix. Both the package name and the identifier must not be [blank](http://docscn.studygolang.com/ref/spec#Blank_identifier).

限定标识符是具有包名称前缀的限定标识符。包名和标识符都不能为空。

```
QualifiedIdent = PackageName "." identifier .
```

A qualified identifier accesses an identifier in a different package, which must be [imported](http://docscn.studygolang.com/ref/spec#Import_declarations). The identifier must be [exported](http://docscn.studygolang.com/ref/spec#Exported_identifiers) and declared in the [package block](http://docscn.studygolang.com/ref/spec#Blocks) of that package.

限定标识符访问不同包中的标识符，该标识符必须导入。标识符必须导出并在该包的包块中声明。

```
math.Sin	// denotes the Sin function in package math
```

### Composite literals 合成文字

Composite literals construct values for structs, arrays, slices, and maps and create a new value each time they are evaluated. They consist of the type of the literal followed by a brace-bound list of elements. Each element may optionally be preceded by a corresponding key.

复合文字构造结构、数组、片和映射的值，并在每次计算它们时创建一个新值。它们由文本类型和元素的大括号绑定列表组成。每个元素前面可以有一个对应的键。

```
CompositeLit  = LiteralType LiteralValue .
LiteralType   = StructType | ArrayType | "[" "..." "]" ElementType |
                SliceType | MapType | TypeName .
LiteralValue  = "{" [ ElementList [ "," ] ] "}" .
ElementList   = KeyedElement { "," KeyedElement } .
KeyedElement  = [ Key ":" ] Element .
Key           = FieldName | Expression | LiteralValue .
FieldName     = identifier .
Element       = Expression | LiteralValue .
```

The LiteralType's underlying type must be a struct, array, slice, or map type (the grammar enforces this constraint except when the type is given as a TypeName). The types of the elements and keys must be [assignable](http://docscn.studygolang.com/ref/spec#Assignability) to the respective field, element, and key types of the literal type; there is no additional conversion. The key is interpreted as a field name for struct literals, an index for array and slice literals, and a key for map literals. For map literals, all elements must have a key. It is an error to specify multiple elements with the same field name or constant key value. For non-constant map keys, see the section on [evaluation order](http://docscn.studygolang.com/ref/spec#Order_of_evaluation).

LiteralType 的基础类型必须是 struct、 array、 slice 或 map 类型(语法强制执行此约束，除非该类型作为 TypeName 给出)。元素和键的类型必须分配给文本类型的各个字段、元素和键类型; 不存在额外的转换。键被解释为结构文本的字段名、数组和片文本的索引以及映射文本的键。对于映射文字，所有元素都必须有一个键。指定具有相同字段名或常量键值的多个元素是错误的。有关非常量映射键，请参阅计算顺序一节。

For struct literals the following rules apply:

对于结构文本，应用以下规则:

- A key must be a field name declared in the struct type. 键必须是在结构类型中声明的字段名
- An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared. 不包含任何键的元素列表必须按声明字段的顺序列出每个结构字段的元素
- If any element has a key, every element must have a key. 如果任何元素都有一个键，那么每个元素都必须有一个键
- An element list that contains keys does not need to have an element for each struct field. Omitted fields get the zero value for that field. 包含键的元素列表不需要为每个结构字段包含元素。被省略的字段将获得该字段的零值
- A literal may omit the element list; such a literal evaluates to the zero value for its type. 常值可以省略元素列表; 这样的常值的计算结果为其类型的零值
- It is an error to specify an element for a non-exported field of a struct belonging to a different package. 为属于不同包的结构的非导出字段指定元素是错误的

Given the declarations

考虑到这些声明

```
type Point3D struct { x, y, z float64 }
type Line struct { p, q Point3D }
```

one may write

你可以写

```
origin := Point3D{}                            // zero value for Point3D
line := Line{origin, Point3D{y: -4, z: 12.3}}  // zero value for line.q.x
```

For array and slice literals the following rules apply:

对于数组和片文本，应用以下规则:

- Each element has an associated integer index marking its position in the array. 每个元素都有一个关联的整数索引，标记它在数组中的位置
- An element with a key uses the key as its index. The key must be a non-negative constant 具有键的元素使用键作为其索引。密钥必须是非负常数[representable 可代表的](http://docscn.studygolang.com/ref/spec#Representability) by a value of type 类型的值`int`; and if it is typed it must be of integer type. ; 如果输入的是 integer 类型，则必须为 integer 类型
- An element without a key uses the previous element's index plus one. If the first element has no key, its index is zero. 没有键的元素使用前一个元素的索引加1。如果第一个元素没有键，则其索引为零

[Taking the address](http://docscn.studygolang.com/ref/spec#Address_operators) of a composite literal generates a pointer to a unique [variable](http://docscn.studygolang.com/ref/spec#Variables) initialized with the literal's value.

获取复合文本的地址会生成一个指向用该文本的值初始化的唯一变量的指针。

```
var pointer *Point3D = &Point3D{y: 1000}
```

Note that the [zero value](http://docscn.studygolang.com/ref/spec#The_zero_value) for a slice or map type is not the same as an initialized but empty value of the same type. Consequently, taking the address of an empty slice or map composite literal does not have the same effect as allocating a new slice or map value with [new](http://docscn.studygolang.com/ref/spec#Allocation).

请注意，片或映射类型的零值与同一类型的已初始化但为空的值不同。因此，获取空片或映射复合文字的地址与使用 new 分配新片或映射值的效果不同。

```
p1 := &[]int{}    // p1 points to an initialized, empty slice with value []int{} and length 0
p2 := new([]int)  // p2 points to an uninitialized slice with value nil and length 0
```

The length of an array literal is the length specified in the literal type. If fewer elements than the length are provided in the literal, the missing elements are set to the zero value for the array element type. It is an error to provide elements with index values outside the index range of the array. The notation `...` specifies an array length equal to the maximum element index plus one.

数组常值的长度是在常值类型中指定的长度。如果文本中提供的元素少于长度，则缺少的元素将设置为数组元素类型的零值。将索引值提供给数组索引范围之外的元素是错误的。表示法... 指定的数组长度等于最大元素索引加1。

```
buffer := [10]string{}             // len(buffer) == 10
intSet := [6]int{1, 2, 3, 5}       // len(intSet) == 6
days := [...]string{"Sat", "Sun"}  // len(days) == 2
```

A slice literal describes the entire underlying array literal. Thus the length and capacity of a slice literal are the maximum element index plus one. A slice literal has the form

片文本描述整个基础数组文本。因此，片文字的长度和容量是最大元素索引加1。一个 slice literal 具有这样的形式

```
[]T{x1, x2, … xn}
```

and is shorthand for a slice operation applied to an array:

它是应用于数组的 slice 操作的简写形式:

```
tmp := [n]T{x1, x2, … xn}
tmp[0 : n]
```

Within a composite literal of array, slice, or map type `T`, elements or map keys that are themselves composite literals may elide the respective literal type if it is identical to the element or key type of `T`. Similarly, elements or keys that are addresses of composite literals may elide the `&T` when the element or key type is `*T`.

在由数组、片或映射类型 t 组成的复合文字中，如果元素或映射键本身是复合文字，那么如果它们与元素或键类型 t 相同，那么这些元素或键可能会省略相应的文字类型。

```
[...]Point{{1.5, -3.5}, {0, 0}}     // same as [...]Point{Point{1.5, -3.5}, Point{0, 0}}
[][]int{{1, 2, 3}, {4, 5}}          // same as [][]int{[]int{1, 2, 3}, []int{4, 5}}
[][]Point{{{0, 1}, {1, 2}}}         // same as [][]Point{[]Point{Point{0, 1}, Point{1, 2}}}
map[string]Point{"orig": {0, 0}}    // same as map[string]Point{"orig": Point{0, 0}}
map[Point]string{{0, 0}: "orig"}    // same as map[Point]string{Point{0, 0}: "orig"}

type PPoint *Point
[2]*Point{{1.5, -3.5}, {}}          // same as [2]*Point{&Point{1.5, -3.5}, &Point{}}
[2]PPoint{{1.5, -3.5}, {}}          // same as [2]PPoint{PPoint(&Point{1.5, -3.5}), PPoint(&Point{})}
```

A parsing ambiguity arises when a composite literal using the TypeName form of the LiteralType appears as an operand between the [keyword](http://docscn.studygolang.com/ref/spec#Keywords) and the opening brace of the block of an "if", "for", or "switch" statement, and the composite literal is not enclosed in parentheses, square brackets, or curly braces. In this rare case, the opening brace of the literal is erroneously parsed as the one introducing the block of statements. To resolve the ambiguity, the composite literal must appear within parentheses.

当使用 LiteralType 的 TypeName 形式的复合文本作为操作数出现在“ if”、“ for”或“ switch”语句的关键字和块的开大括号之间，并且复合文本不包含在括号、方括号或花括号中时，就会出现解析歧义。在这种罕见的情况下，文字的开括号被错误地解析为引入语句块的括号。若要解决歧义，复合文本必须出现在括号中。

```
if x == (T{a,b,c}[i]) { … }
if (x == T{a,b,c}[i]) { … }
```

Examples of valid array, slice, and map literals:

有效数组、片段和映射文字的示例:

```
// list of prime numbers
primes := []int{2, 3, 5, 7, 9, 2147483647}

// vowels[ch] is true if ch is a vowel
vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}

// the array [10]float32{-1, 0, 0, 0, -0.1, -0.1, 0, 0, 0, -1}
filter := [10]float32{-1, 4: -0.1, -0.1, 9: -1}

// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
noteFrequency := map[string]float32{
	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	"G0": 24.50, "A0": 27.50, "B0": 30.87,
}
```

### Function literals 函数文字

A function literal represents an anonymous [function](http://docscn.studygolang.com/ref/spec#Function_declarations).

匿名函数表示一个匿名函数。

```
FunctionLit = "func" Signature FunctionBody .
func(a, b int, z float64) bool { return a*b < int(z) }
```

A function literal can be assigned to a variable or invoked directly.

可以为变量赋值或直接调用匿名函数。

```
f := func(x, y int) int { return x + y }
func(ch chan int) { ch <- ACK }(replyChan)
```

Function literals are *closures*: they may refer to variables defined in a surrounding function. Those variables are then shared between the surrounding function and the function literal, and they survive as long as they are accessible.

函数文字是闭包: 它们可以引用在周围函数中定义的变量。然后，这些变量在周围的函数和匿名函数之间共享，只要它们可以被访问，它们就会继续存在。

### Primary expressions 主要表达式

Primary expressions are the operands for unary and binary expressions.

主表达式是一元和二元表达式的操作数。

```
PrimaryExpr =
	Operand |
	Conversion |
	MethodExpr |
	PrimaryExpr Selector |
	PrimaryExpr Index |
	PrimaryExpr Slice |
	PrimaryExpr TypeAssertion |
	PrimaryExpr Arguments .

Selector       = "." identifier .
Index          = "[" Expression "]" .
Slice          = "[" [ Expression ] ":" [ Expression ] "]" |
                 "[" [ Expression ] ":" Expression ":" Expression "]" .
TypeAssertion  = "." "(" Type ")" .
Arguments      = "(" [ ( ExpressionList | Type [ "," ExpressionList ] ) [ "..." ] [ "," ] ] ")" .
x
2
(s + ".txt")
f(3.1415, true)
Point{1, 2}
m["foo"]
s[i : j + 1]
obj.color
f.p[i].x()
```

### Selectors 选择器

For a [primary expression](http://docscn.studygolang.com/ref/spec#Primary_expressions) `x` that is not a [package name](http://docscn.studygolang.com/ref/spec#Package_clause), the *selector expression*

对于不是包名称的主表达式 x，选择器表达式

```
x.f
```

denotes the field or method `f` of the value `x` (or sometimes `*x`; see below). The identifier `f` is called the (field or method) *selector*; it must not be the [blank identifier](http://docscn.studygolang.com/ref/spec#Blank_identifier). The type of the selector expression is the type of `f`. If `x` is a package name, see the section on [qualified identifiers](http://docscn.studygolang.com/ref/spec#Qualified_identifiers).

表示值 x 的字段或方法 f (有时是 * x; 见下文)。标识符 f 称为(字段或方法)选择器; 它不能是空标识符。选择器表达式的类型是 f 类型。如果 x 是一个包名，请参阅限定标识符部分。

A selector `f` may denote a field or method `f` of a type `T`, or it may refer to a field or method `f` of a nested [embedded field](http://docscn.studygolang.com/ref/spec#Struct_types) of `T`. The number of embedded fields traversed to reach `f` is called its *depth* in `T`. The depth of a field or method `f` declared in `T` is zero. The depth of a field or method `f` declared in an embedded field `A` in `T` is the depth of `f` in `A` plus one.

选择器 f 可以表示类型 t 的字段或方法 f，也可以表示嵌套嵌入字段 t 的字段或方法 f。穿越到达 f 的嵌入域数称为它在 t 中的深度。在 t 中声明的字段或方法 f 的深度为零。在嵌入字段 a 中声明的字段或方法 f 的深度是 a 加1中 f 的深度。

The following rules apply to selectors:

以下规则适用于选择器:

1. For a value 对于一个值`x` of type 类型`T` or 或`*T` where 在哪里`T` is not a pointer or interface type, 不是指针或接口类型,`x.f` denotes the field or method at the shallowest depth in 表示在最浅的深度处的田野或方法`T` where there is such an 哪里有这样一个`f`. If there is not exactly 。如果没有完全[one 一`f`](http://docscn.studygolang.com/ref/spec#Uniqueness_of_identifiers) with shallowest depth, the selector expression is illegal. 深度最浅时，选择器表达式是非法的
2. For a value 对于一个值`x` of type 类型`I` where 在哪里`I` is an interface type, 是一种接口类型,`x.f` denotes the actual method with name 用 name 表示实际的方法`f` of the dynamic value of 的动态值`x`. If there is no method with name 。如果没有带名称的方法`f` in the 在[method set 方法集](http://docscn.studygolang.com/ref/spec#Method_sets) of 的`I`, the selector expression is illegal. ，选择器表达式是非法的
3. As an exception, if the type of 作为一个例外，如果`x` is a 是一个[defined 定义](http://docscn.studygolang.com/ref/spec#Type_definitions) pointer type and 指针类型和`(*x).f` is a valid selector expression denoting a field (but not a method), 是表示字段(但不是方法)的有效选择器表达式,`x.f` is shorthand for 是对`(*x).f`.
4. In all other cases, 在所有其他情况下,`x.f` is illegal. 是违法的
5. If 如果`x` is of pointer type and has the value 是指针类型，并且具有`nil` and 及`x.f` denotes a struct field, assigning to or evaluating 表示一个结构字段，分配给或计算`x.f` causes a 引起[run-time panic 运行时恐慌](http://docscn.studygolang.com/ref/spec#Run_time_panics).
6. If 如果`x` is of interface type and has the value 是接口类型，并具有值`nil`, [calling 呼叫](http://docscn.studygolang.com/ref/spec#Calls) or 或[evaluating 评价](http://docscn.studygolang.com/ref/spec#Method_values) the method 方法`x.f` causes a 引起[run-time panic 运行时恐慌](http://docscn.studygolang.com/ref/spec#Run_time_panics).

For example, given the declarations:

例如，给定声明:

```
type T0 struct {
	x int
}

func (*T0) M0()

type T1 struct {
	y int
}

func (T1) M1()

type T2 struct {
	z int
	T1
	*T0
}

func (*T2) M2()

type Q *T2

var t T2     // with t.T0 != nil
var p *T2    // with p != nil and (*p).T0 != nil
var q Q = p
```

one may write:

有人可能会写道:

```
t.z          // t.z
t.y          // t.T1.y
t.x          // (*t.T0).x

p.z          // (*p).z
p.y          // (*p).T1.y
p.x          // (*(*p).T0).x

q.x          // (*(*q).T0).x        (*q).x is a valid field selector

p.M0()       // ((*p).T0).M0()      M0 expects *T0 receiver
p.M1()       // ((*p).T1).M1()      M1 expects T1 receiver
p.M2()       // p.M2()              M2 expects *T2 receiver
t.M2()       // (&t).M2()           M2 expects *T2 receiver, see section on Calls
```

but the following is invalid:

但下面的内容是无效的:

```
q.M0()       // (*q).M0 is valid but not a field selector
```

### Method expressions 方法表达式

If `M` is in the [method set](http://docscn.studygolang.com/ref/spec#Method_sets) of type `T`, `T.M` is a function that is callable as a regular function with the same arguments as `M` prefixed by an additional argument that is the receiver of the method.

如果 m 在 t 类型的方法集中，T.M 是一个可调用的正则函数，其参数与 m 的参数相同，前缀是一个附加的参数，即方法的接收者。

```
MethodExpr    = ReceiverType "." MethodName .
ReceiverType  = Type .
```

Consider a struct type `T` with two methods, `Mv`, whose receiver is of type `T`, and `Mp`, whose receiver is of type `*T`.

考虑一个具有两个方法的结构类型 t，Mv (其接收器为 t 类型)和 Mp (其接收器为 t 类型)。

```
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver

var t T
```

The expression

这个表情

```
T.Mv
```

yields a function equivalent to `Mv` but with an explicit receiver as its first argument; it has signature

产生一个等价于 Mv 的函数，但第一个参数是显式接收器; 它有签名

```
func(tv T, a int) int
```

That function may be called normally with an explicit receiver, so these five invocations are equivalent:

这个函数可以通过一个显式接收器正常调用，所以这五个调用是等价的:

```
t.Mv(7)
T.Mv(t, 7)
(T).Mv(t, 7)
f1 := T.Mv; f1(t, 7)
f2 := (T).Mv; f2(t, 7)
```

Similarly, the expression

类似地，表达式

```
(*T).Mp
```

yields a function value representing `Mp` with signature

产生一个函数值，表示带有签名的 Mp

```
func(tp *T, f float32) float32
```

For a method with a value receiver, one can derive a function with an explicit pointer receiver, so

对于带有值接收器的方法，可以派生一个带有显式指针接收器的函数，因此

```
(*T).Mv
```

yields a function value representing `Mv` with signature

产生一个函数值，表示带有签名的 Mv

```
func(tv *T, a int) int
```

Such a function indirects through the receiver to create a value to pass as the receiver to the underlying method; the method does not overwrite the value whose address is passed in the function call.

这样的函数通过接收器中转创建一个值，作为接收器传递给底层方法; 该方法不会覆盖地址在函数调用中传递的值。

The final case, a value-receiver function for a pointer-receiver method, is illegal because pointer-receiver methods are not in the method set of the value type.

最后一种情况，即指针接收器方法的值接收器函数，是非法的，因为指针接收器方法不在值类型的方法集中。

Function values derived from methods are called with function call syntax; the receiver is provided as the first argument to the call. That is, given `f := T.Mv`, `f` is invoked as `f(t, 7)` not `t.f(7)`. To construct a function that binds the receiver, use a [function literal](http://docscn.studygolang.com/ref/spec#Function_literals) or [method value](http://docscn.studygolang.com/ref/spec#Method_values).

从方法派生的函数值按照函数调用语法调用; 接收方作为调用的第一个参数提供。也就是说，给定 f: = T.Mv，f 被调用为 f (t，7)而不是 t.f (7)。要构造一个绑定接收方的函数，可以使用匿名函数或方法值。

It is legal to derive a function value from a method of an interface type. The resulting function takes an explicit receiver of that interface type.

从接口类型的方法派生函数值是合法的。结果函数接收该接口类型的显式接收器。

### Method values 方法值

If the expression `x` has static type `T` and `M` is in the [method set](http://docscn.studygolang.com/ref/spec#Method_sets) of type `T`, `x.M` is called a *method value*. The method value `x.M` is a function value that is callable with the same arguments as a method call of `x.M`. The expression `x` is evaluated and saved during the evaluation of the method value; the saved copy is then used as the receiver in any calls, which may be executed later.

如果表达式 x 具有静态类型 t，而 m 位于类型 t 的方法集中，则称 x.M 为方法值。方法值 x.M 是一个可调用的函数值，它具有与方法调用 x.M 相同的参数。表达式 x 在方法值计算期间进行计算和保存; 然后在任何调用中将保存的副本用作接收方，这些调用可能在以后执行。

The type `T` may be an interface or non-interface type.

类型 t 可以是接口类型或非接口类型。

As in the discussion of [method expressions](http://docscn.studygolang.com/ref/spec#Method_expressions) above, consider a struct type `T` with two methods, `Mv`, whose receiver is of type `T`, and `Mp`, whose receiver is of type `*T`.

正如在上面对方法表达式的讨论中一样，考虑一个结构类型 t，它有两个方法，Mv (其接收器为 t 类型)和 Mp (其接收器为 t 类型)。

```
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver

var t T
var pt *T
func makeT() T
```

The expression

这个表情

```
t.Mv
```

yields a function value of type

生成类型的函数值

```
func(int) int
```

These two invocations are equivalent:

这两个调用是等价的:

```
t.Mv(7)
f := t.Mv; f(7)
```

Similarly, the expression

类似地，表达式

```
pt.Mp
```

yields a function value of type

生成类型的函数值

```
func(float32) float32
```

As with [selectors](http://docscn.studygolang.com/ref/spec#Selectors), a reference to a non-interface method with a value receiver using a pointer will automatically dereference that pointer: `pt.Mv` is equivalent to `(*pt).Mv`.

与选择器一样，对具有使用指针的值接收器的非接口方法的引用将自动取消引用该指针: pt.Mv 等效于(* pt)。我。

As with [method calls](http://docscn.studygolang.com/ref/spec#Calls), a reference to a non-interface method with a pointer receiver using an addressable value will automatically take the address of that value: `t.Mp` is equivalent to `(&t).Mp`.

与方法调用一样，对具有使用可寻址值的指针接收器的非接口方法的引用将自动获取该值的地址: t.Mp 等效于(& t)。国会议员。

```
f := t.Mv; f(7)   // like t.Mv(7)
f := pt.Mp; f(7)  // like pt.Mp(7)
f := pt.Mv; f(7)  // like (*pt).Mv(7)
f := t.Mp; f(7)   // like (&t).Mp(7)
f := makeT().Mp   // invalid: result of makeT() is not addressable
```

Although the examples above use non-interface types, it is also legal to create a method value from a value of interface type.

尽管上面的示例使用非接口类型，但是从接口类型的值创建方法值也是合法的。

```
var i interface { M(int) } = myVal
f := i.M; f(7)  // like i.M(7)
```

### Index expressions 索引表达式

A primary expression of the form

形式的主要表达式

```
a[x]
```

denotes the element of the array, pointer to array, slice, string or map `a` indexed by `x`. The value `x` is called the *index* or *map key*, respectively. The following rules apply:

表示数组的元素、指向数组的指针、片、字符串或映射一个 x 索引的数组。值 x 分别称为索引键或映射键。以下规则适用:

If `a` is not a map:

如果 a 不是地图:

- the index 索引`x` must be of integer type or an untyped constant 必须是整数类型或非类型化常量
- a constant index must be non-negative and 常数索引必须是非负的[representable 可代表的](http://docscn.studygolang.com/ref/spec#Representability) by a value of type 类型的值`int`
- a constant index that is untyped is given type 未定义类型的常量索引是给定类型的`int`
- the index 索引`x` is 是*in range 在射程内* if 如果`0 <= x < len(a)`, otherwise it is 否则就是了*out of range 在射程之外*

For `a` of [array type](http://docscn.studygolang.com/ref/spec#Array_types) `A`:

对于 a 类型的数组:

- a [constant 常数](http://docscn.studygolang.com/ref/spec#Constants) index must be in range 索引必须在范围内
- if 如果`x` is out of range at run time, a 在运行时超出了范围[run-time panic 运行时恐慌](http://docscn.studygolang.com/ref/spec#Run_time_panics) occurs 发生
- `a[x]` is the array element at index 是索引处的数组元素`x` and the type of 以及类型`a[x]` is the element type of 是元素类型`A`

For `a` of [pointer](http://docscn.studygolang.com/ref/spec#Pointer_types) to array type:

对于指向数组类型的指针:

- `a[x]` is shorthand for 是对`(*a)[x]`

For `a` of [slice type](http://docscn.studygolang.com/ref/spec#Slice_types) `S`:

对于 s 类型的片:

- if 如果`x` is out of range at run time, a 在运行时超出了范围[run-time panic 运行时恐慌](http://docscn.studygolang.com/ref/spec#Run_time_panics) occurs 发生
- `a[x]` is the slice element at index 是位于索引处的 slice 元素`x` and the type of 以及类型`a[x]` is the element type of 是元素类型`S`

For `a` of [string type](http://docscn.studygolang.com/ref/spec#String_types):

对于字符串类型:

- a [constant 常数](http://docscn.studygolang.com/ref/spec#Constants) index must be in range if the string 索引必须在范围内如果字符串`a` is also constant 也是不变的
- if 如果`x` is out of range at run time, a 在运行时超出了范围[run-time panic 运行时恐慌](http://docscn.studygolang.com/ref/spec#Run_time_panics) occurs 发生
- `a[x]` is the non-constant byte value at index 是索引处的非常量字节值`x` and the type of 以及类型`a[x]` is 是`byte`
- `a[x]` may not be assigned to 可能不会被分配给

For `a` of [map type](http://docscn.studygolang.com/ref/spec#Map_types) `M`:

对于类型为 m 的 a:

- `x`'s type must be 的类型必须[assignable 可分配的](http://docscn.studygolang.com/ref/spec#Assignability) to the key type of 到关键类型`M`
- if the map contains an entry with key 如果地图包含一个带有密钥的条目`x`, `a[x]` is the map element with key 就是带钥匙的地图元素`x` and the type of 以及类型`a[x]` is the element type of 是元素类型`M`
- if the map is 如果地图是`nil` or does not contain such an entry, 或不包含这样的条目,`a[x]` is the 是[zero value 零值](http://docscn.studygolang.com/ref/spec#The_zero_value) for the element type of 对于元素类型`M`

Otherwise `a[x]` is illegal.

否则 a [ x ]是非法的。

An index expression on a map `a` of type `map[K]V` used in an [assignment](http://docscn.studygolang.com/ref/spec#Assignments) or initialization of the special form

类型映射[ k ] v 的映射 a 上的索引表达式，用于特殊形式的赋值或初始化

```
v, ok = a[x]
v, ok := a[x]
var v, ok = a[x]
```

yields an additional untyped boolean value. The value of `ok` is `true` if the key `x` is present in the map, and `false` otherwise.

会产生一个额外的无类型布尔值。如果键 x 出现在 map 中，则 ok 的值为 true，否则为 false。

Assigning to an element of a `nil` map causes a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics).

分配给 nil 映射的元素会导致运行时恐慌。

### Slice expressions 片段表达式

Slice expressions construct a substring or slice from a string, array, pointer to array, or slice. There are two variants: a simple form that specifies a low and high bound, and a full form that also specifies a bound on the capacity.

片表达式从字符串、数组、指向数组的指针或片构造子字符串或片。有两种变体: 一种是指定上下限的简单表单，另一种是也指定容量界限的完整表单。

#### Simple slice expressions

#### 简单片表达式

For a string, array, pointer to array, or slice `a`, the primary expression

对于字符串、数组、指向数组的指针或片 a，则为主表达式

```
a[low : high]
```

constructs a substring or slice. The *indices* `low` and `high` select which elements of operand `a` appear in the result. The result has indices starting at 0 and length equal to `high` - `low`. After slicing the array `a`

构造子字符串或片。索引的低和高选择哪些元素的操作数出现在结果中。结果是指数从0开始，长度等于高低。在对数组进行切片之后,

```
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]
```

the slice `s` has type `[]int`, length 3, capacity 4, and elements

片 s 具有类型[] int、长度3、容量4和元素

```
s[0] == 2
s[1] == 3
s[2] == 4
```

For convenience, any of the indices may be omitted. A missing `low` index defaults to zero; a missing `high` index defaults to the length of the sliced operand:

为了方便起见，任何索引都可以省略。缺少的低索引缺省为零; 缺少的高索引缺省为分片操作数的长度:

```
a[2:]  // same as a[2 : len(a)]
a[:3]  // same as a[0 : 3]
a[:]   // same as a[0 : len(a)]
```

If `a` is a pointer to an array, `a[low : high]` is shorthand for `(*a)[low : high]`.

如果 a 是指向数组的指针，则[ low: high ]表示(* a)[ low: high ]。

For arrays or strings, the indices are *in range* if `0` <= `low` <= `high` <= `len(a)`, otherwise they are *out of range*. For slices, the upper index bound is the slice capacity `cap(a)` rather than the length. A [constant](http://docscn.studygolang.com/ref/spec#Constants) index must be non-negative and [representable](http://docscn.studygolang.com/ref/spec#Representability) by a value of type `int`; for arrays or constant strings, constant indices must also be in range. If both indices are constant, they must satisfy `low <= high`. If the indices are out of range at run time, a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics) occurs.

对于数组或字符串，如果0 < = low < = high < = len (a) ，则索引在范围内，否则它们超出范围。对于切片，上索引界限是切片容量帽(a) ，而不是长度。常量索引必须是非负的，并且可由 int 类型的值表示; 对于数组或常量字符串，常量索引也必须在范围内。如果这两个指数都是常量，那么它们必须满足 low < = high。如果指数在运行时超出范围，就会发生运行时恐慌。

Except for [untyped strings](http://docscn.studygolang.com/ref/spec#Constants), if the sliced operand is a string or slice, the result of the slice operation is a non-constant value of the same type as the operand. For untyped string operands the result is a non-constant value of type `string`. If the sliced operand is an array, it must be [addressable](http://docscn.studygolang.com/ref/spec#Address_operators) and the result of the slice operation is a slice with the same element type as the array.

除了非类型化的字符串之外，如果被分片的操作数是字符串或片，片操作的结果是与操作数类型相同的非常量值。对于非类型字符串操作数，结果是类型为 string 的非常量值。如果切片操作数是一个数组，那么它必须是可寻址的，并且切片操作的结果是一个与数组具有相同元素类型的切片。

If the sliced operand of a valid slice expression is a `nil` slice, the result is a `nil` slice. Otherwise, if the result is a slice, it shares its underlying array with the operand.

如果有效片表达式的分片操作数为 nil 片，则结果为 nil 片。否则，如果结果是 slice，它将与操作数共享其基础数组。

```
var a [10]int
s1 := a[3:7]   // underlying array of s1 is array a; &s1[2] == &a[5]
s2 := s1[1:4]  // underlying array of s2 is underlying array of s1 which is array a; &s2[1] == &a[5]
s2[1] = 42     // s2[1] == s1[2] == a[5] == 42; they all refer to the same underlying array element
```

#### Full slice expressions

#### 完整片段表达式

For an array, pointer to array, or slice `a` (but not a string), the primary expression

对于数组，指向数组的指针或切片(但不是字符串)是主表达式

```
a[low : high : max]
```

constructs a slice of the same type, and with the same length and elements as the simple slice expression `a[low : high]`. Additionally, it controls the resulting slice's capacity by setting it to `max - low`. Only the first index may be omitted; it defaults to 0. After slicing the array `a`

构造与简单片表达式 a [ low: high ]具有相同长度和元素的相同类型的片。此外，它通过将切片的容量设置为最大值-低来控制产生的切片的容量。只有第一个索引可以省略; 它的缺省值为0。对数组进行切片之后

```
a := [5]int{1, 2, 3, 4, 5}
t := a[1:3:5]
```

the slice `t` has type `[]int`, length 2, capacity 4, and elements

片 t 具有类型[] int、长度2、容量4和元素

```
t[0] == 2
t[1] == 3
```

As for simple slice expressions, if `a` is a pointer to an array, `a[low : high : max]` is shorthand for `(*a)[low : high : max]`. If the sliced operand is an array, it must be [addressable](http://docscn.studygolang.com/ref/spec#Address_operators).

对于简单的片表达式，如果 a 是指向数组的指针，则[ low: high: max ]是(* a)[ low: high: max ]的简写。如果被分片的操作数是数组，那么它必须是可寻址的。

The indices are *in range* if `0 <= low <= high <= max <= cap(a)`, otherwise they are *out of range*. A [constant](http://docscn.studygolang.com/ref/spec#Constants) index must be non-negative and [representable](http://docscn.studygolang.com/ref/spec#Representability) by a value of type `int`; for arrays, constant indices must also be in range. If multiple indices are constant, the constants that are present must be in range relative to each other. If the indices are out of range at run time, a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics) occurs.

如果0 < = 低 < = 高 < = max < = cap (a) ，则指数在范围内，否则超出范围。常量索引必须是非负的，并且可由 int 类型的值表示; 对于数组，常量索引也必须在范围内。如果多个索引是常量，则存在的常量必须在相对于彼此的范围内。如果指数在运行时超出范围，就会发生运行时恐慌。

### Type assertions 输入断言

For an expression `x` of [interface type](http://docscn.studygolang.com/ref/spec#Interface_types) and a type `T`, the primary expression

对于接口类型的表达式 x 和类型 t，主表达式

```
x.(T)
```

asserts that `x` is not `nil` and that the value stored in `x` is of type `T`. The notation `x.(T)` is called a *type assertion*.

断言 x 不是零，存储在 x 中的值是 t 类型的。记号 x (t)称为类型断言。

More precisely, if `T` is not an interface type, `x.(T)` asserts that the dynamic type of `x` is [identical](http://docscn.studygolang.com/ref/spec#Type_identity) to the type `T`. In this case, `T` must [implement](http://docscn.studygolang.com/ref/spec#Method_sets) the (interface) type of `x`; otherwise the type assertion is invalid since it is not possible for `x` to store a value of type `T`. If `T` is an interface type, `x.(T)` asserts that the dynamic type of `x` implements the interface `T`.

更准确地说，如果 t 不是接口类型，则 x (t)断言 x 的动态类型与 t 类型相同。在这种情况下，t 必须实现 x 的(接口)类型; 否则类型断言无效，因为 x 不可能存储类型 t 的值。如果 t 是接口类型，则 x (t)断言 x 的动态类型实现了接口 t。

If the type assertion holds, the value of the expression is the value stored in `x` and its type is `T`. If the type assertion is false, a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics) occurs. In other words, even though the dynamic type of `x` is known only at run time, the type of `x.(T)` is known to be `T` in a correct program.

如果类型断言成立，则表达式的值为存储在 x 中的值，其类型为 t。如果类型断言为 false，则会发生运行时恐慌。换句话说，即使只有在运行时才知道 x 的动态类型，但是在正确的程序中，x (t)的类型是 t。

```
var x interface{} = 7          // x has dynamic type int and value 7
i := x.(int)                   // i has type int and value 7

type I interface { m() }

func f(y I) {
	s := y.(string)        // illegal: string does not implement I (missing method m)
	r := y.(io.Reader)     // r has type io.Reader and the dynamic type of y must implement both I and io.Reader
	…
}
```

A type assertion used in an [assignment](http://docscn.studygolang.com/ref/spec#Assignments) or initialization of the special form

在特殊窗体的赋值或初始化中使用的类型断言

```
v, ok = x.(T)
v, ok := x.(T)
var v, ok = x.(T)
var v, ok T1 = x.(T)
```

yields an additional untyped boolean value. The value of `ok` is `true` if the assertion holds. Otherwise it is `false` and the value of `v` is the [zero value](http://docscn.studygolang.com/ref/spec#The_zero_value) for type `T`. No [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics) occurs in this case.

会产生一个额外的无类型布尔值。如果断言成立，则 ok 的值为 true。否则为 false，并且类型 t 的 v 值为零。在这种情况下，不会发生运行时恐慌。

### Calls 电话

Given an expression `f` of function type `F`,

给定函数类型 f 的表达式,

```
f(a1, a2, … an)
```

calls `f` with arguments `a1, a2, … an`. Except for one special case, arguments must be single-valued expressions [assignable](http://docscn.studygolang.com/ref/spec#Assignability) to the parameter types of `F` and are evaluated before the function is called. The type of the expression is the result type of `F`. A method invocation is similar but the method itself is specified as a selector upon a value of the receiver type for the method.

调用 f，参数 a1，a2，... an。除了一个特殊情况外，参数必须是单值表达式，可分配给 f 的参数类型，并在调用函数之前进行计算。表达式的类型是 f 的结果类型。方法调用是类似的，但是方法本身被指定为方法的接收方类型的值上的选择器。

```
math.Atan2(x, y)  // function call
var pt *Point
pt.Scale(3.5)     // method call with receiver pt
```

In a function call, the function value and arguments are evaluated in [the usual order](http://docscn.studygolang.com/ref/spec#Order_of_evaluation). After they are evaluated, the parameters of the call are passed by value to the function and the called function begins execution. The return parameters of the function are passed by value back to the calling function when the function returns.

在函数调用中，函数值和参数按照通常的顺序计算。计算之后，调用的参数通过值传递给函数，被调用的函数开始执行。当函数返回时，函数的返回参数通过值返回给调用函数。

Calling a `nil` function value causes a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics).

调用 nil 函数值会导致运行时恐慌。

As a special case, if the return values of a function or method `g` are equal in number and individually assignable to the parameters of another function or method `f`, then the call `f(g(*parameters_of_g*))` will invoke `f` after binding the return values of `g` to the parameters of `f` in order. The call of `f` must contain no parameters other than the call of `g`, and `g` must have at least one return value. If `f` has a final `...` parameter, it is assigned the return values of `g` that remain after assignment of regular parameters.

作为一种特殊情况，如果一个函数或方法 g 的返回值在数量上是相等的，并且可以单独地分配给另一个函数或方法 f 的参数，那么调用 f (g (参数 _ of _ g))在依次将 g 的返回值绑定到 f 的参数之后将调用 f。F 的调用除了 g 的调用之外不能包含任何参数，g 必须至少有一个返回值。如果 f 有一个 final... 参数，它被赋值为 g 的返回值，这些返回值在赋值常规参数之后仍然存在。

```
func Split(s string, pos int) (string, string) {
	return s[0:pos], s[pos:]
}

func Join(s, t string) string {
	return s + t
}

if Join(Split(value, len(value)/2)) != value {
	log.Panic("test fails")
}
```

A method call `x.m()` is valid if the [method set](http://docscn.studygolang.com/ref/spec#Method_sets) of (the type of) `x` contains `m` and the argument list can be assigned to the parameter list of `m`. If `x` is [addressable](http://docscn.studygolang.com/ref/spec#Address_operators) and `&x`'s method set contains `m`, `x.m()` is shorthand for `(&x).m()`:

如果方法集(类型) x 包含 m，并且参数列表可以分配给参数列表 m，那么方法调用 x.m ()是有效的。如果 x 是可寻址的，而 & x 的方法集包含 m，那么 x.m ()是(& x)的简写形式。M () :

```
var p Point
p.Scale(3.5)
```

There is no distinct method type and there are no method literals.

没有明确的方法类型，也没有方法文本。

### Passing arguments to 将参数传递给`...` parameters 参数

If `f` is [variadic](http://docscn.studygolang.com/ref/spec#Function_types) with a final parameter `p` of type `...T`, then within `f` the type of `p` is equivalent to type `[]T`. If `f` is invoked with no actual arguments for `p`, the value passed to `p` is `nil`. Otherwise, the value passed is a new slice of type `[]T` with a new underlying array whose successive elements are the actual arguments, which all must be [assignable](http://docscn.studygolang.com/ref/spec#Assignability) to `T`. The length and capacity of the slice is therefore the number of arguments bound to `p` and may differ for each call site.

如果 f 是可变参数，最后一个参数 p 是... t 型，那么在 f 中，p 型等价于[] t 型。如果在没有实际参数的情况下调用 f，则传递给 p 的值为 nil。否则，传递的值是一个新的片[] t 类型，带有一个新的基础数组，其后续元素是实际的参数，所有这些参数都必须可以分配给 t。因此，切片的长度和容量是绑定到 p 的参数数量，并且每个调用站点的参数数量可能不同。

Given the function and calls

给定函数和调用

```
func Greeting(prefix string, who ...string)
Greeting("nobody")
Greeting("hello:", "Joe", "Anna", "Eileen")
```

within `Greeting`, `who` will have the value `nil` in the first call, and `[]string{"Joe", "Anna", "Eileen"}` in the second.

在 Greeting 中，第一个调用的值为 nil，第二个调用的值为[]字符串{“ Joe”、“ Anna”、“ Eileen”}。

If the final argument is assignable to a slice type `[]T`, it is passed unchanged as the value for a `...T` parameter if the argument is followed by `...`. In this case no new slice is created.

如果最后一个参数可以分配给一个 slice 类型[] t，那么如果参数后面跟着... ，它将不变地作为... t 参数的值传递。在这种情况下，不会创建新的片。

Given the slice `s` and call

给定切片和调用

```
s := []string{"James", "Jasmine"}
Greeting("goodbye:", s...)
```

within `Greeting`, `who` will have the same value as `s` with the same underlying array.

在 Greeting 中，它将具有与 s 相同的基础数组值。

### Operators 操作员

Operators combine operands into expressions.

运算符将操作数组合成表达式。

```
Expression = UnaryExpr | Expression binary_op Expression .
UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .
```

Comparisons are discussed [elsewhere](http://docscn.studygolang.com/ref/spec#Comparison_operators). For other binary operators, the operand types must be [identical](http://docscn.studygolang.com/ref/spec#Type_identity) unless the operation involves shifts or untyped [constants](http://docscn.studygolang.com/ref/spec#Constants). For operations involving constants only, see the section on [constant expressions](http://docscn.studygolang.com/ref/spec#Constant_expressions).

比较在其他地方讨论。对于其他二进制运算符，操作数类型必须相同，除非该操作涉及移位或非类型化常量。对于仅涉及常量的操作，请参阅常量表达式一节。

Except for shift operations, if one operand is an untyped [constant](http://docscn.studygolang.com/ref/spec#Constants) and the other operand is not, the constant is implicitly [converted](http://docscn.studygolang.com/ref/spec#Conversions) to the type of the other operand.

除了移位操作之外，如果一个操作数是无类型常量，而另一个操作数不是，则该常量将隐式转换为另一个操作数的类型。

The right operand in a shift expression must have integer type or be an untyped constant [representable](http://docscn.studygolang.com/ref/spec#Representability) by a value of type `uint`. If the left operand of a non-constant shift expression is an untyped constant, it is first implicitly converted to the type it would assume if the shift expression were replaced by its left operand alone.

移位表达式中的右操作数必须具有整数类型，或者是可由 uint 类型的值表示的非类型常量。如果非常量移位表达式的左操作数是非类型化常量，则首先隐式转换为如果移位表达式被其左操作数单独替换时它将假定的类型。

```
var s uint = 33
var i = 1<<s                  // 1 has type int
var j int32 = 1<<s            // 1 has type int32; j == 0
var k = uint64(1<<s)          // 1 has type uint64; k == 1<<33
var m int = 1.0<<s            // 1.0 has type int; m == 0 if ints are 32bits in size
var n = 1.0<<s == j           // 1.0 has type int32; n == true
var o = 1<<s == 2<<s          // 1 and 2 have type int; o == true if ints are 32bits in size
var p = 1<<s == 1<<33         // illegal if ints are 32bits in size: 1 has type int, but 1<<33 overflows int
var u = 1.0<<s                // illegal: 1.0 has type float64, cannot shift
var u1 = 1.0<<s != 0          // illegal: 1.0 has type float64, cannot shift
var u2 = 1<<s != 1.0          // illegal: 1 has type float64, cannot shift
var v float32 = 1<<s          // illegal: 1 has type float32, cannot shift
var w int64 = 1.0<<33         // 1.0<<33 is a constant shift expression
var x = a[1.0<<s]             // 1.0 has type int; x == a[0] if ints are 32bits in size
var a = make([]byte, 1.0<<s)  // 1.0 has type int; len(a) == 0 if ints are 32bits in size
```

#### Operator precedence

#### 运算符优先级

Unary operators have the highest precedence. As the `++` and `--` operators form statements, not expressions, they fall outside the operator hierarchy. As a consequence, statement `*p++` is the same as `(*p)++`.

一元运算符的优先级最高。由于 + + 和 -- 运算符形成语句而不是表达式，因此它们不属于运算符层次结构。因此，语句 * p + + 与(* p) + + 相同。

There are five precedence levels for binary operators. Multiplication operators bind strongest, followed by addition operators, comparison operators, `&&` (logical AND), and finally `||` (logical OR):

二进制运算符有五个优先级。乘法运算符绑定最强，其次是加法运算符、比较运算符和 & (逻辑与) ，最后是 | | (逻辑 OR) :

```
Precedence    Operator
    5             *  /  %  <<  >>  &  &^
    4             +  -  |  ^
    3             ==  !=  <  <=  >  >=
    2             &&
    1             ||
```

Binary operators of the same precedence associate from left to right. For instance, `x / y * z` is the same as `(x / y) * z`.

相同优先级的二进制运算符从左到右关联。例如，x/y * z 与(x/y) * z 相同。

```
+x
23 + 3*x[i]
x <= f()
^a >> b
f() || g()
x == y+1 && <-chanPtr > 0
```

### Arithmetic operators 算术运算符

Arithmetic operators apply to numeric values and yield a result of the same type as the first operand. The four standard arithmetic operators (`+`, `-`, `*`, `/`) apply to integer, floating-point, and complex types; `+` also applies to strings. The bitwise logical and shift operators apply to integers only.

算术运算符应用于数值，并产生与第一个操作数类型相同的结果。四个标准算术运算符(+ ,-，* ,/)适用于整数、浮点和复杂类型; + 也适用于字符串。位逻辑运算符和移位运算符仅适用于整数。

```
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&^   bit clear (AND NOT)    integers

<<   left shift             integer << unsigned integer
>>   right shift            integer >> unsigned integer
```

#### Integer operators

#### 整数运算符

For two integer values `x` and `y`, the integer quotient `q = x / y` and remainder `r = x % y` satisfy the following relationships:

对于两个整数值 x 和 y，整数商 q = x/y 和余数 r = x% y 满足以下关系:

```
x = q*y + r  and  |r| < |y|
```

with `x / y` truncated towards zero (["truncated division"](https://en.wikipedia.org/wiki/Modulo_operation)).

X/y 缩短到零(“缩短除法”)。

```
 x     y     x / y     x % y
 5     3       1         2
-5     3      -1        -2
 5    -3      -1         2
-5    -3       1        -2
```

The one exception to this rule is that if the dividend `x` is the most negative value for the int type of `x`, the quotient `q = x / -1` is equal to `x` (and `r = 0`) due to two's-complement [integer overflow](http://docscn.studygolang.com/ref/spec#Integer_overflow):

这个规则的一个例外是，如果红利 x 是整数类型 x 的最负值，由于两个补整数溢出，商 q = x/-1等于 x (和 r = 0) :

```
			 x, q
int8                     -128
int16                  -32768
int32             -2147483648
int64    -9223372036854775808
```

If the divisor is a [constant](http://docscn.studygolang.com/ref/spec#Constants), it must not be zero. If the divisor is zero at run time, a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics) occurs. If the dividend is non-negative and the divisor is a constant power of 2, the division may be replaced by a right shift, and computing the remainder may be replaced by a bitwise AND operation:

如果除数是一个常数，它一定不能为零。如果在运行时除数为零，就会发生运行时恐慌。如果被除数是非负的，除数是2的常数幂，除法可以用右移代替，余数的计算可以用位与运算代替:

```
 x     x / 4     x % 4     x >> 2     x & 3
 11      2         3         2          3
-11     -2        -3        -3          1
```

The shift operators shift the left operand by the shift count specified by the right operand, which must be non-negative. If the shift count is negative at run time, a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics) occurs. The shift operators implement arithmetic shifts if the left operand is a signed integer and logical shifts if it is an unsigned integer. There is no upper limit on the shift count. Shifts behave as if the left operand is shifted `n` times by 1 for a shift count of `n`. As a result, `x << 1` is the same as `x*2` and `x >> 1` is the same as `x/2` but truncated towards negative infinity.

移位操作符通过右操作数指定的移位计数来移位左操作数，移位计数必须是非负的。如果在运行时移位计数为负，则会发生运行时恐慌。如果左操作数是有符号整数，则移位操作符实现算术移位; 如果是无符号整数，则实现逻辑移位。移位计数没有上限。移位表现为左操作数移位 n 次，移位次数 n 次。因此，x < 1与 x * 2相同，x > 1与 x/2相同，只是向负无穷方向截断。

For integer operands, the unary operators `+`, `-`, and `^` are defined as follows:

对于整数操作数，一元运算符 + 、-和 ^ 的定义如下:

```
+x                          is 0 + x
-x    negation              is 0 - x
^x    bitwise complement    is m ^ x  with m = "all bits set to 1" for unsigned x
                                      and  m = -1 for signed x
```

#### Integer overflow

#### 整数溢出

For unsigned integer values, the operations `+`, `-`, `*`, and `<<` are computed modulo 2*n*, where *n* is the bit width of the [unsigned integer](http://docscn.studygolang.com/ref/spec#Numeric_types)'s type. Loosely speaking, these unsigned integer operations discard high bits upon overflow, and programs may rely on "wrap around".

对于无符号整数值，+ 、-、 * 和 < < 都是计算模2n，其中 n 是无符号整数类型的位宽度。不严格地说，这些无符号整数操作在溢出时丢弃高位，而程序可能依赖于“换行”。

For signed integers, the operations `+`, `-`, `*`, `/`, and `<<` may legally overflow and the resulting value exists and is deterministically defined by the signed integer representation, the operation, and its operands. Overflow does not cause a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics). A compiler may not optimize code under the assumption that overflow does not occur. For instance, it may not assume that `x < x + 1` is always true.

对于有符号整数，+ 、-、 * 、/和 < < 操作可以合法地溢出，结果值存在，并由有符号整数表示形式、操作数及其操作数确定。溢出不会引起运行时恐慌。在不发生溢出的假设下，编译器可能不会优化代码。例如，它可能不会假定 x < x + 1总是正确的。

#### Floating-point operators

#### 浮点运算符

For floating-point and complex numbers, `+x` is the same as `x`, while `-x` is the negation of `x`. The result of a floating-point or complex division by zero is not specified beyond the IEEE-754 standard; whether a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics) occurs is implementation-specific.

对于浮点数和复数，+ x 等于 x，而-x 等于 x 的负数。在 IEEE-754标准之外，没有指定浮点或复杂除零的结果; 是否发生运行时恐慌是特定于实现的。

An implementation may combine multiple floating-point operations into a single fused operation, possibly across statements, and produce a result that differs from the value obtained by executing and rounding the instructions individually. An explicit floating-point type [conversion](http://docscn.studygolang.com/ref/spec#Conversions) rounds to the precision of the target type, preventing fusion that would discard that rounding.

一个实现可以将多个浮点操作组合成一个单独的融合操作(可能跨语句) ，并产生与单独执行和舍入指令所获得的值不同的结果。显式的浮点类型转换回合到目标类型的精度，防止融合丢弃舍入。

For instance, some architectures provide a "fused multiply and add" (FMA) instruction that computes `x*y + z` without rounding the intermediate result `x*y`. These examples show when a Go implementation can use that instruction:

例如，一些体系结构提供了“融合乘法和加法”(FMA)指令，该指令计算 x * y + z，而不取整中间结果 x * y。这些例子显示了 Go 实现何时可以使用该指令:

```
// FMA allowed for computing r, because x*y is not explicitly rounded:
r  = x*y + z
r  = z;   r += x*y
t  = x*y; r = t + z
*p = x*y; r = *p + z
r  = x*y + float64(z)

// FMA disallowed for computing r, because it would omit rounding of x*y:
r  = float64(x*y) + z
r  = z; r += float64(x*y)
t  = float64(x*y); r = t + z
```

#### String concatenation

#### 字符串串联

Strings can be concatenated using the `+` operator or the `+=` assignment operator:

字符串可以使用 + 运算符或 + = 赋值运算符串联:

```
s := "hi" + string(c)
s += " and good bye"
```

String addition creates a new string by concatenating the operands.

字符串加法通过串联操作数创建一个新字符串。

### Comparison operators 比较操作符

Comparison operators compare two operands and yield an untyped boolean value.

比较运算符比较两个操作数并产生一个无类型的布尔值。

```
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal
```

In any comparison, the first operand must be [assignable](http://docscn.studygolang.com/ref/spec#Assignability) to the type of the second operand, or vice versa.

在任何比较中，第一个操作数必须可分配给第二个操作数的类型，反之亦然。

The equality operators `==` and `!=` apply to operands that are *comparable*. The ordering operators `<`, `<=`, `>`, and `>=` apply to operands that are *ordered*. These terms and the result of the comparisons are defined as follows:

相等运算符 = = 和！= 适用于具有可比性的操作数。排序运算符 < 、 < = 、 > 和 > = 适用于排序的操作数。这些术语和比较结果的定义如下:

- Boolean values are comparable. Two boolean values are equal if they are either both 布尔值是可比较的。如果两个布尔值都是，那么它们是相等的`true` or both 或者两者都有`false`.
- Integer values are comparable and ordered, in the usual way. 整数值是可比较的，并按照通常的方式排序
- Floating-point values are comparable and ordered, as defined by the IEEE-754 standard. 浮点值是可比的和有序的，正如 IEEE-754标准所定义的
- Complex values are comparable. Two complex values 复数值是可比的。两个复数值`u` and 及`v` are equal if both 都是相等的`real(u) == real(v)` and 及`imag(u) == imag(v)`.
- String values are comparable and ordered, lexically byte-wise. 字符串值是可比较的，并且是按字节排序的
- Pointer values are comparable. Two pointer values are equal if they point to the same variable or if both have value 指针值是可比较的。如果两个指针值指向同一个变量，或者两者都有值，那么两个指针值是相等的`nil`. Pointers to distinct . 指向 distinct 的指针[zero-size 零尺寸的](http://docscn.studygolang.com/ref/spec#Size_and_alignment_guarantees) variables may or may not be equal. 变量可能相等，也可能不相等
- Channel values are comparable. Two channel values are equal if they were created by the same call to 通道值是可比较的。如果两个通道值是由调用[`make`](http://docscn.studygolang.com/ref/spec#Making_slices_maps_and_channels) or if both have value 或者两者都有价值`nil`.
- Interface values are comparable. Two interface values are equal if they have 界面值是可比的。如果有两个界面值，则两个界面值相等[identical 相同的](http://docscn.studygolang.com/ref/spec#Type_identity) dynamic types and equal dynamic values or if both have value 动态类型和相等的动态值，或者两者都有值`nil`.
- A value 一个值`x` of non-interface type 非接口类型`X` and a value 和一个值`t` of interface type 接口类型`T` are comparable when values of type 类型的数值是可比的`X` are comparable and 是可以比较的`X` implements 工具`T`. They are equal if 。他们是平等的，如果`t`'s dynamic type is identical to 的动态类型与`X` and 及`t`'s dynamic value is equal to 的动态值等于`x`.
- Struct values are comparable if all their fields are comparable. Two struct values are equal if their corresponding non- 如果结构值的所有字段都是可比较的，则它们的值是可比较的[blank 空白的](http://docscn.studygolang.com/ref/spec#Blank_identifier) fields are equal. 字段是相等的
- Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal. 如果数组元素类型的值是可比较的，则数组值是可比较的。如果两个数组的相应元素相等，则它们的值相等

A comparison of two interface values with identical dynamic types causes a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics) if values of that type are not comparable. This behavior applies not only to direct interface value comparisons but also when comparing arrays of interface values or structs with interface-valued fields.

如果两个接口值的动态类型相同，则比较该类型的值会导致运行时恐慌。此行为不仅适用于直接的接口值比较，还适用于将接口值或结构的数组与接口值字段进行比较。

Slice, map, and function values are not comparable. However, as a special case, a slice, map, or function value may be compared to the predeclared identifier `nil`. Comparison of pointer, channel, and interface values to `nil` is also allowed and follows from the general rules above.

Slice、 map 和函数值是不可比的。但是，作为一种特殊情况，片、映射或函数值可以与预声明的标识符 nil 进行比较。还允许将指针、通道和接口值比较为 nil，并遵循上面的一般规则。

```
const c = 3 < 4            // c is the untyped boolean constant true

type MyBool bool
var x, y int
var (
	// The result of a comparison is an untyped boolean.
	// The usual assignment rules apply.
	b3        = x == y // b3 has type bool
	b4 bool   = x == y // b4 has type bool
	b5 MyBool = x == y // b5 has type MyBool
)
```

### Logical operators 逻辑运算符

Logical operators apply to [boolean](http://docscn.studygolang.com/ref/spec#Boolean_types) values and yield a result of the same type as the operands. The right operand is evaluated conditionally.

逻辑运算符应用于布尔值，并产生与操作数类型相同的结果。右操作数按条件求值。

```
&&    conditional AND    p && q  is  "if p then q else false"
||    conditional OR     p || q  is  "if p then true else q"
!     NOT                !p      is  "not p"
```

### Address operators 地址操作符

For an operand `x` of type `T`, the address operation `&x` generates a pointer of type `*T` to `x`. The operand must be *addressable*, that is, either a variable, pointer indirection, or slice indexing operation; or a field selector of an addressable struct operand; or an array indexing operation of an addressable array. As an exception to the addressability requirement, `x` may also be a (possibly parenthesized) [composite literal](http://docscn.studygolang.com/ref/spec#Composite_literals). If the evaluation of `x` would cause a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics), then the evaluation of `&x` does too.

对于类型为 t 的操作数 x，地址操作 & x 生成类型为 * t 到 x 的指针。操作数必须是可寻址的，即变量、指针间接或片索引操作; 或可寻址结构操作数的字段选择器; 或可寻址数组的数组索引操作。作为可寻址性要求的一个例外，x 也可以是一个(可能是括号中的)复合文本。如果对 x 的求值会引起运行时恐慌，那么对 & x 的求值也会引起恐慌。

For an operand `x` of pointer type `*T`, the pointer indirection `*x` denotes the [variable](http://docscn.studygolang.com/ref/spec#Variables) of type `T` pointed to by `x`. If `x` is `nil`, an attempt to evaluate `*x` will cause a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics).

对于指针类型 * t 的操作数 x，指针间接方向 * x 表示类型 t 指向 x 的变量。如果 x 为 nil，则尝试计算 * x 将导致运行时恐慌。

```
&x
&a[f(2)]
&Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // causes a run-time panic
&*x  // causes a run-time panic
```

### Receive operator 接收操作员

For an operand `ch` of [channel type](http://docscn.studygolang.com/ref/spec#Channel_types), the value of the receive operation `<-ch` is the value received from the channel `ch`. The channel direction must permit receive operations, and the type of the receive operation is the element type of the channel. The expression blocks until a value is available. Receiving from a `nil` channel blocks forever. A receive operation on a [closed](http://docscn.studygolang.com/ref/spec#Close) channel can always proceed immediately, yielding the element type's [zero value](http://docscn.studygolang.com/ref/spec#The_zero_value) after any previously sent values have been received.

对于通道类型的操作数 ch，接收操作 <-ch 的值是从通道 ch 接收的值。通道方向必须允许接收操作，而接收操作的类型是通道的元素类型。表达式会阻塞，直到一个值可用为止。永远接收无线频道的讯号。关闭通道上的接收操作始终可以立即执行，在接收到任何以前发送的值之后，生成元素类型的零值。

```
v1 := <-ch
v2 = <-ch
f(<-ch)
<-strobe  // wait until clock pulse and discard received value
```

A receive expression used in an [assignment](http://docscn.studygolang.com/ref/spec#Assignments) or initialization of the special form

在特殊窗体的赋值或初始化中使用的接收表达式

```
x, ok = <-ch
x, ok := <-ch
var x, ok = <-ch
var x, ok T = <-ch
```

yields an additional untyped boolean result reporting whether the communication succeeded. The value of `ok` is `true` if the value received was delivered by a successful send operation to the channel, or `false` if it is a zero value generated because the channel is closed and empty.

产生一个额外的无类型布尔结果，报告通信是否成功。如果通过成功的发送操作将接收到的值传递到通道，则 ok 的值为 true; 如果值为零，则为 false，因为通道已关闭且为空。

### Conversions 转化率

A conversion changes the [type](http://docscn.studygolang.com/ref/spec#Types) of an expression to the type specified by the conversion. A conversion may appear literally in the source, or it may be *implied* by the context in which an expression appears.

转换将表达式的类型更改为转换所指定的类型。转换可能按字面意思出现在源中，也可能由表达式出现的上下文所暗示。

An *explicit* conversion is an expression of the form `T(x)` where `T` is a type and `x` is an expression that can be converted to type `T`.

显式转换是形式 t (x)的表达式，其中 t 是一个类型，x 是一个可以转换为类型 t 的表达式。

```
Conversion = Type "(" Expression [ "," ] ")" .
```

If the type starts with the operator `*` or `<-`, or if the type starts with the keyword `func` and has no result list, it must be parenthesized when necessary to avoid ambiguity:

如果类型以操作符 * 或 <-开始，或者类型以关键字 func 开始，没有结果列表，则必须在必要时加上括号以避免歧义:

```
*Point(p)        // same as *(Point(p))
(*Point)(p)      // p is converted to *Point
<-chan int(c)    // same as <-(chan int(c))
(<-chan int)(c)  // c is converted to <-chan int
func()(x)        // function signature func() x
(func())(x)      // x is converted to func()
(func() int)(x)  // x is converted to func() int
func() int(x)    // x is converted to func() int (unambiguous)
```

A [constant](http://docscn.studygolang.com/ref/spec#Constants) value `x` can be converted to type `T` if `x` is [representable](http://docscn.studygolang.com/ref/spec#Representability) by a value of `T`. As a special case, an integer constant `x` can be explicitly converted to a [string type](http://docscn.studygolang.com/ref/spec#String_types) using the [same rule](http://docscn.studygolang.com/ref/spec#Conversions_to_and_from_a_string_type) as for non-constant `x`.

如果 x 可以用 t 的值表示，那么常量值 x 可以转换为类型 t。作为一种特殊情况，整数常量 x 可以使用与非常量 x 相同的规则显式转换为字符串类型。

Converting a constant yields a typed constant as result.

转换一个常量会产生一个类型化的常量。

```
uint(iota)               // iota value of type uint
float32(2.718281828)     // 2.718281828 of type float32
complex128(1)            // 1.0 + 0.0i of type complex128
float32(0.49999999)      // 0.5 of type float32
float64(-1e-1000)        // 0.0 of type float64
string('x')              // "x" of type string
string(0x266c)           // "♬" of type string
MyString("foo" + "bar")  // "foobar" of type MyString
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
int(1.2)                 // illegal: 1.2 cannot be represented as an int
string(65.0)             // illegal: 65.0 is not an integer constant
```

A non-constant value `x` can be converted to type `T` in any of these cases:

在以下任何情况下，非常数值 x 都可以转换为类型 t:

- `x` is 是[assignable 可分配的](http://docscn.studygolang.com/ref/spec#Assignability) to 到`T`.
- ignoring struct tags (see below), 忽略结构标记(见下文) ,`x`'s type and 的类型`T` have 有[identical 相同的](http://docscn.studygolang.com/ref/spec#Type_identity) [underlying types 基本类型](http://docscn.studygolang.com/ref/spec#Types).
- ignoring struct tags (see below), 忽略结构标记(见下文) ,`x`'s type and 的类型`T` are pointer types that are not 是指针类型，不是[defined types 定义类型](http://docscn.studygolang.com/ref/spec#Type_definitions), and their pointer base types have identical underlying types. ，并且它们的指针基类型具有相同的基础类型
- `x`'s type and 的类型`T` are both integer or floating point types. 都是整型或浮点型
- `x`'s type and 的类型`T` are both complex types. 都是复杂的类型
- `x` is an integer or a slice of bytes or runes and 是一个整数或一个字节片段或符文`T` is a string type. 是一个字符串类型
- `x` is a string and 是一根绳子`T` is a slice of bytes or runes. 是字节或符文的切片

[Struct tags](http://docscn.studygolang.com/ref/spec#Struct_types) are ignored when comparing struct types for identity for the purpose of conversion:

为了转换的目的，比较结构类型的标识时忽略结构标记:

```
type Person struct {
	Name    string
	Address *struct {
		Street string
		City   string
	}
}

var data *struct {
	Name    string `json:"name"`
	Address *struct {
		Street string `json:"street"`
		City   string `json:"city"`
	} `json:"address"`
}

var person = (*Person)(data)  // ignoring tags, the underlying types are identical
```

Specific rules apply to (non-constant) conversions between numeric types or to and from a string type. These conversions may change the representation of `x` and incur a run-time cost. All other conversions only change the type but not the representation of `x`.

特定规则适用于数值类型之间或与字符串类型之间的(非常数)转换。这些转换可能会改变 x 的表示形式，并产生运行时成本。所有其他转换只改变类型，而不改变 x 的表示形式。

There is no linguistic mechanism to convert between pointers and integers. The package [`unsafe`](http://docscn.studygolang.com/ref/spec#Package_unsafe) implements this functionality under restricted circumstances.

没有语言机制来在指针和整数之间进行转换。包不安全在受限制的情况下实现这个功能。

#### Conversions between numeric types

#### 数值类型之间的转换

For the conversion of non-constant numeric values, the following rules apply:

对于非常量数值的转换，适用以下规则:

1. When converting between integer types, if the value is a signed integer, it is sign extended to implicit infinite precision; otherwise it is zero extended. It is then truncated to fit in the result type's size. For example, if 在整数类型之间进行转换时，如果值是有符号整数，则将其符号扩展为隐式无限精度; 否则为零扩展。然后将其截断以适应结果类型的大小。例如，如果`v := uint16(0x10F0)`, then 那么`uint32(int8(v)) == 0xFFFFFFF0`. The conversion always yields a valid value; there is no indication of overflow. 。转换总是产生一个有效值; 没有溢出指示
2. When converting a floating-point number to an integer, the fraction is discarded (truncation towards zero). 当将浮点数转换为整数时，分数将被丢弃(截断向零)
3. When converting an integer or floating-point number to a floating-point type, or a complex number to another complex type, the result value is rounded to the precision specified by the destination type. For instance, the value of a variable 当将整数或浮点数转换为浮点数类型，或将复数转换为其他复杂类型时，结果值将四舍五入到目标类型指定的精度。例如，变量的值`x` of type 类型`float32` may be stored using additional precision beyond that of an IEEE-754 32-bit number, but float32(x) represents the result of rounding 可以使用 ieee-75432位数之外的额外精度进行存储，但 float32(x)表示舍入的结果`x`'s value to 32-bit precision. Similarly, 值设置为32位精度,`x + 0.1` may use more than 32 bits of precision, but 可能使用超过32位的精度，但`float32(x + 0.1)` does not. 没有

In all non-constant conversions involving floating-point or complex values, if the result type cannot represent the value the conversion succeeds but the result value is implementation-dependent.

在所有涉及浮点值或复杂值的非常数转换中，如果结果类型不能表示值，则转换成功，但结果值依赖于实现。

#### Conversions to and from a string type

#### 与字符串类型的转换

1. Converting a signed or unsigned integer value to a string type yields a string containing the UTF-8 representation of the integer. Values outside the range of valid Unicode code points are converted to

    

   将有符号或无符号整数值转换为字符串类型会生成包含整数的 UTF-8表示形式的字符串。有效 Unicode 代码点范围以外的值转换为

   ```
   "\uFFFD"
   ```

   .

   ```
   string('a')       // "a"
   string(-1)        // "\ufffd" == "\xef\xbf\xbd"
   string(0xf8)      // "\u00f8" == "ø" == "\xc3\xb8"
   type MyString string
   MyString(0x65e5)  // "\u65e5" == "日" == "\xe6\x97\xa5"
   ```

2. Converting a slice of bytes to a string type yields a string whose successive bytes are the elements of the slice.

    

   将字节片转换为字符串类型会产生一个字符串，其后续的字节是片的元素

   ```
   string([]byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'})   // "hellø"
   string([]byte{})                                     // ""
   string([]byte(nil))                                  // ""
   
   type MyBytes []byte
   string(MyBytes{'h', 'e', 'l', 'l', '\xc3', '\xb8'})  // "hellø"
   ```

3. Converting a slice of runes to a string type yields a string that is the concatenation of the individual rune values converted to strings.

    

   将符文切片转换为字符串类型会生成一个字符串，该字符串是转换为字符串的各个符文值的串联

   ```
   string([]rune{0x767d, 0x9d6c, 0x7fd4})   // "\u767d\u9d6c\u7fd4" == "白鵬翔"
   string([]rune{})                         // ""
   string([]rune(nil))                      // ""
   
   type MyRunes []rune
   string(MyRunes{0x767d, 0x9d6c, 0x7fd4})  // "\u767d\u9d6c\u7fd4" == "白鵬翔"
   ```

4. Converting a value of a string type to a slice of bytes type yields a slice whose successive elements are the bytes of the string.

    

   将字符串类型的值转换为字节类型的片会产生一个片，其后续元素是字符串的字节

   ```
   []byte("hellø")   // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
   []byte("")        // []byte{}
   
   MyBytes("hellø")  // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
   ```

5. Converting a value of a string type to a slice of runes type yields a slice containing the individual Unicode code points of the string.

    

   将一个字符串类型的值转换为一个 runes 类型的片，将产生一个包含该字符串的单个 Unicode 代码点的片

   ```
   []rune(MyString("白鵬翔"))  // []rune{0x767d, 0x9d6c, 0x7fd4}
   []rune("")                 // []rune{}
   
   MyRunes("白鵬翔")           // []rune{0x767d, 0x9d6c, 0x7fd4}
   ```

### Constant expressions 常量表达式

Constant expressions may contain only [constant](http://docscn.studygolang.com/ref/spec#Constants) operands and are evaluated at compile time.

常量表达式可以只包含常量操作数，并在编译时计算。

Untyped boolean, numeric, and string constants may be used as operands wherever it is legal to use an operand of boolean, numeric, or string type, respectively.

非类型化的布尔型、数值型和字符串常量可以用作操作数，只要分别使用布尔型、数值型或字符串类型的操作数是合法的。

A constant [comparison](http://docscn.studygolang.com/ref/spec#Comparison_operators) always yields an untyped boolean constant. If the left operand of a constant [shift expression](http://docscn.studygolang.com/ref/spec#Operators) is an untyped constant, the result is an integer constant; otherwise it is a constant of the same type as the left operand, which must be of [integer type](http://docscn.studygolang.com/ref/spec#Numeric_types).

常量比较总是产生一个无类型的布尔常量。如果常量移位表达式的左操作数是一个无类型常量，则结果是一个整数常量; 否则，它是一个与左操作数类型相同的常量，该常量必须是整数类型。

Any other operation on untyped constants results in an untyped constant of the same kind; that is, a boolean, integer, floating-point, complex, or string constant. If the untyped operands of a binary operation (other than a shift) are of different kinds, the result is of the operand's kind that appears later in this list: integer, rune, floating-point, complex. For example, an untyped integer constant divided by an untyped complex constant yields an untyped complex constant.

对非类型化常量的任何其他操作都会生成同类型的非类型化常量，即布尔型常量、整数型常量、浮点型常量、复数型常量或字符串常量。如果一个二元运算的非类型化操作数(shift 除外)是不同类型的，那么结果就是操作数类型的结果，这将在下面的列表中出现: 整数、符文、浮点数、复数。例如，一个非类型化的整数常量除以一个非类型化的复数常量，就会产生一个非类型化的复数常量。

```
const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 << 3.0         // d == 8     (untyped integer constant)
const e = 1.0 << 3         // e == 8     (untyped integer constant)
const f = int32(1) << 33   // illegal    (constant 8589934592 overflows int32)
const g = float64(2) >> 1  // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" > "bar"    // h == true  (untyped boolean constant)
const j = true             // j == true  (untyped boolean constant)
const k = 'w' + 1          // k == 'x'   (untyped rune constant)
const l = "hi"             // l == "hi"  (untyped string constant)
const m = string(k)        // m == "x"   (type string)
const Σ = 1 - 0.707i       //            (untyped complex constant)
const Δ = Σ + 2.0e-4       //            (untyped complex constant)
const Φ = iota*1i - 1/1i   //            (untyped complex constant)
```

Applying the built-in function `complex` to untyped integer, rune, or floating-point constants yields an untyped complex constant.

将内置函数复数应用于非类型化的整数、符文或浮点常量会生成一个非类型化的复数常量。

```
const ic = complex(0, c)   // ic == 3.75i  (untyped complex constant)
const iΘ = complex(0, Θ)   // iΘ == 1i     (type complex128)
```

Constant expressions are always evaluated exactly; intermediate values and the constants themselves may require precision significantly larger than supported by any predeclared type in the language. The following are legal declarations:

常量表达式始终是精确计算的; 中间值和常量本身所需的精度可能比语言中任何预声明的类型所支持的精度大得多。以下是法律声明:

```
const Huge = 1 << 100         // Huge == 1267650600228229401496703205376  (untyped integer constant)
const Four int8 = Huge >> 98  // Four == 4                                (type int8)
```

The divisor of a constant division or remainder operation must not be zero:

常数除法或余数运算的除数不能为零:

```
3.14 / 0.0   // illegal: division by zero
```

The values of *typed* constants must always be accurately [representable](http://docscn.studygolang.com/ref/spec#Representability) by values of the constant type. The following constant expressions are illegal:

类型化常量的值必须始终能够精确地用常量类型的值表示。下面的常量表达式是非法的:

```
uint(-1)     // -1 cannot be represented as a uint
int(3.14)    // 3.14 cannot be represented as an int
int64(Huge)  // 1267650600228229401496703205376 cannot be represented as an int64
Four * 300   // operand 300 cannot be represented as an int8 (type of Four)
Four * 100   // product 400 cannot be represented as an int8 (type of Four)
```

The mask used by the unary bitwise complement operator `^` matches the rule for non-constants: the mask is all 1s for unsigned constants and -1 for signed and untyped constants.

一元位补运算符 ^ 使用的掩码与非常数规则匹配: 掩码对于无符号常数都是1，对于有符号和无类型的常数都是 -1。

```
^1         // untyped integer constant, equal to -2
uint8(^1)  // illegal: same as uint8(-2), -2 cannot be represented as a uint8
^uint8(1)  // typed uint8 constant, same as 0xFF ^ uint8(1) = uint8(0xFE)
int8(^1)   // same as int8(-2)
^int8(1)   // same as -1 ^ int8(1) = -2
```

Implementation restriction: A compiler may use rounding while computing untyped floating-point or complex constant expressions; see the implementation restriction in the section on [constants](http://docscn.studygolang.com/ref/spec#Constants). This rounding may cause a floating-point constant expression to be invalid in an integer context, even if it would be integral when calculated using infinite precision, and vice versa.

实现限制: 编译器在计算非类型化浮点表达式或复杂常量表达式时可能使用舍入; 请参阅关于常量的部分中的实现限制。这种舍入可能导致浮点常量表达式在整数上下文中无效，即使在使用无限精度计算时它是整数，反之亦然。

### Order of evaluation 评估顺序

At package level, [initialization dependencies](http://docscn.studygolang.com/ref/spec#Package_initialization) determine the evaluation order of individual initialization expressions in [variable declarations](http://docscn.studygolang.com/ref/spec#Variable_declarations). Otherwise, when evaluating the [operands](http://docscn.studygolang.com/ref/spec#Operands) of an expression, assignment, or [return statement](http://docscn.studygolang.com/ref/spec#Return_statements), all function calls, method calls, and communication operations are evaluated in lexical left-to-right order.

在包级别上，初始化依赖关系决定变量声明中各个初始化表达式的计算顺序。否则，在计算表达式、赋值或返回语句的操作数时，所有函数调用、方法调用和通信操作都按从左到右的词法顺序计算。

For example, in the (function-local) assignment

例如，在(function-local)赋值中

```
y[f()], ok = g(h(), i()+x[j()], <-c), k()
```

the function calls and communication happen in the order `f()`, `h()`, `i()`, `j()`, `<-c`, `g()`, and `k()`. However, the order of those events compared to the evaluation and indexing of `x` and the evaluation of `y` is not specified.

函数调用和通信按 f ()、 h ()、 i ()、 j ()、 <-c、 g ()和 k ()的顺序进行。但是，与 x 的计算和索引以及 y 的计算相比，这些事件的顺序并未指定。

```
a := 1
f := func() int { a++; return a }
x := []int{a, f()}            // x may be [1, 2] or [2, 2]: evaluation order between a and f() is not specified
m := map[int]int{a: 1, a: 2}  // m may be {2: 1} or {2: 2}: evaluation order between the two map assignments is not specified
n := map[int]int{a: f()}      // n may be {2: 3} or {3: 3}: evaluation order between the key and the value is not specified
```

At package level, initialization dependencies override the left-to-right rule for individual initialization expressions, but not for operands within each expression:

在包级别，对于单个初始化表达式，初始化依赖关系覆盖从左到右的规则，但是对于每个表达式中的操作数则不是这样:

```
var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int        { return c }
func g() int        { return a }
func sqr(x int) int { return x*x }

// functions u and v are independent of all other variables and functions
```

The function calls happen in the order `u()`, `sqr()`, `v()`, `f()`, `v()`, and `g()`.

函数调用按 u ()、 sqr ()、 v ()、 f ()、 v ()和 g ()的顺序进行。

Floating-point operations within a single expression are evaluated according to the associativity of the operators. Explicit parentheses affect the evaluation by overriding the default associativity. In the expression `x + (y + z)` the addition `y + z` is performed before adding `x`.

单个表达式中的浮点运算是根据运算符的结合性来计算的。显式括号通过覆盖默认关联性来影响计算。在表达式 x + (y + z)中，加法 y + z 是在加法 x 之前进行的。

## Statements 声明

Statements control execution.

语句控制执行。

```
Statement =
	Declaration | LabeledStmt | SimpleStmt |
	GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
	FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
	DeferStmt .

SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
```

### Terminating statements 终止语句

A *terminating statement* prevents execution of all statements that lexically appear after it in the same [block](http://docscn.studygolang.com/ref/spec#Blocks). The following statements are terminating:

终止语句阻止执行在同一块中以词法形式出现在它之后的所有语句。下列声明终止:

1. A

    

   "return" 「归还」

    

   or

    

   或

   "goto" 后藤

    

   statement.

    

   语句

2. A call to the built-in function

    

   对内置函数的调用

   `panic`

   .

3. A

    

   block 块

    

   in which the statement list ends in a terminating statement.

    

   其中语句列表以终止语句结束

4. An

    

   一个

   "if" statement “如果”语句

    

   in which:

    

   其中:

   - the "else" branch is present, and “ else”分支已经存在
   - both branches are terminating statements. 两个分支都在终止语句

5. A

    

   "for" statement “ for”语句

    

   in which:

    

   其中:

   - there are no "break" statements referring to the "for" statement, and 没有提及“ for”语句的“ break”语句，以及
   - the loop condition is absent. 没有循环条件

6. A

    

   "switch" statement “ switch”语句

    

   in which:

    

   其中:

   - there are no "break" statements referring to the "switch" statement, 没有提到“ switch”语句的“ break”语句,
   - there is a default case, and 存在缺省情况
   - the statement lists in each case, including the default, end in a terminating statement, or a possibly labeled 每种情况下的语句列表，包括缺省语句、以终止语句结尾的语句或可能标记的语句["fallthrough" statement “ fallthrough”语句](http://docscn.studygolang.com/ref/spec#Fallthrough_statements).

7. A

    

   "select" statement ”选择”语句

    

   in which:

    

   其中:

   - there are no "break" statements referring to the "select" statement, and 没有引用“ select”语句的“ break”语句，以及
   - the statement lists in each case, including the default if present, end in a terminating statement. 每种情况下的语句列表，包括缺省的 if present，以终止语句结束

8. A [labeled statement 标记语句](http://docscn.studygolang.com/ref/spec#Labeled_statements) labeling a terminating statement. 标记终止语句

All other statements are not terminating.

所有其他语句都不终止。

A [statement list](http://docscn.studygolang.com/ref/spec#Blocks) ends in a terminating statement if the list is not empty and its final non-empty statement is terminating.

如果一个语句列表不是空的，并且它的最后一个非空语句正在终止，那么该语句列表将以终止语句结束。

### Empty statements 空洞的陈述

The empty statement does nothing.

空语句什么也不做。

```
EmptyStmt = .
```

### Labeled statements 标记语句

A labeled statement may be the target of a `goto`, `break` or `continue` statement.

带标签的语句可能是 goto、 break 或 continue 语句的目标。

```
LabeledStmt = Label ":" Statement .
Label       = identifier .
Error: log.Panic("error encountered")
```

### Expression statements 表达式语句

With the exception of specific built-in functions, function and method [calls](http://docscn.studygolang.com/ref/spec#Calls) and [receive operations](http://docscn.studygolang.com/ref/spec#Receive_operator) can appear in statement context. Such statements may be parenthesized.

除了特定的内置函数之外，函数、方法调用和接收操作都可以出现在语句上下文中。这些语句可以用括号括起来。

```
ExpressionStmt = Expression .
```

The following built-in functions are not permitted in statement context:

在语句上下文中不允许使用下列内置函数:

```
append cap complex imag len make new real
unsafe.Alignof unsafe.Offsetof unsafe.Sizeof
h(x+y)
f.Close()
<-ch
(<-ch)
len("foo")  // illegal if len is the built-in function
```

### Send statements 发送声明

A send statement sends a value on a channel. The channel expression must be of [channel type](http://docscn.studygolang.com/ref/spec#Channel_types), the channel direction must permit send operations, and the type of the value to be sent must be [assignable](http://docscn.studygolang.com/ref/spec#Assignability) to the channel's element type.

Send 语句在通道上发送值。通道表达式必须是通道类型，通道方向必须允许发送操作，要发送的值的类型必须可分配给通道的元素类型。

```
SendStmt = Channel "<-" Expression .
Channel  = Expression .
```

Both the channel and the value expression are evaluated before communication begins. Communication blocks until the send can proceed. A send on an unbuffered channel can proceed if a receiver is ready. A send on a buffered channel can proceed if there is room in the buffer. A send on a closed channel proceeds by causing a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics). A send on a `nil` channel blocks forever.

通道和值表达式都在通信开始之前进行计算。通信阻塞，直到发送可以继续。如果接收机准备就绪，在未缓冲信道上的发送可以继续。如果缓冲区中有空间，则可以继续发送缓冲通道上的发送。封闭通道上的发送会引起运行时的恐慌。一个发送无信道永远阻塞。

```
ch <- 3  // send value 3 to channel ch
```

### IncDec statements IncDec 声明

The "++" and "--" statements increment or decrement their operands by the untyped [constant](http://docscn.studygolang.com/ref/spec#Constants) `1`. As with an assignment, the operand must be [addressable](http://docscn.studygolang.com/ref/spec#Address_operators) or a map index expression.

“ + +”和“ --”语句将其操作数递增或递减为非类型化常量1。与赋值一样，操作数必须是可寻址的或映射索引表达式。

```
IncDecStmt = Expression ( "++" | "--" ) .
```

The following [assignment statements](http://docscn.studygolang.com/ref/spec#Assignments) are semantically equivalent:

下面的赋值语句在语义上是等价的:

```
IncDec statement    Assignment
x++                 x += 1
x--                 x -= 1
```

### Assignments 工作分配

```
Assignment = ExpressionList assign_op ExpressionList .

assign_op = [ add_op | mul_op ] "=" .
```

Each left-hand side operand must be [addressable](http://docscn.studygolang.com/ref/spec#Address_operators), a map index expression, or (for `=` assignments only) the [blank identifier](http://docscn.studygolang.com/ref/spec#Blank_identifier). Operands may be parenthesized.

每个左侧操作数必须是可寻址的、映射索引表达式或(仅用于 = assignment)空白标识符。操作数可以是括号。

```
x = 1
*p = f()
a[i] = 23
(k) = <-ch  // same as: k = <-ch
```

An *assignment operation* `x` *op*`=` `y` where *op* is a binary [arithmetic operator](http://docscn.studygolang.com/ref/spec#Arithmetic_operators) is equivalent to `x` `=` `x` *op* `(y)` but evaluates `x` only once. The *op*`=` construct is a single token. In assignment operations, both the left- and right-hand expression lists must contain exactly one single-valued expression, and the left-hand expression must not be the blank identifier.

一个赋值运算 x op = y，其中 op 是一个二元算术运算符，这等价于 x = x op (y) ，但只计算一次 x。Op = 构造是一个单一的令牌。在赋值操作中，左表达式列表和右表达式列表必须恰好包含一个单值表达式，左表达式不能是空白标识符。

```
a[i] <<= 2
i &^= 1<<n
```

A tuple assignment assigns the individual elements of a multi-valued operation to a list of variables. There are two forms. In the first, the right hand operand is a single multi-valued expression such as a function call, a [channel](http://docscn.studygolang.com/ref/spec#Channel_types) or [map](http://docscn.studygolang.com/ref/spec#Map_types) operation, or a [type assertion](http://docscn.studygolang.com/ref/spec#Type_assertions). The number of operands on the left hand side must match the number of values. For instance, if `f` is a function returning two values,

元组赋值将多值操作的单个元素赋给变量列表。有两种形式。在第一种方法中，右操作数是单个多值表达式，如函数调用、通道或映射操作或类型断言。左边的操作数数目必须与值的数目相匹配。例如，如果 f 是一个返回两个值的函数,

```
x, y = f()
```

assigns the first value to `x` and the second to `y`. In the second form, the number of operands on the left must equal the number of expressions on the right, each of which must be single-valued, and the *n*th expression on the right is assigned to the *n*th operand on the left:

将第一个值赋给 x，第二个赋给 y。在第二种形式中，左边的操作数数目必须等于右边的表达式数目，每个表达式必须是单值的，右边的第 n 个表达式分配给左边的第 n 个操作数:

```
one, two, three = '一', '二', '三'
```

The [blank identifier](http://docscn.studygolang.com/ref/spec#Blank_identifier) provides a way to ignore right-hand side values in an assignment:

空白标识符提供了一种在赋值中忽略右边值的方法:

```
_ = x       // evaluate x but ignore it
x, _ = f()  // evaluate f() but ignore second result value
```

The assignment proceeds in two phases. First, the operands of [index expressions](http://docscn.studygolang.com/ref/spec#Index_expressions) and [pointer indirections](http://docscn.studygolang.com/ref/spec#Address_operators) (including implicit pointer indirections in [selectors](http://docscn.studygolang.com/ref/spec#Selectors)) on the left and the expressions on the right are all [evaluated in the usual order](http://docscn.studygolang.com/ref/spec#Order_of_evaluation). Second, the assignments are carried out in left-to-right order.

分配工作分两个阶段进行。首先，左侧的索引表达式和指针间接方向(包括选择器中的隐式指针间接方向)的操作数和右侧的表达式都按照通常的顺序计算。其次，作业是按照从左到右的顺序进行的。

```
a, b = b, a  // exchange a and b

x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2  // set i = 1, x[0] = 2

i = 0
x[i], i = 2, 1  // set x[0] = 2, i = 1

x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end)

x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5.

type Point struct { x, y int }
var p *Point
x[2], p.x = 6, 7  // set x[2] = 6, then panic setting p.x = 7

i = 2
x = []int{3, 5, 7}
for i, x[i] = range x {  // set i, x[2] = 0, x[0]
	break
}
// after this loop, i == 0 and x == []int{3, 5, 3}
```

In assignments, each value must be [assignable](http://docscn.studygolang.com/ref/spec#Assignability) to the type of the operand to which it is assigned, with the following special cases:

在转让中，每个值都必须可以分配给指派给它的操作数的类型，特殊情况如下:

1. Any typed value may be assigned to the blank identifier. 任何类型化的值都可以分配给空白标识符
2. If an untyped constant is assigned to a variable of interface type or the blank identifier, the constant is first implicitly 如果将非类型化常量分配给接口类型的变量或空白标识符，则该常量首先隐式[converted 转换](http://docscn.studygolang.com/ref/spec#Conversions) to its 到它的[default type 默认类型](http://docscn.studygolang.com/ref/spec#Constants).
3. If an untyped boolean value is assigned to a variable of interface type or the blank identifier, it is first implicitly converted to type 如果将非类型化布尔值分配给接口类型的变量或空白标识符，则首先将其隐式转换为类型`bool`.

### If statements 如果语句

"If" statements specify the conditional execution of two branches according to the value of a boolean expression. If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.

“ If”语句根据布尔表达式的值指定两个分支的条件执行。如果表达式的计算结果为 true，则执行“ If”分支，否则，如果存在，则执行“ else”分支。

```
IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .
if x > max {
	x = max
}
```

The expression may be preceded by a simple statement, which executes before the expression is evaluated.

表达式之前可以有一个简单的语句，该语句在表达式求值之前执行。

```
if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}
```

### Switch statements 开关语句

"Switch" statements provide multi-way execution. An expression or type specifier is compared to the "cases" inside the "switch" to determine which branch to execute.

“ Switch”语句提供多路执行。表达式或类型说明符与“开关”中的“ cases”进行比较，以确定要执行哪个分支。

```
SwitchStmt = ExprSwitchStmt | TypeSwitchStmt .
```

There are two forms: expression switches and type switches. In an expression switch, the cases contain expressions that are compared against the value of the switch expression. In a type switch, the cases contain types that are compared against the type of a specially annotated switch expression. The switch expression is evaluated exactly once in a switch statement.

有两种形式: 表达式开关和类型开关。在表达式开关中，案例包含与开关表达式的值进行比较的表达式。在类型开关中，案例包含的类型与专门注释的开关表达式的类型进行比较。在 switch 语句中，开关表达式只计算一次。

#### Expression switches

#### 表达式开关

In an expression switch, the switch expression is evaluated and the case expressions, which need not be constants, are evaluated left-to-right and top-to-bottom; the first one that equals the switch expression triggers execution of the statements of the associated case; the other cases are skipped. If no case matches and there is a "default" case, its statements are executed. There can be at most one default case and it may appear anywhere in the "switch" statement. A missing switch expression is equivalent to the boolean value `true`.

在表达式开关中，对 switch 表达式进行求值，并且对不需要为常量的 case 表达式从左到右和从上到下进行求值; 与 switch 表达式相等的第一个 case 触发相关 case 语句的执行; 跳过其他 case。如果没有大小写匹配，并且存在“缺省”大小写，则执行其语句。最多只有一个默认情况，它可能出现在“ switch”语句的任何地方。缺少的开关表达式等价于布尔值 true。

```
ExprSwitchStmt = "switch" [ SimpleStmt ";" ] [ Expression ] "{" { ExprCaseClause } "}" .
ExprCaseClause = ExprSwitchCase ":" StatementList .
ExprSwitchCase = "case" ExpressionList | "default" .
```

If the switch expression evaluates to an untyped constant, it is first implicitly [converted](http://docscn.studygolang.com/ref/spec#Conversions) to its [default type](http://docscn.studygolang.com/ref/spec#Constants); if it is an untyped boolean value, it is first implicitly converted to type `bool`. The predeclared untyped value `nil` cannot be used as a switch expression.

如果 switch 表达式的计算结果为非类型化常量，则首先隐式转换为其默认类型; 如果是非类型化布尔值，则首先隐式转换为 bool 类型。预先声明的非类型化值 nil 不能用作开关表达式。

If a case expression is untyped, it is first implicitly [converted](http://docscn.studygolang.com/ref/spec#Conversions) to the type of the switch expression. For each (possibly converted) case expression `x` and the value `t` of the switch expression, `x == t` must be a valid [comparison](http://docscn.studygolang.com/ref/spec#Comparison_operators).

如果大小写表达式是非类型化的，则首先将其隐式转换为开关表达式的类型。对于每个(可能转换的)大小写表达式 x 和开关表达式的值 t，x = = t 必须是一个有效的比较。

In other words, the switch expression is treated as if it were used to declare and initialize a temporary variable `t` without explicit type; it is that value of `t` against which each case expression `x` is tested for equality.

换句话说，开关表达式被当作是用来声明和初始化临时变量 t，而不是显式类型; 它是 t 的值，每个案例表达式 x 是否相等。

In a case or default clause, the last non-empty statement may be a (possibly [labeled](http://docscn.studygolang.com/ref/spec#Labeled_statements)) ["fallthrough" statement](http://docscn.studygolang.com/ref/spec#Fallthrough_statements) to indicate that control should flow from the end of this clause to the first statement of the next clause. Otherwise control flows to the end of the "switch" statement. A "fallthrough" statement may appear as the last statement of all but the last clause of an expression switch.

在 case 或 default 子句中，最后一个非空语句可能是(可能标记为)“ fallthrough”语句，以指示控制应该从该子句的结尾流向下一个子句的第一个语句。否则，控制流将流到“ switch”语句的末尾。“ fallthrough”语句可能作为除表达式切换的最后一个子句以外的所有语句的最后一个语句出现。

The switch expression may be preceded by a simple statement, which executes before the expression is evaluated.

在开关表达式之前可以执行一个简单的语句，该语句在表达式求值之前执行。

```
switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {  // missing switch expression means "true"
case x < 0: return -x
default: return x
}

switch {
case x < y: f1()
case x < z: f2()
case x == 4: f3()
}
```

Implementation restriction: A compiler may disallow multiple case expressions evaluating to the same constant. For instance, the current compilers disallow duplicate integer, floating point, or string constants in case expressions.

实现限制: 编译器可能禁止计算到同一常量的多个大小写表达式。例如，当前编译器禁止大小写表达式中的重复整数、浮点数或字符串常量。

#### Type switches

#### 类型开关

A type switch compares types rather than values. It is otherwise similar to an expression switch. It is marked by a special switch expression that has the form of a [type assertion](http://docscn.studygolang.com/ref/spec#Type_assertions) using the reserved word `type` rather than an actual type:

类型开关比较的是类型而不是值。它在其他方面类似于表达式开关。它由一个特殊的 switch 表达式标记，该表达式的形式是使用保留字类型而不是实际类型的类型断言:

```
switch x.(type) {
// cases
}
```

Cases then match actual types `T` against the dynamic type of the expression `x`. As with type assertions, `x` must be of [interface type](http://docscn.studygolang.com/ref/spec#Interface_types), and each non-interface type `T` listed in a case must implement the type of `x`. The types listed in the cases of a type switch must all be [different](http://docscn.studygolang.com/ref/spec#Type_identity).

然后，用例将实际类型 t 与表达式 x 的动态类型匹配。与类型断言一样，x 必须是接口类型，并且案例中列出的每个非接口类型 t 都必须实现 x 的类型。类型开关的例子中列出的类型必须都是不同的。

```
TypeSwitchStmt  = "switch" [ SimpleStmt ";" ] TypeSwitchGuard "{" { TypeCaseClause } "}" .
TypeSwitchGuard = [ identifier ":=" ] PrimaryExpr "." "(" "type" ")" .
TypeCaseClause  = TypeSwitchCase ":" StatementList .
TypeSwitchCase  = "case" TypeList | "default" .
TypeList        = Type { "," Type } .
```

The TypeSwitchGuard may include a [short variable declaration](http://docscn.studygolang.com/ref/spec#Short_variable_declarations). When that form is used, the variable is declared at the end of the TypeSwitchCase in the [implicit block](http://docscn.studygolang.com/ref/spec#Blocks) of each clause. In clauses with a case listing exactly one type, the variable has that type; otherwise, the variable has the type of the expression in the TypeSwitchGuard.

TypeSwitchGuard 可能包括一个简短的变量声明。当使用该表单时，变量在每个子句的隐式块中的 TypeSwitchCase 结尾声明。如果子句中的 case 只列出一种类型，则变量具有该类型; 否则，变量具有 TypeSwitchGuard 中表达式的类型。

Instead of a type, a case may use the predeclared identifier [`nil`](http://docscn.studygolang.com/ref/spec#Predeclared_identifiers); that case is selected when the expression in the TypeSwitchGuard is a `nil` interface value. There may be at most one `nil` case.

Case 可以使用预声明的标识符 nil，而不是类型; 当 TypeSwitchGuard 中的表达式为 nil 接口值时，就会选择这个 case。最多可能有一个零案件。

Given an expression `x` of type `interface{}`, the following type switch:

给定一个接口{}类型的表达式 x，下面的类型切换:

```
switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
```

could be rewritten:

可以重写:

```
v := x  // x is evaluated exactly once
if v == nil {
	i := v                                 // type of i is type of x (interface{})
	printString("x is nil")
} else if i, isInt := v.(int); isInt {
	printInt(i)                            // type of i is int
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i)                        // type of i is float64
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i)                       // type of i is func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v                         // type of i is type of x (interface{})
		printString("type is bool or string")
	} else {
		i := v                         // type of i is type of x (interface{})
		printString("don't know the type")
	}
}
```

The type switch guard may be preceded by a simple statement, which executes before the guard is evaluated.

在类型开关保护之前可以执行一个简单的语句，该语句在评估保护之前执行。

The "fallthrough" statement is not permitted in a type switch.

类型开关中不允许使用“ fallthrough”语句。

### For statements 对于声明

A "for" statement specifies repeated execution of a block. There are three forms: The iteration may be controlled by a single condition, a "for" clause, or a "range" clause.

“ for”语句指定重复执行一个块。有三种形式: 迭代可以由单个条件、“ for”子句或“ range”子句控制。

```
ForStmt = "for" [ Condition | ForClause | RangeClause ] Block .
Condition = Expression .
```

#### For statements with single condition

#### 对于具有单个条件的语句

In its simplest form, a "for" statement specifies the repeated execution of a block as long as a boolean condition evaluates to true. The condition is evaluated before each iteration. If the condition is absent, it is equivalent to the boolean value `true`.

在其最简单的形式中，“ for”语句指定一个块的重复执行，只要布尔条件的计算结果为 true。条件在每次迭代之前进行评估。如果条件不存在，则等效于布尔值 true。

```
for a < b {
	a *= 2
}
```

#### For statements with `for` clause

#### 用于带 For 子句的语句

A "for" statement with a ForClause is also controlled by its condition, but additionally it may specify an *init* and a *post* statement, such as an assignment, an increment or decrement statement. The init statement may be a [short variable declaration](http://docscn.studygolang.com/ref/spec#Short_variable_declarations), but the post statement must not. Variables declared by the init statement are re-used in each iteration.

带有 ForClause 的“ for”语句也受其条件控制，但是它还可以指定 init 和 post 语句，如赋值、递增或递减语句。Init 语句可以是简短的变量声明，但 post 语句不能。Init 语句声明的变量在每次迭代中重用。

```
ForClause = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
InitStmt = SimpleStmt .
PostStmt = SimpleStmt .
for i := 0; i < 10; i++ {
	f(i)
}
```

If non-empty, the init statement is executed once before evaluating the condition for the first iteration; the post statement is executed after each execution of the block (and only if the block was executed). Any element of the ForClause may be empty but the [semicolons](http://docscn.studygolang.com/ref/spec#Semicolons) are required unless there is only a condition. If the condition is absent, it is equivalent to the boolean value `true`.

如果非空，init 语句将在计算第一次迭代的条件之前执行一次; post 语句将在每次执行代码块之后执行(并且只在执行代码块时执行)。ForClause 中的任何元素都可以是空的，但是分号是必需的，除非只有一个条件。如果条件不存在，则等效于布尔值 true。

```
for cond { S() }    is the same as    for ; cond ; { S() }
for      { S() }    is the same as    for true     { S() }
```

#### For statements with `range` clause

#### 对于带 range 子句的语句

A "for" statement with a "range" clause iterates through all entries of an array, slice, string or map, or values received on a channel. For each entry it assigns *iteration values* to corresponding *iteration variables* if present and then executes the block.

带有“ range”子句的“ for”语句迭代通道上接收的数组、片、字符串或映射的所有条目。对于每个条目，它为相应的迭代变量赋值迭代值，如果存在，然后执行该块。

```
RangeClause = [ ExpressionList "=" | IdentifierList ":=" ] "range" Expression .
```

The expression on the right in the "range" clause is called the *range expression*, which may be an array, pointer to an array, slice, string, map, or channel permitting [receive operations](http://docscn.studygolang.com/ref/spec#Receive_operator). As with an assignment, if present the operands on the left must be [addressable](http://docscn.studygolang.com/ref/spec#Address_operators) or map index expressions; they denote the iteration variables. If the range expression is a channel, at most one iteration variable is permitted, otherwise there may be up to two. If the last iteration variable is the [blank identifier](http://docscn.studygolang.com/ref/spec#Blank_identifier), the range clause is equivalent to the same clause without that identifier.

“ range”子句右边的表达式称为 range expression，它可以是数组、指向数组的指针、 slice、 string、 map 或允许接收操作的 channel。与赋值一样，如果左边的操作数是可寻址的或映射索引表达式; 它们表示迭代变量。如果范围表达式是通道，则最多允许一个迭代变量，否则最多可能有两个。如果最后一个迭代变量是空标识符，那么 range 子句等效于没有该标识符的同一个子句。

The range expression `x` is evaluated once before beginning the loop, with one exception: if at most one iteration variable is present and `len(x)` is [constant](http://docscn.studygolang.com/ref/spec#Length_and_capacity), the range expression is not evaluated.

范围表达式 x 在开始循环之前计算一次，但有一个例外: 如果最多只有一个迭代变量，且 len (x)为常数，则不计算范围表达式。

Function calls on the left are evaluated once per iteration. For each iteration, iteration values are produced as follows if the respective iteration variables are present:

左边的函数调用每次迭代计算一次。对于每个迭代，如果存在相应的迭代变量，迭代值将产生如下:

```
Range expression                          1st value          2nd value

array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
string          s  string type            index    i  int    see below  rune
map             m  map[K]V                key      k  K      m[k]       V
channel         c  chan E, <-chan E       element  e  E
```

1. For an array, pointer to array, or slice value 对于数组、指向数组的指针或片值`a`, the index iteration values are produced in increasing order, starting at element index 0. If at most one iteration variable is present, the range loop produces iteration values from 0 up to 在此基础上，以元素索引0为起始点，按递增顺序生成索引迭代值。如果最多只有一个迭代变量，那么范围循环产生的迭代值从0到`len(a)-1` and does not index into the array or slice itself. For a 并且不会索引到数组中或者切片本身`nil` slice, the number of iterations is 0. 切片，迭代次数为0
2. For a string value, the "range" clause iterates over the Unicode code points in the string starting at byte index 0. On successive iterations, the index value will be the index of the first byte of successive UTF-8-encoded code points in the string, and the second value, of type 对于字符串值，“ range”子句在 Unicode 代码上迭代，从字节索引0开始。在连续的迭代中，索引值将是字符串中连续 utf-8编码的代码点的第一个字节的索引，以及类型的第二个值`rune`, will be the value of the corresponding code point. If the iteration encounters an invalid UTF-8 sequence, the second value will be ，就是相应的码点的值。如果迭代遇到无效的 UTF-8序列，则第二个值为`0xFFFD`, the Unicode replacement character, and the next iteration will advance a single byte in the string. ，Unicode 替换字符，下一次迭代将在字符串中前进一个字节
3. The iteration order over maps is not specified and is not guaranteed to be the same from one iteration to the next. If a map entry that has not yet been reached is removed during iteration, the corresponding iteration value will not be produced. If a map entry is created during iteration, that entry may be produced during the iteration or may be skipped. The choice may vary for each entry created and from one iteration to the next. If the map is 没有指定映射上的迭代顺序，也不能保证从一个迭代到下一个迭代是相同的。如果在迭代过程中删除了尚未到达的映射条目，则不会生成相应的迭代值。如果在迭代期间创建了映射条目，那么可以在迭代期间生成该条目，或者跳过该条目。对于创建的每个条目以及每次迭代的选择可能会有所不同。如果地图是`nil`, the number of iterations is 0. ，迭代次数为0
4. For channels, the iteration values produced are the successive values sent on the channel until the channel is 对于通道，生成的迭代值是在通道上发送的连续值，直到通道为[closed 关闭](http://docscn.studygolang.com/ref/spec#Close). If the channel is 。如果频道是`nil`, the range expression blocks forever. 范围表达式永远阻塞

The iteration values are assigned to the respective iteration variables as in an [assignment statement](http://docscn.studygolang.com/ref/spec#Assignments).

迭代值分配给相应的迭代变量，如赋值语句中所示。

The iteration variables may be declared by the "range" clause using a form of [short variable declaration](http://docscn.studygolang.com/ref/spec#Short_variable_declarations) (`:=`). In this case their types are set to the types of the respective iteration values and their [scope](http://docscn.studygolang.com/ref/spec#Declarations_and_scope) is the block of the "for" statement; they are re-used in each iteration. If the iteration variables are declared outside the "for" statement, after execution their values will be those of the last iteration.

迭代变量可以由“ range”子句使用短变量声明(: =)的形式声明。在这种情况下，它们的类型被设置为各自迭代值的类型，它们的范围是“ for”语句的块; 它们在每次迭代中被重用。如果迭代变量在“ for”语句之外声明，那么在执行之后，它们的值将是最后一次迭代的值。

```
var testdata *struct {
	a *[7]int
}
for i, _ := range testdata.a {
	// testdata.a is never evaluated; len(testdata.a) is constant
	// i ranges from 0 to 6
	f(i)
}

var a [10]string
for i, s := range a {
	// type of i is int
	// type of s is string
	// s == a[i]
	g(i, s)
}

var key string
var val interface{}  // element type of m is assignable to val
m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
for key, val = range m {
	h(key, val)
}
// key == last map key encountered in iteration
// val == map[key]

var ch chan Work = producer()
for w := range ch {
	doWork(w)
}

// empty a channel
for range ch {}
```

### Go statements 发布声明

A "go" statement starts the execution of a function call as an independent concurrent thread of control, or *goroutine*, within the same address space.

“ go”语句作为独立的并发控制线程(或 goroutine)在同一地址空间内启动函数调用的执行。

```
GoStmt = "go" Expression .
```

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for [expression statements](http://docscn.studygolang.com/ref/spec#Expression_statements).

表达式必须是函数或方法调用; 它不能被括号括起来。内置函数的调用与表达式语句一样受到限制。

The function value and parameters are [evaluated as usual](http://docscn.studygolang.com/ref/spec#Calls) in the calling goroutine, but unlike with a regular call, program execution does not wait for the invoked function to complete. Instead, the function begins executing independently in a new goroutine. When the function terminates, its goroutine also terminates. If the function has any return values, they are discarded when the function completes.

函数值和参数通常在调用 goroutine 中计算，但与常规调用不同的是，程序执行不等待被调用函数完成。相反，该函数开始在新 goroutine 中独立执行。当函数终止时，它的 goroutine 也终止。如果函数具有任何返回值，则在函数完成时丢弃这些返回值。

```
go Server()
go func(ch chan<- bool) { for { sleep(10); ch <- true }} (c)
```

### Select statements 选择语句

A "select" statement chooses which of a set of possible [send](http://docscn.studygolang.com/ref/spec#Send_statements) or [receive](http://docscn.studygolang.com/ref/spec#Receive_operator) operations will proceed. It looks similar to a ["switch"](http://docscn.studygolang.com/ref/spec#Switch_statements) statement but with the cases all referring to communication operations.

“ select”语句选择一组可能的发送或接收操作中的哪一个将继续进行。它看起来类似于“ switch”语句，但案例都涉及到通信操作。

```
SelectStmt = "select" "{" { CommClause } "}" .
CommClause = CommCase ":" StatementList .
CommCase   = "case" ( SendStmt | RecvStmt ) | "default" .
RecvStmt   = [ ExpressionList "=" | IdentifierList ":=" ] RecvExpr .
RecvExpr   = Expression .
```

A case with a RecvStmt may assign the result of a RecvExpr to one or two variables, which may be declared using a [short variable declaration](http://docscn.studygolang.com/ref/spec#Short_variable_declarations). The RecvExpr must be a (possibly parenthesized) receive operation. There can be at most one default case and it may appear anywhere in the list of cases.

使用 RecvStmt 的情况可以将 RecvExpr 的结果分配给一个或两个变量，这些变量可以使用短变量声明声明。RecvExpr 必须是一个(可能是括号中的)接收操作。最多可以有一个缺省情况，它可能出现在情况列表的任何地方。

Execution of a "select" statement proceeds in several steps:

“ select”语句的执行分为几个步骤:

1. For all the cases in the statement, the channel operands of receive operations and the channel and right-hand-side expressions of send statements are evaluated exactly once, in source order, upon entering the "select" statement. The result is a set of channels to receive from or send to, and the corresponding values to send. Any side effects in that evaluation will occur irrespective of which (if any) communication operation is selected to proceed. Expressions on the left-hand side of a RecvStmt with a short variable declaration or assignment are not yet evaluated. 对于语句中的所有情况，在输入“ select”语句时，接收操作的通道操作数以及发送语句的通道和右侧表达式将按源顺序精确计算一次。结果是一组要从中接收或发送的通道，以及相应的要发送的值。无论选择哪个(如果有的话)通信操作进行，该评估中的任何副作用都将发生。具有短变量声明或赋值的 RecvStmt 左侧的表达式尚未计算
2. If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection. Otherwise, if there is a default case, that case is chosen. If there is no default case, the "select" statement blocks until at least one of the communications can proceed. 如果可以进行一个或多个通信，则通过均匀的伪随机选择选择可以进行的单个通信。否则，如果存在缺省情况，则选择该情况。如果没有默认情况，“ select”语句将阻塞，直到至少有一个通信可以继续
3. Unless the selected case is the default case, the respective communication operation is executed. 除非所选的大小写是缺省大小写，否则将执行相应的通信操作
4. If the selected case is a RecvStmt with a short variable declaration or an assignment, the left-hand side expressions are evaluated and the received value (or values) are assigned. 如果选中的大小写是一个带有短变量声明或赋值的 RecvStmt，则计算左侧表达式并赋值接收到的值(或值)
5. The statement list of the selected case is executed. 执行所选案例的语句列表

Since communication on `nil` channels can never proceed, a select with only `nil` channels and no default case blocks forever.

由于无通道上的通信永远无法进行，选择只有无通道和没有默认情况下永远阻塞。

```
var a []int
var c, c1, c2, c3, c4 chan int
var i1, i2 int
select {
case i1 = <-c1:
	print("received ", i1, " from c1\n")
case c2 <- i2:
	print("sent ", i2, " to c2\n")
case i3, ok := (<-c3):  // same as: i3, ok := <-c3
	if ok {
		print("received ", i3, " from c3\n")
	} else {
		print("c3 is closed\n")
	}
case a[f()] = <-c4:
	// same as:
	// case t := <-c4
	//	a[f()] = t
default:
	print("no communication\n")
}

for {  // send random sequence of bits to c
	select {
	case c <- 0:  // note: no statement, no fallthrough, no folding of cases
	case c <- 1:
	}
}

select {}  // block forever
```

### Return statements 返回语句

A "return" statement in a function `F` terminates the execution of `F`, and optionally provides one or more result values. Any functions [deferred](http://docscn.studygolang.com/ref/spec#Defer_statements) by `F` are executed before `F` returns to its caller.

函数 f 中的“ return”语句终止 f 的执行，并可选地提供一个或多个结果值。F 延迟的任何函数在 f 返回给它的调用者之前都会被执行。

```
ReturnStmt = "return" [ ExpressionList ] .
```

In a function without a result type, a "return" statement must not specify any result values.

在没有结果类型的函数中，“ return”语句不能指定任何结果值。

```
func noResult() {
	return
}
```

There are three ways to return values from a function with a result type:

有三种方法可以从结果类型的函数返回值:

1. The return value or values may be explicitly listed in the "return" statement. Each expression must be single-valued and

    

   返回值可以在“ return”语句中显式列出。每个表达式必须是单值和

   assignable 可分配的

    

   to the corresponding element of the function's result type.

    

   到函数结果类型的对应元素

   ```
   func simpleF() int {
   	return 2
   }
   
   func complexF1() (re float64, im float64) {
   	return -7.0, -4.0
   }
   ```

2. The expression list in the "return" statement may be a single call to a multi-valued function. The effect is as if each value returned from that function were assigned to a temporary variable with the type of the respective value, followed by a "return" statement listing these variables, at which point the rules of the previous case apply.

    

   “ return”语句中的表达式列表可以是对多值函数的单个调用。其效果就好像是将该函数返回的每个值分配给一个临时变量，该临时变量的类型是相应的值，后面跟着一个列出这些变量的“ return”语句，此时适用前一种情况的规则

   ```
   func complexF2() (re float64, im float64) {
   	return complexF1()
   }
   ```

3. The expression list may be empty if the function's result type specifies names for its

    

   如果函数的结果类型为其指定了名称，则表达式列表可能为空

   result parameters 结果参数

   . The result parameters act as ordinary local variables and the function may assign values to them as necessary. The "return" statement returns the values of these variables.

    

   .结果参数作为普通的局部变量，函数可以根据需要为它们赋值。返回语句返回这些变量的值

   ```
   func complexF3() (re float64, im float64) {
   	re = 7.0
   	im = 4.0
   	return
   }
   
   func (devnull) Write(p []byte) (n int, _ error) {
   	n = len(p)
   	return
   }
   ```

Regardless of how they are declared, all the result values are initialized to the [zero values](http://docscn.studygolang.com/ref/spec#The_zero_value) for their type upon entry to the function. A "return" statement that specifies results sets the result parameters before any deferred functions are executed.

无论如何声明它们，在进入函数时，所有结果值都被初始化为其类型的零值。在执行任何延迟函数之前，指定结果的“ return”语句将设置结果参数。

Implementation restriction: A compiler may disallow an empty expression list in a "return" statement if a different entity (constant, type, or variable) with the same name as a result parameter is in [scope](http://docscn.studygolang.com/ref/spec#Declarations_and_scope) at the place of the return.

实现限制: 如果返回位置的范围中有与结果参数同名的不同实体(常量、类型或变量) ，编译器可能会禁用“ return”语句中的空表达式列表。

```
func f(n int) (res int, err error) {
	if _, err := f(n-1); err != nil {
		return  // invalid return statement: err is shadowed
	}
	return
}
```

### Break statements 打破陈述

A "break" statement terminates execution of the innermost ["for"](http://docscn.studygolang.com/ref/spec#For_statements), ["switch"](http://docscn.studygolang.com/ref/spec#Switch_statements), or ["select"](http://docscn.studygolang.com/ref/spec#Select_statements) statement within the same function.

“ break”语句终止同一函数中最里面的“ for”、“ switch”或“ select”语句的执行。

```
BreakStmt = "break" [ Label ] .
```

If there is a label, it must be that of an enclosing "for", "switch", or "select" statement, and that is the one whose execution terminates.

如果有一个标签，它必须是一个封闭的“ for”、“ switch”或“ select”语句的标签，而这个标签的执行将终止。

```
OuterLoop:
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			switch a[i][j] {
			case nil:
				state = Error
				break OuterLoop
			case item:
				state = Found
				break OuterLoop
			}
		}
	}
```

### Continue statements 继续陈述

A "continue" statement begins the next iteration of the innermost ["for" loop](http://docscn.studygolang.com/ref/spec#For_statements) at its post statement. The "for" loop must be within the same function.

“ continue”语句开始其 post 语句中最内层的“ for”循环的下一次迭代。“ for”循环必须位于同一函数内。

```
ContinueStmt = "continue" [ Label ] .
```

If there is a label, it must be that of an enclosing "for" statement, and that is the one whose execution advances.

如果有一个标签，那么它必须是一个封闭的“ for”语句，而这个语句的执行将会进行。

```
RowLoop:
	for y, row := range rows {
		for x, data := range row {
			if data == endOfRow {
				continue RowLoop
			}
			row[x] = data + bias(x, y)
		}
	}
```

### Goto statements 后藤语句

A "goto" statement transfers control to the statement with the corresponding label within the same function.

“ goto”语句将控制转移到同一函数中具有相应标签的语句。

```
GotoStmt = "goto" Label .
goto Error
```

Executing the "goto" statement must not cause any variables to come into [scope](http://docscn.studygolang.com/ref/spec#Declarations_and_scope) that were not already in scope at the point of the goto. For instance, this example:

执行“ goto”语句不能导致任何变量进入 goto 的作用域，这些变量在 goto 的点上已经不在作用域中。例如，这个例子:

```
	goto L  // BAD
	v := 3
L:
```

is erroneous because the jump to label `L` skips the creation of `v`.

是错误的，因为跳转到标签 l 会跳过 v 的创建。

A "goto" statement outside a [block](http://docscn.studygolang.com/ref/spec#Blocks) cannot jump to a label inside that block. For instance, this example:

块外的“ goto”语句不能跳转到块内的标签。例如:

```
if n%2 == 1 {
	goto L1
}
for n > 0 {
	f()
	n--
L1:
	f()
	n--
}
```

is erroneous because the label `L1` is inside the "for" statement's block but the `goto` is not.

是错误的，因为标签 L1位于“ for”语句的块中，而 goto 则不是。

### Fallthrough statements 失败的陈述

A "fallthrough" statement transfers control to the first statement of the next case clause in an [expression "switch" statement](http://docscn.studygolang.com/ref/spec#Expression_switches). It may be used only as the final non-empty statement in such a clause.

“ fallthrough”语句将控制转移到表达式“ switch”语句中下一个 case 子句的第一个语句。它只能用作这种子句中最后的非空语句。

```
FallthroughStmt = "fallthrough" .
```

### Defer statements 推迟语句

A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a [return statement](http://docscn.studygolang.com/ref/spec#Return_statements), reached the end of its [function body](http://docscn.studygolang.com/ref/spec#Function_declarations), or because the corresponding goroutine is [panicking](http://docscn.studygolang.com/ref/spec#Handling_panics).

一个“推迟”语句调用一个函数，该函数的执行被推迟到周围函数返回的那一刻，这要么是因为周围函数执行了一个 return 语句，到达了函数体的末尾，要么是因为相应的 goroutine 感到恐慌。

```
DeferStmt = "defer" Expression .
```

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for [expression statements](http://docscn.studygolang.com/ref/spec#Expression_statements).

表达式必须是函数或方法调用; 它不能被括号括起来。内置函数的调用与表达式语句一样受到限制。

Each time a "defer" statement executes, the function value and parameters to the call are [evaluated as usual](http://docscn.studygolang.com/ref/spec#Calls) and saved anew but the actual function is not invoked. Instead, deferred functions are invoked immediately before the surrounding function returns, in the reverse order they were deferred. That is, if the surrounding function returns through an explicit [return statement](http://docscn.studygolang.com/ref/spec#Return_statements), deferred functions are executed *after* any result parameters are set by that return statement but *before* the function returns to its caller. If a deferred function value evaluates to `nil`, execution [panics](http://docscn.studygolang.com/ref/spec#Handling_panics) when the function is invoked, not when the "defer" statement is executed.

每次执行“ defer”语句时，调用的函数值和参数都会像通常一样进行计算并重新保存，但不会调用实际的函数。相反，延迟函数在周围函数返回之前被调用，调用顺序与延迟的顺序相反。也就是说，如果周围的函数通过显式的 return 语句返回，那么在 return 语句设置任何结果参数之后，在函数返回给调用者之前，将执行延迟函数。如果延迟函数值计算为 nil，则在调用函数时执行会感到恐慌，而不是在执行“ defer”语句时。

For instance, if the deferred function is a [function literal](http://docscn.studygolang.com/ref/spec#Function_literals) and the surrounding function has [named result parameters](http://docscn.studygolang.com/ref/spec#Function_types) that are in scope within the literal, the deferred function may access and modify the result parameters before they are returned. If the deferred function has any return values, they are discarded when the function completes. (See also the section on [handling panics](http://docscn.studygolang.com/ref/spec#Handling_panics).)

例如，如果延迟函数是一个匿名函数参数，并且周围的函数命名了文本范围内的结果参数，那么延迟函数可以在返回结果参数之前访问和修改结果参数。如果延迟函数有任何返回值，则在函数完成时将丢弃这些返回值。(参见处理恐慌的章节)

```
lock(l)
defer unlock(l)  // unlocking happens before surrounding function returns

// prints 3 2 1 0 before surrounding function returns
for i := 0; i <= 3; i++ {
	defer fmt.Print(i)
}

// f returns 42
func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}
```

## Built-in functions 内置功能

Built-in functions are [predeclared](http://docscn.studygolang.com/ref/spec#Predeclared_identifiers). They are called like any other function but some of them accept a type instead of an expression as the first argument.

内置函数是预声明的。它们像其他函数一样被调用，但是其中一些函数接受类型而不是表达式作为第一个参数。

The built-in functions do not have standard Go types, so they can only appear in [call expressions](http://docscn.studygolang.com/ref/spec#Calls); they cannot be used as function values.

内置函数没有标准的 Go 类型，因此它们只能出现在调用表达式中; 它们不能用作函数值。

### Close 关闭

For a channel `c`, the built-in function `close(c)` records that no more values will be sent on the channel. It is an error if `c` is a receive-only channel. Sending to or closing a closed channel causes a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics). Closing the nil channel also causes a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics). After calling `close`, and after any previously sent values have been received, receive operations will return the zero value for the channel's type without blocking. The multi-valued [receive operation](http://docscn.studygolang.com/ref/spec#Receive_operator) returns a received value along with an indication of whether the channel is closed.

对于通道 c，内置函数 close (c)记录没有更多的值将被发送到通道上。如果 c 是一个只接收通道，那么它就是一个错误。发送到或关闭闭合通道会导致运行时恐慌。关闭 nil 频道也会导致运行时恐慌。在调用 close 之后，以及在接收到任何以前发送的值之后，接收操作将返回通道类型的零值，而不会阻塞。多值接收操作返回一个接收值，并指示通道是否关闭。

### Length and capacity 长度和容量

The built-in functions `len` and `cap` take arguments of various types and return a result of type `int`. The implementation guarantees that the result always fits into an `int`.

内置函数 len 和 cap 接受各种类型的参数，并返回 int 类型的结果。实现保证结果始终符合 int。

```
Call      Argument type    Result

len(s)    string type      string length in bytes
          [n]T, *[n]T      array length (== n)
          []T              slice length
          map[K]T          map length (number of defined keys)
          chan T           number of elements queued in channel buffer

cap(s)    [n]T, *[n]T      array length (== n)
          []T              slice capacity
          chan T           channel buffer capacity
```

The capacity of a slice is the number of elements for which there is space allocated in the underlying array. At any time the following relationship holds:

片的容量是在基础数组中为其分配空间的元素数。以下关系在任何时候都有效:

```
0 <= len(s) <= cap(s)
```

The length of a `nil` slice, map or channel is 0. The capacity of a `nil` slice or channel is 0.

Nil 切片、映射或通道的长度为0。无切片或通道的容量为0。

The expression `len(s)` is [constant](http://docscn.studygolang.com/ref/spec#Constants) if `s` is a string constant. The expressions `len(s)` and `cap(s)` are constants if the type of `s` is an array or pointer to an array and the expression `s` does not contain [channel receives](http://docscn.studygolang.com/ref/spec#Receive_operator) or (non-constant) [function calls](http://docscn.studygolang.com/ref/spec#Calls); in this case `s` is not evaluated. Otherwise, invocations of `len` and `cap` are not constant and `s` is evaluated.

如果 s 是一个字符串常量，那么表达式 len (s)是常量。如果表达式的类型是数组或指向数组的指针，且表达式不包含通道接收或(非常量)函数调用，那么 len (s)和 cap (s)是常量; 在这种情况下，s 不计算。否则，调用的镜头和帽子不是常数，并进行了计算。

```
const (
	c1 = imag(2i)                    // imag(2i) = 2.0 is a constant
	c2 = len([10]float64{2})         // [10]float64{2} contains no function calls
	c3 = len([10]float64{c1})        // [10]float64{c1} contains no function calls
	c4 = len([10]float64{imag(2i)})  // imag(2i) is a constant and no function call is issued
	c5 = len([10]float64{imag(z)})   // invalid: imag(z) is a (non-constant) function call
)
var z complex128
```

### Allocation 分配

The built-in function `new` takes a type `T`, allocates storage for a [variable](http://docscn.studygolang.com/ref/spec#Variables) of that type at run time, and returns a value of type `*T` [pointing](http://docscn.studygolang.com/ref/spec#Pointer_types) to it. The variable is initialized as described in the section on [initial values](http://docscn.studygolang.com/ref/spec#The_zero_value).

内置函数 new 采用 t 类型，在运行时为该类型的变量分配存储空间，并返回指向它的 * t 类型的值。变量按照初始值一节中的描述初始化。

```
new(T)
```

For instance

例如

```
type S struct { a int; b float64 }
new(S)
```

allocates storage for a variable of type `S`, initializes it (`a=0`, `b=0.0`), and returns a value of type `*S` containing the address of the location.

为类型为 s 的变量分配存储空间，初始化它(a = 0，b = 0.0) ，并返回一个类型为 * s 的值，该值包含位置的地址。

### Making slices, maps and channels 制作切片、地图和通道

The built-in function `make` takes a type `T`, which must be a slice, map or channel type, optionally followed by a type-specific list of expressions. It returns a value of type `T` (not `*T`). The memory is initialized as described in the section on [initial values](http://docscn.studygolang.com/ref/spec#The_zero_value).

内置函数 make 接受类型 t，该类型必须是片、映射或通道类型，后面还可以跟特定类型的表达式列表。它返回类型为 t (而不是 * t)的值。内存按照初始值一节中的描述进行初始化。

```
Call             Type T     Result

make(T, n)       slice      slice of type T with length n and capacity n
make(T, n, m)    slice      slice of type T with length n and capacity m

make(T)          map        map of type T
make(T, n)       map        map of type T with initial space for approximately n elements

make(T)          channel    unbuffered channel of type T
make(T, n)       channel    buffered channel of type T, buffer size n
```

Each of the size arguments `n` and `m` must be of integer type or an untyped [constant](http://docscn.studygolang.com/ref/spec#Constants). A constant size argument must be non-negative and [representable](http://docscn.studygolang.com/ref/spec#Representability) by a value of type `int`; if it is an untyped constant it is given type `int`. If both `n` and `m` are provided and are constant, then `n` must be no larger than `m`. If `n` is negative or larger than `m` at run time, a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics) occurs.

每个 size 参数 n 和 m 必须是整型类型或非类型化常量。常量大小参数必须是非负的，并且可以用 int 类型的值表示; 如果是非类型化常量，则给定 int 类型。如果 n 和 m 都是常数，那么 n 必须不大于 m。如果 n 在运行时为负或大于 m，就会发生运行时恐慌。

```
s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
s := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
s := make([]int, 1<<63)         // illegal: len(s) is not representable by a value of type int
s := make([]int, 10, 0)         // illegal: len(s) > cap(s)
c := make(chan int, 10)         // channel with a buffer size of 10
m := make(map[string]int, 100)  // map with initial space for approximately 100 elements
```

Calling `make` with a map type and size hint `n` will create a map with initial space to hold `n` map elements. The precise behavior is implementation-dependent.

使用 map 类型和大小 hint n 调用 make 将创建一个初始空间包含 n 个 map 元素的 map。精确的行为依赖于实现。

### Appending to and copying slices 追加并复制切片

The built-in functions `append` and `copy` assist in common slice operations. For both functions, the result is independent of whether the memory referenced by the arguments overlaps.

内置函数在常见片操作中附加和复制辅助。对于这两个函数，结果与参数引用的内存是否重叠无关。

The [variadic](http://docscn.studygolang.com/ref/spec#Function_types) function `append` appends zero or more values `x` to `s` of type `S`, which must be a slice type, and returns the resulting slice, also of type `S`. The values `x` are passed to a parameter of type `...T` where `T` is the [element type](http://docscn.studygolang.com/ref/spec#Slice_types) of `S` and the respective [parameter passing rules](http://docscn.studygolang.com/ref/spec#Passing_arguments_to_..._parameters) apply. As a special case, `append` also accepts a first argument assignable to type `[]byte` with a second argument of string type followed by `...`. This form appends the bytes of the string.

可变参数函数将零个或多个值 x 追加到 s 类型的 s 后面，s 类型必须是片类型，并返回结果片，也是 s 类型。值 x 被传递给类型为... t 的参数，其中 t 是 s 的元素类型，并且适用相应的参数传递规则。作为一种特殊情况，append 还接受第一个参数，这个参数可以分配给 byte 类型，第二个参数是 string 类型，后面跟着... 。此表单追加字符串的字节。

```
append(s S, x ...T) S  // T is the element type of S
```

If the capacity of `s` is not large enough to fit the additional values, `append` allocates a new, sufficiently large underlying array that fits both the existing slice elements and the additional values. Otherwise, `append` re-uses the underlying array.

如果 s 的容量不足以容纳额外的值，那么附加分配一个新的、足够大/值底层数组，它既适合现有的 slice 元素，也适合额外的值。否则，追加重用基础数组。

```
s0 := []int{0, 0}
s1 := append(s0, 2)                // append a single element     s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)            // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

var t []interface{}
t = append(t, 42, 3.1415, "foo")   //                             t == []interface{}{42, 3.1415, "foo"}

var b []byte
b = append(b, "bar"...)            // append string contents      b == []byte{'b', 'a', 'r' }
```

The function `copy` copies slice elements from a source `src` to a destination `dst` and returns the number of elements copied. Both arguments must have [identical](http://docscn.studygolang.com/ref/spec#Type_identity) element type `T` and must be [assignable](http://docscn.studygolang.com/ref/spec#Assignability) to a slice of type `[]T`. The number of elements copied is the minimum of `len(src)` and `len(dst)`. As a special case, `copy` also accepts a destination argument assignable to type `[]byte` with a source argument of a string type. This form copies the bytes from the string into the byte slice.

函数将 slice 元素从源 src 复制到目标 dst，并返回复制的元素数。两个参数必须具有相同的元素类型 t，并且必须可以分配给类型[] t 的一个片段。复制的元素的数量是 len (src)和 len (dst)的最小值。作为一种特殊情况，copy 还接受一个目标参数，这个目标参数可以通过一个字符串类型的源参数来分配字节类型。此表单将字符串中的字节复制到字节片中。

```
copy(dst, src []T) int
copy(dst []byte, src string) int
```

Examples:

例子:

```
var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b == []byte("Hello")
```

### Deletion of map elements 删除地图元素

The built-in function `delete` removes the element with key `k` from a [map](http://docscn.studygolang.com/ref/spec#Map_types) `m`. The type of `k` must be [assignable](http://docscn.studygolang.com/ref/spec#Assignability) to the key type of `m`.

内置函数 delete 从 map m 中删除键为 k 的元素。K 的类型必须可分配给 m 的密钥类型。

```
delete(m, k)  // remove element m[k] from map m
```

If the map `m` is `nil` or the element `m[k]` does not exist, `delete` is a no-op.

如果映射 m 为 nil 或元素 m [ k ]不存在，则 delete 为 no-op。

### Manipulating complex numbers 操纵复数

Three functions assemble and disassemble complex numbers. The built-in function `complex` constructs a complex value from a floating-point real and imaginary part, while `real` and `imag` extract the real and imaginary parts of a complex value.

复数的三种组合和拆解功能。内置函数复合形从浮点实部和虚部构造复数值，而实部和虚部提取复数值的实部和虚部。

```
complex(realPart, imaginaryPart floatT) complexT
real(complexT) floatT
imag(complexT) floatT
```

The type of the arguments and return value correspond. For `complex`, the two arguments must be of the same floating-point type and the return type is the complex type with the corresponding floating-point constituents: `complex64` for `float32` arguments, and `complex128` for `float64` arguments. If one of the arguments evaluates to an untyped constant, it is first implicitly [converted](http://docscn.studygolang.com/ref/spec#Conversions) to the type of the other argument. If both arguments evaluate to untyped constants, they must be non-complex numbers or their imaginary parts must be zero, and the return value of the function is an untyped complex constant.

参数和返回值的类型对应。对于复合形式，两个参数必须是相同的浮点类型，并且返回类型是具有相应的浮点组成部分的复合形式: 对于 float32个参数，complex64对于 float64个参数，complex128对于 float64个参数。如果其中一个参数的计算结果为非类型化常数，则首先将其隐式转换为另一个参数的类型。如果两个参数都计算为无类型的常量，那么它们必须是非复数，或者它们的虚部必须为零，并且函数的返回值是无类型的复数常量。

For `real` and `imag`, the argument must be of complex type, and the return type is the corresponding floating-point type: `float32` for a `complex64` argument, and `float64` for a `complex128` argument. If the argument evaluates to an untyped constant, it must be a number, and the return value of the function is an untyped floating-point constant.

对于 real 和 imag，参数必须是复杂类型，返回类型是对应的浮点类型: float32用于复杂64参数，float64用于复杂128参数。如果参数计算为非类型化常量，则它必须是一个数字，并且函数的返回值为非类型化浮点常量。

The `real` and `imag` functions together form the inverse of `complex`, so for a value `z` of a complex type `Z`, `z == Z(complex(real(z), imag(z)))`.

实函数和 imag 函数一起构成复数的逆函数，因此对于复数类型 z 的值 z，z = = z (复数(实数(z) ，imag (z)))。

If the operands of these functions are all constants, the return value is a constant.

如果这些函数的操作数都是常量，则返回值为常量。

```
var a = complex(2, -2)             // complex128
const b = complex(1.0, -1.4)       // untyped complex constant 1 - 1.4i
x := float32(math.Cos(math.Pi/2))  // float32
var c64 = complex(5, -x)           // complex64
var s int = complex(1, 0)          // untyped complex constant 1 + 0i can be converted to int
_ = complex(1, 2<<s)               // illegal: 2 assumes floating-point type, cannot shift
var rl = real(c64)                 // float32
var im = imag(a)                   // float64
const c = imag(b)                  // untyped constant -1.4
_ = imag(3 << s)                   // illegal: 3 assumes complex type, cannot shift
```

### Handling panics 处理恐慌

Two built-in functions, `panic` and `recover`, assist in reporting and handling [run-time panics](http://docscn.studygolang.com/ref/spec#Run_time_panics) and program-defined error conditions.

两个内置的功能，恐慌和恢复，协助报告和处理运行时恐慌和程序定义的错误条件。

```
func panic(interface{})
func recover() interface{}
```

While executing a function `F`, an explicit call to `panic` or a [run-time panic](http://docscn.studygolang.com/ref/spec#Run_time_panics) terminates the execution of `F`. Any functions [deferred](http://docscn.studygolang.com/ref/spec#Defer_statements) by `F` are then executed as usual. Next, any deferred functions run by `F's` caller are run, and so on up to any deferred by the top-level function in the executing goroutine. At that point, the program is terminated and the error condition is reported, including the value of the argument to `panic`. This termination sequence is called *panicking*.

在执行函数 f 时，显式调用 panic 或运行时 panic 会终止 f 的执行，然后 f 延迟的所有函数都会照常执行。接下来，运行 f 的调用方运行的任何延迟函数，以此类推，直到执行 goroutine 中的顶级函数所延迟的任何函数。此时，程序终止，并报告错误条件，包括参数的值以表示恐慌。这个终止序列被称为恐慌。

```
panic(42)
panic("unreachable")
panic(Error("cannot parse"))
```

The `recover` function allows a program to manage behavior of a panicking goroutine. Suppose a function `G` defers a function `D` that calls `recover` and a panic occurs in a function on the same goroutine in which `G` is executing. When the running of deferred functions reaches `D`, the return value of `D`'s call to `recover` will be the value passed to the call of `panic`. If `D` returns normally, without starting a new `panic`, the panicking sequence stops. In that case, the state of functions called between `G` and the call to `panic` is discarded, and normal execution resumes. Any functions deferred by `G` before `D` are then run and `G`'s execution terminates by returning to its caller.

Recover 函数允许程序管理恐慌的 goroutine 的行为。假设一个函数 g 延迟了一个调用 recover 的函数 d，并且在执行 g 的同一个 goroutine 上的一个函数中发生了一个 panic。当延迟函数的运行达到 d 时，d 调用 recover 的返回值将是传递给 panic 调用的值。如果 d 正常返回，没有启动新的恐慌，恐慌序列停止。在这种情况下，g 和 panic 调用之间调用的函数的状态将被丢弃，正常的执行将恢复。在 d 之前被 g 延迟的任何函数都会被运行，g 的执行通过返回到它的调用者来终止。

The return value of `recover` is `nil` if any of the following conditions holds:

如果下列任何一个条件成立，则 recover 的返回值为 nil:

- `panic`'s argument was 的论点是`nil`;
- the goroutine is not panicking; Goroutine 没有惊慌失措;
- `recover` was not called directly by a deferred function. 不是由延迟函数直接调用的

The `protect` function in the example below invokes the function argument `g` and protects callers from run-time panics raised by `g`.

下面示例中的 protect 函数调用函数参数 g，并保护调用者免受 g 引起的运行时恐慌。

```
func protect(g func()) {
	defer func() {
		log.Println("done")  // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}
```

### Bootstrapping 自力更生

Current implementations provide several built-in functions useful during bootstrapping. These functions are documented for completeness but are not guaranteed to stay in the language. They do not return a result.

当前的实现提供了几个在启动过程中有用的内置函数。这些函数是为了完整性而编制的，但不能保证使用该语言。它们不返回结果。

```
Function   Behavior

print      prints all arguments; formatting of arguments is implementation-specific
println    like print but prints spaces between arguments and a newline at the end
```

Implementation restriction: `print` and `println` need not accept arbitrary argument types, but printing of boolean, numeric, and string [types](http://docscn.studygolang.com/ref/spec#Types) must be supported.

实现限制: print 和 println 不需要接受任意的参数类型，但是必须支持 boolean、 numeric 和 string 类型的打印。

## Packages 包装

Go programs are constructed by linking together *packages*. A package in turn is constructed from one or more source files that together declare constants, types, variables and functions belonging to the package and which are accessible in all files of the same package. Those elements may be [exported](http://docscn.studygolang.com/ref/spec#Exported_identifiers) and used in another package.

Go 程序是由链接在一起的软件包构成的。包依次由一个或多个源文件构成，这些源文件一起声明属于包的常量、类型、变量和函数，这些常量、类型和函数可以在同一个包的所有文件中访问。这些元素可以导出并在另一个包中使用。

### Source file organization 源文件组织

Each source file consists of a package clause defining the package to which it belongs, followed by a possibly empty set of import declarations that declare packages whose contents it wishes to use, followed by a possibly empty set of declarations of functions, types, variables, and constants.

每个源文件包含一个定义它所属的包的包子句，后面跟着一组可能为空的导入声明，这些导入声明声明包的内容是它希望使用的，后面跟着一组可能为空的函数、类型、变量和常量声明。

```
SourceFile       = PackageClause ";" { ImportDecl ";" } { TopLevelDecl ";" } .
```

### Package clause 包装条款

A package clause begins each source file and defines the package to which the file belongs.

Package 子句以每个源文件开始，并定义文件所属的包。

```
PackageClause  = "package" PackageName .
PackageName    = identifier .
```

The PackageName must not be the [blank identifier](http://docscn.studygolang.com/ref/spec#Blank_identifier).

PackageName 不能是空白标识符。

```
package math
```

A set of files sharing the same PackageName form the implementation of a package. An implementation may require that all source files for a package inhabit the same directory.

一组共享相同 PackageName 的文件构成了包的实现。实现可能要求包的所有源文件都位于同一目录中。

### Import declarations 进口报关单

An import declaration states that the source file containing the declaration depends on functionality of the *imported* package ([§Program initialization and execution](http://docscn.studygolang.com/ref/spec#Program_initialization_and_execution)) and enables access to [exported](http://docscn.studygolang.com/ref/spec#Exported_identifiers) identifiers of that package. The import names an identifier (PackageName) to be used for access and an ImportPath that specifies the package to be imported.

导入声明声明包含声明的源文件取决于导入包的功能(程序初始化和执行) ，并允许访问该包的导出标识符。导入命名了一个用于访问的标识符(PackageName)和一个指定要导入的包的 ImportPath。

```
ImportDecl       = "import" ( ImportSpec | "(" { ImportSpec ";" } ")" ) .
ImportSpec       = [ "." | PackageName ] ImportPath .
ImportPath       = string_lit .
```

The PackageName is used in [qualified identifiers](http://docscn.studygolang.com/ref/spec#Qualified_identifiers) to access exported identifiers of the package within the importing source file. It is declared in the [file block](http://docscn.studygolang.com/ref/spec#Blocks). If the PackageName is omitted, it defaults to the identifier specified in the [package clause](http://docscn.studygolang.com/ref/spec#Package_clause) of the imported package. If an explicit period (`.`) appears instead of a name, all the package's exported identifiers declared in that package's [package block](http://docscn.studygolang.com/ref/spec#Blocks) will be declared in the importing source file's file block and must be accessed without a qualifier.

PackageName 用于限定标识符中，以访问导入源文件中的包的导出标识符。它在文件块中声明。如果省略 PackageName，它将缺省为导入包的 package 子句中指定的标识符。如果一个明确的期间(.)出现而不是名称，则在包的包块中声明的所有包的导出标识符都将在导入源文件的文件块中声明，并且必须在没有限定符的情况下访问。

The interpretation of the ImportPath is implementation-dependent but it is typically a substring of the full file name of the compiled package and may be relative to a repository of installed packages.

ImportPath 的解释依赖于实现，但它通常是编译包的完整文件名的子字符串，并且可能相对于已安装包的存储库。

Implementation restriction: A compiler may restrict ImportPaths to non-empty strings using only characters belonging to [Unicode's](https://www.unicode.org/versions/Unicode6.3.0/) L, M, N, P, and S general categories (the Graphic characters without spaces) and may also exclude the characters `!"#$%&'()*,:;<=>?[\]^`{|}` and the Unicode replacement character U+FFFD.

实现限制: 编译器可能仅使用属于 Unicode 的 l、 m、 n、 p 和 s 通用类别(不带空格的图形字符)的字符将 ImportPaths 限制为非空字符串，还可能排除字符!”#$%&'()*,:;<=>？[] ^ ‘{ | }和 Unicode 替换字符 u + fffd。

Assume we have compiled a package containing the package clause `package math`, which exports function `Sin`, and installed the compiled package in the file identified by `"lib/math"`. This table illustrates how `Sin` is accessed in files that import the package after the various types of import declaration.

假设我们已经编译了一个包含包子句包 math 的包，该包导出函数 Sin，并将已编译的包安装到由“ lib/math”标识的文件中。此表说明了在各种类型的导入声明之后，如何在导入包的文件中访问 Sin。

```
Import declaration          Local name of Sin

import   "lib/math"         math.Sin
import m "lib/math"         m.Sin
import . "lib/math"         Sin
```

An import declaration declares a dependency relation between the importing and imported package. It is illegal for a package to import itself, directly or indirectly, or to directly import a package without referring to any of its exported identifiers. To import a package solely for its side-effects (initialization), use the [blank](http://docscn.studygolang.com/ref/spec#Blank_identifier) identifier as explicit package name:

导入声明声明导入包和导入包之间的依赖关系。包直接或间接导入自身，或者直接导入包而不引用其导出的任何标识符，都是非法的。若要只为副作用(初始化)而导入包，请使用空标识符作为显式包名:

```
import _ "lib/math"
```

### An example package 一个示例包

Here is a complete Go package that implements a concurrent prime sieve.

这是一个完整的 Go 包，实现了并发质数筛。

```
package main

import "fmt"

// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i  // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {  // Loop over values received from 'src'.
		if i%prime != 0 {
			dst <- i  // Send 'i' to channel 'dst'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int)  // Create a new channel.
	go generate(ch)       // Start generate() as a subprocess.
	for {
		prime := <-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
	sieve()
}
```

## Program initialization and execution 程序的初始化和执行

### The zero value 零值

When storage is allocated for a [variable](http://docscn.studygolang.com/ref/spec#Variables), either through a declaration or a call of `new`, or when a new value is created, either through a composite literal or a call of `make`, and no explicit initialization is provided, the variable or value is given a default value. Each element of such a variable or value is set to the *zero value* for its type: `false` for booleans, `0` for numeric types, `""` for strings, and `nil` for pointers, functions, interfaces, slices, channels, and maps. This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed if no value is specified.

当通过一个声明或一个 new 调用为一个变量分配存储空间时，或者当通过一个复合文字或一个 make 调用创建一个新值时，没有提供显式的初始化，变量或值被赋予一个默认值。这种变量或值的每个元素的类型都被设置为零值: 对于布尔值为 false，对于数字类型为0，对于字符串为“”，对于指针、函数、接口、片段、通道和映射为 nil。这个初始化是递归完成的，因此，例如，如果没有指定值，结构数组中的每个元素的字段都将为零。

These two simple declarations are equivalent:

这两个简单的声明是等价的:

```
var i int
var i int = 0
```

After

之后

```
type T struct { i int; f float64; next *T }
t := new(T)
```

the following holds:

以下是一些观点:

```
t.i == 0
t.f == 0.0
t.next == nil
```

The same would also be true after

同样的道理也适用于年之后

```
var t T
```

### Package initialization 包初始化

Within a package, package-level variable initialization proceeds stepwise, with each step selecting the variable earliest in *declaration order* which has no dependencies on uninitialized variables.

在包中，包级别的变量初始化是逐步进行的，每一步都是按声明顺序选择最早的变量，这些变量与未初始化的变量没有依赖关系。

More precisely, a package-level variable is considered *ready for initialization* if it is not yet initialized and either has no [initialization expression](http://docscn.studygolang.com/ref/spec#Variable_declarations) or its initialization expression has no *dependencies* on uninitialized variables. Initialization proceeds by repeatedly initializing the next package-level variable that is earliest in declaration order and ready for initialization, until there are no variables ready for initialization.

更确切地说，如果包级变量尚未初始化，或者没有初始化表达式，或者其初始化表达式对未初始化变量没有依赖关系，则认为它已经可以进行初始化了。初始化过程通过重复初始化下一个包级别的变量来进行，这个变量最早按照声明顺序并准备好进行初始化，直到没有可以进行初始化的变量。

If any variables are still uninitialized when this process ends, those variables are part of one or more initialization cycles, and the program is not valid.

如果在这个过程结束时任何变量仍未初始化，那么这些变量是一个或多个初始化周期的一部分，并且程序无效。

Multiple variables on the left-hand side of a variable declaration initialized by single (multi-valued) expression on the right-hand side are initialized together: If any of the variables on the left-hand side is initialized, all those variables are initialized in the same step.

由右侧的单个(多值)表达式初始化的变量声明左侧的多个变量一起初始化: 如果左侧的任何变量被初始化，所有这些变量都在同一步中初始化。

```
var x = a
var a, b = f() // a and b are initialized together, before x is initialized
```

For the purpose of package initialization, [blank](http://docscn.studygolang.com/ref/spec#Blank_identifier) variables are treated like any other variables in declarations.

为了进行包初始化，空变量在声明中被当作任何其他变量处理。

The declaration order of variables declared in multiple files is determined by the order in which the files are presented to the compiler: Variables declared in the first file are declared before any of the variables declared in the second file, and so on.

在多个文件中声明的变量的声明顺序由文件呈现给编译器的顺序决定: 在第一个文件中声明的变量先于在第二个文件中声明的任何变量，依此类推。

Dependency analysis does not rely on the actual values of the variables, only on lexical *references* to them in the source, analyzed transitively. For instance, if a variable `x`'s initialization expression refers to a function whose body refers to variable `y` then `x` depends on `y`. Specifically:

依赖性分析并不依赖于变量的实际值，而仅仅依赖于源中对它们的词法引用，并进行了传递性分析。例如，如果一个变量 x 的初始化表达式引用了一个函数，该函数的主体引用了变量 y，那么 x 取决于 y:

- A reference to a variable or function is an identifier denoting that variable or function. 对变量或函数的引用是表示该变量或函数的标识符
- A reference to a method 方法参考一种方法`m` is a 是一个[method value 方法值](http://docscn.studygolang.com/ref/spec#Method_values) or 或[method expression 方法表达式](http://docscn.studygolang.com/ref/spec#Method_expressions) of the form 的表格`t.m`, where the (static) type of ，其中的(静态)类型`t` is not an interface type, and the method 不是接口类型，并且方法`m` is in the 是在[method set 方法集](http://docscn.studygolang.com/ref/spec#Method_sets) of 的`t`. It is immaterial whether the resulting function value 。这是无关紧要的是，由此产生的函数值`t.m` is invoked. 被调用
- A variable, function, or method 变量、函数或方法`x` depends on a variable 取决于一个变量`y` if 如果`x`'s initialization expression or body (for functions and methods) contains a reference to 的初始化表达式或主体(用于函数和方法)包含对`y` or to a function or method that depends on 或者一个函数或者方法`y`.

For example, given the declarations

例如，给定声明

```
var (
	a = c + b  // == 9
	b = f()    // == 4
	c = f()    // == 5
	d = 3      // == 5 after initialization has finished
)

func f() int {
	d++
	return d
}
```

the initialization order is `d`, `b`, `c`, `a`. Note that the order of subexpressions in initialization expressions is irrelevant: `a = c + b` and `a = b + c` result in the same initialization order in this example.

初始化顺序是 d，b，c，a。注意，初始化表达式中子表达式的顺序是无关的: a = c + b 和 a = b + c 的初始化顺序在本例中是相同的。

Dependency analysis is performed per package; only references referring to variables, functions, and (non-interface) methods declared in the current package are considered. If other, hidden, data dependencies exists between variables, the initialization order between those variables is unspecified.

每个包执行依赖性分析; 只考虑引用当前包中声明的变量、函数和(非接口)方法。如果变量之间存在其他隐藏的数据依赖关系，则未指定这些变量之间的初始化顺序。

For instance, given the declarations

例如，给定声明

```
var x = I(T{}).ab()   // x has an undetected, hidden dependency on a and b
var _ = sideEffect()  // unrelated to x, a, or b
var a = b
var b = 42

type I interface      { ab() []int }
type T struct{}
func (T) ab() []int   { return []int{a, b} }
```

the variable `a` will be initialized after `b` but whether `x` is initialized before `b`, between `b` and `a`, or after `a`, and thus also the moment at which `sideEffect()` is called (before or after `x` is initialized) is not specified.

变量 a 将在 b 之后初始化，但是 x 是在 b 之前初始化，还是在 b 和 a 之间初始化，还是在 a 之后初始化，因此没有指定调用 sideEffect ()的时刻(在初始化 x 之前或之后)。

Variables may also be initialized using functions named `init` declared in the package block, with no arguments and no result parameters.

也可以使用在包块中声明的名为 init 的函数初始化变量，这些函数没有参数和结果参数。

```
func init() { … }
```

Multiple such functions may be defined per package, even within a single source file. In the package block, the `init` identifier can be used only to declare `init` functions, yet the identifier itself is not [declared](http://docscn.studygolang.com/ref/spec#Declarations_and_scope). Thus `init` functions cannot be referred to from anywhere in a program.

可以为每个包定义多个这样的函数，甚至可以在单个源文件中定义。在包块中，init 标识符只能用于声明 init 函数，而不能声明标识符本身。因此 init 函数不能从程序中的任何地方引用。

A package with no imports is initialized by assigning initial values to all its package-level variables followed by calling all `init` functions in the order they appear in the source, possibly in multiple files, as presented to the compiler. If a package has imports, the imported packages are initialized before initializing the package itself. If multiple packages import a package, the imported package will be initialized only once. The importing of packages, by construction, guarantees that there can be no cyclic initialization dependencies.

初始化一个没有导入的包，方法是为其所有的包级变量赋予初始值，然后按照初始值在源代码中出现的顺序调用所有 init 函数，可能是在多个文件中调用初始值，并呈现给编译器。如果包包含导入，则在初始化包本身之前初始化导入的包。如果多个包导入一个包，导入的包只初始化一次。通过构造导入包，可以保证不存在循环初始化依赖关系。

Package initialization—variable initialization and the invocation of `init` functions—happens in a single goroutine, sequentially, one package at a time. An `init` function may launch other goroutines, which can run concurrently with the initialization code. However, initialization always sequences the `init` functions: it will not invoke the next one until the previous one has returned.

包初始化ー变量初始化和 init 函数的调用ー发生在单个 goroutine 中，一次一个包。Init 函数可以启动其他 goroutine，这些 goroutine 可以与初始化代码并发运行。但是，初始化总是对 init 函数进行序列化: 在前一个函数返回之前，它不会调用下一个函数。

To ensure reproducible initialization behavior, build systems are encouraged to present multiple files belonging to the same package in lexical file name order to a compiler.

为了确保可重复的初始化行为，鼓励构建系统以词法文件名顺序向编译器显示属于同一包的多个文件。

### Program execution 程序执行

A complete program is created by linking a single, unimported package called the *main package* with all the packages it imports, transitively. The main package must have package name `main` and declare a function `main` that takes no arguments and returns no value.

一个完整的程序是通过将一个称为 main 包的单个未导入包与它导入的所有包相链接而创建的。Main 包必须具有包名 main 并声明一个不带参数且不返回值的函数 main。

```
func main() { … }
```

Program execution begins by initializing the main package and then invoking the function `main`. When that function invocation returns, the program exits. It does not wait for other (non-`main`) goroutines to complete.

程序执行首先初始化主包，然后调用函数 main。当函数调用返回时，程序退出。它不等待其他(非主要) goroutine 完成。

## Errors 错误

The predeclared type `error` is defined as

预声明的类型错误定义为

```
type error interface {
	Error() string
}
```

It is the conventional interface for representing an error condition, with the nil value representing no error. For instance, a function to read data from a file might be defined:

它是表示错误条件的常规接口，用 nil 值表示没有错误。例如，一个从文件中读取数据的函数可能被定义为:

```
func Read(f *File, b []byte) (n int, err error)
```

## Run-time panics 运行时恐慌

Execution errors such as attempting to index an array out of bounds trigger a *run-time panic* equivalent to a call of the built-in function [`panic`](http://docscn.studygolang.com/ref/spec#Handling_panics) with a value of the implementation-defined interface type `runtime.Error`. That type satisfies the predeclared interface type [`error`](http://docscn.studygolang.com/ref/spec#Errors). The exact error values that represent distinct run-time error conditions are unspecified.

执行错误(如试图为数组建立超出界限的索引)会触发一个运行时错误，该错误等同于内置函数错误的调用，其值为实现定义的接口类型运行时。错误。该类型满足预声明的接口类型错误。未指定表示不同运行时错误条件的确切错误值。

```
package runtime

type Error interface {
	error
	// and perhaps other methods
}
```

## System considerations 系统方面的考虑

### Package 包装`unsafe`

The built-in package `unsafe`, known to the compiler and accessible through the [import path](http://docscn.studygolang.com/ref/spec#Import_declarations) `"unsafe"`, provides facilities for low-level programming including operations that violate the type system. A package using `unsafe` must be vetted manually for type safety and may not be portable. The package provides the following interface:

内置的包不安全，编译器知道，可以通过“不安全”的导入路径访问，为低级编程提供了便利，包括违反类型系统的操作。使用不安全的包必须手动检查类型安全，并且可能不可移植。该包提供了以下接口:

```
package unsafe

type ArbitraryType int  // shorthand for an arbitrary Go type; it is not a real type
type Pointer *ArbitraryType

func Alignof(variable ArbitraryType) uintptr
func Offsetof(selector ArbitraryType) uintptr
func Sizeof(variable ArbitraryType) uintptr
```

A `Pointer` is a [pointer type](http://docscn.studygolang.com/ref/spec#Pointer_types) but a `Pointer` value may not be [dereferenced](http://docscn.studygolang.com/ref/spec#Address_operators). Any pointer or value of [underlying type](http://docscn.studygolang.com/ref/spec#Types) `uintptr` can be converted to a type of underlying type `Pointer` and vice versa. The effect of converting between `Pointer` and `uintptr` is implementation-defined.

指针是一种指针类型，但是指针值不能被解除引用。任何基础类型 uintptr 的指针或值都可以转换为基础类型 Pointer，反之亦然。指针和 uintptr 之间的转换效果是实现定义的。

```
var f float64
bits = *(*uint64)(unsafe.Pointer(&f))

type ptr unsafe.Pointer
bits = *(*uint64)(ptr(&f))

var p ptr = nil
```

The functions `Alignof` and `Sizeof` take an expression `x` of any type and return the alignment or size, respectively, of a hypothetical variable `v` as if `v` was declared via `var v = x`.

函数 Alignof 和 Sizeof 采用任意类型的表达式 x，并分别返回假设变量 v 的对齐方式或大小，就好像 v 是通过 var v = x 声明的一样。

The function `Offsetof` takes a (possibly parenthesized) [selector](http://docscn.studygolang.com/ref/spec#Selectors) `s.f`, denoting a field `f` of the struct denoted by `s` or `*s`, and returns the field offset in bytes relative to the struct's address. If `f` is an [embedded field](http://docscn.studygolang.com/ref/spec#Struct_types), it must be reachable without pointer indirections through fields of the struct. For a struct `s` with field `f`:

函数 Offsetof 接受一个(可能是括号)选择器 s.f，表示由 s 或 * s 表示的 struct 的字段 f，并返回相对于 struct 的地址的字节偏移量。如果 f 是嵌入字段，则必须能够通过 struct 的字段在没有指针间接指向的情况下访问它。对于字段 f 为的结构体:

```
uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f) == uintptr(unsafe.Pointer(&s.f))
```

Computer architectures may require memory addresses to be *aligned*; that is, for addresses of a variable to be a multiple of a factor, the variable's type's *alignment*. The function `Alignof` takes an expression denoting a variable of any type and returns the alignment of the (type of the) variable in bytes. For a variable `x`:

计算机体系结构可能要求内存地址对齐; 也就是说，变量的地址是一个因子的倍数，即变量的类型对齐。Alignof 函数接受表示任何类型的变量的表达式，并返回(类型)变量的对齐方式(以字节为单位)。对于变量 x:

```
uintptr(unsafe.Pointer(&x)) % unsafe.Alignof(x) == 0
```

Calls to `Alignof`, `Offsetof`, and `Sizeof` are compile-time constant expressions of type `uintptr`.

对 Alignof、 Offsetof 和 Sizeof 的调用是 uintptr 类型的编译时常量表达式。

### Size and alignment guarantees 尺寸和对齐保证

For the [numeric types](http://docscn.studygolang.com/ref/spec#Numeric_types), the following sizes are guaranteed:

对于数字类型，保证使用以下大小:

```
type                                 size in bytes

byte, uint8, int8                     1
uint16, int16                         2
uint32, int32, float32                4
uint64, int64, float64, complex64     8
complex128                           16
```

The following minimal alignment properties are guaranteed:

保证以下最小对齐属性:

1. For a variable 对于一个变量`x` of any type: 任何类型:`unsafe.Alignof(x)` is at least 1. 至少是1
2. For a variable 对于一个变量`x` of struct type: 的结构类型:`unsafe.Alignof(x)` is the largest of all the values 是所有价值中最大的`unsafe.Alignof(x.f)` for each field 每个字段`f` of 的`x`, but at least 1. ，但最少1
3. For a variable 对于一个变量`x` of array type: 数组类型:`unsafe.Alignof(x)` is the same as the alignment of a variable of the array's element type. 与数组元素类型的变量对齐方式相同

A struct or array type has size zero if it contains no fields (or elements, respectively) that have a size greater than zero. Two distinct zero-size variables may have the same address in memory.

如果结构或数组类型不包含大小大于零的字段(或元素) ，则其大小为零。两个不同的零大小变量在内存中可能具有相同的地址。
