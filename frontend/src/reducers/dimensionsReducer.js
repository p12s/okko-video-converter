const SHOW_DIMENSIONS = 'SHOW_DIMENSIONS'
const HIDE_DIMENSIONS = 'HIDE_DIMENSIONS'

const defaultState = {
    isVisible: false
}

export default function dimensionsReducer(state = defaultState, action) {
    switch (action.type) {
        case SHOW_DIMENSIONS: return {...state, isVisible: true}
        case HIDE_DIMENSIONS: return {...state, isVisible: false}
        default:
            return state
    }
}

export const showDimensions = () => ({type: SHOW_DIMENSIONS})
export const hideDimensions = () => ({type: HIDE_DIMENSIONS})
