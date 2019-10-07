package minpgb

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
)

var pgb *MinPgb
var winsize *unix.Winsize

const (
	MAX_PERCENT float64 = 100
	CH_RESET_LINE string = "\r\033[K"	
)
	


type MinPgb struct{
	Curr 	float64
	Total 	float64
}
func init(){
	pgb = New()
	winsize = GetWinsize()
}

func New() *MinPgb{
	pg := new(MinPgb)	
	return pg
}

func (pgb *MinPgb)GetCurrent() float64{
	return pgb.Curr
}
func (pgb *MinPgb)SetCurrent(curr float64){
	pgb.Curr = curr

	fmt.Print(CH_RESET_LINE)	
	currPercent := pgb.Curr/pgb.Total*MAX_PERCENT	
		
	strHead := fmt.Sprintf("[%.0f/%.0f] ", pgb.Curr, pgb.Total)
	strEnd := fmt.Sprintf(" %.2f%s", currPercent, "%")
	pgbWidth := int(winsize.Col) - (len(strHead)+len(strHead))

	sProgress := CreateProgressText(currPercent, MAX_PERCENT, float64(pgbWidth))
	
	fmt.Printf("%s%s%s", strHead, sProgress, strEnd)
		
}
func (pgb *MinPgb)Flush(){
	fmt.Print(CH_RESET_LINE)
}
func CreateProgressText(currPercent float64, totalPercent float64, txtWidth float64) string{

	s := "["
	for i:=1; i<= CallTextAppend(txtWidth, totalPercent); i++ {
		if i <= CallTextAppend(txtWidth, currPercent){
			s += "#"	
		}else{
			s += " "
		}		
	}
	s += "]"
	return s
}
/*
// unix.Winsize 
type Winsize struct {
    Row    uint16
    Col    uint16
    Xpixel uint16
    Ypixel uint16
}
*/
func GetWinsize() *unix.Winsize{
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil
	}
	return ws 
}
func CallTextAppend(txtLen float64, percent float64) int{
	v := (percent/100)*txtLen
	return int(v)

}

