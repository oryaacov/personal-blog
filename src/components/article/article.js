import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { spinner } from '../common/spinner';
import axios from 'axios';
import ReactMarkdown from 'react-markdown'
import './article.css';
import NotFound from '../errors/not-found/not-found';
import { baseUrl } from '../../utils/config';

const getArticleById = async (id, setArticleState) => {
  try {
    const article = (await axios.get(`${baseUrl}/api/v1/articles/${id}`)).data;

    setArticleState({ id, article: article.body, loading: false, title: article.title });
  } catch (e) {
    setArticleState({ id, loading: false, error: e.response.status })
  }
}

const Article = (props) => {
  const { id } = useParams();
  const [articleState, setArticleState] = useState({ id, article: "", loading: true });

  useEffect(() => {
    getArticleById(id, setArticleState);
  }, []);

  if (articleState.error){
    return <NotFound/>;
  }

  const article =
    <div className="article-container">
      <div className="article">
      <h1>{articleState.title}</h1>
        {/* eslint-disable-next-line */}
        <ReactMarkdown children={articleState.article} />
      </div>
    </div>

  return articleState.loading ? spinner : article;
}

export default Article;
