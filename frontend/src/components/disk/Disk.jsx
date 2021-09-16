import React, {useEffect} from 'react';
import { useDispatch } from 'react-redux';
import FileList from './filelist/FileList';
import {getFiles} from '../../actions/file';

const Disk = () => {
  const dispatch = useDispatch()
  // в примере вызов запроса по изменеиню теущ директории, 
  // а мне надо при загрузке новых файлов

  useEffect(() => {
    dispatch(getFiles())
  }, [dispatch])

  return (
    <div>
      <FileList/>
    </div>
  );
}

export default Disk;
