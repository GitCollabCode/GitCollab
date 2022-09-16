import React, { createContext, useReducer, ReactNode } from 'react'

import { userLoginContextState } from '../../constants/common'
import { reducer } from './userLoginValues'

const isLoggedLS = localStorage.getItem('isLoggedIn')
const user = localStorage.getItem('user')

const initialState = {
  isLoggedIn: isLoggedLS ? JSON.parse(isLoggedLS) : false,
  user: user ? JSON.parse(user) : null,
  client_id: process.env.REACT_APP_CLIENT_ID,
  redirect_uri: process.env.REACT_APP_REDIRECT_URI,
  client_secret: process.env.REACT_APP_CLIENT_SECRET,
  proxy_url: process.env.REACT_APP_PROXY_URL,
  logOut: () => {
    /** */
  },
  logIn: (user: string) => {
    /** */
  },
}

export const UserLoginContext =
  createContext<userLoginContextState>(initialState)

export function UserLoginProvider({ children }: { children: ReactNode }) {
  const [state, dispatch] = useReducer(reducer, initialState)

  const logOut = () =>
    dispatch({
      type: 'LOGOUT',
    })

  const logIn = (user: string) => {
    dispatch({
      type: 'LOGIN',
      payload: { user: user, isLoggedIn: true },
    })
  }

  return (
    <UserLoginContext.Provider
      value={{
        isLoggedIn: state.isLoggedIn,
        user: state.user,
        client_id: state.client_id,
        redirect_uri: state.redirect_uri,
        client_secret: state.client_secret,
        proxy_url: state.proxy_url,
        logOut,
        logIn,
      }}
    >
      {children}
    </UserLoginContext.Provider>
  )
}
