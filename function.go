package main

import "fmt"

type callable interface{
    arity() int
    call(interp *Interpreter, arguments []interface{}) interface{}
}

type RuntimeFunction struct{
    callable
    declaration Function
    closure *Environment
}

func (f RuntimeFunction) call(interp *Interpreter, arguments []interface{}) (returnVal interface{}) {

    funcEnv := NewEnvironment(interp.globals)

    //We catch the function return in a panic
    defer func() {
        if err := recover(); err != nil {
            if v, ok := err.(ReturnValue); ok {

                returnVal = v.value
                return
            }
            fmt.Println("ERROR: DID NOT RETURN")
            fmt.Println(err)
            panic(err)
        }
    }()

    //callEnv := NewEnvironment(f.closure)
    //callEnv := interp.environment

    for i, param := range f.declaration.params {
        //interp.environment[param.lexeme] = arguments[i]
        funcEnv.Define(param.lexeme, arguments[i])
    }

    interp.executeBlock(f.declaration.body, funcEnv)
    return nil

}

func (f RuntimeFunction) arity() int {
   return len(f.declaration.params)
}
