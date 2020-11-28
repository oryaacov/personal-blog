import React, {Component} from 'react';
import ArticleThumbnail from '../../components/article/article-thumbnail';
import './article-container.css';

class ArticleContainer extends Component {
  render() {
    const thumbnails = this.props.thumbnails.map((articleThumbnail) => (
      <ArticleThumbnail
        key={articleThumbnail.href}
        title={articleThumbnail.title}
        image={articleThumbnail.image}
        summary={articleThumbnail.summary}
        publishDate={articleThumbnail.publishDate}
      ></ArticleThumbnail>
    ));
    return (
      <div className="flex-container thumbnail-container">{thumbnails}</div>
    );
  }
}

export default ArticleContainer;
