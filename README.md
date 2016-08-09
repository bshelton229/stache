## Stache

A CLI tool for rendering mustache templates using ENV as the context.

### Pipe input

```
$ echo "My home directory is: {{ HOME }}" | stache
```

### Read from a file

```
$ echo "My username is {{ USER }}" > test.mustache
$ stache -f test.mustache > test.txt
```

### Write to an output file

```
$ stache -f test.mustace -o test.txt
```
