# -*- coding: utf-8 -*-
import random
from datetime import datetime

import factory
from factory.alchemy import SQLAlchemyModelFactory
from models import session, todo


class TodoFactory(SQLAlchemyModelFactory):
    class Meta:
        model = todo.Todo
        sqlalchemy_session = session

    name = factory.Sequence(lambda n: 'todo %03d' % n)
    duration = factory.LazyAttribute(lambda n: random.randint(1000, 10000))
    started_at = factory.LazyFunction(datetime.now)
    is_completed = False

    user = factory.SubFactory('factories.user.TodoUserFactory')
