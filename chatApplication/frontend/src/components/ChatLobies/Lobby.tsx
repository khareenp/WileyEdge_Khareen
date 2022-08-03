import { Formik } from "formik";
import { FC, useEffect, useState } from "react";
import { ChatMessage } from "../../App";

const Lobby: FC<{
  chatHistory: ChatMessage[];
  newChatroom: (name: string) => void;
  setChatroom: (chatroom?: string) => void;
}> = ({ newChatroom, chatHistory, setChatroom }) => {
  const [rooms, setRooms] = useState<string[]>();

  // generates chatroom list from chat history
  useEffect(() => {
    let tmpRooms = rooms ? [...rooms] : [];
    chatHistory.forEach((message) => {
      if (message.chatroom && !tmpRooms.includes(message.chatroom))
        tmpRooms.push(message.chatroom);
    });
    setRooms(tmpRooms);
  }, [chatHistory]);

  const form = (
    <Formik
      initialValues={{ name: "" }}
      validateOnBlur={false}
      validateOnChange={false}
      onSubmit={(values, actions) => {
        actions.setSubmitting(false);
        newChatroom(values.name);
        // clears form input field after msg sent
        actions.resetForm();
      }}
    >
      {({ handleChange, handleSubmit, isSubmitting, values }) => (
        <form onSubmit={handleSubmit}>
          <div className="lobby__input flex pl-2 mt-2">
            <input
              name="name"
              className=" w-10 p-2 my-1 text-base rounded shadow-md grow"
              onChange={handleChange}
              value={values.name}
              placeholder="New Chatroom"
            />
            {/* submit message on button press */}
            <button
              type="submit"
              className="lobby__add w-fit mt-1 text-md border-2 p-2 bg-slate-300 rounded-xl"
            >
              Add
            </button>
          </div>
        </form>
      )}
    </Formik>
  );

  return (
    <div className="Lobby flex flex-col h-full tracking-wide">
      <div className="Lobby__title  w-full h-full bg-sky-300 p-6 md:pl-10 font-bold">
        <h2 className="md:text-2xl sm:text-2xl flex-1">Chat Lobbies</h2>
      </div>
      <div className="Lobby__room h-full md:pl-7 pl-3 py-2 uppercase text-sm md:text-xl tracking-wide ">
        {form}
        {rooms?.map((room) => (
          <p key={room} onClick={() => setChatroom(room)}>
            {room}
          </p>
        ))}
      </div>
    </div>
  );
};

export default Lobby;
