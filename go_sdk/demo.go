package acms_sdk

import (
	"fmt"
)

// Demo 使用示例
func Demo() {
	client, err := CreateClient("123456", "123456")
	if err != nil {
		fmt.Println("登录失败:", err)
		return
	}
	defer client.Logout()

	fmt.Printf("登录成功: %v (%v)\n", client.UserInfo["name"], client.UserInfo["position"])

	// 访问量
	vc, _ := client.GetVisitCount()
	fmt.Printf("访问量: %v\n", vc["data"])

	// 教师
	teachers, _ := client.GetAllTeachers()
	if data, ok := teachers["data"].([]any); ok {
		names := make([]string, len(data))
		for i, t := range data {
			names[i] = t.(map[string]any)["name"].(string)
		}
		fmt.Printf("教师: %v\n", names)
	}

	// 班级
	names, _ := client.GetClassNames()
	fmt.Printf("班级: %v\n", names["data"])

	// 题库分类
	cats, _ := client.GetQuestionCategories()
	if data, ok := cats["data"].([]any); ok {
		catNames := make([]string, len(data))
		for i, c := range data {
			catNames[i] = c.(map[string]any)["questionCategory"].(string)
		}
		fmt.Printf("题库分类: %v\n", catNames)
	}

	// 题目
	questions, _ := client.GetAllQuestions("", nil)
	fmt.Printf("题目总数: %d\n", len(questions))

	// 查询教师
	teacher, _ := client.QueryTeacher(TeacherQuery{TeacherID: "20050027"})
	if data, ok := teacher["data"].(map[string]any); ok {
		fmt.Printf("查询教师20050027: %v\n", data["name"])
	}

	// 按条件查题目
	catID := 9
	q := QuestionListQuery{PageSize: 3, CategoryID: &catID, Type: "选择题"}
	list, _ := client.GetQuestionList(q)
	if data, ok := list["data"].(map[string]any); ok {
		if items, ok := data["items"].([]any); ok {
			for _, item := range items {
				it := item.(map[string]any)
				text := it["questionText"].(string)
				if len(text) > 50 {
					text = text[:50] + "..."
				}
				fmt.Printf("  [%v] %s\n", it["type"], text)
			}
		}
	}

	// 分页查询学生
	sp, _ := client.GetStudentPage(PageQuery{PageNum: 1, PageSize: 5})
	fmt.Printf("学生分页: total=%v\n", sp["data"].(map[string]any)["total"])

	// AI 提问
	aiResp, _ := client.AIStudentAsk([]AIAskItem{
		{TestID: 142, StudentID: "2407110107", QuestionID: 1040, Index: 1,
			QuestionText: "什么是存储过程？", StudentAnswer: "<p>...</p>", Score: 0, MaxScore: 28},
	})
	fmt.Printf("AI提问: success=%v\n", aiResp["success"])
}
