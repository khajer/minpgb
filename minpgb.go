package minpgb

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"strings"

)

var pgb *MinPgb
var winsize *unix.Winsize
var PgType int
var pgTypeList []ProgressbarType

type ProgressbarType struct {
	MarkCh,Seperator,RemainCh 		string
}


const (
	MAX_PERCENT float64 = 100
	CH_RESET_LINE string = "\r\033[K"	

	PGTYPE_NORMAL = 0
	PGTYPE_DASH = 1
)


type MinPgb struct{
	Curr 	float64
	Total 	float64
}


func init(){
	pgb = New()
	winsize = GetWinsize()

	CreateProgressTypeList()	
	PgType = PGTYPE_NORMAL

	
}
func CreateProgressTypeList(){
	pgTypeList = []ProgressbarType{
		ProgressbarType{
			MarkCh:"=",
			Seperator:">",
			RemainCh:" ",
		},
		ProgressbarType{
			MarkCh:"#",
			Seperator:"#",
			RemainCh:" ",
		},
	}
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
	col := uint16(MAX_PERCENT)
	if winsize != nil{
		col = winsize.Col
	}
	pgbWidth := int(col) - (len(strHead)+len(strHead))

	sProgress := CreateProgressText(currPercent, MAX_PERCENT, float64(pgbWidth))
	
	fmt.Printf("%s%s%s", strHead, sProgress, strEnd)
		
}
func (pgb *MinPgb)Flush(){
	fmt.Print(CH_RESET_LINE)
	pgb.Curr = 0
}

func (pgb *MinPgb)SetStyle(styleID int){
	PgType = styleID
}
func CreateProgressText(currPercent float64, totalPercent float64, txtWidth float64) string{
	s := ""
	remainTxt := ""

	// markCh := "="
	// seperator := ">"
	// remainCh := " "	
	
	markCh := pgTypeList[PgType].MarkCh
	seperator := pgTypeList[PgType].Seperator
	remainCh := pgTypeList[PgType].RemainCh
	
	if currPercent < totalPercent {

		curCnt := CallTextAppend(txtWidth, currPercent)-len(seperator)
		curr := ""
		if curCnt > 0{
			curr = strings.Repeat(markCh, curCnt)	
		}		

		totalRemainCnt := int(txtWidth) - (len(curr)+len(seperator))		
								
		if totalRemainCnt >= 0 {
			remainTxt = strings.Repeat(remainCh, totalRemainCnt)	
		}else{
			remainTxt = strings.Repeat(markCh, CallTextAppend(txtWidth, totalPercent))
		}		
		s = "["+curr+seperator+remainTxt+"]"	
	}else{
		total := strings.Repeat(markCh, CallTextAppend(txtWidth, totalPercent))
		s = "["+total+"]"	
	}	
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
