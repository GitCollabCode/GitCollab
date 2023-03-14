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
  projectId: 0,
  setProjectId: (val:number) => {
     /** */
  }

}

export const ModalContextStateContext =
  createContext<modalContextState>(initialState)

export const ModalContextStateProvider = ({
  children,
}: {
  children: ReactNode
}) => {
  const [displayModal, setdisplayModal] = useState(false)
  const [projectId, setProjectId] = useState(0)
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
        projectId, 
        setProjectId
      }}
    >
      {children}
    </ModalContextStateContext.Provider>
  )
}
