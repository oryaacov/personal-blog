import React, { Component } from "react";
import { Link } from "react-router-dom";
import "./navbar.css";

class Navbar extends Component {
  render() {
    return (
      <div class="flex-container">
        <Link to="/" class="header header-name">Or Yaacov</Link>
        <ul class="flex-container">
          <li> 
          <Link to="/tutorials">Tutorials</Link>
          </li>
          <li>
          <Link to="/about">About</Link>
          </li>
        </ul>
      </div>
    );
  }
}
export default Navbar;
