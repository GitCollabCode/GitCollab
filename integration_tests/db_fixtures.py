import pytest
from sqlalchemy.orm import sessionmaker
from sqlalchemy import create_engine
from models.blacklist import BlacklistFactory



Session = sessionmaker()


@pytest.fixture(scope="session")
def engine():
    return create_engine('postgresql://postgres:postgres@localhost:5432/postgres')

@pytest.fixture
def session(engine):
    """Returns an sqlalchemy session, and after the test tears down everything properly."""
    connection = engine.connect()
    transaction = connection.begin()
    session = Session(bind=connection)
    BlacklistFactory._meta.sqlalchemy_session = session
    yield session
    session.close()
    transaction.rollback()
    connection.close()

@pytest.fixture(scope='session')
def api_route():
    return "http://localhost:8080/api/"
