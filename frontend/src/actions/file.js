import axios from 'axios';
import { setFiles } from "../reducers/fileReducer";
import { config } from '../config';
import { registerUser } from './user';

import { showLoader, hideLoader, changeLoader } from "../reducers/loaderReducer";
import { showProgress, hideProgress } from "../reducers/progressReducer";
import { replaceFile, removeFile } from "../reducers/fileReducer"; // TODO когда не нужно будет перезаписывать, а просто добавлять - addFile
import { showCreateButton } from '../reducers/createButtonReducer';
import { showDownloadButton } from '../reducers/downloadButtonReducer';
import { changeImageSize } from '../reducers/imageSizeReducer';

export const getFiles = () => {
  return async dispatch => {
    try {
      const token = localStorage.getItem('token')

      if (token !== null) {
        const response = await axios.get(`${config.API_URL}/api/v1/files`, {
          headers: {Authorization: `Bearer ${localStorage.getItem('token')}`}
        })
        
        if (response.status === 200) {
          if (response.data && response.data.length > 0) {
            dispatch(showCreateButton())
            dispatch(changeImageSize({
              imageWidth: response.data[0].width,
              imageHeight: response.data[0].height,
            }))
          }
          dispatch(setFiles(response.data))
        }
      }

    } catch (e) {
      console.log('возникла ошибка при запросе файлов, недействительный токен либо проблема с сервером')
      console.log(e)
    }
  }
}

function connectToWebsocket(dispatch) {
  let tokenStr = localStorage.getItem('token')
  if (tokenStr === null) {
    console.log('Cannot sent ws request - token not exist')
    return
  }

  var ws = new WebSocket(`${config.WS_URL}/api/ws`);
  
  ws.onopen = function (event) {
    ws.send(JSON.stringify({ token: tokenStr }))
  };
  
  ws.onmessage = function (event) {
    let json = JSON.parse(event.data)
    if(json.hasOwnProperty('status') && json['status'] == 2){
      dispatch(showDownloadButton())
      dispatch(replaceFile(json))
    } else {
      console.log('Error from websocket')
    }
    console.log(json)
  };
  
  ws.onclose = function (event) {
    dispatch(hideProgress())
  };
}

// Загрузка файла на сервер - отрабатывает на каждый файл
export const uploadFile = (file, extension) => {
  return async dispatch => {
      try {
          let token = localStorage.getItem('token')
          if (token === null) {
            registerUser()
          }

          const formData = new FormData()
          formData.append('files', file)
          formData.append('extension', extension)

          const loaderProgress = {progress: 0}
          dispatch(showLoader())
          dispatch(changeLoader(loaderProgress))
          dispatch(removeFile())

          const response = await axios.post(`${config.API_URL}/api/v1/upload`, formData, {
            headers: {Authorization: `Bearer ${localStorage.getItem('token')}`},
            onUploadProgress: progressEvent => {
              const totalLength = progressEvent.lengthComputable ? progressEvent.total : progressEvent.target.getResponseHeader('content-length') || progressEvent.target.getResponseHeader('x-decompressed-content-length');
              if (totalLength) {
                loaderProgress.progress = Math.round((progressEvent.loaded * 100) / totalLength)
                // здесь приходят % загрузки файла на сервер
                dispatch(changeLoader(loaderProgress))
              }
            }
          })

          if (response.status !== 200) {
            console.log('возникла ошибка при загрузке файла:')
            console.log(response)
            // TODO вывод текста ошибки в интерфейсе
          } else {
            dispatch(showProgress())
            connectToWebsocket(dispatch)
          }
          dispatch(hideLoader())
      } catch (e) {
        alert(e)
      }
  }
}

// Выгрузка архива картинок с сервера
export async function downloadFile() {
  const response = await fetch(`${config.API_URL}/api/v1/download`,{
    headers: {
      Authorization: `Bearer ${localStorage.getItem('token')}`
    }
  })
  if (response.status === 200) {
      const blob = await response.blob()
      const downloadUrl = window.URL.createObjectURL(blob)
      const link = document.createElement('a')
      link.href = downloadUrl
      link.download = 'result.zip'
      document.body.appendChild(link)
      link.click()
      link.remove()

  } else {
    console.log('возникла ошибка при получении архива файлов')
    alert(response)
    // TODO вывод текста ошибки в интерфейсе
  }
}

export const removeAll = () => {
  return async dispatch => {
    try {
        let token = localStorage.getItem('token')
        if (token === null) {
          return
        }

        const response = await axios.get(`${config.API_URL}/api/v1/removeAll`, {
          headers: {Authorization: `Bearer ${localStorage.getItem('token')}`}
        });

        if (response.status !== 200) {
          console.log('возникла ошибка при удалении файлов:')
          console.log(response)
          // TODO вывод текста ошибки в интерфейсе
        } else {
          console.log('remove all')
        }
    } catch (e) {
      alert(e)
    }
  }
}
