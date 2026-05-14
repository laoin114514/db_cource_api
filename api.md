# 《数据库原理》智慧课程 API 接口文档

**Base URL:** `http://172.21.44.162:5174/acms`

**认证方式:** Bearer Token (JWT)，`Authorization: Bearer <token>`

**方法总数:** GET 46 / POST 31 / PUT 12 / DELETE 4

---

## 目录

1. [认证](#1-认证)
2. [访问量](#2-访问量)
3. [用户管理](#3-用户管理)
4. [角色管理](#4-角色管理)
5. [教师管理](#5-教师管理)
6. [班级管理](#6-班级管理)
7. [学生管理](#7-学生管理)
8. [题库分类](#8-题库分类)
9. [题目管理](#9-题目管理)
10. [考试/测试](#10-考试测试)
11. [考试结果](#11-考试结果)
12. [AI 问答](#12-ai-问答)
13. [视频管理](#13-视频管理)
14. [考勤管理](#14-考勤管理)
15. [校历](#15-校历)
16. [知识图谱](#16-知识图谱)
17. [分组管理](#17-分组管理)
18. [文件上传](#18-文件上传)
19. [HTTP 方法对照](#19-http-方法对照)
20. [通用说明](#20-通用说明)

---

## 1. 认证

### POST /auth/login — 登录

```json
// Request
{"id": "123456", "password": "123456"}

// Response 200
{"code": 200, "data": {"user": {"role": 2, "name": "张三", "position": "学生"}, "token": "eyJ..."}}
```

| role | position |
|------|----------|
| 0 | 管理员 |
| 1 | 教师 |
| 2 | 学生 |

---

## 2. 访问量

### GET /visit/count — 获取访问量（公开）

### POST /visit/increment — 增加访问量（公开）

---

## 3. 用户管理

| 方法 | 路径 | 说明 | 权限 |
|------|------|------|------|
| GET | `/user/query?id=` | 按 ID 查询 | 已登录 |
| GET | `/user/all` | 获取所有用户 | 管理员 |
| POST | `/user/insert` | 添加用户 | 已登录 |
| POST | `/user/update` | 更新用户 | 已登录 |
| DELETE | `/user/{id}` | 删除用户 | 管理员 |

### POST /user/insert — `{"id":"学号","role":2,"name":"姓名","password":"123456"}`
### POST /user/update — `{"id":"学号","role":1}`
### DELETE /user/{id} — 路径参数，如 `DELETE /user/88888`

---

## 4. 角色管理

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/role/all` | 获取所有角色 |
| POST | `/role/add` | 添加角色 |
| PUT | `/role/update` | 更新角色 |

### POST /role/add — `{"name":"角色名","description":"描述"}`
### PUT /role/update — `{"id":1,"name":"新名称"}`

---

## 5. 教师管理

| 方法 | 路径 | 说明 | 权限 |
|------|------|------|------|
| GET | `/teacher/findAll` | 获取所有教师 | 已登录 |
| GET | `/teacher/acquireInfo` | 教师简要列表 | 已登录 |
| GET | `/teacher/query?teacherId=` | 按工号查询 | 已登录 |
| GET | `/teacher/showPage?pageNum=&pageSize=&query=` | 分页查询 | 已登录 |
| POST | `/teacher/insert` | 添加教师 | 管理员 |
| POST | `/teacher/update` | 更新教师 | 已登录 |
| POST | `/teacher/join` | 加入班级 | 已登录 |
| POST | `/teacher/remove` | 移出班级 | 已登录 |
| POST | `/teacher/batchInfo` | 批量信息操作 | 已登录 |
| DELETE | `/teacher/{teacherId}` | 删除教师 | 管理员 |

### POST /teacher/insert — `{"teacherId":"工号","name":"姓名","gender":"男","position":"职称","password":"密码"}`
### POST /teacher/update — `{"teacher":{"teacherId":"","name":"","gender":"","position":""},"password":""}`
### POST /teacher/join — `{"classId":27,"teacherId":"20050027"}`
### POST /teacher/remove — `{"classId":27,"teacherId":"20050027"}`
### DELETE /teacher/{teacherId} — 路径参数

---

## 6. 班级管理

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/class/acquireName` | 班级名称列表 |
| GET | `/class/findAll` | 所有班级（含学生/教师） |
| GET | `/class/getByName?className=` | 按名称获取 |
| GET | `/class/queryClass?className=` | 按名称查询 |
| GET | `/class/showPage?pageNum=&pageSize=&query=` | 分页查询 |
| GET | `/class/findByTeacher?teacherId=` | 按教师查班级 |
| GET | `/class/findByTeacherAndName?teacherId=&className=` | 按教师+名称查 |
| GET | `/class/findByStudent?studentId=` | 按学生查班级 |
| GET | `/class/students?classId=` | 班级学生列表 |
| POST | `/class/insert` | 添加班级 |
| PUT | `/class/update` | 更新班级 |
| DELETE | `/class/{classId}` | 删除班级 |

### POST /class/insert — `{"className":"名称","headteacherId":"工号"}`
### PUT /class/update — `{"classId":27,"className":"名称","headteacherId":"工号"}`
### DELETE /class/{classId} — 路径参数

---

## 7. 学生管理

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/student/query?studentId=` | 按学号查询 |
| GET | `/student/showPage?pageNum=&pageSize=&query=` | 分页查询 |
| GET | `/student/findNoClass` | 未分班学生 |
| GET | `/student/searchByStudentIdPrefix?prefix=` | 学号前缀搜索 |
| GET | `/student/export?classId=` | 导出 Excel |
| POST | `/student/insert` | 添加学生 |
| PUT | `/student/update` | 更新学生 |
| POST | `/student/delete` | 批量删除（POST 数组） |
| DELETE | `/student/{studentId}` | 删除单个学生 |
| PUT | `/student/updateClass` | 更新班级 |
| POST | `/student/batchInfo` | 批量信息操作 |
| POST | `/student/batchImport` | 批量导入 Excel |
| POST | `/studentFile/upload` | 上传学生文件 |

### POST /student/insert — `{"studentId":"学号","name":"姓名","gender":"男","phone":"","notes":"","className":"班级"}`
### PUT /student/update — `{"student":{"studentId":"学号","name":"","className":"计科241班"}}`
> className 必传，否则 400
### POST /student/delete — `[2407110101, 2407110102]` （整数数组）
### DELETE /student/{studentId} — 路径参数
### PUT /student/updateClass — `{"studentId":"学号","className":"班级","classId":27}`
> classId=null 表示移出班级

---

## 8. 题库分类

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/question-category/list` | 获取全部分类 |
| POST | `/question-category/add` | 添加分类 |

12 个分类：第1章 数据库概述 ~ 第12章 数据库备份与恢复

### POST /question-category/add — `{"questionCategory":"名称","teacherId":"工号"}`

---

## 9. 题目管理

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/question/list` | 分页查询 |
| GET | `/question/query?id=` | 查询单个 |
| GET | `/question/find?studentId=&questionId=&testId=` | 考试上下文查找 |
| POST | `/question/insert` | 添加题目 |
| PUT | `/question/update` | 更新题目 |
| POST | `/question/batchDelete` | 批量删除 |
| GET | `/questionFile/template/download` | 下载模板 |
| POST | `/questionFile/upload` | 上传题目文件 |

### GET /question/list

| 参数 | 类型 | 说明 |
|------|------|------|
| pageNum | int | 页码 |
| pageSize | int | 每页条数 |
| categoryId | int | 章节 1-12 |
| type | string | 题型 |
| keyword | string | 搜索 |
| isAiGenerate | int | 0=非AI, 1=AI |
| teacherId | string | 出题教师 |

可选题型：选择题/填空题/简答题/实验题/设计题/判断题/证明题

题库统计：总计 339 题（选择 104 / 填空 95 / 简答 94 / 实验 24 / 设计 12 / 判断 5 / 证明 5）

### POST /question/insert — `{"categoryId":1,"type":"选择题","questionText":"...","answer":"A","score":5,"answerTime":60,"teacherId":"20050027"}`
### PUT /question/update — `{"id":1168,"categoryId":1,"type":"简答题","questionText":"...","answer":"...","score":10,"answerTime":120}`
### POST /question/batchDelete — `[1168, 1167, 1166]`

---

## 10. 考试/测试

| 方法 | 路径 | 说明 | 权限 |
|------|------|------|------|
| GET | `/test/getTeacherTest?teacherId=&classId=` | 教师测试列表 | 教师/管理员 |
| GET | `/test/getClassTest?classId=` | 班级测试 | 已登录 |
| GET | `/test/getStudentTest?studentId=` | 学生测试 | 已登录 |
| GET | `/test/publish/teacher?teacherId=` | 教师已发布测试 | 教师/管理员 |
| GET | `/test/publish/ai` | AI 发布测试 | 教师/管理员 |
| GET | `/test/publish/class` | 班级发布测试 | 教师/管理员 |
| GET | `/test/publish/student` | 学生发布测试 | 教师/管理员 |
| POST | `/test/publish` | 发布考试 | 教师/管理员 |
| POST | `/test/extendTestTime` | 延长时间 | 教师/管理员 |
| GET | `/test/exportStudentReport/excel` | 导出学生报告 | 已登录 |
| GET | `/test/exportClassAverageScores/excel` | 导出班级均分 | 已登录 |
| GET | `/test/exportFilteredScoreDetail/excel` | 导出成绩详情 | 已登录 |

### POST /test/publish — `{"testId":null,"title":"期末","classIds":[27],"questionIds":[1168],"duration":120,"startTime":"","endTime":"","isTest":0}`
### POST /test/extendTestTime — `{"testId":100,"teacherId":"123456","minutes":60}`

---

## 11. 考试结果

| 方法 | 路径 | 说明 | 权限 |
|------|------|------|------|
| GET | `/result/student?studentId=` | 学生成绩 | 已登录 |
| GET | `/result/start?studentId=&testId=&isTest=` | 开始考试 | 已登录 |
| GET | `/result/detail?studentId=&testId=` | 结果详情 | 已登录 |
| GET | `/result/find` | 查找结果 | 已登录 |
| GET | `/result/ai/detail?studentId=&testId=` | AI 评分详情 | 已登录 |
| PUT | `/result/saveAnswer` | 保存作答 | 已登录 |
| PUT | `/result/updateAnswer` | 提交作答 | 已登录 |
| POST | `/result/updateScoreAndComment` | 评分 | 教师/管理员 |
| POST | `/result/update/ai` | 更新 AI 结果 | 已登录 |

### PUT /result/saveAnswer — `[{studentId, testId, questionId, answer}, ...]`
### PUT /result/updateAnswer — 同上格式，提交后触发 AI 批改
### POST /result/updateScoreAndComment — `[{id, studentId, testId, questionId, actualScore, comment}, ...]`

---

## 12. AI 问答

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/ai/student/ask` | AI 批改（questions 数组） |
| POST | `/ai/teacher/ask` | 教师 AI 提问 |

### POST /ai/student/ask（批改格式）

```json
{"questions": [{"testId":142,"studentId":"123456","questionId":1040,"index":1,"questionText":"...","studentAnswer":"<p>...</p>","score":0,"maxScore":28}]}
```

### POST /ai/teacher/ask — `{"question":"如何设计索引？","context":""}`

---

## 13. 视频管理

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/video/tree` | 分类树 |
| GET | `/video/tree/select` | 分类树（选择器） |
| POST | `/video/add` | 添加视频 |
| POST | `/video/upload` | 上传视频文件 |

### POST /video/add — `{"title":"标题","url":"地址","description":"描述","categoryId":0}`

---

## 14. 考勤管理

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/attendance/history?classId=` | 班级考勤历史 |
| POST | `/attendance/save` | 保存考勤 |
| POST | `/attendance/update` | 更新记录 |
| POST | `/attendance/updateStatus` | 更新状态 |
| GET | `/attendance/export?classId=` | 导出 Excel |
| GET | `/attendance/exportSingle?id=` | 导出单条 |

---

## 15. 校历

### GET /academicCalendar/semesters

---

## 16. 知识图谱

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/knowledgeGraph/data` | 图谱数据 |
| GET | `/knowledgeGraph/pendingKnowledge` | 待处理图谱 |
| POST | `/upload/knowledge-graph` | 上传文件 |

---

## 17. 分组管理

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/group/student` | 学生分组 |
| GET | `/group/teacher` | 教师分组 |
| POST | `/group/generate` | 生成分组 |

---

## 18. 文件上传

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/upload/image` | 上传图片 |

---

## 19. HTTP 方法对照

| 方法 | 接口 |
|------|------|
| **PUT** | `/student/update`, `/student/updateClass`, `/class/update`, `/question/update`, `/role/update`, `/result/saveAnswer`, `/result/updateAnswer` |
| **DELETE** | `/user/{id}`, `/teacher/{id}`, `/student/{id}`, `/class/{id}` |
| 其余 | POST / GET |

---

## 20. 通用说明

### 响应格式

**成功:** `{"code": 200, "msg": "成功", "data": ...}`
**失败:** `{"code": 400, "msg": "错误描述", "data": null}`
**权限不足:** `{"success": false, "error": "权限不足，需要以下角色之一: ..."}`
**AI 接口:** `{"success": true, "sessionId": "...", "message": "..."}`

### 权限汇总

| 接口 | 最低角色 |
|------|----------|
| `/visit/count`, `/visit/increment` | 公开 |
| `/user/all`, `/teacher/insert`, `/user/{id}`, `/teacher/{id}` | 管理员 |
| `/test/getTeacherTest`, `/test/publish`, `/test/publish/*`, `/test/extendTestTime`, `/result/updateScoreAndComment` | 教师 |
| 其余 | 学生（已登录） |

### Python 客户端

[acms_api.py](acms_api.py) — 102 个方法 | [acms_models.py](acms_models.py) — 38 个 dataclass

```python
from acms_api import create_client
from acms_models import *
client = create_client("123456", "123456")
```
