package minpgb

import (
	"fmt"
)

var pgb *MinPgb

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

	fmt.Print("\r\033[K")	
	percent := float64(pgb.Curr)/float64(pgb.Total)*100	
	s := CreateProgressText(pgb.Curr, pgb.Total)
	fmt.Printf("[%d/%d] %s %.2f%s", pgb.Curr, pgb.Total, s, percent, "%")
	
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
