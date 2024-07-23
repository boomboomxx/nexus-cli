package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Almost completely ripped off https://www.socketloop.com/tutorials/golang-natural-string-sorting-example

type Compare func(str1, str2 string) bool

func (cmp Compare) Sort(strs []string) {
	strSort := &strSorter{
		strs: strs,
		cmp:  cmp,
	}
	sort.Sort(strSort)
}

type strSorter struct {
	strs []string
	cmp  func(str1, str2 string) bool
}

func extractNumberFromString(str string) (num int64) {
	strSlice := make([]string, 0)
	for _, v := range str {
		if unicode.IsDigit(v) {
			strSlice = append(strSlice, string(v))
		}
	}

	// If the tag was all non-digits, the strSlice would be empty (e.g., 'latest')
	// therefore just throw it to the end (1 << 32 is maxint)
	if len(strSlice) == 0 {
		return 1 << 32
	}

	num, err := strconv.ParseInt(strings.Join(strSlice, ""), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

// 提取需要删除的镜像版本
func extractToDelete(tags []string, keep int) []string {
	var toDelete []string

	if len(tags) < 1 {
		return toDelete
	}
	// 定义 map 用于存储分组后的字符串
	groupMap := make(map[string][]string)

	// 分组
	for _, tag := range tags {
		if !checkPattern(tag) {
			fmt.Printf("tag: [%s] 不匹配 SZIS 规则, 跳过 \n", tag)
			continue
		}
		prefix := getPrefix(tag)
		groupMap[prefix] = append(groupMap[prefix], tag)
	}

	// 定义切片用于存储需要删除的部分

	// 对每个分组进行处理
	for _, group := range groupMap {
		// 对每个分组的字符串按照后缀部分进行排序
		sort.Slice(group, func(i, j int) bool {
			return getSuffixNumber(group[i]) < getSuffixNumber(group[j])
		})
		// 剔除最大的那个，即删除第一个元素
		if len(group) > keep {
			toDelete = append(toDelete, group[:len(group)-keep]...)
		}
	}
	return toDelete
}

// 获取前缀部分
func getPrefix(tag string) string {
	parts := strings.Split(tag, ".")
	return strings.Join(parts[:3], ".")
}
func checkPattern(tag string) bool {
	pattern := "^\\d+\\.(\\d+)\\.(\\d+)\\.(\\d+)(?:-\\w+(?:-\\d+)?)?$"
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return regex.MatchString(tag)
}

func extractNumberFromStringForSzis(str string) (num int64) {
	strSlice := make([]string, 0)

	for _, v := range str {
		if unicode.IsDigit(v) {
			strSlice = append(strSlice, string(v))
		}
	}
	parts := strings.Split(str, "-")
	if len(parts) == 1 {
		// 使其最大化
		strSlice = append(strSlice, "99999999")
	}

	// If the tag was all non-digits, the strSlice would be empty (e.g., 'latest')
	// therefore just throw it to the end (1 << 32 is maxint)
	if len(strSlice) == 0 {
		return 1 << 32
	}

	num, err := strconv.ParseInt(strings.Join(strSlice, ""), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

// 获取后缀分数
func getSuffixNumber(tag string) int64 {
	parts := strings.Split(tag, ".")
	return extractNumberFromStringForSzis(parts[3])
}

func (s *strSorter) Len() int { return len(s.strs) }

func (s *strSorter) Swap(i, j int) { s.strs[i], s.strs[j] = s.strs[j], s.strs[i] }

func (s *strSorter) Less(i, j int) bool { return s.cmp(s.strs[i], s.strs[j]) }
