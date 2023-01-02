import React, { useContext, useEffect, useState }  from 'react'

import style from '../Modal.module.css'
import { GitHubRedirectResponse, loginResponse } from '../../../constants/common'
import { GITHUB_REDIRECT, SIGNIN} from '../../../constants/endpoints'

import octocat from  "../../../assets/octocat.png"
import { UserLoginContext } from '../../../context/userLoginContext/userLoginContext'
import { ModalContextStateContext } from '../../../context/modalContext/modalContext'

const LoggedOutModal = () => {    
  const modalContext = useContext(ModalContextStateContext)
  const { logIn, isLoggedIn } = useContext(UserLoginContext);
  const [data, setData] = useState({ errorMessage: '', isLoading: false })

  useEffect(() => {
    // After requesting Github access, Github redirects back to your app with a code parameter
    const url = window.location.href
    const hasCode = url.includes('?code=')

    // If Github API returns the code parameter
    if (hasCode) {
      
      const newUrl = url.split('?code=')
      window.history.pushState({}, '', newUrl[0])

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
        // after login close modal
        modalContext.hideModal()
    }
  }, [isLoggedIn, data, logIn])


  const redirectToGithub = () => {
      fetch(process.env.REACT_APP_API_URI + GITHUB_REDIRECT, {
        method: 'GET',
      })
        .then((response) => response.json())
        .then((data: GitHubRedirectResponse) => { 
          window.location.replace(data.RedirectUrl)
        })
        .catch((error) => {
          console.log(error)
        })
    }

  
  return (
      <>
      <div className={style.modalText}>
          <img className={style.modalLogo} src={octocat}/>
          <p className={style.modalTextContent}>
              Sorry! You’re currently logged out, sign in to continue
          </p>
          <button className={style.modalButton} onClick={redirectToGithub}>
              Login with GitHub
          </button>
      </div>
      </>
    ) 
}

export default LoggedOutModal