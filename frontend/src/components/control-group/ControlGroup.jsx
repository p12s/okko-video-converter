import React from 'react';
import { useSelector } from 'react-redux';
import Button from '../ui/button/Button';
import './ControlGroup.css';
import { downloadFile } from '../../actions/file';
import ProgressBar from '../progress-bar/ProgressBar';

const ControlGroup = () => {
  const isShowDownloadButton = useSelector( state => state.downloadButton.isVisible )

  function downloadClickHandler(event) {
    if (isShowDownloadButton) {
      event.stopPropagation()
      downloadFile()
    }
  }

  return (
    <>
      <ProgressBar/>
      {isShowDownloadButton &&
        <div className="container df jc-c top-control-indent bottom-control-indent">
          <Button onClick={(e) => downloadClickHandler(e)} 
            className="button button_secondary">Download</Button>
        </div>
      }
    </>
  );
}

export default ControlGroup;
