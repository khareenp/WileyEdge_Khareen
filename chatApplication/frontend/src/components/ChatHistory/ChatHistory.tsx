// This component will take in an array of chat messages from our App.js function
// through its' props and will subsequently render them one under the other.

import { FC } from "react";
import Message from "../Message/Message";

type ChatHistoryProps = {
  chatHistory: string[];
  ip: string;
};

const ChatHistory: FC<ChatHistoryProps> = ({ chatHistory, ip }) => {
  var today = new Date();
  const date =
    today.getFullYear() + "-" + (today.getMonth() + 1) + "-" + today.getDate();

  const messages = chatHistory.map((msg, index) => (
    <Message key={index.toString()} message={msg} />
  ));

  return (
    <div className="ChatHistory bg-slate-100 text-left h-5/6 md:text-2xl sm:text-xl md:font-semibold tracking-wide">
      <h2 className=" p-5 bg-sky-200 font-semibold">
        Chat History
        {/* <p>{this.state.currentDateTime}</p> */}
        <p>My IP Address:{ip}</p>
        <p>{date}</p>
      </h2>
      {/* prop for message data */}
      {messages}
    </div>
  );
};

export default ChatHistory;