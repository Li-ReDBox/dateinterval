# Calculate days between two dates from scratch

The request is to calculate the distance in whole days between two dates, counting only
the days in between those dates, i.e. 01/01/2001 to 03/01/2001 yields “1”. The valid
date range is between 01/01/1900 and 31/12/2999, all other dates should be
rejected.

## Usages
### Compile an executable from cmd/main.go
```shell
go build -o demo
./demo 2/6/1983 22/6/1983
./demo 4/7/1984 25/12/1984
./demo 3/1/1989 3/8/1983
./demo 1/3/1989 3/8/1983
```
The output should look like:
```shell
Days between 2/6/1983 and 22/6/1983 is 19
Days between 4/7/1984 and 25/12/1984 is 173
Days between 3/1/1989 and 3/8/1983 is 1979
Days between 1/3/1989 and 3/8/1983 is 2036
```

### Use this as a package
```shell
mkdir verifier
cd verifier
go mod init funmech/verifier
go get github.com/Li-ReDBox/dateinterval
```

In main.go:
```go
import "github.com/Li-ReDBox/dateinterval"

...

func main() {
    d, err := dateinterval.CreateDate("1/1/2000")
    ...
}
```
