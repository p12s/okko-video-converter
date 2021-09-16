import axios from 'axios';
import { setFiles } from "../reducers/fileReducer";
import { config } from '../config';
import { registerUser } from './user';

import { showLoader, hideLoader, changeLoader } from "../reducers/loaderReducer";
import { showProgress, hideProgress, changeProgress } from "../reducers/progressReducer";
import { replaceFile } from "../reducers/fileReducer"; // TODO когда не нужно будет перезаписывать, а просто добавлять - addFile
import { showCreateButton } from '../reducers/createButtonReducer';
import { showDownloadButton } from '../reducers/downloadButtonReducer';
import { showDimensions } from '../reducers/dimensionsReducer';
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
            dispatch(showDimensions())
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

// Загрузка файла на сервер - отрабатывает на каждый файл
export const uploadFile = (file) => {
  return async dispatch => {
      try {
          let token = localStorage.getItem('token')
          if (token === null) {
            registerUser()
          }

          const formData = new FormData()
          formData.append('files', file)

          const loaderProgress = {progress: 0}
          dispatch(showLoader())
          dispatch(changeLoader(loaderProgress))
          
          console.log('upload file start')
          const response = await axios.post(`${config.API_URL}/api/v1/upload`, formData, {
            headers: {Authorization: `Bearer ${localStorage.getItem('token')}`},
            onUploadProgress: progressEvent => {
              const totalLength = progressEvent.lengthComputable ? progressEvent.total : progressEvent.target.getResponseHeader('content-length') || progressEvent.target.getResponseHeader('x-decompressed-content-length');
              if (totalLength) {
                loaderProgress.progress = Math.round((progressEvent.loaded * 100) / totalLength)
                console.log(loaderProgress.progress)
                dispatch(changeLoader(loaderProgress))
              }
            }
          })
          if (response.status === 200) {
            console.log('should be end')
          }

          if (response.status !== 200) {
            console.log('взникла ошибка при загрузке файла:')
            console.log(response)
            // TODO вывод текста ошибки в интерфейсе
          }

          dispatch(hideLoader())
          //dispatch(addFile(response.data.files[0])) - для множественной загрузки, пока не надо
          dispatch(replaceFile(response.data.files[0]))
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

// TODO возврат ответа ОК - что запустилась нарезка
// Запуск нарезки/webp-конвертаци/архивации картинок
export const updateResizeOptions = (widthList, coefList, isAddWebp, isCompressImage) => {
  return async dispatch => {
    try {
      let postData = {
        'widthList': widthList,
        'coefList': coefList,
        'isAddWebp': isAddWebp,
        'isCompress': isCompressImage
      }

      const response = await axios.post(`${config.API_URL}/api/v1/updateResizeOptions`, postData, {
        headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
      })

      if (response.status === 200) {
        dispatch(showProgress())
        //dispatch(resize()) EventSourse
        setTimeout(() => dispatch(getResizeProgress()), 1000);
      }

    } catch (e) {
      console.log(e)
      dispatch(hideProgress())
      alert(e)
      // TODO вывод текста ошибки в интерфейсе
    }
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
          console.log('возникла ошибка при загрузке файла:')
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

export const resize = () => {
  return async dispatch => {
    try {
      let token = localStorage.getItem('token')
      if (token === null) {
        return
      }
    
      const eventSourse = new EventSource(`${config.API_URL}/api/v1/resize/${localStorage.getItem('token')}`)

      eventSourse.onmessage = function (event) {
        const message = JSON.parse(event.data)

        if (message.Current === message.Total) {
          eventSourse.close();
          dispatch(hideProgress())
          dispatch(showCreateButton())
          dispatch(showDownloadButton())
        }

        let progressPercent = Math.round((message.Current * 100)/message.Total)
        if (progressPercent > 100) {
          progressPercent = 100
        }
        console.log(progressPercent)
        dispatch(changeProgress({progress: progressPercent}))
      }
      eventSourse.onerror = function (event) {
        console.log('error and close stream:')
        console.log(event)
        eventSourse.close();
        dispatch(hideProgress())
        dispatch(showCreateButton())
        dispatch(showDownloadButton())
      }
    } catch (e) {
      alert(e)
    }
  }
}

export const getResizeProgress = () => {
  return async dispatch => {
    try {
      let token = localStorage.getItem('token')
      if (token === null) {
        return
      }

      const response = await axios.get(`${config.API_URL}/api/v1/getResizeProgress`, {
        headers: {Authorization: `Bearer ${localStorage.getItem('token')}`}
      });

      if (response.status === 200) {
        if (response.data.hasOwnProperty('progress')) {
          let progress = parseInt(response.data.progress)
          console.log(progress)

          if (progress >= 0 && progress < 100) {
            dispatch(changeProgress({progress: progress})) 
            setTimeout(() => dispatch(getResizeProgress()), 1000);

          } else if (progress === 100) {
            dispatch(hideProgress())
            dispatch(showCreateButton())
            dispatch(showDownloadButton())

          } else {
            console.log('Error progress:')
            console.log(response)
          }
        }
      }

    } catch (e) {
      alert(e)
    }
  }
}
