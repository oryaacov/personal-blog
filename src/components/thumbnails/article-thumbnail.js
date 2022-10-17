import React from 'react';
import { Link } from 'react-router-dom';
import './article-thumbnail.css';

const ArticleThumbnail = (props) => {
  return (
    <div className="article-thumbnail-box">
      <Link to={`/articles/${props.id}`}>
        <h2 className="link">{props.title}</h2>
        <img className="thumbnail-img" src={props.image}></img>
        <div className='thumbnail-summary'>{props.summary}</div>
      </Link>
      <div className='thumbnail-date'>
        {new Date(props.publishDate).toUTCString().replace('GMT', '')}
      </div>
    </div>
  );
}

export default ArticleThumbnail;
