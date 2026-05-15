package go_sdk

import "fmt"

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
	fmt.Printf("访问量: %v\n", vc.Data)

	// 教师
	teachers, _ := client.GetAllTeachers()
	names := make([]string, len(teachers.Data))
	for i, t := range teachers.Data {
		names[i] = t.Name
	}
	fmt.Printf("教师: %v\n", names)

	// 班级
	cnames, _ := client.GetClassNames()
	fmt.Printf("班级: %v\n", cnames.Data)

	// 题库分类
	cats, _ := client.GetQuestionCategories()
	catNames := make([]string, len(cats.Data))
	for i, c := range cats.Data {
		catNames[i] = c.QuestionCategory
	}
	fmt.Printf("题库分类: %v\n", catNames)

	// 题目
	questions, _ := client.GetAllQuestions("", nil)
	fmt.Printf("题目总数: %d\n", len(questions))

	// 查询教师
	teacher, _ := client.QueryTeacher("20050027")
	fmt.Printf("查询教师20050027: %v\n", teacher.Data.Name)

	// 按条件查题目
	catID := 9
	q := QuestionListQuery{PageSize: 3, CategoryID: &catID, Type: "选择题"}
	list, _ := client.GetQuestionList(q)
	for _, it := range list.Data.Items {
		text := it.QuestionText
		if len(text) > 50 {
			text = text[:50] + "..."
		}
		fmt.Printf("  [%v] %s\n", it.Type, text)
	}

	// 分页查询学生
	sp, _ := client.GetStudentPage(PageQuery{PageNum: 1, PageSize: 5})
	fmt.Printf("学生分页: total=%d\n", len(sp.Data))

	// 学生成绩（带结构体）
	result, _ := client.GetStudentResult("123456")
	for _, cls := range result.Data {
		for _, test := range cls.Tests {
			fmt.Printf("考试: %s (得分 %d/%d)\n", test.TestName, scoreSum(test.Results), totalMax(test.Results))
		}
	}

	// AI 提问
	aiResp, _ := client.AIStudentAsk([]AIAskItem{
		{TestID: 142, StudentID: "123456", QuestionID: 1040, Index: 1,
			QuestionText: "什么是存储过程？", StudentAnswer: "<p>...</p>", Score: 0, MaxScore: 28},
	})
	fmt.Printf("AI提问: success=%v\n", aiResp.Data)
}

func scoreSum(items []ResultItem) int {
	s := 0
	for _, it := range items {
		s += it.Score
	}
	return s
}

func totalMax(items []ResultItem) int {
	s := 0
	for _, it := range items {
		s += it.MaxScore
	}
	return s
}
