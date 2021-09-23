import React, {useState, useEffect} from 'react';
import Navbar from './components/navbar/Navbar';
import './styles/App.css';
import Disk from './components/disk/Disk';
import FormatSelect from './components/formatSelect/FormatSelect';
import { UserContext } from './context';
import { uploadFile, removeAll } from './actions/file';
import {registerUser} from './actions/user';
import {useDispatch} from 'react-redux';
import {setUser} from "./reducers/userReducer";
import Loader from './components/loader/Loader';
import ControlGroup from './components/control-group/ControlGroup';

function App() {
  // аутентификации в стандартном понимании нет,
  // просто при попытке загрузки файла в браузере сервер создает нового пользователя
  // и в дальнейшем идентифицирует его по нему
  const [token, setToken] = useState('')
  const dispatch = useDispatch()
  const [dragEnter, setDragEnter] = useState(false)
  const [extension, setExtension] = useState('mp4')
  const [extList] = useState(['mp4', 'avi', 'mpeg', 'mov', 'flv', 'webm', 'mkv', '3gp'])

  useEffect(() => {
    let token = localStorage.getItem('token')
    if (token) {
      setToken(token)
      dispatch(setUser(token))
    }
  }, [dispatch]);

  function fileUploadHandler(event) {
    const files = [...event.target.files]
    removeAll()
    files.forEach(file => dispatch(uploadFile(file, extension)))
  }

  function checkRegisterHandler() {
    if (!localStorage.getItem('token')) {
      registerUser()
    }
  }

  function dragEnterHandler(event) {
    event.preventDefault()
    event.stopPropagation()
    setDragEnter(true)
  }

  function dragLeaveHandler(event) {
    event.preventDefault()
    event.stopPropagation()
    setDragEnter(false)
  }

  function dropHandler(event) {
    event.preventDefault()
    event.stopPropagation()
    let files = [...event.dataTransfer.files]
    files.forEach(file => dispatch(uploadFile(file, extension)))
    setDragEnter(false)
  }

  return (
    <>
      <UserContext.Provider value={{
        token, setToken,
        extension, setExtension,
        extList
      }}>
        <Navbar/>
        <Loader/>
        { !dragEnter ?
           <div onDragEnter={dragEnterHandler} onDragLeave={dragLeaveHandler} onDragOver={dragEnterHandler}>
            <div className="container">
              <h1 className="text-center h1-title">Сервис конвертации видео</h1>
            </div>
            <div className="container df jc-c">
              <input onClick={checkRegisterHandler} onChange={(event) => fileUploadHandler(event)} id="upload-input" type="file" className="upload-input" multiple={false} />
              <label htmlFor="upload-input" className="button button_primary m-12">Загрузить и конвертировать в</label>
              <FormatSelect/>
            </div>
            <ControlGroup/>
            <Disk/>
          </div>
          :
          <div className="drop-area" onDrop={dropHandler} onDragEnter={dragEnterHandler} onDragLeave={dragLeaveHandler} onDragOver={dragEnterHandler}>
            Перетащите файл сюда
          </div>
        }
      </UserContext.Provider>
    </>
  );
}

export default App;
