import React, {Component} from 'react';
import LinksBar from '../links-bar/links-bar';
import './profile.css';

class Profile extends Component {
  render() {
    return (
      <div className="profile-flex-container">
        <div>
          <img
            className="profile-image"
            src="https://avatars.githubusercontent.com/u/48175064?s=400&u=06d14e0f94bd4fbb44cd15704e765ac8866c7a14&v=4"
          ></img>
        </div>
        <div className="profile-info-flex">
          <h1 className="name-h1 profile-text">
            {this.props.profile.fullname}
          </h1>
          <h2 className="profile-text">{this.props.profile.title}</h2>
          <p className="profile-text">{this.props.profile.message}</p>
          <LinksBar links={this.props.profile.links}></LinksBar>
        </div>
      </div>
    );
  }
}

export default Profile;
