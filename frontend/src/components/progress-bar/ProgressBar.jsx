import React from 'react';
import { useSelector } from 'react-redux';
import './ProgressBar.css';
import { CSSTransition, TransitionGroup } from 'react-transition-group';

const ProgressBar = () => {
  const isVisible = useSelector( state => state.progress.isVisible )

  return ( isVisible &&
    <>
      <TransitionGroup>
        <CSSTransition
          key="progress-container"
          timeout={200}
          classNames="container"
        >
          <div className="container df jc-c top-control-indent bottom-control-indent">
            <div className="progress-container mr-12 ml-12">
              <div className="progress-fill" style={{width: "100%"}}></div>
              <span className="progress-text">
                <span className="progress-number"></span>
              </span>
            </div>
          </div>
        </CSSTransition>
      </TransitionGroup>
    </>
  );
}

export default ProgressBar;

