import React, { useContext }  from 'react'

import style from '../Modal.module.css'

import octocat from  "../../../assets/octocat.png"
import { Link } from 'react-router-dom'
import { ModalContextStateContext } from '../../../context/modalContext/modalContext'

const LoggedOutModal = () => {    
  const modalContext = useContext(ModalContextStateContext)
  return (
      <>
      <div className={style.modalText}>
          <img className={style.modalLogo} src={octocat} alt="i love chesburger"/>
          <p className={style.modalTextContent}>
              Sorry! That page does not exist
          </p>
          <Link to="/">
            <button className={style.modalButton} onClick={() => modalContext.hideModal()} >
                Take me back home
            </button>
          </Link>
      </div>
      </>
    ) 
}

export default LoggedOutModal