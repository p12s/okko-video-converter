import React from 'react';
import './Checkbox.css'

const Checkbox = (props) => {
  return (
      <label for={props.for}>
        <input onChange={(event)=> props.setValue(event.target.value)}
          value={props.value}
          type={props.type}
          id={props.value}
          checked={props.checked}/>
        <span className="checkbox-label-text">{props.text}</span>
      </label>
  );
};

export default Checkbox;
