package main

import (
	"fmt"
)

type Any interface{}
type evalFunc func(Any) (Any, Any)

func buildLazyEvaluator(evalFunc evalFunc, initialState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var start Any = initialState
		var returnValue Any
		for {
			returnValue, start = evalFunc(start)
			retValChan <- returnValue
		}
	}

	retFunc := func() Any {
		return <-retValChan
	}

	go loopFunc()
	return retFunc

}

func buildLazyIntEvaluator(evalFunc evalFunc, initialState Any) func() int {
	consume := buildLazyEvaluator(evalFunc, initialState)
	return func() int {
		return consume().(int)
	}
}

func main() {
	evenNum := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}

	even := buildLazyIntEvaluator(evenNum, 0)
	for i := 1; i < 10; i++ {
		fmt.Printf("%v even number is %v\n", i, even())
	}

}
