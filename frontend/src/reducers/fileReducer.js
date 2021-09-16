
const SET_FILES = "SET_FILES"
const ADD_FILE = "ADD_FILE"
const REPLACE_FILE = "REPLACE_FILE"
const REMOVE_FILE = "REMOVE_FILE"
const SET_POPUP_DISPLAY = "SET_POPUP_DISPLAY"

const defaultState = {
  files: [],
  popupDisplay: 'none',
}

export default function fileReducer(state = defaultState, action) {
  switch (action.type) {
    case SET_FILES: return {...state, files: action.payload}
    case ADD_FILE: return {...state, files: [...state.files, action.payload]}
    case REPLACE_FILE: return {...state, files: [action.payload]}
    case REMOVE_FILE: return {...state, files: []}
    case SET_POPUP_DISPLAY: return {...state, popupDisplay: action.payload}
    default:
      return state
  }
}

export const setFiles = (files) => ({type: SET_FILES, payload: files})
export const addFile = (file) => ({type: ADD_FILE, payload: file})
export const replaceFile = (file) => ({type: REPLACE_FILE, payload: file})
export const removeFile = () => ({type: REMOVE_FILE})
export const setPopupDisplay = (display) => ({type: SET_POPUP_DISPLAY, payload: display})
