import React from 'react';
import { useSelector } from 'react-redux';
import './Loader.css';

const Loader = () => {
  const isVisible = useSelector( state => state.loader.isVisible )
  const progress = useSelector(state => state.loader.progress)

  return ( isVisible &&
    <>
      <div className="loader-container">
        <div className="loader-progress" style={{width: progress + "%"}}></div>
      </div>
    </>
  );
}

export default Loader;
