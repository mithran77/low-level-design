from datetime import datetime

class Comment:
    def __init__(self, author, content: str):
        self.id = id(self)
        self.author = author
        self.content = content
        self.creation_date = datetime.now()
