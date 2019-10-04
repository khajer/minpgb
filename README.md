# minpgb
minpgb is a mini progress bar cli 

## example
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

}
```
