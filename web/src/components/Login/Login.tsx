import React, { useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { GITHUB_REDIRECT } from '../../constants/endpoints'
import { UserLoginContext } from '../../context/userLoginContext/userLoginContext'
import styles from './Login.module.css'

const Login = () => {
  const { proxy_url, logIn, user, isLoggedIn, logOut } =
    useContext(UserLoginContext)
  const [data, setData] = useState({ errorMessage: '', isLoading: false })
  const navigate = useNavigate()

  useEffect(() => {
    // After requesting Github access, Github redirects back to your app with a code parameter
    const url = window.location.href
    const hasCode = url.includes('?code=')
    console.log(process.env.REACT_APP_API_URI + GITHUB_REDIRECT);
    const newUri = process.env.REACT_APP_API_URI + GITHUB_REDIRECT

    // If Github API returns the code parameter
    if (hasCode) {
      const newUrl = url.split('?code=')
      window.history.pushState({}, '', newUrl[0])
      setData({ ...data, isLoading: true })

      const requestData = {
        code: newUrl[1],
      }

      type jwtToken = {
        token: string
      }

      // Use code parameter and other parameters to make POST request to BE
      fetch(newUri, {
        method: 'POST',
        body: JSON.stringify(requestData),
      })
        .then((response) => response.json())
        .then((data: jwtToken) => {
          console.log(data)
          logIn(data.token, {})
        })
        .catch((error) => {
          console.log(error)
          setData({
            isLoading: false,
            errorMessage: 'Sorry! Login failed',
          })
        })
    }
  }, [isLoggedIn, data, logIn, proxy_url])

  if (isLoggedIn) {
    navigate('/')
  }

  type githubRedirectURL = {
    redirect: string
  }
  const redirectToGithub = () => {
    fetch(process.env.REACT_APP_API_URI + GITHUB_REDIRECT, {
      method: 'GET',
    })
      .then((response) => response.json())
      .then((data: githubRedirectURL) => {
        window.location.replace(data.redirect)
      })
      .catch((error) => {
        console.log(error)
      })
  }
  return (
    <div className={styles.container}>
      <h3>Login</h3>
      <div className="login-container">
        {data.isLoading ? (
          <div className="loader-container">
            <div className="loader"></div>
          </div>
        ) : (
          <>
            {console.log(user)}
            {!isLoggedIn ? (
              <button
                className={'btn btn-primary'}
                onClick={() => redirectToGithub()}
              >
                <i className="fa fa-trophy"></i> Login with GitHub
              </button>
            ) : (
              <button className={'btn btn-primary'} onClick={() => logOut()}>
                <i className="fa fa-trophy"></i> Logout
              </button>
            )}
          </>
        )}
      </div>
    </div>
  )
}

export default Login
