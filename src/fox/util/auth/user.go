package auth

import "regexp"

/**手机号检测
     * @param string $mobile 手机号
     * @return bool
     */
func CheckMobile(str string) bool{
	// $pattern = "/^1\d{10}$/";
	ok, _ := regexp.Match(`^1[34578][0-9]{9}$`, []byte(str));
	return ok
}

/**邮箱检测
 * @param string $email 邮箱
 * @return bool
 */
func CheckMail(str string) bool {
	ok, _ := regexp.Match(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, []byte(str));
	return ok
}