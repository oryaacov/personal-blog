import React, { Component } from "react";
import { Link } from "react-router-dom";
import "./navbar.css";

class Navbar extends Component {
  render() {
    return (
      <div className="navbar-container">
        <div className="links-flex">
          <Link to="/" className="header header-name">
            {this.props.fullName}
          </Link>
        </div>
        <div className="links-flex">
          <Link className="link" to="/tutorials">
            Tutorials
          </Link>
          <Link className="link" to="/about">
            Articles
          </Link>
          <Link className="link" to="/about">
            About
          </Link>
        </div>
      </div>
    );
  }
}

export default Navbar;
