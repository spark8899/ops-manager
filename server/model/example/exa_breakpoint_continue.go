package example

import (
	"github.com/spark8899/ops-manager/server/global"
)

// file struct, 文件结构体
type ExaFile struct {
	global.OPM_MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.OPM_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
