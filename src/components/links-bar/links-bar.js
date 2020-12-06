import React, { Component } from "react";
import ImageLink from "../base/image-link-component/image-link";
import "./links-bar.css";
import github from "../../assets/images/github.png";
import linkedin from "../../assets/images/linkedin.png";
import stackoverflow from "../../assets/images/stackoverflow.png";

class LinksBar extends Component {
  render() {
    return (
      <div class="flex-container">
          <ImageLink href="https://github.com/oryaacov" text="Github" imageSrc={github}></ImageLink>
          <ImageLink href="https://stackoverflow.com/users/10115847/or-yaacov" text="Stack Overflow" imageSrc={stackoverflow}></ImageLink>
          <ImageLink href="https://www.linkedin.com/in/yaacovor/" text="Linkedin" imageSrc={linkedin}></ImageLink>
      </div>
    );
  }
}
export default LinksBar;
