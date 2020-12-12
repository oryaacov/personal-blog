import React, { Component } from "react";
import "./image-link.css";

class ImageLink extends Component {
  render() {
    return (
      <div class="flex-container link">
        <a href={this.props.href}>
          <img class="link-image" src={this.props.imageSrc}></img>
          {this.props.text}
        </a>
      </div>
    );
  }
}
export default ImageLink;
