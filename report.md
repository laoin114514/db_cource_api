# 《数据库原理》智慧课程 — API 函数使用说明

**文件结构：**

| 文件 | 内容 |
|------|------|
| [acms_api.py](acms_api.py) | `ACMSClient` 客户端类，91 个方法 |
| [acms_models.py](acms_models.py) | 30 个请求参数 dataclass |
| [acms_demo.py](acms_demo.py) | 综合使用示例 |

---

## 快速开始

```python
from acms_api import create_client

# 登录（学生/教师/管理员均可）
client = create_client("123456", "123456")
```

---

## 参数 Dataclass 速查表

所有参数类均从 `acms_models` 导入：

```python
from acms_models import (
    LoginForm,                              # 登录
    PageQuery,                              # 通用分页
    UserQuery, UserInsert, UserUpdate,      # 用户
    RoleForm,                               # 角色
    TeacherQuery, TeacherForm,              # 教师
    TeacherClassRelation,                   # 教师-班级关联
    ClassQuery, ClassForm,                  # 班级
    StudentQuery, StudentForm, StudentUpdate,  # 学生
    QuestionCategoryForm,                   # 题库分类
    QuestionListQuery, QuestionFind,        # 题目
    TestQuery, TestPublish,                 # 考试
    ResultStudent, ResultStart,             # 考试结果
    VideoForm,                              # 视频
    AttendanceSave, AttendanceUpdate,       # 考勤
    AttendanceStatus,                       # 考勤状态
    AIAsk,                                  # AI提问
)
```

### 各 dataclass 字段一览

| 类 | 字段 |
|----|------|
| `LoginForm` | `id: str`, `password: str` |
| `PageQuery` | `page_num: 1`, `page_size: 20`, `query: ""` |
| `UserQuery` | `id: str` |
| `UserInsert` | `id: str`, `role: int`, `name: ""`, `password: ""` |
| `UserUpdate` | `id: str`, `role: int` |
| `RoleForm` | `name: str`, `description: ""` |
| `TeacherQuery` | `teacher_id: str` |
| `TeacherForm` | `teacher_id: str`, `name: str`, `gender: ""`, `position: ""`, `password: ""` |
| `TeacherClassRelation` | `class_id: int`, `teacher_id: str` |
| `ClassQuery` | `class_name: str` |
| `ClassForm` | `class_name: str`, `headteacher_id: ""` |
| `StudentQuery` | `student_id: str` |
| `StudentForm` | `student_id: str`, `name: str`, `gender: ""`, `phone: ""`, `notes: ""`, `class_name: ""` |
| `StudentUpdate` | `student_id: str`, `name: ""`, `gender: ""`, `phone: ""`, `notes: ""`, `class_name: ""` |
| `QuestionCategoryForm` | `question_category: str`, `teacher_id: ""` |
| `QuestionListQuery` | `page_num: 1`, `page_size: 100`, `status: None`, `keyword: ""`, `category_id: None`, `type: ""`, `is_ai_generate: None`, `teacher_id: ""` |
| `QuestionFind` | `student_id: str`, `question_id: int`, `test_id: int` |
| `TestQuery` | `teacher_id: str`, `class_id: int` |
| `TestPublish` | `test_id: None`, `title: ""`, `class_ids: []`, `question_ids: []`, `duration: 120`, `start_time: ""`, `end_time: ""`, `is_test: 0` |
| `ResultStudent` | `student_id: str` |
| `ResultStart` | `student_id: str`, `test_id: int`, `is_test: 0` |
| `VideoForm` | `title: str`, `url: ""`, `description: ""`, `category_id: 0` |
| `AttendanceSave` | `class_id: int`, `student_id: ""`, `status: ""`, `date: ""`, `is_attendance: True` |
| `AttendanceUpdate` | `timestamp: ""`, `student_id: ""`, `is_attendance: True`, `login_time: ""` |
| `AttendanceStatus` | `student_id: ""`, `is_attendance: True`, `date: ""` |
| `AIAsk` | `question: str`, `context: ""` |

---

## 方法速查

### 认证

```python
client.login(LoginForm(id="123456", password="123456"))
client.logout()
client.is_logged_in()          # → bool
client.user_info               # → dict | None
```

### 访问量

```python
client.get_visit_count()       # → {"code": 200, "data": 9358}
client.increment_visit()
```

### 用户管理

```python
client.query_user(UserQuery(id="123456"))
client.get_all_users()                       # 管理员权限
client.insert_user(UserInsert(id="20240001", role=2, name="张三", password="123456"))
client.update_user(UserUpdate(id="123456", role=1))
```

### 角色管理

```python
client.get_all_roles()
client.add_role(RoleForm(name="新角色", description="描述"))
client.update_role({"id": 1, "name": "改名"})
```

### 教师管理

```python
client.get_all_teachers()                    # 完整信息
client.get_teacher_info()                    # 简要 id+name
client.query_teacher(TeacherQuery(teacher_id="20050027"))
client.get_teacher_page(PageQuery(page_num=1, page_size=20, query=""))
client.insert_teacher(TeacherForm(teacher_id="T001", name="王老师", gender="男", position="教授", password="123456"))  # 管理员
client.update_teacher(TeacherForm(teacher_id="20050027", name="蒙祖强", gender="男", position="教授", password=""))
client.teacher_join_class(TeacherClassRelation(class_id=27, teacher_id="20050027"))
client.teacher_leave_class(TeacherClassRelation(class_id=27, teacher_id="20050027"))
```

### 班级管理

```python
client.get_class_names()                     # 仅名称列表
client.get_all_classes()                     # 含学生/教师列表
client.query_class(ClassQuery(class_name="计科241班"))
client.get_class_by_name("计科241班")
client.get_class_page(PageQuery(page_num=1, page_size=20, query=""))
client.find_class_by_teacher("20050027")
client.find_class_by_teacher_and_name("20050027", "计科241班")
client.get_class_students(27)
client.insert_class(ClassForm(class_name="新班级", headteacher_id="20050027"))
client.update_class({"classId": 27, "className": "计科241班", "headteacherId": "20050027"})
```

### 学生管理

```python
client.query_student(StudentQuery(student_id="123456"))
client.get_student_page(PageQuery(page_num=1, page_size=20, query=""))
client.find_no_class_students()
client.search_student_by_prefix("24071101")
client.insert_student(StudentForm(student_id="20240001", name="张三", gender="男", class_name="计科241班"))
client.update_student(StudentUpdate(student_id="123456", name="新名字"))
client.delete_student([20240001, 20240002])         # 整数数组
client.update_student_class(["2407110101", "2407110102"], class_id=28)
client.batch_student_info({...})

# Excel 操作
data = client.export_students(27)              # → bytes
data = client.download_student_template()      # → bytes
client.batch_import_students("students.xlsx")
client.upload_student_file("students.xlsx")
```

### 题库分类

```python
client.get_question_categories()
client.add_question_category(QuestionCategoryForm(question_category="新章节", teacher_id="20050027"))
```

### 题目管理

```python
# 列表查询
client.get_question_list(QuestionListQuery(page_size=20, category_id=9, type="选择题"))
client.get_all_questions()                              # 全部 339 题
client.get_all_questions(teacher_id="20050027")         # 某教师出题
client.get_all_questions(category_id=9)                 # 某章节题
client.query_question(1168)                             # 单个题目
client.find_question(QuestionFind(student_id="123456", question_id=1168, test_id=1))

# 增删改
client.insert_question({"categoryId": 1, "type": "选择题", "questionText": "...", "answer": "A", "score": 5, "answerTime": 60, "teacherId": "20050027"})
client.update_question({"id": 1168, "questionText": "新内容", "answer": "新答案", ...})
client.batch_delete_questions([1168, 1167])

# 文件
data = client.download_question_template()              # → bytes
client.upload_question_file("questions.xlsx")
```

### 考试/测试

```python
# 查询
client.get_teacher_tests(TestQuery(teacher_id="20050027", class_id=27))  # 教师/管理员
client.get_class_test(27)
client.get_student_test("123456")

# 发布
client.publish_test(TestPublish(title="期末考试", duration=120, class_ids=[27], question_ids=[1168, 1167]))  # 教师/管理员
client.publish_test_to_class(test_id=1, class_ids=[27, 28])
client.publish_test_to_student(test_id=1, student_ids=["123456"])
client.publish_test_ai({...})
client.extend_exam_time({"studentId": "123456", "testId": 1, "extendMinutes": 10})

# 导出
data = client.export_student_report("123456", test_id=1)    # → bytes
data = client.export_class_avg_scores(test_id=1)            # → bytes
data = client.export_filtered_score_detail(test_id=1)       # → bytes
```

### 考试结果

```python
client.get_student_result(ResultStudent(student_id="123456"))
client.start_result(ResultStart(student_id="123456", test_id=1, is_test=0))
client.get_result_detail("123456", test_id=1)
client.find_result({"studentId": "123456"})
client.get_ai_result_detail("123456", test_id=1)
client.save_answer({"studentId": "123456", "testId": 1, "questionId": 1168, "answer": "我的答案"})
client.update_answer({"studentId": "123456", "testId": 1, "questionId": 1168, "answer": "修改后"})
client.update_score_comment({"studentId": "123456", "testId": 1, "questionId": 1168, "score": 8, "comment": "评语"})  # 教师/管理员
client.update_ai_result({"studentId": "123456", "testId": 1})
```

### AI 问答

```python
client.ai_student_ask(AIAsk(question="什么是数据库事务？", context="第9章"))
client.ai_teacher_ask("如何设计索引优化查询？")
```

### 视频管理

```python
client.get_video_list({"pageNum": 1, "pageSize": 20})
client.get_video_tree()
client.get_video_tree_select()
client.add_video(VideoForm(title="B+树索引", url="http://example.com/v.mp4", description="讲解", category_id=1))
client.upload_video("video.mp4")
```

### 考勤管理

```python
client.get_attendance_history(27)
client.save_attendance(AttendanceSave(class_id=27, student_id="123456", status="出勤", date="2026-05-13"))
client.update_attendance(AttendanceUpdate(timestamp="2026-05-13T09:00", student_id="123456", is_attendance=True))
client.update_attendance_status(AttendanceStatus(student_id="123456", is_attendance=True, date="2026-05-13"))
data = client.export_attendance(27)                  # → bytes
client.export_attendance_single(1)
```

### 校历

```python
client.get_semesters()
```

### 知识图谱

```python
client.get_knowledge_graph()
client.upload_knowledge_graph("kg.json")
```

### 文件上传

```python
client.upload_image("image.png")
```

### Excel 导出通用写法

```python
# 返回 bytes 的方法保存到文件
data = client.export_students(27)
with open("students.xlsx", "wb") as f:
    f.write(data)
```

---

## 响应格式

| 接口类型 | 成功响应 | 失败响应 |
|----------|---------|---------|
| 普通接口 | `{"code": 200, "data": ...}` | `{"code": 400, "msg": "..."}` |
| 权限不足 | — | `{"success": false, "error": "权限不足..."}` |
| AI 接口 | `{"success": true, "sessionId": "...", "message": "..."}` | — |
| Excel 导出 | `bytes` (直接返回二进制) | — |

---

## 权限汇总

| 接口 | 最低角色 |
|------|----------|
| `/visit/count`, `/visit/increment` | 公开 |
| `/user/all` | 管理员 |
| `/teacher/insert` | 管理员 |
| `/test/getTeacherTest`, `/test/publish` | 教师 |
| 其余全部 | 学生（已登录即可） |
