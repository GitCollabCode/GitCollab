import React, { createContext, useReducer, ReactNode, useEffect } from 'react'

import { userLoginContextState } from '../../constants/common'
import { LOGOUT } from '../../constants/endpoints'
import { reducer } from './userLoginValues'

const isLoggedLS = localStorage.getItem('isLoggedIn')

const initialState = {
  isLoggedIn: isLoggedLS ? JSON.parse(isLoggedLS) : false,
  jwtToken: '',
  logOut: () => {
    /** */
  },
  logIn: (token: string) => {
    /** */
  },
  setToken: (token: string) => {
    /** */
  },
}

export const UserLoginContext =
  createContext<userLoginContextState>(initialState)

export function UserLoginProvider({ children }: { children: ReactNode }) {
  const [state, dispatch] = useReducer(reducer, initialState)

  const logOut = () => {
    const token = localStorage.getItem('gitcollab_jwt')
    fetch(process.env.REACT_APP_API_URI + LOGOUT, {
      method: 'GET',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
        Authorization: 'Bearer ' + token,
      },
    }).then(() => localStorage.removeItem('gitcollab_jwt'))

    dispatch({
      type: 'LOGOUT',
    })
  }

  useEffect(() => {
    const getToken = localStorage.getItem('gitcollab_jwt') || ''
    if (getToken !== '') {
      dispatch({
        type: 'LOGIN',
        payload: { isLoggedIn: true },
      })
    }
  }, [])

  const setToken = (token: string) => {
    localStorage.setItem('gitcollab_jwt', token)
  }

  const logIn = (token: string) => {
    setToken(token)
    dispatch({
      type: 'LOGIN',
      payload: { isLoggedIn: true },
    })
  }

  return (
    <UserLoginContext.Provider
      value={{
        isLoggedIn: state.isLoggedIn,
        jwtToken: state.jwtToken,
        logOut,
        logIn,
        setToken,
      }}
    >
      {children}
    </UserLoginContext.Provider>
  )
}
