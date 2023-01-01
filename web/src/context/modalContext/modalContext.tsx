import React, { createContext, ReactNode, useState } from 'react'

import { modalContextState } from '../../constants/common'

const initialState: modalContextState = {
  modalContents: <></>,
  setModalContents: () => {
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
  const [modalStuff, setModalStuff] = useState<ReactNode>()

  const showModal = () => {
    setdisplayModal(true)
  }

  const hideModal = () => {
    setdisplayModal(false)
  }

  const setModalContents = (reactNode: ReactNode) => {
    setModalStuff(reactNode)
  }

  return (
    <ModalContextStateContext.Provider
      value={{
        modalContents: modalStuff,
        setModalContents,
        displayModal,
        showModal,
        hideModal,
      }}
    >
      {children}
    </ModalContextStateContext.Provider>
  )
}
