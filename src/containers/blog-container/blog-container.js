import React, {Component} from 'react';
import About from '../../components/about/about';
import Article from '../../components/article/article';
import {Route, Routes} from 'react-router-dom';
import ArticleContainer from '../article-container/article-container';

class BlogContainer extends Component {
  render() {
    return (
      <Routes>
        <Route
          path="/"
          exact
          element={<ArticleContainer/>}/>
        <Route path="/about" exact element={<About/>} />
        <Route path="/article" exact element={<Article/>} />
        <Route element={<Error/>} />
      </Routes>
    );
  }
}

export default BlogContainer;
