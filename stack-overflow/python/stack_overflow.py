from typing import Dict

from user import User
from question import Question
from answer import Answer
from tag import Tag


class StackOverflow:

    def __init__(self):
        self.users: Dict[int, User] = {}
        self.questions: Dict[int, Question] = {}
        self.answers: Dict[int, Answer] = {}
        self.tags: Dict[str, Tag] = {}

    def create_user(self, username, email):
        user_id = len(self.users) + 1
        user = User(user_id, username, email)
        self.users[user_id] = user
        return user

    def ask_question(self, user: User, title, content, tags):
        question = user.ask_question(title, content, tags)
        self.questions[question.id] = question
        for tag in question.tags:
            self.tags.setdefault(tag.name, tag)
        return question

    def answer_question(self, user: User, question, content):
        answer = user.answer_question(question, content)
        self.answers[answer.id] = answer
        return answer

    def add_comment(self, user: User, commentable, content):
        return user.comment_on(commentable, content)

    def vote_question(self, user: User, question, value):
        question.vote(user, value)

    def vote_answer(self, user, answer: Answer, value):
        answer.vote(user, value)

    def accept_answer(self, answer):
        answer.accept()

    def search_questions(self, query: str):
        res = []
        for q in self.questions.values():
            tags = set()
            for t in q.tags: tags.add(t.name.lower())
            query = query.lower()
            # In q.title or q.content
            if (query in q.title.lower() or
                query in q.content.lower()):
                res.append(q)
            # OR in tags
            elif (query in tags):
                res.append(q)

        return res

    def get_questions_by_user(self, user):
        return user.questions

    def get_user(self, user_id):
        return self.users.get(user_id)

    def get_question(self, question_id):
        return self.questions.get(question_id)

    def get_answer(self, answer_id):
        return self.answers.get(answer_id)

    def get_tag(self, name: str):
        return self.tags.get(name)
