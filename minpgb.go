package minpgb

import (
	"fmt"
)

var pgb *MinPgb

const (
	MAX_PERCENT float64 = 100
	CH_RESET_LINE string = "\r\033[K"	
)
	


type MinPgb struct{
	Curr 	int
	Total 	int
}
func init(){
	pgb = New()
}

func New() *MinPgb{
	pg := new(MinPgb)	
	return pg
}

func (pgb *MinPgb)GetCurrent() int{
	return pgb.Curr
}
func (pgb *MinPgb)SetCurrent(curr int){
	pgb.Curr = curr

	fmt.Print(CH_RESET_LINE)	
	percent := float64(pgb.Curr)/float64(pgb.Total)*MAX_PERCENT	
	s := CreateProgressText(pgb.Curr, pgb.Total)
	sEnd := "%"
	if percent >= MAX_PERCENT{
		sEnd += "\n"
	}
	fmt.Printf("[%d/%d] %s %.2f%s", pgb.Curr, pgb.Total, s, percent, sEnd)
	
	
}
func CreateProgressText(curr int, total int) string{
	s := "["
	for i:=1; i<= total; i++ {
		if i <= curr{
			s += "#"	
		}else{
			s += " "
		}		
	}
	s += "]"
	return s
}
