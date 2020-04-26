# fistandards
fistandards is a set of methods to validate a number of banking and financial coding standards

To use you initialize a Checker:

```go
var ibanChecker iban.Checker
if err := ibanChecker.Initialize(); err != nil {
log.Fatal("something went wrong: ", err.Error())
```

And you then verify a string:
```go
countryCode, cleanedIban, matches := ibanChecker.Validate("AT-123456789012345678", true)
```

See `main.go` and the tester packages for more examples.
