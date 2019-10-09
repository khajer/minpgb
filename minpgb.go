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
	currTxt := ""
	seperator := pgTypeList[PgType].Seperator
	remainTxt := ""
		
	if currPercent < totalPercent {
		curCnt := CallTextAppend(txtWidth, currPercent)
		if len(seperator) > 0{
			curCnt -= len(seperator)
			if curCnt < 0 {curCnt = 0}
		}
		currTxt = strings.Repeat(pgTypeList[PgType].MarkCh, curCnt)					
		remainCnt := int(txtWidth) - (len(currTxt) + len(seperator))
		remainTxt = strings.Repeat(pgTypeList[PgType].RemainCh, remainCnt)		
	}else{
		currTxt = strings.Repeat(pgTypeList[PgType].MarkCh, int(txtWidth))
		seperator = ""		
	}	
	s = "["+currTxt+seperator+remainTxt+"]"	
	return s
}

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
