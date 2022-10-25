import pytest
from pytest_postgresql import factories

# Global fixtures to go here

#postgres_db = factories.postgresql('postgresql_my_proc')
#
#@pytest.fixture()
#def setup_db(postgres_db):
#    def db_creator():
#        return postgres_db.cursor().connection
#
#    engine = create_engine('postgresql+psycopg2://', creator=dbcreator)
#    Base.metadata.create_all(engine)
#    Session = sessionmaker(bind=engine)
#    session = Session()
#    yield session