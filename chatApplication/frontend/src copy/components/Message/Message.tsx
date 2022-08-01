import { FC } from "react";
import { ChatMessage, useIP } from "../../App";

type MessageProps = {
  message: ChatMessage;
};

//this component is used to display message in chat window
const Message: FC<MessageProps> = ({ message }) => {
  const IP = useIP();
  let currentTime = "";
  if (message.timestamp) {
    const date = new Date(message.timestamp);
    currentTime =
      date.getHours() + ":" + date.getMinutes() + ":" + date.getSeconds();
  }
  if (IP === undefined) return <div>loading...</div>;
  if (IP === null) return <div>error loading this message</div>;

  // Display the user's identity
  const userDetails = message.uip === IP ? "Me" : `User: ${message.uip ?? ""}`;
  return (
    <>
      <div className="flex flex-col items-start p-2 ml-6">
        <div className="text-sm">
          <p>Time stamp:{currentTime}</p>
          <h1>{userDetails}</h1>
        </div>
        <div className="Message block bg-sky-100 m-1 p-2 rounded-sm shadow-sm clear-both w-fit font-light">
          {message.text}
        </div>
      </div>
    </>
  );
};

export default Message;
