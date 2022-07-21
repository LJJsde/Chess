package game

//PositionStruct 局面结构
type PositionStruct struct {
	sdPlayer    int      // 轮到谁走，0=红方，1=黑方
	ucpcSquares [256]int // 棋盘上的棋子
}

//NewPositionStruct 初始化棋局
func NewPositionStruct() *PositionStruct {
	p := &PositionStruct{}
	if p == nil {
		return nil
	}
	return p
}

//startup 初始化棋盘
func (p *PositionStruct) startup() {
	p.sdPlayer = 0
	for sq := 0; sq < 256; sq++ {
		p.ucpcSquares[sq] = cucpcStartup[sq]
	}
}

//changeSide 交换走子方
func (p *PositionStruct) changeSide() {
	p.sdPlayer = 1 - p.sdPlayer
}

//addPiece 在棋盘上放一枚棋子
func (p *PositionStruct) addPiece(sq, pc int) {
	p.ucpcSquares[sq] = pc
}

//delPiece 从棋盘上拿走一枚棋子
func (p *PositionStruct) delPiece(sq int) {
	p.ucpcSquares[sq] = 0
}

//movePiece 搬一步棋的棋子
func (p *PositionStruct) movePiece(mv int) {
	sqSrc := src(mv)
	sqDst := dst(mv)
	p.delPiece(sqDst)
	pc := p.ucpcSquares[sqSrc]
	p.delPiece(sqSrc)
	p.addPiece(sqDst, pc)
}

//makeMove 走一步棋
func (p *PositionStruct) makeMove(mv int) {
	p.movePiece(mv)
	p.changeSide()
}
