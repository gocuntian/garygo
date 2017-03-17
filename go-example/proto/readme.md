#Protocol Buffers Version 3 Language Specification
这是Protocol Buffers语言（proto3）的版本3的语言规范参考。语法使用扩展的Backus-Naur形式（EBNF）指定：
有关使用proto3的更多信息，请参阅语言指南。
词汇元素
协议缓冲区版本3语言规范

这是Protocol Buffers语言（proto3）的版本3的语言规范参考。语法使用扩展的Backus-Naur形式（EBNF）指定：

|   alternation
()  grouping
[]  option (zero or one time)
{}  repetition (any number of times)

有关使用proto3的更多信息，请参阅语言指南。

#词汇元素

##字母和数字

letter =“A”...“Z”| “a”...“z”
decimalDigit =“0”...“9”
octalDigit =“0”...“7”
hexDigit =“0”...“9”| “A”...“F”| “a”...“f”

##身份标识

ident = letter {letter | decimalDigit | “_”}
fullIdent = ident {“。” ident}
messageName = ident
enumName = ident
fieldName = ident
oneofName = ident
mapName = ident
serviceName = ident
rpcName = ident
messageType = [“.” ] {ident “.” } messageName
enumType = [“.” ] {ident “.” } enumName

##整数文字
intLit = decimalLit | octalLit | hexLit
decimalLit =（“1”...“9”）{decimalDigit}
octalLit =“0”{octalDigit}
hexLit =“0”（“x”|“X”）hexDigit {hexDigit}

##浮点文字
floatLit = ( decimals "." [ decimals ] [ exponent ] | decimals exponent | "."decimals [ exponent ] ) | "inf" | "nan"
decimals  = decimalDigit { decimalDigit }
exponent  = ( "e" | "E" ) [ "+" | "-" ] decimals 

##布尔值
boolLit = "true" | "false" 

##字符串文字
strLit = ( "'" { charValue } "'" ) |  ( '"' { charValue } '"' )
charValue = hexEscape | octEscape | charEscape | /[^\0\n\\]/
hexEscape = '\' ( "x" | "X" ) hexDigit hexDigit
octEscape = '\' octalDigit octalDigit octalDigit
charEscape = '\' ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | '\' | "'" | '"' )
quote = "'" | '"'

##EmptyStatement

emptyStatement =“;”

##Constant
constant = fullIdent | ( [ "-" | "+" ] intLit ) | ( [ "-" | "+" ] floatLit ) | strLit | boolLit 

#Syntax 语法语句用于定义protobuf版本。
syntax =“syntax”“=”quote“proto3”quote“;”

##导入语句 import语句用于导入另一个.proto的定义。

import = "import" [ "weak" | "public" ] strLit ";" 
例： import public "other.proto";

##包 包说明符可用于防止协议消息类型之间的名称冲突。

package = "package" fullIdent ";"
例：package foo.bar;

###选项
    选项可以在原型文件，消息，枚举和服务中使用。选项可以是protobuf定义的选项或自定义选项。有关详细信息，请参阅语言指南中的选项。

    option = "option" optionName  "=" constant ";"
    optionName = ( ident | "(" fullIdent ")" ) { "." ident }

    例：option java_package = "com.example.foo";

###字段 
    字段是协议缓冲区消息的基本元素。字段可以是正常字段，其中一个字段或映射字段。字段具有类型和字段号。

    type = "double" | "float" | "int32" | "int64" | "uint32" | "uint64" | "sint32" | "sint64" | "fixed32" | "fixed64" | "sfixed32" | "sfixed64" | "bool" | "string" | "bytes" | messageType | enumType
    fieldNumber = intLit;

###正常字段
    每个字段都有类型，名称和字段号。它可能有字段选项

    field = [ "repeated" ] type fieldName "=" fieldNumber [ "[" fieldOptions "]" ] ";"
    fieldOptions = fieldOption { ","  fieldOption }
    fieldOption = optionName "=" constant
例子：

    foo.bar nested_message = 2;
    repeated int32 samples = 4 [packed=true];

###Oneof and oneof field
    一个由一个字段和一个名字组成。

    oneof = "oneof" oneofName "{" { oneofField | emptyStatement } "}"
    oneofField = type fieldName "=" fieldNumber [ "[" fieldOptions "]" ] ";"
例：
    oneof foo {
        string name = 4;
        SubMessage sub_message = 9;
    }

###地图字段
    映射字段具有键类型，值类型，名称和字段号。键类型可以是任何整数或字符串类型。

    mapField = "map" "<" keyType "," type ">" mapName "=" fieldNumber [ "[" fieldOptions "]" ] ";"
    keyType = "int32" | "int64" | "uint32" | "uint64" | "sint32" | "sint64" | "fixed32" | "fixed64" | "sfixed32" | "sfixed64" | "bool" | "string"
例：
    map<string, Project> projects = 3;
###保留

    保留的语句声明不能在此消息中使用的字段号或字段名的范围。

    reserved = "reserved" ( ranges | fieldNames ) ";"
    fieldNames = fieldName { "," fieldName }
例子：
    reserved 2, 15, 9 to 11;
    reserved "foo", "bar";

##顶级定义

###枚举定义
    枚举定义包括名称和枚举体。枚举体可以有选项和枚举字段。枚举定义必须以枚举值零开始。

    enum = "enum" enumName enumBody
    enumBody = "{" { option | enumField | emptyStatement } "}"
    enumField = ident "=" intLit [ "[" enumValueOption { ","  enumValueOption } "]" ]";"
    enumValueOption = optionName "=" constant
例：

    enum EnumAllowingAlias {
        option allow_alias = true;
        UNKNOWN = 0;
        STARTED = 1;
        RUNNING = 2 [(custom_option) = "hello world"];
    }

###消息定义
    消息由消息名称和消息体组成。消息体可以具有字段，嵌套枚举定义，嵌套消息定义，选项，oneofs，映射字段和保留语句。

    message = "message" messageName messageBody
    messageBody = "{" { field | enum | message | option | oneof | mapField | reserved | emptyStatement } "}"
例：

    message Outer {
        option (my_option).a = true;
        message Inner {   // Level 2
            int64 ival = 1;
        }
        map<int32, string> my_map = 2;
    }
##服务定义
    service = "service" serviceName "{" { option | rpc | stream | emptyStatement } "}"
    rpc = "rpc" rpcName "(" [ "stream" ] messageType ")" "returns" "(" [ "stream" ] messageType ")" (( "{" {option | emptyStatement } "}" ) | ";")
例：
    service SearchService {
       rpc Search (SearchRequest) returns (SearchResponse);
    }

##Proto文件

    proto = syntax { import | package | option | topLevelDef | emptyStatement }
    topLevelDef = message | enum | service

.proto文件示例：

    syntax = “proto3”;
    import public “other.proto”;

    option java_package = "com.example.foo";
    enum EnumAllowingAlias {
        option allow_alias = true;
        UNKNOWN = 0;
        STARTED = 1;
        RUNNING = 2 [(custom_option) = "hello world"];
    }

    message outer {
        option (my_option).a = true;
        message inner {   // Level 2
            int64 ival = 1;
        }
        repeated inner inner_message = 2;
        EnumAllowingAlias enum_field =3;
        map<int32, string> my_map = 4;
    }