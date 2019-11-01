package minpgb


const (
	MAX_PERCENT float64 = 100

	CH_RESET_LINE 			string = "\r\033[K"	

	CH_COLOR_RED			string = "\033[0;31m"
	CH_COLOR_GREEN			string = "\033[0;32m"	
	CH_COLOR_YELLOW			string = "\033[0;33m"
	CH_COLOR_BLUE			string = "\033[0;34m"
	CH_COLOR_PURPLE			string = "\033[0;35m"
	CH_COLOR_CYAN			string = "\033[0;36m"
	CH_COLOR_WHITE			string = "\033[0;37m"

	CH_COLOR_LIGHT_RED		string = "\033[1;31m"
	
	CH_COLOR_NO_COLOUR		string = "\033[0m"
									   
)	

const (
	PGTYPE_NORMAL int = iota
	PGTYPE_ARROW
	PGTYPE_DOT
	PGTYPE_BLOCK
	PGTYPE_B1
	PGTYPE_BLOCK1
	PGTYPE_BLOCK2
	PGTYPE_BEER
)

type ProgressbarType struct {
	BucketBegin, MarkCh, Seperator, RemainCh, BucketEnd string	
	CurrColor, RemainColor string
}

