import React from 'react';
import './Error.css';

const Error = () => {

  function removeError() {
    console.log('remove error')
  }

  return (
    <>
      <div className="error-container">
        <span className="error-close" onClick={removeError}>X</span>
        <div className="error-item">
          Upload file err: open /files/18579-gorutines.gif: no such file or directory
        </div>
        <div className="error-item">
          Upload file err: open /files/18579-gorutines.gif: no such file or directory
        </div>
      </div>
    </>
  );
}

export default Error;
