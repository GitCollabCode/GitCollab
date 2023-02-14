/* eslint-disable  no-unused-vars */
export enum ModalType {
  LoggedOutModal,
  SkillSelectModal,
  LanguagesSelectModal,
  PageNotFoundModal,
  NewProjectModal,
  NewTaskModal,
}


export enum TaskProgress {
  TaskStatusUnassigned = "UNASSIGNED",
  TaskStatusAssigned = "ASSIGNED",
  TaskStatusChangesRequested = "CHANGES_REQUESTED",
  TaskStatusDismissed = "DISMISSED",
  TaskStatusApproved = "APPROVED",
  TaskStatusReadyToMerge = "READY_TO_MERGE",
  TaskStatusCompleted = "COMPLETED",
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

export type IssueResponse = {
  issues: IssueType[]
}

export type IssueType = {
  state:string,
  title:string,
  url:string
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
  project_name: string
  project_description: string
  project_owner: string
  project_skills: string[]
}

export type TaskCardType = {
  name: string
  description: string
  languages: string[]
  assignedTo: string
  assignedToImg: string
  progress: TaskProgress
}

export type SearchProjectResponse = {
  projects: ProjectCardType[]
}

export type ProfileProjectResponse = {
  projects: ProjectCardType[]
}


export type TaskType = {
  "task_id":number,
	"project_id":number,
	"project_name":string,
	"task_status":TaskProgress,
	"completed_by_id":number,
	"date_created_date":Date,
	"completed_date":Date,
	"task_title":string,
	"task_description":String,
	"diffictly":number,
	"priority":number,
	"skills":string[]
}

export type TaskResponse = TaskType[] 

