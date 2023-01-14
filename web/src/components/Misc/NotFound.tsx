import React, { useContext } from 'react'
import { ModalType } from '../../constants/common';
import { ModalContextStateContext } from '../../context/modalContext/modalContext';
const NotFound = () => {
  useContext(ModalContextStateContext).setModalType(ModalType.PageNotFoundModal);
  useContext(ModalContextStateContext).showModal();
  return(<>.</>)
}

export default NotFound
