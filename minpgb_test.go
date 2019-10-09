package minpgb

import (
	"testing"
	"time"
	"fmt"	
)

func TestSimpleProgressBarType0(t *testing.T){
	pgb := New()
	if pgb != nil{
	}		

	pgb.Total = 1000;
	for i:=0; i< 100; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+10)
		time.Sleep(10 * time.Millisecond)
	}
	pgb.Flush()
	fmt.Println("Completed")

	pgb.SetStyle(PGTYPE_DASH)
	for i:=0; i< 100; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+10)
		time.Sleep(10 * time.Millisecond)
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
		time.Sleep(10 * time.Millisecond)
	}
	pgb.Flush()
	fmt.Println("completed")
}


func TestGetWinsize(t *testing.T){
	ws := GetWinsize()
	if ws == nil{
		// t.Errorf("TestGetWinsize Fails ")
	}
	
}

func TestProgressBarLength(t *testing.T){
	for i:=0; i<100; i++{
		str := CreateProgressText(float64(i), 100, 100)	
		if len(str) != 102 {
			t.Errorf("fails ")
		}
	}	
}

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
		time.Sleep(10 * time.Millisecond)
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
		time.Sleep(10 * time.Millisecond)
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