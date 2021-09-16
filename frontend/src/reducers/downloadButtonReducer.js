const SHOW_DOWNLOAD = 'SHOW_DOWNLOAD'
const HIDE_DOWNLOAD = 'HIDE_DOWNLOAD'

const defaultState = {
    isVisible: false
}

export default function downloadButtonReducer(state = defaultState, action) {
    switch (action.type) {
        case SHOW_DOWNLOAD: return {...state, isVisible: true}
        case HIDE_DOWNLOAD: return {...state, isVisible: false}
        default:
            return state
    }
}

export const showDownloadButton = () => ({type: SHOW_DOWNLOAD})
export const hideDownloadButton = () => ({type: HIDE_DOWNLOAD})
