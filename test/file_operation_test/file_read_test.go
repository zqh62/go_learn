package file_operation_test

import (
	"go_learn/file_operation/file_operation"
	"testing"
)

func TestMyOpenFile(t *testing.T) {
	file_operation.MyOpenFile("test.txt")
}
