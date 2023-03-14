import React, { ReactPortal } from 'react'
import { render, screen, waitFor } from '@testing-library/react'

import { ModalContextStateContext } from '../../../context/modalContext/modalContext'
import Modal from '../Modal'
import { ModalType } from '../../../constants/common'
import ReactDOM from 'react-dom'
import { debug } from 'console'

const mockedUsedNavigate = jest.fn()

jest.mock('react-router-dom', () => ({
  ...(jest.requireActual('react-router-dom') as any),
  useNavigate: () => mockedUsedNavigate,
}))

describe('Modal', () => {
  beforeAll(() => {
    ReactDOM.createPortal = jest.fn((element): ReactPortal => {
      return element as ReactPortal
    })
  })

  it('Test LoggedOut Modal render', () => {
    render(
      <ModalContextStateContext.Provider
        value={{
          modalType: ModalType.LoggedOutModal,
          displayModal: true,
          setModalType: jest.fn(),
          showModal: jest.fn(),
          hideModal: jest.fn(),
        }}
      >
        <div id="modal-root"></div>
        <Modal />
      </ModalContextStateContext.Provider>
    )
    waitFor(() =>
      expect(screen.getByText('Login with GitHub')).toBeInTheDocument()
    )
  })

  it('Test LanguageSelectModal render', () => {
    render(
      <ModalContextStateContext.Provider
        value={{
          modalType: ModalType.LanguagesSelectModal,
          setModalType: jest.fn(),
          displayModal: true,
          showModal: jest.fn(),
          hideModal: jest.fn(),
        }}
      >
        <div id="modal-root"></div>
        <Modal />
      </ModalContextStateContext.Provider>
    )
  })

  it('Test NewProjectModal render', () => {
    render(
      <ModalContextStateContext.Provider
        value={{
          modalType: ModalType.NewProjectModal,
          setModalType: jest.fn(),
          displayModal: true,
          showModal: jest.fn(),
          hideModal: jest.fn(),
        }}
      >
        <Modal />
      </ModalContextStateContext.Provider>
    )
  })

  it('Test PageNotFound render', () => {
    render(
      <ModalContextStateContext.Provider
        value={{
          modalType: ModalType.PageNotFoundModal,
          setModalType: jest.fn(),
          displayModal: true,
          showModal: jest.fn(),
          hideModal: jest.fn(),
        }}
      >
        <Modal />
      </ModalContextStateContext.Provider>
    )
    waitFor(() =>
      expect(screen.getByText('Take me back home')).toBeInTheDocument()
    )
  })
  it('Test SkillsSelectModal render', () => {
    render(
      <ModalContextStateContext.Provider
        value={{
          modalType: ModalType.SkillSelectModal,
          setModalType: jest.fn(),
          displayModal: true,
          showModal: jest.fn(),
          hideModal: jest.fn(),
        }}
      >
        <Modal />
      </ModalContextStateContext.Provider>
    )
  })
})
