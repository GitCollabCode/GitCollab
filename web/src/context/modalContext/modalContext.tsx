import React, { createContext, ReactNode, useState } from 'react'

import { modalContextState, ModalType } from '../../constants/common'

const initialState: modalContextState = {
  modalType: ModalType.LoggedOutModal,
  setModalType: (type: ModalType) => {
    /** */
  },
  displayModal: false,
  showModal: () => {
    /** */
  },
  hideModal: () => {
    /** */
  },
}

export const ModalContextStateContext =
  createContext<modalContextState>(initialState)

export const ModalContextStateProvider = ({
  children,
}: {
  children: ReactNode
}) => {
  const [displayModal, setdisplayModal] = useState(false)
  const [modalType, setModalType] = useState<ModalType>(
    ModalType.LoggedOutModal
  )

  const showModal = () => {
    setdisplayModal(true)
  }

  const hideModal = () => {
    setdisplayModal(false)
  }

  return (
    <ModalContextStateContext.Provider
      value={{
        modalType: modalType,
        setModalType,
        displayModal,
        showModal,
        hideModal,
      }}
    >
      {children}
    </ModalContextStateContext.Provider>
  )
}
