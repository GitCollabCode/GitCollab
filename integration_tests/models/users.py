from sqlalchemy import Column, Integer, String
import factory
from sqlalchemy.orm import declarative_base

Base = declarative_base()

class Profiles(Base):
    __tablename__ = "profiles"

    github_user_id = Column(Integer, primary_key=True)
    github_token = Column(String)
    username = Column(String)
    email = Column(String)
    avatar_url = Column(String)


#class Profiles(Base):
#    github_user_id = factory.sequence(lambda n: '%s' % n)
#    github_token = factory.Faker('github_token')
#    username = factory.Faker('username')
#    email = factory.Faker('email')
#    avatar_url = factory.Faker('avatar_url')
