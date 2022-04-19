Simple Leveled Logger
=====================

`github.com/andrewslotin/llog` is a library that provides leveled logs to the Go standard library's `log.Logger`.

Installation
------------

```bash
go get github.com/andrewslotin/llog
```

Usage
-----

The leveled logger is implemented as an `io.Writer` to be provided to the stdlib's `log.SetOutput()`. The writer determines log record level by looking at the message prefix. For example, a message that starts with `warn` will be considered as a `llog.WarnLevel` and sent to the underlying writer, if the writer has the minimal level set to `llog.WarnLevel` or below.

```
// configure standard logger to write warning messages and above to STDERR
log.SetOutput(llog.NewWriter(os.Stderr, llog.WarnLevel))

log.Printf("warn: failed to close the output file: %s", err) // warn level, will be printed out
log.Println("something just happened") // info level, will be ignored
```

If a message does not contain any known prefixes, its level is considered to be `llog.InfoLevel`.

Please consult [the package documentation][godoc] for the list of available log levels as well as how they are mapped to the message prefixes.

[godoc]: https://pkg.go.dev/github.com/andrewslotin/llog
