import "./App.css";
import Home from "@/pages/home";
import Search from "@/pages/search";
import AddMerchant from "./pages/add-merchant";
import Analytics from "./pages/analytics";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route>
            <Route path="/" element={<Home />} />
            <Route path="/search" element={<Search />} />
            <Route path="/add-merchant" element={<AddMerchant />} />
            <Route path="/analytics" element={<Analytics />} />
          </Route>
        </Routes>
      </Router>
    </>
  );
}

export default App;
