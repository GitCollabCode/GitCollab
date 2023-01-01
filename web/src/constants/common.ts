/* eslint-disable  no-unused-vars */
export enum ModalType {
  LoggedOutModal,
  SkillSelectModal,
}

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
  modalType: ModalType
  setModalType: (type: ModalType) => void
  displayModal: boolean
  showModal: () => void
  hideModal: () => void
}
