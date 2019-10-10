package minpgb

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"strings"

)

var pgb *MinPgb
var winsize *unix.Winsize
var pgType int
var pgTypeList []ProgressbarType
var pgPreText string

type ProgressbarType struct {
	MarkCh, Seperator, RemainCh 		string
}

const (
	MAX_PERCENT float64 = 100
	CH_RESET_LINE string = "\r\033[K"	

	PGTYPE_NORMAL = 0
	PGTYPE_DASH = 1
)

type MinPgb struct{
	Curr, Total 	float64
}

func init(){
	pgb = New()
	winsize = GetWinsize()

	CreateProgressTypeList()	
	pgType = PGTYPE_NORMAL
	pgPreText = ""
	
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
		
	strHead := CreatePreLoadingText(pgPreText, pgb.Curr, pgb.Total)
	strEnd := fmt.Sprintf(" %2.2f%s", currPercent, "%")

	col := uint16(MAX_PERCENT)
	if winsize != nil{
		col = winsize.Col
	}
	spacer := 4
	pgbWidth := int(col) - (len(strHead)+len(strEnd)+spacer) 

	sProgress := CreateProgressText(currPercent, MAX_PERCENT, float64(pgbWidth))	
	fmt.Printf("%s %s %s", strHead, sProgress, strEnd)		
}
func (pgb *MinPgb)SetPreText(pretext string){
	pgPreText = pretext
}
func (pgb *MinPgb)Flush(){
	fmt.Print(CH_RESET_LINE)
	pgb.Curr = 0
}

func (pgb *MinPgb)SetStyle(styleID int){
	pgType = styleID
}

func CreateProgressText(currPercent float64, totalPercent float64, txtWidth float64) string{
	s := ""
	currTxt := ""
	seperator := pgTypeList[pgType].Seperator
	remainTxt := ""
		
	if currPercent < totalPercent {
		curCnt := CallTextAppend(txtWidth, currPercent)
		if len(seperator) > 0{
			curCnt -= len(seperator)
			if curCnt < 0 {curCnt = 0}
		}
		currTxt = strings.Repeat(pgTypeList[pgType].MarkCh, curCnt)					
		remainCnt := int(txtWidth) - (len(currTxt) + len(seperator))
		remainTxt = strings.Repeat(pgTypeList[pgType].RemainCh, remainCnt)		
	}else{
		currTxt = strings.Repeat(pgTypeList[pgType].MarkCh, int(txtWidth))
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

func CreatePreLoadingText(pretext string, curr float64, total float64) string{
	if len(pretext) > 0 {
		pretext += " "
	}
	return fmt.Sprintf("%s[%.0f/%.0f]", pretext, curr, total)
}
