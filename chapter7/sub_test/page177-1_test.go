package sub_test

/**
验证自测试
使用t.run() 来进行子测试
*/
import "testing"

func Add(a int, b int) int {
	return a + b
}

func sub1(t *testing.T) {
	var a, b = 1, 2
	var expected = 3

	if actual := Add(a, b); actual != expected {

		t.Errorf("Add(%d, %d) = %d; expected %d", a, b, actual, expected) // 期望值不等于实际值= Add(a, b)

	}

}
func sub2(t *testing.T) {
	var a, b = 1, 2
	var expected = 3

	if actual := Add(a, b); actual != expected {

		t.Errorf("Add(%d, %d) = %d; expected %d", a, b, actual, expected) // 期望值不等于实际值= Add(a, b)

	}

}
func sub3(t *testing.T) {
	var a, b = 1, 2
	var expected = 3

	if actual := Add(a, b); actual != expected {

		t.Errorf("Add(%d, %d) = %d; expected %d", a, b, actual, expected) // 期望值不等于实际值= Add(a, b)

	}

}

func TestSub(t *testing.T) {

	//set up

	t.Run("A=1", sub1)
	t.Run("A=2", sub2)
	t.Run("A=3", sub3)
	//tear down
}
