# -*- coding: utf-8 -*-
from models import session
from factories import todo, user

if __name__ == '__main__':
    at = user.AccessTokenFactory.create()
    for _ in range(10):
        t = todo.TodoFactory.create(user=at.user)
        print(t.name)
    session.commit()
