# 《数据库原理》智慧课程 API 接口文档

**Base URL:** `http://172.21.44.162:5174/acms`

**认证方式:** 登录后获取 JWT Token，在请求头携带 `Authorization: Bearer <token>`

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
17. [文件上传](#17-文件上传)
18. [通用说明](#18-通用说明)

---

## 1. 认证

### POST /auth/login

> 登录获取 token

**Request**

```json
{
  "id": "123456",
  "password": "123456"
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | string | 是 | 学号或工号 |
| password | string | 是 | 密码 |

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功",
  "total": 0,
  "data": {
    "user": {
      "role": 2,
      "gender": "null",
      "name": "李四",
      "id": "123456",
      "position": "学生"
    },
    "token": "eyJ0eXAiOiJKV1Q..."
  }
}
```

| role | position | 说明 |
|------|----------|------|
| 0 | — | 管理员 |
| 1 | — | 教师 |
| 2 | 学生 | 学生 |

---

## 2. 访问量

### GET /visit/count

> 获取网站访问量（无需登录）

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功",
  "total": 0,
  "data": 9358
}
```

---

### POST /visit/increment

> 增加访问量（无需登录）

**Request**

无参数。

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

## 3. 用户管理

### GET /user/query

> 按 ID 查询用户

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | string | 是 | 学号/工号 |

**Response** `200 OK`

```json
{
  "code": 200,
  "data": {
    "role": 2,
    "gender": "null",
    "name": "李四",
    "id": "123456",
    "position": "学生"
  }
}
```

---

### GET /user/all

> 获取所有用户

**权限:** 管理员

**Response** `200 OK`

```json
{
  "code": 200,
  "data": [
    {"role": 0, "name": "管理员", "id": "114514", "position": null},
    {"role": 1, "name": "王六", "id": "20050027", "position": "教授"},
    {"role": 2, "name": "李四", "id": "123456", "position": "学生"}
  ]
}
```

**错误**

```json
{
  "success": false,
  "error": "权限不足，需要以下角色之一: 管理员"
}
```

---

### POST /user/insert

> 添加用户

**Request**

```json
{
  "id": "20240001",
  "role": 2,
  "name": "张三",
  "password": "123456"
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | string | 是 | 学号/工号 |
| role | int | 是 | 0=管理员 1=教师 2=学生 |
| name | string | 是 | 姓名 |
| password | string | 是 | 密码 |

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功",
  "data": 1
}
```

---

### POST /user/update

> 更新用户角色

**Request**

```json
{
  "id": "123456",
  "role": 1
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | string | 是 | 学号/工号 |
| role | int | 是 | 目标角色 |

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功",
  "data": 1
}
```

---

## 4. 角色管理

### GET /role/all

> 获取所有角色

**Response** `200 OK`

```json
{
  "code": 200,
  "data": []
}
```

---

### POST /role/add

> 添加角色

**Request**

```json
{
  "name": "角色名",
  "description": "描述"
}
```

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功",
  "data": 1
}
```

---

### POST /role/update

> 更新角色

**Request**

```json
{
  "id": 1,
  "name": "新名称"
}
```

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

## 5. 教师管理

### GET /teacher/findAll

> 获取所有教师

**Response** `200 OK`

```json
{
  "code": 200,
  "data": [
    {
      "teacherId": "20050027",
      "name": "王六",
      "gender": "男",
      "phone": null,
      "position": "教授",
      "listOfClasses": "[\"计科241班\"]"
    },
    {
      "teacherId": "19900025",
      "name": "梁正友",
      "gender": "男",
      "phone": null,
      "position": "教授",
      "listOfClasses": "[\"计科242班\"]"
    },
    {
      "teacherId": "20070059",
      "name": "孙宇",
      "gender": "男",
      "phone": null,
      "position": "教授",
      "listOfClasses": "[\"计科243班\"]"
    }
  ]
}
```

---

### GET /teacher/acquireInfo

> 获取教师简要列表（id + name）

**Response** `200 OK`

```json
{
  "code": 200,
  "data": [
    {"teacher_id": "20050027", "name": "王六"},
    {"teacher_id": "19900025", "name": "梁正友"},
    {"teacher_id": "20070059", "name": "孙宇"}
  ]
}
```

---

### GET /teacher/query

> 按工号查询教师详情

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| teacherId | string | 是 | 教师工号 |

**Response** `200 OK`

```json
{
  "code": 200,
  "data": {
    "teacherId": "20050027",
    "name": "王六",
    "gender": "男",
    "position": "教授"
  }
}
```

---

### GET /teacher/showPage

> 分页查询教师

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| pageNum | int | 是 | 页码 |
| pageSize | int | 是 | 每页条数 |
| query | string | 是 | 搜索关键词，空字符串表示全部 |

---

### POST /teacher/insert

> 添加教师

**权限:** 管理员

**Request**

```json
{
  "teacherId": "20250001",
  "name": "王老师",
  "gender": "男",
  "position": "讲师",
  "password": "123456"
}
```

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

### POST /teacher/update

> 更新教师信息

**Request**

```json
{
  "teacher": {
    "teacherId": "20050027",
    "name": "王六",
    "gender": "男",
    "position": "教授"
  },
  "password": ""
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| teacher.teacherId | string | 是 | 工号 |
| teacher.name | string | 是 | 姓名 |
| teacher.gender | string | 否 | 男/女 |
| teacher.position | string | 否 | 职称 |
| password | string | 否 | 空字符串表示不修改密码 |

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

### POST /teacher/join

> 教师加入班级

**Request**

```json
{
  "classId": 27,
  "teacherId": "20050027"
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| classId | int | 是 | 班级 ID |
| teacherId | string | 是 | 教师工号 |

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

### POST /teacher/remove

> 教师移出班级

**Request**

```json
{
  "classId": 27,
  "teacherId": "20050027"
}
```

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

## 6. 班级管理

### GET /class/acquireName

> 获取所有班级名称

**Response** `200 OK`

```json
{
  "code": 200,
  "data": ["计科241班", "计科242班", "计科243班", "测试班级"]
}
```

---

### GET /class/findAll

> 获取所有班级（含学生/教师列表）

**Response** `200 OK`

```json
{
  "code": 200,
  "data": [
    {
      "classId": 27,
      "className": "计科241班",
      "listOfStudents": "[\"2407110101\",\"2407110102\",...]",
      "listOfTeachers": "[\"20050027\"]",
      "headteacherId": "20050027"
    }
  ]
}
```

> `listOfStudents` 和 `listOfTeachers` 为 JSON 字符串，需 `JSON.parse` 解析

---

### GET /class/getByName

> 按名称获取班级

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| className | string | 是 | 班级名称 |

---

### GET /class/queryClass

> 按名称查询班级

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| className | string | 是 | 班级名称 |

---

### GET /class/showPage

> 分页查询班级

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| pageNum | int | 是 | 页码 |
| pageSize | int | 是 | 每页条数 |
| query | string | 是 | 搜索关键词，空字符串表示全部 |

---

### GET /class/findByTeacher

> 按教师工号查找其负责的班级

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| teacherId | string | 是 | 教师工号 |

**Response** `200 OK`

```json
{
  "code": 200,
  "data": [
    {
      "classId": 27,
      "className": "计科241班",
      "listOfStudents": "[\"2407110101\",...]",
      "listOfTeachers": "[\"20050027\"]",
      "headteacherId": "20050027"
    }
  ]
}
```

---

### GET /class/findByTeacherAndName

> 按教师和班级名称查找

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| teacherId | string | 是 | 教师工号 |
| className | string | 是 | 班级名称 |

---

### GET /class/students

> 获取班级学生列表

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| classId | int | 是 | 班级 ID |

**Response** `200 OK`

```json
{
  "code": 200,
  "data": [
    {"studentId": "2407110101", "name": "...", "gender": "...", ...}
  ]
}
```

---

### POST /class/insert

> 添加班级

**Request**

```json
{
  "className": "新班级",
  "headteacherId": "20050027"
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| className | string | 是 | 班级名称 |
| headteacherId | string | 是 | 班主任工号 |

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功",
  "data": "班级添加成功"
}
```

---

### POST /class/update

> 更新班级信息

**Request**

```json
{
  "classId": 27,
  "className": "计科241班",
  "headteacherId": "20050027"
}
```

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

## 7. 学生管理

### GET /student/query

> 按学号查询学生

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |

---

### GET /student/showPage

> 分页查询学生

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| pageNum | int | 是 | 页码 |
| pageSize | int | 是 | 每页条数 |
| query | string | 是 | 搜索关键词，空字符串表示全部 |

---

### GET /student/findNoClass

> 查找所有未分班的学生

**Response** `200 OK`

```json
{
  "code": 200,
  "data": [
    {"studentId": "20240001", "name": "张三", ...}
  ]
}
```

---

### GET /student/searchByStudentIdPrefix

> 按学号前缀搜索学生

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| prefix | string | 是 | 学号前缀 |

**Response** `200 OK`

```json
{
  "code": 200,
  "data": [
    {"studentId": "123456", "name": "李四", ...}
  ]
}
```

---

### GET /student/export

> 导出学生 Excel

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| classId | int | 是 | 班级 ID |

**Response** `200 OK`

```
Content-Type: application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
```

---

### POST /student/insert

> 添加学生

**Request**

```json
{
  "studentId": "20240001",
  "name": "张三",
  "gender": "男",
  "phone": "13800138000",
  "notes": "备注信息",
  "className": "计科241班"
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |
| name | string | 是 | 姓名 |
| gender | string | 否 | 男/女 |
| phone | string | 否 | 手机号 |
| notes | string | 否 | 备注 |
| className | string | 否 | 班级名称 |

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

### POST /student/update

> 更新学生信息

**Request**

```json
{
  "student": {
    "studentId": "123456",
    "name": "李四",
    "gender": "男",
    "phone": "13800000000",
    "notes": "",
    "className": "计科241班"
  }
}
```

> 注意：字段必须包裹在 `student` 对象内

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| student.studentId | string | 是 | 学号 |
| student.name | string | 否 | 姓名 |
| student.gender | string | 否 | 男/女 |
| student.phone | string | 否 | 手机号 |
| student.notes | string | 否 | 备注 |
| student.className | string | 否 | 班级名称 |

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

### POST /student/delete

> 批量删除学生

**Request**

```json
[2407110101, 2407110102, 2407110103]
```

> 学号的整数数组（非字符串）

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

### POST /student/updateClass

> 批量更新学生班级

**Request**

```json
{
  "studentIds": ["2407110101", "2407110102"],
  "classId": 28
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentIds | string[] | 是 | 学号数组 |
| classId | int | 是 | 目标班级 ID |

---

### POST /student/batchInfo

> 批量学生信息操作

**Request**

```json
{}
```

---

### POST /student/batchImport

> 批量导入学生 Excel

**Request**

```
Content-Type: multipart/form-data
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| file | file | 是 | Excel 文件 |

---

### POST /studentFile/upload

> 上传学生文件

**Request**

```
Content-Type: multipart/form-data
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| file | file | 是 | Excel 文件 |

---

## 8. 题库分类

### GET /question-category/list

> 获取全部题库分类

**Response** `200 OK`

```json
{
  "code": 200,
  "data": [
    {"id": 1,  "questionCategory": "第1章 数据库概述"},
    {"id": 2,  "questionCategory": "第2章 关系数据库理论基础"},
    {"id": 3,  "questionCategory": "第3章 数据库设计技术"},
    {"id": 4,  "questionCategory": "第4章 数据库查询语言SQL"},
    {"id": 5,  "questionCategory": "第5章 Transact-SQL程序设计"},
    {"id": 6,  "questionCategory": "第6章 数据库的创建和管理"},
    {"id": 7,  "questionCategory": "第7章 索引与视图"},
    {"id": 8,  "questionCategory": "第8章 存储过程和触发器"},
    {"id": 9,  "questionCategory": "第9章 事务管理与并发控制"},
    {"id": 10, "questionCategory": "第10章 数据的完整性管理"},
    {"id": 11, "questionCategory": "第11章 数据的安全性控制"},
    {"id": 12, "questionCategory": "第12章 数据库备份与恢复"}
  ]
}
```

---

### POST /question-category/add

> 添加题库分类

**Request**

```json
{
  "questionCategory": "新章节名称",
  "teacherId": "20050027"
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| questionCategory | string | 是 | 分类名称 |
| teacherId | string | 否 | 教师工号 |

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

## 9. 题目管理

### GET /question/list

> 分页查询题目

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| pageNum | int | 是 | 页码 |
| pageSize | int | 是 | 每页条数（最大 500） |
| categoryId | int | 否 | 章节分类 ID（1-12） |
| type | string | 否 | 题型 |
| keyword | string | 否 | 关键词搜索 |
| isAiGenerate | int | 否 | 0=非AI, 1=AI生成（不传=全部） |
| status | int | 否 | 状态过滤（不传=全部） |
| teacherId | string | 否 | 出题教师工号 |

**可选题型:** `选择题` `填空题` `简答题` `实验题` `设计题` `判断题` `证明题`

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功",
  "data": {
    "total": 339,
    "pageSize": 500,
    "items": [
      {
        "id": 1168,
        "categoryId": 1,
        "type": "简答题",
        "answerTime": 120,
        "questionText": "简述"不可重复读"与"读脏数据"的主要区别...",
        "answer": "区别：①读脏数据是指事务读取了另一个未提交事务修改后的数据...",
        "keywords": "[\"不可重复读\",\"读脏数据\",\"事务\"]",
        "score": 10,
        "teacherId": "20050027",
        "isAiGenerate": true
      }
    ]
  }
}
```

**题库统计**

| 章节 | 数量 | | 题型 | 数量 |
|------|------|-|------|------|
| 第1章 数据库概述 | 83 | | 选择题 | 104 |
| 第2章 关系数据库理论基础 | 53 | | 填空题 | 95 |
| 第3章 数据库设计技术 | 12 | | 简答题 | 94 |
| 第4章 SQL | 32 | | 实验题 | 24 |
| 第5章 Transact-SQL | 17 | | 设计题 | 12 |
| 第6章 数据库创建和管理 | 21 | | 判断题 | 5 |
| 第7章 索引与视图 | 29 | | 证明题 | 5 |
| 第8章 存储过程和触发器 | 22 | | | |
| 第9章 事务管理与并发控制 | 18 | | | |
| 第10章 数据完整性 | 20 | | | |
| 第11章 数据安全性 | 19 | | | |
| 第12章 备份与恢复 | 13 | | | |
| **总计** | **339** | | **总计** | **339** |

---

### GET /question/query

> 查询单个题目详情

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | int | 是 | 题目 ID |

---

### GET /question/find

> 查找题目（考试上下文）

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |
| questionId | int | 是 | 题目 ID |
| testId | int | 是 | 考试 ID |

> 需要进行中的考试，否则返回 400

**Response** `200 OK`

```json
{
  "code": 200,
  "data": {
    "id": 1168,
    "type": "简答题",
    "questionText": "...",
    "answerTime": 120,
    "score": 10,
    "teacherId": "20050027"
  }
}
```

---

### POST /question/insert

> 添加题目

**Request**

```json
{
  "categoryId": 1,
  "type": "选择题",
  "questionText": "题目内容",
  "answer": "A",
  "score": 5,
  "answerTime": 60,
  "teacherId": "20050027"
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| categoryId | int | 是 | 章节分类 ID |
| type | string | 是 | 题型 |
| questionText | string | 是 | 题目内容 |
| answer | string | 是 | 参考答案 |
| score | int | 是 | 分值 |
| answerTime | int | 是 | 答题时间（秒） |
| teacherId | string | 是 | 出题教师工号 |
| keywords | string | 否 | 关键词 JSON 数组字符串 |
| isAiGenerate | bool | 否 | 是否 AI 生成 |

也支持 `multipart/form-data` 带文件上传。

---

### POST /question/update

> 更新题目

**Request**

```json
{
  "id": 1168,
  "categoryId": 1,
  "type": "简答题",
  "questionText": "更新后的题目内容",
  "answer": "新答案",
  "score": 10,
  "answerTime": 120
}
```

---

### POST /question/batchDelete

> 批量删除题目

**Request**

```json
[1168, 1167, 1166]
```

> 题目 ID 的整数数组

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

### GET /questionFile/template/download

> 下载题目导入模板

**Response** `200 OK`

```
Content-Type: application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
```

---

### POST /questionFile/upload

> 上传题目 Excel 文件

**Request**

```
Content-Type: multipart/form-data
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| file | file | 是 | Excel 文件 |

---

## 10. 考试/测试

### GET /test/getTeacherTest

> 获取教师发布的测试

**权限:** 教师或管理员

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| teacherId | string | 是 | 教师工号 |
| classId | int | 是 | 班级 ID |

---

### GET /test/getClassTest

> 获取班级的测试列表

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| classId | int | 是 | 班级 ID |

---

### GET /test/getStudentTest

> 获取学生的测试列表

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |

---

### POST /test/publish

> 发布考试/测试

**权限:** 教师或管理员

**Request**

```json
{
  "testId": null,
  "title": "期末考试",
  "classIds": [27, 28],
  "questionIds": [1168, 1167, 1166],
  "duration": 120,
  "startTime": "2026-06-01 09:00",
  "endTime": "2026-06-01 11:00",
  "isTest": 0
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| testId | int\|null | 否 | null=新建 |
| title | string | 是 | 考试标题 |
| classIds | int[] | 是 | 参与班级 ID 数组 |
| questionIds | int[] | 是 | 题目 ID 数组 |
| duration | int | 是 | 考试时长（分钟） |
| startTime | string | 否 | 开始时间 |
| endTime | string | 否 | 结束时间 |
| isTest | int | 否 | 0=考试, 1=测试 |

---

### POST /test/publish/class

> 发布测试到班级

**权限:** 教师或管理员

**Request**

```json
{
  "testId": 1,
  "classIds": [27, 28]
}
```

---

### POST /test/publish/student

> 发布测试到指定学生

**权限:** 教师或管理员

**Request**

```json
{
  "testId": 1,
  "studentIds": ["123456", "2407110108"]
}
```

---

### GET /test/exportStudentReport/excel

> 导出学生成绩报告

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |
| testId | int | 是 | 考试 ID |

**Response** `200 OK`

```
Content-Type: application/octet-stream
```

---

### POST /test/extendTestTime

> 延长测试/考试时间

**权限:** 教师或管理员

**Request**

```json
{
  "testId": 100,
  "teacherId": "123456",
  "minutes": 60
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| testId | int | 是 | 考试 ID |
| teacherId | string | 是 | 教师工号 |
| minutes | int | 是 | 延长分钟数 |

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功"
}
```

---

### GET /test/exportClassAverageScores/excel

> 导出班级平均分

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| testId | int | 是 | 考试 ID |

---

### GET /test/exportFilteredScoreDetail/excel

> 导出筛选成绩详情

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| testId | int | 是 | 考试 ID |

---

## 11. 考试结果

### GET /result/student

> 获取学生考试成绩

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |

---

### GET /result/start

> 开始考试

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |
| testId | int | 是 | 考试 ID |
| isTest | int | 是 | 0=考试, 1=测试 |

> 需要已分配的考试记录

---

### GET /result/detail

> 获取考试结果详情

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |
| testId | int | 是 | 考试 ID |

---

### GET /result/find

> 查找考试结果

**Query Parameters**

| 参数 | 类型 | 说明 |
|------|------|------|
| — | — | 具体参数待补充 |

---

### GET /result/ai/detail

> 获取 AI 评分详情

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |
| testId | int | 是 | 考试 ID |

---

### PUT /result/saveAnswer

> 批量保存作答（PUT 请求，数组格式）

**Request**

```json
[
  {
    "studentId": "123456",
    "testId": 1,
    "questionId": 1168,
    "answer": "我的作答内容"
  }
]
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |
| testId | int | 是 | 考试 ID |
| questionId | int | 是 | 题目 ID |
| answer | string | 否 | 作答内容 |

> 需要进行中的考试；可一次传多道题的作答

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "成功",
  "data": {"code": 200, "successCount": 1, "message": "更新成功，共更新 1 条记录", "totalCount": 1}
}
```

---

### PUT /result/updateAnswer

> 批量更新作答（PUT 请求，数组格式）

**Request**

```json
[
  {
    "studentId": "123456",
    "testId": 1,
    "questionId": 1168,
    "answer": "修改后的答案"
  }
]
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |
| testId | int | 是 | 考试 ID |
| questionId | int | 是 | 题目 ID |
| answer | string | 否 | 作答内容 |

---

### POST /result/updateScoreAndComment

> 批量更新分数和评语（数组格式）

**权限:** 教师或管理员

**Request**

```json
[
  {
    "id": "2407110107839100",
    "studentId": "123456",
    "testId": 100,
    "questionId": 839,
    "actualScore": 30,
    "comment": "答案正确"
  }
]
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | string | 是 | 结果记录 ID |
| studentId | string | 是 | 学号 |
| testId | int | 是 | 考试 ID |
| questionId | int | 是 | 题目 ID |
| actualScore | int | 是 | 得分 |
| comment | string | 否 | 评语 |

---

### POST /result/update/ai

> 更新 AI 评分结果

**Request**

```json
{
  "studentId": "123456",
  "testId": 1
}
```

---

## 12. AI 问答

### POST /ai/student/ask

> 学生向 AI 提问

**Request**

```json
{
  "question": "什么是数据库事务的 ACID 特性？",
  "context": "第9章 事务管理与并发控制"
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| question | string | 是 | 问题内容 |
| context | string | 否 | 上下文信息 |

**Response** `200 OK`

```json
{
  "success": true,
  "sessionId": "f44b7d4b28ad49ba",
  "message": "请求已接收，正在异步处理",
  "timestamp": 1778663577502
}
```

> 响应 `success` 为 `true`（非 `code: 200` 格式）

---

### POST /ai/teacher/ask

> 教师向 AI 提问

**Request**

```json
{
  "question": "如何设计索引来优化这个查询？",
  "context": ""
}
```

**Response** `200 OK`

```json
{
  "success": true,
  "sessionId": "...",
  "message": "请求已接收，处理完成后将通过SSE推送结果",
  "timestamp": 1778663577502
}
```

---

## 13. 视频管理

### GET /video/tree

> 获取视频分类树

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "获取成功",
  "data": [...]
}
```

---

### GET /video/tree/select

> 获取视频分类树（下拉选择器用）

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "获取成功",
  "data": [...]
}
```

---

### POST /video/add

> 添加视频

**Request**

```json
{
  "title": "B+树索引原理",
  "url": "http://example.com/video.mp4",
  "description": "讲解B+树索引的数据结构与查询优化",
  "categoryId": 1
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| title | string | 是 | 视频标题 |
| url | string | 是 | 视频地址 |
| description | string | 否 | 描述 |
| categoryId | int | 否 | 分类 ID |

**Response** `200 OK`

```json
{
  "code": 200,
  "msg": "添加成功"
}
```

---

### POST /video/upload

> 上传视频文件

**Request**

```
Content-Type: multipart/form-data
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| file | file | 是 | 视频文件 |

---

## 14. 考勤管理

### GET /attendance/history

> 获取班级考勤历史

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| classId | int | 是 | 班级 ID |

---

### POST /attendance/save

> 保存考勤记录

**Request**

```json
{
  "classId": 27,
  "studentId": "123456",
  "status": "出勤",
  "date": "2026-05-13",
  "isAttendance": true
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| classId | int | 是 | 班级 ID |
| studentId | string | 是 | 学号 |
| status | string | 是 | 出勤/迟到/缺勤 |
| date | string | 是 | 日期 (yyyy-MM-dd) |
| isAttendance | bool | 否 | 是否计入考勤 |

---

### POST /attendance/update

> 更新考勤记录

**Request**

```json
{
  "timestamp": "2026-05-13T09:00:00",
  "studentId": "123456",
  "isAttendance": true,
  "loginTime": "2026-05-13T09:05:00"
}
```

---

### POST /attendance/updateStatus

> 批量更新考勤状态

**Request**

```json
{
  "studentId": "123456",
  "isAttendance": true,
  "date": "2026-05-13"
}
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| studentId | string | 是 | 学号 |
| isAttendance | bool | 是 | 是否出勤 |
| date | string | 是 | 日期 (yyyy-MM-dd) |

---

### GET /attendance/export

> 导出考勤 Excel

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| classId | int | 是 | 班级 ID |

**Response** `200 OK`

```
Content-Type: application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
```

---

### GET /attendance/exportSingle

> 导出单条考勤记录

**Query Parameters**

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | int | 是 | 考勤记录 ID |

---

## 15. 校历

### GET /academicCalendar/semesters

> 获取学期校历

**Response** `200 OK`

```json
{
  "code": 200,
  "data": [...]
}
```

---

## 16. 知识图谱

### GET /knowledgeGraph/data

> 获取知识图谱数据

**Response** `200 OK`

```json
{
  "code": 200,
  "data": {...}
}
```

---

### POST /upload/knowledge-graph

> 上传知识图谱文件

**Request**

```
Content-Type: multipart/form-data
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| file | file | 是 | 知识图谱文件 |

---

## 17. 文件上传

### POST /upload/image

> 上传图片

**Request**

```
Content-Type: multipart/form-data
```

| 字段 | 类型 | 必填 | 说明 |
|------|------|------|------|
| file | file | 是 | 图片文件 |

---

## 18. 通用说明

### 响应格式

所有接口遵循统一的响应结构：

**成功**

```json
{
  "code": 200,
  "msg": "成功",
  "total": 0,
  "data": ...
}
```

**失败**

```json
{
  "code": 400,
  "msg": "错误描述",
  "total": 0,
  "data": null
}
```

**权限不足**

```json
{
  "success": false,
  "error": "权限不足，需要以下角色之一: 管理员"
}
```

> 注：AI 接口 (`/ai/*`) 返回 `success` 字段而非 `code` 字段

### 认证

登录后将 token 放入请求头：

```
Authorization: Bearer eyJ0eXAiOiJKV1Q...
```

### 权限汇总

| 接口 | 最低角色 |
|------|----------|
| `/visit/count`, `/visit/increment` | 公开 |
| `/user/all` | 管理员 |
| `/teacher/insert` | 管理员 |
| `/test/getTeacherTest` | 教师 |
| `/test/publish` | 教师 |
| `/test/publish/ai` | 教师 |
| `/test/publish/class` | 教师 |
| `/test/publish/student` | 教师 |
| `/result/updateScoreAndComment` | 教师 |
| 其余全部接口 | 学生（已登录即可） |

### 状态图例

| 状态 | 说明 |
|------|------|
| ✅ | 已测试通过 |
| ⚠️ | 接口存在，需要特定参数或上下文 |
| ⏭️ | 需要已分配的进行中考试上下文 |

---

### Python 客户端

封装见 [acms_api.py](acms_api.py)，总计 91 个方法，38 个参数 dataclass。

```python
from acms_api import create_client
from acms_models import *

# 管理员登录
client = create_client("123456", "123456")
```
