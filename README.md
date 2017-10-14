# rsraid

Inelegant quick application to call utamaro/rsraid. Yay, my first (and LAST?!)
go program

## Build

```
go get github.com/andyneff/rsraid
go build github.com/andyneff/rsraid
```

## Encode

```
ls / > test.txt
rsraid -n 2 -m 1 -e test.txt
```

## Decode

```
rsraid -n 2 -d -s <size of original file> -o recombined.txt test_0.txt test_1.txt
```

## Bugs

Well YEAH! The underlining library does it wrong. For example a 1299 size file
split with `-n 2` breaks (it drops two bytes), but size 1300 works.