import React, { Component } from "react";
import "./article-thumbnail.css";

class ArticleThumbnail extends Component {
  render() {
    return (
      <div className="article-thumbnail-box">
        <h1 className="link">{this.props.title}</h1>
        <img className=" thumbnail-img" src={this.props.image}></img>
        <p>{this.props.summary}</p>
        <p>
          {new Date(this.props.publishDate).toUTCString().replace("GMT", "")}
        </p>
      </div>
    );
  }
}
export default ArticleThumbnail;
