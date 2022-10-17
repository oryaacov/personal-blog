
import React, { useEffect, useState } from 'react';
import ArticleThumbnail from '../../components/thumbnails/article-thumbnail';
import './thumbnails-container.css';
import axios from 'axios';
import { spinner } from '../../components/common/spinner';
import { baseUrl } from '../../utils/config';
const loadThumbnails = async () => {
  console.log("test", `${baseUrl}/api/v1/thumbnails`)
  return (
    await axios.get(`${baseUrl}/api/v1/thumbnails`)
  ).data;
}

const generateThumbnails = (thumbnails) => {
  return thumbnails.map((articleThumbnail) => {
    return <ArticleThumbnail
        key={articleThumbnail.id}
        id={articleThumbnail.id}
        title={articleThumbnail.title}
        image={articleThumbnail.img}
        summary={articleThumbnail.summary}
        publishDate={articleThumbnail.date}>
      </ArticleThumbnail>
  })
}


const ArticleContainer = (props) => {
  const [state, setState] = useState({ thumbnails: [], loading: true })

  useEffect(async () => {
    const thumbnails = await loadThumbnails();
    setState({ thumbnails, loading: false })
  }, []);

  return state.loading ? spinner : <div className='thumbnails-container'>{generateThumbnails(state.thumbnails)}</div>;
}

export default ArticleContainer;
