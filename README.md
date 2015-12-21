## Stache

Compile a mustache template using ENV as the context.

```
# Download and make executable
$ curl -o stache -L https://github.com/bshelton229/stache/blob/master/build/stache.linux.amd64\?raw\=true && chmod 755 stache

# Pipe input
$ echo "My path is: {{ PATH }}" | ./stache

# Read from a file
$ echo "My username is {{ USER }}" > test.mustache
$ ./stache -f test.mustache > test.txt
```
