package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"syscall/js"

	"blueprint" // Import the blueprint package from LayerForge
)

// Wrapper function to expose Blueprint methods to JavaScript
func methodWrapper(bp *blueprint.Blueprint, methodName string) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		method := reflect.ValueOf(bp).MethodByName(methodName)
		if !method.IsValid() {
			return fmt.Sprintf("Method %s not found", methodName)
		}

		// Assume JSON input; parse args[0] as JSON if provided
		if len(args) == 0 {
			return "No parameters provided"
		}
		paramJSON := args[0].String()
		var paramValues []interface{}
		if err := json.Unmarshal([]byte(paramJSON), &paramValues); err != nil {
			return fmt.Sprintf("Invalid JSON input: %v", err)
		}

		// Prepare reflect.Value inputs based on expected parameter types
		inputs := make([]reflect.Value, len(paramValues))
		for i, param := range paramValues {
			inputs[i] = reflect.ValueOf(param)
		}

		// Call the method and capture results
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
	// Create Blueprint instance
	bp := &blueprint.Blueprint{}

	// Retrieve all methods using reflection and expose them
	methods, err := bp.GetBlueprintMethods()
	if err != nil {
		fmt.Println("Error getting methods:", err)
		return
	}

	for _, methodInfo := range methods {
		js.Global().Set(methodInfo.MethodName, methodWrapper(bp, methodInfo.MethodName))
	}

	// Keep the WASM program running
	select {}
}
