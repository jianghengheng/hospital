/*
 * @Author: jiangheng jh@pzds.com
 * @Date: 2025-02-06 15:43:38
 * @LastEditors: jiangheng jh@pzds.com
 * @LastEditTime: 2025-02-06 15:49:36
 * @FilePath: \hospital\utils\register.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// models/register.go
package utils

var registeredModels []interface{}

func RegisterModel(model interface{}) {
	registeredModels = append(registeredModels, model)
}

func GetRegisteredModels() []interface{} {
	return registeredModels
}
