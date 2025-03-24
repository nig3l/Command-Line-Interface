package main

import (
	"fmt"
	"os"
	"strings"
)
// Task represents a single todo item
type Task struct {
	ID        int
	Desc      string
	Completed bool
}

