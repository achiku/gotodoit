# -*- coding: utf-8 -*-
from sqlalchemy import Column, DateTime, ForeignKey, Boolean
from sqlalchemy.dialects.postgresql import BIGINT, TEXT, UUID
from sqlalchemy.sql.expression import text
from sqlalchemy.orm import backref, relationship

from . import Base


class Todo(Base):

    __tablename__ = 'todo'

    uuid = Column(UUID, primary_key=True, server_default=text('uuid_generate_v1mc()'))
    user_id = Column(UUID, ForeignKey('todo_user.uuid'), nullable=False)
    name = Column(TEXT, nullable=False)
    duration = Column(BIGINT, nullable=False)
    started_at = Column(DateTime(timezone=True), nullable=False)
    is_completed = Column(Boolean, nullable=False)

    user = relationship('TodoUser', backref=backref('todos'))
