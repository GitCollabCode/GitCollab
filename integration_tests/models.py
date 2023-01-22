from sqlalchemy import Column, Integer, VARCHAR, TIMESTAMP
from sqlalchemy.orm import declarative_base

Base = declarative_base()

class Blacklist(Base):
    __tablename__ = "blacklist"
    invalidated_time = Column(TIMESTAMP())
    jwt = Column(VARCHAR(), primary_key=True)
