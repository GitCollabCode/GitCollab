import { userValue } from '../../constants/common'

const isLoggedLS = localStorage.getItem('isLoggedIn')
const user = localStorage.getItem('user')

export const initialState = {
  isLoggedIn: isLoggedLS ? JSON.parse(isLoggedLS) : false,
  user: user ? JSON.parse(user) : null,
  client_id: process.env.REACT_APP_CLIENT_ID,
  redirect_uri: process.env.REACT_APP_REDIRECT_URI,
  client_secret: process.env.REACT_APP_CLIENT_SECRET,
  proxy_url: process.env.REACT_APP_PROXY_URL,
  jwtToken: '',
}

type LogIn = {
  type: 'LOGIN'
  payload: any
}

type LogOut = {
  type: 'LOGOUT'
}

export type ValidAction = LogIn | LogOut

export const reducer = (state: userValue, action: ValidAction): userValue => {
  switch (action.type) {
    case 'LOGIN': {
      return {
        ...state,
        isLoggedIn: action.payload.isLoggedIn,
        user: action.payload.user,
      }
    }
    case 'LOGOUT': {
      localStorage.clear()
      return {
        ...state,
        isLoggedIn: false,
        user: null,
      }
    }
    default:
      return state
  }
}
