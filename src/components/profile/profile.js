import React from 'react';
import LinksBar from '../links-bar/links-bar';
import './profile.css';

const Profile = (props) => {
  return (
    <div className="profile-flex-container">
      <div>
        <img
          className="profile-image"
          src="https://avatars.githubusercontent.com/u/48175064?s=400&u=06d14e0f94bd4fbb44cd15704e765ac8866c7a14&v=4"
        ></img>
      </div>
      <div className="profile-info-flex">
        <h1 className='profile-text'>
          {props.profile.fullname}
        </h1>
        <h3 className='profile-text'>{props.profile.title}</h3>
        <h4 className='profile-text'>{props.profile.message}</h4>
        <LinksBar links={props.profile.links}></LinksBar>
      </div>
    </div>
  );
}

export default Profile;
