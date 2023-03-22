package ast

import (
	"testing"

	"github.com/spark8899/ops-manager/server/global"
	"github.com/spark8899/ops-manager/server/model/example"
)

const A = 123

func TestAddRegisterTablesAst(t *testing.T) {
	AddRegisterTablesAst("D:\\ops-manager\\server\\utils\\ast_test.go", "Register", "test", "testDB", "testModel")
}

func Register() {
	test := global.GetGlobalDBByDBName("test")
	test.AutoMigrate(example.ExaFile{})
}
