import React, {Component} from 'react';
import {HashRouter} from 'react-router-dom';
import BlogContainer from '../../containers/blog-container/blog-container';
import Footer from '../footer/footer';
import Data from '../../public/data.json';
import Navbar from '../navbar/navbar';
import Profile from '../profile/profile';
import './app.css';
import ReactGA from 'react-ga';

class App extends Component {
  constructor(props) {
    ReactGA.initialize('UA-000000-01');
    ReactGA.pageview(window.location.pathname + window.location.search);
    super(props);
    this.state = {};
  }

  render() {
    return (
      <div className="flex-app-container">
        <HashRouter>
          <Navbar fullName={Data.profile.fullname}></Navbar>
          <Profile profile={Data.profile}></Profile>
          <BlogContainer
            profile={Data.profile}
          ></BlogContainer>
        </HashRouter>
        <Footer profile={Data.profile}></Footer>
      </div>
    );
  }
}

export default App;
