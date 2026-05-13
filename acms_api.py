"""
《数据库原理》智慧课程 API 客户端
基地地址: http://172.21.44.162:5174/acms
"""

from typing import Optional

import requests

from acms_models import (
    to_dict,
    LoginForm,
    UserQuery, UserInsert, UserUpdate,
    RoleForm, RoleUpdate,
    TeacherQuery, TeacherForm, TeacherClassRelation,
    ClassQuery, ClassForm, ClassUpdate,
    StudentQuery, StudentForm, StudentUpdate, BatchStudentInfo,
    QuestionCategoryForm, QuestionListQuery, QuestionFind,
    QuestionInsert, QuestionUpdate,
    TestQuery, TestPublish, ExtendExamTime,
    ResultStudent, ResultStart,
    SaveAnswer, ScoreComment, AIResultUpdate,
    VideoForm, VideoListQuery,
    AttendanceSave, AttendanceUpdate, AttendanceStatus,
    AIAsk,
    PageQuery,
)


class ACMSClient:
    """数据库原理智慧课程 API 客户端"""

    def __init__(self, base_url: str = "http://172.21.44.162:5174/acms"):
        self.base_url = base_url.rstrip("/")
        self.session = requests.Session()
        self.token: Optional[str] = None
        self.user_info: Optional[dict] = None

    # ── 认证 ──────────────────────────────────────────────

    def login(self, form: LoginForm) -> dict:
        """登录获取 token，自动保存到 session"""
        resp = self.session.post(
            f"{self.base_url}/auth/login",
            json=to_dict(form),
        )
        data = resp.json()
        if data.get("code") == 200:
            self.token = data["data"]["token"]
            self.user_info = data["data"]["user"]
            self.session.headers.update(
                {"Authorization": f"Bearer {self.token}"}
            )
        return data

    def logout(self) -> None:
        """清除登录状态"""
        self.token = None
        self.user_info = None
        self.session.headers.pop("Authorization", None)

    # ── 访问量 ────────────────────────────────────────────

    def get_visit_count(self) -> dict:
        """获取网站访问量"""
        return self.session.get(f"{self.base_url}/visit/count").json()

    def increment_visit(self) -> dict:
        """增加访问量"""
        return self.session.post(f"{self.base_url}/visit/increment").json()

    # ── 用户 ──────────────────────────────────────────────

    def query_user(self, query: UserQuery) -> dict:
        """按 ID 查询用户"""
        return self.session.get(
            f"{self.base_url}/user/query",
            params=to_dict(query),
        ).json()

    def get_all_users(self) -> dict:
        """获取所有用户（需要管理员权限）"""
        return self.session.get(f"{self.base_url}/user/all").json()

    def insert_user(self, form: UserInsert) -> dict:
        """添加用户"""
        return self.session.post(
            f"{self.base_url}/user/insert", json=to_dict(form)
        ).json()

    def update_user(self, form: UserUpdate) -> dict:
        """更新用户"""
        return self.session.post(
            f"{self.base_url}/user/update", json=to_dict(form)
        ).json()

    # ── 角色 ──────────────────────────────────────────────

    def get_all_roles(self) -> dict:
        """获取所有角色"""
        return self.session.get(f"{self.base_url}/role/all").json()

    def add_role(self, form: RoleForm) -> dict:
        """添加角色"""
        return self.session.post(
            f"{self.base_url}/role/add", json=to_dict(form)
        ).json()

    def update_role(self, form: RoleUpdate) -> dict:
        """更新角色"""
        return self.session.post(
            f"{self.base_url}/role/update", json=to_dict(form)
        ).json()

    # ── 教师 ──────────────────────────────────────────────

    def query_teacher(self, query: TeacherQuery) -> dict:
        """按工号查询教师"""
        return self.session.get(
            f"{self.base_url}/teacher/query",
            params=to_dict(query),
        ).json()

    def get_all_teachers(self) -> dict:
        """获取所有教师"""
        return self.session.get(f"{self.base_url}/teacher/findAll").json()

    def get_teacher_info(self) -> dict:
        """获取教师简要信息列表"""
        return self.session.get(f"{self.base_url}/teacher/acquireInfo").json()

    def get_teacher_page(self, page: PageQuery = PageQuery()) -> dict:
        """分页获取教师"""
        return self.session.get(
            f"{self.base_url}/teacher/showPage",
            params=page.as_params(),
        ).json()

    def insert_teacher(self, form: TeacherForm) -> dict:
        """添加教师（需要管理员权限）"""
        return self.session.post(
            f"{self.base_url}/teacher/insert", json=to_dict(form)
        ).json()

    def update_teacher(self, form: TeacherForm) -> dict:
        """更新教师信息（后端期望 {teacher: {...}, password: ...}）"""
        d = to_dict(form)
        password = d.pop("password", "")
        return self.session.post(
            f"{self.base_url}/teacher/update",
            json={"teacher": d, "password": password},
        ).json()

    def teacher_join_class(self, rel: TeacherClassRelation) -> dict:
        """教师加入班级"""
        return self.session.post(
            f"{self.base_url}/teacher/join", json=to_dict(rel)
        ).json()

    def teacher_leave_class(self, rel: TeacherClassRelation) -> dict:
        """从班级移除教师"""
        return self.session.post(
            f"{self.base_url}/teacher/remove", json=to_dict(rel)
        ).json()

    # ── 班级 ──────────────────────────────────────────────

    def query_class(self, query: ClassQuery) -> dict:
        """按名称查询班级"""
        return self.session.get(
            f"{self.base_url}/class/queryClass",
            params=to_dict(query),
        ).json()

    def get_class_names(self) -> dict:
        """获取所有班级名称"""
        return self.session.get(f"{self.base_url}/class/acquireName").json()

    def get_class_page(self, page: PageQuery = PageQuery()) -> dict:
        """分页获取班级"""
        return self.session.get(
            f"{self.base_url}/class/showPage",
            params=page.as_params(),
        ).json()

    def find_class_by_teacher(self, teacher_id: str) -> dict:
        """按教师 ID 查找其负责的班级"""
        return self.session.get(
            f"{self.base_url}/class/findByTeacher",
            params={"teacherId": teacher_id},
        ).json()

    def get_all_classes(self) -> dict:
        """获取所有班级（含学生/教师列表）"""
        return self.session.get(f"{self.base_url}/class/findAll").json()

    def get_class_by_name(self, class_name: str) -> dict:
        """按名称获取班级"""
        return self.session.get(
            f"{self.base_url}/class/getByName",
            params={"className": class_name},
        ).json()

    def find_class_by_teacher_and_name(self, teacher_id: str, class_name: str) -> dict:
        """按教师和名称查班级"""
        return self.session.get(
            f"{self.base_url}/class/findByTeacherAndName",
            params={"teacherId": teacher_id, "className": class_name},
        ).json()

    def get_class_students(self, class_id: int) -> dict:
        """获取班级学生列表"""
        return self.session.get(
            f"{self.base_url}/class/students",
            params={"classId": class_id},
        ).json()

    def insert_class(self, form: ClassForm) -> dict:
        """添加班级"""
        return self.session.post(
            f"{self.base_url}/class/insert", json=to_dict(form)
        ).json()

    def update_class(self, form: ClassUpdate) -> dict:
        """更新班级信息"""
        return self.session.post(
            f"{self.base_url}/class/update", json=to_dict(form)
        ).json()

    # ── 学生 ──────────────────────────────────────────────

    def query_student(self, query: StudentQuery) -> dict:
        """按学号查询学生"""
        return self.session.get(
            f"{self.base_url}/student/query",
            params=to_dict(query),
        ).json()

    def get_student_page(self, page: PageQuery = PageQuery()) -> dict:
        """分页获取学生"""
        return self.session.get(
            f"{self.base_url}/student/showPage",
            params=page.as_params(),
        ).json()

    def find_no_class_students(self) -> dict:
        """查找未分班学生"""
        return self.session.get(f"{self.base_url}/student/findNoClass").json()

    def search_student_by_prefix(self, prefix: str) -> dict:
        """按学号前缀搜索学生"""
        return self.session.get(
            f"{self.base_url}/student/searchByStudentIdPrefix",
            params={"prefix": prefix},
        ).json()

    def insert_student(self, form: StudentForm) -> dict:
        """添加学生"""
        return self.session.post(
            f"{self.base_url}/student/insert", json=to_dict(form)
        ).json()

    def update_student(self, form: StudentUpdate) -> dict:
        """更新学生信息（后端期望 {student: {...}}）"""
        return self.session.post(
            f"{self.base_url}/student/update",
            json={"student": to_dict(form)},
        ).json()

    def delete_student(self, student_ids: list[int]) -> dict:
        """删除学生（传入学号整数列表）"""
        return self.session.post(
            f"{self.base_url}/student/delete", json=student_ids
        ).json()

    def update_student_class(self, student_ids: list[str], class_id: int) -> dict:
        """批量更新学生班级"""
        return self.session.post(
            f"{self.base_url}/student/updateClass",
            json={"studentIds": student_ids, "classId": class_id},
        ).json()

    def batch_student_info(self, form: BatchStudentInfo) -> dict:
        """批量学生信息操作"""
        return self.session.post(
            f"{self.base_url}/student/batchInfo", json=to_dict(form)
        ).json()

    def export_students(self, class_id: int) -> bytes:
        """导出学生 Excel（返回 bytes）"""
        return self.session.get(
            f"{self.base_url}/student/export",
            params={"classId": class_id},
        ).content

    def download_student_template(self) -> bytes:
        """下载学生导入模板（返回 bytes）"""
        return self.session.get(
            f"{self.base_url}/student/downloadTemplate"
        ).content

    def batch_import_students(self, file_path: str) -> dict:
        """批量导入学生 Excel"""
        with open(file_path, "rb") as f:
            return self.session.post(
                f"{self.base_url}/student/batchImport",
                files={"file": f},
            ).json()

    def upload_student_file(self, file_path: str) -> dict:
        """上传学生 Excel 文件"""
        with open(file_path, "rb") as f:
            return self.session.post(
                f"{self.base_url}/studentFile/upload",
                files={"file": f},
            ).json()

    # ── 题库分类 ──────────────────────────────────────────

    def get_question_categories(self) -> dict:
        """获取所有题库分类"""
        return self.session.get(
            f"{self.base_url}/question-category/list"
        ).json()

    def add_question_category(self, form: QuestionCategoryForm) -> dict:
        """添加题库分类"""
        return self.session.post(
            f"{self.base_url}/question-category/add", json=to_dict(form)
        ).json()

    # ── 题目 ──────────────────────────────────────────────

    def get_question_list(self, query: QuestionListQuery = QuestionListQuery()) -> dict:
        """获取题目列表（支持分类/题型/教师/关键词筛选）"""
        return self.session.get(
            f"{self.base_url}/question/list",
            params=query.as_params(),
        ).json()

    def get_all_questions(self, teacher_id: str = "", category_id: int = None) -> list[dict]:
        """一次性获取所有题目"""
        q = QuestionListQuery(page_size=500, teacher_id=teacher_id)
        if category_id is not None:
            q.category_id = category_id
        resp = self.get_question_list(q)
        if resp.get("code") != 200:
            return []
        return resp["data"]["items"]

    def query_question(self, question_id: int) -> dict:
        """查询单个题目详情"""
        return self.session.get(
            f"{self.base_url}/question/query",
            params={"id": question_id},
        ).json()

    def find_question(self, query: QuestionFind) -> dict:
        """查找单个题目（考试上下文）"""
        return self.session.get(
            f"{self.base_url}/question/find",
            params=to_dict(query),
        ).json()

    def insert_question(self, form: QuestionInsert, file_path: str = None) -> dict:
        """添加题目（支持文件上传）"""
        d = to_dict(form)
        if file_path:
            with open(file_path, "rb") as f:
                return self.session.post(
                    f"{self.base_url}/question/insert",
                    data=d,
                    files={"file": f},
                ).json()
        return self.session.post(
            f"{self.base_url}/question/insert",
            json=d,
        ).json()

    def update_question(self, form: QuestionUpdate) -> dict:
        """更新题目"""
        return self.session.post(
            f"{self.base_url}/question/update", json=to_dict(form)
        ).json()

    def batch_delete_questions(self, ids: list[int]) -> dict:
        """批量删除题目"""
        return self.session.post(
            f"{self.base_url}/question/batchDelete", json=ids
        ).json()

    def download_question_template(self) -> bytes:
        """下载题目导入模板 Excel（返回 bytes）"""
        return self.session.get(
            f"{self.base_url}/questionFile/template/download"
        ).content

    def upload_question_file(self, file_path: str) -> dict:
        """上传题目文件"""
        with open(file_path, "rb") as f:
            return self.session.post(
                f"{self.base_url}/questionFile/upload",
                files={"file": f},
            ).json()

    # ── 考试/测试 ─────────────────────────────────────────

    def get_teacher_tests(self, query: TestQuery) -> dict:
        """获取教师发布的测试（需要教师/管理员权限）"""
        return self.session.get(
            f"{self.base_url}/test/getTeacherTest",
            params=to_dict(query),
        ).json()

    def get_class_test(self, class_id: int) -> dict:
        """获取班级测试"""
        return self.session.get(
            f"{self.base_url}/test/getClassTest",
            params={"classId": class_id},
        ).json()

    def get_student_test(self, student_id: str) -> dict:
        """获取学生测试列表"""
        return self.session.get(
            f"{self.base_url}/test/getStudentTest",
            params={"studentId": student_id},
        ).json()

    def publish_test(self, form: TestPublish) -> dict:
        """发布测试（需要教师/管理员权限）"""
        return self.session.post(
            f"{self.base_url}/test/publish", json=to_dict(form)
        ).json()

    def publish_test_ai(self, form: TestPublish) -> dict:
        """发布 AI 测试"""
        return self.session.post(
            f"{self.base_url}/test/publish/ai", json=to_dict(form)
        ).json()

    def publish_test_to_class(self, test_id: int, class_ids: list[int]) -> dict:
        """发布测试到班级"""
        return self.session.post(
            f"{self.base_url}/test/publish/class",
            json={"testId": test_id, "classIds": class_ids},
        ).json()

    def publish_test_to_student(self, test_id: int, student_ids: list[str]) -> dict:
        """发布测试到指定学生"""
        return self.session.post(
            f"{self.base_url}/test/publish/student",
            json={"testId": test_id, "studentIds": student_ids},
        ).json()

    def extend_exam_time(self, form: ExtendExamTime) -> dict:
        """延长考试时间"""
        return self.session.post(
            f"{self.base_url}/test/extendExamTime", json=to_dict(form)
        ).json()

    def export_student_report(self, student_id: str, test_id: int = 0) -> bytes:
        """导出学生成绩报告 Excel（返回 bytes）"""
        return self.session.get(
            f"{self.base_url}/test/exportStudentReport/excel",
            params={"studentId": student_id, "testId": test_id},
        ).content

    def export_class_avg_scores(self, test_id: int) -> bytes:
        """导出班级平均分 Excel（返回 bytes）"""
        return self.session.get(
            f"{self.base_url}/test/exportClassAverageScores/excel",
            params={"testId": test_id},
        ).content

    def export_filtered_score_detail(self, test_id: int) -> bytes:
        """导出筛选成绩详情 Excel（返回 bytes）"""
        return self.session.get(
            f"{self.base_url}/test/exportFilteredScoreDetail/excel",
            params={"testId": test_id},
        ).content

    # ── 考试结果 ──────────────────────────────────────────

    def get_student_result(self, query: ResultStudent) -> dict:
        """获取学生考试结果"""
        return self.session.get(
            f"{self.base_url}/result/student",
            params=to_dict(query),
        ).json()

    def start_result(self, query: ResultStart) -> dict:
        """开始考试"""
        return self.session.get(
            f"{self.base_url}/result/start",
            params=to_dict(query),
        ).json()

    def get_result_detail(self, student_id: str, test_id: int = 0) -> dict:
        """获取考试结果详情"""
        return self.session.get(
            f"{self.base_url}/result/detail",
            params={"studentId": student_id, "testId": test_id},
        ).json()

    def find_result(self, params: dict = None) -> dict:
        """查找考试结果"""
        return self.session.get(
            f"{self.base_url}/result/find", params=params
        ).json()

    def get_ai_result_detail(self, student_id: str, test_id: int = 0) -> dict:
        """获取 AI 评分详情"""
        return self.session.get(
            f"{self.base_url}/result/ai/detail",
            params={"studentId": student_id, "testId": test_id},
        ).json()

    def save_answer(self, form: SaveAnswer) -> dict:
        """保存作答"""
        return self.session.post(
            f"{self.base_url}/result/saveAnswer", json=to_dict(form)
        ).json()

    def update_answer(self, form: SaveAnswer) -> dict:
        """更新作答"""
        return self.session.post(
            f"{self.base_url}/result/updateAnswer", json=to_dict(form)
        ).json()

    def update_score_comment(self, form: ScoreComment) -> dict:
        """更新分数和评语"""
        return self.session.post(
            f"{self.base_url}/result/updateScoreAndComment", json=to_dict(form)
        ).json()

    def update_ai_result(self, form: AIResultUpdate) -> dict:
        """更新 AI 评分结果"""
        return self.session.post(
            f"{self.base_url}/result/update/ai", json=to_dict(form)
        ).json()

    # ── AI ────────────────────────────────────────────────

    def ai_student_ask(self, form: AIAsk) -> dict:
        """AI 学生提问"""
        return self.session.post(
            f"{self.base_url}/ai/student/ask", json=to_dict(form)
        ).json()

    def ai_teacher_ask(self, question: str, context: str = "") -> dict:
        """AI 教师提问"""
        return self.session.post(
            f"{self.base_url}/ai/teacher/ask",
            json={"question": question, "context": context},
        ).json()

    # ── 视频 ──────────────────────────────────────────────

    def get_video_list(self, query: VideoListQuery = VideoListQuery()) -> dict:
        """获取视频列表"""
        return self.session.get(
            f"{self.base_url}/video/list",
            params=query.as_params(),
        ).json()

    def get_video_tree(self) -> dict:
        """获取视频分类树"""
        return self.session.get(f"{self.base_url}/video/tree").json()

    def get_video_tree_select(self) -> dict:
        """获取视频分类树（选择器用）"""
        return self.session.get(
            f"{self.base_url}/video/tree/select"
        ).json()

    def add_video(self, form: VideoForm) -> dict:
        """添加视频"""
        return self.session.post(
            f"{self.base_url}/video/add", json=to_dict(form)
        ).json()

    def upload_video(self, file_path: str) -> dict:
        """上传视频文件"""
        with open(file_path, "rb") as f:
            return self.session.post(
                f"{self.base_url}/video/upload",
                files={"file": f},
            ).json()

    # ── 考勤 ──────────────────────────────────────────────

    def get_attendance_history(self, class_id: int) -> dict:
        """获取班级考勤历史"""
        return self.session.get(
            f"{self.base_url}/attendance/history",
            params={"classId": class_id},
        ).json()

    def save_attendance(self, form: AttendanceSave) -> dict:
        """保存考勤"""
        return self.session.post(
            f"{self.base_url}/attendance/save", json=to_dict(form)
        ).json()

    def update_attendance(self, form: AttendanceUpdate) -> dict:
        """更新考勤记录"""
        return self.session.post(
            f"{self.base_url}/attendance/update", json=to_dict(form)
        ).json()

    def update_attendance_status(self, form: AttendanceStatus) -> dict:
        """批量更新考勤状态"""
        return self.session.post(
            f"{self.base_url}/attendance/updateStatus", json=to_dict(form)
        ).json()

    def export_attendance(self, class_id: int) -> bytes:
        """导出考勤 Excel（返回 bytes）"""
        return self.session.get(
            f"{self.base_url}/attendance/export",
            params={"classId": class_id},
        ).content

    def export_attendance_single(self, record_id: int) -> dict:
        """导出单条考勤记录"""
        return self.session.get(
            f"{self.base_url}/attendance/exportSingle",
            params={"id": record_id},
        ).json()

    # ── 校历 ──────────────────────────────────────────────

    def get_semesters(self) -> dict:
        """获取学期校历"""
        return self.session.get(
            f"{self.base_url}/academicCalendar/semesters"
        ).json()

    # ── 知识图谱 ──────────────────────────────────────────

    def get_knowledge_graph(self) -> dict:
        """获取知识图谱数据"""
        return self.session.get(
            f"{self.base_url}/knowledgeGraph/data"
        ).json()

    def upload_knowledge_graph(self, file_path: str) -> dict:
        """上传知识图谱文件"""
        with open(file_path, "rb") as f:
            return self.session.post(
                f"{self.base_url}/upload/knowledge-graph",
                files={"file": f},
            ).json()

    # ── 上传 ──────────────────────────────────────────────

    def upload_image(self, file_path: str) -> dict:
        """上传图片"""
        with open(file_path, "rb") as f:
            return self.session.post(
                f"{self.base_url}/upload/image",
                files={"file": f},
            ).json()

    # ── 工具 ──────────────────────────────────────────────

    def is_logged_in(self) -> bool:
        """检查是否已登录"""
        return self.token is not None


def create_client(user_id: str, password: str) -> ACMSClient:
    """创建客户端并自动登录"""
    client = ACMSClient()
    result = client.login(LoginForm(id=user_id, password=password))
    if result.get("code") != 200:
        raise Exception(f"登录失败: {result.get('msg', '未知错误')}")
    return client
