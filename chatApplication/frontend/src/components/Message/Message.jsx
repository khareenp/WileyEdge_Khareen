import React, { Component } from "react";

class Message extends Component {
  constructor(props) {
    super(props);
    let temp = JSON.parse(this.props.message);
    this.state = {
      message: temp,
    };
  }

  render() {
    return (
      <div className="Message block bg-sky-100 m-1 p-2 rounded-sm shadow-sm clear-both w-fit font-light">
        {this.state.message.body}
      </div>
    );
  }
}

export default Message;
