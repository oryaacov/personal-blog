import React from 'react';
import LinksBar from '../links-bar/links-bar';
import './footer.css';

const Footer = (props) => {
  return (
    <div className="footer-container">
      <LinksBar links={props.profile.links}></LinksBar>
      <h4>
        © {new Date().getFullYear()} {props.profile.fullname}
      </h4>
    </div>
  );
}

export default Footer;
