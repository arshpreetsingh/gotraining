// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to store and retrieve
// values from a context.
package main

import (
	"context"
	"fmt"
)

// TraceID is represents the trace id.
type TraceID string

// TraceIDKey is the type of value to use for the key. The key is
// type specific and only values of the same type will match.
type TraceIDKey int

func main() {

	// Create a traceID for this request.
	traceID := TraceID("f47ac10b-58cc-0372-8567-0e02b2c3d479")

	// Declare a key with the value of zero of type userKey.
	var traceIDKey TraceIDKey = 0

	// Store the traceID value inside the context with a value of
	// zero for the key type.
	ctx := context.WithValue(context.Background(), traceIDKey, traceID)

	// Retrieve that traceID value from the Context value bag.
	if uuid, ok := ctx.Value(traceIDKey).(TraceID); ok {
		fmt.Println("TraceID:", uuid)
		fmt.Println("this is anoter stuf we need to look for as well")
		fmt.Println(ctx.Value(traceIDKey))
		fmt.Println("////////////////////////")
		fmt.Println(ctx.Value(traceIDKey).(TraceID))
	}

	// Retrieve that traceID value from the Context value bag not
	// using the proper key type.
	// if _, ok := ctx.Value(0).(TraceID); !ok {
	// 	fmt.Println("TraceID Not Found")
	// 	fmt.Println("this is ok", ok)
	// 	//	fmt.Println("this is anoter Moreeeeeeeeee   stuff we need to look for as well")
	// 	//	fmt.Println(ctx.Value(traceIDKey).(TraceID))
	// }
}
