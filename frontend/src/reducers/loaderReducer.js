const SHOW_LOADER = 'SHOW_LOADER'
const HIDE_LOADER = 'HIDE_LOADER'
const CHANGE_LOADER = 'CHANGE_LOADER'

const defaultState = {
    isVisible: false,
    progress: 0
}

export default function loaderReducer(state = defaultState, action) {
    switch (action.type) {
        case SHOW_LOADER: return {...state, isVisible: true, progress: 0}
        case HIDE_LOADER: return {...state, isVisible: false, progress: 0}
        case CHANGE_LOADER: // TODO progress == 100 - isVisible: false
            return {
                ...state,
                progress: action.payload.progress // payload - ???
            }
        default:
            return state
    }
}


export const showLoader = () => ({type: SHOW_LOADER})
export const hideLoader = () => ({type: HIDE_LOADER})
export const changeLoader = (payload) => ({type: CHANGE_LOADER, payload: payload})
