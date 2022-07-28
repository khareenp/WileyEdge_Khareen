import React, { Component } from "react";
import GetIP from "../GetIP/GetIP";

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
      <>
        <div className="flex flex-col items-start p-2 ml-6">
          <GetIP />
          <div className="Message block bg-sky-100 m-1 p-2 rounded-sm shadow-sm clear-both w-fit font-light">
            {this.state.message.body}
          </div>
        </div>
      </>
    );
  }
}

export default Message;
