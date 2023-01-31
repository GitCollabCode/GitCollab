/* eslint-disable  no-unused-vars */
export enum ModalType {
  LoggedOutModal,
  SkillSelectModal,
  LanguagesSelectModal,
  PageNotFoundModal,
  NewProjectModal,
}

export enum IssueProgress {
  UnAssigned,
  Assigned,
  InProgress,
  InReview,
  NeedsChanges,
  ReadyToMerge,
  Done,
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

export type profileResponse = {
  username: string
  gitID: number
  email: string
  avatarUrl: string
  languages: string[]
  skills: string[]
  bio: string
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

export type SkillListResponse = {
  skills: string[]
}

export type LanguageListResponse = {
  languages: string[]
}

export type ReposResponse = {
  repos: string[]
}

export type SelectType = {
  value: string
  label: string
}

export type ContributerType = {
  name: string
  url: string
}

export type ProjectCardType = {
  name: string
  description: string
  languages: string[]
  url: string
}

export type IssueCardType = {
  name: string
  description: string
  languages: string[]
  assignedTo: string
  assignedToImg: string
  progress: IssueProgress
}
