import React, { Component } from "react";
import LinksBar from "../links-bar/links-bar";
import "./profile.css";

class Profile extends Component {
  render() {
    return (
      <div class="profile-flex-container">
        <div>
          <img
            class="profile-image"
            src="https://media-exp1.licdn.com/dms/image/C4E03AQGM1OrFrgewkg/profile-displayphoto-shrink_200_200/0/1561146151867?e=1612396800&v=beta&t=1gaBVdcyORr5geahLKggsmFPA_Eghk4EexCMg6Xj_fg"
          ></img>
        </div>
        <div class="profile-info-flex">
          <h1 class="name-h1 profile-text">{this.props.profile.fullname}</h1>
          <h2 class="profile-text">{this.props.profile.title}</h2>
          <p class="profile-text">{this.props.profile.message}</p>
          <LinksBar links={this.props.profile.links}></LinksBar>
        </div>
      </div>
    );
  }
}
export default Profile;
