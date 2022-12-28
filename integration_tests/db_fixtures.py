import pytest
from sqlalchemy.orm import sessionmaker
from sqlalchemy import create_engine
from models.blacklist import BlacklistFactory



Session = sessionmaker()


@pytest.fixture(scope="function")
def engine():
    return create_engine('postgresql://postgres:postgres@localhost:5432/postgres')

@pytest.fixture(scope="function")
def session(engine):
    """Returns an sqlalchemy session, and after the test tears down everything properly."""
    connection = engine.connect()
    transaction = connection.begin()
    engine.create_scoped_session()
    session = Session(bind=engine)
    BlacklistFactory._meta.sqlalchemy_session = session
    yield session
    session.close()
    session.rollback()
    connection.close()

@pytest.fixture(scope='session')
def api_route():
    return "http://localhost:8080/api/"
