 Elem() 函数的用法   
    var x interface{} = new(interface{})

	//返回动态值，调用elem()方法的接收器只能是 接口和指针
	valueOf := reflect.ValueOf(x).Elem()

	//fmt.Printf("%v\n", valueOf)

	valueOf.Set(reflect.ValueOf("hello"))

	fmt.Println(x.string())


#### 获取结构体标签的用法

   field value.type().Field(i)