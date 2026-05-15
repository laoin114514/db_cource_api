package go_sdk

// ── 通用分页 ──

type PageQuery struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Query    string `json:"query"`
}

func NewPageQuery() PageQuery {
	return PageQuery{PageNum: 1, PageSize: 20}
}

// ── 认证 ──

type LoginForm struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

// ── 用户 ──

type UserQuery struct {
	ID string `json:"id"`
}

type UserInsert struct {
	ID       string `json:"id"`
	Role     int    `json:"role"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserUpdate struct {
	ID   string `json:"id"`
	Role int    `json:"role"`
}

// ── 角色 ──

type RoleForm struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type RoleUpdate struct {
	ID   int    `json:"id"`
	Name string `json:"name,omitempty"`
}

// ── 教师 ──

type TeacherQuery struct {
	TeacherID string `json:"teacherId"`
}

type TeacherForm struct {
	TeacherID string `json:"teacherId"`
	Name      string `json:"name"`
	Gender    string `json:"gender,omitempty"`
	Position  string `json:"position,omitempty"`
	Password  string `json:"password,omitempty"`
}

type TeacherClassRelation struct {
	ClassID   int    `json:"classId"`
	TeacherID string `json:"teacherId"`
}

// ── 班级 ──

type ClassQuery struct {
	ClassName string `json:"className"`
}

type ClassForm struct {
	ClassName     string `json:"className"`
	HeadteacherID string `json:"headteacherId,omitempty"`
}

type ClassUpdate struct {
	ClassID       int    `json:"classId"`
	ClassName     string `json:"className,omitempty"`
	HeadteacherID string `json:"headteacherId,omitempty"`
}

// ── 学生 ──

type StudentQuery struct {
	StudentID string `json:"studentId"`
}

type StudentForm struct {
	StudentID string  `json:"studentId"`
	Name      string  `json:"name"`
	Gender    *string `json:"gender,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Notes     *string `json:"notes,omitempty"`
	ClassName *string `json:"className,omitempty"`
}

func StrPtr(s string) *string { return &s }

type StudentUpdate struct {
	StudentID string  `json:"studentId"`
	Name      *string `json:"name,omitempty"`
	Gender    *string `json:"gender,omitempty"`
	Phone     *string `json:"phone,omitempty"`
	Notes     *string `json:"notes,omitempty"`
	ClassName *string `json:"className,omitempty"`
}

type StudentUpdateClass struct {
	StudentID string `json:"studentId"`
	ClassID   *int   `json:"classId"`
	ClassName string `json:"className,omitempty"`
}

type BatchStudentInfo struct {
	Students  []map[string]any `json:"students,omitempty"`
	Operation string           `json:"operation,omitempty"`
}

// ── 题库分类 ──

type QuestionCategoryForm struct {
	QuestionCategory string `json:"questionCategory"`
	TeacherID        string `json:"teacherId,omitempty"`
}

// ── 题目 ──

type QuestionListQuery struct {
	PageNum      int    `json:"pageNum"`
	PageSize     int    `json:"pageSize"`
	Status       *int   `json:"status,omitempty"`
	Keyword      string `json:"keyword,omitempty"`
	CategoryID   *int   `json:"categoryId,omitempty"`
	Type         string `json:"type,omitempty"`
	IsAIGenerate *int   `json:"isAiGenerate,omitempty"`
	TeacherID    string `json:"teacherId,omitempty"`
}

func NewQuestionListQuery() QuestionListQuery {
	return QuestionListQuery{PageNum: 1, PageSize: 100}
}

func IntPtr(i int) *int { return &i }

type QuestionFind struct {
	StudentID  string `json:"studentId"`
	QuestionID int    `json:"questionId"`
	TestID     int    `json:"testId"`
}

type QuestionInsert struct {
	CategoryID   int    `json:"categoryId"`
	Type         string `json:"type"`
	QuestionText string `json:"questionText"`
	Answer       string `json:"answer"`
	Score        int    `json:"score"`
	AnswerTime   int    `json:"answerTime"`
	TeacherID    string `json:"teacherId"`
	Keywords     string `json:"keywords,omitempty"`
	IsAIGenerate bool   `json:"isAiGenerate,omitempty"`
}

type QuestionUpdate struct {
	ID           int    `json:"id"`
	CategoryID   int    `json:"categoryId,omitempty"`
	Type         string `json:"type,omitempty"`
	QuestionText string `json:"questionText,omitempty"`
	Answer       string `json:"answer,omitempty"`
	Score        int    `json:"score,omitempty"`
	AnswerTime   int    `json:"answerTime,omitempty"`
}

// ── 考试/测试 ──

type TestQuery struct {
	TeacherID string `json:"teacherId"`
	ClassID   *int   `json:"classId,omitempty"`
}

type TestPublish struct {
	TestID      *int   `json:"testId"`
	Title       string `json:"title"`
	ClassIDs    []int  `json:"classIds"`
	QuestionIDs []int  `json:"questionIds"`
	Duration    int    `json:"duration"`
	StartTime   string `json:"startTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	IsTest      int    `json:"isTest,omitempty"`
}

type ExtendTestTime struct {
	TestID    int    `json:"testId"`
	TeacherID string `json:"teacherId"`
	Minutes   int    `json:"minutes"`
}

// ── 考试结果 ──

type ResultStudent struct {
	StudentID string `json:"studentId"`
}

type ResultStart struct {
	StudentID string `json:"studentId"`
	TestID    int    `json:"testId"`
	IsTest    int    `json:"isTest,omitempty"`
}

// ── 作答 ──

type AnswerItem struct {
	StudentID  string `json:"studentId"`
	TestID     int    `json:"testId"`
	QuestionID int    `json:"questionId"`
	Answer     string `json:"answer,omitempty"`
}

type SaveAnswer struct {
	Items []AnswerItem
}

// ── 分数评语 ──

type ScoreItem struct {
	ID          string `json:"id"`
	QuestionID  int    `json:"questionId"`
	StudentID   string `json:"studentId"`
	TestID      int    `json:"testId"`
	ActualScore int    `json:"actualScore"`
	Comment     string `json:"comment,omitempty"`
}

type ScoreComment struct {
	Items []ScoreItem
}

type AIResultUpdate struct {
	StudentID string `json:"studentId"`
	TestID    int    `json:"testId"`
}

// ── AI ──

type AIAskItem struct {
	TestID        int    `json:"testId"`
	StudentID     string `json:"studentId"`
	QuestionID    int    `json:"questionId"`
	Index         int    `json:"index"`
	QuestionText  string `json:"questionText,omitempty"`
	StudentAnswer string `json:"studentAnswer,omitempty"`
	Score         int    `json:"score"`
	MaxScore      int    `json:"maxScore"`
}

type AIAsk struct {
	Questions []AIAskItem
}

// ── 视频 ──

type VideoForm struct {
	Title       string `json:"title"`
	URL         string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	CategoryID  int    `json:"categoryId,omitempty"`
}

type VideoListQuery struct {
	PageNum  int    `json:"pageNum"`
	PageSize int    `json:"pageSize"`
	Keyword  string `json:"keyword,omitempty"`
}

func NewVideoListQuery() VideoListQuery {
	return VideoListQuery{PageNum: 1, PageSize: 20}
}

// ── 考勤 ──

type AttendanceSave struct {
	ClassID      int    `json:"classId"`
	StudentID    string `json:"studentId,omitempty"`
	Status       string `json:"status,omitempty"`
	Date         string `json:"date,omitempty"`
	IsAttendance bool   `json:"isAttendance,omitempty"`
}

type AttendanceUpdate struct {
	Timestamp    string `json:"timestamp,omitempty"`
	StudentID    string `json:"studentId,omitempty"`
	IsAttendance bool   `json:"isAttendance,omitempty"`
	LoginTime    string `json:"loginTime,omitempty"`
}

type AttendanceStatus struct {
	StudentID    string `json:"studentId,omitempty"`
	IsAttendance bool   `json:"isAttendance,omitempty"`
	Date         string `json:"date,omitempty"`
}
