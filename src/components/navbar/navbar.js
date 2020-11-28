import React, {Component} from 'react';
import {Link} from 'react-router-dom';
class Navbar extends Component {
  render() {
    return (
      <div>
        <h4>header</h4>
        <div>
          <Link to='/'>Home</Link>
          <Link to='/about'>About</Link>
        </div>
      </div>
    );
  }
}
export default Navbar;
