import React from 'react';
import { Link } from 'react-router-dom';
import './navbar.css';

const Navbar = (props) => {
  return (
    <div className="navbar-container">
      <div className="links-flex">
        <Link to="/" className="header header-name">
          {props.fullName}
        </Link>
      </div>
      <div className="links-flex">
        <Link className="link" to="/">
          Articles
        </Link>
      </div>
    </div>
  );
}

export default Navbar;
