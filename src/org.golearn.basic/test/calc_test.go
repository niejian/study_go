package main

import  "testing"

func TestAdd(t *testing.T) {
	t.Helper()
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}
}

func TestMul(t *testing.T) {
	//if ans := Mul(10, 20); ans != 100 {
	//	t.Errorf("10 * 20 expected be 600, but %d got", ans)
	//}

	// 子测试
	// 正数
	t.Run("正数相乘结果", func(t *testing.T) {
		if ans := Mul(2, 3); ans != 6{
			t.Errorf("10 * 20 expected be 6, but %d got", ans)
		}
	})

	// 子测试-负数
	t.Run("负数相乘结果", func(t *testing.T) {
		if ans := Mul(-20, 3); ans != -60 {
			t.Errorf("10 * 20 expected be -60, but %d got", ans)
		}
	})

}

func TestSub(t *testing.T) {
	if ans := Sub(10, 1); ans != 9 {
		t.Errorf("10 - 1 expected be 9, but %d got ", ans)

	}
}

// 多个子测试场景
func TestMul2(t *testing.T) {

	mulCases := []struct{
		Name string
		Num1, Num2, Expect int
	}{
		{"pos", 1, 2, 2},
		{"neg", 1, -2, -2},
		{"zero", 0, -2, 10},

	}

	for _, r := range mulCases {
		t.Run(r.Name, func(t *testing.T) {
			t.Helper()
			if ans := Mul(r.Num1, r.Num2); ans != r.Expect {
				t.Fatalf("%d * %d expected %d, but %d got",
					r.Num1, r.Num2, r.Expect, ans)
			}
		})
	}

}

//
/*
测试用例
1. 在相同的包路径中创建 xxx_test.go文件
2. 编写测试用例
3. 运行测试用例，查看测试用例中所有函数的运行结果 go test -v
*/

