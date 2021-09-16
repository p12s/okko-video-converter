import React, { useState, useContext } from 'react';
import { useSelector } from 'react-redux';
import Button from '../ui/button/Button';
import './ControlGroup.css';
import { downloadFile, updateResizeOptions } from '../../actions/file';
import { useDispatch } from 'react-redux';
import ProgressBar from '../progress-bar/ProgressBar';
import { hideCreateButton } from "../../reducers/createButtonReducer";
import { hideDownloadButton } from "../../reducers/downloadButtonReducer";
import { UserContext } from '../../context/index';

const ControlGroup = () => {
  const dispatch = useDispatch()
  const {widthList, coefList, removedFiles} = useContext(UserContext);
  const isShowCreateButton = useSelector( state => state.createButton.isVisible )
  const isShowDownloadButton = useSelector( state => state.downloadButton.isVisible )

  const [isAddWebpCheckbox, setIsAddWebpCheckbox] = useState(false)
  const [isCompressCheckbox, setIsCompressImageCheckbox] = useState(true)

  function downloadClickHandler(event) {
    if (isShowDownloadButton) {
      event.stopPropagation()
      downloadFile()
    }
  }

  function resizeImagesHandler(event) {
    event.stopPropagation()
    if (isShowCreateButton) {
      dispatch(hideCreateButton())
      dispatch(hideDownloadButton())
      dispatch(updateResizeOptions(widthList, coefList, isAddWebpCheckbox, isCompressCheckbox))
    }
  }

  function changeIsAddWebp() {
    setIsAddWebpCheckbox(!isAddWebpCheckbox)
  }

  function changeIsCompress() {
    setIsCompressImageCheckbox(!isCompressCheckbox)
  }

  return (
    <>
      <div className="container df jc-c top-control-indent">
        <div className="df direction-col checkbox-width">
          <label htmlFor="add-webp">
            <input onChange={changeIsAddWebp}
              type="checkbox" id="add-webp"checked={isAddWebpCheckbox} />
            <span className="checkbox-label-text">Add .webp</span>
          </label>
          <label htmlFor="compress-img">
            <input onChange={changeIsCompress}
              type="checkbox" id="compress-img" checked={isCompressCheckbox} />
            <span className="checkbox-label-text">Compress images</span>
          </label>
        </div>
        <Button onClick={(e) => resizeImagesHandler(e)} 
          className={isShowCreateButton ? "button button_primary mr-12 ml-12" : "button button_primary mr-12 ml-12 disabled"}>
          Create images ({widthList.length * coefList.length - removedFiles})
        </Button>
        <div className="counterweight checkbox-width">
          <p className="counterweight-invisible-text">Lilo Pic - image video</p>
        </div>
      </div>
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
