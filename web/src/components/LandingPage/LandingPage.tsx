import React from 'react'
import LoginButton from '../LoginButton/LoginButton'
import style from './LandingPage.module.css'
import Media from 'react-media'

const LandingPage = () => {
  return (
    <>
      <div className={style.container}>
        <Media query={{ maxWidth: 1023 }}>
          <div className={style.login}>
            <div className={style.logo}>GitCollab</div>
            <LoginButton />
          </div>
        </Media>
        <div className={style.textContainer}>
          <div className={style.topText}>Coding is</div>
          <div className={style.animatedText}>
            <div className={style.line}>wagwan</div>
            <div className={style.line}>cringe</div>
            <div className={style.line}>collaborative</div>
            <div className={style.line}>in your walls</div>
            <div className={style.line}>maged sus</div>
          </div>
        </div>
      </div>
    </>
  )
}

export default LandingPage
