package go_utils

// 将十进制转换为 任意进制,需要注意的是，域名总不能有 下划线(_)，但是可以有减号(-)
// 0 -- > 0
// 1 -- > 1
// 10-- > a
// 61-- > Z
//  id 需要转换的数字
//  szTemplate 模版
//  szTemplate 的长度决定进制 数据, 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ 表示 62 进制度
func TransInt64ToN(id int64, szTemplate string) string {
	n := int64(len(szTemplate))
	var shortUrl []byte
	for {
		var result byte
		number := id % n
		result = szTemplate[number]
		var tmp []byte
		tmp = append(tmp, result)
		shortUrl = append(tmp, shortUrl...)
		id = id / n
		if id == 0 {
			break
		}
	}
	return string(shortUrl)
}

// 数字转 62 进制
func TransInt64To62(id int64) string {
	return TransInt64ToN(id, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
