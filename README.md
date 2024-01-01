## Table of Contents
- [Overview](#overview)
- [How to run](#how_to_run)
    - [App](#app)
    - [Using as a library](#library)
    - [Tests](#tests)
## Overview <a name="#overview"></a>

This app takes time string in format 12:01:00AM(AM/PM can be uppercase or lowercase) and returns time string in format 00:01:00.
It will not accept any other time format and will throw an error and end the execution if the format is wrong.
It was unit tested in parallel by table driven unit test.

## How to run <a name="how_to_run"></a>

### App <a name="app"></a>

Clone the repository:

    git clone https://github.com/sosshik/work-time.git

Run app(instead of 12:01:00AM you can put your time string as argument, but in the same format):

    go run time_format.go -time=12:01:00AM  

### Using as a library <a name="library"></a> 

`go get` it:

    go get github.com/sosshik/work-time

And import to your project to use the function:

    import "github.com/sosshik/work-time"

### Tests <a name="tests"></a> 

Clone the repository:

    git clone https://github.com/sosshik/work-time.git

Run the tests:

    go test -v

Run benchmark:

    go test -bench=. -benchmem

