package main

import "strconv"

type ID int64

func (f ID) Int64() int64 {
	return int64(f)
}

func ParseInt64(id int64) ID {
	return ID(id)
}

func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}

func ParseString(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	return ID(i), err
}

func (f ID) Base2() string {
	return strconv.FormatInt(int64(f), 2)
}

func ParseBase2(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 2, 64)
	return ID(i), err
}

func (f ID) Base36() string {
	return strconv.FormatInt(int64(f), 36)
}

func ParseBase36(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 36, 64)
	return ID(i), err
}

// TODO:
func (f ID) Time() int64 {
	return Epoch
}

// TODO:
func (f ID) Node() int64 {
	return Epoch
}
