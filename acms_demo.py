"""
《数据库原理》智慧课程 API — 使用示例
"""

import json
import sys
from collections import Counter

from acms_api import create_client
from acms_models import (
    TeacherQuery, ClassQuery, StudentQuery, StudentUpdate,
    QuestionListQuery, QuestionFind, ResultStudent, ResultStart,
    PageQuery, AIAsk,
)

if __name__ == "__main__":
    sys.stdout.reconfigure(encoding="utf-8")

    client = create_client("123456", "123456")
    print(f"登录成功: {client.user_info['name']} ({client.user_info['position']})")

    # ── 访问量 ──
    vc = client.get_visit_count()
    print(f"访问量: {vc['data']}")

    # ── 教师 ──
    teachers = client.get_all_teachers()
    print(f"教师: {[t['name'] for t in teachers['data']]}")

    # ── 班级 ──
    names = client.get_class_names()
    print(f"班级: {names['data']}")

    for t in teachers["data"]:
        c = client.find_class_by_teacher(t["teacherId"])
        if c["code"] == 200:
            for cls in c["data"]:
                print(f"  {t['name']} -> {cls['className']} "
                      f"(学生数: {len(json.loads(cls['listOfStudents']))})")

    # ── 题库分类 ──
    cats = client.get_question_categories()
    print(f"题库分类: {[c['questionCategory'] for c in cats['data']]}")

    # ── 题目 ──
    questions = client.get_all_questions()
    print(f"题目总数: {len(questions)}")

    type_dist = Counter(q["type"] for q in questions)
    print(f"题型分布: {dict(type_dist)}")

    cat_dist = Counter(q["categoryId"] for q in questions)
    print(f"章节分布: {dict(sorted(cat_dist.items()))}")

    # ── 参数类调用演示 ──
    print("\n--- 参数类调用演示 ---")

    r = client.query_teacher(TeacherQuery(teacher_id="20050027"))
    print(f"查询教师20050027: {r['data']['name'] if r['code'] == 200 else r}")

    r = client.get_question_list(QuestionListQuery(
        page_size=3, category_id=9, type="选择题",
    ))
    if r["code"] == 200:
        for item in r["data"]["items"]:
            print(f"  [{item['type']}] {item['questionText'][:50]}...")

    # ── 更多示例 ──
    print("\n--- 更多操作示例 ---")

    # 分页
    r = client.get_student_page(PageQuery(page_num=1, page_size=5, query=""))
    print(f"学生分页: total={r['data']['total']}")

    # 搜索
    r = client.search_student_by_prefix("24071101")
    print(f"前缀搜索: {len(r['data'])} 条")

    # 视频树
    r = client.get_video_tree()
    print(f"视频树: code={r.get('code')}")

    # 校历
    r = client.get_semesters()
    print(f"校历: code={r.get('code')}")

    # AI 提问
    r = client.ai_student_ask(AIAsk(question="什么是数据库事务？"))
    print(f"AI提问: success={r.get('success')}")
