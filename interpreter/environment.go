package interpreter

type Environment struct{
    enclosing *Environment
    values map[string]interface{}
}

func NewEnvironment(enclosing *Environment) *Environment{
    environment := new(Environment)
    environment.values = make(map[string]interface{})
    environment.enclosing = enclosing
    return environment
}

func (env Environment) Get(name string) interface{} {

    if localValue, ok := env.values[name]; ok {
        return localValue
    }

    //Look recursively into parent scope
    if env.enclosing != nil {
        return env.enclosing.Get(name)
    }

    return nil
    //panic("Undefined variable '" + name+ "'")
}


func (env Environment) Define(name string, value interface{}) {

    env.values[name] = value
}

func (env Environment) Assign(name string, value interface{}) {

    if _, ok := env.values[name]; ok {
        env.values[name] = value
        return
    }

    if env.enclosing != nil {
        env.enclosing.Assign(name, value)
        return
    }

    panic("Undefined variable '" +name+"'")
}
