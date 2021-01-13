package binchunk

const (
	LUA_SIGNATURE    = "\x1bLua"
	LUAC_VERSION     = 0x53
	LUA_FORMAT       = 0
	LUAC_DATA        = "\x19\x93\r\n\x1a\n"
	CINT_SIZE        = 4
	CSIZET_SIZE      = 8
	INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE = 8
	LUA_NUMBER_SIZE  = 0
	LUAC_INT         = 0X5678
	LUAC_NUM         = 370.5
)

// binaryChunk 二进制Chunk
type binaryChunk struct {
	header                  // 头部
	sizeUpvalues byte       // 主函数upvalue变量
	mainFunc     *Prototype // 主函数原型
}

// header 头部
type header struct {
	signature       [4]byte // 签名 Lua二进制chunk魔数 4个字节。分别是ESC、L、u、a的ASCII码，用十六进制表示是0x1B4C7561,写出字面量就是"\x1bLua"
	version         byte    // 版本号
	format          byte    // 格式号 Lua官方使用的格式号是0
	luacData        [6]byte // LUAC_DATA "\x19\x93\r\n\x1a\n" 前两个字节是 0x1993(lua1.0发布的年份) 后四个字节依次是回车符(0x0D)、换行符(0x0A)、替换符(0x1A)和另一个换行符
	cintSize        byte    // cint 占用字节数
	sizetSize       byte    // size_t 占用字节数
	instructionSize byte    // Lua虚拟机指令 占用字节数
	luaIntegerSize  byte    // Lua整数 占用字节数
	luaNumberSize   byte    // Lua浮点数 占用字节数
	luacInt         int64   // LUAC_INT lua整数值 0x5678, 存储这个整数的目的是为了检测二进制chunk的大小端方式
	luacNum         float64 // LUAC_NUM Lua浮点数370.5, 存储这个浮点数的目的是为了检测二进制chunk所使用的浮点数格式
}

// Prototype 主函数原型
type Prototype struct {
	Source          string
	LineDefined     uint32
	LastLineDefined uint32
	NumParams       byte
	IsVararg        byte
	MaxStackSize    byte
	Code            []uint32
	Constants       []interface{}
	Upvalues        []Upvalue
	Protos          []*Prototype
	LineInfo        []uint32
	LocVars         []LocVar
	UpvalueNames    []string
}

type Upvalue struct {
}

type LocVar struct {
}
