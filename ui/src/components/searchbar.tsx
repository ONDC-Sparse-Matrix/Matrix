import { Input } from "./ui/input";
import { Search } from "lucide-react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useSessionStore } from "@/lib/session-store";
import { generateRandomId } from "@/lib/utils";

interface SearchBarProps {
  pincode?: string;
}

export function SearchBar(props: SearchBarProps) {
  const [searchedPincode, setSearchedPincode] = useState("");
  const { sessionId, setSessionId } = useSessionStore();
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (sessionId) {
      navigate(`/search?session=${sessionId}&pincode=${searchedPincode}`);
    } else {
      const newSessionId = generateRandomId(5);
      navigate(`/search?session=${newSessionId}&pincode=${searchedPincode}`);
      setSessionId(newSessionId);
    }
  };

  return (
    <>
      <div className="relative w-full mb-10">
        <div className="max-w-xl mx-auto">
          <div className="absolute inset-y-0 flex items-center ps-3 pointer-events-none">
            <Search className="text-sm text-gray-400 dark:text-gray-600" />
          </div>
          <form onSubmit={handleSubmit}>
            <Input
              placeholder="Enter your pincode"
              id="pincode"
              type="number"
              required
              onChange={(e) => setSearchedPincode(e.target.value)}
              defaultValue={props.pincode ? props.pincode : ""}
              className="bg-sky-50 dark:bg-gray-800 transition-all duration-300 p-6 ps-12 text-md rounded-full"
            />
          </form>
        </div>
      </div>
    </>
  );
}
