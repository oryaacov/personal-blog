import React, { Component } from 'react';
import About from '../../components/about/about';
import Article from '../../components/article/article';
import { Route, Routes, Navigate } from 'react-router-dom';
import ArticleContainer from '../thumbnails-container/thumbnails-container';

const BlogContainer = () => {
  return (
    <Routes>
      <Route path="/about" exact element={<About />} />
      <Route path="/articles/:id" element={<Article />} />
      <Route path="/:category" element={<ArticleContainer />} />
      <Route path="*" element={<Navigate to="/tech" replace />} />
    </Routes>
  );
}

export default BlogContainer;
