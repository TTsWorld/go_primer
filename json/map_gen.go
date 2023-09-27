package json

//func mapGen(m map[string]interface{}) (string, error) {
//
//	buf := writeStart([]byte{})
//	i := 0
//	for k, v := range m {
//		if i != 0 {
//			buf = writeChar(buf, []byte{','})
//		}
//		buf = writeChar(buf, k)
//		buf = writeChar(buf, ":")
//		buf = writeChar(buf, cast.ToString(v))
//		i++
//		//vType := reflect.TypeOf(v)
//		//switch vType.(type) {
//		//cast reflect.Int:
//		//
//		//}
//	}
//	return buf, nil
//}
//
//func writeStart(s []byte) []byte {
//	return append(s, '{')
//}
//
//func writeChar(s []byte, b byte) []byte {
//	return append(s, toWrite)
//}
//
//func d() {
//	cast.ToString("")
//}
