import "./App.css";
import Home from "@/pages/home";
import Search from "@/pages/search";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

function App() {
  return (
    <>
      <Router>
        <Routes>
          <Route>
            <Route path="/" element={<Home />} />
            <Route path="/search" element={<Search />} />
          </Route>
        </Routes>
      </Router>
    </>
  );
}

export default App;
