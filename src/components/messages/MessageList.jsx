import React, { Component } from 'react';
import Message from './Message.jsx'
import PropTypes from 'prop-types';


class MessageList extends Component {

  render(){
    return(
      <ul>
        {this.props.messages.map( msg => {
          return(
            <Message key={msg.id} message={msg} />
          )
        })}
      </ul>
    )
  }
}

MessageList.propTypes = {
  messages: PropTypes.array.isRequired
}

export default MessageList;
