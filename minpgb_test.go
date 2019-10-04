package minpgb

import (
	"testing"
	// "github.com/stretchr/testify/assert"
	"time"
)


func TestSimpleProgressBar(t *testing.T){
	pgb := New()
	if pgb != nil{
		// assert.Equal(t, "x", "x")
	}

	pgb.Total = 20
	pgb.SetCurrent(0)

	for i:=0; i< 20; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+1)
		time.Sleep(1 * time.Second)
	}

}