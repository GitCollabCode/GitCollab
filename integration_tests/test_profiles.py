from db_fixtures import *
from models.blacklist import BlacklistFactory, BlacklistModel
import requests


class TestProfiles:
    
    def test_monkey(self, api_route):
        '''
        monkey
        '''
        