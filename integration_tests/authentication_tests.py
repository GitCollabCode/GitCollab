from db_fixtures import *
from models.users import Profiles
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
        # check if blocked, make sure not a 404 error
        assert resp.status_code != 200 and resp.status_code != 404

    def test_blacklist_invalid_jwt(self, api_route):
        '''
        Ensure request with invalid jwt is blocked by blacklist
        '''
        url = api_route + "test/test-blacklist"
        resp = requests.get(url, headers={"Authorization":"Bearer ILoveCrackCocaine"})
        # check if blocked, make sure not a 404 error
        assert resp.status_code != 200 and resp.status_code != 404

    def test_login_db(self, session):
        '''
        Ensure 
        '''
        result = session.query(Profiles).one_or_none()
        assert result is not None
