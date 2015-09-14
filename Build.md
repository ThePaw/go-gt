#Build instructions

# Build Instructions #

At the moment the code can be compiled with the Go tip. The differences to the Release are mostly minor, such as the package "rand" is in tip known as "math/rand", and the type "os.Error" is simply "error".

In Unix environments use the Makefiles and in Plan 9 we have the mkfile using the Plan 9 from http://code.google.com/p/go-plan9/.