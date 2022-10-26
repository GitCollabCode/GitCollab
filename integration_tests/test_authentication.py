import os
import requests
from pytest_postgresql import factories

ENDPOINT = "http://localhost:8080/auth/"

DB_HOST = "host.docker.internal"
DB_NAME = os.getenv("POSTGRES_NAME")
DB_USER = os.getenv("POSTGRES_USER")
DB_PASS = os.getenv("POSTGRES_PASSWORD")

# session scope fixture for live postgres in docker, pass postgres to all tests
postgresql_in_docker = factories.postgresql("postgresql_noproc", dbname="postgres")
postgresql = factories.postgresql("postgresql_in_docker")


def test_get_redirect_url():
    url = ENDPOINT + "redirect-url"
    req = requests.get(url)
    assert req.status_code == 200
    assert req != ""
        
def test_logout():
    # expired jwt
    TEST_JWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOiIyMDIyLTEwLTI3VDAxOjAzOjA0LjU1MDY4MzI5OFoiLCJnaXRodWJJRCI6Mjk1NTE5NzQsInVzZXIiOiJyb2JvdGV2YW4ifQ.M5XB5s8o3HFEewh-5ae6_jJpMpIuXhc3jfkNsJnKYvA"

    url = ENDPOINT + "logout"
    req = requests.get(url, headers={"Authorization": "Bearer " + TEST_JWT})
    assert req.status_code == 200
    #with postgresql.cursor() as cur:
    #    cur.execute("select jwt from jwt_blacklist where jwt=" + "\"" + TEST_JWT + "\"")
    #    print(cur.fetchall())