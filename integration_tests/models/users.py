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

