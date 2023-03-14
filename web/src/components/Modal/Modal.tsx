import React, { useContext } from 'react'
import ReactDOM from 'react-dom'
import { useNavigate } from 'react-router-dom'
import style from '../Modal/Modal.module.css'
import { ModalContextStateContext } from '../../context/modalContext/modalContext'
import { ModalType } from '../../constants/common'
import LoggedOutModal from './ModalTypes/LoggedOutModal'
import LoginModal from './ModalTypes/SkillSelectModal'
import LanguagesSelectModal from './ModalTypes/LanguagesSelectModal'
import PageNotFoundModal from './ModalTypes/PageNotFoundModal'
import NewProjectModal from './ModalTypes/NewProjectModal'

const Modal = () => {
  const { modalType, displayModal, hideModal /*showModal*/ } = useContext(
    ModalContextStateContext
  )
  const navigate = useNavigate()

  const onBackgroundClick = () => {
    navigate('/')
    hideModal()
  }

  const renderModalType = (modalType: any) => {
    switch (modalType) {
      case ModalType.LoggedOutModal:
        return <LoggedOutModal />
      case ModalType.SkillSelectModal:
        return <LoginModal />
      case ModalType.LanguagesSelectModal:
        return <LanguagesSelectModal />
      case ModalType.PageNotFoundModal:
        return <PageNotFoundModal />
      case ModalType.NewProjectModal:
        return <NewProjectModal />
    }
  }
  return displayModal ? (
    ReactDOM.createPortal(
      <>
        <div className={style.modalContainer}>
          <div className={style.bg} onClick={onBackgroundClick}></div>
          <div className={style.overlayCard}>{renderModalType(modalType)}</div>
        </div>
        ,
      </>,
      document.getElementById('modal-root') as Element
    )
  ) : (
    <></>
  )
}

export default Modal
