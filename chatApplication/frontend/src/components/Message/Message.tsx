import axios from "axios";
import { FC, useEffect, useState } from "react";
import { useIP } from "../../App";

type MessageMetadata = {
  country: string;
  IP: string; // query
  isUserMsg: boolean;
};

type MessageProps = {
  message: string;
};

//this component is used to display message in chat window
const Message: FC<MessageProps> = ({ message }) => {
  const IP = useIP();
  const [metadata, setMetaData] = useState<MessageMetadata | null>();
  useEffect(() => {
    getIpInfo();
  }, []);

  const getIpInfo = () => {
    axios
      .get("http://ip-api.com/json")
      .then((response) => {
        if (response.status === 200) {
          setMetaData({
            country: response.data.country,
            IP: response.data.query,
            isUserMsg: IP === response.data.query,
          });
        } else {
          setMetaData(null);
        }
      })
      .catch((err) => {
        console.log("get ip error", err);
        setMetaData(null);
      });
  };

  var today = new Date();
  const currentTime =
    today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();

  if (metadata === undefined) return <div>loading...</div>;
  if (metadata === null) return <div>error loading this message</div>;
  const userDetails = metadata.isUserMsg
    ? "Me"
    : `User: ${metadata.IP}, ${metadata.country}`;
  return (
    <>
      <div className="flex flex-col items-start p-2 ml-6">
        <div className="text-sm">
          <p>Time stamp:{currentTime}</p>
          <h1>{userDetails}</h1>
        </div>
        <div className="Message block bg-sky-100 m-1 p-2 rounded-sm shadow-sm clear-both w-fit font-light">
          {message}
        </div>
      </div>
    </>
  );
};

export default Message;
