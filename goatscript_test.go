package goatscript

import (
	"fmt"
	goatdata "goatscript/ast"
)

// https://blog.golang.org/examples

func Example() {
	script := goatdata.Condition{
		Test: goatdata.IntLessThan{
				A: goatdata.IntVariable{ "var1" },
				B: goatdata.Int{ 2 },
			},
		True: goatdata.IntIncrement{
				Variable: "var1",
				Amount: 1,
			},
		False: goatdata.NoOp{},
	}

	templ, err := Precompile(script)
	if err != nil {
		panic(err)
	}

	inst1 := templ.CreateInst()
	inst1.SetIntVariable("var1", 1)
	inst1.Run()
	var1, _ := inst1.GetIntVariable("var1")
	fmt.Println(var1)

	inst2 := templ.CreateInst()
	inst2.SetIntVariable("var1", 6)
	inst2.Run()
	var1again, _ := inst2.GetIntVariable("var1")
	fmt.Println(var1again)

	// Output:
	// 2
	// 6
}
