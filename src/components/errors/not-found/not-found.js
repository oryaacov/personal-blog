import React from 'react';
import notFoundSvg from '../../../assets/404.svg'
import './not-found.css';

const NotFound = (props) => {
  return (
    <div className="image-container">
        <img className="not-found-img" src={notFoundSvg}></img>
    </div>
  );
}

export default NotFound;
