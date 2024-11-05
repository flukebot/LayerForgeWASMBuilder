package main

import (
    "encoding/json"
    "fmt"
    "reflect"
    "syscall/js"
    "blueprint" // Adjust the import path if necessary
)

// methodWrapper dynamically wraps each method of Blueprint to expose it to JavaScript
func methodWrapper(bp *blueprint.Blueprint, methodName string) js.Func {
    return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        method := reflect.ValueOf(bp).MethodByName(methodName)
        if !method.IsValid() {
            return fmt.Sprintf("Method %s not found", methodName)
        }

        methodType := method.Type()
        numParams := methodType.NumIn()

        var inputs []reflect.Value

        if numParams > 0 { // Method expects parameters
            if len(args) == 0 {
                return fmt.Sprintf("Method %s requires parameters but none were provided", methodName)
            }

            paramJSON := args[0].String()
            var paramValues []interface{}
            if err := json.Unmarshal([]byte(paramJSON), &paramValues); err != nil {
                return fmt.Sprintf("Invalid JSON input: %v", err)
            }

            // Prepare reflect.Value inputs based on expected parameter types
            inputs = make([]reflect.Value, numParams)
            for i := 0; i < numParams; i++ {
                if i < len(paramValues) {
                    inputs[i] = reflect.ValueOf(paramValues[i])
                } else {
                    inputs[i] = reflect.Zero(methodType.In(i)) // Provide zero value for missing params
                }
            }
        }

        // Call the method with or without inputs, depending on numParams
        results := method.Call(inputs)

        // Format the results as JSON for output
        output := make([]interface{}, len(results))
        for i, result := range results {
            output[i] = result.Interface()
        }
        resultJSON, _ := json.Marshal(output)
        return string(resultJSON)
    })
}

func main() {
    // Create an instance of Blueprint
    bp := &blueprint.Blueprint{}

    // Retrieve all methods using reflection from introspection
    methods, err := bp.GetBlueprintMethods()
    if err != nil {
        fmt.Println("Error getting methods:", err)
        return
    }

    // Expose each method as a JavaScript function
    for _, method := range methods {
        js.Global().Set(method.MethodName, methodWrapper(bp, method.MethodName))
    }

    // Keep the WebAssembly program running
    select {}
}
