package tests

import (
	"github.com/godrealms/apple-store-sdk/pkg/utils"
	"net/url"
	"testing"
)

// 测试 BuildQueryParams 方法，支持多层嵌套结构体
func TestBuildQueryParamsNested(t *testing.T) {
	// 测试用例 1：多层嵌套结构体
	type Address struct {
		City    string
		ZipCode int
	}
	type User struct {
		Name    string
		Age     int
		Active  bool
		Address Address
		Tags    []string
	}
	params1 := User{
		Name:   "Alice",
		Age:    25,
		Active: true,
		Address: Address{
			City:    "Wonderland",
			ZipCode: 12345,
		},
		Tags: []string{"golang", "developer"},
	}
	// 直接写预期结果，稍后会进行编码对比
	expected1 := "Active=true&Address[City]=Wonderland&Address[ZipCode]=12345&Age=25&Name=Alice&Tags[]=golang&Tags[]=developer"
	result1 := utils.BuildQueryParams(params1)
	t.Logf("Test Case 1: Input: %+v, Output: %s", params1, result1)
	if result1 != url.QueryEscape(expected1) {
		t.Errorf("Test case 1 failed. Expected: %s, Got: %s", url.QueryEscape(expected1), result1)

	}

	// 测试用例 2：嵌套 map
	params2 := map[string]any{
		"user": map[string]any{
			"name": "Bob",
			"age":  30,
			"address": map[string]any{
				"city":    "Metropolis",
				"zipCode": 54321,
			},
		},
		"active": false,
	}
	expected2 := "active=false&user[address][city]=Metropolis&user[address][zipCode]=54321&user[age]=30&user[name]=Bob"
	result2 := utils.BuildQueryParams(params2)
	t.Logf("Test Case 2: Input: %+v, Output: %s", params2, result2)
	if result2 != url.QueryEscape(expected2) {
		t.Errorf("Test case 2 failed. Expected: %s, Got: %s", url.QueryEscape(expected2), result2)
	}

	// 测试用例 3：数组嵌套
	params3 := map[string]any{
		"names": []string{"Alice", "Bob", "Charlie"},
	}
	expected3 := "names[]=Alice&names[]=Bob&names[]=Charlie"
	result3 := utils.BuildQueryParams(params3)
	t.Logf("Test Case 3: Input: %+v, Output: %s", params3, result3)
	if result3 != url.QueryEscape(expected3) {
		t.Errorf("Test case 3 failed. Expected: %s, Got: %s", url.QueryEscape(expected3), result3)
	}
}
