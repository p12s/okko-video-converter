import {config} from '../config';

const CHANGE_IMAGE_SIZE = 'CHANGE_IMAGE_SIZE'
const DROP_IMAGE_SIZE = 'CHANGE_IS_ADD_WEBP'

const defaultState = {
    imageWidth: 0,
    imageHeight: 0
}

export default function imageSizeReducer(state = defaultState, action) {
    switch (action.type) {
        case DROP_IMAGE_SIZE: return {...state, imageWidth: 0, imageHeight: 0}
        case CHANGE_IMAGE_SIZE:
            let width = action.payload.imageWidth
            if (width > config.MAX_PIXEL_SIZE) {
                width = config.MAX_PIXEL_SIZE
            }
            
            let height = action.payload.imageHeight
            if (height > config.MAX_PIXEL_SIZE) {
                height = config.MAX_PIXEL_SIZE
            }

            return {
                ...state,
                imageWidth: width,
                imageHeight: height,
            }
        default:
            return state
    }
}

export const changeImageSize = (payload) => ({type: CHANGE_IMAGE_SIZE, payload: payload})
export const dropImageSize = () => ({type: DROP_IMAGE_SIZE})
