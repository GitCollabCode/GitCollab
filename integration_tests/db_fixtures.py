import pytest
from sqlalchemy.orm import sessionmaker
from sqlalchemy import create_engine

engine = create_engine('postgresql://postgres:postgres@localhost:5432/postgres')
Session = sessionmaker()

@pytest.fixture(scope='module')
def connection():
    connection = engine.connect()
    yield connection
    connection.close()

@pytest.fixture(scope='function')
def temp_session(connection):
    transaction = connection.begin()
    session = Session(bind=connection)
    yield session
    session.close()
    transaction.rollback()

@pytest.fixture(scope='function')
def session(connection):
    connection.begin()
    session = Session(bind=connection)
    yield session
    session.close()

@pytest.fixture(scope='function')
def api_route():
    return "http://localhost:8080/api/"
