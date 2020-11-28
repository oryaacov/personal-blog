import React, {Component} from 'react';
import './article-thumbnail.css';
import {useHistory} from 'react-router-dom';

class ArticleThumbnail extends Component {

  constructor(props){
    super(props);
    this.changeRoute = this.changeRoute.bind(this);
  }
  changeRoute() {
    console.log(this)
    //history = useHistory();
    //this.history.push(this.props.key);
  }

  render() {
    return (
      <div onClick={this.changeRoute} className="article-thumbnail-box">
        <div>
          <h1 className="link">{this.props.title}</h1>
          <img className=" thumbnail-img" src={this.props.image}></img>
        </div>
        <p>{this.props.summary}</p>
        <p>
          {new Date(this.props.publishDate).toUTCString().replace('GMT', '')}
        </p>
      </div>
    );
  }
}
export default ArticleThumbnail;
