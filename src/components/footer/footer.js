import React, { Component } from "react";
import LinksBar from "../links-bar/links-bar";
import "./footer.css";

class Footer extends Component {
  render() {
    return (
      <div className="footer-container">
        <LinksBar links={this.props.profile.links}></LinksBar>
        <h4>
          © {new Date().getFullYear()} {this.props.profile.fullname}
        </h4>
      </div>
    );
  }
}

export default Footer;
