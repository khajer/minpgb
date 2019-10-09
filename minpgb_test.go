package minpgb

import (
	"testing"
	"time"
	"fmt"
	// "golang.org/x/crypto/ssh/terminal"
	// "golang.org/x/sys/unix"
	// "os"
)

func TestSimpleProgressBarType0(t *testing.T){
	pgb := New()
	if pgb != nil{
	}
		

	pgb.Total = 100;
	for i:=0; i< 10; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+10)
		time.Sleep(100 * time.Millisecond)
	}
	pgb.Flush()
	fmt.Println("Completed")

	pgb.SetStyle(PGTYPE_DASH)
	for i:=0; i< 10; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+10)
		time.Sleep(100 * time.Millisecond)
	}
	pgb.Flush()
	fmt.Println("completed")
}

func TestSimpleProgressBar(t *testing.T){
	pgb := New()
	if pgb != nil{
	}

	pgb.Total = 100;
	for i:=0; i< 10; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+10)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("completed")
}

func TestSimpleProgressBarFlush(t *testing.T){
	pgb := New()
	if pgb != nil{
	}	
	
	pgb.Total = 100;
	for i:=0; i< 100; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+1)
		time.Sleep(100 * time.Millisecond)
	}
	pgb.Flush()
	fmt.Println("completed")
}


func TestGetWinsize(t *testing.T){
	ws := GetWinsize()
	if ws == nil{
		// t.Errorf("TestGetWinsize Fails ")
	}
	// fmt.Println(ws)
}

func TestProgressBarLength(t *testing.T){
	for i:=0; i<100; i++{
		str := CreateProgressText(float64(i), 100, 100)	
		if len(str) != 102 {
			t.Errorf("fails ")
		}
	}	

	// str := CreateProgressText(71,100, 100)
	// fmt.Println(str, len(str))
}

// func TestCallTextAppend(t *testing.T){
// 	for i:=0; i<100; i++{
// 		d := CallTextAppend(float64(50), float64(i))
// 	}	
// }

func TestPercentTextAppend(t *testing.T){
	if 50 != CallTextAppend(100, 50){
		t.Errorf("Calculate Value Fails ")
	}
	if 21 != CallTextAppend(42, 50){
		t.Errorf("Calculate Value Fails ")
	}

	if 42 != CallTextAppend(84, 50){
		t.Errorf("Calculate Value Fails ")
	}
}

func TestProgressBarMore500(t *testing.T){
	pgb := New()
	if pgb != nil{
	}

	pgb.Total =  500
	pgb.SetCurrent(0)
	for i:=0; i< 53; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+10)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("completed")
}

func TestProgressBarMore501AndWithFush(t *testing.T){
	pgb := New()

	if pgb != nil{
	}

	pgb.Total =  500
	pgb.SetCurrent(0)
	for i:=0; i< 53; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+10)
		time.Sleep(100 * time.Millisecond)
	}
	pgb.Flush()
	fmt.Println("completed")
}

func TestCreateProgressText(t *testing.T){
	txtLen100percent := 50.00
	s := CreateProgressText(0, 100, txtLen100percent)
	if len(s) != int(txtLen100percent)+2{
		t.Errorf("Length txt error , curr len = %d", len(s))
	}
	txtLen100percent = 30.00
	s = CreateProgressText(0, 100, txtLen100percent)
	if len(s) != int(txtLen100percent)+2{
		t.Errorf("Length txt error , curr len = %d", len(s))
	}
}