# -*- coding: utf-8 -*-
from __future__ import with_statement

import os
import sys
from logging.config import fileConfig

from alembic import context
from sqlalchemy import engine_from_config, pool

# this is the Alembic Config object, which provides
# access to the values within the .ini file in use.
config = context.config

# Interpret the config file for Python logging.
# This line sets up loggers basically.
fileConfig(config.config_file_name)

# add your model's MetaData object here
# for 'autogenerate' support
# from myapp import mymodel
# target_metadata = mymodel.Base.metadata
project_dir = os.path.abspath(os.getcwd())
sys.path.append(project_dir)

from dbschema.models import Base  # NOQA
target_metadata = Base.metadata

# other values from the config, defined by the needs of env.py,
# can be acquired:
# my_important_option = config.get_main_option("my_important_option")
# ... etc.


def get_db_url():
    # prd/stgにはAPP_ENVを含むDB接続用情報が環境変数として定義済み
    env_name = os.getenv('APP_ENV', None)
    if env_name is None:
        return config.get_main_option("sqlalchemy.url")
    db_url = os.getenv('DB_URL')
    db_name = os.getenv('API_DB_NAME')
    db_user_name = os.getenv('API_DB_USER')
    db_user_pass = os.getenv('API_DB_PASS')
    url = 'postgres://{0}:{1}@{2}/{3}'.format(db_user_name, db_user_pass, db_url, db_name)
    return url


def run_migrations_offline():
    """Run migrations in 'offline' mode.

    This configures the context with just a URL
    and not an Engine, though an Engine is acceptable
    here as well.  By skipping the Engine creation
    we don't even need a DBAPI to be available.

    Calls to context.execute() here emit the given string to the
    script output.

    """
    url = get_db_url()
    context.configure(
        url=url,
        target_metadata=target_metadata,
        compare_type=True,
        compare_server_default=True,
        literal_binds=True
    )

    with context.begin_transaction():
        context.run_migrations()


def run_migrations_online():
    """Run migrations in 'online' mode.

    In this scenario we need to create an Engine
    and associate a connection with the context.

    """
    # alembic.iniから取得したコンフィグのURL部分のみ上書き
    config.set_section_option('alembic', 'sqlalchemy.url', get_db_url())
    connectable = engine_from_config(
        config.get_section(config.config_ini_section),
        prefix='sqlalchemy.',
        poolclass=pool.NullPool)

    with connectable.connect() as connection:
        context.configure(
            connection=connection,
            compare_type=True,
            compare_server_default=True,
            target_metadata=target_metadata
        )

        with context.begin_transaction():
            context.run_migrations()


if context.is_offline_mode():
    run_migrations_offline()
else:
    run_migrations_online()
