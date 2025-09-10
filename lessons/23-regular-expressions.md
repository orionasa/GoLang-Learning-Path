# 23. Regular Expressions in Go

Regular expressions are a powerful tool for matching patterns in text. Go provides a built-in `regexp` package for working with regular expressions.

## What is a Regular Expression?

A regular expression (or "regex") is a sequence of characters that defines a search pattern. They are used to find, match, and manipulate strings.

## The `regexp` Package

The `regexp` package in Go implements regular expression search.

### Compiling a Regular Expression

The first step in using a regular expression is to compile it. You can do this using the `regexp.Compile()` function. Compiling a regex creates a `regexp.Regexp` object that can be used for matching.

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    // Compile a regex to find numbers in a string.
    re, err := regexp.Compile("[0-9]+")
    if err != nil {
        // handle error
    }
    // ...
}
```

If the regular expression is a constant, you can use `regexp.MustCompile()`. This function will panic if the expression fails to compile, which is useful for global variables that are initialized at the start of your program.

```go
var re = regexp.MustCompile("[0-9]+")
```

## Matching Text

Once you have a compiled `regexp.Regexp` object, you can use its methods to match text.

### `MatchString()`

The `MatchString()` method returns a boolean indicating whether the regex matches the string.

```go
matched := re.MatchString("abc123def") // returns true
```

### `FindString()`

The `FindString()` method returns the first substring that matches the regex.

```go
match := re.FindString("abc123def456") // returns "123"
```

### `FindAllString()`

The `FindAllString()` method returns a slice of all successive matches of the regex.

```go
matches := re.FindAllString("abc123def456", -1) // returns ["123", "456"]
```
The second argument is the number of matches to find. A value of -1 means to find all matches.

## Submatches

You can use parentheses in a regular expression to create capturing groups. This allows you to extract specific parts of a match.

The `FindStringSubmatch()` method returns a slice of strings holding the text of the leftmost match and the matches for its subexpressions.

```go
re := regexp.MustCompile("([a-z]+) ([0-9]+)")
submatches := re.FindStringSubmatch("hello 123")
// submatches will be ["hello 123", "hello", "123"]
```

The first element of the slice (`submatches[0]`) is the full match. The subsequent elements are the matches for the capturing groups.

## Replacing Text

The `regexp` package also provides functions for replacing text that matches a pattern.

### `ReplaceAllString()`

The `ReplaceAllString()` method returns a copy of the string, replacing all matches of the regex with the provided replacement string.

```go
re := regexp.MustCompile("a[a-z]+")
replaced := re.ReplaceAllString("an apple a day", "the")
// replaced will be "the apple the day"
```

### `ReplaceAllStringFunc()`

The `ReplaceAllStringFunc()` method allows you to use a function to generate the replacement string for each match.

```go
import "strings"

re := regexp.MustCompile("a[a-z]+")
replacedFunc := re.ReplaceAllStringFunc("an apple a day", strings.ToUpper)
// replacedFunc will be "AN APPLE A DAY"
```

Regular expressions are a powerful and flexible tool for text processing. The `regexp` package in Go provides a comprehensive and easy-to-use interface for working with them.
