import { userValue } from '../../constants/common'

const isLoggedLS = localStorage.getItem('isLoggedIn')
const user = localStorage.getItem('user')

export const initialState = {
  isLoggedIn: isLoggedLS ? JSON.parse(isLoggedLS) : false,
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
      }
    }
    case 'LOGOUT': {
      return {
        ...state,
        isLoggedIn: false,
      }
    }
    default:
      return state
  }
}
