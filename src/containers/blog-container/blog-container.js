import React, {Component} from 'react';
import Gallery from '../../components/gallery/gallery';
import About from '../../components/about/about';
import FirstStory from '../../articles/first-story/first-story';
import {Route, Switch} from 'react-router-dom';
class BlogContainer extends Component {
  render() {
    return (
      <Switch>
        <Route path='/' exact component={Gallery} />
        <Route path='/about' exact component={About} />
        <Route path='/first-story' exact component={FirstStory} />
        <Route component={Error} />
      </Switch>
    );
  }
}
export default BlogContainer;
