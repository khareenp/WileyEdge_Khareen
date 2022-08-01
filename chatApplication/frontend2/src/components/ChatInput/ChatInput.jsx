import React, { Component } from "react";

class ChatInput extends Component {
  render() {
    return (
      <div className="ChatInput block">
        <input
          name={this.props.name}
          className=" w-11/12 p-3 text-base rounded shadow-md"
          onChange={this.props.handleChange}
          value={this.props.value}
          onKeyDown={this.props.send}
          placeholder="Please type a message, hit enter to send"
        />
      </div>
    );
  }
}

export default ChatInput;
