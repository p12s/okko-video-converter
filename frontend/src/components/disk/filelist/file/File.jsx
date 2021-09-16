import React from 'react';
import './File.css';
import preloader from '../../../../assets/img/preloader.gif'

const File = ({file}) => {
  let fileExtension = ''
  if (file.name) {
    fileExtension = file.name.split('.').pop().toUpperCase()
  }
  let megaByteSize = 0
  if (file.kilo_byte_size) {
    megaByteSize = Math.round(file.kilo_byte_size / 1024, 2)
  }
  if (file.prev_image === "") {
    file.prev_image = 'https://placehold.co/133x100/random/random'
  }
  let videoType = 'video/' + fileExtension.toLowerCase()

  return (
    <div className="file" key={file.name}>
      <div className="file-content df">
        <div className="file-preview" style={{ backgroundImage: `url(${preloader})` }}>
            <video width="133" controls poster={file.prev_image} >
                <source src={file.path} type={videoType} />
                Your browser doesn't support HTML5 video tag
            </video>
        </div>
        <div className="df fg-3">
          <div className="df fd-col ml-12 mr-12 color-gray vertical-sb">
            <span>Filename:</span>
            <span>Format:</span>
            <span>Size:</span>
          </div>
          <div className="df fd-col ml-12 mr-12 vertical-sb">
            <span>{file.name}</span>
            <span>{fileExtension}</span>
            <span>{megaByteSize} Mb</span>
          </div>
        </div>
        <div className="df fd-col fg-1 vertical-c align-end"></div>
      </div>
    </div>
  );
}

export default File;
