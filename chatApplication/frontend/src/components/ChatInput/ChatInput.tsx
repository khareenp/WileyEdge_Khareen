import { FC } from "react";

// props
type ChatInputProps = {
  name: string;
  value: string;
  handleChange: any;
  handleSubmit: any;
};

// simple input to handle chat messages
const ChatInput: FC<ChatInputProps> = ({
  name,
  handleChange,
  handleSubmit,
  value,
}) => {
  return (
    <div className="ChatInput ml-6">
      <input
        name={name}
        className=" w-11/12 p-2 my-1 text-base rounded shadow-md "
        onChange={handleChange}
        value={value}
        onSubmit={handleSubmit}
        placeholder="Please type a message, hit enter to send"
      />
    </div>
  );
};

export default ChatInput;
