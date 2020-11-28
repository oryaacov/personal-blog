import React, {Component} from 'react';
import {HashRouter} from 'react-router-dom';
import BlogContainer from '../../containers/blog-container/blog-container';
import Footer from '../footer/footer';
import Navbar from '../navbar/navbar';
class App extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <div>
        <HashRouter>
          <Navbar></Navbar>
          <BlogContainer></BlogContainer>
        </HashRouter>
        <Footer></Footer>
      </div>
    );
  }
}

export default App;
