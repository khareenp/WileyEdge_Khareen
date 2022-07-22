import React, { Component } from "react";

class ChatInput extends Component {
  render() {
    return (
      <div className="ChatInput block">
        <input
          className=" w-11/12 p-3 text-base rounded shadow-md"
          onKeyDown={this.props.send}
        />
      </div>
    );
  }
}

export default ChatInput;
