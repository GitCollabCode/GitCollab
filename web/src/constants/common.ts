import { ReactNode } from 'react'

export type userValue = {
  isLoggedIn: boolean | null
  jwtToken: string | undefined
}

export type userDispatch = {
  logOut: () => void
  logIn: (token: string) => void
  setToken: (token: string) => void
}

export type loginResponse = {
  Token: string
  NewUser: boolean
}

export type GitHubRedirectResponse = {
  RedirectUrl: string
}

export type userLoginContextState = userValue & userDispatch

export type modalContextState = {
  modalContents: ReactNode
  setModalContents: (content: ReactNode) => void
  displayModal: boolean
  showModal: () => void
  hideModal: () => void
}
