import React from 'react'
import style from './LandingPage.module.css'

const LandingPage = () => {
  return (
    <>
      <div className={style.container}>
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
