import React, { useContext, useEffect, useState } from 'react'
import { GITHUB_REDIRECT, SIGNIN } from '../../constants/endpoints'
import { UserLoginContext } from '../../context/userLoginContext/userLoginContext'
import { loginResponse } from '../../constants/common'
import { ReactComponent as GithubIcon } from '../../assets/github.svg'
import { ReactComponent as LogoutIcon } from '../../assets/logout.svg'
import style from './LoginButton.module.css'

const LoginButton = () => {
  const { logIn, isLoggedIn, logOut } = useContext(UserLoginContext)
  const [data, setData] = useState({ errorMessage: '', isLoading: false })

  useEffect(() => {
    // After requesting Github access, Github redirects back to your app with a code parameter
    const url = window.location.href
    const hasCode = url.includes('?code=')

    // If Github API returns the code parameter
    if (hasCode) {
      const newUrl = url.split('?code=')
      window.history.pushState({}, '', newUrl[0])
      setData({ ...data, isLoading: true })

      const requestData = {
        code: newUrl[1],
      }

      // Use code parameter and other parameters to make POST request to BE
      fetch(process.env.REACT_APP_API_URI + SIGNIN, {
        method: 'POST',
        body: JSON.stringify(requestData),
      })
        .then((response) => response.json())
        .then((data: loginResponse) => {
          if (data.NewUser) {
            //TODO ADD the Modal for getting user stuff
            console.log('New USER ')
          }
          logIn(data.Token)
        })
        .catch((error) => {
          console.log(error)
          setData({
            isLoading: false,
            errorMessage: 'Sorry! Login failed',
          })
        })
    }
  }, [isLoggedIn, data, logIn])

  const redirectToGithub = () => {
    fetch(process.env.REACT_APP_API_URI + GITHUB_REDIRECT, {
      method: 'GET',
    })
      .then((response) => response.text())
      .then((data) => {
        window.location.replace(data)
      })
      .catch((error) => {
        console.log(error)
      })
  }

  return (
    <>
      {!isLoggedIn ? (
        <button className={style.button} onClick={() => redirectToGithub()}>
          <GithubIcon /> <p className={style.githubButtonText}>login</p>
        </button>
      ) : (
        <button className={style.button} onClick={() => logOut()}>
          <LogoutIcon /> <p className={style.githubButtonText}>logout</p>
        </button>
      )}
    </>
  )
}

export default LoginButton
