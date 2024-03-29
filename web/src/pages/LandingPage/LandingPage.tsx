import React, { useState } from 'react'
import LoginButton from '../../components/LoginButton/LoginButton'
import style from './LandingPage.module.css'
import Media from 'react-media'
import LoadingSpinner from '../../components/LoadingSpinner/LoadingSpinner'

const LandingPage = () => {
  const [isLoading, setIsLoading] = useState(false)
  return (
    <>
      <div className={style.container}>
        <Media query={{ maxWidth: 1023 }}>
          <div className={style.login}>
            <div className={style.logo}>GitCollab</div>
            <LoginButton setIsLoading={setIsLoading} />
            {isLoading && <LoadingSpinner isLoading={isLoading} />}
          </div>
        </Media>
        <div className={style.textContainer}>
          <div className={style.topText}>Coding is</div>
          <div className={style.animatedText}>
            <div className={style.line}>cooperative</div>
            <div className={style.line}>innovative</div>
            <div className={style.line}>collaborative</div>
            <div className={style.line}>interactive</div>
            <div className={style.line}>productive</div>
          </div>
        </div>
      </div>
    </>
  )
}

export default LandingPage
