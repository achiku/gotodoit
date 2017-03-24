# -*- coding: utf-8 -*-
from models import session
from factories import user


if __name__ == '__main__':
    t = user.AccessTokenFactory.create()
    print(t.token)
    session.commit()
