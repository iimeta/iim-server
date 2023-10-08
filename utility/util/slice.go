package util

import (
	"strconv"
	"strings"
)

type IntInterface interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64
}

// 切片去重
func Unique[T IntInterface | string](data []T) []T {

	list, hash := make([]T, 0), make(map[T]struct{})

	for _, value := range data {
		if _, ok := hash[value]; !ok {
			list = append(list, value)
			hash[value] = struct{}{}
		}
	}

	return list
}

// 切片转 map
func ToMap[T any, K int | string](arr []T, fn func(T) K) map[K]T {
	var m = make(map[K]T)

	for _, t := range arr {
		m[fn(t)] = t
	}

	return m
}

func ParseIds(str string) []int {
	str = strings.TrimSpace(str)
	ids := make([]int, 0)

	if str == "" {
		return ids
	}

	for _, value := range strings.Split(str, ",") {
		if id, err := strconv.Atoi(value); err == nil {
			ids = append(ids, id)
		}
	}

	return ids
}
