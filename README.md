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

The leveled logger is implemented as an `io.Writer` to be provided to the stdlib's `log.SetOutput()`. The writer determines log record level by looking at the message prefix. For example, a message that starts with `warn` will be considered as a `llog.WarnLevel` and sent to the underlying writer, if the writer has the minimal level set to `llog.WarnLevel` or lower.

```
// configure standard logger to write warning messages and above to STDERR
log.SetOutput(llog.NewWriter(os.Stderr, llog.WarnLevel))

log.Printf("warn: failed to close the output file: %s", err) // warn level, will be printed out
log.Println("something just happened") // info level, will be ignored
```

If a message does not contain any known prefixes, its level is considered to be `llog.InfoLevel`.

Please consult [the package documentation][godoc] for the list of available log levels as well as how they are mapped to the message prefixes.

Performance
-----------

There is nearly no overhead in using the `llog.Writer` to write log messages.

Below is the `benchstat` output comparing the benchmarks of `log.Println()` with `io.Discard` and `llog.Writer` set as outputs. The `llog.Writer` is configured to write to `io.Discard` as well with the minimal log level sufficient to print out all messages used in the benchmark.

```
name            old time/op    new time/op    delta
Writer_Write-4     306ns ± 0%     362ns ± 0%   ~     (p=1.000 n=1+1)

name            old alloc/op   new alloc/op   delta
Writer_Write-4     16.0B ± 0%     16.0B ± 0%   ~     (all equal)

name            old allocs/op  new allocs/op  delta
Writer_Write-4      1.00 ± 0%      1.00 ± 0%   ~     (all equal)
```

[godoc]: https://pkg.go.dev/github.com/andrewslotin/llog
