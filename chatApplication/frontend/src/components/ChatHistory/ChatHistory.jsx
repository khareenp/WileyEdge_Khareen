// This component will take in an array of chat messages from our App.js function
// through its' props and will subsequently render them one under the other.

import React, { Component } from "react";
import GetIP from "../GetIP/GetIP";
import Message from "../Message/Message";

class ChatHistory extends Component {
  constructor() {
    super();
    var today = new Date(),
      date =
        today.getFullYear() +
        "-" +
        (today.getMonth() + 1) +
        "-" +
        today.getDate();
    this.state = {
      currentDate: date,
      // currentDateTime: Date().toLocaleString(),
    };
  }

  render() {
    console.log(this.props.chatHistory);
    const messages = this.props.chatHistory.map((msg) => (
      <Message message={msg.data} />
    ));

    return (
      <div className="ChatHistory bg-slate-100 text-left h-5/6 md:text-2xl sm:text-xl md:font-semibold tracking-wide">
        <h2 className=" p-5 bg-sky-200 font-semibold">
          Chat History
          {/* <p>{this.state.currentDateTime}</p> */}
          <p>{this.state.currentDate}</p>
        </h2>
        {/* prop for message data */}
        {messages}
      </div>
    );
  }
}

export default ChatHistory;
