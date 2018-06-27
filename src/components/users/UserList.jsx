import React, { Component } from 'react';
import User from './User.jsx';
import PropTypes from 'prop-types';

class UserList extends Component {

  render(){
    return(
      <ul>
        {this.props.users.map(user => {
          return(
            <User
              key={user.id}
              user={user}
          />)
        })}
      </ul>
    )
  }
}

UserList.propTypes = {
  users: PropTypes.array.isRequired
}

export default UserList;
