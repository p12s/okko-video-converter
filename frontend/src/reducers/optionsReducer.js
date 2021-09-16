const CHANGE_OPTIONS = 'CHANGE_OPTIONS'
// TODO стоит переделать на useState - и прокидвать через useContext

const defaultState = {
    widthList: [],
    coefList: [],
}

export default function optionsReducer(state = defaultState, action) {

    switch (action.type) {
        case CHANGE_OPTIONS:
            return {
                ...state,
                widthList: action.payload.widthList,
                coefList: action.payload.coefList,
            }
        default:
            return state
    }
}

export const changeOptions = (payload) => ({type: CHANGE_OPTIONS, payload: payload})
