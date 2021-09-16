const SET_USER = "SET_USER"

const defaultState = {
  token: null
}

export default function userReducer(state = defaultState, action) {
  switch (action.type) {
    case SET_USER:
      return {
        ...state,
        token: action.payload
      }
    default:
      return state
  }
}

export const setUser = (user) => ({type: SET_USER, payload: user})