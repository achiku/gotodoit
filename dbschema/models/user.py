# -*- coding: utf-8 -*-
from sqlalchemy import Column, DateTime, ForeignKey, Boolean
from sqlalchemy.dialects.postgresql import TEXT, UUID
from sqlalchemy.sql.expression import text
from sqlalchemy.orm import backref, relationship

from . import Base


class TodoUser(Base):

    __tablename__ = 'todo_user'

    uuid = Column(UUID, primary_key=True, server_default=text('uuid_generate_v1mc()'))
    username = Column(TEXT, nullable=False)
    email = Column(TEXT, nullable=False)
    password = Column(TEXT, nullable=False)
    status = Column(TEXT, nullable=False)


class AccessToken(Base):

    __tablename__ = 'access_token'

    token = Column(TEXT, primary_key=True)
    user_id = Column(UUID, ForeignKey('todo_user.uuid'), nullable=False)
    generated_at = Column(DateTime(timezone=True), nullable=False)
    is_active = Column(Boolean, nullable=False)

    user = relationship('TodoUser', backref=backref('tokens'))
