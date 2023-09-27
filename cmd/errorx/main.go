package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type temporary interface {
	Temporary() bool
}

func main() {
	var err error

	err = demo1()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("demo1 == ErrNoRows")
	}

	err = demo2()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("demo2 == ErrNoRows")
	}
	//for err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		fmt.Println("demo2 == ErrNoRows")
	//		break
	//	}
	//	err = errors.Unwrap(err)
	//}

	var te temporary
	if errors.As(err, &te) {
		if te.Temporary() {
			fmt.Println("this is temporary")
		}
	}
}

func demo1() error {
	// Error doing a query leads to the error
	return sql.ErrNoRows
}

func demo2() error {
	err := demo1()
	if err != nil {
		return fmt.Errorf("demo2: %w", err)
	}
	return nil
}
