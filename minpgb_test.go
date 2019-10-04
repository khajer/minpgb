package minpgb

import (
	"testing"
	"time"
	"fmt"
)


func TestSimpleProgressBar(t *testing.T){
	pgb := New()
	if pgb != nil{
	}

	pgb.Total = 20
	pgb.SetCurrent(0)

	for i:=0; i< 20; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+1)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("completed")

}