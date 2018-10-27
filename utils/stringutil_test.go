package utils

import (
    "testing"
)

func TestCapCamel2Const(t *testing.T) {
	str:="TestOne"
	constStr:=CapCamel2Const(str)
	if constStr != "TEST_ONE" {
		t.Error("expect1:",constStr)
	}
	
	str = "TEST1"
	constStr=CapCamel2Const(str)
	if constStr != "T_E_S_T1" {
		t.Error("expect1:",constStr)
	}
	str = "Test1One2"
	constStr=CapCamel2Const(str)
	if constStr != "TEST1_ONE2" {
		t.Error("expect1:",constStr)
	}
}

