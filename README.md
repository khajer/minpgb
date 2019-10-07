# minpgb
minpgb is a mini progress bar cli 

## install 
``` go get -u github.com/khajer/minpgb ```

## example
```
package main

import (
	pgb "github.com/khajer/minpgb"
	"time"
)

func main(){
	pb := pgb.New()
	pb.Total = 100

	for i:=0; i< 100; i++{
		curr := pb.GetCurrent()
		pb.SetCurrent(curr+1)
		time.Sleep(100 * time.Millisecond)
	}

}
```

if you want to clear progress bar when 100% completed 
```
	pb.Flush()
```
example
```
package main

import (
	pgb "github.com/khajer/minpgb"
	"time"
	"fmt"
)

func main(){
	pb := pgb.New()
	pb.Total = 100

	for i:=0; i< 100; i++{
		curr := pb.GetCurrent()
		pb.SetCurrent(curr+1)
		time.Sleep(100 * time.Millisecond)
	}
	pb.Flush()
	fmt.Println("100 percent completed")

}
```


## screenshot

![alt text](https://user-images.githubusercontent.com/797258/66182419-354df200-e69f-11e9-88cb-9a339a81f7e0.png)
