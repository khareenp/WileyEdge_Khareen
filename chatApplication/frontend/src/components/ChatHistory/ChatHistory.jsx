// This component will take in an array of chat messages from our App.js function
// through its' props and will subsequently render them one under the other.

import React, { Component } from "react";
import Message from "../Message/Message";

class ChatHistory extends Component {
  render() {
    console.log(this.props.chatHistory);
    const messages = this.props.chatHistory.map((msg) => (
      <Message message={msg.data} />
    ));

    return (
      <div className="ChatHistory bg-slate-100 text-left h-5/6 md:text-2xl sm:text-xl md:font-semibold tracking-wide">
        <h2 className=" p-5 bg-sky-200 min-w-max font-semibold">
          Chat History
        </h2>
        {/* prop for message data */}
        {messages}
      </div>
    );
  }
}

export default ChatHistory;
