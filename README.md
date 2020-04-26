# fistandards
fistandards is an implementation of a number of banking standards.

Here's some example code:

```go
var iban iban.Iban
var normalized string
var ok bool

if normalized, ok = iban.Set("GB82-WEST 1234 5698 7654 32"); !ok {
	fmt.Printf("%v is not a valid IBAN\n", normalized)
} else {
	fmt.Printf("%v is a valid IBAN\n", normalized)
}
```

See `main.go` and the tester packages for more examples.
