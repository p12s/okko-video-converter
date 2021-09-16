"use strict";

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.store = void 0;

var _redux = require("redux");

var _reduxDevtoolsExtension = require("redux-devtools-extension");

var _reduxThunk = _interopRequireDefault(require("redux-thunk"));

var _userReducer = _interopRequireDefault(require("./userReducer"));

var _fileReducer = _interopRequireDefault(require("./fileReducer"));

var _loaderReducer = _interopRequireDefault(require("./loaderReducer"));

var _progressReducer = _interopRequireDefault(require("./progressReducer"));

var _optionsReducer = _interopRequireDefault(require("./optionsReducer"));

var _createButtonReducer = _interopRequireDefault(require("./createButtonReducer"));

var _downloadButtonReducer = _interopRequireDefault(require("./downloadButtonReducer"));

var _dimensionsReducer = _interopRequireDefault(require("./dimensionsReducer"));

var _imageSizeReducer = _interopRequireDefault(require("./imageSizeReducer"));

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { "default": obj }; }

var rootReducer = (0, _redux.combineReducers)({
  user: _userReducer["default"],
  files: _fileReducer["default"],
  loader: _loaderReducer["default"],
  progress: _progressReducer["default"],
  options: _optionsReducer["default"],
  createButton: _createButtonReducer["default"],
  downloadButton: _downloadButtonReducer["default"],
  dimensions: _dimensionsReducer["default"],
  imageSize: _imageSizeReducer["default"]
});
var store = (0, _redux.createStore)(rootReducer, (0, _reduxDevtoolsExtension.composeWithDevTools)((0, _redux.applyMiddleware)(_reduxThunk["default"])));
exports.store = store;