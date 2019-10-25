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

	pgb.SetStyle(PGTYPE_BEER)
	pgb.Total = 100;
	for i:=0; i< 10; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+10)
		time.Sleep(100 * time.Millisecond)
	}
	pgb.End()
	fmt.Println("progress completed")	
}

func TestSimpleProgressBarType1(t *testing.T){
	pgb := New()
	if pgb != nil{
	}		

	pgb.Total = 100;
	for i:=0; i< 10; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+10)
		time.Sleep(100 * time.Millisecond)
	}
	pgb.End()
	fmt.Println("Completed")	
}
func TestSimpleProgressBarType2(t *testing.T){
	pgb := New()
	if pgb != nil{
	}		

	pgb.SetStyle(PGTYPE_BLOCK2)
	pgb.Total = 100;
	for i:=0; i< 10; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+10)
		time.Sleep(100 * time.Millisecond)
	}
	pgb.End()
	fmt.Println("Completed")	
}

func TestSimpleProgressBar(t *testing.T){
	pgb := New()
	if pgb != nil{
	}
	pgb.Total = 100;
	pgb.SetStyle(PGTYPE_BLOCK)
	for i:=0; i< 100; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+1)
		time.Sleep(10 * time.Millisecond)
	}
	pgb.Flush()
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
		time.Sleep(1 * time.Millisecond)
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

func TestCallTextAppend(t *testing.T){
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
	pgb.SetStyle(PGTYPE_NORMAL)
	pgb.Total =  500
	pgb.SetCurrent(0)
	for i:=0; i< 53; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+10)
		time.Sleep(1 * time.Millisecond)
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
		time.Sleep(1 * time.Millisecond)
	}
	pgb.Flush()
	fmt.Println("completed")
}

func TestCreateProgressText(t *testing.T){
	txtLen100percent := 50
	borderBeginEnd := 2
	s := CreateProgressText(0, 100, txtLen100percent)
	if len(s) != txtLen100percent+borderBeginEnd{
		t.Errorf("Length txt error , curr len = %d", len(s))
	}
	txtLen100percent = 30.00
	s = CreateProgressText(0, 100, txtLen100percent)
	if len(s) != txtLen100percent+borderBeginEnd{
		t.Errorf("Length txt error , curr len = %d", len(s))
	}
}

func TestCreatePreLoadingText(t *testing.T){
	
	str := CreatePreLoadingText("readme.txt", 100, 100)
	if str != "readme.txt [100/100]" {
		t.Errorf("String Prefix error : '%s'", str)
	}
}

func TestShowPreloadingText(t *testing.T){
	pgb := New()
	if pgb != nil{
	}
	filename := "readme.txt"
	pgb.SetPreText(filename)
	pgb.SetStyle(PGTYPE_ARROW)
	pgb.Total =  100
	pgb.SetCurrent(0)
	for i:=0; i< 100; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+1)
		time.Sleep(10 * time.Millisecond)
	}
	pgb.Flush()
	fmt.Printf("'%s': download completed\n", filename)
	filename = "main.go"

	pgb.SetPreText(filename)
	for i:=0; i< 100; i++{
		curr := pgb.GetCurrent()
		pgb.SetCurrent(curr+1)
		time.Sleep(10 * time.Millisecond)
	}
	pgb.Flush()
	fmt.Printf("'%s': download completed\n", filename)

}