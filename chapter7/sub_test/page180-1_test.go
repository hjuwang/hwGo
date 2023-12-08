package sub_test

/**
验证自测试
使用t.run(name,f func(t *T)) 在f 中嵌套 t。
*/
import (
	"os"
	"testing"
	"time"
)

func Add1(a int, b int) int {
	return a + b
}

func sub11(t *testing.T) {

	t.Parallel()
	time.Sleep(time.Second * 10)
	//do some testing

}
func sub22(t *testing.T) {

	t.Parallel()

	time.Sleep(time.Second * 2)

	//do some testing

}
func sub33(t *testing.T) {
	t.Parallel()

	time.Sleep(time.Second * 1)

	//do some testing

}

func TestSub1(t *testing.T) {

	//set up
	t.Logf("set up")

	t.Run("Test1", sub11)
	t.Run("Test2", sub22)
	t.Run("Test3", sub33)
	//tear down

	t.Logf("tear down")
}

func TestMain(m *testing.M) {

	println("TestMain set up")

	retCode := m.Run()

	println("TestMain tear down")

	os.Exit(retCode)

}
