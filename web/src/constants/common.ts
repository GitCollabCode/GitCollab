export type userValue = {
  isLoggedIn: boolean | null
  user: string | null
  client_id: string | undefined
  redirect_uri: string | undefined
  client_secret: string | undefined
  proxy_url: string | undefined
}

export type userDispatch = {
  logOut: () => void
  logIn: (user: string) => void
}

export type userLoginContextState = userValue & userDispatch
