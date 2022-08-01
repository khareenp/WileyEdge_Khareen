import { FC } from "react";

type ChatInputProps = {
  name: string;
  value: string;
  handleChange: any;
  handleSubmit: any;
};

const ChatInput: FC<ChatInputProps> = ({
  name,
  handleChange,
  handleSubmit,
  value,
}) => {
  return (
    <div className="ChatInput block ml-6">
      <input
        name={name}
        className=" w-11/12 p-2 text-base rounded shadow-md"
        onChange={handleChange}
        value={value}
        onSubmit={handleSubmit}
        placeholder="Please type a message, hit enter to send"
      />
    </div>
  );
};

export default ChatInput;
