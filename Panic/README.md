# Panic impl

This is a simple Go program that demonstrates the use of the panic and recover functions to handle unexpected errors.

The program defines a main function that calls a `divideByZero` function. The `divideByZero` function attempts to divide 1 by 0, which would normally cause a runtime error. However, the `divideByZero` function uses the `defer` keyword to call a function that uses the recover function to catch the panic and log the error message, allowing the program to continue executing.

The `divide` function is also defined which will check if the denominator is 0 and panic if it is 0.

Note that the `recover` function should only be used in deferred functions, as it only works when called directly from the goroutine that panicked.
