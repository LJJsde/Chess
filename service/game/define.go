package game

const (
	//ImgChessBoard 棋盘
	ImgChessBoard = 1
	//ImgSelect 选中
	ImgSelect = 2
	//ImgRedShuai 红帅
	ImgRedShuai = 8
	//ImgRedShi 红士
	ImgRedShi = 9
	//ImgRedXiang 红相
	ImgRedXiang = 10
	//ImgRedMa 红马
	ImgRedMa = 11
	//ImgRedJu 红车
	ImgRedJu = 12
	//ImgRedPao 红炮
	ImgRedPao = 13
	//ImgRedBing 红兵
	ImgRedBing = 14
	//ImgBlackJiang 黑将
	ImgBlackJiang = 16
	//ImgBlackShi 黑士
	ImgBlackShi = 17
	//ImgBlackXiang 黑相
	ImgBlackXiang = 18
	//ImgBlackMa 黑马
	ImgBlackMa = 19
	//ImgBlackJu 黑车
	ImgBlackJu = 20
	//ImgBlackPao 黑炮
	ImgBlackPao = 21
	//ImgBlackBing 黑兵
	ImgBlackBing = 22
)

const (
	//MusicSelect 选子
	MusicSelect = 100
	//MusicPut 落子
	MusicPut = 101
	//MusicEat 吃子
	MusicEat = 102
	//MusicJiang 将军
	MusicJiang = 103
	//MusicGameWin 胜利
	MusicGameWin = 104
	//MusicGameLose 失败
	MusicGameLose = 105
)

const (
	SquareSize  = 56
	BoardEdge   = 8
	BoardWidth  = BoardEdge + SquareSize*9 + BoardEdge
	BoardHeight = BoardEdge + SquareSize*10 + BoardEdge
)

const (
	Top    = 3
	Bottom = 12
	Left   = 3
	Right  = 11
)

var cucpcStartup = [256]int{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 20, 19, 18, 17, 16, 17, 18, 19, 20, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 21, 0, 0, 0, 0, 0, 21, 0, 0, 0, 0, 0,
	0, 0, 0, 22, 0, 22, 0, 22, 0, 22, 0, 22, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 14, 0, 14, 0, 14, 0, 14, 0, 14, 0, 0, 0, 0,
	0, 0, 0, 0, 13, 0, 0, 0, 0, 0, 13, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 12, 11, 10, 9, 8, 9, 10, 11, 12, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

var ccInBoard = [256]int{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//获得格子的Y
func getY(sq int) int {
	return sq >> 4
}

//获得格子的X
func getX(sq int) int {
	return sq & 15
}

//根据纵坐标和横坐标获得格子
func squareXY(x, y int) int {
	return x + (y << 4)
}

//翻转格子
func squareFlip(sq int) int {
	return 254 - sq
}

//X水平镜像
func xFlip(x int) int {
	return 14 - x
}

//Y垂直镜像
func yFlip(y int) int {
	return 15 - y
}

//获得红黑标记(红子是8，黑子是16)
func sideTag(sd int) int {
	return 8 + (sd << 3)
}

//获得对方红黑标记
func oppSideTag(sd int) int {
	return 16 - (sd << 3)
}

//获得走法的起点
func src(mv int) int {
	return mv & 255
}

//获得走法的终点
func dst(mv int) int {
	return mv >> 8
}

//根据起点和终点获得走法
func move(sqSrc, sqDst int) int {
	return sqSrc + sqDst*256
}
