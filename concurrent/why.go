package concurrent

// Why build concurrency on the ideas of CSP? https://golang.org/doc/faq#csp
// Why goroutines instead of threads? https://golang.org/doc/faq#goroutines

// Why are map operations not defined to be atomic? read often and update slow down , few safety  https://golang.org/doc/faq#atomic_maps
// key not comparable will cause panic https://github.com/golang/go/issues/23734

//   Why is there no goroutine ID? https://golang.org/doc/faq#no_goroutine_id
//   For those cases where a particular goroutine is truly special, the language provides features such as channels that can be used in flexible ways to interact with it.

//  Why doesn't my program run faster with more CPUs?  https://golang.org/doc/faq#parallel_slow
//intrinsically sequential   pipeline may be dramatically
