import React, { Component } from "react";
import { Link } from "react-router-dom";
import "./navbar.css";

class Navbar extends Component {
  render() {
    return (
      <div class="navbar-container">
       <div class="links-flex">
        <Link to="/" class="header header-name">
          {this.props.fullName}
        </Link>
        </div>
        <div class="links-flex">
          <Link class="link" to="/tutorials">
            Tutorials
          </Link>
          <Link class="link" to="/about">
            Articles
          </Link>
          <Link class="link" to="/about">
            About
          </Link>
        </div>
      </div>
    );
  }
}
export default Navbar;
