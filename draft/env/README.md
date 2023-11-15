# Reading from environment variables

## Motivation
This environment variable uses Go annotations, but I feel that is overkill compared to simple functions.
https://github.com/caarlos0/env


## Notes and Learnings by reading caarlos0
* unset sensitive envirionment variables after having read them into the config
* I feel the package does too much: envSeparator, fromfile
* Required should also mean not emtpy
*

## Questions
* How is it used in real life?
  * Port is an env string, converted to integer, with a default value provided
  * Password is an env string, must exit if not provided
* What is the recommended Go way to wrap errors and print them again?
* How to cast variable to a generic type?
  `animal,_ := reflect.ValueOf(cat).Interface().(T)`
  When the value's type is only known at runtime, the  reflect.Value  API must be used.
