import React, { createContext, useReducer, ReactNode, useEffect } from 'react'

import { userLoginContextState } from '../../constants/common'
import { LOGOUT } from '../../constants/endpoints'
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
  jwtToken: '',
  logOut: () => {
    /** */
  },
  logIn: (token: string, user: any) => {
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
    localStorage.removeItem('gitcollab_jwt')

    fetch(process.env.REACT_APP_API_URI+LOGOUT, {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + token,
      }
    }).then((response)=>console.log(response)).then(()=>localStorage.removeItem("gitcollab_jwt"))
    
    dispatch({
      type: 'LOGOUT',
    })
  }

  useEffect(()=>{
    const getToken = localStorage.getItem('gitcollab_jwt') || "";
    if(getToken !== "") {
      dispatch({
        type:'LOGIN',
        payload:{user:{}, isLoggedIn:true},
      });
    }
  },[])
  
  const setToken = (token: string) => {
    localStorage.setItem('gitcollab_jwt', token)
  }

  const logIn = (token: string, user: any) => {
    setToken(token)
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
