package factory

import (
	"fmt"
)

//Any is nickname of interface
type Any interface{}

//EvalFunc is nickname of factory func
type EvalFunc func(Any) (Any, Any)

//Fib to store fibnachi value
type Fib struct {
	front int
	end   int
}

//BuildLazyFibEvaluator is Fib lazy builder
func BuildLazyFibEvaluator(evalFunc EvalFunc, initState Any) func() Fib {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() Fib {
		return ef().(Fib)
	}
}

//BuildFibEvaluator is Fib Evaluator Test
func BuildFibEvaluator() {
	evalFunc := func(state Any) (Any, Any) {
		os := state.(Fib)
		var ns Fib
		ns.front = os.end
		ns.end = os.end + os.front
		return os, ns
	}
	even := BuildLazyFibEvaluator(evalFunc, Fib{0, 1})
	for i := 0; i < 30; i++ {
		fmt.Println(even().end)
	}
}

//BuildEvenEvaluator is int Evaluator Test
func BuildEvenEvaluator() {
	evalFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}
	even := BuildLazyIntEvaluator(evalFunc, 0)
	for i := 0; i < 10; i++ {
		fmt.Println(even())
	}
}

//BuildLazyIntEvaluator is int Evaluator Builder
func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}

//BuildLazyEvaluator is Factory
func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <-retValChan
	}

	go loopFunc()
	return retFunc
}
