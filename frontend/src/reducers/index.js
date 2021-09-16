import {applyMiddleware, combineReducers, createStore} from 'redux';
import {composeWithDevTools} from 'redux-devtools-extension';
import thunk from 'redux-thunk';
import userReducer from './userReducer';
import fileReducer from './fileReducer';
import loaderReducer from './loaderReducer';
import progressReducer from './progressReducer';
import optionsReducer from './optionsReducer';
import createButtonReducer from './createButtonReducer';
import downloadButtonReducer from './downloadButtonReducer';
import dimensionsReducer from './dimensionsReducer';
import imageSizeReducer from './imageSizeReducer';

const rootReducer = combineReducers({
  user: userReducer,
  files: fileReducer,
  loader: loaderReducer,
  progress: progressReducer,
  options: optionsReducer,
  createButton: createButtonReducer,
  downloadButton: downloadButtonReducer,
  dimensions: dimensionsReducer,
  imageSize: imageSizeReducer,
})

export const store = createStore(rootReducer, 
  composeWithDevTools(applyMiddleware(thunk)))
