import React, {Component} from 'react';
import {HashRouter} from 'react-router-dom';
import BlogContainer from '../../containers/blog-container/blog-container';
import Footer from '../footer/footer';
import Data from "../../public/data.json";
import Navbar from '../navbar/navbar';
import Profile from '../profile/profile';
import './app.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <div class="flex-app-container">
        <HashRouter>
          <Navbar fullName={Data.profile.fullname}></Navbar>
          <Profile profile={Data.profile}></Profile>
          <BlogContainer></BlogContainer>
        </HashRouter>
        <Footer></Footer>
      </div>
    );
  }
}

export default App;
