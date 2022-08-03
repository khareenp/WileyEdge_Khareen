// This component will take in an array of chat messages from our App.js function
// through its' props and will subsequently render them one under the other.

import { FC } from "react";
import { ChatMessage } from "../../App";
import Message from "../Message/Message";
``;
type ChatHistoryProps = {
  chatHistory: ChatMessage[];
  ip: string;
};

const ChatHistory: FC<ChatHistoryProps> = ({ chatHistory, ip }) => {
  var today = new Date();
  const date =
    today.getFullYear() + "-" + (today.getMonth() + 1) + "-" + today.getDate();

  // message lst
  const messages = chatHistory.map((msg, index) => (
    <Message key={index.toString()} message={msg} />
  ));

  return (
    <div className="ChatHistory bg-slate-100 text-left h-full md:text-2xl sm:text-xl md:font-semibold tracking-wide">
      <h2 className="p-6 md:p-4 bg-sky-200 font-semibold text-xl">
        Chat History
        {/* <p>{this.state.currentDateTime}</p> */}
        <p className=" text-sm">Todays Date: {date}</p>
      </h2>
      <p className="text-sm">
        {/* prop for message data */}
        {messages}
      </p>
    </div>
  );
};

export default ChatHistory;
