# -*- coding: utf-8 -*-
from datetime import datetime

import factory
from factory.alchemy import SQLAlchemyModelFactory

from models import user, session


class TodoUserFactory(SQLAlchemyModelFactory):
    class Meta:
        model = user.TodoUser
        sqlalchemy_session = session

    email = factory.Faker('email', locale='ja_JP')
    username = factory.Faker('user_name', locale='ja_JP')
    password = factory.Faker('md5')
    status = 'active'


class AccessTokenFactory(SQLAlchemyModelFactory):
    class Meta:
        model = user.AccessToken
        sqlalchemy_session = session

    token = factory.Faker('md5')
    is_active = True
    generated_at = factory.LazyFunction(datetime.now)

    user = factory.SubFactory('factories.user.TodoUserFactory')
