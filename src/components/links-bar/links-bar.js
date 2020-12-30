import React, { Component } from "react";
import ImageLink from "../base/image-link-component/image-link";
import "./links-bar.css";


class LinksBar extends Component {
  
  render() {
    
    const links = this.props.links.map(link =>
      <ImageLink href={link.url} text={link.text} imageSrc={link.icon}></ImageLink>
       );
    return (
      <div class="flex-container">
          {links}
      </div>
    );
  }
}
export default LinksBar;
 