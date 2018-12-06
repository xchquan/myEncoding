/*
 * @Author: huqing
 * @Date: 2018-12-06 09:48:31
 * @Last Modified by: huqing
 * @Last Modified time: 2018-12-06 09:54:38
 */

package myEncoding_test

import (
	"fmt"
	mycod "myself/myEncoding"
	"testing"
)

func TestA(t *testing.T) {
	str := "中国"
	res, err := mycod.GbkToUtf8([]byte(str))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(res))
	}
}
