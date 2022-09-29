package test

import (
	"fmt"
	"testing"
)

func TestStringVar(t *testing.T) {
	name := "ming"
	fmt.Printf("&name = %v \n", &name)
	StringVar(&name, name)
}

func StringVar(p *string, name string)  {	//这里的 name是值复制, 虽然值相同, 但地址不同
	if p == &name {
		fmt.Printf("yes, p = &name = %v \n", p)
	}else {
		fmt.Printf("no, p = %v, &name = %v \n", p, &name)
	}
}

func TestNewStringValue(t *testing.T) {
	val := "ming"
	var name string
	p := &name
/*	fmt.Printf("Test: val = %v \n" +
		"&val = %v \n" +
		"p = %v \n", val, &val, p)*/
	fmt.Printf("return %v, %T \n", newStringValue(val, p), newStringValue(val, p))
}

type stringValue string	//类型别名, 虽然是别名, 但这就成了两种类型, 传递时需要强制转换
func newStringValue(val string, p *string) *stringValue {	//这里的 val也是值复制
	*p = val
/*	fmt.Printf("val = %v \n" +
		"&val = %v \n" +
		"p = %v \n" +
		"*p = %v \n", val, &val, p, *p)*/
	return (*stringValue)(p)	//强制类型转换, p是 *string类型, 需要强制转换成 *stringValue类型
}

func TestCommandLine_Var(t *testing.T) {
	println("不测这个了, 涉及太多")
}

// Value is the interface to the dynamic value stored in a flag.
// (The default value is represented as a string.)
//
// If a Value has an IsBoolFlag() bool method returning true,
// the command-line parser makes -name equivalent to -name=true
// rather than using the next command-line argument.
//
// Set is called once, in command line order, for each flag present.
// The flag package may call the String method with a zero-valued receiver,
// such as a nil pointer.
/*
翻译：
值是存储在标志中的动态值的接口。(默认值以字符串形式表示。)
如果Value有一个IsBoolFlag() bool方法返回true，
命令行解析器将使name等价于name=true，而不是使用下一个命令行参数。
Set按命令行顺序为每个标记调用一次。标志包可以使用零值的接收器调用String方法，例如nil指针。
*/

/*type Value interface {
	String() string
	Set(string) error
}

type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler. What happens after Usage is called depends
	// on the ErrorHandling setting; for the command line, this defaults
	// to ExitOnError, which exits the program after calling Usage.
	Usage func()

	name          string
	parsed        bool
	actual        map[string]*Flag
	formal        map[string]*Flag
	args          []string // arguments after flags
	errorHandling ErrorHandling
	output        io.Writer // nil means stderr; use Output() accessor
}

type ErrorHandling int
type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}*/

/*func (f *FlagSet) Var(value Value, name string, usage string) {
	// Flag must not begin "-" or contain "=".
	if strings.HasPrefix(name, "-") {
		panic(f.sprintf("flag %q begins with -", name))
	} else if strings.Contains(name, "=") {
		panic(f.sprintf("flag %q contains =", name))
	}

	// Remember the default value as a string; it won't change.
	flag := &Flag{name, usage, value, value.String()}
	_, alreadythere := f.formal[name]
	if alreadythere {
		var msg string
		if f.name == "" {
			msg = f.sprintf("flag redefined: %s", name)
		} else {
			msg = f.sprintf("%s flag redefined: %s", f.name, name)
		}
		panic(msg) // Happens only if flags are declared with identical names
	}
	if f.formal == nil {
		f.formal = make(map[string]*Flag)
	}
	f.formal[name] = flag
}
*/
// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-
/* 翻译
Var用指定的名称和用法字符串定义一个标志。
标志的类型和值由第一个value类型的参数表示，该参数通常保存value的用户定义实现。
例如，调用者可以创建一个标志，将逗号分隔的字符串转换为字符串的slice，
方法是给slice Value;特别地，Set将把逗号分隔的字符串分解为片。
 */