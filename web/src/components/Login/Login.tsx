import React, { useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { UserLoginContext } from '../../context/userLoginContext/userLoginContext'
import styles from './Login.module.css'

const Login = () => {
  const { proxy_url, logIn, user, isLoggedIn, client_id, redirect_uri } =
    useContext(UserLoginContext)
  const [data, setData] = useState({ errorMessage: '', isLoading: false })
  const navigate = useNavigate()

  useEffect(() => {
    // After requesting Github access, Github redirects back to your app with a code parameter
    const url = window.location.href
    const hasCode = url.includes('?code=')
    const newUri = "http://localhost:8080/auth/signin/" || ''

    // If Github API returns the code parameter
    if (hasCode) {
      const newUrl = url.split('?code=')
      window.history.pushState({}, '', newUrl[0])
      setData({ ...data, isLoading: true })

      const requestData = {
        code: newUrl[1],
      }

      console.log(requestData)

      // Use code parameter and other parameters to make POST request to BE
      fetch(newUri, {
        method: 'POST',
        body: JSON.stringify(requestData),
      })
        .then((response) => response.json())
        .then((data) => {
          logIn(data)
        })
        .catch((error) => {
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
            
              <button className={"btn btn-primary"}>
              
              <i className="fa fa-trophy"></i> | Signin with GitHub
              </button>
              <a
              
              href={`https://github.com/login/oauth/authorize?scope=user&client_id=${client_id}&redirect_uri=${redirect_uri}`}
              onClick={() => {
                setData({ ...data, errorMessage: '' })
              }}
            >
            </a>
          </>
        )}
      </div>
    </div>
  )
}

export default Login
