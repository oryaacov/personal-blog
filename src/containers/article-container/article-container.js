import React, {Component} from 'react';
import ArticleThumbnail from '../../components/article/article-thumbnail';
import './article-container.css';
import axios from 'axios';
import {SpinnerInfinity} from 'spinners-react';
class ArticleContainer extends Component {
  constructor() {
    super();
    this.state = {thumbnails: [], loading: true};
    this.loadData();
  }

   async loadData() {
    const thumbnails = await (await axios.get(`http://localhost:8080/api/v1/articles`)).data;
    this.setState({loading: false, thumbnails});
  }

  render() {
    const {thumbnails, loading} = this.state;
    
    const articleThumbnails = thumbnails.map((articleThumbnail) => (
      <ArticleThumbnail
        key={articleThumbnail.id}
        title={articleThumbnail.title}
        image={articleThumbnail.img}
        summary={articleThumbnail.summary}
        publishDate={articleThumbnail.date}
      ></ArticleThumbnail>
    ));
    const spinner =  <div class="spinner-container"><SpinnerInfinity size={69} thickness={141} speed={100} color="#36ad47" secondaryColor="#EDEFF0" /></div>;

    const renderSite = () => {
      return loading ? spinner : articleThumbnails;
    };

    return renderSite();
  }
}

export default ArticleContainer;
