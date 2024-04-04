package main

import "fmt"

type callable interface{
    arity() int
    call(interp *Interpreter, arguments []interface{}) interface{}
}

type RuntimeFunction struct{
    callable
    declaration Function
}

func (f RuntimeFunction) call(interp *Interpreter, arguments []interface{}) (returnVal interface{}) {

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

    for i, param := range f.declaration.params {
        interp.env[param.lexeme] = arguments[i]
    }

    interp.executeBlock(f.declaration.body)//, interp.env)
    return nil

}

func (f RuntimeFunction) arity() int {
   return len(f.declaration.params)
}
