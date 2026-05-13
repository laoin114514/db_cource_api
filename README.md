# 《数据库原理》智慧课程 — API 函数使用说明

**文件：**

| 文件 | 内容 |
|------|------|
| [acms_models.py](acms_models.py) | 38 个请求参数 dataclass |
| [acms_api.py](acms_api.py) | ACMSClient 客户端，91 个方法 |
| [acms_demo.py](acms_demo.py) | 综合使用示例 |

---

## 快速开始

```python
from acms_api import create_client

client = create_client("123456", "123456")
```

---

## 参数 Dataclass 速查

| 类 | 字段 |
|----|------|
| `LoginForm` | `id: str`, `password: str` |
| `PageQuery` | `page_num: 1`, `page_size: 20`, `query: ""` |
| `UserQuery` | `id: str` |
| `UserInsert` | `id: str`, `role: int`, `name: ""`, `password: ""` |
| `UserUpdate` | `id: str`, `role: int` |
| `RoleForm` | `name: str`, `description: ""` |
| `RoleUpdate` | `id: int`, `name: ""` |
| `TeacherQuery` | `teacher_id: str` |
| `TeacherForm` | `teacher_id: str`, `name: str`, `gender: ""`, `position: ""`, `password: ""` |
| `TeacherClassRelation` | `class_id: int`, `teacher_id: str` |
| `ClassQuery` | `class_name: str` |
| `ClassForm` | `class_name: str`, `headteacher_id: ""` |
| `ClassUpdate` | `class_id: int`, `class_name: ""`, `headteacher_id: ""` |
| `StudentQuery` | `student_id: str` |
| `StudentForm` | `student_id: str`, `name: str`, `gender: ""`, `phone: ""`, `notes: ""`, `class_name: ""` |
| `StudentUpdate` | `student_id: str`, `name: ""`, `gender: ""`, `phone: ""`, `notes: ""`, `class_name: ""` |
| `BatchStudentInfo` | `students: []`, `operation: ""` |
| `QuestionCategoryForm` | `question_category: str`, `teacher_id: ""` |
| `QuestionListQuery` | `page_num: 1`, `page_size: 100`, `status: None`, `keyword: ""`, `category_id: None`, `type: ""`, `is_ai_generate: None`, `teacher_id: ""` |
| `QuestionFind` | `student_id: str`, `question_id: int`, `test_id: int` |
| `QuestionInsert` | `category_id: int`, `type: str`, `question_text: str`, `answer: str`, `score: int`, `answer_time: int`, `teacher_id: str`, `keywords: ""`, `is_ai_generate: False` |
| `QuestionUpdate` | `id: int`, `category_id: 0`, `type: ""`, `question_text: ""`, `answer: ""`, `score: 0`, `answer_time: 0` |
| `TestQuery` | `teacher_id: str`, `class_id: int` |
| `TestPublish` | `test_id: None`, `title: ""`, `class_ids: []`, `question_ids: []`, `duration: 120`, `start_time: ""`, `end_time: ""`, `is_test: 0` |
| `ExtendTestTime` | `test_id: int`, `teacher_id: str`, `minutes: 0` |
| `ResultStudent` | `student_id: str` |
| `ResultStart` | `student_id: str`, `test_id: int`, `is_test: 0` |
| `AnswerItem` | `student_id: ""`, `test_id: 0`, `question_id: 0`, `answer: ""` |
| `SaveAnswer` | `items: list[AnswerItem]` — 数组包装，序列化为 `[{...}]` |
| `ScoreItem` | `id: ""`, `question_id: 0`, `student_id: ""`, `test_id: 0`, `actual_score: 0`, `comment: ""` |
| `ScoreComment` | `items: list[ScoreItem]` — 数组包装，序列化为 `[{...}]` |
| `AIResultUpdate` | `student_id: str`, `test_id: int` |
| `VideoForm` | `title: str`, `url: ""`, `description: ""`, `category_id: 0` |
| `VideoListQuery` | `page_num: 1`, `page_size: 20`, `keyword: ""` |
| `AttendanceSave` | `class_id: int`, `student_id: ""`, `status: ""`, `date: ""`, `is_attendance: True` |
| `AttendanceUpdate` | `timestamp: ""`, `student_id: ""`, `is_attendance: True`, `login_time: ""` |
| `AttendanceStatus` | `student_id: ""`, `is_attendance: True`, `date: ""` |
| `AIAsk` | `question: str`, `context: ""` |

---

## 方法速查

### 认证 / 访问量

```python
client.login(LoginForm(id="123456", password="123456"))
client.logout()
client.is_logged_in()
client.get_visit_count()
client.increment_visit()
```

### 用户管理

```python
client.query_user(UserQuery(id="123456"))
client.get_all_users()                         # 管理员
client.insert_user(UserInsert(id="20240001", role=2, name="张三", password="123456"))
client.update_user(UserUpdate(id="123456", role=1))
```

### 角色管理

```python
client.get_all_roles()
client.add_role(RoleForm(name="新角色"))
client.update_role(RoleUpdate(id=1, name="改名"))
```

### 教师管理

```python
client.get_all_teachers()
client.get_teacher_info()
client.query_teacher(TeacherQuery(teacher_id="20050027"))
client.get_teacher_page(PageQuery(page_num=1, page_size=20, query=""))
client.insert_teacher(TeacherForm(teacher_id="T001", name="王老师", gender="男", position="教授", password="123456"))
client.update_teacher(TeacherForm(teacher_id="20050027", name="蒙祖强", gender="男", position="教授", password=""))
client.teacher_join_class(TeacherClassRelation(class_id=27, teacher_id="20050027"))
client.teacher_leave_class(TeacherClassRelation(class_id=27, teacher_id="20050027"))
```

### 班级管理

```python
client.get_class_names()
client.get_all_classes()
client.query_class(ClassQuery(class_name="计科241班"))
client.get_class_by_name("计科241班")
client.get_class_page(PageQuery(page_num=1, page_size=20, query=""))
client.find_class_by_teacher("20050027")
client.find_class_by_teacher_and_name("20050027", "计科241班")
client.get_class_students(27)
client.insert_class(ClassForm(class_name="新班级", headteacher_id="20050027"))
client.update_class(ClassUpdate(class_id=27, class_name="计科241班", headteacher_id="20050027"))
```

### 学生管理

```python
client.query_student(StudentQuery(student_id="123456"))
client.get_student_page(PageQuery(page_num=1, page_size=20, query=""))
client.find_no_class_students()
client.search_student_by_prefix("24071101")
client.insert_student(StudentForm(student_id="20240001", name="张三", gender="男", class_name="计科241班"))
client.update_student(StudentUpdate(student_id="123456", name="新名字"))
client.delete_student([20240001, 20240002])
client.update_student_class(["2407110101"], class_id=28)
client.batch_student_info(BatchStudentInfo(students=[]))

data = client.export_students(27)
data = client.download_student_template()
client.batch_import_students("students.xlsx")
client.upload_student_file("students.xlsx")
```

### 题库分类

```python
client.get_question_categories()
client.add_question_category(QuestionCategoryForm(question_category="新章节"))
```

### 题目管理

```python
client.get_question_list(QuestionListQuery(page_size=20, category_id=9, type="选择题"))
client.get_all_questions()
client.get_all_questions(teacher_id="20050027")
client.get_all_questions(category_id=9)
client.query_question(1168)
client.find_question(QuestionFind(student_id="123456", question_id=1168, test_id=1))
client.insert_question(QuestionInsert(category_id=1, type="选择题", question_text="...", answer="A", score=5, answer_time=60, teacher_id="20050027"))
client.update_question(QuestionUpdate(id=1168, question_text="新内容"))
client.batch_delete_questions([1168, 1167])
data = client.download_question_template()
client.upload_question_file("questions.xlsx")
```

### 考试/测试

```python
client.get_teacher_tests(TestQuery(teacher_id="20050027", class_id=27))  # 教师/管理员
client.get_class_test(27)
client.get_student_test("123456")
client.publish_test(TestPublish(title="期末考试", duration=120, class_ids=[27], question_ids=[1168, 1167]))
client.publish_test_ai(TestPublish(title="AI测试", duration=60))
client.publish_test_to_class(test_id=1, class_ids=[27, 28])
client.publish_test_to_student(test_id=1, student_ids=["123456"])
client.extend_test_time(ExtendTestTime(test_id=100, teacher_id="123456", minutes=60))

data = client.export_student_report("123456", test_id=1)
data = client.export_class_avg_scores(test_id=1)
data = client.export_filtered_score_detail(test_id=1)
```

### 考试结果

```python
client.get_student_result(ResultStudent(student_id="123456"))
client.start_result(ResultStart(student_id="123456", test_id=1, is_test=0))
client.get_result_detail("123456", test_id=1)
client.get_ai_result_detail("123456", test_id=1)

# 保存作答 — PUT 请求，发送 AnswerItem 数组
client.save_answer(SaveAnswer(items=[
    AnswerItem(student_id="123456", test_id=1, question_id=839, answer="C"),
]))

# 更新作答 — PUT 请求
client.update_answer(SaveAnswer(items=[
    AnswerItem(student_id="123456", test_id=1, question_id=839, answer="修改后"),
]))

# 更新分数评语 — POST 请求，发送 ScoreItem 数组
client.update_score_comment(ScoreComment(items=[
    ScoreItem(id="2407110107839100", question_id=839, student_id="123456",
              test_id=100, actual_score=30, comment="答案正确"),
]))

client.update_ai_result(AIResultUpdate(student_id="123456", test_id=1))
```

### AI 问答

```python
client.ai_student_ask(AIAsk(question="什么是数据库事务？", context="第9章"))
client.ai_teacher_ask("如何设计索引优化查询？")
```

### 视频管理

```python
client.get_video_list(VideoListQuery(page_num=1, page_size=20, keyword=""))
client.get_video_tree()
client.get_video_tree_select()
client.add_video(VideoForm(title="B+树索引", url="http://example.com/v.mp4", description="讲解"))
client.upload_video("video.mp4")
```

### 考勤管理

```python
client.get_attendance_history(27)
client.save_attendance(AttendanceSave(class_id=27, student_id="123456", status="出勤", date="2026-05-13"))
client.update_attendance(AttendanceUpdate(timestamp="2026-05-13T09:00", student_id="123456"))
client.update_attendance_status(AttendanceStatus(student_id="123456", is_attendance=True, date="2026-05-13"))
data = client.export_attendance(27)
client.export_attendance_single(1)
```

### 校历 / 知识图谱 / 上传

```python
client.get_semesters()
client.get_knowledge_graph()
client.upload_knowledge_graph("kg.json")
client.upload_image("image.png")
```

### Excel 导出通用写法

```python
data = client.export_students(27)
with open("students.xlsx", "wb") as f:
    f.write(data)
```

---

## HTTP 方法对照

| 接口 | 方法 | 请求体格式 |
|------|------|-----------|
| saveAnswer | **PUT** | `[{studentId, testId, questionId, answer}, ...]` |
| updateAnswer | **PUT** | 同上 |
| updateScoreAndComment | POST | `[{id, studentId, testId, questionId, actualScore, comment}, ...]` |
| 其余全部 | POST/GET | — |

---

## 响应格式

| 接口类型 | 成功 | 失败 |
|----------|------|------|
| 普通接口 | `{"code": 200, "data": ...}` | `{"code": 400, "msg": "..."}` |
| 权限不足 | — | `{"success": false, "error": "权限不足..."}` |
| AI 接口 | `{"success": true, ...}` | — |
| Excel 导出 | `bytes` | — |

## 权限汇总

| 接口 | 最低角色 |
|------|----------|
| `/visit/count`, `/visit/increment` | 公开 |
| `/user/all` | 管理员 |
| `/teacher/insert` | 管理员 |
| `/test/getTeacherTest`, `/test/publish` | 教师 |
| 其余全部 | 学生（已登录即可） |
