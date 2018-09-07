import React, { Component } from 'react';
import PropTypes from 'prop-types';

class User extends Component {
  render(){
    let user = this.props.user;
    return(
      <li>
        {user.name}
      </li>
    )
  }
}

User.propTypes = {
  user: PropTypes.object.isRequired
}

export default User;
