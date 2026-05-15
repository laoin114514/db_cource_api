package go_sdk

// ═══════════════════════════════════════════════════════════════
// 通用响应
// ═══════════════════════════════════════════════════════════════

type ApiResponse[T any] struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Total int    `json:"total"`
	Data  T      `json:"data"`
}

func (r ApiResponse[T]) OK() bool { return r.Code == 200 }

type AIChatResponse struct {
	Success   bool   `json:"success"`
	SessionID string `json:"sessionId"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

// ═══════════════════════════════════════════════════════════════
// 认证
// ═══════════════════════════════════════════════════════════════

type LoginData struct {
	Token string         `json:"token"`
	User  map[string]any `json:"user"`
}

// ═══════════════════════════════════════════════════════════════
// 用户
// ═══════════════════════════════════════════════════════════════

type UserData struct {
	Role     int    `json:"role"`
	Gender   string `json:"gender"`
	Name     string `json:"name"`
	ID       string `json:"id"`
	Position string `json:"position"`
}

// ═══════════════════════════════════════════════════════════════
// 教师
// ═══════════════════════════════════════════════════════════════

type TeacherData struct {
	TeacherID     string `json:"teacherId"`
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	Phone         any    `json:"phone"`
	Position      string `json:"position"`
	ListOfClasses string `json:"listOfClasses"`
}

type TeacherBrief struct {
	TeacherID string `json:"teacher_id"`
	Name      string `json:"name"`
}

type TeacherPageData struct {
	Total int           `json:"total"`
	Items []TeacherData `json:"items"`
	// or it might be a flat array; handled by the caller
}

// ═══════════════════════════════════════════════════════════════
// 班级
// ═══════════════════════════════════════════════════════════════

type ClassData struct {
	ClassID        int    `json:"classId"`
	ClassName      string `json:"className"`
	ListOfStudents string `json:"listOfStudents"`
	ListOfTeachers string `json:"listOfTeachers"`
	HeadteacherID  string `json:"headteacherId"`
}

type ClassPageData struct {
	Total int         `json:"total"`
	Items []ClassData `json:"items"`
}

// ═══════════════════════════════════════════════════════════════
// 学生
// ═══════════════════════════════════════════════════════════════

type StudentData struct {
	StudentID string `json:"studentId"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Phone     any    `json:"phone"`
	Notes     any    `json:"notes"`
	ClassName string `json:"className"`
	ClassID   int    `json:"classId"`
}

type StudentPageData struct {
	Total int           `json:"total"`
	Items []StudentData `json:"items"`
}

// ═══════════════════════════════════════════════════════════════
// 题库分类
// ═══════════════════════════════════════════════════════════════

type QuestionCategoryData struct {
	ID               int    `json:"id"`
	QuestionCategory string `json:"questionCategory"`
}

// ═══════════════════════════════════════════════════════════════
// 题目
// ═══════════════════════════════════════════════════════════════

type QuestionData struct {
	ID           int    `json:"id"`
	CategoryID   int    `json:"categoryId"`
	Type         string `json:"type"`
	AnswerTime   int    `json:"answerTime"`
	QuestionText string `json:"questionText"`
	Answer       string `json:"answer"`
	Keywords     string `json:"keywords"`
	Score        int    `json:"score"`
	TeacherID    string `json:"teacherId"`
	IsAIGenerate bool   `json:"isAiGenerate"`
}

type QuestionPageData struct {
	Total    int            `json:"total"`
	PageSize int            `json:"pageSize"`
	Items    []QuestionData `json:"items"`
	PageNum  int            `json:"pageNum"`
}

// ═══════════════════════════════════════════════════════════════
// 考试/测试
// ═══════════════════════════════════════════════════════════════

type TestData struct {
	ID          int    `json:"id"`
	TestName    string `json:"testName"`
	ClassName   string `json:"className"`
	TotalScore  int    `json:"totalScore"`
	TeacherID   string `json:"teacherId"`
	PublishTime string `json:"publishTime"`
	Deadline    string `json:"deadline"`
	AnswerTime  int    `json:"answerTime"`
	Purpose     string `json:"purpose"`
}

type TeacherPublishData struct {
	TestCount int           `json:"testCount"`
	ClassID   int           `json:"classId"`
	Tests     []TestSummary `json:"tests"`
}

type TestSummary struct {
	PublishTime    string           `json:"publishTime"`
	AnswerTime     int              `json:"answerTime"`
	Purpose        string           `json:"purpose"`
	StudentDetails []map[string]any `json:"studentDetails"`
}

// ═══════════════════════════════════════════════════════════════
// 考试结果
// ═══════════════════════════════════════════════════════════════

type StudentResultGroup struct {
	ClassID int                 `json:"classId"`
	Tests   []StudentTestResult `json:"tests"`
}

type StudentTestResult struct {
	TestID      int          `json:"testId"`
	TestName    string       `json:"testName"`
	PublishTime string       `json:"publishTime"`
	Deadline    string       `json:"deadline"`
	AnswerTime  int          `json:"answerTime"`
	Purpose     string       `json:"purpose"`
	TotalScore  int          `json:"totalScore"`
	Results     []ResultItem `json:"results"`
}

type ResultItem struct {
	ClassID    int    `json:"classId"`
	StudentID  string `json:"studentId"`
	QuestionID int    `json:"questionId"`
	TestID     int    `json:"testId"`
	Score      int    `json:"score"`
	MaxScore   int    `json:"maxScore"`
	Answer     string `json:"answer"`
	Comment    string `json:"comment"`
}

type ResultDetailData struct {
	QuestionID    int    `json:"questionId"`
	QuestionText  string `json:"questionText"`
	StudentAnswer any    `json:"studentAnswer"`
	CorrectAnswer string `json:"correctAnswer"`
	Keywords      string `json:"keywords"`
	Score         int    `json:"score"`
	MaxScore      int    `json:"maxScore"`
	Comment       string `json:"comment"`
}

// ═══════════════════════════════════════════════════════════════
// 视频
// ═══════════════════════════════════════════════════════════════

type VideoTreeNode struct {
	ID       int             `json:"id"`
	Label    string          `json:"label"`
	Children []VideoTreeNode `json:"children"`
}

// ═══════════════════════════════════════════════════════════════
// 考勤
// ═══════════════════════════════════════════════════════════════

type AttendanceRecord struct {
	ID           int    `json:"id"`
	StudentID    string `json:"studentId"`
	StudentName  string `json:"studentName"`
	Status       string `json:"status"`
	Date         string `json:"date"`
	IsAttendance bool   `json:"isAttendance"`
}

// ═══════════════════════════════════════════════════════════════
// 校历
// ═══════════════════════════════════════════════════════════════

type SemesterData struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
