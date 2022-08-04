# get-testable-pkg command

Get your testable packages in your repo.  
Maybe, your incomplete test code runs a little bit faster?

## Usage
```shell
$ go install github.com/daichitakahashi/get-testable-pkg
$ go test `get-testable-pkg -v "exclude-pattern1" "exclude-pattern2"`
```

## Options
| option           | description                                                                       |
|------------------|-----------------------------------------------------------------------------------|
| `-v`             | output non-testable packages and other info                                       |
| `-shuffle`       | shuffle testable package (default is true)                                        |
| exclude patterns | patterns that matches excluding paths.<br>pattens are compiled as `regexp.Regexp` |
