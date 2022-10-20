export type userValue = {
  isLoggedIn: boolean | null
  user: string | null
  client_id: string | undefined
  redirect_uri: string | undefined
  client_secret: string | undefined
  proxy_url: string | undefined
  jwtToken: string | undefined
}

export type userDispatch = {
  logOut: () => void
  logIn: (token: string, user: any) => void
  setToken: (token: string) => void
}

export type userLoginContextState = userValue & userDispatch
