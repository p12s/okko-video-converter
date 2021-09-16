import React from 'react';
import { useSelector } from 'react-redux';
import File from './file/File';
import { CSSTransition, TransitionGroup } from 'react-transition-group';

const FileList = () => {
  let files = useSelector(state => state.files.files)

  return ( files &&
    <TransitionGroup>
      {files.map((file, index) => 
        <CSSTransition
          key={index+1}
          timeout={200}
          classNames="file"
        >
          <File file={file} number={index+1} key={index+1}/>
        </CSSTransition>
      )}
    </TransitionGroup>
  );
}

export default FileList;
