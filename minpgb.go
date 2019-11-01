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

	pgTypeList = make([]ProgressbarType, PGTYPE_BEER+1)		
	pgTypeList[PGTYPE_NORMAL] = ProgressbarType{"[", "#", " ", " ", "]", "", ""}
	pgTypeList[PGTYPE_ARROW] = ProgressbarType{"[", "=", ">", " ", "]", "", ""}
	pgTypeList[PGTYPE_DOT] = ProgressbarType{"[", ".", "", " ", "]", "", ""}
	pgTypeList[PGTYPE_BLOCK] = ProgressbarType{"|", "▓", "▒", " ", "|", "", ""}
	pgTypeList[PGTYPE_B1] = ProgressbarType{"[", "|", "|", "-", "]", "", ""}
	pgTypeList[PGTYPE_BLOCK1] = ProgressbarType{"", "█", "▒", "░", "", "", ""}
	pgTypeList[PGTYPE_BLOCK2] = ProgressbarType{"|", "▓", "▒", "░", "|", "", ""}
	pgTypeList[PGTYPE_BEER] = ProgressbarType{"|", "=", "🍺", "-", "|", "", ""}

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
	currPercent := pgb.Curr/pgb.Total*MAX_PERCENT	
		
	strHead := CreatePreLoadingText(pgPreText, pgb.Curr, pgb.Total)
	strEnd := fmt.Sprintf(" %3.2f%s", currPercent, "%")

	col := uint16(MAX_PERCENT)
	if winsize != nil{
		col = winsize.Col
	}

	pgbWidth := int(col) - len(strHead) - len(strEnd)
	sProgress := CreateProgressText(currPercent, MAX_PERCENT, pgbWidth)	

	// process 
	fmt.Print(CH_RESET_LINE)		

	fmt.Printf("%s%s%s", strHead, sProgress, strEnd)

}

func (pgb *MinPgb)SetPreText(pretext string){
	pgPreText = pretext
}
func (pgb *MinPgb)Flush(){
	fmt.Print(CH_RESET_LINE)
	pgb.Curr = 0
}
func (pgb *MinPgb)End(){
	fmt.Println("")
	pgb.Curr = 0
}

func (pgb *MinPgb)SetStyle(styleID int){
	pgType = styleID
}

func CreateProgressText(currPercent float64, totalPercent float64, txtWidth int) string{
	
	beginTxt := pgTypeList[pgType].BucketBegin
	chCurrent := pgTypeList[pgType].MarkCh
	seperator := pgTypeList[pgType].Seperator
	chRemain := pgTypeList[pgType].RemainCh	
	endTxt := pgTypeList[pgType].BucketEnd

	remainTxt := ""

	TxtLenAll := txtWidth - len(beginTxt) - len(endTxt)
	if currPercent >= totalPercent{
		return beginTxt+strings.Repeat(chCurrent, TxtLenAll)+endTxt
	}

	curCnt := CallTextAppend(TxtLenAll, currPercent)
	remainCnt := TxtLenAll - curCnt - len(seperator)
	currTxt := strings.Repeat(chCurrent, curCnt)
	if remainCnt > 0{
		remainTxt =  strings.Repeat(chRemain, remainCnt)	
	}

	pgtxt := currTxt+seperator+remainTxt

	return beginTxt+pgtxt+endTxt
}

func GetWinsize() *unix.Winsize{
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil
	}
	return ws 
}
func CallTextAppend(txtLen int, percent float64) int{	
	v := (percent/100)*float64(txtLen)
	return int(v)
}

func CreatePreLoadingText(pretext string, curr float64, total float64) string{
	if len(pretext) > 0 {
		pretext += " "
	}
	return fmt.Sprintf("%s[%3.0f/%3.0f]", pretext, curr, total)
}
