package acms_sdk

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

// NewClient 创建客户端
func NewClient() *ACMSClient {
	return &ACMSClient{BaseURL: DefaultBaseURL}
}

// ── 认证 ──────────────────────────────────────────────

// Login 登录获取 token
func (c *ACMSClient) Login(form LoginForm) (map[string]any, error) {
	data, err := c.doPOST("/auth/login", form)
	if err != nil {
		return nil, err
	}
	if code, _ := data["code"].(float64); code == 200 {
		if d, ok := data["data"].(map[string]any); ok {
			if tok, ok := d["token"].(string); ok {
				c.token = tok
			}
			if u, ok := d["user"].(map[string]any); ok {
				c.UserInfo = u
			}
		}
	}
	return data, nil
}

// Logout 清除登录状态
func (c *ACMSClient) Logout() {
	c.token = ""
	c.UserInfo = nil
}

// IsLoggedIn 检查是否已登录
func (c *ACMSClient) IsLoggedIn() bool {
	return c.token != ""
}

// ── 访问量 ────────────────────────────────────────────

func (c *ACMSClient) GetVisitCount() (map[string]any, error) {
	return c.doGET("/visit/count", nil)
}

func (c *ACMSClient) IncrementVisit() (map[string]any, error) {
	return c.doPOST("/visit/increment", nil)
}

// ── 用户 ──────────────────────────────────────────────

func (c *ACMSClient) QueryUser(q UserQuery) (map[string]any, error) {
	return c.doGET("/user/query", map[string]string{"id": q.ID})
}

func (c *ACMSClient) GetAllUsers() (map[string]any, error) {
	return c.doGET("/user/all", nil)
}

func (c *ACMSClient) InsertUser(form UserInsert) (map[string]any, error) {
	return c.doPOST("/user/insert", form)
}

func (c *ACMSClient) UpdateUser(form UserUpdate) (map[string]any, error) {
	return c.doPOST("/user/update", form)
}

func (c *ACMSClient) DeleteUser(userID string) (map[string]any, error) {
	return c.doDELETE(fmt.Sprintf("/user/%s", userID))
}

// ── 角色 ──────────────────────────────────────────────

func (c *ACMSClient) GetAllRoles() (map[string]any, error) {
	return c.doGET("/role/all", nil)
}

func (c *ACMSClient) AddRole(form RoleForm) (map[string]any, error) {
	return c.doPOST("/role/add", form)
}

func (c *ACMSClient) UpdateRole(form RoleUpdate) (map[string]any, error) {
	return c.doPUT("/role/update", form)
}

// ── 教师 ──────────────────────────────────────────────

func (c *ACMSClient) QueryTeacher(q TeacherQuery) (map[string]any, error) {
	return c.doGET("/teacher/query", map[string]string{"teacherId": q.TeacherID})
}

func (c *ACMSClient) GetAllTeachers() (map[string]any, error) {
	return c.doGET("/teacher/findAll", nil)
}

func (c *ACMSClient) GetTeacherInfo() (map[string]any, error) {
	return c.doGET("/teacher/acquireInfo", nil)
}

func (c *ACMSClient) GetTeacherPage(p PageQuery) (map[string]any, error) {
	return c.doGET("/teacher/showPage", p)
}

func (c *ACMSClient) InsertTeacher(form TeacherForm) (map[string]any, error) {
	return c.doPOST("/teacher/insert", form)
}

func (c *ACMSClient) UpdateTeacher(form TeacherForm) (map[string]any, error) {
	return c.doPOST("/teacher/update", map[string]any{
		"teacher": map[string]string{
			"teacherId": form.TeacherID,
			"name":      form.Name,
			"gender":    form.Gender,
			"position":  form.Position,
		},
		"password": form.Password,
	})
}

func (c *ACMSClient) TeacherJoinClass(rel TeacherClassRelation) (map[string]any, error) {
	return c.doPOST("/teacher/join", rel)
}

func (c *ACMSClient) TeacherLeaveClass(rel TeacherClassRelation) (map[string]any, error) {
	return c.doPOST("/teacher/remove", rel)
}

func (c *ACMSClient) TeacherBatchInfo(payload map[string]any) (map[string]any, error) {
	return c.doPOST("/teacher/batchInfo", payload)
}

func (c *ACMSClient) DeleteTeacher(teacherID string) (map[string]any, error) {
	return c.doDELETE(fmt.Sprintf("/teacher/%s", teacherID))
}

// ── 班级 ──────────────────────────────────────────────

func (c *ACMSClient) QueryClass(q ClassQuery) (map[string]any, error) {
	return c.doGET("/class/queryClass", map[string]string{"className": q.ClassName})
}

func (c *ACMSClient) GetClassNames() (map[string]any, error) {
	return c.doGET("/class/acquireName", nil)
}

func (c *ACMSClient) GetClassPage(p PageQuery) (map[string]any, error) {
	return c.doGET("/class/showPage", p)
}

func (c *ACMSClient) FindClassByTeacher(teacherID string) (map[string]any, error) {
	return c.doGET("/class/findByTeacher", map[string]string{"teacherId": teacherID})
}

func (c *ACMSClient) GetAllClasses() (map[string]any, error) {
	return c.doGET("/class/findAll", nil)
}

func (c *ACMSClient) GetClassByName(className string) (map[string]any, error) {
	return c.doGET("/class/getByName", map[string]string{"className": className})
}

func (c *ACMSClient) FindClassByTeacherAndName(teacherID, className string) (map[string]any, error) {
	return c.doGET("/class/findByTeacherAndName", map[string]string{
		"teacherId": teacherID, "className": className,
	})
}

func (c *ACMSClient) GetClassStudents(classID int) (map[string]any, error) {
	return c.doGET("/class/students", map[string]int{"classId": classID})
}

func (c *ACMSClient) InsertClass(form ClassForm) (map[string]any, error) {
	return c.doPOST("/class/insert", form)
}

func (c *ACMSClient) UpdateClass(form ClassUpdate) (map[string]any, error) {
	return c.doPUT("/class/update", form)
}

func (c *ACMSClient) FindClassByStudent(studentID string) (map[string]any, error) {
	return c.doGET("/class/findByStudent", map[string]string{"studentId": studentID})
}

func (c *ACMSClient) DeleteClass(classID int) (map[string]any, error) {
	return c.doDELETE(fmt.Sprintf("/class/%d", classID))
}

// ── 学生 ──────────────────────────────────────────────

func (c *ACMSClient) QueryStudent(q StudentQuery) (map[string]any, error) {
	return c.doGET("/student/query", map[string]string{"studentId": q.StudentID})
}

func (c *ACMSClient) GetStudentPage(p PageQuery) (map[string]any, error) {
	return c.doGET("/student/showPage", p)
}

func (c *ACMSClient) FindNoClassStudents() (map[string]any, error) {
	return c.doGET("/student/findNoClass", nil)
}

func (c *ACMSClient) SearchStudentByPrefix(prefix string) (map[string]any, error) {
	return c.doGET("/student/searchByStudentIdPrefix", map[string]string{"prefix": prefix})
}

func (c *ACMSClient) InsertStudent(form StudentForm) (map[string]any, error) {
	return c.doPOST("/student/insert", form)
}

func (c *ACMSClient) UpdateStudent(form StudentUpdate) (map[string]any, error) {
	return c.doPUT("/student/update", map[string]any{"student": form})
}

func (c *ACMSClient) DeleteStudent(studentIDs []int) (map[string]any, error) {
	return c.doPOST("/student/delete", studentIDs)
}

func (c *ACMSClient) DeleteStudentByID(studentID string) (map[string]any, error) {
	return c.doDELETE(fmt.Sprintf("/student/%s", studentID))
}

func (c *ACMSClient) UpdateStudentClass(studentID string, classID *int, className string) (map[string]any, error) {
	return c.doPUT("/student/updateClass", map[string]any{
		"studentId": studentID, "classId": classID, "className": className,
	})
}

func (c *ACMSClient) BatchStudentInfo(form BatchStudentInfo) (map[string]any, error) {
	return c.doPOST("/student/batchInfo", form)
}

func (c *ACMSClient) ExportStudents(classID int) ([]byte, error) {
	return c.doGETBytes("/student/export", map[string]int{"classId": classID})
}

func (c *ACMSClient) DownloadStudentTemplate() ([]byte, error) {
	return c.doGETBytes("/student/downloadTemplate", nil)
}

func (c *ACMSClient) BatchImportStudents(filePath string) (map[string]any, error) {
	return c.uploadFile("/student/batchImport", "file", filePath)
}

func (c *ACMSClient) UploadStudentFile(filePath string) (map[string]any, error) {
	return c.uploadFile("/studentFile/upload", "file", filePath)
}

// ── 题库分类 ──────────────────────────────────────────

func (c *ACMSClient) GetQuestionCategories() (map[string]any, error) {
	return c.doGET("/question-category/list", nil)
}

func (c *ACMSClient) AddQuestionCategory(form QuestionCategoryForm) (map[string]any, error) {
	return c.doPOST("/question-category/add", form)
}

// ── 题目 ──────────────────────────────────────────────

func (c *ACMSClient) GetQuestionList(q QuestionListQuery) (map[string]any, error) {
	return c.doGET("/question/list", q)
}

func (c *ACMSClient) GetAllQuestions(teacherID string, categoryID *int) ([]map[string]any, error) {
	q := QuestionListQuery{PageNum: 1, PageSize: 500, TeacherID: teacherID, CategoryID: categoryID}
	resp, err := c.GetQuestionList(q)
	if err != nil {
		return nil, err
	}
	if code, _ := resp["code"].(float64); code != 200 {
		return nil, nil
	}
	data, _ := resp["data"].(map[string]any)
	items, _ := data["items"].([]any)
	result := make([]map[string]any, len(items))
	for i, item := range items {
		result[i], _ = item.(map[string]any)
	}
	return result, nil
}

func (c *ACMSClient) QueryQuestion(questionID int) (map[string]any, error) {
	return c.doGET("/question/query", map[string]int{"id": questionID})
}

func (c *ACMSClient) FindQuestion(q QuestionFind) (map[string]any, error) {
	return c.doGET("/question/find", q)
}

func (c *ACMSClient) InsertQuestion(form QuestionInsert) (map[string]any, error) {
	return c.doPOST("/question/insert", form)
}

func (c *ACMSClient) UpdateQuestion(form QuestionUpdate) (map[string]any, error) {
	return c.doPUT("/question/update", form)
}

func (c *ACMSClient) BatchDeleteQuestions(ids []int) (map[string]any, error) {
	return c.doPOST("/question/batchDelete", ids)
}

func (c *ACMSClient) DownloadQuestionTemplate() ([]byte, error) {
	return c.doGETBytes("/questionFile/template/download", nil)
}

func (c *ACMSClient) UploadQuestionFile(filePath string) (map[string]any, error) {
	return c.uploadFile("/questionFile/upload", "file", filePath)
}

// ── 考试/测试 ─────────────────────────────────────────

func (c *ACMSClient) GetTeacherTests(q TestQuery) (map[string]any, error) {
	return c.doGET("/test/getTeacherTest", q)
}

func (c *ACMSClient) GetClassTest(classID int) (map[string]any, error) {
	return c.doGET("/test/getClassTest", map[string]int{"classId": classID})
}

func (c *ACMSClient) GetStudentTest(studentID string) (map[string]any, error) {
	return c.doGET("/test/getStudentTest", map[string]string{"studentId": studentID})
}

func (c *ACMSClient) PublishTest(form TestPublish) (map[string]any, error) {
	return c.doPOST("/test/publish", form)
}

func (c *ACMSClient) GetPublishTeacher(teacherID string) (map[string]any, error) {
	return c.doGET("/test/publish/teacher", map[string]string{"teacherId": teacherID})
}

func (c *ACMSClient) GetPublishAI(params map[string]any) (map[string]any, error) {
	return c.doGET("/test/publish/ai", params)
}

func (c *ACMSClient) GetPublishClass(params map[string]any) (map[string]any, error) {
	return c.doGET("/test/publish/class", params)
}

func (c *ACMSClient) GetPublishStudent(params map[string]any) (map[string]any, error) {
	return c.doGET("/test/publish/student", params)
}

func (c *ACMSClient) ExtendTestTime(form ExtendTestTime) (map[string]any, error) {
	return c.doPOST("/test/extendTestTime", form)
}

func (c *ACMSClient) ExportStudentReport(studentID string, testID int) ([]byte, error) {
	return c.doGETBytes("/test/exportStudentReport/excel",
		map[string]any{"studentId": studentID, "testId": testID})
}

func (c *ACMSClient) ExportClassAvgScores(testID int) ([]byte, error) {
	return c.doGETBytes("/test/exportClassAverageScores/excel",
		map[string]int{"testId": testID})
}

func (c *ACMSClient) ExportFilteredScoreDetail(testID int) ([]byte, error) {
	return c.doGETBytes("/test/exportFilteredScoreDetail/excel",
		map[string]int{"testId": testID})
}

// ── 考试结果 ──────────────────────────────────────────

func (c *ACMSClient) GetStudentResult(q ResultStudent) (map[string]any, error) {
	return c.doGET("/result/student", q)
}

func (c *ACMSClient) StartResult(q ResultStart) (map[string]any, error) {
	return c.doGET("/result/start", q)
}

func (c *ACMSClient) GetResultDetail(studentID string, testID int) (map[string]any, error) {
	return c.doGET("/result/detail", map[string]any{"studentId": studentID, "testId": testID})
}

func (c *ACMSClient) FindResult(params map[string]any) (map[string]any, error) {
	return c.doGET("/result/find", params)
}

func (c *ACMSClient) GetAIResultDetail(studentID string, testID int) (map[string]any, error) {
	return c.doGET("/result/ai/detail", map[string]any{"studentId": studentID, "testId": testID})
}

func (c *ACMSClient) SaveAnswer(items []AnswerItem) (map[string]any, error) {
	return c.doPUT("/result/saveAnswer", items)
}

func (c *ACMSClient) UpdateAnswer(items []AnswerItem) (map[string]any, error) {
	return c.doPUT("/result/updateAnswer", items)
}

func (c *ACMSClient) UpdateScoreComment(items []ScoreItem) (map[string]any, error) {
	return c.doPOST("/result/updateScoreAndComment", items)
}

func (c *ACMSClient) UpdateAIResult(form AIResultUpdate) (map[string]any, error) {
	return c.doPOST("/result/update/ai", form)
}

// ── AI ────────────────────────────────────────────────

func (c *ACMSClient) AIStudentAsk(items []AIAskItem) (map[string]any, error) {
	return c.doPOST("/ai/student/ask", map[string]any{"questions": items})
}

func (c *ACMSClient) AITeacherAsk(question, context string) (map[string]any, error) {
	return c.doPOST("/ai/teacher/ask", map[string]any{
		"question": question, "context": context,
	})
}

// ── 视频 ──────────────────────────────────────────────

func (c *ACMSClient) GetVideoList(q VideoListQuery) (map[string]any, error) {
	return c.doGET("/video/list", q)
}

func (c *ACMSClient) GetVideoTree() (map[string]any, error) {
	return c.doGET("/video/tree", nil)
}

func (c *ACMSClient) GetVideoTreeSelect() (map[string]any, error) {
	return c.doGET("/video/tree/select", nil)
}

func (c *ACMSClient) AddVideo(form VideoForm) (map[string]any, error) {
	return c.doPOST("/video/add", form)
}

func (c *ACMSClient) UploadVideo(filePath string) (map[string]any, error) {
	return c.uploadFile("/video/upload", "file", filePath)
}

// ── 考勤 ──────────────────────────────────────────────

func (c *ACMSClient) GetAttendanceHistory(classID int) (map[string]any, error) {
	return c.doGET("/attendance/history", map[string]int{"classId": classID})
}

func (c *ACMSClient) SaveAttendance(form AttendanceSave) (map[string]any, error) {
	return c.doPOST("/attendance/save", form)
}

func (c *ACMSClient) UpdateAttendance(form AttendanceUpdate) (map[string]any, error) {
	return c.doPOST("/attendance/update", form)
}

func (c *ACMSClient) UpdateAttendanceStatus(form AttendanceStatus) (map[string]any, error) {
	return c.doPOST("/attendance/updateStatus", form)
}

func (c *ACMSClient) ExportAttendance(classID int) ([]byte, error) {
	return c.doGETBytes("/attendance/export", map[string]int{"classId": classID})
}

func (c *ACMSClient) ExportAttendanceSingle(recordID int) (map[string]any, error) {
	return c.doGET("/attendance/exportSingle", map[string]int{"id": recordID})
}

// ── 校历 ──────────────────────────────────────────────

func (c *ACMSClient) GetSemesters() (map[string]any, error) {
	return c.doGET("/academicCalendar/semesters", nil)
}

// ── 知识图谱 ──────────────────────────────────────────

func (c *ACMSClient) GetKnowledgeGraph() (map[string]any, error) {
	return c.doGET("/knowledgeGraph/data", nil)
}

func (c *ACMSClient) GetPendingKnowledge() (map[string]any, error) {
	return c.doGET("/knowledgeGraph/pendingKnowledge", nil)
}

func (c *ACMSClient) UploadKnowledgeGraph(filePath string) (map[string]any, error) {
	return c.uploadFile("/upload/knowledge-graph", "file", filePath)
}

// ── 分组 ──────────────────────────────────────────────

func (c *ACMSClient) GetGroupStudent(params map[string]any) (map[string]any, error) {
	return c.doGET("/group/student", params)
}

func (c *ACMSClient) GetGroupTeacher(params map[string]any) (map[string]any, error) {
	return c.doGET("/group/teacher", params)
}

func (c *ACMSClient) GenerateGroup(payload map[string]any) (map[string]any, error) {
	return c.doPOST("/group/generate", payload)
}

// ── 上传 ──────────────────────────────────────────────

func (c *ACMSClient) UploadImage(filePath string) (map[string]any, error) {
	return c.uploadFile("/upload/image", "file", filePath)
}

// ═══════════════════════════════════════════════════════════════
// 内部 HTTP 方法
// ═══════════════════════════════════════════════════════════════

func (c *ACMSClient) doGET(path string, params any) (map[string]any, error) {
	return c.request(http.MethodGet, path, params, nil)
}

func (c *ACMSClient) doPOST(path string, body any) (map[string]any, error) {
	return c.request(http.MethodPost, path, nil, body)
}

func (c *ACMSClient) doPUT(path string, body any) (map[string]any, error) {
	return c.request(http.MethodPut, path, nil, body)
}

func (c *ACMSClient) doDELETE(path string) (map[string]any, error) {
	return c.request(http.MethodDelete, path, nil, nil)
}

func (c *ACMSClient) doGETBytes(path string, params any) ([]byte, error) {
	fullURL := c.BaseURL + path
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}
	c.setAuth(req)
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

func (c *ACMSClient) request(method, path string, params, body any) (map[string]any, error) {
	fullURL := c.BaseURL + path
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, fullURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	c.setAuth(req)
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
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result map[string]any
	if err := json.Unmarshal(raw, &result); err != nil {
		var arrResult []any
		if err2 := json.Unmarshal(raw, &arrResult); err2 != nil {
			return map[string]any{"_raw": string(raw)}, nil
		}
		return map[string]any{"_array": arrResult}, nil
	}
	return result, nil
}

func (c *ACMSClient) uploadFile(path, fieldName, filePath string) (map[string]any, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	part, err := w.CreateFormFile(fieldName, filePath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	if _, err := io.Copy(part, f); err != nil {
		return nil, err
	}
	w.Close()

	req, err := http.NewRequest(http.MethodPost, c.BaseURL+path, &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	c.setAuth(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	var result map[string]any
	json.Unmarshal(raw, &result)
	return result, nil
}

func (c *ACMSClient) setAuth(req *http.Request) {
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
}

func toQueryMap(params any) map[string]string {
	result := make(map[string]string)
	if params == nil {
		return result
	}

	switch v := params.(type) {
	case map[string]string:
		return v
	case map[string]int:
		for k, val := range v {
			result[k] = fmt.Sprintf("%d", val)
		}
	case map[string]any:
		for k, val := range v {
			if s, ok := val.(string); ok {
				result[k] = s
			} else {
				result[k] = fmt.Sprintf("%v", val)
			}
		}
	default:
		b, _ := json.Marshal(params)
		var m map[string]any
		json.Unmarshal(b, &m)
		for k, val := range m {
			if s, ok := val.(string); ok {
				result[k] = s
			} else {
				result[k] = fmt.Sprintf("%v", val)
			}
		}
	}
	return result
}

// CreateClient 创建客户端并自动登录
func CreateClient(userID, password string) (*ACMSClient, error) {
	c := NewClient()
	data, err := c.Login(LoginForm{ID: userID, Password: password})
	if err != nil {
		return nil, err
	}
	if code, _ := data["code"].(float64); code != 200 {
		msg, _ := data["msg"].(string)
		return nil, fmt.Errorf("登录失败: %s", msg)
	}
	return c, nil
}
