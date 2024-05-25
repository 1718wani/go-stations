package model

import "fmt"

// ErrNotFound は対象のTODOが存在しない場合のエラーを表します。
type ErrNotFound struct {
	ID int64
}

// Error メソッドはエラーメッセージを返します。
func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("TODO with ID %d not found", e.ID)
}