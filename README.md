# json-summarize

Merge compatible array elements to sumarize a JSON object.

## example

```
$ cat test.json
{
	"a" : [
		1,
		2
	],
	"b" : [
		{"b1" : 1},
		{"b2" : 2}
	]
}
$ cat example.json | go run main.go
{"a":[1],"b":[{"b1":1,"b2":2}]}
```

