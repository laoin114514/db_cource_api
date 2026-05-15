package go_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const DefaultBaseURL = "http://172.21.44.162:5174/acms"

// ACMSClient 数据库原理智慧课程 API 客户端
type ACMSClient struct {
	BaseURL  string
	token    string
	UserInfo map[string]any
}

func NewClient() *ACMSClient {
	return &ACMSClient{BaseURL: DefaultBaseURL}
}

// ═══════════════════════════════════════════════════════════════
// 内部 HTTP 方法
// ═══════════════════════════════════════════════════════════════

func doRequest[T any](c *ACMSClient, method, path string, params, body any) (ApiResponse[T], error) {
	fullURL := c.BaseURL + path
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return ApiResponse[T]{}, err
		}
		bodyReader = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, fullURL, bodyReader)
	if err != nil {
		return ApiResponse[T]{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	if params != nil {
		q := req.URL.Query()
		for k, v := range toQueryMap(params) {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ApiResponse[T]{}, err
	}
	defer resp.Body.Close()
	return decodeResponse[T](resp)
}

func decodeResponse[T any](resp *http.Response) (ApiResponse[T], error) {
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return ApiResponse[T]{}, err
	}
	var result ApiResponse[T]
	if err := json.Unmarshal(raw, &result); err != nil {
		return ApiResponse[T]{}, fmt.Errorf("解析响应失败: %w, body=%s", err, string(raw[:min(len(raw), 200)]))
	}
	return result, nil
}

func doGETBytes(c *ACMSClient, path string, params any) ([]byte, error) {
	fullURL := c.BaseURL + path
	req, _ := http.NewRequest(http.MethodGet, fullURL, nil)
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	if params != nil {
		q := req.URL.Query()
		for k, v := range toQueryMap(params) {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func doUpload(c *ACMSClient, path, fieldName, filePath string) (ApiResponse[any], error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	part, _ := w.CreateFormFile(fieldName, filePath)
	f, err := os.Open(filePath)
	if err != nil {
		return ApiResponse[any]{}, err
	}
	defer f.Close()
	io.Copy(part, f)
	w.Close()

	req, _ := http.NewRequest(http.MethodPost, c.BaseURL+path, &buf)
	req.Header.Set("Content-Type", w.FormDataContentType())
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ApiResponse[any]{}, err
	}
	defer resp.Body.Close()
	return decodeResponse[any](resp)
}

func toQueryMap(params any) map[string]string {
	result := make(map[string]string)
	if params == nil {
		return result
	}
	b, _ := json.Marshal(params)
	var m map[string]any
	json.Unmarshal(b, &m)
	for k, val := range m {
		if s, ok := val.(string); ok {
			result[k] = s
			continue
		}
		result[k] = fmt.Sprintf("%v", val)
	}
	return result
}

// ═══════════════════════════════════════════════════════════════
// 认证
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) Login(form LoginForm) (ApiResponse[LoginData], error) {
	r, err := doRequest[LoginData](c, http.MethodPost, "/auth/login", nil, form)
	if err == nil && r.Code == 200 {
		c.token = r.Data.Token
		c.UserInfo = r.Data.User
	}
	return r, err
}

func (c *ACMSClient) Logout() {
	c.token = ""
	c.UserInfo = nil
}

func (c *ACMSClient) IsLoggedIn() bool { return c.token != "" }

// ═══════════════════════════════════════════════════════════════
// 访问量
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) GetVisitCount() (ApiResponse[int], error) {
	return doRequest[int](c, http.MethodGet, "/visit/count", nil, nil)
}

func (c *ACMSClient) IncrementVisit() (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/visit/increment", nil, nil)
}

// ═══════════════════════════════════════════════════════════════
// 用户
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) QueryUser(id string) (ApiResponse[UserData], error) {
	return doRequest[UserData](c, http.MethodGet, "/user/query", map[string]string{"id": id}, nil)
}

func (c *ACMSClient) GetAllUsers() (ApiResponse[[]UserData], error) {
	return doRequest[[]UserData](c, http.MethodGet, "/user/all", nil, nil)
}

func (c *ACMSClient) InsertUser(form UserInsert) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/user/insert", nil, form)
}

func (c *ACMSClient) UpdateUser(form UserUpdate) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/user/update", nil, form)
}

func (c *ACMSClient) DeleteUser(userID string) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodDelete, "/user/"+userID, nil, nil)
}

// ═══════════════════════════════════════════════════════════════
// 角色
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) GetAllRoles() (ApiResponse[[]any], error) {
	return doRequest[[]any](c, http.MethodGet, "/role/all", nil, nil)
}

func (c *ACMSClient) AddRole(form RoleForm) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/role/add", nil, form)
}

func (c *ACMSClient) UpdateRole(form RoleUpdate) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPut, "/role/update", nil, form)
}

// ═══════════════════════════════════════════════════════════════
// 教师
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) QueryTeacher(teacherID string) (ApiResponse[TeacherData], error) {
	return doRequest[TeacherData](c, http.MethodGet, "/teacher/query",
		map[string]string{"teacherId": teacherID}, nil)
}

func (c *ACMSClient) GetAllTeachers() (ApiResponse[[]TeacherData], error) {
	return doRequest[[]TeacherData](c, http.MethodGet, "/teacher/findAll", nil, nil)
}

func (c *ACMSClient) GetTeacherInfo() (ApiResponse[[]TeacherBrief], error) {
	return doRequest[[]TeacherBrief](c, http.MethodGet, "/teacher/acquireInfo", nil, nil)
}

func (c *ACMSClient) GetTeacherPage(p PageQuery) (ApiResponse[[]TeacherData], error) {
	return doRequest[[]TeacherData](c, http.MethodGet, "/teacher/showPage", p, nil)
}

func (c *ACMSClient) InsertTeacher(form TeacherForm) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/teacher/insert", nil, form)
}

func (c *ACMSClient) UpdateTeacher(form TeacherForm) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/teacher/update", nil, map[string]any{
		"teacher":  map[string]string{"teacherId": form.TeacherID, "name": form.Name, "gender": form.Gender, "position": form.Position},
		"password": form.Password,
	})
}

func (c *ACMSClient) TeacherJoinClass(rel TeacherClassRelation) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/teacher/join", nil, rel)
}

func (c *ACMSClient) TeacherLeaveClass(rel TeacherClassRelation) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/teacher/remove", nil, rel)
}

func (c *ACMSClient) TeacherBatchInfo(payload map[string]any) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/teacher/batchInfo", nil, payload)
}

func (c *ACMSClient) DeleteTeacher(teacherID string) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodDelete, "/teacher/"+teacherID, nil, nil)
}

// ═══════════════════════════════════════════════════════════════
// 班级
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) QueryClass(className string) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/class/queryClass",
		map[string]string{"className": className}, nil)
}

func (c *ACMSClient) GetClassNames() (ApiResponse[[]string], error) {
	return doRequest[[]string](c, http.MethodGet, "/class/acquireName", nil, nil)
}

func (c *ACMSClient) GetClassPage(p PageQuery) (ApiResponse[[]ClassData], error) {
	return doRequest[[]ClassData](c, http.MethodGet, "/class/showPage", p, nil)
}

func (c *ACMSClient) FindClassByTeacher(teacherID string) (ApiResponse[[]ClassData], error) {
	return doRequest[[]ClassData](c, http.MethodGet, "/class/findByTeacher",
		map[string]string{"teacherId": teacherID}, nil)
}

func (c *ACMSClient) GetAllClasses() (ApiResponse[[]ClassData], error) {
	return doRequest[[]ClassData](c, http.MethodGet, "/class/findAll", nil, nil)
}

func (c *ACMSClient) GetClassByName(className string) (ApiResponse[ClassData], error) {
	return doRequest[ClassData](c, http.MethodGet, "/class/getByName",
		map[string]string{"className": className}, nil)
}

func (c *ACMSClient) FindClassByTeacherAndName(teacherID, className string) (ApiResponse[[]ClassData], error) {
	return doRequest[[]ClassData](c, http.MethodGet, "/class/findByTeacherAndName",
		map[string]string{"teacherId": teacherID, "className": className}, nil)
}

func (c *ACMSClient) GetClassStudents(classID int) (ApiResponse[[]StudentData], error) {
	return doRequest[[]StudentData](c, http.MethodGet, "/class/students",
		map[string]int{"classId": classID}, nil)
}

func (c *ACMSClient) InsertClass(form ClassForm) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/class/insert", nil, form)
}

func (c *ACMSClient) UpdateClass(form ClassUpdate) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPut, "/class/update", nil, form)
}

func (c *ACMSClient) FindClassByStudent(studentID string) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/class/findByStudent",
		map[string]string{"studentId": studentID}, nil)
}

func (c *ACMSClient) DeleteClass(classID int) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodDelete, fmt.Sprintf("/class/%d", classID), nil, nil)
}

// ═══════════════════════════════════════════════════════════════
// 学生
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) QueryStudent(studentID string) (ApiResponse[StudentData], error) {
	return doRequest[StudentData](c, http.MethodGet, "/student/query",
		map[string]string{"studentId": studentID}, nil)
}

func (c *ACMSClient) GetStudentPage(p PageQuery) (ApiResponse[[]StudentData], error) {
	return doRequest[[]StudentData](c, http.MethodGet, "/student/showPage", p, nil)
}

func (c *ACMSClient) FindNoClassStudents() (ApiResponse[[]StudentData], error) {
	return doRequest[[]StudentData](c, http.MethodGet, "/student/findNoClass", nil, nil)
}

func (c *ACMSClient) SearchStudentByPrefix(prefix string) (ApiResponse[[]StudentData], error) {
	return doRequest[[]StudentData](c, http.MethodGet, "/student/searchByStudentIdPrefix",
		map[string]string{"prefix": prefix}, nil)
}

func (c *ACMSClient) InsertStudent(form StudentForm) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/student/insert", nil, form)
}

func (c *ACMSClient) UpdateStudent(form StudentUpdate) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPut, "/student/update", nil, map[string]any{"student": form})
}

func (c *ACMSClient) DeleteStudent(studentIDs []int) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/student/delete", nil, studentIDs)
}

func (c *ACMSClient) DeleteStudentByID(studentID string) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodDelete, "/student/"+studentID, nil, nil)
}

func (c *ACMSClient) UpdateStudentClass(studentID string, classID *int, className string) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPut, "/student/updateClass", nil, map[string]any{
		"studentId": studentID, "classId": classID, "className": className,
	})
}

func (c *ACMSClient) BatchStudentInfo(form BatchStudentInfo) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/student/batchInfo", nil, form)
}

func (c *ACMSClient) ExportStudents(classID int) ([]byte, error) {
	return doGETBytes(c, "/student/export", map[string]int{"classId": classID})
}

func (c *ACMSClient) DownloadStudentTemplate() ([]byte, error) {
	return doGETBytes(c, "/student/downloadTemplate", nil)
}

func (c *ACMSClient) BatchImportStudents(filePath string) (ApiResponse[any], error) {
	return doUpload(c, "/student/batchImport", "file", filePath)
}

func (c *ACMSClient) UploadStudentFile(filePath string) (ApiResponse[any], error) {
	return doUpload(c, "/studentFile/upload", "file", filePath)
}

// ═══════════════════════════════════════════════════════════════
// 题库分类
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) GetQuestionCategories() (ApiResponse[[]QuestionCategoryData], error) {
	return doRequest[[]QuestionCategoryData](c, http.MethodGet, "/question-category/list", nil, nil)
}

func (c *ACMSClient) AddQuestionCategory(form QuestionCategoryForm) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/question-category/add", nil, form)
}

// ═══════════════════════════════════════════════════════════════
// 题目
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) GetQuestionList(q QuestionListQuery) (ApiResponse[QuestionPageData], error) {
	return doRequest[QuestionPageData](c, http.MethodGet, "/question/list", q, nil)
}

func (c *ACMSClient) GetAllQuestions(teacherID string, categoryID *int) ([]QuestionData, error) {
	q := QuestionListQuery{PageNum: 1, PageSize: 500, TeacherID: teacherID, CategoryID: categoryID}
	r, err := c.GetQuestionList(q)
	if err != nil || !r.OK() {
		return nil, err
	}
	return r.Data.Items, nil
}

func (c *ACMSClient) QueryQuestion(questionID int) (ApiResponse[QuestionData], error) {
	return doRequest[QuestionData](c, http.MethodGet, "/question/query",
		map[string]int{"id": questionID}, nil)
}

func (c *ACMSClient) FindQuestion(studentID string, questionID, testID int) (ApiResponse[QuestionData], error) {
	return doRequest[QuestionData](c, http.MethodGet, "/question/find",
		map[string]any{"studentId": studentID, "questionId": questionID, "testId": testID}, nil)
}

func (c *ACMSClient) InsertQuestion(form QuestionInsert) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/question/insert", nil, form)
}

func (c *ACMSClient) UpdateQuestion(form QuestionUpdate) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPut, "/question/update", nil, form)
}

func (c *ACMSClient) BatchDeleteQuestions(ids []int) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/question/batchDelete", nil, ids)
}

func (c *ACMSClient) DownloadQuestionTemplate() ([]byte, error) {
	return doGETBytes(c, "/questionFile/template/download", nil)
}

func (c *ACMSClient) UploadQuestionFile(filePath string) (ApiResponse[any], error) {
	return doUpload(c, "/questionFile/upload", "file", filePath)
}

// ═══════════════════════════════════════════════════════════════
// 考试/测试
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) GetTeacherTests(teacherID string, classID *int) (ApiResponse[[]TestData], error) {
	return doRequest[[]TestData](c, http.MethodGet, "/test/getTeacherTest",
		map[string]any{"teacherId": teacherID, "classId": classID}, nil)
}

func (c *ACMSClient) GetClassTest(classID int) (ApiResponse[[]TestData], error) {
	return doRequest[[]TestData](c, http.MethodGet, "/test/getClassTest",
		map[string]int{"classId": classID}, nil)
}

func (c *ACMSClient) GetStudentTest(studentID string) (ApiResponse[[]TestData], error) {
	return doRequest[[]TestData](c, http.MethodGet, "/test/getStudentTest",
		map[string]string{"studentId": studentID}, nil)
}

func (c *ACMSClient) PublishTest(form TestPublish) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/test/publish", nil, form)
}

func (c *ACMSClient) GetPublishTeacher(teacherID string) (ApiResponse[[]TeacherPublishData], error) {
	return doRequest[[]TeacherPublishData](c, http.MethodGet, "/test/publish/teacher",
		map[string]string{"teacherId": teacherID}, nil)
}

func (c *ACMSClient) GetPublishAI(params map[string]any) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/test/publish/ai", params, nil)
}

func (c *ACMSClient) GetPublishClass(params map[string]any) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/test/publish/class", params, nil)
}

func (c *ACMSClient) GetPublishStudent(params map[string]any) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/test/publish/student", params, nil)
}

func (c *ACMSClient) ExtendTestTime(form ExtendTestTime) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/test/extendTestTime", nil, form)
}

func (c *ACMSClient) ExportStudentReport(studentID string, testID int) ([]byte, error) {
	return doGETBytes(c, "/test/exportStudentReport/excel",
		map[string]any{"studentId": studentID, "testId": testID})
}

func (c *ACMSClient) ExportClassAvgScores(testID int) ([]byte, error) {
	return doGETBytes(c, "/test/exportClassAverageScores/excel",
		map[string]int{"testId": testID})
}

func (c *ACMSClient) ExportFilteredScoreDetail(testID int) ([]byte, error) {
	return doGETBytes(c, "/test/exportFilteredScoreDetail/excel",
		map[string]int{"testId": testID})
}

// ═══════════════════════════════════════════════════════════════
// 考试结果
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) GetStudentResult(studentID string) (ApiResponse[[]StudentResultGroup], error) {
	return doRequest[[]StudentResultGroup](c, http.MethodGet, "/result/student",
		map[string]string{"studentId": studentID}, nil)
}

func (c *ACMSClient) StartResult(studentID string, testID, isTest int) (ApiResponse[string], error) {
	return doRequest[string](c, http.MethodGet, "/result/start",
		map[string]any{"studentId": studentID, "testId": testID, "isTest": isTest}, nil)
}

func (c *ACMSClient) GetResultDetail(studentID string, testID int) (ApiResponse[[]ResultDetailData], error) {
	return doRequest[[]ResultDetailData](c, http.MethodGet, "/result/detail",
		map[string]any{"studentId": studentID, "testId": testID}, nil)
}

func (c *ACMSClient) FindResult(params map[string]any) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/result/find", params, nil)
}

func (c *ACMSClient) GetAIResultDetail(studentID string, testID int) (ApiResponse[[]ResultDetailData], error) {
	return doRequest[[]ResultDetailData](c, http.MethodGet, "/result/ai/detail",
		map[string]any{"studentId": studentID, "testId": testID}, nil)
}

func (c *ACMSClient) SaveAnswer(items []AnswerItem) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPut, "/result/saveAnswer", nil, items)
}

func (c *ACMSClient) UpdateAnswer(items []AnswerItem) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPut, "/result/updateAnswer", nil, items)
}

func (c *ACMSClient) UpdateScoreComment(items []ScoreItem) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/result/updateScoreAndComment", nil, items)
}

func (c *ACMSClient) UpdateAIResult(form AIResultUpdate) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/result/update/ai", nil, form)
}

// ═══════════════════════════════════════════════════════════════
// AI
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) AIStudentAsk(items []AIAskItem) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/ai/student/ask", nil, map[string]any{"questions": items})
}

func (c *ACMSClient) AITeacherAsk(question, context string) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/ai/teacher/ask", nil,
		map[string]string{"question": question, "context": context})
}

// ═══════════════════════════════════════════════════════════════
// 视频
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) GetVideoList(q VideoListQuery) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/video/list", q, nil)
}

func (c *ACMSClient) GetVideoTree() (ApiResponse[[]VideoTreeNode], error) {
	return doRequest[[]VideoTreeNode](c, http.MethodGet, "/video/tree", nil, nil)
}

func (c *ACMSClient) GetVideoTreeSelect() (ApiResponse[[]VideoTreeNode], error) {
	return doRequest[[]VideoTreeNode](c, http.MethodGet, "/video/tree/select", nil, nil)
}

func (c *ACMSClient) AddVideo(form VideoForm) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/video/add", nil, form)
}

func (c *ACMSClient) UploadVideo(filePath string) (ApiResponse[any], error) {
	return doUpload(c, "/video/upload", "file", filePath)
}

// ═══════════════════════════════════════════════════════════════
// 考勤
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) GetAttendanceHistory(classID int) (ApiResponse[[]AttendanceRecord], error) {
	return doRequest[[]AttendanceRecord](c, http.MethodGet, "/attendance/history",
		map[string]int{"classId": classID}, nil)
}

func (c *ACMSClient) SaveAttendance(form AttendanceSave) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/attendance/save", nil, form)
}

func (c *ACMSClient) UpdateAttendance(form AttendanceUpdate) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/attendance/update", nil, form)
}

func (c *ACMSClient) UpdateAttendanceStatus(form AttendanceStatus) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/attendance/updateStatus", nil, form)
}

func (c *ACMSClient) ExportAttendance(classID int) ([]byte, error) {
	return doGETBytes(c, "/attendance/export", map[string]int{"classId": classID})
}

func (c *ACMSClient) ExportAttendanceSingle(recordID int) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/attendance/exportSingle",
		map[string]int{"id": recordID}, nil)
}

// ═══════════════════════════════════════════════════════════════
// 校历
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) GetSemesters() (ApiResponse[[]SemesterData], error) {
	return doRequest[[]SemesterData](c, http.MethodGet, "/academicCalendar/semesters", nil, nil)
}

// ═══════════════════════════════════════════════════════════════
// 知识图谱
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) GetKnowledgeGraph() (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/knowledgeGraph/data", nil, nil)
}

func (c *ACMSClient) GetPendingKnowledge() (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/knowledgeGraph/pendingKnowledge", nil, nil)
}

func (c *ACMSClient) UploadKnowledgeGraph(filePath string) (ApiResponse[any], error) {
	return doUpload(c, "/upload/knowledge-graph", "file", filePath)
}

// ═══════════════════════════════════════════════════════════════
// 分组
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) GetGroupStudent(params map[string]any) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/group/student", params, nil)
}

func (c *ACMSClient) GetGroupTeacher(params map[string]any) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodGet, "/group/teacher", params, nil)
}

func (c *ACMSClient) GenerateGroup(payload map[string]any) (ApiResponse[any], error) {
	return doRequest[any](c, http.MethodPost, "/group/generate", nil, payload)
}

// ═══════════════════════════════════════════════════════════════
// 上传
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) UploadImage(filePath string) (ApiResponse[any], error) {
	return doUpload(c, "/upload/image", "file", filePath)
}

// ═══════════════════════════════════════════════════════════════
// 便捷函数
// ═══════════════════════════════════════════════════════════════

func CreateClient(userID, password string) (*ACMSClient, error) {
	c := NewClient()
	r, err := c.Login(LoginForm{ID: userID, Password: password})
	if err != nil {
		return nil, err
	}
	if !r.OK() {
		return nil, fmt.Errorf("登录失败: %s", r.Msg)
	}
	return c, nil
}
