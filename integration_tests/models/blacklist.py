from sqlalchemy import Column, Integer, String, TIMESTAMP, VARCHAR
import factory
import datetime
import factory.fuzzy as fuzzy
from db_fixtures import *
from sqlalchemy.orm import declarative_base

Base = declarative_base()

class BlacklistModel(Base):
    __tablename__ = "jwt_blacklist"
    
    uuid = Column(Integer)
    invalidated_time = Column(TIMESTAMP)
    jwt = Column(VARCHAR, primary_key=True)


class BlacklistFactory(factory.alchemy.SQLAlchemyModelFactory):
    class Meta:
        model = BlacklistModel
    
    uuid = factory.Sequence(lambda n: '%d' % n)
    invalidated_time = fuzzy.FuzzyDate(datetime.date(2022, 3, 20))
    jwt = factory.Faker('pystr')