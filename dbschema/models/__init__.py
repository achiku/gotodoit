# -*- coding: utf-8 -*-
from sqlalchemy import MetaData, create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import scoped_session, sessionmaker

engine = create_engine(
    "postgresql+psycopg2://gotodoit_api@localhost/gotodoit"
)
meta = MetaData(engine)
Base = declarative_base(metadata=meta)
session = scoped_session(sessionmaker(bind=engine))

from .todo import Todo  # NOQA
from .user import TodoUser, AccessToken  # NOQA

__all__ = [
    # todo
    'Todo',

    # user
    'TodoUser', 'AccessToken',
]
