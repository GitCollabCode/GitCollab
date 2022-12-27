from db_fixtures import *
from models.blacklist import BlacklistFactory, BlacklistModel
import requests



class TestAuthentication:
    
    def test_get_login_redirect(self, api_route):
        '''
        Ensure redirect url is created and returned
        '''
        url = api_route + "auth/redirect-url"
        resp = requests.get(url)
        assert "RedirectUrl" in resp.json()
        assert resp.json()["RedirectUrl"] is not None

    def test_blacklist_no_jwt(self, api_route):
        '''
        Ensure request with no jwt is blocked by blacklist
        '''
        url = api_route + "test/test-blacklist"
        resp = requests.get(url)
        assert resp.status_code == 401

    def test_blacklist_valid_jwt(self, api_route):
        '''
        Check if request served with jwt not in blacklist
        '''
        url = api_route + "test/test-blacklist"
        resp = requests.get(url, headers={"Authorization":"Bearer CringeBurgerKing"})
        assert resp.status_code == 200

    def test_blacklist_blocked_jwt(self, api_route, session):
        '''
        Check if request blocked when jwt in blacklist
        '''
        url = api_route + "test/test-blacklist"
        session.add(BlacklistFactory(jwt="abc"))
        session.commit()
        #connection.execute("INSERT INTO jwt_blacklist VALUES (1234, 'abc')")

        resp = requests.get(url, headers={"Authorization":"Bearer abc"})
        assert resp.status_code == 401

 #   def test_login_db(self, session):
 #       '''
 #       Ensure s
 #       '''
 #       result = session.query(Profiles).one_or_none()
 #       assert result is not None
