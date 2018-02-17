# envLoader

[documentation](https://godoc.org/github.com/nikosEfthias/envLoader).
**Env loader** is a simple package that help you load custom environment variables from a csv file.


`.env`:

```csv
a,b
c,d
e,f
```
`main.go`:

```go
package main
import "fmt"
import "log"
import "os"
import "github.com/nikosEfthias/envLoader"

func main(){
//by default check for .env file but you can specify a custom filename
	if err:=envLoader.Load();nil!=err{
		log.Fatalln(err)
	}
//here we can use all the variables we specified in `.env` file

os.GetEnv("a")//b
}

```
