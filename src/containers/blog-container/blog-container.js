import React, {Component} from 'react';
import About from '../../components/about/about';
import Article from '../../components/article/article';
import {Route, Switch} from 'react-router-dom';
import ArticleContainer from '../article-container/article-container';

class BlogContainer extends Component {
  render() {
    return (
      <Switch>
        <Route
          path="/"
          exact
          component={() => (
            <ArticleContainer
              key={'A' + Math.random()}
              thumbnails={this.props.thumbnails}
            />
          )}
        />
        <Route path="/about" exact component={About} />
        <Route path="/article" exact component={Article} />
        <Route component={Error} />
      </Switch>
    );
  }
}

export default BlogContainer;
