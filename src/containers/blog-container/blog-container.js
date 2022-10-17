import React, { Component } from 'react';
import About from '../../components/about/about';
import Article from '../../components/article/article';
import { Route, Routes } from 'react-router-dom';
import ArticleContainer from '../thumbnails-container/thumbnails-container';

const BlogContainer = () => {
  return (
    <Routes>
      <Route
        path="/"
        exact
        element={<ArticleContainer />} />
      <Route path="/about" exact element={<About />} />
      <Route path="/articles/:id" element={<Article />} />
    </Routes>
  );
}

export default BlogContainer;
