package _go

func SafeGo(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				println(err.(string))
			}
		}()
	}()
}
