package user

import (
	"encoding/json"
	"test/dao"
)

func GetUserList() ([]string, error) {
	users_online := dao.GetUserList("上线")
	users_outline := dao.GetUserList("下线")

	// 这里传一个切片
	// sname 切片用于存储上线用户
	// 因为不会proto 里面存储user 所以使用[]string
	result := make([]string, 0)
	for _, val := range users_online {
		message := map[string]interface{}{
			"name":   val.Name,
			"status": val.Status,
		}
		res_val, _ := json.Marshal(message)
		//fmt.Println("res_val", res_val)
		result = append(result, string(res_val))
	}

	for _, val := range users_outline {
		message := map[string]interface{}{
			"name":   val.Name,
			"status": val.Status,
		}
		res_val, _ := json.Marshal(message)

		result = append(result, string(res_val))
	}

	return result, nil
}
