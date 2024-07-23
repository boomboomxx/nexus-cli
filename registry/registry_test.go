package registry

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

func Test_Image_List(t *testing.T) {
	r, err := NewRegistry()
	if err != nil {
		log.Print(err)
	}
	images, err := r.ListImages()
	if err != nil {
		log.Print(err)
	}
	log.Print(images)
}

func Test_image_prod_del(t *testing.T) {
	r, err := NewRegistry()
	if err != nil {
		log.Print(err)
	}
	images, err := r.ListImages()
	if err != nil {
		log.Print(err)
		return
	}
	var prod_images []string

	for _, image := range images {
		if strings.HasPrefix(image, "prod") {
			prod_images = append(prod_images, image)
		}
	}

	var cnt = 0
	for _, imgName := range prod_images {
		tags, err := r.ListTagsByImage(imgName)
		if err != nil {
			log.Fatal("无tag")
			return
		}
		toBeDeleteTags := extractToDelete(tags, 1)
		if len(toBeDeleteTags) > 0 {
			for _, tag := range toBeDeleteTags {
				cnt++
				log.Printf("%s:%s image will be deleted ...\n", imgName, tag)
				r.DeleteImageByTag(imgName, tag)
			}
		}
	}
	log.Println("total deleted images: ", cnt)

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
			fmt.Printf("tag: [%s] dose not matching szis rules, skip \n", tag)
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

func extractNumberFromStringForSzis(str string) (num int) {
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

	num, err := strconv.Atoi(strings.Join(strSlice, ""))
	if err != nil {
		log.Fatal(err)
	}
	return num
}

// 获取后缀分数
func getSuffixNumber(tag string) int {
	parts := strings.Split(tag, ".")
	return extractNumberFromStringForSzis(parts[3])
}
