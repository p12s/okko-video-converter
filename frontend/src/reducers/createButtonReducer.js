const SHOW_BUTTON = 'SHOW_BUTTON'
const HIDE_BUTTON = 'HIDE_BUTTON'

const defaultState = {
    isVisible: true
}

export default function createButtonReducer(state = defaultState, action) {
    switch (action.type) {
        case SHOW_BUTTON: return {...state, isVisible: true}
        case HIDE_BUTTON: return {...state, isVisible: false}
        default:
            return state
    }
}

export const showCreateButton = () => ({type: SHOW_BUTTON})
export const hideCreateButton = () => ({type: HIDE_BUTTON})
