import React, { useContext } from 'react';
import { UserContext } from '../../context/index';
import './FormatSelect.css';

const FormatSelect = () => {
  const {extList, extension, setExtension} = useContext(UserContext);

  return (
    <>
      <select className="format-select m-12" id="format-select" onChange={e => setExtension(e.target.value)} value={extension}>
        {extList.map((format, index) => 
            <option key={index+1}>{format}</option>
        )}
      </select>
    </>
  );
}

export default FormatSelect;
