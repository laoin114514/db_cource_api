"""
《数据库原理》智慧课程 API — 请求参数 dataclass
"""

from dataclasses import dataclass, field, asdict
from typing import Optional, Any


def to_dict(obj: Any) -> dict:
    """dataclass → dict，过滤 None 值，snake_case → camelCase"""
    if obj is None:
        return {}
    d = {}
    for k, v in asdict(obj).items():
        if v is None:
            continue
        parts = k.split("_")
        camel = parts[0] + "".join(p.capitalize() for p in parts[1:])
        d[camel] = v
    return d


# ── 通用分页 ──

@dataclass
class PageQuery:
    """通用分页参数（用于 teacher/class/student 的 showPage）"""
    page_num: int = 1
    page_size: int = 20
    query: str = ""           # 搜索关键词（后端要求此字段，空字符串表示无过滤）

    def as_params(self) -> dict:
        return {"pageNum": self.page_num, "pageSize": self.page_size, "query": self.query}


# ── 认证 ──

@dataclass
class LoginForm:
    """登录表单"""
    id: str          # 学号/工号
    password: str


# ── 用户 ──

@dataclass
class UserQuery:
    """用户查询"""
    id: str          # 用户ID（学号/工号）


@dataclass
class UserInsert:
    """添加用户"""
    id: str
    role: int
    name: str = ""
    password: str = ""


@dataclass
class UserUpdate:
    """更新用户"""
    id: str
    role: int


# ── 角色 ──

@dataclass
class RoleForm:
    """角色表单"""
    name: str
    description: str = ""


# ── 教师 ──

@dataclass
class TeacherQuery:
    """教师查询"""
    teacher_id: str      # 教师工号


@dataclass
class TeacherForm:
    """教师表单（新增/编辑共用）"""
    teacher_id: str      # 教师工号
    name: str
    gender: str = ""     # 男/女
    position: str = ""   # 职称
    password: str = ""   # 仅新增/更新时传


@dataclass
class TeacherClassRelation:
    """教师-班级关联"""
    class_id: int
    teacher_id: str


# ── 班级 ──

@dataclass
class ClassQuery:
    """按班级名称查询"""
    class_name: str


@dataclass
class ClassForm:
    """班级表单"""
    class_name: str
    headteacher_id: str = ""   # 班主任工号


# ── 学生 ──

@dataclass
class StudentQuery:
    """学生查询"""
    student_id: str     # 学号


@dataclass
class StudentForm:
    """学生表单（新增）"""
    student_id: str      # 学号
    name: str
    gender: Optional[str] = None
    phone: Optional[str] = None
    notes: Optional[str] = None
    class_name: Optional[str] = None


@dataclass
class StudentUpdate:
    """更新学生信息"""
    student_id: str
    name: Optional[str] = None
    gender: Optional[str] = None
    phone: Optional[str] = None
    notes: Optional[str] = None
    class_name: Optional[str] = None


# ── 题库分类 ──

@dataclass
class QuestionCategoryForm:
    """题库分类表单"""
    question_category: str   # 分类名称
    teacher_id: str = ""     # 可选，教师ID


# ── 题目 ──

@dataclass
class QuestionListQuery:
    """题目列表查询"""
    page_num: int = 1
    page_size: int = 100
    status: Optional[int] = None          # 状态（None=不过滤）
    keyword: str = ""                      # 关键词搜索
    category_id: Optional[int] = None      # 章节分类ID
    type: str = ""                         # 题型
    is_ai_generate: Optional[int] = None   # None=不过滤, 0=非AI, 1=AI生成
    teacher_id: str = ""                   # 出题教师ID

    def as_params(self) -> dict:
        p = {
            "pageNum": self.page_num,
            "pageSize": self.page_size,
            "status": self.status if self.status is not None else None,
            "isAiGenerate": self.is_ai_generate,
            "keyword": self.keyword or None,
            "categoryId": self.category_id or None,
            "type": self.type or None,
        }
        if self.teacher_id:
            p["teacherId"] = self.teacher_id
        return {k: v for k, v in p.items() if v is not None}


@dataclass
class QuestionFind:
    """查找题目（考试上下文）"""
    student_id: str
    question_id: int
    test_id: int


# ── 考试/测试 ──

@dataclass
class TestQuery:
    """教师测试查询"""
    teacher_id: str
    class_id: int


@dataclass
class TestPublish:
    """发布测试"""
    test_id: Optional[int] = None
    title: str = ""
    class_ids: list[int] = field(default_factory=list)
    question_ids: list[int] = field(default_factory=list)
    duration: int = 120        # 考试时长（分钟）
    start_time: str = ""       # 开始时间
    end_time: str = ""         # 结束时间
    is_test: int = 0           # 0=考试, 1=测试


# ── 考试结果 ──

@dataclass
class ResultStudent:
    """学生成绩查询"""
    student_id: str


@dataclass
class ResultStart:
    """开始考试"""
    student_id: str
    test_id: int
    is_test: int = 0           # 0=考试, 1=测试


# ── 视频 ──

@dataclass
class VideoForm:
    """视频表单"""
    title: str
    url: str = ""
    description: str = ""
    category_id: int = 0


# ── 考勤 ──

@dataclass
class AttendanceSave:
    """保存考勤"""
    class_id: int
    student_id: str = ""
    status: str = ""          # 出勤/迟到/缺勤
    date: str = ""            # yyyy-MM-dd
    is_attendance: bool = True


@dataclass
class AttendanceUpdate:
    """更新考勤记录"""
    timestamp: str = ""
    student_id: str = ""
    is_attendance: bool = True
    login_time: str = ""


@dataclass
class AttendanceStatus:
    """更新考勤状态"""
    student_id: str = ""
    is_attendance: bool = True
    date: str = ""


# ── AI ──

@dataclass
class AIAsk:
    """AI 提问"""
    question: str
    context: str = ""


# ── 角色 ──

@dataclass
class RoleUpdate:
    """更新角色"""
    id: int
    name: str = ""


# ── 班级更新 ──

@dataclass
class ClassUpdate:
    """更新班级"""
    class_id: int
    class_name: str = ""
    headteacher_id: str = ""


# ── 批量学生信息 ──

@dataclass
class BatchStudentInfo:
    """批量学生信息操作"""
    students: list[dict] = None
    operation: str = ""

    def __post_init__(self):
        if self.students is None:
            self.students = []


# ── 题目增改 ──

@dataclass
class QuestionInsert:
    """添加题目"""
    category_id: int
    type: str                   # 题型
    question_text: str
    answer: str
    score: int
    answer_time: int
    teacher_id: str
    keywords: str = ""
    is_ai_generate: bool = False


@dataclass
class QuestionUpdate:
    """更新题目"""
    id: int
    category_id: int = 0
    type: str = ""
    question_text: str = ""
    answer: str = ""
    score: int = 0
    answer_time: int = 0


# ── 考试扩展 ──

@dataclass
class ExtendTestTime:
    """延长测试/考试时间"""
    test_id: int
    teacher_id: str
    minutes: int = 0


# ── 作答 ──

@dataclass
class AnswerItem:
    """单条作答"""
    student_id: str = ""
    test_id: int = 0
    question_id: int = 0
    answer: str = ""


@dataclass
class SaveAnswer:
    """批量保存/更新作答（数组）"""
    items: list[AnswerItem] = None

    def __post_init__(self):
        if self.items is None:
            self.items = []


@dataclass
class ScoreItem:
    """单条分数评语"""
    id: str = ""               # 记录ID（如 "2407110107839100"）
    question_id: int = 0
    student_id: str = ""
    test_id: int = 0
    actual_score: int = 0
    comment: str = ""


@dataclass
class ScoreComment:
    """批量更新分数和评语（数组）"""
    items: list[ScoreItem] = None

    def __post_init__(self):
        if self.items is None:
            self.items = []


@dataclass
class AIResultUpdate:
    """更新 AI 评分结果"""
    student_id: str
    test_id: int


# ── 视频列表查询 ──

@dataclass
class VideoListQuery:
    """视频列表查询"""
    page_num: int = 1
    page_size: int = 20
    keyword: str = ""

    def as_params(self) -> dict:
        return {"pageNum": self.page_num, "pageSize": self.page_size, "keyword": self.keyword}
