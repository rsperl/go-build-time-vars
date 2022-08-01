# Build-Time Vars with `go:generate`

## The `go:generate` comment

The generate comment must start at the beginning of the line with no space between `//` and `go:generate`:

```go
//go:generate ...
```

This comment will be treated like an ordinary comment (ignored)

```go
   //go:generate ...
```

## Generating Files to Embed

When `go generate .` is run, go executes the command after the `go:generate` by treats the first word as the program and all other tokens after that as arguments to the program. It does not recognize `>` as shell redirection, since this is not being run by the shell. To get around this, use `bash` as the program to call, and put the actual command quoted in `-c`:

```go
//go:generate bash -c "date --iso-8601=seconds > tmp/build_time"
```

## Stripping Newlines

Because `go:embed` adds a new line to whatever it reads in, the following pattern is used:

```go
// 1. Create the file in a file with the same name as the var that receives the value
// 2. embed in a lowercase variable
// 3. Set the titlecase variable to the trimmed result of the lowercase variable

//go:generate bash -c "date > tmp/varName"
//go:embed tmp/varName
var varName string

// VarName with upper case first letter is the real var, but
// the lowercase var actually receives the embed value
var VarName = strings.TrimRight(varName, "\n")
```
