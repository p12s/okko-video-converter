const SHOW_PROGRESS = 'SHOW_PROGRESS'
const HIDE_PROGRESS = 'HIDE_PROGRESS'
const CHANGE_PROGRESS = 'CHANGE_PROGRESS'

const defaultState = {
    isVisible: false,
    progress: 0
}

export default function progressReducer(state = defaultState, action) {
    switch (action.type) {
        case SHOW_PROGRESS: return {...state, isVisible: true, progress: 0}
        case HIDE_PROGRESS: return {...state, isVisible: false, progress: 0}
        case CHANGE_PROGRESS: // TODO progress == 100 - isVisible: false
            let progres = action.payload.progress
            if (progres > 100) {
                progres = 100
            }
            return {
                ...state,
                progress: progres
            }
        default:
            return state
    }
}

export const showProgress = () => ({type: SHOW_PROGRESS})
export const hideProgress = () => ({type: HIDE_PROGRESS})
export const changeProgress = (payload) => ({type: CHANGE_PROGRESS, payload: payload})
