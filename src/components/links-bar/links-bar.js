import React from 'react';
import ImageLink from '../base/image-link-component/image-link';

const LinksBar = (props) => {
  const links = props.links.map((link) => (
    <ImageLink
      key={link.text}
      href={link.url}
      text={link.text}
      imageSrc={link.icon}
    ></ImageLink>
  ));

  return <div className="flex-container">{links}</div>;
}

export default LinksBar;
