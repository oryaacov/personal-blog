import React from 'react';
import './image-link.css';

const ImageLink = (props) => {
  return (
      <a className="image-link" href={props.href}>
        <img className="link-image" src={props.imageSrc}></img>
        <div className="link-text"><p>{props.text}</p></div>
      </a>
  );
}

export default ImageLink;
